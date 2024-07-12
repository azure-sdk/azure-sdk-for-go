//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListDeletedWebAppsByLocation.json
func ExampleDeletedWebAppsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedWebAppsClient().NewListByLocationPager("West US 2", nil)
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
		// page.DeletedWebAppCollection = armappservice.DeletedWebAppCollection{
		// 	Value: []*armappservice.DeletedSite{
		// 		{
		// 			Name: to.Ptr("wussite6"),
		// 			Type: to.Ptr("Microsoft.Web/locations/deletedSites"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg1/providers/Microsoft.Web/locations/West US 2/deletedwebapps/wussite6"),
		// 			Properties: &armappservice.DeletedSiteProperties{
		// 				DeletedSiteID: to.Ptr[int32](9),
		// 				DeletedSiteName: to.Ptr("wussite6"),
		// 				DeletedTimestamp: to.Ptr("2019-05-09T22:29:05.1337007"),
		// 				GeoRegionName: to.Ptr("West US 2"),
		// 				Kind: to.Ptr("app"),
		// 				ResourceGroup: to.Ptr("rg1"),
		// 				Slot: to.Ptr("Production"),
		// 				Subscription: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetDeletedWebAppByLocation.json
func ExampleDeletedWebAppsClient_GetDeletedWebAppByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedWebAppsClient().GetDeletedWebAppByLocation(ctx, "West US 2", "9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedSite = armappservice.DeletedSite{
	// 	Name: to.Ptr("wussite6"),
	// 	Type: to.Ptr("Microsoft.Web/locations/deletedSites"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg1/providers/Microsoft.Web/locations/West US 2/deletedwebapps/wussite6"),
	// 	Properties: &armappservice.DeletedSiteProperties{
	// 		DeletedSiteID: to.Ptr[int32](9),
	// 		DeletedSiteName: to.Ptr("wussite6"),
	// 		DeletedTimestamp: to.Ptr("2019-05-09T22:29:05.1337007"),
	// 		GeoRegionName: to.Ptr("West US 2"),
	// 		Kind: to.Ptr("app"),
	// 		ResourceGroup: to.Ptr("rg1"),
	// 		Slot: to.Ptr("Production"),
	// 		Subscription: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	},
	// }
}