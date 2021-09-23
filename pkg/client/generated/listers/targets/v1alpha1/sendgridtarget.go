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

// SendGridTargetLister helps list SendGridTargets.
// All objects returned here must be treated as read-only.
type SendGridTargetLister interface {
	// List lists all SendGridTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SendGridTarget, err error)
	// SendGridTargets returns an object that can list and get SendGridTargets.
	SendGridTargets(namespace string) SendGridTargetNamespaceLister
	SendGridTargetListerExpansion
}

// sendGridTargetLister implements the SendGridTargetLister interface.
type sendGridTargetLister struct {
	indexer cache.Indexer
}

// NewSendGridTargetLister returns a new SendGridTargetLister.
func NewSendGridTargetLister(indexer cache.Indexer) SendGridTargetLister {
	return &sendGridTargetLister{indexer: indexer}
}

// List lists all SendGridTargets in the indexer.
func (s *sendGridTargetLister) List(selector labels.Selector) (ret []*v1alpha1.SendGridTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SendGridTarget))
	})
	return ret, err
}

// SendGridTargets returns an object that can list and get SendGridTargets.
func (s *sendGridTargetLister) SendGridTargets(namespace string) SendGridTargetNamespaceLister {
	return sendGridTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SendGridTargetNamespaceLister helps list and get SendGridTargets.
// All objects returned here must be treated as read-only.
type SendGridTargetNamespaceLister interface {
	// List lists all SendGridTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SendGridTarget, err error)
	// Get retrieves the SendGridTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SendGridTarget, error)
	SendGridTargetNamespaceListerExpansion
}

// sendGridTargetNamespaceLister implements the SendGridTargetNamespaceLister
// interface.
type sendGridTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SendGridTargets in the indexer for a given namespace.
func (s sendGridTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SendGridTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SendGridTarget))
	})
	return ret, err
}

// Get retrieves the SendGridTarget from the indexer for a given namespace and name.
func (s sendGridTargetNamespaceLister) Get(name string) (*v1alpha1.SendGridTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sendgridtarget"), name)
	}
	return obj.(*v1alpha1.SendGridTarget), nil
}
