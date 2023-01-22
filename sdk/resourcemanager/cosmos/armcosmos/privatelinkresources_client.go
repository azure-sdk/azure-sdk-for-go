//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the private link resources that need to be created for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - groupName - The name of the private link resource.
//   - options - PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get
//     method.
func (client *PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, accountName string, groupName string, options *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, groupName, options)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, groupName string, options *PrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateLinkResources/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (PrivateLinkResourcesClientGetResponse, error) {
	result := PrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabaseAccountPager - Gets the private link resources that need to be created for a Cosmos DB account.
//
// Generated from API version 2022-08-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - options - PrivateLinkResourcesClientListByDatabaseAccountOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByDatabaseAccountPager
//     method.
func (client *PrivateLinkResourcesClient) NewListByDatabaseAccountPager(resourceGroupName string, accountName string, options *PrivateLinkResourcesClientListByDatabaseAccountOptions) *runtime.Pager[PrivateLinkResourcesClientListByDatabaseAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientListByDatabaseAccountResponse]{
		More: func(page PrivateLinkResourcesClientListByDatabaseAccountResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientListByDatabaseAccountResponse) (PrivateLinkResourcesClientListByDatabaseAccountResponse, error) {
			req, err := client.listByDatabaseAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return PrivateLinkResourcesClientListByDatabaseAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkResourcesClientListByDatabaseAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkResourcesClientListByDatabaseAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseAccountHandleResponse(resp)
		},
	})
}

// listByDatabaseAccountCreateRequest creates the ListByDatabaseAccount request.
func (client *PrivateLinkResourcesClient) listByDatabaseAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PrivateLinkResourcesClientListByDatabaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseAccountHandleResponse handles the ListByDatabaseAccount response.
func (client *PrivateLinkResourcesClient) listByDatabaseAccountHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByDatabaseAccountResponse, error) {
	result := PrivateLinkResourcesClientListByDatabaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientListByDatabaseAccountResponse{}, err
	}
	return result, nil
}
