/*
Copyright 2023.

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
	"context"

	v1 "github.com/zhaoxin-BF/cloud-operator/apis/example.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBars implements BarInterface
type FakeBars struct {
	Fake *FakeExampleV1
	ns   string
}

var barsResource = v1.SchemeGroupVersion.WithResource("bars")

var barsKind = v1.SchemeGroupVersion.WithKind("Bar")

// Get takes name of the bar, and returns the corresponding bar object, and an error if there is any.
func (c *FakeBars) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(barsResource, c.ns, name), &v1.Bar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Bar), err
}

// List takes label and field selectors, and returns the list of Bars that match those selectors.
func (c *FakeBars) List(ctx context.Context, opts metav1.ListOptions) (result *v1.BarList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(barsResource, barsKind, c.ns, opts), &v1.BarList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.BarList{ListMeta: obj.(*v1.BarList).ListMeta}
	for _, item := range obj.(*v1.BarList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bars.
func (c *FakeBars) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(barsResource, c.ns, opts))

}

// Create takes the representation of a bar and creates it.  Returns the server's representation of the bar, and an error, if there is any.
func (c *FakeBars) Create(ctx context.Context, bar *v1.Bar, opts metav1.CreateOptions) (result *v1.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(barsResource, c.ns, bar), &v1.Bar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Bar), err
}

// Update takes the representation of a bar and updates it. Returns the server's representation of the bar, and an error, if there is any.
func (c *FakeBars) Update(ctx context.Context, bar *v1.Bar, opts metav1.UpdateOptions) (result *v1.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(barsResource, c.ns, bar), &v1.Bar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Bar), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBars) UpdateStatus(ctx context.Context, bar *v1.Bar, opts metav1.UpdateOptions) (*v1.Bar, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(barsResource, "status", c.ns, bar), &v1.Bar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Bar), err
}

// Delete takes name of the bar and deletes it. Returns an error if one occurs.
func (c *FakeBars) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(barsResource, c.ns, name, opts), &v1.Bar{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBars) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(barsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.BarList{})
	return err
}

// Patch applies the patch and returns the patched bar.
func (c *FakeBars) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(barsResource, c.ns, name, pt, data, subresources...), &v1.Bar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Bar), err
}
