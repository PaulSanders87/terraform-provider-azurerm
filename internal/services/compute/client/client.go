package client

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/availabilitysets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/dedicatedhostgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/dedicatedhosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/sshpublickeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-01/capacityreservationgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-01/capacityreservations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-01/images"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-01/proximityplacementgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/diskaccesses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/diskencryptionsets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/disks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/snapshots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleryapplications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleryapplicationversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/marketplaceordering/2015-06-01/agreements"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
	"github.com/tombuildsstuff/kermit/sdk/compute/2023-03-01/compute"
)

type Client struct {
	AvailabilitySetsClient           *availabilitysets.AvailabilitySetsClient
	CapacityReservationsClient       *capacityreservations.CapacityReservationsClient
	CapacityReservationGroupsClient  *capacityreservationgroups.CapacityReservationGroupsClient
	DedicatedHostsClient             *dedicatedhosts.DedicatedHostsClient
	DedicatedHostGroupsClient        *dedicatedhostgroups.DedicatedHostGroupsClient
	DisksClient                      *disks.DisksClient
	DiskAccessClient                 *diskaccesses.DiskAccessesClient
	DiskEncryptionSetsClient         *diskencryptionsets.DiskEncryptionSetsClient
	GalleriesClient                  *galleries.GalleriesClient
	GalleryApplicationsClient        *galleryapplications.GalleryApplicationsClient
	GalleryApplicationVersionsClient *galleryapplicationversions.GalleryApplicationVersionsClient
	GalleryImagesClient              *compute.GalleryImagesClient
	GalleryImageVersionsClient       *compute.GalleryImageVersionsClient
	ImagesClient                     *images.ImagesClient
	MarketplaceAgreementsClient      *agreements.AgreementsClient
	ProximityPlacementGroupsClient   *proximityplacementgroups.ProximityPlacementGroupsClient
	SkusClient                       *skus.SkusClient
	SSHPublicKeysClient              *sshpublickeys.SshPublicKeysClient
	SnapshotsClient                  *snapshots.SnapshotsClient
	VirtualMachinesClient            *virtualmachines.VirtualMachinesClient
	VMExtensionImageClient           *compute.VirtualMachineExtensionImagesClient
	VMExtensionClient                *compute.VirtualMachineExtensionsClient
	VMScaleSetClient                 *compute.VirtualMachineScaleSetsClient
	VMScaleSetExtensionsClient       *compute.VirtualMachineScaleSetExtensionsClient
	VMScaleSetRollingUpgradesClient  *compute.VirtualMachineScaleSetRollingUpgradesClient
	VMScaleSetVMsClient              *compute.VirtualMachineScaleSetVMsClient
	VMClient                         *compute.VirtualMachinesClient
	VMImageClient                    *compute.VirtualMachineImagesClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	availabilitySetsClient := availabilitysets.NewAvailabilitySetsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&availabilitySetsClient.Client, o.ResourceManagerAuthorizer)

	capacityReservationsClient := capacityreservations.NewCapacityReservationsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&capacityReservationsClient.Client, o.ResourceManagerAuthorizer)

	capacityReservationGroupsClient := capacityreservationgroups.NewCapacityReservationGroupsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&capacityReservationGroupsClient.Client, o.ResourceManagerAuthorizer)

	dedicatedHostsClient := dedicatedhosts.NewDedicatedHostsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&dedicatedHostsClient.Client, o.ResourceManagerAuthorizer)

	dedicatedHostGroupsClient := dedicatedhostgroups.NewDedicatedHostGroupsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&dedicatedHostGroupsClient.Client, o.ResourceManagerAuthorizer)

	disksClient := disks.NewDisksClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&disksClient.Client, o.ResourceManagerAuthorizer)

	diskAccessClient := diskaccesses.NewDiskAccessesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&diskAccessClient.Client, o.ResourceManagerAuthorizer)

	diskEncryptionSetsClient := diskencryptionsets.NewDiskEncryptionSetsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&diskEncryptionSetsClient.Client, o.ResourceManagerAuthorizer)

	galleriesClient := galleries.NewGalleriesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&galleriesClient.Client, o.ResourceManagerAuthorizer)

	galleryApplicationsClient := galleryapplications.NewGalleryApplicationsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&galleryApplicationsClient.Client, o.ResourceManagerAuthorizer)

	galleryApplicationVersionsClient := galleryapplicationversions.NewGalleryApplicationVersionsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&galleryApplicationVersionsClient.Client, o.ResourceManagerAuthorizer)

	galleryImagesClient := compute.NewGalleryImagesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&galleryImagesClient.Client, o.ResourceManagerAuthorizer)

	galleryImageVersionsClient := compute.NewGalleryImageVersionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&galleryImageVersionsClient.Client, o.ResourceManagerAuthorizer)

	imagesClient := images.NewImagesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&imagesClient.Client, o.ResourceManagerAuthorizer)

	marketplaceAgreementsClient, err := agreements.NewAgreementsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building MarketplaceAgreementsClient client: %+v", err)
	}
	o.Configure(marketplaceAgreementsClient.Client, o.Authorizers.ResourceManager)

	proximityPlacementGroupsClient := proximityplacementgroups.NewProximityPlacementGroupsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&proximityPlacementGroupsClient.Client, o.ResourceManagerAuthorizer)

	skusClient := skus.NewSkusClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&skusClient.Client, o.ResourceManagerAuthorizer)

	snapshotsClient := snapshots.NewSnapshotsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&snapshotsClient.Client, o.ResourceManagerAuthorizer)

	sshPublicKeysClient := sshpublickeys.NewSshPublicKeysClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&sshPublicKeysClient.Client, o.ResourceManagerAuthorizer)

	usageClient := compute.NewUsageClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&usageClient.Client, o.ResourceManagerAuthorizer)

	virtualMachinesClient := virtualmachines.NewVirtualMachinesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&virtualMachinesClient.Client, o.ResourceManagerAuthorizer)

	vmExtensionImageClient := compute.NewVirtualMachineExtensionImagesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmExtensionImageClient.Client, o.ResourceManagerAuthorizer)

	vmExtensionClient := compute.NewVirtualMachineExtensionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmExtensionClient.Client, o.ResourceManagerAuthorizer)

	vmImageClient := compute.NewVirtualMachineImagesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmImageClient.Client, o.ResourceManagerAuthorizer)

	vmScaleSetClient := compute.NewVirtualMachineScaleSetsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmScaleSetClient.Client, o.ResourceManagerAuthorizer)

	vmScaleSetExtensionsClient := compute.NewVirtualMachineScaleSetExtensionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmScaleSetExtensionsClient.Client, o.ResourceManagerAuthorizer)

	vmScaleSetRollingUpgradesClient := compute.NewVirtualMachineScaleSetRollingUpgradesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmScaleSetRollingUpgradesClient.Client, o.ResourceManagerAuthorizer)

	vmScaleSetVMsClient := compute.NewVirtualMachineScaleSetVMsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmScaleSetVMsClient.Client, o.ResourceManagerAuthorizer)

	vmClient := compute.NewVirtualMachinesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vmClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		AvailabilitySetsClient:           &availabilitySetsClient,
		CapacityReservationsClient:       &capacityReservationsClient,
		CapacityReservationGroupsClient:  &capacityReservationGroupsClient,
		DedicatedHostsClient:             &dedicatedHostsClient,
		DedicatedHostGroupsClient:        &dedicatedHostGroupsClient,
		DisksClient:                      &disksClient,
		DiskAccessClient:                 &diskAccessClient,
		DiskEncryptionSetsClient:         &diskEncryptionSetsClient,
		GalleriesClient:                  &galleriesClient,
		GalleryApplicationsClient:        &galleryApplicationsClient,
		GalleryApplicationVersionsClient: &galleryApplicationVersionsClient,
		GalleryImagesClient:              &galleryImagesClient,
		GalleryImageVersionsClient:       &galleryImageVersionsClient,
		ImagesClient:                     &imagesClient,
		MarketplaceAgreementsClient:      marketplaceAgreementsClient,
		ProximityPlacementGroupsClient:   &proximityPlacementGroupsClient,
		SkusClient:                       &skusClient,
		SSHPublicKeysClient:              &sshPublicKeysClient,
		SnapshotsClient:                  &snapshotsClient,
		VirtualMachinesClient:            &virtualMachinesClient,
		VMExtensionImageClient:           &vmExtensionImageClient,
		VMExtensionClient:                &vmExtensionClient,
		VMScaleSetClient:                 &vmScaleSetClient,
		VMScaleSetExtensionsClient:       &vmScaleSetExtensionsClient,
		VMScaleSetRollingUpgradesClient:  &vmScaleSetRollingUpgradesClient,
		VMScaleSetVMsClient:              &vmScaleSetVMsClient,
		VMImageClient:                    &vmImageClient,

		// NOTE: use `VirtualMachinesClient` instead
		VMClient: &vmClient,
	}, nil
}
