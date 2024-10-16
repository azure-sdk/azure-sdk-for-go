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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/agreementByName.json
func ExampleAgreementsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgreementsClient().Get(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "ABC123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Agreement = armbilling.Agreement{
	// 	Name: to.Ptr("ABC123"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/agreements"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/agreements/ABC123"),
	// 	Properties: &armbilling.AgreementProperties{
	// 		AcceptanceMode: to.Ptr(armbilling.AcceptanceModeClickToAccept),
	// 		AgreementLink: to.Ptr("https://contoso.com"),
	// 		DisplayName: to.Ptr("Microsoft Customer Agreement"),
	// 		EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-01T00:00:00.000Z"); return t}()),
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-01T00:00:00.000Z"); return t}()),
	// 		LeadBillingAccountName: to.Ptr("20000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 		Status: to.Ptr("Active"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/agreementsListByBillingAccount.json
func ExampleAgreementsClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAgreementsClient().NewListByBillingAccountPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", &armbilling.AgreementsClientListByBillingAccountOptions{Expand: to.Ptr("Participants")})
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
		// page.AgreementListResult = armbilling.AgreementListResult{
		// 	Value: []*armbilling.Agreement{
		// 		{
		// 			Name: to.Ptr("ABC123"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/agreements"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/agreements/ABC123"),
		// 			Properties: &armbilling.AgreementProperties{
		// 				AcceptanceMode: to.Ptr(armbilling.AcceptanceModeClickToAccept),
		// 				AgreementLink: to.Ptr("https://contoso.com"),
		// 				DisplayName: to.Ptr("Microsoft Customer Agreement"),
		// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-01T00:00:00.000Z"); return t}()),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-01T00:00:00.000Z"); return t}()),
		// 				LeadBillingAccountName: to.Ptr("20000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				Participants: []*armbilling.Participant{
		// 					{
		// 						Email: to.Ptr("abc@contoso.com"),
		// 						Status: to.Ptr("Accepted"),
		// 						StatusDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-01T00:00:00.000Z"); return t}()),
		// 					},
		// 					{
		// 						Email: to.Ptr("xtz@contoso.com"),
		// 						Status: to.Ptr("Declined"),
		// 						StatusDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-02T00:00:00.000Z"); return t}()),
		// 				}},
		// 				Status: to.Ptr("Active"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DEF456"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/agreements"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/agreements/DEF456"),
		// 			Properties: &armbilling.AgreementProperties{
		// 				AcceptanceMode: to.Ptr(armbilling.AcceptanceModeESignEmbedded),
		// 				AgreementLink: to.Ptr("https://contoso.com"),
		// 				DisplayName: to.Ptr("Microsoft Customer Agreement"),
		// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-05T00:00:00.000Z"); return t}()),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-05T00:00:00.000Z"); return t}()),
		// 				LeadBillingAccountName: to.Ptr("20000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				Participants: []*armbilling.Participant{
		// 					{
		// 						Email: to.Ptr("abc@contoso.com"),
		// 						Status: to.Ptr("Pending"),
		// 						StatusDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-05T00:00:00.000Z"); return t}()),
		// 				}},
		// 				Status: to.Ptr("PendingSignature"),
		// 			},
		// 	}},
		// }
	}
}