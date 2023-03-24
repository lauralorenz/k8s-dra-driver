/*
 * Copyright (c) 2023, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/NVIDIA/k8s-dra-driver/api/nvidia.com/resource/gpu/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMigDeviceClaimParameters implements MigDeviceClaimParametersInterface
type FakeMigDeviceClaimParameters struct {
	Fake *FakeGpuV1alpha1
	ns   string
}

var migdeviceclaimparametersResource = schema.GroupVersionResource{Group: "gpu.resource.nvidia.com", Version: "v1alpha1", Resource: "migdeviceclaimparameters"}

var migdeviceclaimparametersKind = schema.GroupVersionKind{Group: "gpu.resource.nvidia.com", Version: "v1alpha1", Kind: "MigDeviceClaimParameters"}

// Get takes name of the migDeviceClaimParameters, and returns the corresponding migDeviceClaimParameters object, and an error if there is any.
func (c *FakeMigDeviceClaimParameters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MigDeviceClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(migdeviceclaimparametersResource, c.ns, name), &v1alpha1.MigDeviceClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MigDeviceClaimParameters), err
}

// List takes label and field selectors, and returns the list of MigDeviceClaimParameters that match those selectors.
func (c *FakeMigDeviceClaimParameters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MigDeviceClaimParametersList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(migdeviceclaimparametersResource, migdeviceclaimparametersKind, c.ns, opts), &v1alpha1.MigDeviceClaimParametersList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MigDeviceClaimParametersList{ListMeta: obj.(*v1alpha1.MigDeviceClaimParametersList).ListMeta}
	for _, item := range obj.(*v1alpha1.MigDeviceClaimParametersList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested migDeviceClaimParameters.
func (c *FakeMigDeviceClaimParameters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(migdeviceclaimparametersResource, c.ns, opts))

}

// Create takes the representation of a migDeviceClaimParameters and creates it.  Returns the server's representation of the migDeviceClaimParameters, and an error, if there is any.
func (c *FakeMigDeviceClaimParameters) Create(ctx context.Context, migDeviceClaimParameters *v1alpha1.MigDeviceClaimParameters, opts v1.CreateOptions) (result *v1alpha1.MigDeviceClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(migdeviceclaimparametersResource, c.ns, migDeviceClaimParameters), &v1alpha1.MigDeviceClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MigDeviceClaimParameters), err
}

// Update takes the representation of a migDeviceClaimParameters and updates it. Returns the server's representation of the migDeviceClaimParameters, and an error, if there is any.
func (c *FakeMigDeviceClaimParameters) Update(ctx context.Context, migDeviceClaimParameters *v1alpha1.MigDeviceClaimParameters, opts v1.UpdateOptions) (result *v1alpha1.MigDeviceClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(migdeviceclaimparametersResource, c.ns, migDeviceClaimParameters), &v1alpha1.MigDeviceClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MigDeviceClaimParameters), err
}

// Delete takes name of the migDeviceClaimParameters and deletes it. Returns an error if one occurs.
func (c *FakeMigDeviceClaimParameters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(migdeviceclaimparametersResource, c.ns, name, opts), &v1alpha1.MigDeviceClaimParameters{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMigDeviceClaimParameters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(migdeviceclaimparametersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MigDeviceClaimParametersList{})
	return err
}

// Patch applies the patch and returns the patched migDeviceClaimParameters.
func (c *FakeMigDeviceClaimParameters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MigDeviceClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(migdeviceclaimparametersResource, c.ns, name, pt, data, subresources...), &v1alpha1.MigDeviceClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MigDeviceClaimParameters), err
}