// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcontoso_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/contoso/armcontoso"
	"log"
)

// Generated from example definition: /mnt/vss/_work/1/s/azure-sdk-for-go/src/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/contoso/armcontoso/TempTypeSpecFiles/Contoso.Management/examples/2021-10-01-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontoso.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armcontoso.OperationsClientListResponse{
		// 	OperationListResult: armcontoso.OperationListResult{
		// 		Value: []*armcontoso.Operation{
		// 			{
		// 				Name: to.Ptr("ymeow"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armcontoso.OperationDisplay{
		// 					Provider: to.Ptr("qxyznq"),
		// 					Resource: to.Ptr("bqfwkox"),
		// 					Operation: to.Ptr("td"),
		// 					Description: to.Ptr("yvgkhsuwartgxb"),
		// 				},
		// 				Origin: to.Ptr(armcontoso.OriginUser),
		// 				ActionType: to.Ptr(armcontoso.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://sample.com/nextLink"),
		// 	},
		// }
	}
}