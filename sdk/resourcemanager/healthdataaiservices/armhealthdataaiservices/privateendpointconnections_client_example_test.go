// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhealthdataaiservices_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
	"log"
)

// Generated from example definition: /mnt/vss/_work/1/s/azure-sdk-for-go/src/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices/TempTypeSpecFiles/HealthDataAIServices.Management/examples/2024-02-28-preview/PrivateEndpointConnections_Create_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "rgopenapi", "deidTest", "kgwgrrpabvrsrrvpcgcnfmyfgyrl", armhealthdataaiservices.PrivateEndpointConnectionResource{
		Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{},
			PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
				Status:          to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
				ActionsRequired: to.Ptr("ulb"),
				Description:     to.Ptr("xr"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhealthdataaiservices.PrivateEndpointConnectionsClientCreateResponse{
	// 	PrivateEndpointConnectionResource: &armhealthdataaiservices.PrivateEndpointConnectionResource{
	// 		Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
	// 				ID: to.Ptr("gpnxxbbtsysdhhclm"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
	// 				ActionsRequired: to.Ptr("ulb"),
	// 				Description: to.Ptr("xr"),
	// 			},
	// 			GroupIDs: []*string{
	// 				to.Ptr("xbdyjqg"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/deidTest/privateEndpointConnections/aduyb"),
	// 		Name: to.Ptr("aduyb"),
	// 		Type: to.Ptr("umjjkodjuhccrngl"),
	// 		SystemData: &armhealthdataaiservices.SystemData{
	// 			CreatedBy: to.Ptr("p"),
	// 			CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 			LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: /mnt/vss/_work/1/s/azure-sdk-for-go/src/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices/TempTypeSpecFiles/HealthDataAIServices.Management/examples/2024-02-28-preview/PrivateEndpointConnections_Delete_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginDelete(ctx, "rgopenapi", "deidTest", "kgwgrrpabvrsrrvpcgcnfmyfgyrl", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: /mnt/vss/_work/1/s/azure-sdk-for-go/src/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices/TempTypeSpecFiles/HealthDataAIServices.Management/examples/2024-02-28-preview/PrivateEndpointConnections_Get_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "rgopenapi", "deidTest", "kgwgrrpabvrsrrvpcgcnfmyfgyrl", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhealthdataaiservices.PrivateEndpointConnectionsClientGetResponse{
	// 	PrivateEndpointConnectionResource: &armhealthdataaiservices.PrivateEndpointConnectionResource{
	// 		Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
	// 				ID: to.Ptr("gpnxxbbtsysdhhclm"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
	// 				ActionsRequired: to.Ptr("ulb"),
	// 				Description: to.Ptr("xr"),
	// 			},
	// 			GroupIDs: []*string{
	// 				to.Ptr("xbdyjqg"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/deidTest/privateEndpointConnections/aduyb"),
	// 		Name: to.Ptr("aduyb"),
	// 		Type: to.Ptr("umjjkodjuhccrngl"),
	// 		SystemData: &armhealthdataaiservices.SystemData{
	// 			CreatedBy: to.Ptr("p"),
	// 			CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 			LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: /mnt/vss/_work/1/s/azure-sdk-for-go/src/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices/TempTypeSpecFiles/HealthDataAIServices.Management/examples/2024-02-28-preview/PrivateEndpointConnections_ListByDeidService_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_NewListByDeidServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByDeidServicePager("rgopenapi", "deidTest", nil)
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
		// page = armhealthdataaiservices.PrivateEndpointConnectionsClientListByDeidServiceResponse{
		// 	PrivateEndpointConnectionResourceListResult: armhealthdataaiservices.PrivateEndpointConnectionResourceListResult{
		// 		Value: []*armhealthdataaiservices.PrivateEndpointConnectionResource{
		// 			{
		// 				Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
		// 						ID: to.Ptr("gpnxxbbtsysdhhclm"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
		// 						Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
		// 						ActionsRequired: to.Ptr("ulb"),
		// 						Description: to.Ptr("mmvcleuufspfrojjveuith"),
		// 					},
		// 					GroupIDs: []*string{
		// 						to.Ptr("xbdyjqg"),
		// 					},
		// 					ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/deidTest/privateEndpointConnections/aduyb"),
		// 				Name: to.Ptr("aduyb"),
		// 				Type: to.Ptr("umjjkodjuhccrngl"),
		// 				SystemData: &armhealthdataaiservices.SystemData{
		// 					CreatedBy: to.Ptr("p"),
		// 					CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
		// 					LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
