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

// FirewallPolicyRuleCollectionGroupsClient contains the methods for the FirewallPolicyRuleCollectionGroups group.
// Don't use this type directly, use NewFirewallPolicyRuleCollectionGroupsClient() instead.
type FirewallPolicyRuleCollectionGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFirewallPolicyRuleCollectionGroupsClient creates a new instance of FirewallPolicyRuleCollectionGroupsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFirewallPolicyRuleCollectionGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallPolicyRuleCollectionGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".FirewallPolicyRuleCollectionGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallPolicyRuleCollectionGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
//   - parameters - Parameters supplied to the create or update FirewallPolicyRuleCollectionGroup operation.
//   - options - FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdate
//     method.
func (client *FirewallPolicyRuleCollectionGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
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
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdateDraft - Creates or updates a new draft version of the specified Firewall Policy Rule Collection Group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
//   - parameters - Parameters supplied to the create or update FirewallPolicyRuleCollectionGroup operation.
//   - options - FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateDraftOptions contains the optional parameters for
//     the FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdateDraft method.
func (client *FirewallPolicyRuleCollectionGroupsClient) BeginCreateOrUpdateDraft(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroupDraft, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateDraftOptions) (*runtime.Poller[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateDraftResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateDraft(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateDraftResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateDraftResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateDraft - Creates or updates a new draft version of the specified Firewall Policy Rule Collection Group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdateDraft(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroupDraft, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateDraftOptions) (*http.Response, error) {
	var err error
	const operationName = "FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdateDraft"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateDraftCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateDraftCreateRequest creates the CreateOrUpdateDraft request.
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdateDraftCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroupDraft, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateDraftOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}/putDraft"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
//   - options - FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.BeginDelete
//     method.
func (client *FirewallPolicyRuleCollectionGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions) (*runtime.Poller[FirewallPolicyRuleCollectionGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FirewallPolicyRuleCollectionGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[FirewallPolicyRuleCollectionGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FirewallPolicyRuleCollectionGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
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
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteDraft - Get the current draft version of the specified Firewall Policy Rule Collection Group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
//   - parameters - Parameters supplied to the create or update FirewallPolicyRuleCollectionGroup operation.
//   - options - FirewallPolicyRuleCollectionGroupsClientDeleteDraftOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.DeleteDraft
//     method.
func (client *FirewallPolicyRuleCollectionGroupsClient) DeleteDraft(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientDeleteDraftOptions) (FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse, error) {
	var err error
	const operationName = "FirewallPolicyRuleCollectionGroupsClient.DeleteDraft"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteDraftCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse{}, err
	}
	resp, err := client.deleteDraftHandleResponse(httpResp)
	return resp, err
}

// deleteDraftCreateRequest creates the DeleteDraft request.
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteDraftCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientDeleteDraftOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}/deleteDraft"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// deleteDraftHandleResponse handles the DeleteDraft response.
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteDraftHandleResponse(resp *http.Response) (FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse, error) {
	result := FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyRuleCollectionGroup); err != nil {
		return FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse{}, err
	}
	return result, nil
}

// Get - Gets the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
//   - options - FirewallPolicyRuleCollectionGroupsClientGetOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.Get
//     method.
func (client *FirewallPolicyRuleCollectionGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientGetOptions) (FirewallPolicyRuleCollectionGroupsClientGetResponse, error) {
	var err error
	const operationName = "FirewallPolicyRuleCollectionGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallPolicyRuleCollectionGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FirewallPolicyRuleCollectionGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FirewallPolicyRuleCollectionGroupsClient) getHandleResponse(resp *http.Response) (FirewallPolicyRuleCollectionGroupsClientGetResponse, error) {
	result := FirewallPolicyRuleCollectionGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyRuleCollectionGroup); err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetResponse{}, err
	}
	return result, nil
}

// GetDraft - Get the current draft version of the specified Firewall Policy Rule Collection Group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
//   - parameters - Parameters supplied to the create or update FirewallPolicyRuleCollectionGroup operation.
//   - options - FirewallPolicyRuleCollectionGroupsClientGetDraftOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.GetDraft
//     method.
func (client *FirewallPolicyRuleCollectionGroupsClient) GetDraft(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientGetDraftOptions) (FirewallPolicyRuleCollectionGroupsClientGetDraftResponse, error) {
	var err error
	const operationName = "FirewallPolicyRuleCollectionGroupsClient.GetDraft"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDraftCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetDraftResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetDraftResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallPolicyRuleCollectionGroupsClientGetDraftResponse{}, err
	}
	resp, err := client.getDraftHandleResponse(httpResp)
	return resp, err
}

// getDraftCreateRequest creates the GetDraft request.
func (client *FirewallPolicyRuleCollectionGroupsClient) getDraftCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientGetDraftOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}/getDraft"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// getDraftHandleResponse handles the GetDraft response.
func (client *FirewallPolicyRuleCollectionGroupsClient) getDraftHandleResponse(resp *http.Response) (FirewallPolicyRuleCollectionGroupsClientGetDraftResponse, error) {
	result := FirewallPolicyRuleCollectionGroupsClientGetDraftResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyRuleCollectionGroupDraft); err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetDraftResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - options - FirewallPolicyRuleCollectionGroupsClientListOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.NewListPager
//     method.
func (client *FirewallPolicyRuleCollectionGroupsClient) NewListPager(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsClientListOptions) *runtime.Pager[FirewallPolicyRuleCollectionGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FirewallPolicyRuleCollectionGroupsClientListResponse]{
		More: func(page FirewallPolicyRuleCollectionGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallPolicyRuleCollectionGroupsClientListResponse) (FirewallPolicyRuleCollectionGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FirewallPolicyRuleCollectionGroupsClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FirewallPolicyRuleCollectionGroupsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FirewallPolicyRuleCollectionGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FirewallPolicyRuleCollectionGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *FirewallPolicyRuleCollectionGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FirewallPolicyRuleCollectionGroupsClient) listHandleResponse(resp *http.Response) (FirewallPolicyRuleCollectionGroupsClientListResponse, error) {
	result := FirewallPolicyRuleCollectionGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyRuleCollectionGroupListResult); err != nil {
		return FirewallPolicyRuleCollectionGroupsClientListResponse{}, err
	}
	return result, nil
}
