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

type FederationData struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	OrigOPFederationId string `json:"origOPFederationId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Z]{2}$`
	OrigOPCountryCode string `json:"origOPCountryCode,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format=date-time
	InitialDate string `json:"initialDate"`

	// +kubebuilder:validation:Required
	PartnerStatusLink string `json:"partnerStatusLink"`

	// +kubebuilder:validation:Optional
	PartnerCallbackCredentials *Credentials `json:"partnerCallbackCredentials,omitempty"`
}

type Credentials struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ClientId string `json:"clientId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ClientSecret string `json:"clientSecret"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^https?://[^\s/$.?#].[^\s]*$`
	TokenUrl string `json:"tokenUrl"`
}

type AppIdList struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppId string `json:"appId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppProvId string `json:"appProvId"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ZoneIds []string `json:"zoneIds,omitempty"`
}

type AssocPolicy struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	PolicyId string `json:"policyId"`

	// +kubebuilder:validation:Required
	AppIdList AppIdList `json:"appIdList"`
}

type UpdateData struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=MOBILE_NETWORK_CODES;FIXED_NETWORK_CODES;OPS_POLICY;APP_POLICY
	ObjectType string `json:"objectType,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=ADD_CODES;REMOVE_CODES;UPDATE_CODES;ADD_POLICY;REMOVE_POLICY;UPDATE_POLICY
	OperationType string `json:"operationType,omitempty"`

	// +kubebuilder:validation:Optional
	AssocAppPolicies *AssocPolicy `json:"assocAppPolicies,omitempty"`

	// +kubebuilder:validation:Optional
	AssocOpsPolicies *AssocPolicy `json:"assocOpsPolicies,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=date-time
	ModificationDate string `json:"modificationDate,omitempty"`
}

type MobileNetworkIds struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^\d{3}$`
	Mcc string `json:"mcc,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	Mncs []string `json:"mncs,omitempty"`
}

// FederationSpec defines the desired state of Federation
type FederationSpec struct {
	// +kubebuilder:validation:Optional
	FederationData *FederationData `json:"federationData,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateData *UpdateData `json:"updateData,omitempty"`

	// +kubebuilder:validation:Optional
	MobileNetworkIds *MobileNetworkIds `json:"mobileNetworkIds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	FixedNetworkIds []string `json:"fixedNetworkIds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=NW_CAP_CONN_STATE_CHANGE;NW_CAP_LOCATION_RETRIEVAL;NW_CAP_USERPLANE_MGMT_EVENTS;NW_CAP_DYNAMIC_QOS
	CapType string `json:"capType,omitempty"`
}

// +kubebuilder:validation:Pattern=`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`
type IPv4String string

type IPv6String string

type ServiceEndpoint struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=edge;lcm
	EndpointType string `json:"endpointType"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=0
	Port int32 `json:"port"`

	// +kubebuilder:validation:Optional
	Fqdn string `json:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	Ipv4Addresses []IPv4String `json:"ipv4Addresses,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	Ipv6Addresses []IPv6String `json:"ipv6Addresses,omitempty"`
}

type ZoneDetail struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ZoneId string `json:"zoneId"`

	// +kubebuilder:validation:Required
	GeographyDetails string `json:"geographyDetails"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^([-+]?)([\d]{1,2})((((\.)([\d]{1,4}))?(,)))(([-+]?)([\d]{1,3})((\.)([\d]{1,4}))?)$`
	Geolocation string `json:"geolocation,omitempty"`
}

type AlarmStatus struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=RAISED;UPDATED;CLEAR
	AlarmState string `json:"alarmState"`
}

type FederationHealthInfo struct {
	// +kubebuilder:validation:Required
	Status AlarmStatus `json:"status"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format=date-time
	DateAndTimeZoneObject string `json:"dateAndTimeZoneObject"`

	// +kubebuilder:validation:Required
	NumOfAcceptedZones string `json:"numOfAcceptedZones"`

	// +kubebuilder:validation:Optional
	NumOfActiveAlarms string `json:"numOfActiveAlarms,omitempty"`

	// +kubebuilder:validation:Optional
	NumOfApplications string `json:"numOfApplications,omitempty"`
}

type Caps struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=NW_CAP_CONN_STATE_CHANGE;NW_CAP_LOCATION_RETRIEVAL;NW_CAP_USERPLANE_MGMT_EVENTS;NW_CAP_DYNAMIC_QOS
	CapabilityId string `json:"capabilityId"`

	// +kubebuilder:validation:Required
	MaxiDetectionTime string `json:"maxiDetectionTime"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=CELL_LEVEL_ACCURACY;REGISTRATION_AREA_ACCURACY;TRACKING_AREA_ACCURACY;GEO_LOCATION_ACCURACY
	LocationType string `json:"locationType"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=LAST_KNOWN_LOCATION;CURRENT_LOCATION;INITIAL_LOCATION
	LocationAccuracy string `json:"locationAccuracy,omitempty"`

	// +kubebuilder:validation:Required
	MaxUserPlaneLatency string `json:"maxUserPlaneLatency"`

	// +kubebuilder:validation:Required
	SupportedQoS string `json:"supportedQoS"`
}

type Service struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ServiceCaps []string `json:"serviceCaps,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=api_federation
	ServiceType string `json:"serviceType,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ApiRoutingInfo []string `json:"apiRoutingInfo,omitempty"`
}

// FederationStatus defines the observed state of Federation.
type FederationStatus struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	PartnerOPFederationId string `json:"partnerOPFederationId,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Z]{2}$`
	PartnerOPCountryCode string `json:"partnerOPCountryCode,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceEndpoint *ServiceEndpoint `json:"serviceEndpoint,omitempty"`

	// +kubebuilder:validation:Optional
	MobileNetworkIds *MobileNetworkIds `json:"mobileNetworkIds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	FixedNetworkIds []string `json:"fixedNetworkIds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ZoneDetails []ZoneDetail `json:"zoneDetails,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformCaps []string `json:"platformCaps,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=date-time
	FederationExpiryDate string `json:"federationExpiryDate,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=date-time
	FederationRenewalDate string `json:"federationRenewalDate,omitempty"`

	// +kubebuilder:validation:Optional
	FederationHealthInfo *FederationHealthInfo `json:"federationHealthInfo,omitempty"`

	// +kubebuilder:validation:Optional
	Caps *Caps `json:"caps,omitempty"`

	// +kubebuilder:validation:Optional
	Service *Service `json:"service,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=fed,scope=Namespaced

// Federation is the Schema for the federations API
type Federation struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// +required
	Spec FederationSpec `json:"spec"`

	// +optional
	Status FederationStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// FederationList contains a list of Federation
type FederationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Federation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Federation{}, &FederationList{})
}
