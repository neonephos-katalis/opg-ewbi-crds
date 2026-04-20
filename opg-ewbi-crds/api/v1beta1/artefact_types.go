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

// +kubebuilder:validation:Format=uuid
type ImageUUID string

// CommandString definisce una stringa per gli array dei comandi
type CommandString string

type ArtefactRepoLocation struct {
	// +kubebuilder:validation:Optional
	RepoURL string `json:"repoURL,omitempty"`

	// +kubebuilder:validation:Optional
	UserName string `json:"userName,omitempty"`

	// +kubebuilder:validation:Optional
	Password string `json:"password,omitempty"`

	// +kubebuilder:validation:Optional
	Token string `json:"token,omitempty"`
}

// CommandLineParams descrive i parametri da riga di comando
type CommandLineParams struct {
	// +kubebuilder:validation:Required
	Command []CommandString `json:"command"`

	// +kubebuilder:validation:Optional
	CommandArgs []CommandString `json:"commandArgs,omitempty"`
}

type ExposedInterface struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,30}[A-Za-z0-9]$`
	InterfaceId string `json:"interfaceId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=TCP;UDP;HTTP_HTTPS
	CommProtocol string `json:"commProtocol"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	CommPort int32 `json:"commPort"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=VISIBILITY_EXTERNAL;VISIBILITY_INTERNAL
	VisibilityType string `json:"visibilityType"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{6,30}[A-Za-z0-9]$`
	Network string `json:"network,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[a-z][a-z0-9]{3}$`
	InterfaceName string `json:"InterfaceName,omitempty"`
}

type CompEnvParam struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,30}[A-Za-z0-9]$`
	EnvVarName string `json:"envVarName"` // Nota: Nello YAML era envVarNam in required e envVarName in properties. Uso la property.

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=USER_DEFINED;PLATFORM_DEFINED_DYNAMIC_PORT;PLATFORM_DEFINED_DNS;PLATFORM_DEFINED_IP
	EnvValueType string `json:"envValueType"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,62}[A-Za-z0-9]$`
	EnvVarValue string `json:"envVarValue,omitempty"`

	// +kubebuilder:validation:Optional
	EnvVarSrc string `json:"envVarSrc,omitempty"`
}

type DeploymentConfig struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=DOCKER_COMPOSE;KUBERNETES_MANIFEST;CLOUD_INIT;HELM_VALUES
	ConfigType string `json:"configType"`

	// +kubebuilder:validation:Required
	Contents string `json:"contents"`
}

type PersistentVolume struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum="10Gi";"20Gi";"50Gi";"100Gi"
	VolumeSize string `json:"volumeSize"`

	// +kubebuilder:validation:Required
	VolumeMountPath string `json:"volumeMountPath"`

	// +kubebuilder:validation:Required
	VolumeName string `json:"volumeName"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	EphemeralType bool `json:"ephemeralType,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=RW;RO
	// +kubebuilder:default=RW
	AccessMode string `json:"accessMode,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=EXCLUSIVE;SHARED
	// +kubebuilder:default=EXCLUSIVE
	SharingPolicy string `json:"sharingPolicy,omitempty"`
}

type ComponentSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9_]{6,62}[A-Za-z0-9]$`
	ComponentName string `json:"componentName"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	Images []ImageUUID `json:"images"`

	// +kubebuilder:validation:Required
	NumOfInstances int32 `json:"numOfInstances"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=RESTART_POLICY_ALWAYS;RESTART_POLICY_NEVER
	RestartPolicy string `json:"restartPolicy"`

	// +kubebuilder:validation:Required
	ComputeResourceProfile ComputeResourceInfo `json:"computeResourceProfile"`

	// +kubebuilder:validation:Optional
	CommandLineParams *CommandLineParams `json:"commandLineParams,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	ExposedInterfaces []ExposedInterface `json:"exposedInterfaces,omitempty"`

	// +kubebuilder:validation:Optional
	CompEnvParams []CompEnvParam `json:"compEnvParams,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentConfig *DeploymentConfig `json:"deploymentConfig,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinItems=1
	PersistentVolumes []PersistentVolume `json:"persistentVolumes,omitempty"`
}

type ArtefactBody struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,63}$`
	AppProviderId string `json:"appProviderId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z][A-Za-z0-9_]{7,31}$`
	ArtefactName string `json:"artefactName"`

	// +kubebuilder:validation:Required
	ArtefactVersionInfo string `json:"artefactVersionInfo"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MaxLength=256
	ArtefactDescription string `json:"artefactDescription,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=VM_TYPE;CONTAINER_TYPE
	ArtefactVirtType string `json:"artefactVirtType"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinLength=8
	// +kubebuilder:validation:MaxLength=32
	ArtefactFileName string `json:"artefactFileName,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=ZIP;TAR;TEXT;TARGZ
	ArtefactFileFormat string `json:"artefactFileFormat,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=HELM;TERRAFORM;ANSIBLE;SHELL;COMPONENTSPEC
	ArtefactDescriptorType string `json:"artefactDescriptorType"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=PRIVATEREPO;PUBLICREPO;UPLOAD
	RepoType string `json:"repoType,omitempty"`

	// +kubebuilder:validation:Optional
	ArtefactRepoLocation *ArtefactRepoLocation `json:"artefactRepoLocation,omitempty"`

	// +kubebuilder:validation:Required
	ComponentSpec ComponentSpec `json:"componentSpec"`
}

// ArtefactSpec defines the desired state of Artefact
type ArtefactSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9][A-Za-z0-9-]*$`
	FederationContextId string `json:"federationContextId"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format=uuid
	ArtefactId string `json:"artefactId"`

	// +kubebuilder:validation:Optional
	ArtefactBody *ArtefactBody `json:"artefactBody,omitempty"`
}

// ArtefactStatus defines the observed state of Artefact.
type ArtefactStatus struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=Pending;Uploading;Uploaded;Failed
	State string `json:"state,omitempty"`
	// +kubebuilder:validation:Optional
	Message     string      `json:"message,omitempty"`
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Artefact is the Schema for the artefacts API
type Artefact struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`
	// +required
	Spec ArtefactSpec `json:"spec"`

	// +optional
	Status ArtefactStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=art,scope=Namespaced
// +kubebuilder:printcolumn:name="FederationContextId",type=string,JSONPath=`.spec.federationContextId`
// +kubebuilder:printcolumn:name="ArtefactId",type=string,JSONPath=`.spec.artefactId`
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="LastUpdated",type=string,JSONPath=`.status.lastUpdated`

// ArtefactList contains a list of Artefact
type ArtefactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Artefact `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Artefact{}, &ArtefactList{})
}
