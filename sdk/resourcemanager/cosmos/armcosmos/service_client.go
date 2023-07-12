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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServiceClient contains the methods for the Service group.
// Don't use this type directly, use NewServiceClient() instead.
type ServiceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceClient creates a new instance of ServiceClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceClient, error) {
	cl, err := arm.NewClient(moduleName+".ServiceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - serviceName - Cosmos DB service name.
//   - createUpdateParameters - The Service resource parameters.
//   - options - ServiceClientBeginCreateOptions contains the optional parameters for the ServiceClient.BeginCreate method.
func (client *ServiceClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, serviceName string, createUpdateParameters ServiceResourceCreateUpdateParameters, options *ServiceClientBeginCreateOptions) (*runtime.Poller[ServiceClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, serviceName, createUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
func (client *ServiceClient) create(ctx context.Context, resourceGroupName string, accountName string, serviceName string, createUpdateParameters ServiceResourceCreateUpdateParameters, options *ServiceClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, serviceName, createUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ServiceClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, serviceName string, createUpdateParameters ServiceResourceCreateUpdateParameters, options *ServiceClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/services/{serviceName}"
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
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, createUpdateParameters)
}

// BeginDelete - Deletes service with the given serviceName.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - serviceName - Cosmos DB service name.
//   - options - ServiceClientBeginDeleteOptions contains the optional parameters for the ServiceClient.BeginDelete method.
func (client *ServiceClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, serviceName string, options *ServiceClientBeginDeleteOptions) (*runtime.Poller[ServiceClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, serviceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes service with the given serviceName.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
func (client *ServiceClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, serviceName string, options *ServiceClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, serviceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, serviceName string, options *ServiceClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/services/{serviceName}"
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
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the status of service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - serviceName - Cosmos DB service name.
//   - options - ServiceClientGetOptions contains the optional parameters for the ServiceClient.Get method.
func (client *ServiceClient) Get(ctx context.Context, resourceGroupName string, accountName string, serviceName string, options *ServiceClientGetOptions) (ServiceClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, serviceName, options)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, serviceName string, options *ServiceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/services/{serviceName}"
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
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceClient) getHandleResponse(resp *http.Response) (ServiceClientGetResponse, error) {
	result := ServiceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServiceClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the status of service.
//
// Generated from API version 2023-08-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - options - ServiceClientListOptions contains the optional parameters for the ServiceClient.NewListPager method.
func (client *ServiceClient) NewListPager(resourceGroupName string, accountName string, options *ServiceClientListOptions) *runtime.Pager[ServiceClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceClientListResponse]{
		More: func(page ServiceClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServiceClientListResponse) (ServiceClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return ServiceClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *ServiceClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/services"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceClient) listHandleResponse(resp *http.Response) (ServiceClientListResponse, error) {
	result := ServiceClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceListResult); err != nil {
		return ServiceClientListResponse{}, err
	}
	return result, nil
}
