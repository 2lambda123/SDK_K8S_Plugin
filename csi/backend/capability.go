/*
 *  Copyright (c) Huawei Technologies Co., Ltd. 2020-2023. All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package backend

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"sync"

	"huawei-csi-driver/csi/app"
	pkgUtils "huawei-csi-driver/pkg/utils"
	"huawei-csi-driver/utils/log"
)

func updatePoolCapabilitiesByBackend(backend *Backend, backendCapabilities, poolCapabilities map[string]interface{}) {
	for _, pool := range backend.Pools {
		for k, v := range backendCapabilities {
			if cur, exist := pool.Capabilities[k]; !exist || cur != v {
				log.Infof("Update backend capability [%s] of pool [%s] of backend [%s] from %v to %v",
					k, pool.Name, pool.Parent, cur, v)
				pool.Capabilities[k] = v
			}
		}

		// The storage pool capability does not need to be updated in the DTree scenario.
		if backend.Storage == "oceanstor-dtree" {
			continue
		}

		capabilities, exist := poolCapabilities[pool.Name].(map[string]interface{})
		if !exist {
			log.Warningf("Pool %s of backend %s does not exist, set it unavailable", pool.Name, pool.Parent)
			pool.Capabilities["FreeCapacity"] = 0
			continue
		}

		for k, v := range capabilities {
			if cur, exist := pool.Capabilities[k]; !exist || !reflect.DeepEqual(cur, v) {
				log.FilteredLog(context.TODO(), false, k == "FreeCapacity",
					fmt.Sprintf("Update pool capability [%s] of pool [%s] of backend [%s] from %v to %v",
						k, pool.Name, pool.Parent, cur, v))
				pool.Capabilities[k] = v
			}
		}
	}
}

func updateBackendCapabilities(backend *Backend, sync bool) error {
	// If sbct is offline, delete the backend from the csiBackends.
	backendID := pkgUtils.MakeMetaWithNamespace(app.GetGlobalConfig().Namespace, backend.Name)
	online, err := pkgUtils.GetSBCTOnlineStatusByClaim(context.TODO(), backendID)
	if !online {
		RemoveOneBackend(context.TODO(), backend.Name)
		log.Infof("SBCT: [%s] online status is false, RemoveOneBackend: [%s]", backendID, backend.Name)
		return nil
	}
	backendCapabilities, _, err := backend.Plugin.UpdateBackendCapabilities()
	if err != nil {
		log.Errorf("Cannot update backend %s capabilities: %v", backend.Name, err)
		return err
	}

	var poolNames []string
	for _, pool := range backend.Pools {
		poolNames = append(poolNames, pool.Name)
	}
	poolCapabilities, err := backend.Plugin.UpdatePoolCapabilities(poolNames)
	if err != nil {
		log.Errorf("Cannot update pool capabilities of backend %s: %v", backend.Name, err)
		return err
	}

	if sync && len(poolCapabilities) < len(poolNames) {
		msg := fmt.Sprintf("There're pools not available for backend %s", backend.Name)
		log.Errorln(msg)
		return errors.New(msg)
	}

	updatePoolCapabilitiesByBackend(backend, backendCapabilities, poolCapabilities)

	return nil
}

func SyncUpdateCapabilities() error {
	log.Debugln("main SyncUpdateCapabilities")
	for _, backend := range csiBackends {
		err := updateBackendCapabilities(backend, true)
		if err != nil {
			return err
		}

		backend.Available = true
	}

	return nil
}

func AsyncUpdateCapabilities() {
	log.Debugln("main AsyncUpdateCapabilities")
	var wait sync.WaitGroup

	for _, backend := range csiBackends {
		wait.Add(1)

		go func(b *Backend) {
			defer func() {
				wait.Done()

				if r := recover(); r != nil {
					log.Errorf("Runtime error caught in loop routine: %v", r)
					log.Errorf("%s", debug.Stack())
				}

				log.Flush()
			}()

			err := updateBackendCapabilities(b, false)
			if err != nil {
				log.Warningf("update backend %s capabilities failed, error: %v", b.Name, err)
				b.Available = false
			} else {
				b.Available = true
			}
		}(backend)
	}

	wait.Wait()
}

// LogoutBackend is to logout all storage backend
func LogoutBackend() {
	for _, backend := range csiBackends {
		log.Infof("Start to logout the backend %s", backend.Name)
		backend.Plugin.Logout(context.Background())
		backend.Available = false
	}
}
