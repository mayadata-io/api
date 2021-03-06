/*
Copyright 2020 The OpenEBS Authors

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
	"time"

	v1alpha1 "github.com/openebs/api/pkg/apis/openebs.io/v1alpha1"
	scheme "github.com/openebs/api/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MigrationTasksGetter has a method to return a MigrationTaskInterface.
// A group's client should implement this interface.
type MigrationTasksGetter interface {
	MigrationTasks(namespace string) MigrationTaskInterface
}

// MigrationTaskInterface has methods to work with MigrationTask resources.
type MigrationTaskInterface interface {
	Create(*v1alpha1.MigrationTask) (*v1alpha1.MigrationTask, error)
	Update(*v1alpha1.MigrationTask) (*v1alpha1.MigrationTask, error)
	UpdateStatus(*v1alpha1.MigrationTask) (*v1alpha1.MigrationTask, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MigrationTask, error)
	List(opts v1.ListOptions) (*v1alpha1.MigrationTaskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MigrationTask, err error)
	MigrationTaskExpansion
}

// migrationTasks implements MigrationTaskInterface
type migrationTasks struct {
	client rest.Interface
	ns     string
}

// newMigrationTasks returns a MigrationTasks
func newMigrationTasks(c *OpenebsV1alpha1Client, namespace string) *migrationTasks {
	return &migrationTasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the migrationTask, and returns the corresponding migrationTask object, and an error if there is any.
func (c *migrationTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.MigrationTask, err error) {
	result = &v1alpha1.MigrationTask{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("migrationtasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MigrationTasks that match those selectors.
func (c *migrationTasks) List(opts v1.ListOptions) (result *v1alpha1.MigrationTaskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MigrationTaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("migrationtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested migrationTasks.
func (c *migrationTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("migrationtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a migrationTask and creates it.  Returns the server's representation of the migrationTask, and an error, if there is any.
func (c *migrationTasks) Create(migrationTask *v1alpha1.MigrationTask) (result *v1alpha1.MigrationTask, err error) {
	result = &v1alpha1.MigrationTask{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("migrationtasks").
		Body(migrationTask).
		Do().
		Into(result)
	return
}

// Update takes the representation of a migrationTask and updates it. Returns the server's representation of the migrationTask, and an error, if there is any.
func (c *migrationTasks) Update(migrationTask *v1alpha1.MigrationTask) (result *v1alpha1.MigrationTask, err error) {
	result = &v1alpha1.MigrationTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("migrationtasks").
		Name(migrationTask.Name).
		Body(migrationTask).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *migrationTasks) UpdateStatus(migrationTask *v1alpha1.MigrationTask) (result *v1alpha1.MigrationTask, err error) {
	result = &v1alpha1.MigrationTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("migrationtasks").
		Name(migrationTask.Name).
		SubResource("status").
		Body(migrationTask).
		Do().
		Into(result)
	return
}

// Delete takes name of the migrationTask and deletes it. Returns an error if one occurs.
func (c *migrationTasks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("migrationtasks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *migrationTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("migrationtasks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched migrationTask.
func (c *migrationTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MigrationTask, err error) {
	result = &v1alpha1.MigrationTask{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("migrationtasks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
