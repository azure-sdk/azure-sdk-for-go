// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armiotoperations_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
	"log"
)

// Generated from example definition: 2025-04-01/Instance_CreateOrUpdate_MaximumSet_Gen.json
func ExampleInstanceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstanceClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "aio-instance", armiotoperations.InstanceResource{
		Properties: &armiotoperations.InstanceProperties{
			SchemaRegistryRef: &armiotoperations.SchemaRegistryRef{
				ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"),
			},
			Description: to.Ptr("kpqtgocs"),
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		},
		Identity: &armiotoperations.ManagedServiceIdentity{
			Type:                   to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{},
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("xvewadyhycrjpu"),
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
	// res = armiotoperations.InstanceClientCreateOrUpdateResponse{
	// 	InstanceResource: &armiotoperations.InstanceResource{
	// 		Properties: &armiotoperations.InstanceProperties{
	// 			SchemaRegistryRef: &armiotoperations.SchemaRegistryRef{
	// 				ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"),
	// 			},
	// 			Description: to.Ptr("kpqtgocs"),
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 			Version: to.Ptr("vjjbmunthiphfmekvxgxcxkzdwjti"),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		Identity: &armiotoperations.ManagedServiceIdentity{
	// 			Type: to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("4a6e4195-75b8-4685-aa0c-0b5704779327"),
	// 			TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("xvewadyhycrjpu"),
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123"),
	// 		Name: to.Ptr("llptmlifnqqwairx"),
	// 		Type: to.Ptr("qwrfzxjfxvismlqvigot"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-04-01/Instance_Delete_MaximumSet_Gen.json
func ExampleInstanceClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstanceClient().BeginDelete(ctx, "rgiotoperations", "aio-instance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2025-04-01/Instance_Get_MaximumSet_Gen.json
func ExampleInstanceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstanceClient().Get(ctx, "rgiotoperations", "aio-instance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.InstanceClientGetResponse{
	// 	InstanceResource: &armiotoperations.InstanceResource{
	// 		Properties: &armiotoperations.InstanceProperties{
	// 			SchemaRegistryRef: &armiotoperations.SchemaRegistryRef{
	// 				ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"),
	// 			},
	// 			Description: to.Ptr("rlfvvnnhcypp"),
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 			Version: to.Ptr("vjjbmunthiphfmekvxgxcxkzdwjti"),
	// 			Features: map[string]*armiotoperations.InstanceFeature{
	// 				"dataFlows": &armiotoperations.InstanceFeature{
	// 					Mode: to.Ptr(armiotoperations.InstanceFeatureModeDisabled),
	// 				},
	// 				"akri": &armiotoperations.InstanceFeature{
	// 					Mode: to.Ptr(armiotoperations.InstanceFeatureModePreview),
	// 				},
	// 				"mqttBroker": &armiotoperations.InstanceFeature{
	// 					Settings: map[string]*armiotoperations.OperationalMode{
	// 						"preview": to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		Identity: &armiotoperations.ManagedServiceIdentity{
	// 			Type: to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("4a6e4195-75b8-4685-aa0c-0b5704779327"),
	// 			TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("xvewadyhycrjpu"),
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123"),
	// 		Name: to.Ptr("llptmlifnqqwairx"),
	// 		Type: to.Ptr("qwrfzxjfxvismlqvigot"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-04-01/Instance_ListByResourceGroup_MaximumSet_Gen.json
func ExampleInstanceClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstanceClient().NewListByResourceGroupPager("rgiotoperations", nil)
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
		// page = armiotoperations.InstanceClientListByResourceGroupResponse{
		// 	InstanceResourceListResult: armiotoperations.InstanceResourceListResult{
		// 		Value: []*armiotoperations.InstanceResource{
		// 			{
		// 				Properties: &armiotoperations.InstanceProperties{
		// 					ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
		// 					Version: to.Ptr("vjjbmunthiphfmekvxgxcxkzdwjti"),
		// 					SchemaRegistryRef: &armiotoperations.SchemaRegistryRef{
		// 						ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"),
		// 					},
		// 					Description: to.Ptr("vmujggxdvxk"),
		// 				},
		// 				ExtendedLocation: &armiotoperations.ExtendedLocation{
		// 					Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
		// 					Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		// 				},
		// 				Identity: &armiotoperations.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("4a6e4195-75b8-4685-aa0c-0b5704779327"),
		// 					TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
		// 					Type: to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("xvewadyhycrjpu"),
		// 				ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123"),
		// 				Name: to.Ptr("llptmlifnqqwairx"),
		// 				Type: to.Ptr("qwrfzxjfxvismlqvigot"),
		// 				SystemData: &armiotoperations.SystemData{
		// 					CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
		// 					CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("gnicpuszwd"),
		// 					LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2025-04-01/Instance_ListBySubscription_MaximumSet_Gen.json
func ExampleInstanceClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstanceClient().NewListBySubscriptionPager(nil)
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
		// page = armiotoperations.InstanceClientListBySubscriptionResponse{
		// 	InstanceResourceListResult: armiotoperations.InstanceResourceListResult{
		// 		Value: []*armiotoperations.InstanceResource{
		// 			{
		// 				Properties: &armiotoperations.InstanceProperties{
		// 					ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
		// 					Version: to.Ptr("vjjbmunthiphfmekvxgxcxkzdwjti"),
		// 					SchemaRegistryRef: &armiotoperations.SchemaRegistryRef{
		// 						ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"),
		// 					},
		// 					Description: to.Ptr("empgqmbhvklcqlyahmdsjemlep"),
		// 				},
		// 				ExtendedLocation: &armiotoperations.ExtendedLocation{
		// 					Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
		// 					Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		// 				},
		// 				Identity: &armiotoperations.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("4a6e4195-75b8-4685-aa0c-0b5704779327"),
		// 					TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
		// 					Type: to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("xvewadyhycrjpu"),
		// 				ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123"),
		// 				Name: to.Ptr("llptmlifnqqwairx"),
		// 				Type: to.Ptr("qwrfzxjfxvismlqvigot"),
		// 				SystemData: &armiotoperations.SystemData{
		// 					CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
		// 					CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("gnicpuszwd"),
		// 					LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2025-04-01/Instance_Update_MaximumSet_Gen.json
func ExampleInstanceClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstanceClient().Update(ctx, "rgiotoperations", "aio-instance", armiotoperations.InstancePatchModel{
		Tags: map[string]*string{},
		Identity: &armiotoperations.ManagedServiceIdentity{
			Type:                   to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.InstanceClientUpdateResponse{
	// 	InstanceResource: &armiotoperations.InstanceResource{
	// 		Properties: &armiotoperations.InstanceProperties{
	// 			SchemaRegistryRef: &armiotoperations.SchemaRegistryRef{
	// 				ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"),
	// 			},
	// 			Description: to.Ptr("wwihkapmgjbyrtyaj"),
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 			Version: to.Ptr("vjjbmunthiphfmekvxgxcxkzdwjti"),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		Identity: &armiotoperations.ManagedServiceIdentity{
	// 			Type: to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("4a6e4195-75b8-4685-aa0c-0b5704779327"),
	// 			TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("xvewadyhycrjpu"),
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123"),
	// 		Name: to.Ptr("llptmlifnqqwairx"),
	// 		Type: to.Ptr("qwrfzxjfxvismlqvigot"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}
