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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GoogleCloudFirestoreTargetLister helps list GoogleCloudFirestoreTargets.
// All objects returned here must be treated as read-only.
type GoogleCloudFirestoreTargetLister interface {
	// List lists all GoogleCloudFirestoreTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleCloudFirestoreTarget, err error)
	// GoogleCloudFirestoreTargets returns an object that can list and get GoogleCloudFirestoreTargets.
	GoogleCloudFirestoreTargets(namespace string) GoogleCloudFirestoreTargetNamespaceLister
	GoogleCloudFirestoreTargetListerExpansion
}

// googleCloudFirestoreTargetLister implements the GoogleCloudFirestoreTargetLister interface.
type googleCloudFirestoreTargetLister struct {
	indexer cache.Indexer
}

// NewGoogleCloudFirestoreTargetLister returns a new GoogleCloudFirestoreTargetLister.
func NewGoogleCloudFirestoreTargetLister(indexer cache.Indexer) GoogleCloudFirestoreTargetLister {
	return &googleCloudFirestoreTargetLister{indexer: indexer}
}

// List lists all GoogleCloudFirestoreTargets in the indexer.
func (s *googleCloudFirestoreTargetLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleCloudFirestoreTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleCloudFirestoreTarget))
	})
	return ret, err
}

// GoogleCloudFirestoreTargets returns an object that can list and get GoogleCloudFirestoreTargets.
func (s *googleCloudFirestoreTargetLister) GoogleCloudFirestoreTargets(namespace string) GoogleCloudFirestoreTargetNamespaceLister {
	return googleCloudFirestoreTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GoogleCloudFirestoreTargetNamespaceLister helps list and get GoogleCloudFirestoreTargets.
// All objects returned here must be treated as read-only.
type GoogleCloudFirestoreTargetNamespaceLister interface {
	// List lists all GoogleCloudFirestoreTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleCloudFirestoreTarget, err error)
	// Get retrieves the GoogleCloudFirestoreTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GoogleCloudFirestoreTarget, error)
	GoogleCloudFirestoreTargetNamespaceListerExpansion
}

// googleCloudFirestoreTargetNamespaceLister implements the GoogleCloudFirestoreTargetNamespaceLister
// interface.
type googleCloudFirestoreTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GoogleCloudFirestoreTargets in the indexer for a given namespace.
func (s googleCloudFirestoreTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleCloudFirestoreTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleCloudFirestoreTarget))
	})
	return ret, err
}

// Get retrieves the GoogleCloudFirestoreTarget from the indexer for a given namespace and name.
func (s googleCloudFirestoreTargetNamespaceLister) Get(name string) (*v1alpha1.GoogleCloudFirestoreTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googlecloudfirestoretarget"), name)
	}
	return obj.(*v1alpha1.GoogleCloudFirestoreTarget), nil
}
