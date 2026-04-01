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

type K8sOperations struct {
	// +kubebuilder:validation:Required
	Href string `json:"href"`

	// +kubebuilder:validation:Required
	Version string `json:"version"`

	// +kubebuilder:validation:Required
	Group string `json:"group"`

	// +kubebuilder:validation:Required
	Plural string `json:"plural"`
}

type APIOperation struct {
	// +kubebuilder:validation:Required
	Href string `json:"href"`

	// List of HTTP Methods supported for the given API category
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	HttpMethods []string `json:"httpMethods"`
}

type APICategory struct {
	// +kubebuilder:validation:Optional
	K8sOperations *K8sOperations `json:"k8sOperations,omitempty"`

	// List of endpoints for the given API category
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ApiOperations []APIOperation `json:"apiOperations,omitempty"`
}

type APIList struct {
	// +kubebuilder:validation:Required
	FederationManagement APICategory `json:"federationmanagement"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *APICategory `json:"availabilityzone,omitempty"`

	// +kubebuilder:validation:Optional
	ArtefactManagement *APICategory `json:"artefactmanagement,omitempty"`

	// +kubebuilder:validation:Optional
	FileManagement *APICategory `json:"filemanagement,omitempty"`

	// +kubebuilder:validation:Optional
	OnboardingManagement *APICategory `json:"onboardingmanagement,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentManagement *APICategory `json:"deploymentmanagement,omitempty"`
}

// FederationAPISpec defines the desired state of FederationAPI
type FederationAPISpec struct {
	// This identifier shall be provided by the partner OP on successful verification and validation of the federation create request and is used by partner op to identify this newly created federation context. Originating OP shall provide this identifier in any subsequent request towards the partner op.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId,omitempty"`
}

// FederationAPIStatus defines the observed state of FederationAPI.
type FederationAPIStatus struct {
	// +kubebuilder:validation:Optional
	Api *APIList `json:"api,omitempty"`

	// Current state of the artefact upload
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=Pending;Uploading;Uploaded;Failed
	State string `json:"state,omitempty"`

	// message indicating details about the current state
	// +kubebuilder:validation:Optional
	Message string `json:"message,omitempty"`

	// Timestamp of the last status update
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=date-time
	LastUpdated string `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=fedapi,scope=Namespaced

// FederationAPI is the Schema for the federationapis API
type FederationAPI struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of FederationAPI
	// +required
	Spec FederationAPISpec `json:"spec"`

	// status defines the observed state of FederationAPI
	// +optional
	Status FederationAPIStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// FederationAPIList contains a list of FederationAPI
type FederationAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []FederationAPI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FederationAPI{}, &FederationAPIList{})
}
