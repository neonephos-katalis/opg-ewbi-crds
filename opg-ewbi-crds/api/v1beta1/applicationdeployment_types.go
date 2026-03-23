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

// ZoneInfo definisce le informazioni della zona per l'applicazione
type ZoneInfo struct {
	// +kubebuilder:validation:Optional
	FlavourId string `json:"flavourId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=RESERVED_RES_SHALL;RESERVED_RES_PREFER;RESERVED_RES_AVOID;RESERVED_RES_FORBID
	// +kubebuilder:default=RESERVED_RES_AVOID
	ResourceConsumption string `json:"resourceConsumption,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,30}[A-Za-z0-9]$`
	ResPool string `json:"resPool,omitempty"`
}

// AppDetails definisce i dettagli dell'applicazione
type AppDetails struct {
	// +kubebuilder:validation:Optional
	AppVersion string `json:"appVersion,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneInfo *ZoneInfo `json:"zoneInfo,omitempty"`

	// +kubebuilder:validation:Optional
	AppInstCallbackLink string `json:"appInstCallbackLink,omitempty"`
}

type AccessPoints struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=0
	Port int32 `json:"port"`

	// +kubebuilder:validation:Optional
	Fqdn string `json:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	Ipv4Addresses []IPv4String `json:"ipv4Addresses,omitempty"` // Utilizzo del tipo custom

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	// Nota: Le regole pattern 'allOf' originali andranno verificate tramite Webhook.
	Ipv6Addresses []IPv6String `json:"ipv6Addresses,omitempty"`
}

type AccessPointInfo struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,30}[A-Za-z0-9]$`
	InterfaceId string `json:"interfaceId"`

	// +kubebuilder:validation:Required
	AccessPoints AccessPoints `json:"accessPoints"`
}

type AppInstanceInfo struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,62}[A-Za-z0-9]$`
	AppInstIdentifier string `json:"appInstIdentifier"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=PENDING;READY;FAILED;TERMINATING
	AppInstanceState string `json:"appInstanceState"`
}

// ApplicationDeploymentSpec defines the desired state of ApplicationDeployment
type ApplicationDeploymentSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=uuid
	TransactionId string `json:"transactionId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppId string `json:"appId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,62}[A-Za-z0-9]$`
	AppInstanceId string `json:"appInstanceId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ZoneId string `json:"zoneId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppProviderId string `json:"appProviderId,omitempty"`

	// +kubebuilder:validation:Optional
	AppDetails *AppDetails `json:"appDetails,omitempty"`
}

// ApplicationDeploymentStatus defines the observed state of ApplicationDeployment.
type ApplicationDeploymentStatus struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	AccessPointInfo []AccessPointInfo `json:"accessPointInfo,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	AppInstanceInfo []AppInstanceInfo `json:"appInstanceInfo,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=deploy,scope=Namespaced

// ApplicationDeployment is the Schema for the applicationdeployments API
type ApplicationDeployment struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`
	// +required
	Spec ApplicationDeploymentSpec `json:"spec"`

	// +optional
	Status ApplicationDeploymentStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ApplicationDeploymentList contains a list of ApplicationDeployment
type ApplicationDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ApplicationDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApplicationDeployment{}, &ApplicationDeploymentList{})
}
