// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armneonpostgres_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
	"log"
)

// Generated from example definition: 2025-03-01-preview/Computes_CreateOrUpdate_MaximumSet_Gen.json
func ExampleComputesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("671936A4-ED6C-445D-ACEE-5637920E7051", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputesClient().BeginCreateOrUpdate(ctx, "rgneon", "testOrg123", "test-entity", "test-entity", "test-entity", armneonpostgres.Compute{
		Properties: &armneonpostgres.ComputeProperties{
			EntityName: to.Ptr("test-name"),
			Attributes: []*armneonpostgres.Attributes{
				{
					Name:  to.Ptr("ioyjfywmt"),
					Value: to.Ptr("sfbpcr"),
				},
			},
			Region:   to.Ptr("zgwbivhuqnyyy"),
			CPUCores: to.Ptr[int32](26),
			Memory:   to.Ptr[int32](4),
			Status:   to.Ptr("lvivnemgmdnwshhpfimppxmmpkv"),
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
	// res = armneonpostgres.ComputesClientCreateOrUpdateResponse{
	// 	Compute: &armneonpostgres.Compute{
	// 		Properties: &armneonpostgres.ComputeProperties{
	// 			EntityID: to.Ptr("test-id"),
	// 			EntityName: to.Ptr("test-name"),
	// 			CreatedAt: to.Ptr("fbzbgflvipeobpssylfqrudf"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("ioyjfywmt"),
	// 					Value: to.Ptr("sfbpcr"),
	// 				},
	// 			},
	// 			Region: to.Ptr("zgwbivhuqnyyy"),
	// 			CPUCores: to.Ptr[int32](26),
	// 			Memory: to.Ptr[int32](4),
	// 			Status: to.Ptr("lvivnemgmdnwshhpfimppxmmpkv"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/671936A4-ED6C-445D-ACEE-5637920E7051/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/testOrg123/projects/test-entity/branches/test-entity/computes/test-entity"),
	// 		Name: to.Ptr("ufxkalyygutzcotcddhrztilbupf"),
	// 		Type: to.Ptr("qzwktbvuukocftualaereycqzpanxa"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("dknvmdx"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("zhachretjgocylfh"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-01-preview/Computes_Delete_MaximumSet_Gen.json
func ExampleComputesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("671936A4-ED6C-445D-ACEE-5637920E7051", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputesClient().Delete(ctx, "rgneon", "testOrg123", "test-entity", "test-entity", "test-entity", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.ComputesClientDeleteResponse{
	// }
}

// Generated from example definition: 2025-03-01-preview/Computes_Get_MaximumSet_Gen.json
func ExampleComputesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("671936A4-ED6C-445D-ACEE-5637920E7051", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputesClient().Get(ctx, "rgneon", "testOrg123", "test-entity", "test-entity", "test-entity", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.ComputesClientGetResponse{
	// 	Compute: &armneonpostgres.Compute{
	// 		Properties: &armneonpostgres.ComputeProperties{
	// 			EntityID: to.Ptr("test-id"),
	// 			EntityName: to.Ptr("test-name"),
	// 			CreatedAt: to.Ptr("fbzbgflvipeobpssylfqrudf"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("ioyjfywmt"),
	// 					Value: to.Ptr("sfbpcr"),
	// 				},
	// 			},
	// 			Region: to.Ptr("zgwbivhuqnyyy"),
	// 			CPUCores: to.Ptr[int32](26),
	// 			Memory: to.Ptr[int32](4),
	// 			Status: to.Ptr("lvivnemgmdnwshhpfimppxmmpkv"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/671936A4-ED6C-445D-ACEE-5637920E7051/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/testOrg123/projects/test-entity/branches/test-entity/computes/test-entity"),
	// 		Name: to.Ptr("ufxkalyygutzcotcddhrztilbupf"),
	// 		Type: to.Ptr("qzwktbvuukocftualaereycqzpanxa"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("dknvmdx"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("zhachretjgocylfh"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-01-preview/Computes_List_MaximumSet_Gen.json
func ExampleComputesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("671936A4-ED6C-445D-ACEE-5637920E7051", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewComputesClient().NewListPager("rgneon", "testOrg123", "test-entity", "test-entity", nil)
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
		// page = armneonpostgres.ComputesClientListResponse{
		// 	ComputeListResult: armneonpostgres.ComputeListResult{
		// 		Value: []*armneonpostgres.Compute{
		// 			{
		// 				Properties: &armneonpostgres.ComputeProperties{
		// 					EntityID: to.Ptr("test-id"),
		// 					EntityName: to.Ptr("test-name"),
		// 					CreatedAt: to.Ptr("fbzbgflvipeobpssylfqrudf"),
		// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 					Attributes: []*armneonpostgres.Attributes{
		// 						{
		// 							Name: to.Ptr("ioyjfywmt"),
		// 							Value: to.Ptr("sfbpcr"),
		// 						},
		// 					},
		// 					Region: to.Ptr("zgwbivhuqnyyy"),
		// 					CPUCores: to.Ptr[int32](26),
		// 					Memory: to.Ptr[int32](4),
		// 					Status: to.Ptr("lvivnemgmdnwshhpfimppxmmpkv"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/671936A4-ED6C-445D-ACEE-5637920E7051/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/testOrg123/projects/test-entity/branches/test-entity/computes/test-entity"),
		// 				Name: to.Ptr("ufxkalyygutzcotcddhrztilbupf"),
		// 				Type: to.Ptr("qzwktbvuukocftualaereycqzpanxa"),
		// 				SystemData: &armneonpostgres.SystemData{
		// 					CreatedBy: to.Ptr("dknvmdx"),
		// 					CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("zhachretjgocylfh"),
		// 					LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2025-03-01-preview/Computes_Update_MaximumSet_Gen.json
func ExampleComputesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("671936A4-ED6C-445D-ACEE-5637920E7051", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputesClient().BeginUpdate(ctx, "rgneon", "testOrg123", "test-entity", "test-entity", "test-entity", armneonpostgres.Compute{
		Properties: &armneonpostgres.ComputeProperties{
			EntityName: to.Ptr("test-name"),
			Attributes: []*armneonpostgres.Attributes{
				{
					Name:  to.Ptr("ioyjfywmt"),
					Value: to.Ptr("sfbpcr"),
				},
			},
			Region:   to.Ptr("zgwbivhuqnyyy"),
			CPUCores: to.Ptr[int32](26),
			Memory:   to.Ptr[int32](4),
			Status:   to.Ptr("lvivnemgmdnwshhpfimppxmmpkv"),
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
	// res = armneonpostgres.ComputesClientUpdateResponse{
	// 	Compute: &armneonpostgres.Compute{
	// 		Properties: &armneonpostgres.ComputeProperties{
	// 			EntityID: to.Ptr("test-id"),
	// 			EntityName: to.Ptr("test-name"),
	// 			CreatedAt: to.Ptr("fbzbgflvipeobpssylfqrudf"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("ioyjfywmt"),
	// 					Value: to.Ptr("sfbpcr"),
	// 				},
	// 			},
	// 			Region: to.Ptr("zgwbivhuqnyyy"),
	// 			CPUCores: to.Ptr[int32](26),
	// 			Memory: to.Ptr[int32](4),
	// 			Status: to.Ptr("lvivnemgmdnwshhpfimppxmmpkv"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/671936A4-ED6C-445D-ACEE-5637920E7051/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/testOrg123/projects/test-entity/branches/test-entity/computes/test-entity"),
	// 		Name: to.Ptr("ufxkalyygutzcotcddhrztilbupf"),
	// 		Type: to.Ptr("qzwktbvuukocftualaereycqzpanxa"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("dknvmdx"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("zhachretjgocylfh"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-18T04:09:40.329Z"); return t}()),
	// 		},
	// 	},
	// }
}
