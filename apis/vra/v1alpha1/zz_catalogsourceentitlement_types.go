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

type CatalogSourceEntitlementDefinitionInitParameters struct {
}

type CatalogSourceEntitlementDefinitionObservation struct {

	// Description of the catalog source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Id of the catalog source.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Icon id of associated catalog source.
	IconID *string `json:"iconId,omitempty" tf:"icon_id,omitempty"`

	// Name of the catalog source.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of items in the associated catalog source.
	NumberOfItems *float64 `json:"numberOfItems,omitempty" tf:"number_of_items,omitempty"`

	// Catalog source name.
	SourceName *string `json:"sourceName,omitempty" tf:"source_name,omitempty"`

	// Catalog source type.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Content definition type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CatalogSourceEntitlementDefinitionParameters struct {
}

type CatalogSourceEntitlementInitParameters struct {

	// The id of the catalog source to create the entitlement.
	// Catalog source id.
	CatalogSourceID *string `json:"catalogSourceId,omitempty" tf:"catalog_source_id,omitempty"`

	// The id of the project this entity belongs to.
	// Project id.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type CatalogSourceEntitlementObservation struct {

	// The id of the catalog source to create the entitlement.
	// Catalog source id.
	CatalogSourceID *string `json:"catalogSourceId,omitempty" tf:"catalog_source_id,omitempty"`

	// Represents a catalog source that is linked to a project via an entitlement.
	Definition []CatalogSourceEntitlementDefinitionObservation `json:"definition,omitempty" tf:"definition,omitempty"`

	// Id of the catalog source.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of the project this entity belongs to.
	// Project id.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type CatalogSourceEntitlementParameters struct {

	// The id of the catalog source to create the entitlement.
	// Catalog source id.
	// +kubebuilder:validation:Optional
	CatalogSourceID *string `json:"catalogSourceId,omitempty" tf:"catalog_source_id,omitempty"`

	// The id of the project this entity belongs to.
	// Project id.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// CatalogSourceEntitlementSpec defines the desired state of CatalogSourceEntitlement
type CatalogSourceEntitlementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogSourceEntitlementParameters `json:"forProvider"`
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
	InitProvider CatalogSourceEntitlementInitParameters `json:"initProvider,omitempty"`
}

// CatalogSourceEntitlementStatus defines the observed state of CatalogSourceEntitlement.
type CatalogSourceEntitlementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogSourceEntitlementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CatalogSourceEntitlement is the Schema for the CatalogSourceEntitlements API. A resource that can be used to create a VMware Aria Automation catalog source entitlement.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type CatalogSourceEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.catalogSourceId) || (has(self.initProvider) && has(self.initProvider.catalogSourceId))",message="spec.forProvider.catalogSourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	Spec   CatalogSourceEntitlementSpec   `json:"spec"`
	Status CatalogSourceEntitlementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogSourceEntitlementList contains a list of CatalogSourceEntitlements
type CatalogSourceEntitlementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogSourceEntitlement `json:"items"`
}

// Repository type metadata.
var (
	CatalogSourceEntitlement_Kind             = "CatalogSourceEntitlement"
	CatalogSourceEntitlement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogSourceEntitlement_Kind}.String()
	CatalogSourceEntitlement_KindAPIVersion   = CatalogSourceEntitlement_Kind + "." + CRDGroupVersion.String()
	CatalogSourceEntitlement_GroupVersionKind = CRDGroupVersion.WithKind(CatalogSourceEntitlement_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogSourceEntitlement{}, &CatalogSourceEntitlementList{})
}
