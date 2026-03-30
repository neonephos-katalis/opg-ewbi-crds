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
	// Globally unique identifier allocated to an operator platform. This is valid and used only in context of MEC federation interface.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	OrigOPFederationId string `json:"origOPFederationId,omitempty"`

	// ISO 3166-1 Alpha-2 code for the country of Partner operator
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Z]{2}$`
	OrigOPCountryCode string `json:"origOPCountryCode,omitempty"`

	// Time zone info of the federation initiated by the originating OP
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

// (ZoneIdentifier) Human readable name of the zone.
// +kubebuilder:validation:Type=string
type ZoneId string

type AppIdList struct {
	// (AppIdentifier) Identifier used to refer to an application.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppId string `json:"appId"`

	// (AppProviderId) UserId of the app provider.  Identifier is relevant only in context of this federation.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppProvId string `json:"appProvId"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:items:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ZoneIds []ZoneId `json:"zoneIds,omitempty"`
}

type AssocPolicy struct {
	// (ApplPolicyIdentifier) Application-level Policy unique identifier"
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	PolicyId string `json:"policyId"`

	// AppIdLocList
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

	// Date and time of the federation modification by the originating partner OP
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
	// +kubebuilder:validation:items:Pattern=`^\d{2,3}$`
	Mncs []string `json:"mncs,omitempty"`
}

// List of network identifier associated with the fixedline network of the operator platform.
// +kubebuilder:validation:Type=string
type FixedNetworkId string

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
	FixedNetworkIds []FixedNetworkId `json:"fixedNetworkIds,omitempty"`

	// The enumerated list of network capabilities that an OP can use for various services via SBI-NR.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=NW_CAP_CONN_STATE_CHANGE;NW_CAP_LOCATION_RETRIEVAL;NW_CAP_USERPLANE_MGMT_EVENTS;NW_CAP_DYNAMIC_QOS
	CapType string `json:"capType,omitempty"`
}

// +kubebuilder:validation:Pattern=`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`
type IPv4String string

type IPv6String string

// Service Endpoint
type ServiceEndpoint struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=edge;lcm
	EndpointType string `json:"endpointType"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=0
	Port int `json:"port"`

	// (EdgeAppFQDN) DNS FQDN assigned to application instances in an availability zone. User Clients can resolve the FQDN to communicate with the edge instances of the application
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
	// Human readable name of the zone.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	ZoneId string `json:"zoneId"`

	// Details about cities or state covered by the edge. Details about the type of locality for eg rural, urban, industrial etc. This information is defined in human readable form.
	// +kubebuilder:validation:Required
	GeographyDetails string `json:"geographyDetails"`

	// Latitude,Longitude as decimal fraction up to 4 digit precision
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^([-+]?)([\d]{1,2})((((\.)([\d]{1,4}))?(,)))(([-+]?)([\d]{1,3})((\.)([\d]{1,4}))?)$`
	Geolocation string `json:"geolocation,omitempty"`
}

type AlarmStatus struct {
	// Defines the alarm state during its life cycle (raised | updated | cleared).
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=RAISED;UPDATED;CLEAR
	AlarmState string `json:"alarmState"`
}

type FederationHealthInfo struct {
	// +kubebuilder:validation:Required
	Status AlarmStatus `json:"status"`

	// Date and Time zone info format
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
	// The enumerated list of network capabilities that an OP can use for various services via SBI-NR.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=NW_CAP_CONN_STATE_CHANGE;NW_CAP_LOCATION_RETRIEVAL;NW_CAP_USERPLANE_MGMT_EVENTS;NW_CAP_DYNAMIC_QOS
	CapabilityId string `json:"capabilityId"`

	// The maximum detection time in seconds that the OP can determine the UE change of connectivity with the mobile network.
	// +kubebuilder:validation:Required
	MaxiDetectionTime string `json:"maxiDetectionTime"`

	// The enumerated list of UE location accuracy that an OP can determine via SBI-NR.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=CELL_LEVEL_ACCURACY;REGISTRATION_AREA_ACCURACY;TRACKING_AREA_ACCURACY;GEO_LOCATION_ACCURACY
	LocationType string `json:"locationType"`

	// The enumerated list of type of network location of an UE that an OP can determine via SBI-NR.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=LAST_KNOWN_LOCATION;CURRENT_LOCATION;INITIAL_LOCATION
	LocationAccuracy string `json:"locationAccuracy,omitempty"`

	// Indicates the maximum user plane latency in units of milliseconds to decide whether edge relocation is needed to ascertain latency remain in this range.
	// +kubebuilder:validation:Required
	MaxUserPlaneLatency string `json:"maxUserPlaneLatency"`

	// Set of one or more 5G QoS Identifier (5QI or 4G QCI) created via concatanation of Resource Type and 5QI values i.e., GBR1, GBR2, GBR65, NONGBR79 etc.
	// +kubebuilder:validation:Required
	SupportedQoS string `json:"supportedQoS"`
}

type Service struct {
	// (serviceAPINames) List of Service API capability names an OP supports and offers to other OPs "quality_on_demand", "device_location" etc.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:items:Enum=QualityOnDemand;DeviceLocation;DeviceStatus;SimSwap;NumberVerification;DeviceIdentifier
	ServiceCaps []string `json:"serviceCaps,omitempty"`

	// An identifier to refer to partner OP capabilities for application providers.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=api_federation
	ServiceType string `json:"serviceType,omitempty"`

	// (serviceRoutingInfo) List of public IP addresses MNO manages for UEs to connect with public data networks
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:items:Pattern=`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\\/([0-9]|[1-2][0-9]|3[0-2]))?$`
	ApiRoutingInfo []string `json:"apiRoutingInfo,omitempty"`
}

// FederationStatus defines the observed state of Federation.
type FederationStatus struct {
	// This identifier shall be provided by the partner OP on successful verification and validation of the federation create request and is used by partner op to identify this newly created federation context. Originating OP shall provide this identifier in any subsequent request towards the partner op.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId,omitempty"`

	// (FederationIdentifier) Globally unique identifier allocated to an operator platform. This is valid and used only in context of MEC federation interface.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	PartnerOPFederationId string `json:"partnerOPFederationId,omitempty"`

	// (CountryCode) ISO 3166-1 Alpha-2 code for the country of Partner operator
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Z]{2}$`
	PartnerOPCountryCode string `json:"partnerOPCountryCode,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceEndpoint *ServiceEndpoint `json:"serviceEndpoint,omitempty"`

	// MobileNetworkIds
	// +kubebuilder:validation:Optional
	MobileNetworkIds *MobileNetworkIds `json:"mobileNetworkIds,omitempty"`

	// (FixedNetworkIds) List of network identifier associated with the fixed line network of the operator platform.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	FixedNetworkIds []string `json:"fixedNetworkIds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ZoneDetails []ZoneDetail `json:"zoneDetails,omitempty"`

	// Home routing - Operator platform is capable of routing edge application data traffic from its edges to user device in their home location. This is the case where user devices are served in their home region (requesting platform region, non-roaming) but the corresponding edge application are in operator platform edges. Anchoring - Operator platform is capable of routing edge application traffic for roaming user devices to edge application in user device home network. Service APIs - Capability to handle Service APIs (e.g., CAMARA APIs) from the Leading OP
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:items:Enum=homeRouting;Anchoring;serviceAPIs;faultMgmt;eventMgmt;resourceMonitor;networkEventMgmt;appNotificationMgmt;appLevelPolicyMgmt;opsLevelPolicyMgmt
	PlatformCaps []string `json:"platformCaps,omitempty"`

	//Date and Time zone info of the existing federation expiry
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=date-time
	FederationExpiryDate string `json:"federationExpiryDate,omitempty"`

	// Date and Time zone info of the existing federation renewal. Shall be less than federationExpiryDate
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
