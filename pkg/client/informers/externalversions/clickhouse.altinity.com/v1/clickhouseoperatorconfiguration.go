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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	clickhousealtinitycomv1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	versioned "github.com/altinity/clickhouse-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/altinity/clickhouse-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/altinity/clickhouse-operator/pkg/client/listers/clickhouse.altinity.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClickHouseOperatorConfigurationInformer provides access to a shared informer and lister for
// ClickHouseOperatorConfigurations.
type ClickHouseOperatorConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClickHouseOperatorConfigurationLister
}

type clickHouseOperatorConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClickHouseOperatorConfigurationInformer constructs a new informer for ClickHouseOperatorConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClickHouseOperatorConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClickHouseOperatorConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClickHouseOperatorConfigurationInformer constructs a new informer for ClickHouseOperatorConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClickHouseOperatorConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClickhouseV1().ClickHouseOperatorConfigurations(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClickhouseV1().ClickHouseOperatorConfigurations(namespace).Watch(options)
			},
		},
		&clickhousealtinitycomv1.ClickHouseOperatorConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *clickHouseOperatorConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClickHouseOperatorConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clickHouseOperatorConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clickhousealtinitycomv1.ClickHouseOperatorConfiguration{}, f.defaultInformer)
}

func (f *clickHouseOperatorConfigurationInformer) Lister() v1.ClickHouseOperatorConfigurationLister {
	return v1.NewClickHouseOperatorConfigurationLister(f.Informer().GetIndexer())
}