// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/contiv/vpp/plugins/crd/pkg/apis/contivtelemetry/v1"
	scheme "github.com/contiv/vpp/plugins/crd/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ContivTelemetryReportsGetter has a method to return a ContivTelemetryReportInterface.
// A group's client should implement this interface.
type ContivTelemetryReportsGetter interface {
	ContivTelemetryReports(namespace string) ContivTelemetryReportInterface
}

// ContivTelemetryReportInterface has methods to work with ContivTelemetryReport resources.
type ContivTelemetryReportInterface interface {
	Create(*v1.ContivTelemetryReport) (*v1.ContivTelemetryReport, error)
	Update(*v1.ContivTelemetryReport) (*v1.ContivTelemetryReport, error)
	UpdateStatus(*v1.ContivTelemetryReport) (*v1.ContivTelemetryReport, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ContivTelemetryReport, error)
	List(opts metav1.ListOptions) (*v1.ContivTelemetryReportList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ContivTelemetryReport, err error)
	ContivTelemetryReportExpansion
}

// contivTelemetryReports implements ContivTelemetryReportInterface
type contivTelemetryReports struct {
	client rest.Interface
	ns     string
}

// newContivTelemetryReports returns a ContivTelemetryReports
func newContivTelemetryReports(c *ContivtelemetryV1Client, namespace string) *contivTelemetryReports {
	return &contivTelemetryReports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the contivTelemetryReport, and returns the corresponding contivTelemetryReport object, and an error if there is any.
func (c *contivTelemetryReports) Get(name string, options metav1.GetOptions) (result *v1.ContivTelemetryReport, err error) {
	result = &v1.ContivTelemetryReport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ContivTelemetryReports that match those selectors.
func (c *contivTelemetryReports) List(opts metav1.ListOptions) (result *v1.ContivTelemetryReportList, err error) {
	result = &v1.ContivTelemetryReportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested contivTelemetryReports.
func (c *contivTelemetryReports) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a contivTelemetryReport and creates it.  Returns the server's representation of the contivTelemetryReport, and an error, if there is any.
func (c *contivTelemetryReports) Create(contivTelemetryReport *v1.ContivTelemetryReport) (result *v1.ContivTelemetryReport, err error) {
	result = &v1.ContivTelemetryReport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		Body(contivTelemetryReport).
		Do().
		Into(result)
	return
}

// Update takes the representation of a contivTelemetryReport and updates it. Returns the server's representation of the contivTelemetryReport, and an error, if there is any.
func (c *contivTelemetryReports) Update(contivTelemetryReport *v1.ContivTelemetryReport) (result *v1.ContivTelemetryReport, err error) {
	result = &v1.ContivTelemetryReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		Name(contivTelemetryReport.Name).
		Body(contivTelemetryReport).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *contivTelemetryReports) UpdateStatus(contivTelemetryReport *v1.ContivTelemetryReport) (result *v1.ContivTelemetryReport, err error) {
	result = &v1.ContivTelemetryReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		Name(contivTelemetryReport.Name).
		SubResource("status").
		Body(contivTelemetryReport).
		Do().
		Into(result)
	return
}

// Delete takes name of the contivTelemetryReport and deletes it. Returns an error if one occurs.
func (c *contivTelemetryReports) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *contivTelemetryReports) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched contivTelemetryReport.
func (c *contivTelemetryReports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ContivTelemetryReport, err error) {
	result = &v1.ContivTelemetryReport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("contivtelemetryreports").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}