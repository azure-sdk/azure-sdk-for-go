//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ExpressRoutePortAuthorizationsClient contains the methods for the ExpressRoutePortAuthorizations group.
// Don't use this type directly, use NewExpressRoutePortAuthorizationsClient() instead.
type ExpressRoutePortAuthorizationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRoutePortAuthorizationsClient creates a new instance of ExpressRoutePortAuthorizationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRoutePortAuthorizationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRoutePortAuthorizationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRoutePortAuthorizationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an authorization in the specified express route port.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the express route port.
// authorizationName - The name of the authorization.
// authorizationParameters - Parameters supplied to the create or update express route port authorization operation.
// options - ExpressRoutePortAuthorizationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRoutePortAuthorizationsClient.BeginCreateOrUpdate
// method.
func (client *ExpressRoutePortAuthorizationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, authorizationParameters ExpressRoutePortAuthorization, options *ExpressRoutePortAuthorizationsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ExpressRoutePortAuthorizationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, expressRoutePortName, authorizationName, authorizationParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ExpressRoutePortAuthorizationsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ExpressRoutePortAuthorizationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an authorization in the specified express route port.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRoutePortAuthorizationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, authorizationParameters ExpressRoutePortAuthorization, options *ExpressRoutePortAuthorizationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRoutePortName, authorizationName, authorizationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRoutePortAuthorizationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, authorizationParameters ExpressRoutePortAuthorization, options *ExpressRoutePortAuthorizationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, authorizationParameters)
}

// BeginDelete - Deletes the specified authorization from the specified express route port.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the express route port.
// authorizationName - The name of the authorization.
// options - ExpressRoutePortAuthorizationsClientBeginDeleteOptions contains the optional parameters for the ExpressRoutePortAuthorizationsClient.BeginDelete
// method.
func (client *ExpressRoutePortAuthorizationsClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, options *ExpressRoutePortAuthorizationsClientBeginDeleteOptions) (*armruntime.Poller[ExpressRoutePortAuthorizationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, expressRoutePortName, authorizationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ExpressRoutePortAuthorizationsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ExpressRoutePortAuthorizationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified authorization from the specified express route port.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRoutePortAuthorizationsClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, options *ExpressRoutePortAuthorizationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRoutePortName, authorizationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRoutePortAuthorizationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, options *ExpressRoutePortAuthorizationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified authorization from the specified express route port.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the express route port.
// authorizationName - The name of the authorization.
// options - ExpressRoutePortAuthorizationsClientGetOptions contains the optional parameters for the ExpressRoutePortAuthorizationsClient.Get
// method.
func (client *ExpressRoutePortAuthorizationsClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, options *ExpressRoutePortAuthorizationsClientGetOptions) (ExpressRoutePortAuthorizationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRoutePortName, authorizationName, options)
	if err != nil {
		return ExpressRoutePortAuthorizationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRoutePortAuthorizationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRoutePortAuthorizationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRoutePortAuthorizationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, options *ExpressRoutePortAuthorizationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRoutePortAuthorizationsClient) getHandleResponse(resp *http.Response) (ExpressRoutePortAuthorizationsClientGetResponse, error) {
	result := ExpressRoutePortAuthorizationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortAuthorization); err != nil {
		return ExpressRoutePortAuthorizationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all authorizations in an express route port.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the express route port.
// options - ExpressRoutePortAuthorizationsClientListOptions contains the optional parameters for the ExpressRoutePortAuthorizationsClient.List
// method.
func (client *ExpressRoutePortAuthorizationsClient) NewListPager(resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortAuthorizationsClientListOptions) *runtime.Pager[ExpressRoutePortAuthorizationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ExpressRoutePortAuthorizationsClientListResponse]{
		More: func(page ExpressRoutePortAuthorizationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRoutePortAuthorizationsClientListResponse) (ExpressRoutePortAuthorizationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRoutePortAuthorizationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExpressRoutePortAuthorizationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRoutePortAuthorizationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRoutePortAuthorizationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortAuthorizationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRoutePortAuthorizationsClient) listHandleResponse(resp *http.Response) (ExpressRoutePortAuthorizationsClientListResponse, error) {
	result := ExpressRoutePortAuthorizationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortAuthorizationListResult); err != nil {
		return ExpressRoutePortAuthorizationsClientListResponse{}, err
	}
	return result, nil
}
