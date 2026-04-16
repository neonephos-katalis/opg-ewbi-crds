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

// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
type ZoneIdString string

type ZoneSetting struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ZoneId []ZoneIdString `json:"zoneId,omitempty"`

	// +kubebuilder:validation:Optional
	Forbid bool `json:"forbid,omitempty"`
}

type AppDeploymentZone struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ZoneId string `json:"zoneId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Z]{2}$`
	CountryCode string `json:"countryCode,omitempty"`
}

type AppMetaData struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,31}$`
	AppName string `json:"appName"`

	// +kubebuilder:validation:Required
	Version string `json:"version"`

	// +kubebuilder:validation:Required
	AccessToken string `json:"accessToken"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinLength=16
	// +kubebuilder:validation:MaxLength=256
	AppDescription string `json:"appDescription,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	MobilitySupport bool `json:"mobilitySupport,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=IOT;HEALTH_CARE;GAMING;VIRTUAL_REALITY;SOCIALIZING;SURVEILLANCE;ENTERTAINMENT;CONNECTIVITY;PRODUCTIVITY;SECURITY;INDUSTRIAL;EDUCATION;OTHERS
	Category string `json:"category,omitempty"`
}

type AppQoSProfile struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=NONE;LOW;ULTRALOW
	LatencyConstraints string `json:"latencyConstraints"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	BandwidthRequired int32 `json:"bandwidthRequired,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=APP_TYPE_SINGLE_USER;APP_TYPE_MULTI_USER
	MultiUserClients string `json:"multiUserClients,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	NoOfUsersPerAppInst int32 `json:"noOfUsersPerAppInst,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	AppProvisioning bool `json:"appProvisioning,omitempty"`
}

type AppComponentSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])\.)*([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])$`
	ServiceNameNB string `json:"serviceNameNB,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])\.)*([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])$`
	ServiceNameEW string `json:"serviceNameEW,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,62}[A-Za-z0-9]$`
	ComponentName string `json:"componentName,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=uuid
	ArtefactId string `json:"artefactId,omitempty"`
}

type AppInfo struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppId string `json:"appId"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppProviderId string `json:"appProviderId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	AppDeploymentZones []AppDeploymentZone `json:"appDeploymentZones,omitempty"`

	// +kubebuilder:validation:Optional
	AppMetaData *AppMetaData `json:"appMetaData,omitempty"`

	// +kubebuilder:validation:Optional
	AppQoSProfile *AppQoSProfile `json:"appQoSProfile,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	AppComponentSpecs []AppComponentSpec `json:"appComponentSpecs,omitempty"`

	// +kubebuilder:validation:Optional
	AppStatusCallbackLink string `json:"appStatusCallbackLink,omitempty"`

	// +kubebuilder:validation:Optional
	EdgeAppFQDN string `json:"edgeAppFQDN,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=PENDING;ONBOARDED;DEBOARDING;REMOVED;FAILED
	OnboardStatusInfo string `json:"onboardStatusInfo,omitempty"`
}

// ApplicationOnboardingSpec defines the desired state of ApplicationOnboarding
type ApplicationOnboardingSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ZoneSettings []ZoneSetting `json:"zoneSettings,omitempty"`

	// +kubebuilder:validation:Required
	AppInfo AppInfo `json:"appInfo"`
}

// ApplicationOnboardingStatus defines the observed state of ApplicationOnboarding.
type ApplicationOnboardingStatus struct {
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
// +kubebuilder:resource:shortName=onboard,scope=Namespaced

// ApplicationOnboarding is the Schema for the applicationonboardings API
type ApplicationOnboarding struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`
	// +required
	Spec ApplicationOnboardingSpec `json:"spec"`
	// +optional
	Status ApplicationOnboardingStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ApplicationOnboardingList contains a list of ApplicationOnboarding
type ApplicationOnboardingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ApplicationOnboarding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApplicationOnboarding{}, &ApplicationOnboardingList{})
}
