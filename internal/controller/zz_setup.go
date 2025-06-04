// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/terasky/provider-vra/internal/controller/providerconfig"
	awscloudaccount "github.com/terasky/provider-vra/internal/controller/vra/awscloudaccount"
	azurecloudaccount "github.com/terasky/provider-vra/internal/controller/vra/azurecloudaccount"
	blockdevice "github.com/terasky/provider-vra/internal/controller/vra/blockdevice"
	blockdevicesnapshot "github.com/terasky/provider-vra/internal/controller/vra/blockdevicesnapshot"
	blueprint "github.com/terasky/provider-vra/internal/controller/vra/blueprint"
	blueprintversion "github.com/terasky/provider-vra/internal/controller/vra/blueprintversion"
	catalogitementitlement "github.com/terasky/provider-vra/internal/controller/vra/catalogitementitlement"
	catalogitemvmimage "github.com/terasky/provider-vra/internal/controller/vra/catalogitemvmimage"
	catalogitemvroworkflow "github.com/terasky/provider-vra/internal/controller/vra/catalogitemvroworkflow"
	catalogsourceblueprint "github.com/terasky/provider-vra/internal/controller/vra/catalogsourceblueprint"
	catalogsourceentitlement "github.com/terasky/provider-vra/internal/controller/vra/catalogsourceentitlement"
	contentsharingpolicy "github.com/terasky/provider-vra/internal/controller/vra/contentsharingpolicy"
	contentsource "github.com/terasky/provider-vra/internal/controller/vra/contentsource"
	deployment "github.com/terasky/provider-vra/internal/controller/vra/deployment"
	fabriccompute "github.com/terasky/provider-vra/internal/controller/vra/fabriccompute"
	fabricdatastorevsphere "github.com/terasky/provider-vra/internal/controller/vra/fabricdatastorevsphere"
	fabricnetworkvsphere "github.com/terasky/provider-vra/internal/controller/vra/fabricnetworkvsphere"
	flavorprofile "github.com/terasky/provider-vra/internal/controller/vra/flavorprofile"
	gcpcloudaccount "github.com/terasky/provider-vra/internal/controller/vra/gcpcloudaccount"
	imageprofile "github.com/terasky/provider-vra/internal/controller/vra/imageprofile"
	integration "github.com/terasky/provider-vra/internal/controller/vra/integration"
	loadbalancer "github.com/terasky/provider-vra/internal/controller/vra/loadbalancer"
	machine "github.com/terasky/provider-vra/internal/controller/vra/machine"
	network "github.com/terasky/provider-vra/internal/controller/vra/network"
	networkiprange "github.com/terasky/provider-vra/internal/controller/vra/networkiprange"
	networkprofile "github.com/terasky/provider-vra/internal/controller/vra/networkprofile"
	nsxtcloudaccount "github.com/terasky/provider-vra/internal/controller/vra/nsxtcloudaccount"
	nsxvcloudaccount "github.com/terasky/provider-vra/internal/controller/vra/nsxvcloudaccount"
	policyapproval "github.com/terasky/provider-vra/internal/controller/vra/policyapproval"
	policyday2action "github.com/terasky/provider-vra/internal/controller/vra/policyday2action"
	policyiaasresource "github.com/terasky/provider-vra/internal/controller/vra/policyiaasresource"
	policylease "github.com/terasky/provider-vra/internal/controller/vra/policylease"
	project "github.com/terasky/provider-vra/internal/controller/vra/project"
	storageprofile "github.com/terasky/provider-vra/internal/controller/vra/storageprofile"
	storageprofileaws "github.com/terasky/provider-vra/internal/controller/vra/storageprofileaws"
	storageprofileazure "github.com/terasky/provider-vra/internal/controller/vra/storageprofileazure"
	storageprofilevsphere "github.com/terasky/provider-vra/internal/controller/vra/storageprofilevsphere"
	vmccloudaccount "github.com/terasky/provider-vra/internal/controller/vra/vmccloudaccount"
	vspherecloudaccount "github.com/terasky/provider-vra/internal/controller/vra/vspherecloudaccount"
	zone "github.com/terasky/provider-vra/internal/controller/vra/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		awscloudaccount.Setup,
		azurecloudaccount.Setup,
		blockdevice.Setup,
		blockdevicesnapshot.Setup,
		blueprint.Setup,
		blueprintversion.Setup,
		catalogitementitlement.Setup,
		catalogitemvmimage.Setup,
		catalogitemvroworkflow.Setup,
		catalogsourceblueprint.Setup,
		catalogsourceentitlement.Setup,
		contentsharingpolicy.Setup,
		contentsource.Setup,
		deployment.Setup,
		fabriccompute.Setup,
		fabricdatastorevsphere.Setup,
		fabricnetworkvsphere.Setup,
		flavorprofile.Setup,
		gcpcloudaccount.Setup,
		imageprofile.Setup,
		integration.Setup,
		loadbalancer.Setup,
		machine.Setup,
		network.Setup,
		networkiprange.Setup,
		networkprofile.Setup,
		nsxtcloudaccount.Setup,
		nsxvcloudaccount.Setup,
		policyapproval.Setup,
		policyday2action.Setup,
		policyiaasresource.Setup,
		policylease.Setup,
		project.Setup,
		storageprofile.Setup,
		storageprofileaws.Setup,
		storageprofileazure.Setup,
		storageprofilevsphere.Setup,
		vmccloudaccount.Setup,
		vspherecloudaccount.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
