//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_List.json
func ExampleClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListPager("group1", "cloud1", nil)
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
		// page.ClusterList = armavs.ClusterList{
		// 	Value: []*armavs.Cluster{
		// 		{
		// 			Name: to.Ptr("cluster1"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/clusters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1"),
		// 			Properties: &armavs.ClusterProperties{
		// 				ClusterSize: to.Ptr[int32](3),
		// 				Hosts: []*string{
		// 					to.Ptr("fakehost22.nyc1.kubernetes.center"),
		// 					to.Ptr("fakehost23.nyc1.kubernetes.center"),
		// 					to.Ptr("fakehost24.nyc1.kubernetes.center")},
		// 					ProvisioningState: to.Ptr(armavs.ClusterProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armavs.SKU{
		// 					Name: to.Ptr("AV20"),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_Get.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "group1", "cloud1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armavs.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/clusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1"),
	// 	Properties: &armavs.ClusterProperties{
	// 		ClusterSize: to.Ptr[int32](4),
	// 		Hosts: []*string{
	// 			to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost24.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost25.nyc1.kubernetes.center")},
	// 			ProvisioningState: to.Ptr(armavs.ClusterProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armavs.SKU{
	// 			Name: to.Ptr("AV20"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_CreateOrUpdate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "cluster1", armavs.Cluster{
		Properties: &armavs.ClusterProperties{
			ClusterSize: to.Ptr[int32](3),
		},
		SKU: &armavs.SKU{
			Name: to.Ptr("AV20"),
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
	// res.Cluster = armavs.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/clusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1"),
	// 	Properties: &armavs.ClusterProperties{
	// 		ClusterSize: to.Ptr[int32](3),
	// 		Hosts: []*string{
	// 			to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost24.nyc1.kubernetes.center")},
	// 			ProvisioningState: to.Ptr(armavs.ClusterProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armavs.SKU{
	// 			Name: to.Ptr("AV20"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_Update.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "group1", "cloud1", "cluster1", armavs.ClusterUpdate{
		Properties: &armavs.ClusterUpdateProperties{
			ClusterSize: to.Ptr[int32](4),
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
	// res.Cluster = armavs.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/clusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1"),
	// 	Properties: &armavs.ClusterProperties{
	// 		ClusterSize: to.Ptr[int32](4),
	// 		Hosts: []*string{
	// 			to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost24.nyc1.kubernetes.center"),
	// 			to.Ptr("fakehost25.nyc1.kubernetes.center")},
	// 			ProvisioningState: to.Ptr(armavs.ClusterProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armavs.SKU{
	// 			Name: to.Ptr("AV20"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_Delete.json
func ExampleClustersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginDelete(ctx, "group1", "cloud1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_ListZones.json
func ExampleClustersClient_ListZones_clustersListZones() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListZones(ctx, "group1", "cloud1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterZoneList = armavs.ClusterZoneList{
	// 	Zones: []*armavs.ClusterZone{
	// 		{
	// 			Hosts: []*string{
	// 				to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost24.nyc1.kubernetes.center")},
	// 				Zone: to.Ptr("2"),
	// 		}},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_ListZones_Stretched.json
func ExampleClustersClient_ListZones_clustersListZonesStretched() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListZones(ctx, "group1", "cloud1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterZoneList = armavs.ClusterZoneList{
	// 	Zones: []*armavs.ClusterZone{
	// 		{
	// 			Hosts: []*string{
	// 				to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost24.nyc1.kubernetes.center")},
	// 				Zone: to.Ptr("2"),
	// 			},
	// 			{
	// 				Hosts: []*string{
	// 					to.Ptr("fakehost74.nyc2.kubernetes.center"),
	// 					to.Ptr("fakehost75.nyc2.kubernetes.center"),
	// 					to.Ptr("fakehost76.nyc2.kubernetes.center")},
	// 					Zone: to.Ptr("1"),
	// 			}},
	// 		}
}