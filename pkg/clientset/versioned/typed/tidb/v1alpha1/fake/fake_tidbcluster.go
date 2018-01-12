/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/gaocegege/kubetidb/pkg/apis/tidb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTiDBClusters implements TiDBClusterInterface
type FakeTiDBClusters struct {
	Fake *FakeKubetidbV1alpha1
	ns   string
}

var tidbclustersResource = schema.GroupVersionResource{Group: "kubetidb.gaocegege.com", Version: "v1alpha1", Resource: "tidbclusters"}

var tidbclustersKind = schema.GroupVersionKind{Group: "kubetidb.gaocegege.com", Version: "v1alpha1", Kind: "TiDBCluster"}

// Get takes name of the tiDBCluster, and returns the corresponding tiDBCluster object, and an error if there is any.
func (c *FakeTiDBClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.TiDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tidbclustersResource, c.ns, name), &v1alpha1.TiDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBCluster), err
}

// List takes label and field selectors, and returns the list of TiDBClusters that match those selectors.
func (c *FakeTiDBClusters) List(opts v1.ListOptions) (result *v1alpha1.TiDBClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tidbclustersResource, tidbclustersKind, c.ns, opts), &v1alpha1.TiDBClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TiDBClusterList{}
	for _, item := range obj.(*v1alpha1.TiDBClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tiDBClusters.
func (c *FakeTiDBClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tidbclustersResource, c.ns, opts))

}

// Create takes the representation of a tiDBCluster and creates it.  Returns the server's representation of the tiDBCluster, and an error, if there is any.
func (c *FakeTiDBClusters) Create(tiDBCluster *v1alpha1.TiDBCluster) (result *v1alpha1.TiDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tidbclustersResource, c.ns, tiDBCluster), &v1alpha1.TiDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBCluster), err
}

// Update takes the representation of a tiDBCluster and updates it. Returns the server's representation of the tiDBCluster, and an error, if there is any.
func (c *FakeTiDBClusters) Update(tiDBCluster *v1alpha1.TiDBCluster) (result *v1alpha1.TiDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tidbclustersResource, c.ns, tiDBCluster), &v1alpha1.TiDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBCluster), err
}

// Delete takes name of the tiDBCluster and deletes it. Returns an error if one occurs.
func (c *FakeTiDBClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tidbclustersResource, c.ns, name), &v1alpha1.TiDBCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTiDBClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tidbclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TiDBClusterList{})
	return err
}

// Patch applies the patch and returns the patched tiDBCluster.
func (c *FakeTiDBClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TiDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tidbclustersResource, c.ns, name, data, subresources...), &v1alpha1.TiDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBCluster), err
}