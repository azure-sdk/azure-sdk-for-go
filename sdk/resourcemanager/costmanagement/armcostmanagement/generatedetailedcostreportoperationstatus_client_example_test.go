//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportOperationStatusBySubscriptionScope.json
func ExampleGenerateDetailedCostReportOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGenerateDetailedCostReportOperationStatusClient().Get(ctx, "00000000-0000-0000-0000-000000000000", "subscriptions/00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GenerateDetailedCostReportOperationStatuses = armcostmanagement.GenerateDetailedCostReportOperationStatuses{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.Consumption/operationStatus"),
	// 	Error: &armcostmanagement.ErrorDetails{
	// 		Code: to.Ptr("0"),
	// 	},
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/operationStatus/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armcostmanagement.DownloadURL{
	// 		DownloadURL: to.Ptr("https://ccmreportstorageeastus.blob.core.windows.net/armreports/20201207/00000000-0000-0000-0000-000000000000?sv=2020-05-31&sr=b&sig=abcd"),
	// 		ValidTill: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-08T05:55:59.439Z"); return t}()),
	// 	},
	// 	Status: &armcostmanagement.Status{
	// 	},
	// }
}