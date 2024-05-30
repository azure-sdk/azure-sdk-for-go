//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_ListByFactory.json
func ExampleCredentialOperationsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCredentialOperationsClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.CredentialListResponse = armdatafactory.CredentialListResponse{
		// 	Value: []*armdatafactory.CredentialResource{
		// 		{
		// 			Name: to.Ptr("exampleLinkedService"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/credentials"),
		// 			Etag: to.Ptr("0a0064d4-0000-0000-0000-5b245bd00000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/credentials/exampleCredential"),
		// 			Properties: &armdatafactory.ManagedIdentityCredential{
		// 				Type: to.Ptr("ManagedIdentity"),
		// 				Description: to.Ptr("Example description"),
		// 				AdditionalProperties: map[string]any{
		// 					"typeProperties": map[string]any{
		// 						"resourceId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami",
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Create.json
func ExampleCredentialOperationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCredentialOperationsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleCredential", armdatafactory.CredentialResource{
		Properties: &armdatafactory.ManagedIdentityCredential{
			Type: to.Ptr("ManagedIdentity"),
			AdditionalProperties: map[string]any{
				"typeProperties": map[string]any{
					"resourceId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami",
				},
			},
		},
	}, &armdatafactory.CredentialOperationsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CredentialResource = armdatafactory.CredentialResource{
	// 	Name: to.Ptr("exampleCredential"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/credentials"),
	// 	Etag: to.Ptr("0a0062d4-0000-0000-0000-5b245bcf0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/credentials/exampleCredential"),
	// 	Properties: &armdatafactory.ManagedIdentityCredential{
	// 		Type: to.Ptr("ManagedIdentity"),
	// 		AdditionalProperties: map[string]any{
	// 			"typeProperties": map[string]any{
	// 				"resourceId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami",
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Get.json
func ExampleCredentialOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCredentialOperationsClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleCredential", &armdatafactory.CredentialOperationsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CredentialResource = armdatafactory.CredentialResource{
	// 	Name: to.Ptr("exampleLinkedService"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/credentials"),
	// 	Etag: to.Ptr("1500474f-0000-0200-0000-5cbe090d0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/credentials/exampleCredential"),
	// 	Properties: &armdatafactory.ManagedIdentityCredential{
	// 		Type: to.Ptr("ManagedIdentity"),
	// 		Description: to.Ptr("Example description"),
	// 		AdditionalProperties: map[string]any{
	// 			"typeProperties": map[string]any{
	// 				"resourceId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami",
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Delete.json
func ExampleCredentialOperationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCredentialOperationsClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleCredential", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}