/*
Copyright 2026.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GPUResource struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=GPU_PROVIDER_NVIDIA;GPU_PROVIDER_AMD
	GpuVendorType string `json:"gpuVendorType"`

	// +kubebuilder:validation:Required
	GpuModeName string `json:"gpuModeName"`

	// +kubebuilder:validation:Required
	GpuMemory int32 `json:"gpuMemory"`

	// +kubebuilder:validation:Required
	NumGPU int32 `json:"numGPU"`
}

type Hugepage struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum="2MB";"4MB";"1GB"
	PageSize string `json:"pageSize"`

	// +kubebuilder:validation:Required
	Number int32 `json:"number"`
}

type ComputeResourceInfo struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ISA_X86_64;ISA_ARM_64
	CpuArchType string `json:"cpuArchType"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^\d+((\.\d{1,3})|(m))?$`
	NumCPU string `json:"numCPU"`

	// +kubebuilder:validation:Required
	Memory int64 `json:"memory"`

	// +kubebuilder:validation:Optional
	DiskStorage int32 `json:"diskStorage,omitempty"`

	// +kubebuilder:validation:Optional
	Gpu []GPUResource `json:"gpu,omitempty"`

	// +kubebuilder:validation:Optional
	Vpu int32 `json:"vpu,omitempty"`

	// +kubebuilder:validation:Optional
	Fpga int32 `json:"fpga,omitempty"`

	// +kubebuilder:validation:Optional
	Hugepages []Hugepage `json:"hugepages,omitempty"`

	// +kubebuilder:validation:Optional
	CpuExclusivity bool `json:"cpuExclusivity,omitempty"`
}

type SupportedOSType struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=x86_64;x86
	Architecture string `json:"architecture"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=RHEL;UBUNTU;COREOS;FEDORA;WINDOWS;OTHER
	Distribution string `json:"distribution"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=OS_VERSION_UBUNTU_2204_LTS;OS_VERSION_RHEL_8;OS_VERSION_RHEL_7;OS_VERSION_DEBIAN_11;OS_VERSION_COREOS_STABLE;OS_MS_WINDOWS_2012_R2;OTHER
	Version string `json:"version"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=OS_LICENSE_TYPE_FREE;OS_LICENSE_TYPE_ON_DEMAND;NOT_SPECIFIED
	License string `json:"license"`
}

type FlavourSupported struct {
	// +kubebuilder:validation:Required
	FlavourId string `json:"flavourId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ISA_X86;ISA_X86_64;ISA_ARM_64
	CpuArchType string `json:"cpuArchType"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	SupportedOSTypes []SupportedOSType `json:"supportedOSTypes"`

	// +kubebuilder:validation:Required
	NumCPU int32 `json:"numCPU"`

	// +kubebuilder:validation:Required
	MemorySize int32 `json:"memorySize"`

	// +kubebuilder:validation:Required
	StorageSize int32 `json:"storageSize"`

	// +kubebuilder:validation:Optional
	Gpu []GPUResource `json:"gpu,omitempty"`

	// +kubebuilder:validation:Optional
	Fpga int32 `json:"fpga,omitempty"`

	// +kubebuilder:validation:Optional
	Vpu int32 `json:"vpu,omitempty"`

	// +kubebuilder:validation:Optional
	Hugepages []Hugepage `json:"hugepages,omitempty"`

	// +kubebuilder:validation:Optional
	CpuExclusivity bool `json:"cpuExclusivity,omitempty"`
}

type NetworkResources struct {
	// +kubebuilder:validation:Required
	EgressBandWidth int32 `json:"egressBandWidth"`

	// +kubebuilder:validation:Required
	DedicatedNIC int32 `json:"dedicatedNIC"`

	// +kubebuilder:validation:Required
	SupportSriov bool `json:"supportSriov"`

	// +kubebuilder:validation:Required
	SupportDPDK bool `json:"supportDPDK"`
}

type LatencyRanges struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	MinLatency int32 `json:"minLatency,omitempty"`

	// +kubebuilder:validation:Optional
	MaxLatency int32 `json:"maxLatency,omitempty"`
}

type JitterRanges struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	MinJitter int32 `json:"minJitter,omitempty"`

	// +kubebuilder:validation:Optional
	MaxJitter int32 `json:"maxJitter,omitempty"`
}

type ThroughputRanges struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	MinThroughput int32 `json:"minThroughput,omitempty"`

	// +kubebuilder:validation:Optional
	MaxThroughput int32 `json:"maxThroughput,omitempty"`
}

type ZoneServiceLevelObjsInfo struct {
	// +kubebuilder:validation:Required
	LatencyRanges LatencyRanges `json:"latencyRanges"`

	// +kubebuilder:validation:Required
	JitterRanges JitterRanges `json:"jitterRanges"`

	// +kubebuilder:validation:Required
	ThroughputRanges ThroughputRanges `json:"throughputRanges"`
}

// AvaiabilityZoneSpec defines the desired state of AvaiabilityZone
type AvaiabilityZoneSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ZoneId string `json:"zoneId,omitempty"`

	// +kubebuilder:validation:Optional
	Link string `json:"link,omitempty"`
}

// AvaiabilityZoneStatus defines the observed state of AvaiabilityZone.
type AvaiabilityZoneStatus struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ReservedComputeResources []ComputeResourceInfo `json:"reservedComputeResources,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ComputeResourceQuotaLimits []ComputeResourceInfo `json:"computeResourceQuotaLimits,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	FlavoursSupported []FlavourSupported `json:"flavoursSupported,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkResources *NetworkResources `json:"networkResources,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneServiceLevelObjsInfo *ZoneServiceLevelObjsInfo `json:"zoneServiceLevelObjsInfo,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=zone,scope=Namespaced
// +kubebuilder:printcolumn:name="FederationContextId",type=string,JSONPath=`.spec.federationContextId`

// AvaiabilityZone is the Schema for the avaiabilityzones API
type AvaiabilityZone struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`
	// +required
	Spec AvaiabilityZoneSpec `json:"spec"`
	// +optional
	Status AvaiabilityZoneStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// AvaiabilityZoneList contains a list of AvaiabilityZone
type AvaiabilityZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []AvaiabilityZone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AvaiabilityZone{}, &AvaiabilityZoneList{})
}
