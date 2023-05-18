// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContainerRuntimeConfigs implements ContainerRuntimeConfigInterface
type FakeContainerRuntimeConfigs struct {
	Fake *FakeMachineconfigurationV1
}

var containerruntimeconfigsResource = v1.SchemeGroupVersion.WithResource("containerruntimeconfigs")

var containerruntimeconfigsKind = v1.SchemeGroupVersion.WithKind("ContainerRuntimeConfig")

// Get takes name of the containerRuntimeConfig, and returns the corresponding containerRuntimeConfig object, and an error if there is any.
func (c *FakeContainerRuntimeConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(containerruntimeconfigsResource, name), &v1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ContainerRuntimeConfig), err
}

// List takes label and field selectors, and returns the list of ContainerRuntimeConfigs that match those selectors.
func (c *FakeContainerRuntimeConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ContainerRuntimeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(containerruntimeconfigsResource, containerruntimeconfigsKind, opts), &v1.ContainerRuntimeConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ContainerRuntimeConfigList{ListMeta: obj.(*v1.ContainerRuntimeConfigList).ListMeta}
	for _, item := range obj.(*v1.ContainerRuntimeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerRuntimeConfigs.
func (c *FakeContainerRuntimeConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(containerruntimeconfigsResource, opts))
}

// Create takes the representation of a containerRuntimeConfig and creates it.  Returns the server's representation of the containerRuntimeConfig, and an error, if there is any.
func (c *FakeContainerRuntimeConfigs) Create(ctx context.Context, containerRuntimeConfig *v1.ContainerRuntimeConfig, opts metav1.CreateOptions) (result *v1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(containerruntimeconfigsResource, containerRuntimeConfig), &v1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ContainerRuntimeConfig), err
}

// Update takes the representation of a containerRuntimeConfig and updates it. Returns the server's representation of the containerRuntimeConfig, and an error, if there is any.
func (c *FakeContainerRuntimeConfigs) Update(ctx context.Context, containerRuntimeConfig *v1.ContainerRuntimeConfig, opts metav1.UpdateOptions) (result *v1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(containerruntimeconfigsResource, containerRuntimeConfig), &v1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ContainerRuntimeConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerRuntimeConfigs) UpdateStatus(ctx context.Context, containerRuntimeConfig *v1.ContainerRuntimeConfig, opts metav1.UpdateOptions) (*v1.ContainerRuntimeConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(containerruntimeconfigsResource, "status", containerRuntimeConfig), &v1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ContainerRuntimeConfig), err
}

// Delete takes name of the containerRuntimeConfig and deletes it. Returns an error if one occurs.
func (c *FakeContainerRuntimeConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(containerruntimeconfigsResource, name, opts), &v1.ContainerRuntimeConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerRuntimeConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(containerruntimeconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ContainerRuntimeConfigList{})
	return err
}

// Patch applies the patch and returns the patched containerRuntimeConfig.
func (c *FakeContainerRuntimeConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(containerruntimeconfigsResource, name, pt, data, subresources...), &v1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ContainerRuntimeConfig), err
}
