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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyGetMCA.json
func ExamplePropertyClient_Get_billingPropertyGetMca() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPropertyClient().Get(ctx, &armbilling.PropertyClientGetOptions{IncludeBillingCountry: nil,
		IncludeTransitionStatus: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Property = armbilling.Property{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingProperty"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.Billing/billingProperty/default"),
	// 	Properties: &armbilling.PropertyProperties{
	// 		BillingAccountAgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
	// 		BillingAccountDisplayName: to.Ptr("My Account"),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000"),
	// 		BillingAccountSoldToCountry: to.Ptr("US"),
	// 		BillingAccountType: to.Ptr(armbilling.AccountTypeBusiness),
	// 		BillingProfileDisplayName: to.Ptr("Contoso"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000/billingProfiles/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 		BillingProfileSpendingLimit: to.Ptr(armbilling.SpendingLimitOff),
	// 		BillingProfileStatus: to.Ptr(armbilling.BillingProfileStatusActive),
	// 		CostCenter: to.Ptr("1010"),
	// 		InvoiceSectionDisplayName: to.Ptr("Operations"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000/billingProfiles/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
	// 		SKUDescription: to.Ptr("Microsoft Azure Plan"),
	// 		SKUID: to.Ptr("0001"),
	// 		SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusActive),
	// 		SubscriptionBillingStatusDetails: []*armbilling.SubscriptionStatusDetails{
	// 			{
	// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-01T17:32:28.000Z"); return t}()),
	// 				Reason: to.Ptr(armbilling.SubscriptionStatusReasonCancelled),
	// 		}},
	// 		SubscriptionBillingType: to.Ptr(armbilling.SubscriptionBillingTypeFree),
	// 		SubscriptionWorkloadType: to.Ptr(armbilling.SubscriptionWorkloadTypeProduction),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyGetMOSP.json
func ExamplePropertyClient_Get_billingPropertyGetMosp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPropertyClient().Get(ctx, &armbilling.PropertyClientGetOptions{IncludeBillingCountry: nil,
		IncludeTransitionStatus: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Property = armbilling.Property{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingProperty"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.Billing/billingProperty/default"),
	// 	Properties: &armbilling.PropertyProperties{
	// 		BillingAccountAgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftOnlineServicesProgram),
	// 		BillingAccountDisplayName: to.Ptr("My Account"),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000"),
	// 		BillingAccountSoldToCountry: to.Ptr("US"),
	// 		BillingAccountType: to.Ptr(armbilling.AccountTypeIndividual),
	// 		SubscriptionBillingType: to.Ptr(armbilling.SubscriptionBillingTypeFree),
	// 		SubscriptionServiceUsageAddress: &armbilling.PropertyPropertiesSubscriptionServiceUsageAddress{
	// 			AddressLine1: to.Ptr("Address line 1"),
	// 			AddressLine2: to.Ptr("Address line 2"),
	// 			City: to.Ptr("City"),
	// 			Country: to.Ptr("US"),
	// 			FirstName: to.Ptr("Jenny"),
	// 			LastName: to.Ptr("Doe"),
	// 			MiddleName: to.Ptr("Ann"),
	// 			PostalCode: to.Ptr("12345"),
	// 			Region: to.Ptr("State"),
	// 		},
	// 		SubscriptionWorkloadType: to.Ptr(armbilling.SubscriptionWorkloadTypeProduction),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyGetMPA.json
func ExamplePropertyClient_Get_billingPropertyGetMpa() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPropertyClient().Get(ctx, &armbilling.PropertyClientGetOptions{IncludeBillingCountry: nil,
		IncludeTransitionStatus: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Property = armbilling.Property{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingProperty"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.Billing/billingProperty/default"),
	// 	Properties: &armbilling.PropertyProperties{
	// 		BillingAccountAgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
	// 		BillingAccountDisplayName: to.Ptr("My Account"),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000"),
	// 		BillingAccountSoldToCountry: to.Ptr("US"),
	// 		BillingAccountType: to.Ptr(armbilling.AccountTypeBusiness),
	// 		BillingProfileDisplayName: to.Ptr("Contoso"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000/billingProfiles/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 		BillingProfileSpendingLimit: to.Ptr(armbilling.SpendingLimitOff),
	// 		BillingProfileStatus: to.Ptr(armbilling.BillingProfileStatusActive),
	// 		CostCenter: to.Ptr("1010"),
	// 		CustomerDisplayName: to.Ptr("Operations"),
	// 		CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000/billingProfiles/billingProfiles/xxxx-xxxx-xxx-xxx/customers/yyyy-yyyy-yyy-yyy"),
	// 		SKUDescription: to.Ptr("Microsoft Azure Plan"),
	// 		SKUID: to.Ptr("0001"),
	// 		SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusActive),
	// 		SubscriptionBillingStatusDetails: []*armbilling.SubscriptionStatusDetails{
	// 			{
	// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-01T17:32:28.000Z"); return t}()),
	// 				Reason: to.Ptr(armbilling.SubscriptionStatusReasonCancelled),
	// 		}},
	// 		SubscriptionBillingType: to.Ptr(armbilling.SubscriptionBillingTypeFree),
	// 		SubscriptionWorkloadType: to.Ptr(armbilling.SubscriptionWorkloadTypeProduction),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyPatchCostCenter.json
func ExamplePropertyClient_Update_billingPropertyPatchCostCenter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPropertyClient().Update(ctx, armbilling.Property{
		Properties: &armbilling.PropertyProperties{
			CostCenter: to.Ptr("1010"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Property = armbilling.Property{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingProperty"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.Billing/billingProperty/default"),
	// 	Properties: &armbilling.PropertyProperties{
	// 		BillingAccountAgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
	// 		BillingAccountDisplayName: to.Ptr("My Account"),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000"),
	// 		BillingAccountSoldToCountry: to.Ptr("US"),
	// 		BillingAccountType: to.Ptr(armbilling.AccountTypeBusiness),
	// 		BillingProfileDisplayName: to.Ptr("Contoso"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000/billingProfiles/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 		BillingProfileSpendingLimit: to.Ptr(armbilling.SpendingLimitOff),
	// 		BillingProfileStatus: to.Ptr(armbilling.BillingProfileStatusActive),
	// 		CostCenter: to.Ptr("1010"),
	// 		InvoiceSectionDisplayName: to.Ptr("Operations"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000/billingProfiles/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
	// 		SKUDescription: to.Ptr("Microsoft Azure Plan"),
	// 		SKUID: to.Ptr("0001"),
	// 		SubscriptionBillingStatus: to.Ptr(armbilling.BillingSubscriptionStatusActive),
	// 		SubscriptionBillingStatusDetails: []*armbilling.SubscriptionStatusDetails{
	// 			{
	// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-01T17:32:28.000Z"); return t}()),
	// 				Reason: to.Ptr(armbilling.SubscriptionStatusReasonCancelled),
	// 		}},
	// 		SubscriptionBillingType: to.Ptr(armbilling.SubscriptionBillingTypeFree),
	// 		SubscriptionWorkloadType: to.Ptr(armbilling.SubscriptionWorkloadTypeProduction),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyPatchSubscriptionServiceUsageAddress.json
func ExamplePropertyClient_Update_billingPropertyPatchSubscriptionServiceUsageAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPropertyClient().Update(ctx, armbilling.Property{
		Properties: &armbilling.PropertyProperties{
			SubscriptionServiceUsageAddress: &armbilling.PropertyPropertiesSubscriptionServiceUsageAddress{
				AddressLine1: to.Ptr("Address line 1"),
				AddressLine2: to.Ptr("Address line 2"),
				City:         to.Ptr("City"),
				Country:      to.Ptr("US"),
				FirstName:    to.Ptr("Jenny"),
				LastName:     to.Ptr("Doe"),
				MiddleName:   to.Ptr("Ann"),
				PostalCode:   to.Ptr("12345"),
				Region:       to.Ptr("State"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Property = armbilling.Property{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingProperty"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.Billing/billingProperty/default"),
	// 	Properties: &armbilling.PropertyProperties{
	// 		BillingAccountAgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftOnlineServicesProgram),
	// 		BillingAccountDisplayName: to.Ptr("My Account"),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000"),
	// 		BillingAccountSoldToCountry: to.Ptr("US"),
	// 		BillingAccountType: to.Ptr(armbilling.AccountTypeIndividual),
	// 		SubscriptionBillingType: to.Ptr(armbilling.SubscriptionBillingTypeFree),
	// 		SubscriptionServiceUsageAddress: &armbilling.PropertyPropertiesSubscriptionServiceUsageAddress{
	// 			AddressLine1: to.Ptr("Address line 1"),
	// 			AddressLine2: to.Ptr("Address line 2"),
	// 			City: to.Ptr("City"),
	// 			Country: to.Ptr("US"),
	// 			FirstName: to.Ptr("Jenny"),
	// 			LastName: to.Ptr("Doe"),
	// 			MiddleName: to.Ptr("Ann"),
	// 			PostalCode: to.Ptr("12345"),
	// 			Region: to.Ptr("State"),
	// 		},
	// 		SubscriptionWorkloadType: to.Ptr(armbilling.SubscriptionWorkloadTypeProduction),
	// 	},
	// }
}