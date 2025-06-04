package resources

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_block_device", func(r *config.Resource) {
		r.Kind = "BlockDevice"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_block_device_snapshot", func(r *config.Resource) {
		r.Kind = "BlockDeviceSnapshot"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_blueprint", func(r *config.Resource) {
		r.Kind = "Blueprint"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_blueprint_version", func(r *config.Resource) {
		r.Kind = "BlueprintVersion"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_catalog_item_entitlement", func(r *config.Resource) {
		r.Kind = "CatalogItemEntitlement"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_catalog_item_vm_image", func(r *config.Resource) {
		r.Kind = "CatalogItemVMImage"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_catalog_item_vro_workflow", func(r *config.Resource) {
		r.Kind = "CatalogItemVROWorkflow"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_catalog_source_blueprint", func(r *config.Resource) {
		r.Kind = "CatalogSourceBlueprint"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_catalog_source_entitlement", func(r *config.Resource) {
		r.Kind = "CatalogSourceEntitlement"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_cloud_account_aws", func(r *config.Resource) {
		r.Kind = "AWSCloudAccount"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_cloud_account_azure", func(r *config.Resource) {
		r.Kind = "AzureCloudAccount"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_cloud_account_gcp", func(r *config.Resource) {
		r.Kind = "GCPCloudAccount"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_cloud_account_nsxt", func(r *config.Resource) {
		r.Kind = "NSXTCloudAccount"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_cloud_account_nsxv", func(r *config.Resource) {
		r.Kind = "NSXVCloudAccount"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_cloud_account_vmc", func(r *config.Resource) {
		r.Kind = "VMCCloudAccount"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_cloud_account_vsphere", func(r *config.Resource) {
		r.Kind = "VsphereCloudAccount"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_content_sharing_policy", func(r *config.Resource) {
		r.Kind = "ContentSharingPolicy"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_content_source", func(r *config.Resource) {
		r.Kind = "ContentSource"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_deployment", func(r *config.Resource) {
		r.Kind = "Deployment"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_fabric_compute", func(r *config.Resource) {
		r.Kind = "FabricCompute"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_fabric_datastore_vsphere", func(r *config.Resource) {
		r.Kind = "FabricDatastoreVSphere"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_fabric_network_vsphere", func(r *config.Resource) {
		r.Kind = "FabricNetworkVSphere"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_flavor_profile", func(r *config.Resource) {
		r.Kind = "FlavorProfile"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_image_profile", func(r *config.Resource) {
		r.Kind = "ImageProfile"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_integration", func(r *config.Resource) {
		r.Kind = "Integration"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_load_balancer", func(r *config.Resource) {
		r.Kind = "LoadBalancer"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_machine", func(r *config.Resource) {
		r.Kind = "Machine"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_network", func(r *config.Resource) {
		r.Kind = "Network"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_network_ip_range", func(r *config.Resource) {
		r.Kind = "NetworkIPRange"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_network_profile", func(r *config.Resource) {
		r.Kind = "NetworkProfile"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_policy_approval", func(r *config.Resource) {
		r.Kind = "PolicyApproval"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_policy_day2_action", func(r *config.Resource) {
		r.Kind = "PolicyDay2Action"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_policy_iaas_resource", func(r *config.Resource) {
		r.Kind = "PolicyIaaSResource"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_policy_lease", func(r *config.Resource) {
		r.Kind = "PolicyLease"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_project", func(r *config.Resource) {
		r.Kind = "Project"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_storage_profile", func(r *config.Resource) {
		r.Kind = "StorageProfile"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_storage_profile_aws", func(r *config.Resource) {
		r.Kind = "StorageProfileAWS"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_storage_profile_azure", func(r *config.Resource) {
		r.Kind = "StorageProfileAzure"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_storage_profile_vsphere", func(r *config.Resource) {
		r.Kind = "StorageProfileVSphere"
		r.ShortGroup = "vra"
	})
	p.AddResourceConfigurator("vra_zone", func(r *config.Resource) {
		r.Kind = "Zone"
		r.ShortGroup = "vra"
	})
}
