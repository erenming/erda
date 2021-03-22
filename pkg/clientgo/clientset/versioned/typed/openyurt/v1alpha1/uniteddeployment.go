/*
Copyright 2020 The OpenYurt Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "github.com/erda-project/erda/pkg/clientgo/apis/openyurt/v1alpha1"
	scheme "github.com/erda-project/erda/pkg/clientgo/clientset/versioned/scheme"
)

// UnitedDeploymentsGetter has a method to return a UnitedDeploymentInterface.
// A group's client should implement this interface.
type UnitedDeploymentsGetter interface {
	UnitedDeployments(namespace string) UnitedDeploymentInterface
}

// UnitedDeploymentInterface has methods to work with UnitedDeployment resources.
type UnitedDeploymentInterface interface {
	Create(context.Context, *v1alpha1.UnitedDeployment) (*v1alpha1.UnitedDeployment, error)
	Update(context.Context, *v1alpha1.UnitedDeployment) (*v1alpha1.UnitedDeployment, error)
	UpdateStatus(context.Context, *v1alpha1.UnitedDeployment) (*v1alpha1.UnitedDeployment, error)
	Delete(ctx context.Context, name string, options *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(ctx context.Context, name string, options v1.GetOptions) (*v1alpha1.UnitedDeployment, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UnitedDeploymentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UnitedDeployment, err error)
	UnitedDeploymentExpansion
}

// unitedDeployments implements UnitedDeploymentInterface
type unitedDeployments struct {
	client rest.Interface
	ns     string
}

// newUnitedDeployments returns a UnitedDeployments
func newUnitedDeployments(c *AppsV1alpha1Client, namespace string) *unitedDeployments {
	return &unitedDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the unitedDeployment, and returns the corresponding unitedDeployment object, and an error if there is any.
func (c *unitedDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UnitedDeployments that match those selectors.
func (c *unitedDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UnitedDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UnitedDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("uniteddeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested unitedDeployments.
func (c *unitedDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("uniteddeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a unitedDeployment and creates it.  Returns the server's representation of the unitedDeployment, and an error, if there is any.
func (c *unitedDeployments) Create(ctx context.Context, unitedDeployment *v1alpha1.UnitedDeployment) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Body(unitedDeployment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a unitedDeployment and updates it. Returns the server's representation of the unitedDeployment, and an error, if there is any.
func (c *unitedDeployments) Update(ctx context.Context, unitedDeployment *v1alpha1.UnitedDeployment) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(unitedDeployment.Name).
		Body(unitedDeployment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *unitedDeployments) UpdateStatus(ctx context.Context, unitedDeployment *v1alpha1.UnitedDeployment) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(unitedDeployment.Name).
		SubResource("status").
		Body(unitedDeployment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the unitedDeployment and deletes it. Returns an error if one occurs.
func (c *unitedDeployments) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("uniteddeployments").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *unitedDeployments) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("uniteddeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched unitedDeployment.
func (c *unitedDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UnitedDeployment, err error) {
	result = &v1alpha1.UnitedDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("uniteddeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
