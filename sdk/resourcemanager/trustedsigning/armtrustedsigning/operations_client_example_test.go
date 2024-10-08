// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armtrustedsigning_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning"
	"log"
)

// Generated from example definition: 2024-02-05-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page = armtrustedsigning.OperationsClientListResponse{
		// 	OperationListResult: armtrustedsigning.OperationListResult{
		// 		Value: []*armtrustedsigning.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/write"),
		// 				Display: &armtrustedsigning.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.CodeSigning"),
		// 					Resource: to.Ptr("codeSigningAccounts"),
		// 					Operation: to.Ptr("Create CodeSigningAccount"),
		// 					Description: to.Ptr("Create any CodeSigningAccount"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/read"),
		// 				Display: &armtrustedsigning.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.CodeSigning"),
		// 					Resource: to.Ptr("codeSigningAccounts"),
		// 					Operation: to.Ptr("Get CodeSigningAccount"),
		// 					Description: to.Ptr("Get any CodeSigningAccount"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts/delete"),
		// 				Display: &armtrustedsigning.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.CodeSigning"),
		// 					Resource: to.Ptr("codeSigningAccounts"),
		// 					Operation: to.Ptr("Delete CodeSigningAccount"),
		// 					Description: to.Ptr("Delete any CodeSigningAccount"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
