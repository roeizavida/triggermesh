/*
Copyright 2021 TriggerMesh Inc.

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

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAWSLambdaTargets implements AWSLambdaTargetInterface
type FakeAWSLambdaTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var awslambdatargetsResource = schema.GroupVersionResource{Group: "targets.triggermesh.io", Version: "v1alpha1", Resource: "awslambdatargets"}

var awslambdatargetsKind = schema.GroupVersionKind{Group: "targets.triggermesh.io", Version: "v1alpha1", Kind: "AWSLambdaTarget"}

// Get takes name of the aWSLambdaTarget, and returns the corresponding aWSLambdaTarget object, and an error if there is any.
func (c *FakeAWSLambdaTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSLambdaTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslambdatargetsResource, c.ns, name), &v1alpha1.AWSLambdaTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSLambdaTarget), err
}

// List takes label and field selectors, and returns the list of AWSLambdaTargets that match those selectors.
func (c *FakeAWSLambdaTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSLambdaTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslambdatargetsResource, awslambdatargetsKind, c.ns, opts), &v1alpha1.AWSLambdaTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSLambdaTargetList{ListMeta: obj.(*v1alpha1.AWSLambdaTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSLambdaTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSLambdaTargets.
func (c *FakeAWSLambdaTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslambdatargetsResource, c.ns, opts))

}

// Create takes the representation of a aWSLambdaTarget and creates it.  Returns the server's representation of the aWSLambdaTarget, and an error, if there is any.
func (c *FakeAWSLambdaTargets) Create(ctx context.Context, aWSLambdaTarget *v1alpha1.AWSLambdaTarget, opts v1.CreateOptions) (result *v1alpha1.AWSLambdaTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslambdatargetsResource, c.ns, aWSLambdaTarget), &v1alpha1.AWSLambdaTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSLambdaTarget), err
}

// Update takes the representation of a aWSLambdaTarget and updates it. Returns the server's representation of the aWSLambdaTarget, and an error, if there is any.
func (c *FakeAWSLambdaTargets) Update(ctx context.Context, aWSLambdaTarget *v1alpha1.AWSLambdaTarget, opts v1.UpdateOptions) (result *v1alpha1.AWSLambdaTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslambdatargetsResource, c.ns, aWSLambdaTarget), &v1alpha1.AWSLambdaTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSLambdaTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSLambdaTargets) UpdateStatus(ctx context.Context, aWSLambdaTarget *v1alpha1.AWSLambdaTarget, opts v1.UpdateOptions) (*v1alpha1.AWSLambdaTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awslambdatargetsResource, "status", c.ns, aWSLambdaTarget), &v1alpha1.AWSLambdaTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSLambdaTarget), err
}

// Delete takes name of the aWSLambdaTarget and deletes it. Returns an error if one occurs.
func (c *FakeAWSLambdaTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslambdatargetsResource, c.ns, name), &v1alpha1.AWSLambdaTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSLambdaTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslambdatargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSLambdaTargetList{})
	return err
}

// Patch applies the patch and returns the patched aWSLambdaTarget.
func (c *FakeAWSLambdaTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSLambdaTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslambdatargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSLambdaTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSLambdaTarget), err
}
