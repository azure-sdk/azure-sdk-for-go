// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armservicefabricmanagedclusters_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
	"log"
)

// Generated from example definition: 2025-03-01-preview/managedUnsupportedVMSizesGet_example.json
func ExampleManagedUnsupportedVMSizesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedUnsupportedVMSizesClient().Get(ctx, "eastus", "Standard_B1ls1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientGetResponse{
	// 	ManagedVMSize: &armservicefabricmanagedclusters.ManagedVMSize{
	// 		Name: to.Ptr("Standard_B1ls1"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
	// 		ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B1ls1"),
	// 		Properties: &armservicefabricmanagedclusters.VMSize{
	// 			Size: to.Ptr("Standard_B1ls1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-01-preview/managedUnsupportedVMSizesList_example.json
func ExampleManagedUnsupportedVMSizesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedUnsupportedVMSizesClient().NewListPager("eastus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armservicefabricmanagedclusters.ManagedUnsupportedVMSizesClientListResponse{
		// 	ManagedVMSizesResult: armservicefabricmanagedclusters.ManagedVMSizesResult{
		// 		Value: []*armservicefabricmanagedclusters.ManagedVMSize{
		// 			{
		// 				Name: to.Ptr("Standard_B1ls1"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B1ls1"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_B1ls1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_B1s"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B1s"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_B1s"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_B1ms"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B1ms"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_B1ms"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_B2s"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B2s"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_B2s"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_B2ms"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B2ms"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_B2ms"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_B4ms"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B4ms"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_B4ms"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_B8ms"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_B8ms"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_B8ms"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A1_v2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A1_v2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A1_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A2_v2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A2_v2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A2_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A4_v2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A4_v2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A4_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A8_v2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A8_v2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A8_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A2m_v2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A2m_v2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A2m_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A4m_v2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A4m_v2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A4m_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A8m_v2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A8m_v2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A8m_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Basic_A0"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Basic_A0"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Basic_A0"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Basic_A1"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Basic_A1"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Basic_A1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Basic_A2"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Basic_A2"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Basic_A2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Basic_A3"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Basic_A3"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Basic_A3"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Basic_A4"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Basic_A4"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Basic_A4"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_A0"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/locations/managedVMSizes"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/managedVMSizes/Standard_A0"),
		// 				Properties: &armservicefabricmanagedclusters.VMSize{
		// 					Size: to.Ptr("Standard_A0"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
