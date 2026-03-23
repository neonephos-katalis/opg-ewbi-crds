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

type ImgOSType struct {
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

type ImageRepoLocation struct {
	// +kubebuilder:validation:Optional
	RepoURL string `json:"repoURL,omitempty"`

	// +kubebuilder:validation:Optional
	UserName string `json:"userName,omitempty"`

	// +kubebuilder:validation:Optional
	Password string `json:"password,omitempty"`

	// +kubebuilder:validation:Optional
	Token string `json:"token,omitempty"`
}

type ImageBody struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppProviderId string `json:"appProviderId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,31}$`
	ImageName string `json:"imageName"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinLength=8
	// +kubebuilder:validation:MaxLength=128
	ImageDescription string `json:"imageDescription,omitempty"`

	// +kubebuilder:validation:Required
	ImageVersionInfo string `json:"imageVersionInfo"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=QCOW2;DOCKER;OVA
	ImageType string `json:"imageType"`

	// +kubebuilder:validation:Optional
	Checksum string `json:"checksum,omitempty"`

	// +kubebuilder:validation:Required
	ImgOSType ImgOSType `json:"imgOSType"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ISA_X86_64;ISA_ARM_64
	ImgInsSetArch string `json:"imgInsSetArch"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=PRIVATEREPO;PUBLICREPO;UPLOAD
	RepoType string `json:"repoType,omitempty"`

	// +kubebuilder:validation:Optional
	ImageRepoLocation []ImageRepoLocation `json:"imageRepoLocation,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=byte
	Image []byte `json:"image,omitempty"` // Nota: In Go, 'type: string, format: byte' nello YAML di solito si mappa con []byte, che JSON serializza in base64.
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format=uuid
	ImageId string `json:"imageId"`

	// +kubebuilder:validation:Optional
	ImageBody *ImageBody `json:"imageBody,omitempty"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=Pending;Uploading;Uploaded;Failed
	State string `json:"state,omitempty"`

	// +kubebuilder:validation:Optional
	Message string `json:"message,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=date-time
	LastUpdated string `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Image is the Schema for the images API
type Image struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// +required
	Spec ImageSpec `json:"spec"`

	// +optional
	Status ImageStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=image,scope=Namespaced

// ImageList contains a list of Image
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Image `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
