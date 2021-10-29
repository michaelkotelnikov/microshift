// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	networkv1 "github.com/openshift/api/network/v1"
	versioned "github.com/openshift/client-go/network/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/network/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/network/listers/network/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EgressNetworkPolicyInformer provides access to a shared informer and lister for
// EgressNetworkPolicies.
type EgressNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EgressNetworkPolicyLister
}

type egressNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEgressNetworkPolicyInformer constructs a new informer for EgressNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEgressNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEgressNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEgressNetworkPolicyInformer constructs a new informer for EgressNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEgressNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().EgressNetworkPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().EgressNetworkPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&networkv1.EgressNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *egressNetworkPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEgressNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *egressNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1.EgressNetworkPolicy{}, f.defaultInformer)
}

func (f *egressNetworkPolicyInformer) Lister() v1.EgressNetworkPolicyLister {
	return v1.NewEgressNetworkPolicyLister(f.Informer().GetIndexer())
}