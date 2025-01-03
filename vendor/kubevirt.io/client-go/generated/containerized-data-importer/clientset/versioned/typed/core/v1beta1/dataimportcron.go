/*
Copyright The KubeVirt Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "kubevirt.io/client-go/generated/containerized-data-importer/clientset/versioned/scheme"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// DataImportCronsGetter has a method to return a DataImportCronInterface.
// A group's client should implement this interface.
type DataImportCronsGetter interface {
	DataImportCrons(namespace string) DataImportCronInterface
}

// DataImportCronInterface has methods to work with DataImportCron resources.
type DataImportCronInterface interface {
	Create(ctx context.Context, dataImportCron *v1beta1.DataImportCron, opts v1.CreateOptions) (*v1beta1.DataImportCron, error)
	Update(ctx context.Context, dataImportCron *v1beta1.DataImportCron, opts v1.UpdateOptions) (*v1beta1.DataImportCron, error)
	UpdateStatus(ctx context.Context, dataImportCron *v1beta1.DataImportCron, opts v1.UpdateOptions) (*v1beta1.DataImportCron, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DataImportCron, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DataImportCronList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataImportCron, err error)
	DataImportCronExpansion
}

// dataImportCrons implements DataImportCronInterface
type dataImportCrons struct {
	client rest.Interface
	ns     string
}

// newDataImportCrons returns a DataImportCrons
func newDataImportCrons(c *CdiV1beta1Client, namespace string) *dataImportCrons {
	return &dataImportCrons{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataImportCron, and returns the corresponding dataImportCron object, and an error if there is any.
func (c *dataImportCrons) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataImportCron, err error) {
	result = &v1beta1.DataImportCron{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataimportcrons").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataImportCrons that match those selectors.
func (c *dataImportCrons) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataImportCronList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DataImportCronList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataimportcrons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataImportCrons.
func (c *dataImportCrons) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dataimportcrons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dataImportCron and creates it.  Returns the server's representation of the dataImportCron, and an error, if there is any.
func (c *dataImportCrons) Create(ctx context.Context, dataImportCron *v1beta1.DataImportCron, opts v1.CreateOptions) (result *v1beta1.DataImportCron, err error) {
	result = &v1beta1.DataImportCron{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dataimportcrons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataImportCron).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dataImportCron and updates it. Returns the server's representation of the dataImportCron, and an error, if there is any.
func (c *dataImportCrons) Update(ctx context.Context, dataImportCron *v1beta1.DataImportCron, opts v1.UpdateOptions) (result *v1beta1.DataImportCron, err error) {
	result = &v1beta1.DataImportCron{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dataimportcrons").
		Name(dataImportCron.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataImportCron).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dataImportCrons) UpdateStatus(ctx context.Context, dataImportCron *v1beta1.DataImportCron, opts v1.UpdateOptions) (result *v1beta1.DataImportCron, err error) {
	result = &v1beta1.DataImportCron{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dataimportcrons").
		Name(dataImportCron.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataImportCron).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dataImportCron and deletes it. Returns an error if one occurs.
func (c *dataImportCrons) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataimportcrons").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataImportCrons) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataimportcrons").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dataImportCron.
func (c *dataImportCrons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataImportCron, err error) {
	result = &v1beta1.DataImportCron{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dataimportcrons").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
