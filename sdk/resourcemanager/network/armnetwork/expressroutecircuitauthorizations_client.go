//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// ExpressRouteCircuitAuthorizationsClient contains the methods for the ExpressRouteCircuitAuthorizations group.
// Don't use this type directly, use NewExpressRouteCircuitAuthorizationsClient() instead.
type ExpressRouteCircuitAuthorizationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExpressRouteCircuitAuthorizationsClient creates a new instance of ExpressRouteCircuitAuthorizationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExpressRouteCircuitAuthorizationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteCircuitAuthorizationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ExpressRouteCircuitAuthorizationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRouteCircuitAuthorizationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an authorization in the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceGroupName - The name of the resource group.
//   - circuitName - The name of the express route circuit.
//   - authorizationName - The name of the authorization.
//   - authorizationParameters - Parameters supplied to the create or update express route circuit authorization operation.
//   - options - ExpressRouteCircuitAuthorizationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCircuitAuthorizationsClient.BeginCreateOrUpdate
//     method.
func (client *ExpressRouteCircuitAuthorizationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExpressRouteCircuitAuthorizationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, circuitName, authorizationName, authorizationParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteCircuitAuthorizationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCircuitAuthorizationsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an authorization in the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ExpressRouteCircuitAuthorizationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, authorizationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCircuitAuthorizationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, authorizationParameters)
}

// BeginDelete - Deletes the specified authorization from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceGroupName - The name of the resource group.
//   - circuitName - The name of the express route circuit.
//   - authorizationName - The name of the authorization.
//   - options - ExpressRouteCircuitAuthorizationsClientBeginDeleteOptions contains the optional parameters for the ExpressRouteCircuitAuthorizationsClient.BeginDelete
//     method.
func (client *ExpressRouteCircuitAuthorizationsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsClientBeginDeleteOptions) (*runtime.Poller[ExpressRouteCircuitAuthorizationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, circuitName, authorizationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteCircuitAuthorizationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCircuitAuthorizationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified authorization from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ExpressRouteCircuitAuthorizationsClient) deleteOperation(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, options)
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
func (client *ExpressRouteCircuitAuthorizationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified authorization from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceGroupName - The name of the resource group.
//   - circuitName - The name of the express route circuit.
//   - authorizationName - The name of the authorization.
//   - options - ExpressRouteCircuitAuthorizationsClientGetOptions contains the optional parameters for the ExpressRouteCircuitAuthorizationsClient.Get
//     method.
func (client *ExpressRouteCircuitAuthorizationsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsClientGetOptions) (ExpressRouteCircuitAuthorizationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, options)
	if err != nil {
		return ExpressRouteCircuitAuthorizationsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCircuitAuthorizationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCircuitAuthorizationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCircuitAuthorizationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCircuitAuthorizationsClient) getHandleResponse(resp *http.Response) (ExpressRouteCircuitAuthorizationsClientGetResponse, error) {
	result := ExpressRouteCircuitAuthorizationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitAuthorization); err != nil {
		return ExpressRouteCircuitAuthorizationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all authorizations in an express route circuit.
//
// Generated from API version 2023-02-01
//   - resourceGroupName - The name of the resource group.
//   - circuitName - The name of the circuit.
//   - options - ExpressRouteCircuitAuthorizationsClientListOptions contains the optional parameters for the ExpressRouteCircuitAuthorizationsClient.NewListPager
//     method.
func (client *ExpressRouteCircuitAuthorizationsClient) NewListPager(resourceGroupName string, circuitName string, options *ExpressRouteCircuitAuthorizationsClientListOptions) *runtime.Pager[ExpressRouteCircuitAuthorizationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRouteCircuitAuthorizationsClientListResponse]{
		More: func(page ExpressRouteCircuitAuthorizationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCircuitAuthorizationsClientListResponse) (ExpressRouteCircuitAuthorizationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, circuitName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRouteCircuitAuthorizationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExpressRouteCircuitAuthorizationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRouteCircuitAuthorizationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCircuitAuthorizationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitAuthorizationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCircuitAuthorizationsClient) listHandleResponse(resp *http.Response) (ExpressRouteCircuitAuthorizationsClientListResponse, error) {
	result := ExpressRouteCircuitAuthorizationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationListResult); err != nil {
		return ExpressRouteCircuitAuthorizationsClientListResponse{}, err
	}
	return result, nil
}
