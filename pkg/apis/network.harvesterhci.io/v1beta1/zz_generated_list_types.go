/*
Copyright 2019 Harvester Network Controller Authors

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

// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=network.harvesterhci.io
package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterNetworkList is a list of ClusterNetwork resources
type ClusterNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterNetwork `json:"items"`
}

func NewClusterNetwork(namespace, name string, obj ClusterNetwork) *ClusterNetwork {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ClusterNetwork").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkList is a list of NodeNetwork resources
type NodeNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeNetwork `json:"items"`
}

func NewNodeNetwork(namespace, name string, obj NodeNetwork) *NodeNetwork {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("NodeNetwork").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VlanConfigList is a list of VlanConfig resources
type VlanConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VlanConfig `json:"items"`
}

func NewVlanConfig(namespace, name string, obj VlanConfig) *VlanConfig {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VlanConfig").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VlanStatusList is a list of VlanStatus resources
type VlanStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VlanStatus `json:"items"`
}

func NewVlanStatus(namespace, name string, obj VlanStatus) *VlanStatus {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VlanStatus").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}
