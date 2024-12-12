// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster/v2"
	"log"
)

// Generated from example definition: 2024-07-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armmongocluster.OperationsClientListResponse{
		// 	OperationListResult: armmongocluster.OperationListResult{
		// 		Value: []*armmongocluster.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.DocumentDB/mongoClusters/read"),
		// 				Display: &armmongocluster.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.DocumentDB"),
		// 					Resource: to.Ptr("Mongo Cluster"),
		// 					Operation: to.Ptr("Read Mongo Clusters"),
		// 					Description: to.Ptr("Reads a Mongo Cluster or list all Mongo Clusters."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.DocumentDB/mongoClusters/write"),
		// 				Display: &armmongocluster.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.DocumentDB"),
		// 					Resource: to.Ptr("Mongo Cluster"),
		// 					Operation: to.Ptr("Create or Update Mongo Cluster"),
		// 					Description: to.Ptr("Create or Update the properties or tags of the specified Mongo Cluster."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
