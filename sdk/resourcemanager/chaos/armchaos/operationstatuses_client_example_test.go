// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armchaos_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
	"log"
)

// Generated from example definition: 2025-01-01/OperationStatuses_Get.json
func ExampleOperationStatusesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("e25c0d12-0335-4fec-8ef8-3b4f9a10649e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusesClient().Get(ctx, "westus2", "4bdadd97-207c-4de8-9bba-08339ae099c7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.OperationStatusesClientGetResponse{
	// 	OperationStatusResult: &armchaos.OperationStatusResult{
	// 		ID: to.Ptr("/subscriptions/e25c0d12-0335-4fec-8ef8-3b4f9a10649e/providers/Microsoft.Chaos/locations/westus2/operationStatuses/4bdadd97-207c-4de8-9bba-08339ae099c7"),
	// 		Name: to.Ptr("4bdadd97-207c-4de8-9bba-08339ae099c7"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T21:52:52.2552574Z"); return t}()),
	// 		Status: to.Ptr("Creating"),
	// 	},
	// }
}
