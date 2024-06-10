//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_Get_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_Get_virtualMachineImageGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().Get(ctx, "aaaaaa", "aaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImage = armcompute.VirtualMachineImage{
	// 	ID: to.Ptr("aaaaaaaaaaa"),
	// 	Name: to.Ptr("aaaaaaaaa"),
	// 	ExtendedLocation: &armcompute.ExtendedLocation{
	// 		Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 		Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 	},
	// 	Location: to.Ptr("aaaaa"),
	// 	Tags: map[string]*string{
	// 		"key6817": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineImageProperties{
	// 		AutomaticOSUpgradeProperties: &armcompute.AutomaticOSUpgradeProperties{
	// 			AutomaticOSUpgradeSupported: to.Ptr(true),
	// 		},
	// 		DataDiskImages: []*armcompute.DataDiskImage{
	// 			{
	// 				Lun: to.Ptr[int32](17),
	// 		}},
	// 		Disallowed: &armcompute.DisallowedConfiguration{
	// 			VMDiskType: to.Ptr(armcompute.VMDiskTypesNone),
	// 		},
	// 		Features: []*armcompute.VirtualMachineImageFeature{
	// 			{
	// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 				Value: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		}},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypesV1),
	// 		ImageDeprecationStatus: &armcompute.ImageDeprecationStatus{
	// 			AlternativeOption: &armcompute.AlternativeOption{
	// 				Type: to.Ptr(armcompute.AlternativeTypeOffer),
	// 				Value: to.Ptr("aaaaaaa"),
	// 			},
	// 			ImageState: to.Ptr(armcompute.ImageStateScheduledForDeprecation),
	// 			ScheduledDeprecationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-13T00:00:00.000Z"); return t}()),
	// 		},
	// 		OSDiskImage: &armcompute.OSDiskImage{
	// 			OperatingSystem: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		},
	// 		Plan: &armcompute.PurchasePlan{
	// 			Name: to.Ptr("aaaaaaaaa"),
	// 			Product: to.Ptr("aaaaaaaaaaaaaa"),
	// 			Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_Get_MinimumSet_Gen.json
func ExampleVirtualMachineImagesClient_Get_virtualMachineImageGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().Get(ctx, "aaaaaaaaaaaa", "aaaaaaaaaaa", "aa", "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImage = armcompute.VirtualMachineImage{
	// 	ID: to.Ptr("aaaaaaaaaaa"),
	// 	Name: to.Ptr("aaaaaaaaa"),
	// 	Location: to.Ptr("aaaaa"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_List_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_List_virtualMachineImageListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().List(ctx, "aaaaaaaaaaaaaaa", "aaaaaa", "aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineImagesClientListOptions{Expand: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		Top:     to.Ptr[int32](18),
		Orderby: to.Ptr("aa"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_List_MinimumSet_Gen.json
func ExampleVirtualMachineImagesClient_List_virtualMachineImageListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().List(ctx, "aaaaaaa", "aaaaaaaaaaa", "aaaaaaaaaa", "aaaaaa", &armcompute.VirtualMachineImagesClientListOptions{Expand: nil,
		Top:     nil,
		Orderby: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListOffers_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListOffers_virtualMachineImageListOffersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListOffers(ctx, "aaaaaaa", "aaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListOffers_MinimumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListOffers_virtualMachineImageListOffersMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListOffers(ctx, "aaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListPublishers_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListPublishers_virtualMachineImageListPublishersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListPublishers(ctx, "aaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListPublishers_MinimumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListPublishers_virtualMachineImageListPublishersMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListPublishers(ctx, "aaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListSkus_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListSKUs_virtualMachineImageListSkusMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListSKUs(ctx, "aaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListSkus_MinimumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListSKUs_virtualMachineImageListSkusMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListSKUs(ctx, "aaaa", "aaaaaaaaaaaaa", "aaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListByEdgeZone_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListByEdgeZone_virtualMachineImagesEdgeZoneListByEdgeZoneMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListByEdgeZone(ctx, "WestUS", "microsoftlosangeles1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMImagesInEdgeZoneListResult = armcompute.VMImagesInEdgeZoneListResult{
	// 	Value: []*armcompute.VirtualMachineImageResource{
	// 		{
	// 			ID: to.Ptr("/Subscriptions/5ece5940-d962-4dad-a98f-ca9ac0f021a5/Providers/Microsoft.Compute/Locations/westus/Publishers/CANONICAL/ArtifactTypes/VMImage/Offers/UBUNTUSERVER/Skus/18_04-LTS-GEN2/Versions/18.04.202107200"),
	// 			Name: to.Ptr("18.04.202107200"),
	// 			ExtendedLocation: &armcompute.ExtendedLocation{
	// 				Name: to.Ptr("microsoftlosangeles1"),
	// 				Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 			},
	// 			Location: to.Ptr("WestUS"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListByEdgeZone_MinimumSet_Gen.json
func ExampleVirtualMachineImagesClient_ListByEdgeZone_virtualMachineImagesEdgeZoneListByEdgeZoneMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().ListByEdgeZone(ctx, "WestUS", "microsoftlosangeles1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMImagesInEdgeZoneListResult = armcompute.VMImagesInEdgeZoneListResult{
	// 	Value: []*armcompute.VirtualMachineImageResource{
	// 	},
	// }
}