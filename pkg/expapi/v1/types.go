/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package v1

import "k8s.io/kubernetes/pkg/api/v1"

// ScaleSpec describes the attributes a Scale subresource
type ScaleSpec struct {
	// Replicas is the number of desired replicas.
	Replicas int `json:"replicas,omitempty" description:"number of replicas desired; see http://releases.k8s.io/HEAD/docs/user-guide/replication-controller.md#what-is-a-replication-controller"`
}

// ScaleStatus represents the current status of a Scale subresource.
type ScaleStatus struct {
	// Replicas is the number of actual replicas.
	Replicas int `json:"replicas" description:"most recently oberved number of replicas; see http://releases.k8s.io/HEAD/docs/user-guide/replication-controller.md#what-is-a-replication-controller"`

	// Selector is a label query over pods that should match the replicas count.
	Selector map[string]string `json:"selector,omitempty" description:"label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template; see http://releases.k8s.io/HEAD/docs/user-guide/labels.md#label-selectors"`
}

// Scale subresource, applicable to ReplicationControllers and (in future) Deployment.
type Scale struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata"`

	// Spec defines the behavior of the scale.
	Spec ScaleSpec `json:"spec,omitempty" description:"specification of the desired behavior of the scale; http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status"`

	// Status represents the current status of the scale.
	Status ScaleStatus `json:"status,omitempty" description:"most recently observed status of the service; populated by the system, read-only; http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status"`
}

// Dummy definition
type ReplicationControllerDummy struct {
	v1.TypeMeta `json:",inline"`
}
