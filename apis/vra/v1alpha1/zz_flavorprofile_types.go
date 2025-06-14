// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FlavorMappingInitParameters struct {

	// Number of CPU cores. Mandatory for private clouds such as vSphere. Only instance_type or cpu_count/memory must be specified.
	// Number of CPU cores. Mandatory for private clouds such as vSphere. Only `instance_type` or `cpu_count`/`memory` must be specified.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`

	// The value of the instance type in the corresponding cloud. Mandatory for public clouds. Only instance_type or cpu_count/memory must be specified.
	// The value of the instance type in the corresponding cloud. Mandatory for public clouds. Only `instance_type` or `cpu_count`/`memory` must be specified.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Only instance_type or cpu_count/memory must be specified.
	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Only `instance_type` or `cpu_count`/`memory` must be specified.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// The name of the flavor mapping.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FlavorMappingObservation struct {

	// Number of CPU cores. Mandatory for private clouds such as vSphere. Only instance_type or cpu_count/memory must be specified.
	// Number of CPU cores. Mandatory for private clouds such as vSphere. Only `instance_type` or `cpu_count`/`memory` must be specified.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`

	// The value of the instance type in the corresponding cloud. Mandatory for public clouds. Only instance_type or cpu_count/memory must be specified.
	// The value of the instance type in the corresponding cloud. Mandatory for public clouds. Only `instance_type` or `cpu_count`/`memory` must be specified.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Only instance_type or cpu_count/memory must be specified.
	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Only `instance_type` or `cpu_count`/`memory` must be specified.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// The name of the flavor mapping.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FlavorMappingParameters struct {

	// Number of CPU cores. Mandatory for private clouds such as vSphere. Only instance_type or cpu_count/memory must be specified.
	// Number of CPU cores. Mandatory for private clouds such as vSphere. Only `instance_type` or `cpu_count`/`memory` must be specified.
	// +kubebuilder:validation:Optional
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`

	// The value of the instance type in the corresponding cloud. Mandatory for public clouds. Only instance_type or cpu_count/memory must be specified.
	// The value of the instance type in the corresponding cloud. Mandatory for public clouds. Only `instance_type` or `cpu_count`/`memory` must be specified.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Only instance_type or cpu_count/memory must be specified.
	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Only `instance_type` or `cpu_count`/`memory` must be specified.
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// The name of the flavor mapping.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type FlavorProfileInitParameters struct {

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of the flavor mappings defined for the corresponding cloud end-point region.
	// A list of the flavor mappings defined for the corresponding cloud end-point region.
	FlavorMapping []FlavorMappingInitParameters `json:"flavorMapping,omitempty" tf:"flavor_mapping,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the region for which this profile is defined.
	// The id of the region for which this profile is defined
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`
}

type FlavorProfileLinksInitParameters struct {
}

type FlavorProfileLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type FlavorProfileLinksParameters struct {
}

type FlavorProfileObservation struct {

	// Id of the cloud account this flavor profile belongs to.
	// Id of the cloud account this flavor profile belongs to.
	CloudAccountID *string `json:"cloudAccountId,omitempty" tf:"cloud_account_id,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the region for which this profile is defined.
	// The id of the region for which this profile is defined.
	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	// A list of the flavor mappings defined for the corresponding cloud end-point region.
	// A list of the flavor mappings defined for the corresponding cloud end-point region.
	FlavorMapping []FlavorMappingObservation `json:"flavorMapping,omitempty" tf:"flavor_mapping,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Hypermedia as the Engine of Application State (HATEOAS) of entity.
	Links []FlavorProfileLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of the user that owns the entity.
	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The id of the region for which this profile is defined.
	// The id of the region for which this profile is defined
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type FlavorProfileParameters struct {

	// A human-friendly description.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of the flavor mappings defined for the corresponding cloud end-point region.
	// A list of the flavor mappings defined for the corresponding cloud end-point region.
	// +kubebuilder:validation:Optional
	FlavorMapping []FlavorMappingParameters `json:"flavorMapping,omitempty" tf:"flavor_mapping,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name used as an identifier in APIs that support this option.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the region for which this profile is defined.
	// The id of the region for which this profile is defined
	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`
}

// FlavorProfileSpec defines the desired state of FlavorProfile
type FlavorProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlavorProfileParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FlavorProfileInitParameters `json:"initProvider,omitempty"`
}

// FlavorProfileStatus defines the observed state of FlavorProfile.
type FlavorProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlavorProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FlavorProfile is the Schema for the FlavorProfiles API. Provides a data lookup for vra_flavor_profile.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type FlavorProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionId) || (has(self.initProvider) && has(self.initProvider.regionId))",message="spec.forProvider.regionId is a required parameter"
	Spec   FlavorProfileSpec   `json:"spec"`
	Status FlavorProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlavorProfileList contains a list of FlavorProfiles
type FlavorProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlavorProfile `json:"items"`
}

// Repository type metadata.
var (
	FlavorProfile_Kind             = "FlavorProfile"
	FlavorProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlavorProfile_Kind}.String()
	FlavorProfile_KindAPIVersion   = FlavorProfile_Kind + "." + CRDGroupVersion.String()
	FlavorProfile_GroupVersionKind = CRDGroupVersion.WithKind(FlavorProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&FlavorProfile{}, &FlavorProfileList{})
}
