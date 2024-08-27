//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/SourceControls_ListByContainer.json
func ExampleContainerAppsSourceControlsClient_NewListByContainerAppPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsSourceControlsClient().NewListByContainerAppPager("workerapps-rg-xj", "testcanadacentral", nil)
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
		// page.SourceControlCollection = armappcontainers.SourceControlCollection{
		// 	Value: []*armappcontainers.SourceControl{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.App/containerapps/sourcecontrols"),
		// 			ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/testcanadacentral/sourcecontrols/current"),
		// 			Properties: &armappcontainers.SourceControlProperties{
		// 				Branch: to.Ptr("master"),
		// 				GithubActionConfiguration: &armappcontainers.GithubActionConfiguration{
		// 					BuildEnvironmentVariables: []*armappcontainers.EnvironmentVariable{
		// 						{
		// 							Name: to.Ptr("foo1"),
		// 							Value: to.Ptr("bar1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("foo2"),
		// 							Value: to.Ptr("bar2"),
		// 					}},
		// 					ContextPath: to.Ptr("./"),
		// 					Image: to.Ptr("image/tag"),
		// 					RegistryInfo: &armappcontainers.RegistryInfo{
		// 						RegistryURL: to.Ptr("xwang971reg.azurecr.io"),
		// 						RegistryUserName: to.Ptr("xwang971reg"),
		// 					},
		// 				},
		// 				RepoURL: to.Ptr("https://github.com/xwang971/ghatest"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/SourceControls_Get.json
func ExampleContainerAppsSourceControlsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsSourceControlsClient().Get(ctx, "workerapps-rg-xj", "testcanadacentral", "current", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SourceControl = armappcontainers.SourceControl{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.App/containerapps/sourcecontrols"),
	// 	ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/testcanadacentral/sourcecontrols/current"),
	// 	Properties: &armappcontainers.SourceControlProperties{
	// 		Branch: to.Ptr("master"),
	// 		GithubActionConfiguration: &armappcontainers.GithubActionConfiguration{
	// 			BuildEnvironmentVariables: []*armappcontainers.EnvironmentVariable{
	// 				{
	// 					Name: to.Ptr("foo1"),
	// 					Value: to.Ptr("bar1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("foo2"),
	// 					Value: to.Ptr("bar2"),
	// 			}},
	// 			ContextPath: to.Ptr("./"),
	// 			Image: to.Ptr("image/tag"),
	// 			RegistryInfo: &armappcontainers.RegistryInfo{
	// 				RegistryURL: to.Ptr("xwang971reg.azurecr.io"),
	// 				RegistryUserName: to.Ptr("xwang971reg"),
	// 			},
	// 		},
	// 		RepoURL: to.Ptr("https://github.com/xwang971/ghatest"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/SourceControls_CreateOrUpdate.json
func ExampleContainerAppsSourceControlsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsSourceControlsClient().BeginCreateOrUpdate(ctx, "workerapps-rg-xj", "testcanadacentral", "current", armappcontainers.SourceControl{
		Properties: &armappcontainers.SourceControlProperties{
			Branch: to.Ptr("master"),
			GithubActionConfiguration: &armappcontainers.GithubActionConfiguration{
				AzureCredentials: &armappcontainers.AzureCredentials{
					ClientID:     to.Ptr("<clientid>"),
					ClientSecret: to.Ptr("<clientsecret>"),
					Kind:         to.Ptr("feaderated"),
					TenantID:     to.Ptr("<tenantid>"),
				},
				BuildEnvironmentVariables: []*armappcontainers.EnvironmentVariable{
					{
						Name:  to.Ptr("foo1"),
						Value: to.Ptr("bar1"),
					},
					{
						Name:  to.Ptr("foo2"),
						Value: to.Ptr("bar2"),
					}},
				ContextPath:               to.Ptr("./"),
				GithubPersonalAccessToken: to.Ptr("test"),
				Image:                     to.Ptr("image/tag"),
				RegistryInfo: &armappcontainers.RegistryInfo{
					RegistryPassword: to.Ptr("<registrypassword>"),
					RegistryURL:      to.Ptr("test-registry.azurecr.io"),
					RegistryUserName: to.Ptr("test-registry"),
				},
			},
			RepoURL: to.Ptr("https://github.com/xwang971/ghatest"),
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
	// res.SourceControl = armappcontainers.SourceControl{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.App/containerapps/sourcecontrols"),
	// 	ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/myapp/sourcecontrols/current"),
	// 	Properties: &armappcontainers.SourceControlProperties{
	// 		Branch: to.Ptr("master"),
	// 		GithubActionConfiguration: &armappcontainers.GithubActionConfiguration{
	// 			BuildEnvironmentVariables: []*armappcontainers.EnvironmentVariable{
	// 				{
	// 					Name: to.Ptr("foo1"),
	// 					Value: to.Ptr("bar1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("foo2"),
	// 					Value: to.Ptr("bar2"),
	// 			}},
	// 			ContextPath: to.Ptr("./"),
	// 			Image: to.Ptr("image/tag"),
	// 			RegistryInfo: &armappcontainers.RegistryInfo{
	// 				RegistryURL: to.Ptr("test-registry.azurecr.io"),
	// 				RegistryUserName: to.Ptr("testreg"),
	// 			},
	// 		},
	// 		OperationState: to.Ptr(armappcontainers.SourceControlOperationStateInProgress),
	// 		RepoURL: to.Ptr("https://github.com/xwang971/ghatest"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/SourceControls_Delete.json
func ExampleContainerAppsSourceControlsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsSourceControlsClient().BeginDelete(ctx, "workerapps-rg-xj", "testcanadacentral", "current", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}