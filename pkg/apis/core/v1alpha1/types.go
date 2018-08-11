/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroup is a collection of Pod; used for batch workload.
type PodGroup struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the pod group.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec PodGroupSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current information about a pod group.
	// This data may not be up to date.
	// +optional
	Status PodGroupStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// PodGroupSpec represents the template of a pod group.
type PodGroupSpec struct {
	// NumMember defines the number of members/tasks to run the pod group;
	// if there's not enough resources to start all tasks, the scheduler
	// will not start anyone.
	NumMember int32 `json:"numMember,omitempty" protobuf:"bytes,1,opt,name=numMember"`
}

// PodGroupStatus represents the current state of a pod group.
type PodGroupStatus struct {
	// The number of actively running pods.
	// +optional
	Running int32 `json:"running,omitempty" protobuf:"bytes,3,opt,name=running"`
	// The number of pods which reached phase Succeeded.
	// +optional
	Succeeded int32 `json:"succeeded,omitempty" protobuf:"bytes,3,opt,name=succeeded"`
	// The number of pods which reached phase Failed.
	// +optional
	Failed int32 `json:"failed,omitempty" protobuf:"bytes,3,opt,name=failed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroupList is a collection of pod groups.
type PodGroupList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of PodGroup
	Items []PodGroup `json:"items" protobuf:"bytes,2,rep,name=items"`
}