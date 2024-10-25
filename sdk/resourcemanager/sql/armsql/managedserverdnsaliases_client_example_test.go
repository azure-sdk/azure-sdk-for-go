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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedServerDnsAliasList.json
func ExampleManagedServerDNSAliasesClient_NewListByManagedInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedServerDNSAliasesClient().NewListByManagedInstancePager("Default", "dns-mi", nil)
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
		// page.ManagedServerDNSAliasListResult = armsql.ManagedServerDNSAliasListResult{
		// 	Value: []*armsql.ManagedServerDNSAlias{
		// 		{
		// 			Name: to.Ptr("dns-alias-mi"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/dnsAliases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/dns-mi/dnsAliases/dns-alias-mi"),
		// 			Properties: &armsql.ManagedServerDNSAliasProperties{
		// 				AzureDNSRecord: to.Ptr("dns-alias-mi.abcd1234.database.windows.net"),
		// 				PublicAzureDNSRecord: to.Ptr("dns-alias-mi.public.abcd1234.database.windows.net"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dns-alias-mi-2"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/dnsAliases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/dns-mi/dnsAliases/dns-alias-mi-2"),
		// 			Properties: &armsql.ManagedServerDNSAliasProperties{
		// 				AzureDNSRecord: to.Ptr("dns-alias-mi-2.abcd1234.database.windows.net"),
		// 				PublicAzureDNSRecord: to.Ptr("dns-alias-mi-2.public.abcd1234.database.windows.net"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedServerDnsAliasGet.json
func ExampleManagedServerDNSAliasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedServerDNSAliasesClient().Get(ctx, "Default", "dns-mi", "dns-alias-mi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedServerDNSAlias = armsql.ManagedServerDNSAlias{
	// 	Name: to.Ptr("dns-alias-mi"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dnsAliases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/dns-mi/dnsAliases/dns-alias-mi"),
	// 	Properties: &armsql.ManagedServerDNSAliasProperties{
	// 		AzureDNSRecord: to.Ptr("dns-alias-mi.abcd1234.database.windows.net"),
	// 		PublicAzureDNSRecord: to.Ptr("dns-alias-mi.public.abcd1234.database.windows.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedServerDnsAliasCreateOrUpdate.json
func ExampleManagedServerDNSAliasesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedServerDNSAliasesClient().BeginCreateOrUpdate(ctx, "Default", "dns-mi", "dns-alias-mi", armsql.ManagedServerDNSAliasCreation{}, nil)
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
	// res.ManagedServerDNSAlias = armsql.ManagedServerDNSAlias{
	// 	Name: to.Ptr("dns-alias-mi"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dnsAliases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/dns-mi/dnsAliases/dns-alias-mi"),
	// 	Properties: &armsql.ManagedServerDNSAliasProperties{
	// 		AzureDNSRecord: to.Ptr("dns-alias-mi.abcd1234.database.windows.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedServerDnsAliasDelete.json
func ExampleManagedServerDNSAliasesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedServerDNSAliasesClient().BeginDelete(ctx, "Default", "dns-mi", "dns-alias-mi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedServerDnsAliasAcquire.json
func ExampleManagedServerDNSAliasesClient_BeginAcquire() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedServerDNSAliasesClient().BeginAcquire(ctx, "Default", "new-mi", "dns-alias-mi", armsql.ManagedServerDNSAliasAcquisition{
		OldManagedServerDNSAliasResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/old-mi/dnsAliases/alias1"),
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
	// res.ManagedServerDNSAlias = armsql.ManagedServerDNSAlias{
	// 	Name: to.Ptr("alias1"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dnsAliases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/new-mi/dnsAliases/alias1"),
	// 	Properties: &armsql.ManagedServerDNSAliasProperties{
	// 		AzureDNSRecord: to.Ptr("alias1.abcd1234.database.windows.net"),
	// 	},
	// }
}