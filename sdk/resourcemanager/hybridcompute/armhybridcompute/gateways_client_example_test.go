//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_CreateOrUpdate.json
func ExampleGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGatewaysClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{gatewayName}", armhybridcompute.Gateway{
		Location: to.Ptr("eastus2euap"),
		Properties: &armhybridcompute.GatewayProperties{
			AllowedFeatures: []*string{
				to.Ptr("*")},
			GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
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
	// res.Gateway = armhybridcompute.Gateway{
	// 	Name: to.Ptr("{gatewayName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/gateways"),
	// 	ID: to.Ptr("/subscriptions/ffd506c8-3415-42d3-9612-fdb423fb17df/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/gateways/{gatewayName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.GatewayProperties{
	// 		AllowedFeatures: []*string{
	// 			to.Ptr("*")},
	// 			GatewayEndpoint: to.Ptr("https://uniqueValue.contoso.com"),
	// 			GatewayID: to.Ptr("<generated Guid>"),
	// 			GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
	// 			ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_Update.json
func ExampleGatewaysClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().Update(ctx, "myResourceGroup", "{gatewayName}", armhybridcompute.GatewayUpdate{
		Properties: &armhybridcompute.GatewayUpdateProperties{
			AllowedFeatures: []*string{
				to.Ptr("*")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gateway = armhybridcompute.Gateway{
	// 	Name: to.Ptr("{gatewayName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/gateways"),
	// 	ID: to.Ptr("/subscriptions/ffd506c8-3415-42d3-9612-fdb423fb17df/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/gateways/{gatewayName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.GatewayProperties{
	// 		AllowedFeatures: []*string{
	// 			to.Ptr("*")},
	// 			GatewayEndpoint: to.Ptr("https://uniqueValue.contoso.com"),
	// 			GatewayID: to.Ptr("<generated Guid>"),
	// 			GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
	// 			ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_Get.json
func ExampleGatewaysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().Get(ctx, "myResourceGroup", "{gatewayName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gateway = armhybridcompute.Gateway{
	// 	Name: to.Ptr("{gatewayName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/gateways"),
	// 	ID: to.Ptr("/subscriptions/ffd506c8-3415-42d3-9612-fdb423fb17df/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/gateways/{gatewayName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.GatewayProperties{
	// 		AllowedFeatures: []*string{
	// 			to.Ptr("*")},
	// 			GatewayEndpoint: to.Ptr("https://uniqueValue.contoso.com"),
	// 			GatewayID: to.Ptr("<generated Guid>"),
	// 			GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
	// 			ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_Delete.json
func ExampleGatewaysClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGatewaysClient().BeginDelete(ctx, "myResourceGroup", "{gatewayName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_ListByResourceGroup.json
func ExampleGatewaysClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewaysClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.GatewaysListResult = armhybridcompute.GatewaysListResult{
		// 	Value: []*armhybridcompute.Gateway{
		// 		{
		// 			Name: to.Ptr("{gatewayName}"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/gateways"),
		// 			ID: to.Ptr("/subscriptions/ffd506c8-3415-42d3-9612-fdb423fb17df/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/gateways/{gatewayName}"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armhybridcompute.GatewayProperties{
		// 				AllowedFeatures: []*string{
		// 					to.Ptr("*")},
		// 					GatewayEndpoint: to.Ptr("https://uniqueValue.contoso.com"),
		// 					GatewayID: to.Ptr("<generated Guid>"),
		// 					GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
		// 					ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_ListBySubscription.json
func ExampleGatewaysClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewaysClient().NewListBySubscriptionPager(nil)
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
		// page.GatewaysListResult = armhybridcompute.GatewaysListResult{
		// 	Value: []*armhybridcompute.Gateway{
		// 		{
		// 			Name: to.Ptr("{gatewayName}"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/gateways"),
		// 			ID: to.Ptr("/subscriptions/ffd506c8-3415-42d3-9612-fdb423fb17df/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/gateways/{gatewayName}"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armhybridcompute.GatewayProperties{
		// 				AllowedFeatures: []*string{
		// 					to.Ptr("*")},
		// 					GatewayEndpoint: to.Ptr("https://uniqueValue.contoso.com"),
		// 					GatewayID: to.Ptr("<generated Guid>"),
		// 					GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
		// 					ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}