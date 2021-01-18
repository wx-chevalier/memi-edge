/*
Copyright The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/baetyl/baetyl-cloud/v2/plugin/kube/apis/cloud/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeReports implements NodeReportInterface
type FakeNodeReports struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var nodereportsResource = schema.GroupVersionResource{Group: "cloud.baetyl.io", Version: "v1alpha1", Resource: "nodereports"}

var nodereportsKind = schema.GroupVersionKind{Group: "cloud.baetyl.io", Version: "v1alpha1", Kind: "NodeReport"}

// Get takes name of the nodeReport, and returns the corresponding nodeReport object, and an error if there is any.
func (c *FakeNodeReports) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodereportsResource, c.ns, name), &v1alpha1.NodeReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeReport), err
}

// List takes label and field selectors, and returns the list of NodeReports that match those selectors.
func (c *FakeNodeReports) List(opts v1.ListOptions) (result *v1alpha1.NodeReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodereportsResource, nodereportsKind, c.ns, opts), &v1alpha1.NodeReportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeReportList{ListMeta: obj.(*v1alpha1.NodeReportList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeReports.
func (c *FakeNodeReports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodereportsResource, c.ns, opts))

}

// Create takes the representation of a nodeReport and creates it.  Returns the server's representation of the nodeReport, and an error, if there is any.
func (c *FakeNodeReports) Create(nodeReport *v1alpha1.NodeReport) (result *v1alpha1.NodeReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodereportsResource, c.ns, nodeReport), &v1alpha1.NodeReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeReport), err
}

// Update takes the representation of a nodeReport and updates it. Returns the server's representation of the nodeReport, and an error, if there is any.
func (c *FakeNodeReports) Update(nodeReport *v1alpha1.NodeReport) (result *v1alpha1.NodeReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodereportsResource, c.ns, nodeReport), &v1alpha1.NodeReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeReport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeReports) UpdateStatus(nodeReport *v1alpha1.NodeReport) (*v1alpha1.NodeReport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodereportsResource, "status", c.ns, nodeReport), &v1alpha1.NodeReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeReport), err
}

// Delete takes name of the nodeReport and deletes it. Returns an error if one occurs.
func (c *FakeNodeReports) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodereportsResource, c.ns, name), &v1alpha1.NodeReport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeReports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodereportsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeReportList{})
	return err
}

// Patch applies the patch and returns the patched nodeReport.
func (c *FakeNodeReports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodereportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NodeReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeReport), err
}