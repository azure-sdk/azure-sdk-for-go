// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasefleetmanager_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasefleetmanager/armdatabasefleetmanager"
	"log"
)

// Generated from example definition: 2025-02-01-preview/Fleets_CreateOrUpdate_MaximumSet_Gen.json
func ExampleFleetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasefleetmanager.NewClientFactory("C3897315-3847-4D8A-B2FC-7307B066AD63", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginCreateOrUpdate(ctx, "rg-database-fleet-manager", "production-fleet-01", armdatabasefleetmanager.Fleet{
		Properties: &armdatabasefleetmanager.FleetProperties{
			Description: to.Ptr("Production fleet for high availability and scalability."),
		},
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
			"owner":       to.Ptr("team-database"),
		},
		Location: to.Ptr("East US"),
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
	// res = armdatabasefleetmanager.FleetsClientCreateOrUpdateResponse{
	// 	Fleet: &armdatabasefleetmanager.Fleet{
	// 		Properties: &armdatabasefleetmanager.FleetProperties{
	// 			Description: to.Ptr("Production fleet for high availability and scalability."),
	// 			ProvisioningState: to.Ptr(armdatabasefleetmanager.AzureProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"owner": to.Ptr("team-database"),
	// 		},
	// 		Location: to.Ptr("East US"),
	// 		ID: to.Ptr("/subscriptions/C3897315-3847-4D8A-B2FC-7307B066AD63/resourceGroups/rg-database-fleet-manager/providers/Microsoft.DatabaseFleetManager/fleet/production-fleet-01"),
	// 		Name: to.Ptr("production-fleet-01"),
	// 		Type: to.Ptr("Microsoft.DatabaseFleetManager/fleets"),
	// 		SystemData: &armdatabasefleetmanager.SystemData{
	// 			CreatedBy: to.Ptr("adminuser"),
	// 			CreatedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.802Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("adminuser"),
	// 			LastModifiedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.803Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-02-01-preview/Fleets_Delete_MaximumSet_Gen.json
func ExampleFleetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasefleetmanager.NewClientFactory("C3897315-3847-4D8A-B2FC-7307B066AD63", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginDelete(ctx, "rg-database-fleet-manager", "production-fleet-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2025-02-01-preview/Fleets_Get_MaximumSet_Gen.json
func ExampleFleetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasefleetmanager.NewClientFactory("C3897315-3847-4D8A-B2FC-7307B066AD63", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetsClient().Get(ctx, "rg-database-fleet-manager", "production-fleet-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasefleetmanager.FleetsClientGetResponse{
	// 	Fleet: &armdatabasefleetmanager.Fleet{
	// 		Properties: &armdatabasefleetmanager.FleetProperties{
	// 			Description: to.Ptr("Production fleet for high availability and scalability."),
	// 			ProvisioningState: to.Ptr(armdatabasefleetmanager.AzureProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"owner": to.Ptr("team-database"),
	// 		},
	// 		Location: to.Ptr("East US"),
	// 		ID: to.Ptr("/subscriptions/C3897315-3847-4D8A-B2FC-7307B066AD63/resourceGroups/rg-database-fleet-manager/providers/Microsoft.DatabaseFleetManager/fleet/production-fleet-01"),
	// 		Name: to.Ptr("production-fleet-01"),
	// 		Type: to.Ptr("Microsoft.DatabaseFleetManager/fleets"),
	// 		SystemData: &armdatabasefleetmanager.SystemData{
	// 			CreatedBy: to.Ptr("adminuser"),
	// 			CreatedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.802Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("adminuser"),
	// 			LastModifiedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.803Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-02-01-preview/Fleets_List_MaximumSet_Gen.json
func ExampleFleetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasefleetmanager.NewClientFactory("C3897315-3847-4D8A-B2FC-7307B066AD63", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListPager(nil)
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
		// page = armdatabasefleetmanager.FleetsClientListResponse{
		// 	FleetListResult: armdatabasefleetmanager.FleetListResult{
		// 		Value: []*armdatabasefleetmanager.Fleet{
		// 			{
		// 				Properties: &armdatabasefleetmanager.FleetProperties{
		// 					Description: to.Ptr("Fleet containing critical production databases."),
		// 					ProvisioningState: to.Ptr(armdatabasefleetmanager.AzureProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 					"owner": to.Ptr("team-database"),
		// 				},
		// 				Location: to.Ptr("East US"),
		// 				ID: to.Ptr("/subscriptions/C3897315-3847-4D8A-B2FC-7307B066AD63/resourceGroups/rg-database-fleet-manager/providers/Microsoft.DatabaseFleetManager/fleet/critical-production-fleet"),
		// 				Name: to.Ptr("critical-production-fleet"),
		// 				Type: to.Ptr("Microsoft.DatabaseFleetManager/fleets"),
		// 				SystemData: &armdatabasefleetmanager.SystemData{
		// 					CreatedBy: to.Ptr("adminuser"),
		// 					CreatedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.802Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("adminuser"),
		// 					LastModifiedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.803Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2025-02-01-preview/Fleets_ListByResourceGroup_MaximumSet_Gen.json
func ExampleFleetsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasefleetmanager.NewClientFactory("C3897315-3847-4D8A-B2FC-7307B066AD63", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListByResourceGroupPager("rg-database-fleet-manager", &FleetsClientListByResourceGroupOptions{
		Skip:      to.Ptr[int64](6),
		Top:       to.Ptr[int64](30),
		Skiptoken: to.Ptr("ovlavzakdncfvvbdhqkal")})
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
		// page = armdatabasefleetmanager.FleetsClientListByResourceGroupResponse{
		// 	FleetListResult: armdatabasefleetmanager.FleetListResult{
		// 		Value: []*armdatabasefleetmanager.Fleet{
		// 			{
		// 				Properties: &armdatabasefleetmanager.FleetProperties{
		// 					Description: to.Ptr("Production fleet for critical workloads."),
		// 					ProvisioningState: to.Ptr(armdatabasefleetmanager.AzureProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 					"owner": to.Ptr("team-database"),
		// 				},
		// 				Location: to.Ptr("East US"),
		// 				ID: to.Ptr("/subscriptions/C3897315-3847-4D8A-B2FC-7307B066AD63/resourceGroups/rg-database-fleet-manager/providers/Microsoft.DatabaseFleetManager/fleet/production-fleet-01"),
		// 				Name: to.Ptr("production-fleet-01"),
		// 				Type: to.Ptr("Microsoft.DatabaseFleetManager/fleets"),
		// 				SystemData: &armdatabasefleetmanager.SystemData{
		// 					CreatedBy: to.Ptr("adminuser"),
		// 					CreatedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.802Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("adminuser"),
		// 					LastModifiedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.803Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2025-02-01-preview/Fleets_Update_MaximumSet_Gen.json
func ExampleFleetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasefleetmanager.NewClientFactory("C3897315-3847-4D8A-B2FC-7307B066AD63", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginUpdate(ctx, "rg-database-fleet-manager", "critical-production-fleet", armdatabasefleetmanager.FleetUpdate{
		Properties: &armdatabasefleetmanager.FleetProperties{
			Description: to.Ptr("Fleet containing critical production databases and high availability configurations."),
		},
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
			"owner":       to.Ptr("team-database"),
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
	// res = armdatabasefleetmanager.FleetsClientUpdateResponse{
	// 	Fleet: &armdatabasefleetmanager.Fleet{
	// 		Properties: &armdatabasefleetmanager.FleetProperties{
	// 			Description: to.Ptr("Fleet containing critical production databases and high availability configurations."),
	// 			ProvisioningState: to.Ptr(armdatabasefleetmanager.AzureProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"owner": to.Ptr("team-database"),
	// 		},
	// 		Location: to.Ptr("East US"),
	// 		ID: to.Ptr("/subscriptions/C3897315-3847-4D8A-B2FC-7307B066AD63/resourceGroups/rg-database-fleet-manager/providers/Microsoft.DatabaseFleetManager/fleet/critical-production-fleet"),
	// 		Name: to.Ptr("critical-production-fleet"),
	// 		Type: to.Ptr("Microsoft.DatabaseFleetManager/fleets"),
	// 		SystemData: &armdatabasefleetmanager.SystemData{
	// 			CreatedBy: to.Ptr("adminuser"),
	// 			CreatedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.802Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("adminuser"),
	// 			LastModifiedByType: to.Ptr(armdatabasefleetmanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-06T09:16:01.803Z"); return t}()),
	// 		},
	// 	},
	// }
}
