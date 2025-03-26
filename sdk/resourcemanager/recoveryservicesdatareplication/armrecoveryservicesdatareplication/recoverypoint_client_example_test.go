// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armrecoveryservicesdatareplication_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"log"
)

// Generated from example definition: 2024-09-01/RecoveryPoints_Get.json
func ExampleRecoveryPointClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryPointClient().Get(ctx, "rgrecoveryservicesdatareplication", "4", "d", "1X", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesdatareplication.RecoveryPointClientGetResponse{
	// 	RecoveryPointModel: &armrecoveryservicesdatareplication.RecoveryPointModel{
	// 		Properties: &armrecoveryservicesdatareplication.RecoveryPointModelProperties{
	// 			RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.403Z"); return t}()),
	// 			RecoveryPointType: to.Ptr(armrecoveryservicesdatareplication.RecoveryPointTypeApplicationConsistent),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/recoveryPoints/point1"),
	// 		Name: to.Ptr("mfnjlwvvfnrsllcyyeslxxcchhvld"),
	// 		Type: to.Ptr("zoeadplqxtonvqgwfqmeblh"),
	// 		SystemData: &armrecoveryservicesdatareplication.SystemData{
	// 			CreatedBy: to.Ptr("nykpygxolffv"),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.403Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("tipxxgz"),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.404Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-09-01/RecoveryPoints_List.json
func ExampleRecoveryPointClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecoveryPointClient().NewListPager("rgrecoveryservicesdatareplication", "4", "d", nil)
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
		// page = armrecoveryservicesdatareplication.RecoveryPointClientListResponse{
		// 	RecoveryPointModelListResult: armrecoveryservicesdatareplication.RecoveryPointModelListResult{
		// 		Value: []*armrecoveryservicesdatareplication.RecoveryPointModel{
		// 			{
		// 				Properties: &armrecoveryservicesdatareplication.RecoveryPointModelProperties{
		// 					RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.403Z"); return t}()),
		// 					RecoveryPointType: to.Ptr(armrecoveryservicesdatareplication.RecoveryPointTypeApplicationConsistent),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/recoveryPoints/point1"),
		// 				Name: to.Ptr("mfnjlwvvfnrsllcyyeslxxcchhvld"),
		// 				Type: to.Ptr("zoeadplqxtonvqgwfqmeblh"),
		// 				SystemData: &armrecoveryservicesdatareplication.SystemData{
		// 					CreatedBy: to.Ptr("nykpygxolffv"),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.403Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("tipxxgz"),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.404Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
