//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

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

// ServersClient contains the methods for the AnalysisServicesServers group.
// Don't use this type directly, use NewServersClient() instead.
type ServersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServersClient creates a new instance of ServersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Provisions the specified Analysis Services server based on the configuration specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - resource - Resource create parameters.
//   - options - ServersClientBeginCreateOptions contains the optional parameters for the ServersClient.BeginCreate method.
func (client *ServersClient) BeginCreate(ctx context.Context, resourceGroupName string, serverName string, resource Server, options *ServersClientBeginCreateOptions) (*runtime.Poller[ServersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, serverName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServersClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServersClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Provisions the specified Analysis Services server based on the configuration specified in the request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
func (client *ServersClient) create(ctx context.Context, resourceGroupName string, serverName string, resource Server, options *ServersClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServersClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, serverName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ServersClient) createCreateRequest(ctx context.Context, resourceGroupName string, serverName string, resource Server, options *ServersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified Analysis Services server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - options - ServersClientBeginDeleteOptions contains the optional parameters for the ServersClient.BeginDelete method.
func (client *ServersClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, options *ServersClientBeginDeleteOptions) (*runtime.Poller[ServersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified Analysis Services server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
func (client *ServersClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, options *ServersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ServersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DissociateGateway - Dissociates a Unified Gateway associated with the server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - body - The content of the action request
//   - options - ServersClientDissociateGatewayOptions contains the optional parameters for the ServersClient.DissociateGateway
//     method.
func (client *ServersClient) DissociateGateway(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientDissociateGatewayOptions) (ServersClientDissociateGatewayResponse, error) {
	var err error
	const operationName = "ServersClient.DissociateGateway"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.dissociateGatewayCreateRequest(ctx, resourceGroupName, serverName, body, options)
	if err != nil {
		return ServersClientDissociateGatewayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServersClientDissociateGatewayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServersClientDissociateGatewayResponse{}, err
	}
	resp, err := client.dissociateGatewayHandleResponse(httpResp)
	return resp, err
}

// dissociateGatewayCreateRequest creates the DissociateGateway request.
func (client *ServersClient) dissociateGatewayCreateRequest(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientDissociateGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/dissociateGateway"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// dissociateGatewayHandleResponse handles the DissociateGateway response.
func (client *ServersClient) dissociateGatewayHandleResponse(resp *http.Response) (ServersClientDissociateGatewayResponse, error) {
	result := ServersClientDissociateGatewayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return ServersClientDissociateGatewayResponse{}, err
	}
	return result, nil
}

// GetDetails - Gets details about the specified Analysis Services server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - options - ServersClientGetDetailsOptions contains the optional parameters for the ServersClient.GetDetails method.
func (client *ServersClient) GetDetails(ctx context.Context, resourceGroupName string, serverName string, options *ServersClientGetDetailsOptions) (ServersClientGetDetailsResponse, error) {
	var err error
	const operationName = "ServersClient.GetDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDetailsCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServersClientGetDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServersClientGetDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServersClientGetDetailsResponse{}, err
	}
	resp, err := client.getDetailsHandleResponse(httpResp)
	return resp, err
}

// getDetailsCreateRequest creates the GetDetails request.
func (client *ServersClient) getDetailsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServersClientGetDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDetailsHandleResponse handles the GetDetails response.
func (client *ServersClient) getDetailsHandleResponse(resp *http.Response) (ServersClientGetDetailsResponse, error) {
	result := ServersClientGetDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Server); err != nil {
		return ServersClientGetDetailsResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all the Analysis Services servers for the given resource group.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ServersClientListByResourceGroupOptions contains the optional parameters for the ServersClient.NewListByResourceGroupPager
//     method.
func (client *ServersClient) NewListByResourceGroupPager(resourceGroupName string, options *ServersClientListByResourceGroupOptions) *runtime.Pager[ServersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServersClientListByResourceGroupResponse]{
		More: func(page ServersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServersClientListByResourceGroupResponse) (ServersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ServersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServersClient) listByResourceGroupHandleResponse(resp *http.Response) (ServersClientListByResourceGroupResponse, error) {
	result := ServersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerListResult); err != nil {
		return ServersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListGatewayStatus - Return the gateway status of the specified Analysis Services server instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - body - The content of the action request
//   - options - ServersClientListGatewayStatusOptions contains the optional parameters for the ServersClient.ListGatewayStatus
//     method.
func (client *ServersClient) ListGatewayStatus(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientListGatewayStatusOptions) (ServersClientListGatewayStatusResponse, error) {
	var err error
	const operationName = "ServersClient.ListGatewayStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listGatewayStatusCreateRequest(ctx, resourceGroupName, serverName, body, options)
	if err != nil {
		return ServersClientListGatewayStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServersClientListGatewayStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServersClientListGatewayStatusResponse{}, err
	}
	resp, err := client.listGatewayStatusHandleResponse(httpResp)
	return resp, err
}

// listGatewayStatusCreateRequest creates the ListGatewayStatus request.
func (client *ServersClient) listGatewayStatusCreateRequest(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientListGatewayStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/listGatewayStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// listGatewayStatusHandleResponse handles the ListGatewayStatus response.
func (client *ServersClient) listGatewayStatusHandleResponse(resp *http.Response) (ServersClientListGatewayStatusResponse, error) {
	result := ServersClientListGatewayStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayListStatusLive); err != nil {
		return ServersClientListGatewayStatusResponse{}, err
	}
	return result, nil
}

// ListSKUsForExisting - Lists eligible SKUs for an Analysis Services resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the Azure Resource group of which a given Analysis Services server is part. This name must
//     be at least 1 character in length, and no more than 90.
//   - serverName - The name of the Analysis Services server. It must be at least 3 characters in length, and no more than 63.
//   - subscriptionID - A unique identifier for a Microsoft Azure subscription. The subscription ID forms part of the URI for
//     every service call.
//   - options - ServersClientListSKUsForExistingOptions contains the optional parameters for the ServersClient.ListSKUsForExisting
//     method.
func (client *ServersClient) ListSKUsForExisting(ctx context.Context, resourceGroupName string, serverName string, subscriptionID string, options *ServersClientListSKUsForExistingOptions) (ServersClientListSKUsForExistingResponse, error) {
	var err error
	const operationName = "ServersClient.ListSKUsForExisting"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSKUsForExistingCreateRequest(ctx, resourceGroupName, serverName, subscriptionID, options)
	if err != nil {
		return ServersClientListSKUsForExistingResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServersClientListSKUsForExistingResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServersClientListSKUsForExistingResponse{}, err
	}
	resp, err := client.listSKUsForExistingHandleResponse(httpResp)
	return resp, err
}

// listSKUsForExistingCreateRequest creates the ListSKUsForExisting request.
func (client *ServersClient) listSKUsForExistingCreateRequest(ctx context.Context, resourceGroupName string, serverName string, subscriptionID string, options *ServersClientListSKUsForExistingOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/skus/{resourceGroupName}/{serverName}/{subscriptionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsForExistingHandleResponse handles the ListSKUsForExisting response.
func (client *ServersClient) listSKUsForExistingHandleResponse(resp *http.Response) (ServersClientListSKUsForExistingResponse, error) {
	result := ServersClientListSKUsForExistingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUEnumerationForExistingResourceResult); err != nil {
		return ServersClientListSKUsForExistingResponse{}, err
	}
	return result, nil
}

// BeginResume - Resumes operation of the specified Analysis Services server instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - body - The content of the action request
//   - options - ServersClientBeginResumeOptions contains the optional parameters for the ServersClient.BeginResume method.
func (client *ServersClient) BeginResume(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientBeginResumeOptions) (*runtime.Poller[ServersClientResumeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resume(ctx, resourceGroupName, serverName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServersClientResumeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServersClientResumeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Resume - Resumes operation of the specified Analysis Services server instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
func (client *ServersClient) resume(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientBeginResumeOptions) (*http.Response, error) {
	var err error
	const operationName = "ServersClient.BeginResume"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, serverName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *ServersClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/resume"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginSuspend - Suspends operation of the specified Analysis Services server instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - body - The content of the action request
//   - options - ServersClientBeginSuspendOptions contains the optional parameters for the ServersClient.BeginSuspend method.
func (client *ServersClient) BeginSuspend(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientBeginSuspendOptions) (*runtime.Poller[ServersClientSuspendResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.suspend(ctx, resourceGroupName, serverName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServersClientSuspendResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServersClientSuspendResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Suspend - Suspends operation of the specified Analysis Services server instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
func (client *ServersClient) suspend(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientBeginSuspendOptions) (*http.Response, error) {
	var err error
	const operationName = "ServersClient.BeginSuspend"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.suspendCreateRequest(ctx, resourceGroupName, serverName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// suspendCreateRequest creates the Suspend request.
func (client *ServersClient) suspendCreateRequest(ctx context.Context, resourceGroupName string, serverName string, body any, options *ServersClientBeginSuspendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/suspend"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Updates the current state of the specified Analysis Services server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
//   - properties - The resource properties to be updated.
//   - options - ServersClientBeginUpdateOptions contains the optional parameters for the ServersClient.BeginUpdate method.
func (client *ServersClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerUpdate, options *ServersClientBeginUpdateOptions) (*runtime.Poller[ServersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serverName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the current state of the specified Analysis Services server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01
func (client *ServersClient) update(ctx context.Context, resourceGroupName string, serverName string, properties ServerUpdate, options *ServersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ServersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, properties ServerUpdate, options *ServersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
