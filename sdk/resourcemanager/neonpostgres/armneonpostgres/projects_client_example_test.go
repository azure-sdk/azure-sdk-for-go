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

// Generated from example definition: 2025-03-01/Projects_CreateOrUpdate_MaximumSet_Gen.json
func ExampleProjectsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginCreateOrUpdate(ctx, "rgneon", "test-org", "entity-name", armneonpostgres.Project{
		Properties: &armneonpostgres.ProjectProperties{
			EntityName: to.Ptr("entity-name"),
			Attributes: []*armneonpostgres.Attributes{
				{
					Name:  to.Ptr("trhvzyvaqy"),
					Value: to.Ptr("evpkgsskyavybxwwssm"),
				},
			},
			RegionID:         to.Ptr("tlcltldfrnxh"),
			Storage:          to.Ptr[int64](7),
			PgVersion:        to.Ptr[int32](10),
			HistoryRetention: to.Ptr[int32](7),
			DefaultEndpointSettings: &armneonpostgres.DefaultEndpointSettings{
				AutoscalingLimitMinCu: to.Ptr[float32](26),
				AutoscalingLimitMaxCu: to.Ptr[float32](20),
			},
			Branch: &armneonpostgres.BranchProperties{
				EntityName: to.Ptr("entity-name"),
				Attributes: []*armneonpostgres.Attributes{
					{
						Name:  to.Ptr("trhvzyvaqy"),
						Value: to.Ptr("evpkgsskyavybxwwssm"),
					},
				},
				ProjectID:    to.Ptr("oik"),
				ParentID:     to.Ptr("entity-id"),
				RoleName:     to.Ptr("qrrairsupyosxnqotdwhbpc"),
				DatabaseName: to.Ptr("duhxebzhd"),
				Roles: []*armneonpostgres.NeonRoleProperties{
					{
						EntityName: to.Ptr("entity-name"),
						Attributes: []*armneonpostgres.Attributes{
							{
								Name:  to.Ptr("trhvzyvaqy"),
								Value: to.Ptr("evpkgsskyavybxwwssm"),
							},
						},
						BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
						Permissions: []*string{
							to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
						},
						IsSuperUser: to.Ptr(true),
					},
				},
				Databases: []*armneonpostgres.NeonDatabaseProperties{
					{
						EntityName: to.Ptr("entity-name"),
						Attributes: []*armneonpostgres.Attributes{
							{
								Name:  to.Ptr("trhvzyvaqy"),
								Value: to.Ptr("evpkgsskyavybxwwssm"),
							},
						},
						BranchID:  to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
						OwnerName: to.Ptr("odmbeg"),
					},
				},
				Endpoints: []*armneonpostgres.EndpointProperties{
					{
						EntityName: to.Ptr("entity-name"),
						Attributes: []*armneonpostgres.Attributes{
							{
								Name:  to.Ptr("trhvzyvaqy"),
								Value: to.Ptr("evpkgsskyavybxwwssm"),
							},
						},
						ProjectID:    to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
						BranchID:     to.Ptr("rzsyrhpfbydxtfkpaa"),
						EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
					},
				},
			},
			Roles: []*armneonpostgres.NeonRoleProperties{
				{
					EntityName: to.Ptr("entity-name"),
					Attributes: []*armneonpostgres.Attributes{
						{
							Name:  to.Ptr("trhvzyvaqy"),
							Value: to.Ptr("evpkgsskyavybxwwssm"),
						},
					},
					BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
					Permissions: []*string{
						to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
					},
					IsSuperUser: to.Ptr(true),
				},
			},
			Databases: []*armneonpostgres.NeonDatabaseProperties{
				{
					EntityName: to.Ptr("entity-name"),
					Attributes: []*armneonpostgres.Attributes{
						{
							Name:  to.Ptr("trhvzyvaqy"),
							Value: to.Ptr("evpkgsskyavybxwwssm"),
						},
					},
					BranchID:  to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
					OwnerName: to.Ptr("odmbeg"),
				},
			},
			Endpoints: []*armneonpostgres.EndpointProperties{
				{
					EntityName: to.Ptr("entity-name"),
					Attributes: []*armneonpostgres.Attributes{
						{
							Name:  to.Ptr("trhvzyvaqy"),
							Value: to.Ptr("evpkgsskyavybxwwssm"),
						},
					},
					ProjectID:    to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
					BranchID:     to.Ptr("rzsyrhpfbydxtfkpaa"),
					EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
				},
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
	// res = armneonpostgres.ProjectsClientCreateOrUpdateResponse{
	// 	Project: &armneonpostgres.Project{
	// 		Properties: &armneonpostgres.ProjectProperties{
	// 			EntityID: to.Ptr("entity-id"),
	// 			EntityName: to.Ptr("entity-name"),
	// 			CreatedAt: to.Ptr("eazudrgcnzbydedhwcmgwoauc"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("trhvzyvaqy"),
	// 					Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 				},
	// 			},
	// 			RegionID: to.Ptr("tlcltldfrnxh"),
	// 			Storage: to.Ptr[int64](7),
	// 			PgVersion: to.Ptr[int32](10),
	// 			HistoryRetention: to.Ptr[int32](7),
	// 			DefaultEndpointSettings: &armneonpostgres.DefaultEndpointSettings{
	// 				AutoscalingLimitMinCu: to.Ptr[float32](26),
	// 				AutoscalingLimitMaxCu: to.Ptr[float32](20),
	// 			},
	// 			Branch: &armneonpostgres.BranchProperties{
	// 				EntityID: to.Ptr("entity-id"),
	// 				EntityName: to.Ptr("entity-name"),
	// 				CreatedAt: to.Ptr("dzbqaiixq"),
	// 				ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 				Attributes: []*armneonpostgres.Attributes{
	// 					{
	// 						Name: to.Ptr("trhvzyvaqy"),
	// 						Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 					},
	// 				},
	// 				ProjectID: to.Ptr("oik"),
	// 				ParentID: to.Ptr("entity-id"),
	// 				RoleName: to.Ptr("qrrairsupyosxnqotdwhbpc"),
	// 				DatabaseName: to.Ptr("duhxebzhd"),
	// 				Roles: []*armneonpostgres.NeonRoleProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 						Permissions: []*string{
	// 							to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 						},
	// 						IsSuperUser: to.Ptr(true),
	// 					},
	// 				},
	// 				Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("wgdmylla"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 						OwnerName: to.Ptr("odmbeg"),
	// 					},
	// 				},
	// 				Endpoints: []*armneonpostgres.EndpointProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("vhcilurdd"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 						BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 						EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 					},
	// 				},
	// 			},
	// 			Roles: []*armneonpostgres.NeonRoleProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 					Permissions: []*string{
	// 						to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 					},
	// 					IsSuperUser: to.Ptr(true),
	// 				},
	// 			},
	// 			Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("wgdmylla"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 					OwnerName: to.Ptr("odmbeg"),
	// 				},
	// 			},
	// 			Endpoints: []*armneonpostgres.EndpointProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("vhcilurdd"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 					BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 					EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org/projects/entity-name"),
	// 		Name: to.Ptr("cajluydhtylhjatsexnmxmxhwfs"),
	// 		Type: to.Ptr("voaddrcrqtyqae"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("hnyidmqyvvtsddrwkmrqlwtlew"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("szuncyyauzxhpzlbcvjkeamp"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-01/Projects_Delete_MaximumSet_Gen.json
func ExampleProjectsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().Delete(ctx, "rgneon", "test-org", "entity-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.ProjectsClientDeleteResponse{
	// }
}

// Generated from example definition: 2025-03-01/Projects_Get_MaximumSet_Gen.json
func ExampleProjectsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().Get(ctx, "rgneon", "test-org", "entity-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.ProjectsClientGetResponse{
	// 	Project: &armneonpostgres.Project{
	// 		Properties: &armneonpostgres.ProjectProperties{
	// 			EntityID: to.Ptr("entity-id"),
	// 			EntityName: to.Ptr("entity-name"),
	// 			CreatedAt: to.Ptr("eazudrgcnzbydedhwcmgwoauc"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("trhvzyvaqy"),
	// 					Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 				},
	// 			},
	// 			RegionID: to.Ptr("tlcltldfrnxh"),
	// 			Storage: to.Ptr[int64](7),
	// 			PgVersion: to.Ptr[int32](10),
	// 			HistoryRetention: to.Ptr[int32](7),
	// 			DefaultEndpointSettings: &armneonpostgres.DefaultEndpointSettings{
	// 				AutoscalingLimitMinCu: to.Ptr[float32](26),
	// 				AutoscalingLimitMaxCu: to.Ptr[float32](20),
	// 			},
	// 			Branch: &armneonpostgres.BranchProperties{
	// 				EntityID: to.Ptr("entity-id"),
	// 				EntityName: to.Ptr("entity-name"),
	// 				CreatedAt: to.Ptr("dzbqaiixq"),
	// 				ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 				Attributes: []*armneonpostgres.Attributes{
	// 					{
	// 						Name: to.Ptr("trhvzyvaqy"),
	// 						Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 					},
	// 				},
	// 				ProjectID: to.Ptr("oik"),
	// 				ParentID: to.Ptr("entity-id"),
	// 				RoleName: to.Ptr("qrrairsupyosxnqotdwhbpc"),
	// 				DatabaseName: to.Ptr("duhxebzhd"),
	// 				Roles: []*armneonpostgres.NeonRoleProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 						Permissions: []*string{
	// 							to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 						},
	// 						IsSuperUser: to.Ptr(true),
	// 					},
	// 				},
	// 				Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("wgdmylla"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 						OwnerName: to.Ptr("odmbeg"),
	// 					},
	// 				},
	// 				Endpoints: []*armneonpostgres.EndpointProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("vhcilurdd"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 						BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 						EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 					},
	// 				},
	// 			},
	// 			Roles: []*armneonpostgres.NeonRoleProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 					Permissions: []*string{
	// 						to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 					},
	// 					IsSuperUser: to.Ptr(true),
	// 				},
	// 			},
	// 			Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("wgdmylla"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 					OwnerName: to.Ptr("odmbeg"),
	// 				},
	// 			},
	// 			Endpoints: []*armneonpostgres.EndpointProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("vhcilurdd"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 					BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 					EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org/projects/entity-name"),
	// 		Name: to.Ptr("cajluydhtylhjatsexnmxmxhwfs"),
	// 		Type: to.Ptr("voaddrcrqtyqae"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("hnyidmqyvvtsddrwkmrqlwtlew"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("szuncyyauzxhpzlbcvjkeamp"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2025-03-01/Projects_GetConnectionUri_MaximumSet_Gen.json
func ExampleProjectsClient_GetConnectionURI() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().GetConnectionURI(ctx, "rgneon", "test-org", "entity-name", armneonpostgres.ConnectionURIProperties{
		ProjectID:    to.Ptr("riuifmoqtorrcffgksvfcobia"),
		BranchID:     to.Ptr("iimmlbqv"),
		DatabaseName: to.Ptr("xc"),
		RoleName:     to.Ptr("xhmcvsgtp"),
		EndpointID:   to.Ptr("jcpdvsyjcn"),
		IsPooled:     to.Ptr(true),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.ProjectsClientGetConnectionURIResponse{
	// 	ConnectionURIProperties: &armneonpostgres.ConnectionURIProperties{
	// 		ProjectID: to.Ptr("riuifmoqtorrcffgksvfcobia"),
	// 		BranchID: to.Ptr("iimmlbqv"),
	// 		DatabaseName: to.Ptr("xc"),
	// 		RoleName: to.Ptr("xhmcvsgtp"),
	// 		EndpointID: to.Ptr("jcpdvsyjcn"),
	// 		IsPooled: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: 2025-03-01/Projects_List_MaximumSet_Gen.json
func ExampleProjectsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectsClient().NewListPager("rgneon", "test-org", nil)
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
		// page = armneonpostgres.ProjectsClientListResponse{
		// 	ProjectListResult: armneonpostgres.ProjectListResult{
		// 		Value: []*armneonpostgres.Project{
		// 			{
		// 				Properties: &armneonpostgres.ProjectProperties{
		// 					EntityID: to.Ptr("entity-id"),
		// 					EntityName: to.Ptr("entity-name"),
		// 					CreatedAt: to.Ptr("eazudrgcnzbydedhwcmgwoauc"),
		// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 					Attributes: []*armneonpostgres.Attributes{
		// 						{
		// 							Name: to.Ptr("trhvzyvaqy"),
		// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 						},
		// 					},
		// 					RegionID: to.Ptr("tlcltldfrnxh"),
		// 					Storage: to.Ptr[int64](7),
		// 					PgVersion: to.Ptr[int32](10),
		// 					HistoryRetention: to.Ptr[int32](7),
		// 					DefaultEndpointSettings: &armneonpostgres.DefaultEndpointSettings{
		// 						AutoscalingLimitMinCu: to.Ptr[float32](26),
		// 						AutoscalingLimitMaxCu: to.Ptr[float32](20),
		// 					},
		// 					Branch: &armneonpostgres.BranchProperties{
		// 						EntityID: to.Ptr("entity-id"),
		// 						EntityName: to.Ptr("entity-name"),
		// 						CreatedAt: to.Ptr("dzbqaiixq"),
		// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 						Attributes: []*armneonpostgres.Attributes{
		// 							{
		// 								Name: to.Ptr("trhvzyvaqy"),
		// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 							},
		// 						},
		// 						ProjectID: to.Ptr("oik"),
		// 						ParentID: to.Ptr("entity-id"),
		// 						RoleName: to.Ptr("qrrairsupyosxnqotdwhbpc"),
		// 						DatabaseName: to.Ptr("duhxebzhd"),
		// 						Roles: []*armneonpostgres.NeonRoleProperties{
		// 							{
		// 								EntityID: to.Ptr("entity-id"),
		// 								EntityName: to.Ptr("entity-name"),
		// 								CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
		// 								ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 								Attributes: []*armneonpostgres.Attributes{
		// 									{
		// 										Name: to.Ptr("trhvzyvaqy"),
		// 										Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 									},
		// 								},
		// 								BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
		// 								Permissions: []*string{
		// 									to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
		// 								},
		// 								IsSuperUser: to.Ptr(true),
		// 							},
		// 						},
		// 						Databases: []*armneonpostgres.NeonDatabaseProperties{
		// 							{
		// 								EntityID: to.Ptr("entity-id"),
		// 								EntityName: to.Ptr("entity-name"),
		// 								CreatedAt: to.Ptr("wgdmylla"),
		// 								ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 								Attributes: []*armneonpostgres.Attributes{
		// 									{
		// 										Name: to.Ptr("trhvzyvaqy"),
		// 										Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 									},
		// 								},
		// 								BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
		// 								OwnerName: to.Ptr("odmbeg"),
		// 							},
		// 						},
		// 						Endpoints: []*armneonpostgres.EndpointProperties{
		// 							{
		// 								EntityID: to.Ptr("entity-id"),
		// 								EntityName: to.Ptr("entity-name"),
		// 								CreatedAt: to.Ptr("vhcilurdd"),
		// 								ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 								Attributes: []*armneonpostgres.Attributes{
		// 									{
		// 										Name: to.Ptr("trhvzyvaqy"),
		// 										Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 									},
		// 								},
		// 								ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
		// 								BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
		// 								EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
		// 							},
		// 						},
		// 					},
		// 					Roles: []*armneonpostgres.NeonRoleProperties{
		// 						{
		// 							EntityID: to.Ptr("entity-id"),
		// 							EntityName: to.Ptr("entity-name"),
		// 							CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
		// 							ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 							Attributes: []*armneonpostgres.Attributes{
		// 								{
		// 									Name: to.Ptr("trhvzyvaqy"),
		// 									Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 								},
		// 							},
		// 							BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
		// 							Permissions: []*string{
		// 								to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
		// 							},
		// 							IsSuperUser: to.Ptr(true),
		// 						},
		// 					},
		// 					Databases: []*armneonpostgres.NeonDatabaseProperties{
		// 						{
		// 							EntityID: to.Ptr("entity-id"),
		// 							EntityName: to.Ptr("entity-name"),
		// 							CreatedAt: to.Ptr("wgdmylla"),
		// 							ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 							Attributes: []*armneonpostgres.Attributes{
		// 								{
		// 									Name: to.Ptr("trhvzyvaqy"),
		// 									Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 								},
		// 							},
		// 							BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
		// 							OwnerName: to.Ptr("odmbeg"),
		// 						},
		// 					},
		// 					Endpoints: []*armneonpostgres.EndpointProperties{
		// 						{
		// 							EntityID: to.Ptr("entity-id"),
		// 							EntityName: to.Ptr("entity-name"),
		// 							CreatedAt: to.Ptr("vhcilurdd"),
		// 							ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 							Attributes: []*armneonpostgres.Attributes{
		// 								{
		// 									Name: to.Ptr("trhvzyvaqy"),
		// 									Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 								},
		// 							},
		// 							ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
		// 							BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
		// 							EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org/projects/cajluydhtylhjatsexnmxmxhwfs"),
		// 				Name: to.Ptr("cajluydhtylhjatsexnmxmxhwfs"),
		// 				Type: to.Ptr("voaddrcrqtyqae"),
		// 				SystemData: &armneonpostgres.SystemData{
		// 					CreatedBy: to.Ptr("hnyidmqyvvtsddrwkmrqlwtlew"),
		// 					CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("szuncyyauzxhpzlbcvjkeamp"),
		// 					LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/azyigg"),
		// 	},
		// }
	}
}

// Generated from example definition: 2025-03-01/Projects_Update_MaximumSet_Gen.json
func ExampleProjectsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginUpdate(ctx, "rgneon", "test-org", "test-project", armneonpostgres.Project{
		Properties: &armneonpostgres.ProjectProperties{
			EntityName: to.Ptr("entity-name"),
			Attributes: []*armneonpostgres.Attributes{
				{
					Name:  to.Ptr("trhvzyvaqy"),
					Value: to.Ptr("evpkgsskyavybxwwssm"),
				},
			},
			RegionID:         to.Ptr("vxvmjwuttpiakirzdf"),
			Storage:          to.Ptr[int64](23),
			PgVersion:        to.Ptr[int32](16),
			HistoryRetention: to.Ptr[int32](16),
			DefaultEndpointSettings: &armneonpostgres.DefaultEndpointSettings{
				AutoscalingLimitMinCu: to.Ptr[float32](8),
				AutoscalingLimitMaxCu: to.Ptr[float32](4),
			},
			Branch: &armneonpostgres.BranchProperties{
				EntityName: to.Ptr("entity-name"),
				Attributes: []*armneonpostgres.Attributes{
					{
						Name:  to.Ptr("trhvzyvaqy"),
						Value: to.Ptr("evpkgsskyavybxwwssm"),
					},
				},
				ProjectID:    to.Ptr("oik"),
				ParentID:     to.Ptr("entity-id"),
				RoleName:     to.Ptr("qrrairsupyosxnqotdwhbpc"),
				DatabaseName: to.Ptr("duhxebzhd"),
				Roles: []*armneonpostgres.NeonRoleProperties{
					{
						EntityName: to.Ptr("entity-name"),
						Attributes: []*armneonpostgres.Attributes{
							{
								Name:  to.Ptr("trhvzyvaqy"),
								Value: to.Ptr("evpkgsskyavybxwwssm"),
							},
						},
						BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
						Permissions: []*string{
							to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
						},
						IsSuperUser: to.Ptr(true),
					},
				},
				Databases: []*armneonpostgres.NeonDatabaseProperties{
					{
						EntityName: to.Ptr("entity-name"),
						Attributes: []*armneonpostgres.Attributes{
							{
								Name:  to.Ptr("trhvzyvaqy"),
								Value: to.Ptr("evpkgsskyavybxwwssm"),
							},
						},
						BranchID:  to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
						OwnerName: to.Ptr("odmbeg"),
					},
				},
				Endpoints: []*armneonpostgres.EndpointProperties{
					{
						EntityName: to.Ptr("entity-name"),
						Attributes: []*armneonpostgres.Attributes{
							{
								Name:  to.Ptr("trhvzyvaqy"),
								Value: to.Ptr("evpkgsskyavybxwwssm"),
							},
						},
						ProjectID:    to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
						BranchID:     to.Ptr("rzsyrhpfbydxtfkpaa"),
						EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
					},
				},
			},
			Roles: []*armneonpostgres.NeonRoleProperties{
				{
					EntityName: to.Ptr("entity-name"),
					Attributes: []*armneonpostgres.Attributes{
						{
							Name:  to.Ptr("trhvzyvaqy"),
							Value: to.Ptr("evpkgsskyavybxwwssm"),
						},
					},
					BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
					Permissions: []*string{
						to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
					},
					IsSuperUser: to.Ptr(true),
				},
			},
			Databases: []*armneonpostgres.NeonDatabaseProperties{
				{
					EntityName: to.Ptr("entity-name"),
					Attributes: []*armneonpostgres.Attributes{
						{
							Name:  to.Ptr("trhvzyvaqy"),
							Value: to.Ptr("evpkgsskyavybxwwssm"),
						},
					},
					BranchID:  to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
					OwnerName: to.Ptr("odmbeg"),
				},
			},
			Endpoints: []*armneonpostgres.EndpointProperties{
				{
					EntityName: to.Ptr("entity-name"),
					Attributes: []*armneonpostgres.Attributes{
						{
							Name:  to.Ptr("trhvzyvaqy"),
							Value: to.Ptr("evpkgsskyavybxwwssm"),
						},
					},
					ProjectID:    to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
					BranchID:     to.Ptr("rzsyrhpfbydxtfkpaa"),
					EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
				},
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
	// res = armneonpostgres.ProjectsClientUpdateResponse{
	// 	Project: &armneonpostgres.Project{
	// 		Properties: &armneonpostgres.ProjectProperties{
	// 			EntityID: to.Ptr("entity-id"),
	// 			EntityName: to.Ptr("entity-name"),
	// 			CreatedAt: to.Ptr("eazudrgcnzbydedhwcmgwoauc"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("trhvzyvaqy"),
	// 					Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 				},
	// 			},
	// 			RegionID: to.Ptr("tlcltldfrnxh"),
	// 			Storage: to.Ptr[int64](7),
	// 			PgVersion: to.Ptr[int32](10),
	// 			HistoryRetention: to.Ptr[int32](7),
	// 			DefaultEndpointSettings: &armneonpostgres.DefaultEndpointSettings{
	// 				AutoscalingLimitMinCu: to.Ptr[float32](8),
	// 				AutoscalingLimitMaxCu: to.Ptr[float32](4),
	// 			},
	// 			Branch: &armneonpostgres.BranchProperties{
	// 				EntityID: to.Ptr("entity-id"),
	// 				EntityName: to.Ptr("entity-name"),
	// 				CreatedAt: to.Ptr("dzbqaiixq"),
	// 				ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 				Attributes: []*armneonpostgres.Attributes{
	// 					{
	// 						Name: to.Ptr("trhvzyvaqy"),
	// 						Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 					},
	// 				},
	// 				ProjectID: to.Ptr("oik"),
	// 				ParentID: to.Ptr("entity-id"),
	// 				RoleName: to.Ptr("qrrairsupyosxnqotdwhbpc"),
	// 				DatabaseName: to.Ptr("duhxebzhd"),
	// 				Roles: []*armneonpostgres.NeonRoleProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 						Permissions: []*string{
	// 							to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 						},
	// 						IsSuperUser: to.Ptr(true),
	// 					},
	// 				},
	// 				Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("wgdmylla"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 						OwnerName: to.Ptr("odmbeg"),
	// 					},
	// 				},
	// 				Endpoints: []*armneonpostgres.EndpointProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("vhcilurdd"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 						BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 						EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 					},
	// 				},
	// 			},
	// 			Roles: []*armneonpostgres.NeonRoleProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 					Permissions: []*string{
	// 						to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 					},
	// 					IsSuperUser: to.Ptr(true),
	// 				},
	// 			},
	// 			Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("wgdmylla"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 					OwnerName: to.Ptr("odmbeg"),
	// 				},
	// 			},
	// 			Endpoints: []*armneonpostgres.EndpointProperties{
	// 				{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("vhcilurdd"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 					BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 					EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org/projects/cajluydhtylhjatsexnmxmxhwfs"),
	// 		Name: to.Ptr("cajluydhtylhjatsexnmxmxhwfs"),
	// 		Type: to.Ptr("voaddrcrqtyqae"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("hnyidmqyvvtsddrwkmrqlwtlew"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("szuncyyauzxhpzlbcvjkeamp"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 		},
	// 	},
	// }
}
