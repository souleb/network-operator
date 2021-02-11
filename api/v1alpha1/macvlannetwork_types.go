/*
Copyright 2021 NVIDIA

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

const (
	MacvlanNetworkCRDName = "MacvlanNetwork"
)

// MacvlanNetworkSpec defines the desired state of MacvlanNetwork
type MacvlanNetworkSpec struct {
	// Namespace of the NetworkAttachmentDefinition custom resource
	NetworkNamespace string `json:"networkNamespace,omitempty"`
	// Name of the host interface to enslave. Defaults to default route interface
	Master string `json:"master,omitempty"`
	// +kubebuilder:validation:Enum={"bridge", "private", "vepa", "passthru"}
	// Mode of interface one of "bridge", "private", "vepa", "passthru"
	Mode string `json:"mode,omitempty"`
	// MTU of interface to the specified value. 0 for master's MTU
	// +kubebuilder:validation:Minimum=0
	Mtu int `json:"mtu,omitempty"`
	// IPAM configuration to be used for this network.
	IPAM string `json:"ipam,omitempty"`
}

// MacvlanNetworkStatus defines the observed state of MacvlanNetwork
type MacvlanNetworkStatus struct {
	// Reflects the state of the MacvlanNetwork
	// +kubebuilder:validation:Enum={"notReady", "ready", "error"}
	State State `json:"state"`
	// Network attachment definition generated from MacvlanNetworkSpec
	MacvlanNetworkAttachmentDef string `json:"macvlanNetworkAttachmentDef,omitempty"`
	// Informative string in case the observed state is error
	Reason string `json:"reason,omitempty"`
}

// +kubebuilder:object:root=true
// kubebuilder:object:generate
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.state`,priority=0

// MacvlanNetwork is the Schema for the macvlannetworks API
type MacvlanNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MacvlanNetworkSpec   `json:"spec,omitempty"`
	Status MacvlanNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// kubebuilder:object:generate

// MacvlanNetworkList contains a list of MacvlanNetwork
type MacvlanNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MacvlanNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MacvlanNetwork{}, &MacvlanNetworkList{})
}
