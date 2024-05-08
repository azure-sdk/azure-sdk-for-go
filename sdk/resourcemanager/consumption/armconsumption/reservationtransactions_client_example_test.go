//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationTransactionsListByEnrollmentNumber.json
func ExampleReservationTransactionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationTransactionsClient().NewListPager("123456", &armconsumption.ReservationTransactionsClientListOptions{Filter: to.Ptr("properties/eventDate+ge+2020-05-20+AND+properties/eventDate+le+2020-05-30")})
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
		// page.ReservationTransactionsListResult = armconsumption.ReservationTransactionsListResult{
		// 	Value: []*armconsumption.ReservationTransaction{
		// 		{
		// 			Name: to.Ptr("201909091919"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationTransactions"),
		// 			ID: to.Ptr("/billingAccounts/123456/providers/Microsoft.Consumption/reservationtransactions/201909091919"),
		// 			Tags: []*string{
		// 			},
		// 			Properties: &armconsumption.LegacyReservationTransactionProperties{
		// 				Description: to.Ptr("Standard_DS1_v2 westus 1 Year"),
		// 				AccountName: to.Ptr("Microsoft Infrastructure"),
		// 				AccountOwnerEmail: to.Ptr("admin@microsoft.com"),
		// 				Amount: to.Ptr[float64](-21),
		// 				ArmSKUName: to.Ptr("Standard_DS1_v2"),
		// 				BillingFrequency: to.Ptr("recurring"),
		// 				BillingMonth: to.Ptr[int32](20190901),
		// 				CostCenter: to.Ptr(""),
		// 				Currency: to.Ptr("USD"),
		// 				CurrentEnrollment: to.Ptr("123456"),
		// 				DepartmentName: to.Ptr("Unassigned"),
		// 				EventDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-09T19:19:04.000Z"); return t}()),
		// 				EventType: to.Ptr("Refund"),
		// 				MonetaryCommitment: to.Ptr[float64](523123.9),
		// 				Overage: to.Ptr[float64](23234.49),
		// 				PurchasingEnrollment: to.Ptr("123456"),
		// 				PurchasingSubscriptionGUID: to.Ptr("a838a8c3-a408-49e1-ac90-42cb95bff9b2"),
		// 				PurchasingSubscriptionName: to.Ptr("Infrastructure Subscription"),
		// 				Quantity: to.Ptr[float64](1),
		// 				Region: to.Ptr("westus"),
		// 				ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ReservationOrderName: to.Ptr("Transaction-DS1_v2"),
		// 				Term: to.Ptr("P1Y"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("201909091919"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationTransactions"),
		// 			ID: to.Ptr("/billingAccounts/123456/providers/Microsoft.Consumption/reservationtransactions/201909091919"),
		// 			Tags: []*string{
		// 			},
		// 			Properties: &armconsumption.LegacyReservationTransactionProperties{
		// 				Description: to.Ptr("Standard_DS1_v2 westus 1 Year"),
		// 				AccountName: to.Ptr("Microsoft Infrastructure"),
		// 				AccountOwnerEmail: to.Ptr("admin@microsoft.com"),
		// 				Amount: to.Ptr[float64](21),
		// 				ArmSKUName: to.Ptr("Standard_DS1_v2"),
		// 				BillingFrequency: to.Ptr("recurring"),
		// 				BillingMonth: to.Ptr[int32](20190901),
		// 				CostCenter: to.Ptr(""),
		// 				Currency: to.Ptr("USD"),
		// 				CurrentEnrollment: to.Ptr("123456"),
		// 				DepartmentName: to.Ptr("Unassigned"),
		// 				EventDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-09T19:19:04.000Z"); return t}()),
		// 				EventType: to.Ptr("Purchase"),
		// 				MonetaryCommitment: to.Ptr[float64](523123.9),
		// 				Overage: to.Ptr[float64](23234.49),
		// 				PurchasingEnrollment: to.Ptr("123456"),
		// 				PurchasingSubscriptionGUID: to.Ptr("a838a8c3-a408-49e1-ac90-42cb95bff9b2"),
		// 				PurchasingSubscriptionName: to.Ptr("Infrastructure Subscription"),
		// 				Quantity: to.Ptr[float64](1),
		// 				Region: to.Ptr("westus"),
		// 				ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ReservationOrderName: to.Ptr("Transaction-DS1_v2"),
		// 				Term: to.Ptr("P1Y"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationTransactionsListByBillingProfileId.json
func ExampleReservationTransactionsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationTransactionsClient().NewListByBillingProfilePager("fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30", "Z76D-SGAF-BG7-TGB", &armconsumption.ReservationTransactionsClientListByBillingProfileOptions{Filter: to.Ptr("properties/eventDate+ge+2020-05-20+AND+properties/eventDate+le+2020-05-30")})
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
		// page.ModernReservationTransactionsListResult = armconsumption.ModernReservationTransactionsListResult{
		// 	Value: []*armconsumption.ModernReservationTransaction{
		// 		{
		// 			Name: to.Ptr("a838a8c3-a408-49e1-ac90-42cb95bff9b2"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationTransactions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/billingProfiles/Z76D-SGAF-BG7-TGB/providers/Microsoft.Consumption/reservationTransactions"),
		// 			Properties: &armconsumption.ModernReservationTransactionProperties{
		// 				Description: to.Ptr("Reserved VM Instance, Standard_B1ls, US East, 3 Years"),
		// 				Amount: to.Ptr[float64](1.44),
		// 				ArmSKUName: to.Ptr("Standard_B1ls"),
		// 				BillingFrequency: to.Ptr("Recurring"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/billingProfiles/Z76D-SGAF-BG7-TGB"),
		// 				BillingProfileName: to.Ptr("IT Department*"),
		// 				Currency: to.Ptr("USD"),
		// 				EventDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-25T21:21:38.000Z"); return t}()),
		// 				EventType: to.Ptr("Purchase"),
		// 				Invoice: to.Ptr("T000456437"),
		// 				InvoiceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/billingProfiles/Z76D-SGAF-BG7-TGB/invoices/T000456437"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/invoiceSections/QBTB-EYAK-PJA-TGB"),
		// 				InvoiceSectionName: to.Ptr("IT Department"),
		// 				PurchasingSubscriptionGUID: to.Ptr("d924ad15-4a3d-4047-971d-c8b1b300a97b"),
		// 				PurchasingSubscriptionName: to.Ptr("contoso"),
		// 				Quantity: to.Ptr[float64](1),
		// 				Region: to.Ptr("eastus"),
		// 				ReservationOrderID: to.Ptr("a838a8c3-a408-49e1-ac90-42cb95bff9b2"),
		// 				ReservationOrderName: to.Ptr("VM_RI_03-25-2020_14-18"),
		// 				Term: to.Ptr("P3Y"),
		// 			},
		// 	}},
		// }
	}
}