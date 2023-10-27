//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// VirtualNetworkGatewayNatRulesClient contains the methods for the VirtualNetworkGatewayNatRules group.
// Don't use this type directly, use NewVirtualNetworkGatewayNatRulesClient() instead.
type VirtualNetworkGatewayNatRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworkGatewayNatRulesClient creates a new instance of VirtualNetworkGatewayNatRulesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworkGatewayNatRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkGatewayNatRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualNetworkGatewayNatRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworkGatewayNatRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a nat rule to a scalable virtual network gateway if it doesn't exist else updates the existing
// nat rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The resource group name of the Virtual Network Gateway.
//   - virtualNetworkGatewayName - The name of the gateway.
//   - natRuleName - The name of the nat rule.
//   - natRuleParameters - Parameters supplied to create or Update a Nat Rule.
//   - options - VirtualNetworkGatewayNatRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkGatewayNatRulesClient.BeginCreateOrUpdate
//     method.
func (client *VirtualNetworkGatewayNatRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule, options *VirtualNetworkGatewayNatRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworkGatewayNatRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, natRuleParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworkGatewayNatRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkGatewayNatRulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a nat rule to a scalable virtual network gateway if it doesn't exist else updates the existing
// nat rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *VirtualNetworkGatewayNatRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule, options *VirtualNetworkGatewayNatRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualNetworkGatewayNatRulesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, natRuleParameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkGatewayNatRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters VirtualNetworkGatewayNatRule, options *VirtualNetworkGatewayNatRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, natRuleParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The resource group name of the Virtual Network Gateway.
//   - virtualNetworkGatewayName - The name of the gateway.
//   - natRuleName - The name of the nat rule.
//   - options - VirtualNetworkGatewayNatRulesClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkGatewayNatRulesClient.BeginDelete
//     method.
func (client *VirtualNetworkGatewayNatRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworkGatewayNatRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworkGatewayNatRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkGatewayNatRulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *VirtualNetworkGatewayNatRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualNetworkGatewayNatRulesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, options)
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
func (client *VirtualNetworkGatewayNatRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The resource group name of the Virtual Network Gateway.
//   - virtualNetworkGatewayName - The name of the gateway.
//   - natRuleName - The name of the nat rule.
//   - options - VirtualNetworkGatewayNatRulesClientGetOptions contains the optional parameters for the VirtualNetworkGatewayNatRulesClient.Get
//     method.
func (client *VirtualNetworkGatewayNatRulesClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesClientGetOptions) (VirtualNetworkGatewayNatRulesClientGetResponse, error) {
	var err error
	const operationName = "VirtualNetworkGatewayNatRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, natRuleName, options)
	if err != nil {
		return VirtualNetworkGatewayNatRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkGatewayNatRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualNetworkGatewayNatRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkGatewayNatRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *VirtualNetworkGatewayNatRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkGatewayNatRulesClient) getHandleResponse(resp *http.Response) (VirtualNetworkGatewayNatRulesClientGetResponse, error) {
	result := VirtualNetworkGatewayNatRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkGatewayNatRule); err != nil {
		return VirtualNetworkGatewayNatRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVirtualNetworkGatewayPager - Retrieves all nat rules for a particular virtual network gateway.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The resource group name of the virtual network gateway.
//   - virtualNetworkGatewayName - The name of the gateway.
//   - options - VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayOptions contains the optional parameters for the
//     VirtualNetworkGatewayNatRulesClient.NewListByVirtualNetworkGatewayPager method.
func (client *VirtualNetworkGatewayNatRulesClient) NewListByVirtualNetworkGatewayPager(resourceGroupName string, virtualNetworkGatewayName string, options *VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayOptions) *runtime.Pager[VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse]{
		More: func(page VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse) (VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualNetworkGatewayNatRulesClient.NewListByVirtualNetworkGatewayPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVirtualNetworkGatewayCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVirtualNetworkGatewayHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVirtualNetworkGatewayCreateRequest creates the ListByVirtualNetworkGateway request.
func (client *VirtualNetworkGatewayNatRulesClient) listByVirtualNetworkGatewayCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, options *VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/natRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkGatewayName == "" {
		return nil, errors.New("parameter virtualNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayName}", url.PathEscape(virtualNetworkGatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVirtualNetworkGatewayHandleResponse handles the ListByVirtualNetworkGateway response.
func (client *VirtualNetworkGatewayNatRulesClient) listByVirtualNetworkGatewayHandleResponse(resp *http.Response) (VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse, error) {
	result := VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualNetworkGatewayNatRulesResult); err != nil {
		return VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse{}, err
	}
	return result, nil
}
