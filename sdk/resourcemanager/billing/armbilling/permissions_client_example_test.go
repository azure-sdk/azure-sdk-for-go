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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByBillingAccount.json
func ExamplePermissionsClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByBillingAccountPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", nil)
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
		// page.PermissionListResult = armbilling.PermissionListResult{
		// 	Value: []*armbilling.Permission{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000006"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000003"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000002"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000001"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 				to.Ptr("20000000-aaaa-bbbb-cccc-200000000002"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("20000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000008"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000001")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByBillingProfile.json
func ExamplePermissionsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByBillingProfilePager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", nil)
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
		// page.PermissionListResult = armbilling.PermissionListResult{
		// 	Value: []*armbilling.Permission{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000006"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000003"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000002"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000001"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 				to.Ptr("20000000-aaaa-bbbb-cccc-200000000002"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("20000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000008"),
		// 				to.Ptr("10000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 				to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000001")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByBillingProfile.json
func ExamplePermissionsClient_CheckAccessByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPermissionsClient().CheckAccessByBillingProfile(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", armbilling.CheckAccessRequest{
		Actions: []*string{
			to.Ptr("Microsoft.Billing/billingAccounts/read"),
			to.Ptr("Microsoft.Subscription/subscriptions/write")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckAccessResponseArray = []*armbilling.CheckAccessResponse{
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionAllowed),
	// 		Action: to.Ptr("Microsoft.Billing/billingAccounts/read"),
	// 	},
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionNotAllowed),
	// 		Action: to.Ptr("Microsoft.Subscription/subscriptions/write"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByCustomer.json
func ExamplePermissionsClient_NewListByCustomerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByCustomerPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "BKM6-54VH-BG7-PGB", "11111111-1111-1111-1111-111111111111", nil)
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
		// page.PermissionListResult = armbilling.PermissionListResult{
		// 	Value: []*armbilling.Permission{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000001")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByCustomer.json
func ExamplePermissionsClient_CheckAccessByCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPermissionsClient().CheckAccessByCustomer(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "703ab484-dda2-4402-827b-a74513b61e2d", armbilling.CheckAccessRequest{
		Actions: []*string{
			to.Ptr("Microsoft.Billing/billingAccounts/read"),
			to.Ptr("Microsoft.Subscription/subscriptions/write")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckAccessResponseArray = []*armbilling.CheckAccessResponse{
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionAllowed),
	// 		Action: to.Ptr("Microsoft.Billing/billingAccounts/read"),
	// 	},
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionNotAllowed),
	// 		Action: to.Ptr("Microsoft.Subscription/subscriptions/write"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByInvoiceSection.json
func ExamplePermissionsClient_NewListByInvoiceSectionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByInvoiceSectionPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "XXXX-XXXX-XXX-XXX", nil)
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
		// page.PermissionListResult = armbilling.PermissionListResult{
		// 	Value: []*armbilling.Permission{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000001")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByInvoiceSection.json
func ExamplePermissionsClient_CheckAccessByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPermissionsClient().CheckAccessByInvoiceSection(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "Q7GV-UUVA-PJA-TGB", armbilling.CheckAccessRequest{
		Actions: []*string{
			to.Ptr("Microsoft.Billing/billingAccounts/read"),
			to.Ptr("Microsoft.Subscription/subscriptions/write")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckAccessResponseArray = []*armbilling.CheckAccessResponse{
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionAllowed),
	// 		Action: to.Ptr("Microsoft.Billing/billingAccounts/read"),
	// 	},
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionNotAllowed),
	// 		Action: to.Ptr("Microsoft.Subscription/subscriptions/write"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByBillingAccount.json
func ExamplePermissionsClient_CheckAccessByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPermissionsClient().CheckAccessByBillingAccount(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", armbilling.CheckAccessRequest{
		Actions: []*string{
			to.Ptr("Microsoft.Billing/billingAccounts/read"),
			to.Ptr("Microsoft.Subscription/subscriptions/write")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckAccessResponseArray = []*armbilling.CheckAccessResponse{
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionAllowed),
	// 		Action: to.Ptr("Microsoft.Billing/billingAccounts/read"),
	// 	},
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionNotAllowed),
	// 		Action: to.Ptr("Microsoft.Subscription/subscriptions/write"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByCustomerAtBillingAccount.json
func ExamplePermissionsClient_NewListByCustomerAtBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByCustomerAtBillingAccountPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "11111111-1111-1111-1111-111111111111", nil)
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
		// page.PermissionListResult = armbilling.PermissionListResult{
		// 	Value: []*armbilling.Permission{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 				to.Ptr("30000000-aaaa-bbbb-cccc-200000000001")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByDepartment.json
func ExamplePermissionsClient_NewListByDepartmentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByDepartmentPager("6100092", "123456", nil)
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
		// page.PermissionListResult = armbilling.PermissionListResult{
		// 	Value: []*armbilling.Permission{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("Microsoft.Billing/billingAccounts/departments/read"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/departments/write"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/read"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/write"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/read"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/write")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByDepartment.json
func ExamplePermissionsClient_CheckAccessByDepartment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPermissionsClient().CheckAccessByDepartment(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "12345", armbilling.CheckAccessRequest{
		Actions: []*string{
			to.Ptr("Microsoft.Billing/billingAccounts/read"),
			to.Ptr("Microsoft.Subscription/subscriptions/write")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckAccessResponseArray = []*armbilling.CheckAccessResponse{
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionAllowed),
	// 		Action: to.Ptr("Microsoft.Billing/billingAccounts/read"),
	// 	},
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionNotAllowed),
	// 		Action: to.Ptr("Microsoft.Subscription/subscriptions/write"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByEnrollmentAccount.json
func ExamplePermissionsClient_NewListByEnrollmentAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPermissionsClient().NewListByEnrollmentAccountPager("6100092", "123456", nil)
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
		// page.PermissionListResult = armbilling.PermissionListResult{
		// 	Value: []*armbilling.Permission{
		// 		{
		// 			Actions: []*string{
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/read"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/write"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/read"),
		// 				to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/write")},
		// 				NotActions: []*string{
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByEnrollmentAccount.json
func ExamplePermissionsClient_CheckAccessByEnrollmentAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPermissionsClient().CheckAccessByEnrollmentAccount(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "123456", armbilling.CheckAccessRequest{
		Actions: []*string{
			to.Ptr("Microsoft.Billing/billingAccounts/read"),
			to.Ptr("Microsoft.Subscription/subscriptions/write")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckAccessResponseArray = []*armbilling.CheckAccessResponse{
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionAllowed),
	// 		Action: to.Ptr("Microsoft.Billing/billingAccounts/read"),
	// 	},
	// 	{
	// 		AccessDecision: to.Ptr(armbilling.AccessDecisionNotAllowed),
	// 		Action: to.Ptr("Microsoft.Subscription/subscriptions/write"),
	// }}
}