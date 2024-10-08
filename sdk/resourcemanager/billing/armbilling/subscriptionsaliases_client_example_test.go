//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionAliasGet.json
func ExampleSubscriptionsAliasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsAliasesClient().Get(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "c356b7c7-7545-4686-b843-c1a49cf853fc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionAlias = armbilling.SubscriptionAlias{
	// 	Name: to.Ptr("c356b7c7-7545-4686-b843-c1a49cf853fc"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptionAliases"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/c356b7c7-7545-4686-b843-c1a49cf853fc"),
	// 	SystemData: &armbilling.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-25T22:39:34.260Z"); return t}()),
	// 	},
	// 	Properties: &armbilling.SubscriptionAliasProperties{
	// 		AutoRenew: to.Ptr(armbilling.AutoRenewOn),
	// 		BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 		BillingProfileName: to.Ptr("xxxx-xxxx-xxx-xxx"),
	// 		DisplayName: to.Ptr("Billing Subscription Display Name"),
	// 		InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
	// 		InvoiceSectionName: to.Ptr("yyyy-yyyy-yyy-yyy"),
	// 		ProductCategory: to.Ptr("SeatBased"),
	// 		ProductType: to.Ptr("Seat-Based Product Type"),
	// 		ProductTypeID: to.Ptr("XYZ56789"),
	// 		PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		Quantity: to.Ptr[int64](1),
	// 		SKUDescription: to.Ptr("SKU Description"),
	// 		SKUID: to.Ptr("0001"),
	// 		Status: to.Ptr(armbilling.BillingSubscriptionStatusActive),
	// 		SystemOverrides: &armbilling.SystemOverrides{
	// 			CancellationAllowedEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T22:39:34.260Z"); return t}()),
	// 		},
	// 		TermDuration: to.Ptr("P1M"),
	// 		TermEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T22:39:34.260Z"); return t}()),
	// 		TermStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		BillingSubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionAliasCreateOrUpdate.json
func ExampleSubscriptionsAliasesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSubscriptionsAliasesClient().BeginCreateOrUpdate(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "c356b7c7-7545-4686-b843-c1a49cf853fc", armbilling.SubscriptionAlias{
		Properties: &armbilling.SubscriptionAliasProperties{
			BillingFrequency: to.Ptr("P1M"),
			DisplayName:      to.Ptr("Subscription 3"),
			Quantity:         to.Ptr[int64](1),
			SKUID:            to.Ptr("0001"),
			TermDuration:     to.Ptr("P1M"),
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
	// res.SubscriptionAlias = armbilling.SubscriptionAlias{
	// 	Name: to.Ptr("c356b7c7-7545-4686-b843-c1a49cf853fc"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptionAliases"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/c356b7c7-7545-4686-b843-c1a49cf853fc"),
	// 	SystemData: &armbilling.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-25T22:39:34.260Z"); return t}()),
	// 	},
	// 	Properties: &armbilling.SubscriptionAliasProperties{
	// 		AutoRenew: to.Ptr(armbilling.AutoRenewOn),
	// 		BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 		BillingProfileName: to.Ptr("xxxx-xxxx-xxx-xxx"),
	// 		DisplayName: to.Ptr("Billing Subscription Display Name"),
	// 		InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
	// 		InvoiceSectionName: to.Ptr("yyyy-yyyy-yyy-yyy"),
	// 		ProductCategory: to.Ptr("SeatBased"),
	// 		ProductType: to.Ptr("Seat-Based Product Type"),
	// 		ProductTypeID: to.Ptr("XYZ56789"),
	// 		PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		Quantity: to.Ptr[int64](1),
	// 		SKUDescription: to.Ptr("SKU Description"),
	// 		SKUID: to.Ptr("0001"),
	// 		Status: to.Ptr(armbilling.BillingSubscriptionStatusActive),
	// 		SystemOverrides: &armbilling.SystemOverrides{
	// 			CancellationAllowedEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T22:39:34.260Z"); return t}()),
	// 		},
	// 		TermDuration: to.Ptr("P1M"),
	// 		TermEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T22:39:34.260Z"); return t}()),
	// 		TermStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		BillingSubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionAliasList.json
func ExampleSubscriptionsAliasesClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsAliasesClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", &armbilling.SubscriptionsAliasesClientListByBillingAccountOptions{IncludeDeleted: nil,
		Filter:  nil,
		OrderBy: nil,
		Top:     nil,
		Skip:    nil,
		Count:   nil,
		Search:  nil,
	})
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
		// page.SubscriptionAliasListResult = armbilling.SubscriptionAliasListResult{
		// 	Value: []*armbilling.SubscriptionAlias{
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptionAliases"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/90000000-0000-0000-0000-000000000000"),
		// 			SystemData: &armbilling.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.260Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 			},
		// 			Properties: &armbilling.SubscriptionAliasProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				BillingProfileName: to.Ptr("xxxx-xxxx-xxx-xxx"),
		// 				DisplayName: to.Ptr("My subscription"),
		// 				InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				InvoiceSectionName: to.Ptr("yyyy-yyyy-yyy-yyy"),
		// 				ProductCategory: to.Ptr("SeatBased"),
		// 				ProductType: to.Ptr("Seat-Based Product Type"),
		// 				ProductTypeID: to.Ptr("XYZ56789"),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.260Z"); return t}()),
		// 				Quantity: to.Ptr[int64](1),
		// 				ResourceURI: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/11111111-1111-1111-1111-111111111111"),
		// 				SKUDescription: to.Ptr("SKU Description"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.BillingSubscriptionStatusActive),
		// 				SystemOverrides: &armbilling.SystemOverrides{
		// 					CancellationAllowedEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T22:39:34.260Z"); return t}()),
		// 				},
		// 				TermDuration: to.Ptr("P1M"),
		// 				TermEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T22:39:34.260Z"); return t}()),
		// 				TermStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 				BillingSubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptionAliases"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/90000000-0000-0000-0000-000000000001"),
		// 			SystemData: &armbilling.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-06T22:39:34.260Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-07T22:39:34.260Z"); return t}()),
		// 			},
		// 			Properties: &armbilling.SubscriptionAliasProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				BillingProfileName: to.Ptr("xxxx-xxxx-xxx-xxx"),
		// 				DisplayName: to.Ptr("Test subscription"),
		// 				InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				InvoiceSectionName: to.Ptr("yyyy-yyyy-yyy-yyy"),
		// 				ProductCategory: to.Ptr("Software"),
		// 				ProductType: to.Ptr("Software Product Type"),
		// 				ProductTypeID: to.Ptr("EFG456"),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-06T22:39:34.260Z"); return t}()),
		// 				Quantity: to.Ptr[int64](1),
		// 				ResourceURI: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/22222222-2222-2222-2222-222222222222"),
		// 				SKUDescription: to.Ptr("SKU Description"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.BillingSubscriptionStatusActive),
		// 				SystemOverrides: &armbilling.SystemOverrides{
		// 					CancellationAllowedEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T22:39:34.260Z"); return t}()),
		// 				},
		// 				TermDuration: to.Ptr("P1M"),
		// 				TermEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-07T22:39:34.260Z"); return t}()),
		// 				TermStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-07T22:39:34.260Z"); return t}()),
		// 				BillingSubscriptionID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000002"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptionAliases"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/90000000-0000-0000-0000-000000000002"),
		// 			SystemData: &armbilling.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-25T22:39:34.260Z"); return t}()),
		// 			},
		// 			Properties: &armbilling.SubscriptionAliasProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				BillingProfileName: to.Ptr("xxxx-xxxx-xxx-xxx"),
		// 				DisplayName: to.Ptr("Dev subscription"),
		// 				InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				InvoiceSectionName: to.Ptr("yyyy-yyyy-yyy-yyy"),
		// 				ProductCategory: to.Ptr("ReservationOrder"),
		// 				ProductType: to.Ptr("Reservation Product Type"),
		// 				ProductTypeID: to.Ptr("JKL789"),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 				Quantity: to.Ptr[int64](1),
		// 				ResourceURI: to.Ptr("/providers/Microsoft.Capacity/reservationOrders/33333333-3333-3333-3333-333333333333"),
		// 				SKUDescription: to.Ptr("SKU Description"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.BillingSubscriptionStatusSuspended),
		// 				SuspensionReasonDetails: []*armbilling.SubscriptionStatusDetails{
		// 					{
		// 						EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-15T22:39:34.260Z"); return t}()),
		// 						Reason: to.Ptr(armbilling.SubscriptionStatusReasonCancelled),
		// 				}},
		// 				SuspensionReasons: []*string{
		// 					to.Ptr("Cancelled")},
		// 					SystemOverrides: &armbilling.SystemOverrides{
		// 						CancellationAllowedEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T22:39:34.260Z"); return t}()),
		// 					},
		// 					TermDuration: to.Ptr("P1M"),
		// 					TermEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T22:39:34.260Z"); return t}()),
		// 					TermStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
		// 					BillingSubscriptionID: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 				},
		// 		}},
		// 	}
	}
}
