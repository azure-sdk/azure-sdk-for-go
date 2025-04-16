// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armoracledatabase_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
	"log"
)

// Generated from example definition: 2025-03-01/operations_list.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armoracledatabase.OperationsClientListResponse{
		// 	OperationListResult: armoracledatabase.OperationListResult{
		// 		Value: []*armoracledatabase.Operation{
		// 			{
		// 				Name: to.Ptr("Oracle.Database/cloudExadataInfrastructures/Read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armoracledatabase.OperationDisplay{
		// 					Provider: to.Ptr("Oracle.Database"),
		// 					Resource: to.Ptr("cloudExadataInfrastructures"),
		// 					Operation: to.Ptr("Get/list Exadata Infrastructure resources"),
		// 					Description: to.Ptr("Reads Exadata Infrastructure"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
