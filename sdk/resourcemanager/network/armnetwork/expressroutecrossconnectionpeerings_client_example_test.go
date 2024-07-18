//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ExpressRouteCrossConnectionBgpPeeringList.json
func ExampleExpressRouteCrossConnectionPeeringsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteCrossConnectionPeeringsClient().NewListPager("CrossConnection-SiliconValley", "<circuitServiceKey>", nil)
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
		// page.ExpressRouteCrossConnectionPeeringList = armnetwork.ExpressRouteCrossConnectionPeeringList{
		// 	Value: []*armnetwork.ExpressRouteCrossConnectionPeering{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/CrossConnection-SiliconValley/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>/peerings/AzurePrivatePeering"),
		// 			Name: to.Ptr("AzurePrivatePeering"),
		// 			Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
		// 			Properties: &armnetwork.ExpressRouteCrossConnectionPeeringProperties{
		// 				AzureASN: to.Ptr[int32](12076),
		// 				GatewayManagerEtag: to.Ptr(""),
		// 				IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
		// 					PrimaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::/126"),
		// 					SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
		// 					State: to.Ptr(armnetwork.ExpressRouteCircuitPeeringStateEnabled),
		// 				},
		// 				LastModifiedBy: to.Ptr("Customer"),
		// 				PeerASN: to.Ptr[int64](200),
		// 				PeeringType: to.Ptr(armnetwork.ExpressRoutePeeringTypeAzurePrivatePeering),
		// 				PrimaryAzurePort: to.Ptr(""),
		// 				PrimaryPeerAddressPrefix: to.Ptr("192.168.16.252/30"),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				SecondaryAzurePort: to.Ptr(""),
		// 				SecondaryPeerAddressPrefix: to.Ptr("192.168.18.252/30"),
		// 				State: to.Ptr(armnetwork.ExpressRoutePeeringStateEnabled),
		// 				VlanID: to.Ptr[int32](200),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ExpressRouteCrossConnectionBgpPeeringDelete.json
func ExampleExpressRouteCrossConnectionPeeringsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCrossConnectionPeeringsClient().BeginDelete(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ExpressRouteCrossConnectionBgpPeeringGet.json
func ExampleExpressRouteCrossConnectionPeeringsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteCrossConnectionPeeringsClient().Get(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteCrossConnectionPeering = armnetwork.ExpressRouteCrossConnectionPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/CrossConnection-Boydton1DC/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>/peerings/AzurePrivatePeering"),
	// 	Name: to.Ptr("AzurePrivatePeering"),
	// 	Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
	// 	Properties: &armnetwork.ExpressRouteCrossConnectionPeeringProperties{
	// 		AzureASN: to.Ptr[int32](12076),
	// 		GatewayManagerEtag: to.Ptr(""),
	// 		IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
	// 			PrimaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::/126"),
	// 			SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
	// 			State: to.Ptr(armnetwork.ExpressRouteCircuitPeeringStateEnabled),
	// 		},
	// 		LastModifiedBy: to.Ptr("Customer"),
	// 		PeerASN: to.Ptr[int64](200),
	// 		PeeringType: to.Ptr(armnetwork.ExpressRoutePeeringTypeAzurePrivatePeering),
	// 		PrimaryAzurePort: to.Ptr(""),
	// 		PrimaryPeerAddressPrefix: to.Ptr("192.168.16.252/30"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SecondaryAzurePort: to.Ptr(""),
	// 		SecondaryPeerAddressPrefix: to.Ptr("192.168.18.252/30"),
	// 		State: to.Ptr(armnetwork.ExpressRoutePeeringStateEnabled),
	// 		VlanID: to.Ptr[int32](200),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ExpressRouteCrossConnectionBgpPeeringCreate.json
func ExampleExpressRouteCrossConnectionPeeringsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCrossConnectionPeeringsClient().BeginCreateOrUpdate(ctx, "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", armnetwork.ExpressRouteCrossConnectionPeering{
		Properties: &armnetwork.ExpressRouteCrossConnectionPeeringProperties{
			IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
				PrimaryPeerAddressPrefix:   to.Ptr("3FFE:FFFF:0:CD30::/126"),
				SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
			},
			PeerASN:                    to.Ptr[int64](200),
			PrimaryPeerAddressPrefix:   to.Ptr("192.168.16.252/30"),
			SecondaryPeerAddressPrefix: to.Ptr("192.168.18.252/30"),
			VlanID:                     to.Ptr[int32](200),
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
	// res.ExpressRouteCrossConnectionPeering = armnetwork.ExpressRouteCrossConnectionPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/CrossConnection-Boydton1DC/providers/Microsoft.Network/expressRouteCrossConnections/<circuitServiceKey>/peerings/AzurePrivatePeering"),
	// 	Name: to.Ptr("AzurePrivatePeering"),
	// 	Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
	// 	Properties: &armnetwork.ExpressRouteCrossConnectionPeeringProperties{
	// 		AzureASN: to.Ptr[int32](12076),
	// 		GatewayManagerEtag: to.Ptr(""),
	// 		IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
	// 			PrimaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::/126"),
	// 			SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
	// 			State: to.Ptr(armnetwork.ExpressRouteCircuitPeeringStateEnabled),
	// 		},
	// 		LastModifiedBy: to.Ptr("Customer"),
	// 		PeerASN: to.Ptr[int64](200),
	// 		PeeringType: to.Ptr(armnetwork.ExpressRoutePeeringTypeAzurePrivatePeering),
	// 		PrimaryAzurePort: to.Ptr(""),
	// 		PrimaryPeerAddressPrefix: to.Ptr("192.168.16.252/30"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SecondaryAzurePort: to.Ptr(""),
	// 		SecondaryPeerAddressPrefix: to.Ptr("192.168.18.252/30"),
	// 		State: to.Ptr(armnetwork.ExpressRoutePeeringStateEnabled),
	// 		VlanID: to.Ptr[int32](200),
	// 	},
	// }
}