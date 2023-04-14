/*
 Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "huawei-csi-driver/client/apis/xuanwu/v1"
	scheme "huawei-csi-driver/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StorageBackendClaimsGetter has a method to return a StorageBackendClaimInterface.
// A group's client should implement this interface.
type StorageBackendClaimsGetter interface {
	StorageBackendClaims(namespace string) StorageBackendClaimInterface
}

// StorageBackendClaimInterface has methods to work with StorageBackendClaim resources.
type StorageBackendClaimInterface interface {
	Create(ctx context.Context, storageBackendClaim *v1.StorageBackendClaim, opts metav1.CreateOptions) (*v1.StorageBackendClaim, error)
	Update(ctx context.Context, storageBackendClaim *v1.StorageBackendClaim, opts metav1.UpdateOptions) (*v1.StorageBackendClaim, error)
	UpdateStatus(ctx context.Context, storageBackendClaim *v1.StorageBackendClaim, opts metav1.UpdateOptions) (*v1.StorageBackendClaim, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.StorageBackendClaim, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.StorageBackendClaimList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.StorageBackendClaim, err error)
	StorageBackendClaimExpansion
}

// storageBackendClaims implements StorageBackendClaimInterface
type storageBackendClaims struct {
	client rest.Interface
	ns     string
}

// newStorageBackendClaims returns a StorageBackendClaims
func newStorageBackendClaims(c *XuanwuV1Client, namespace string) *storageBackendClaims {
	return &storageBackendClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storageBackendClaim, and returns the corresponding storageBackendClaim object, and an error if there is any.
func (c *storageBackendClaims) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.StorageBackendClaim, err error) {
	result = &v1.StorageBackendClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageBackendClaims that match those selectors.
func (c *storageBackendClaims) List(ctx context.Context, opts metav1.ListOptions) (result *v1.StorageBackendClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.StorageBackendClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageBackendClaims.
func (c *storageBackendClaims) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a storageBackendClaim and creates it.  Returns the server's representation of the storageBackendClaim, and an error, if there is any.
func (c *storageBackendClaims) Create(ctx context.Context, storageBackendClaim *v1.StorageBackendClaim, opts metav1.CreateOptions) (result *v1.StorageBackendClaim, err error) {
	result = &v1.StorageBackendClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageBackendClaim).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a storageBackendClaim and updates it. Returns the server's representation of the storageBackendClaim, and an error, if there is any.
func (c *storageBackendClaims) Update(ctx context.Context, storageBackendClaim *v1.StorageBackendClaim, opts metav1.UpdateOptions) (result *v1.StorageBackendClaim, err error) {
	result = &v1.StorageBackendClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		Name(storageBackendClaim.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageBackendClaim).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *storageBackendClaims) UpdateStatus(ctx context.Context, storageBackendClaim *v1.StorageBackendClaim, opts metav1.UpdateOptions) (result *v1.StorageBackendClaim, err error) {
	result = &v1.StorageBackendClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		Name(storageBackendClaim.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageBackendClaim).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the storageBackendClaim and deletes it. Returns an error if one occurs.
func (c *storageBackendClaims) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageBackendClaims) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagebackendclaims").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched storageBackendClaim.
func (c *storageBackendClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.StorageBackendClaim, err error) {
	result = &v1.StorageBackendClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storagebackendclaims").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
