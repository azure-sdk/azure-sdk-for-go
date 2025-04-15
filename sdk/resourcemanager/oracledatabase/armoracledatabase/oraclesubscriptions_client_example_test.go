// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armoracledatabase_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
	"log"
)

// Generated from example definition: 2025-03-01/oracleSubscriptions_addAzureSubscriptions.json
func ExampleOracleSubscriptionsClient_BeginAddAzureSubscriptions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOracleSubscriptionsClient().BeginAddAzureSubscriptions(ctx, armoracledatabase.AzureSubscriptions{
		AzureSubscriptionIDs: []*string{
			to.Ptr("00000000-0000-0000-0000-000000000001"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2025-03-01/oracleSubscriptions_listActivationLinks.json
func ExampleOracleSubscriptionsClient_BeginListActivationLinks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOracleSubscriptionsClient().BeginListActivationLinks(ctx, nil)
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
	// res = armoracledatabase.OracleSubscriptionsClientListActivationLinksResponse{
	// }
}

// Generated from example definition: 2025-03-01/oracleSubscriptions_listCloudAccountDetails.json
func ExampleOracleSubscriptionsClient_BeginListCloudAccountDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOracleSubscriptionsClient().BeginListCloudAccountDetails(ctx, nil)
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
	// res = armoracledatabase.OracleSubscriptionsClientListCloudAccountDetailsResponse{
	// }
}

// Generated from example definition: 2025-03-01/oracleSubscriptions_listSaasSubscriptionDetails.json
func ExampleOracleSubscriptionsClient_BeginListSaasSubscriptionDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOracleSubscriptionsClient().BeginListSaasSubscriptionDetails(ctx, nil)
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
	// res = armoracledatabase.OracleSubscriptionsClientListSaasSubscriptionDetailsResponse{
	// }
}

// Generated from example definition: 2025-03-01/oracleSubscriptions_patch.json
func ExampleOracleSubscriptionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOracleSubscriptionsClient().BeginUpdate(ctx, armoracledatabase.OracleSubscriptionUpdate{}, nil)
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
	// res = armoracledatabase.OracleSubscriptionsClientUpdateResponse{
	// 	OracleSubscription: &armoracledatabase.OracleSubscription{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/oracleSubscriptions/default"),
	// 		Type: to.Ptr("Oracle.Database/oracleSubscriptions"),
	// 		Properties: &armoracledatabase.OracleSubscriptionProperties{
	// 			ProvisioningState: to.Ptr(armoracledatabase.OracleSubscriptionProvisioningStateSucceeded),
	// 			SaasSubscriptionID: to.Ptr("saas1"),
	// 			CloudAccountID: to.Ptr("ocid1..aaaaaa"),
	// 			CloudAccountState: to.Ptr(armoracledatabase.CloudAccountProvisioningStatePending),
	// 		},
	// 		Plan: &armoracledatabase.Plan{
	// 			Name: to.Ptr("plan1"),
	// 			Publisher: to.Ptr("publisher1"),
	// 			Product: to.Ptr("product1"),
	// 			PromotionCode: to.Ptr("none"),
	// 			Version: to.Ptr("alpha"),
	// 		},
	// 	},
	// }
}
