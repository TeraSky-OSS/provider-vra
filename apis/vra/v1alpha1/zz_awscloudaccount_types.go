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

type AWSCloudAccountInitParameters struct {

	// Access key ID for AWS.
	// Aws Access key ID.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of AWS cloud account.
	// The name of this resource instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of region names enabled for the cloud account.
	// The set of region ids that will be enabled for this cloud account.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// AWS Secret Access Key
	// Aws Secret Access Key.
	SecretKeySecretRef v1.SecretKeySelector `json:"secretKeySecretRef" tf:"-"`

	// Set of tag keys and values to apply to the cloud account. Example: [ { "key" : "vmware", "value": "provider" } ]
	Tags []AWSCloudAccountTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AWSCloudAccountLinksInitParameters struct {
}

type AWSCloudAccountLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type AWSCloudAccountLinksParameters struct {
}

type AWSCloudAccountObservation struct {

	// Access key ID for AWS.
	// Aws Access key ID.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of AWS cloud account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Hypermedia as the Engine of Application State (HATEOAS) of entity.
	Links []AWSCloudAccountLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// Name of AWS cloud account.
	// The name of this resource instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of organization that entity belongs to.
	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of entity owner.
	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Set of region names enabled for the cloud account.
	// The set of region ids that will be enabled for this cloud account.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// Set of tag keys and values to apply to the cloud account. Example: [ { "key" : "vmware", "value": "provider" } ]
	Tags []AWSCloudAccountTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type AWSCloudAccountParameters struct {

	// Access key ID for AWS.
	// Aws Access key ID.
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Human-friendly description.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of AWS cloud account.
	// The name of this resource instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of region names enabled for the cloud account.
	// The set of region ids that will be enabled for this cloud account.
	// +kubebuilder:validation:Optional
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// AWS Secret Access Key
	// Aws Secret Access Key.
	// +kubebuilder:validation:Optional
	SecretKeySecretRef v1.SecretKeySelector `json:"secretKeySecretRef" tf:"-"`

	// Set of tag keys and values to apply to the cloud account. Example: [ { "key" : "vmware", "value": "provider" } ]
	// +kubebuilder:validation:Optional
	Tags []AWSCloudAccountTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AWSCloudAccountTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AWSCloudAccountTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AWSCloudAccountTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// AWSCloudAccountSpec defines the desired state of AWSCloudAccount
type AWSCloudAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AWSCloudAccountParameters `json:"forProvider"`
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
	InitProvider AWSCloudAccountInitParameters `json:"initProvider,omitempty"`
}

// AWSCloudAccountStatus defines the observed state of AWSCloudAccount.
type AWSCloudAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AWSCloudAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AWSCloudAccount is the Schema for the AWSCloudAccounts API. Creates a vra_cloud_account_aws resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type AWSCloudAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessKey) || (has(self.initProvider) && has(self.initProvider.accessKey))",message="spec.forProvider.accessKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regions) || (has(self.initProvider) && has(self.initProvider.regions))",message="spec.forProvider.regions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.secretKeySecretRef)",message="spec.forProvider.secretKeySecretRef is a required parameter"
	Spec   AWSCloudAccountSpec   `json:"spec"`
	Status AWSCloudAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSCloudAccountList contains a list of AWSCloudAccounts
type AWSCloudAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSCloudAccount `json:"items"`
}

// Repository type metadata.
var (
	AWSCloudAccount_Kind             = "AWSCloudAccount"
	AWSCloudAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AWSCloudAccount_Kind}.String()
	AWSCloudAccount_KindAPIVersion   = AWSCloudAccount_Kind + "." + CRDGroupVersion.String()
	AWSCloudAccount_GroupVersionKind = CRDGroupVersion.WithKind(AWSCloudAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&AWSCloudAccount{}, &AWSCloudAccountList{})
}
