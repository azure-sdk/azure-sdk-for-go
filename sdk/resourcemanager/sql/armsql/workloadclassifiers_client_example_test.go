//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetWorkloadClassifier.json
func ExampleWorkloadClassifiersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadClassifiersClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", "wlm_classifier", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkloadClassifier = armsql.WorkloadClassifier{
	// 	Name: to.Ptr("wlm_classifier"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups/workloadClassifiers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/wlm_workloadgroup/workloadClassifiers/wlm_classifier"),
	// 	Properties: &armsql.WorkloadClassifierProperties{
	// 		Context: to.Ptr("test_context"),
	// 		EndTime: to.Ptr("14:00"),
	// 		Importance: to.Ptr("high"),
	// 		Label: to.Ptr("test_label"),
	// 		MemberName: to.Ptr("dbo"),
	// 		StartTime: to.Ptr("12:00"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadClassifierMax.json
func ExampleWorkloadClassifiersClient_BeginCreateOrUpdate_createAWorkloadGroupWithAllPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadClassifiersClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", "wlm_workloadclassifier", armsql.WorkloadClassifier{
		Properties: &armsql.WorkloadClassifierProperties{
			Context:    to.Ptr("test_context"),
			EndTime:    to.Ptr("14:00"),
			Importance: to.Ptr("high"),
			Label:      to.Ptr("test_label"),
			MemberName: to.Ptr("dbo"),
			StartTime:  to.Ptr("12:00"),
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
	// res.WorkloadClassifier = armsql.WorkloadClassifier{
	// 	Name: to.Ptr("wlm_workloadclassifier"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups/workloadClassifiers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/wlm_workloadgroup/workloadClassifiers/wlm_workloadclassifier"),
	// 	Properties: &armsql.WorkloadClassifierProperties{
	// 		Context: to.Ptr("test_context"),
	// 		EndTime: to.Ptr("14:00"),
	// 		Importance: to.Ptr("high"),
	// 		Label: to.Ptr("test_label"),
	// 		MemberName: to.Ptr("dbo"),
	// 		StartTime: to.Ptr("12:00"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadClassifierMin.json
func ExampleWorkloadClassifiersClient_BeginCreateOrUpdate_createAWorkloadGroupWithTheRequiredPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadClassifiersClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", "wlm_workloadclassifier", armsql.WorkloadClassifier{
		Properties: &armsql.WorkloadClassifierProperties{
			MemberName: to.Ptr("dbo"),
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
	// res.WorkloadClassifier = armsql.WorkloadClassifier{
	// 	Name: to.Ptr("wlm_workloadclassifier"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups/workloadClassifiers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/wlm_workloadgroup/workloadClassifiers/wlm_workloadclassifier"),
	// 	Properties: &armsql.WorkloadClassifierProperties{
	// 		Context: to.Ptr(""),
	// 		EndTime: to.Ptr(""),
	// 		Importance: to.Ptr(""),
	// 		Label: to.Ptr(""),
	// 		MemberName: to.Ptr("dbo"),
	// 		StartTime: to.Ptr(""),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeleteWorkloadClassifier.json
func ExampleWorkloadClassifiersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadClassifiersClient().BeginDelete(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", "wlm_workloadclassifier", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetWorkloadClassifierList.json
func ExampleWorkloadClassifiersClient_NewListByWorkloadGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadClassifiersClient().NewListByWorkloadGroupPager("Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", nil)
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
		// page.WorkloadClassifierListResult = armsql.WorkloadClassifierListResult{
		// 	Value: []*armsql.WorkloadClassifier{
		// 		{
		// 			Name: to.Ptr("classifier3"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups/workloadClassifiers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/wlm_workloadgroup/workloadClassifiers/classifier3"),
		// 			Properties: &armsql.WorkloadClassifierProperties{
		// 				Context: to.Ptr(""),
		// 				EndTime: to.Ptr(""),
		// 				Importance: to.Ptr("high"),
		// 				Label: to.Ptr(""),
		// 				MemberName: to.Ptr("dbo"),
		// 				StartTime: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("classifier1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups/workloadClassifiers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/wlm_workloadgroup/workloadClassifiers/classifier1"),
		// 			Properties: &armsql.WorkloadClassifierProperties{
		// 				Context: to.Ptr("test_context"),
		// 				EndTime: to.Ptr("14:00"),
		// 				Importance: to.Ptr("high"),
		// 				Label: to.Ptr("test_label"),
		// 				MemberName: to.Ptr("dbo"),
		// 				StartTime: to.Ptr("12:00"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("classifier2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups/workloadClassifiers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/wlm_workloadgroup/workloadClassifiers/classifier2"),
		// 			Properties: &armsql.WorkloadClassifierProperties{
		// 				Context: to.Ptr(""),
		// 				EndTime: to.Ptr("17:00"),
		// 				Importance: to.Ptr("high"),
		// 				Label: to.Ptr(""),
		// 				MemberName: to.Ptr("dbo"),
		// 				StartTime: to.Ptr("11:00"),
		// 			},
		// 	}},
		// }
	}
}