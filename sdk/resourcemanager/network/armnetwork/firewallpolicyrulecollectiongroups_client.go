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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FirewallPolicyRuleCollectionGroupsClient contains the methods for the FirewallPolicyRuleCollectionGroups group.
// Don't use this type directly, use NewFirewallPolicyRuleCollectionGroupsClient() instead.
type FirewallPolicyRuleCollectionGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFirewallPolicyRuleCollectionGroupsClient creates a new instance of FirewallPolicyRuleCollectionGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFirewallPolicyRuleCollectionGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallPolicyRuleCollectionGroupsClient, error) {
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
	client := &FirewallPolicyRuleCollectionGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
// parameters - Parameters supplied to the create or update FirewallPolicyRuleCollectionGroup operation.
// options - FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdate
// method.
func (client *FirewallPolicyRuleCollectionGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
// options - FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.BeginDelete
// method.
func (client *FirewallPolicyRuleCollectionGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions) (*runtime.Poller[FirewallPolicyRuleCollectionGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FirewallPolicyRuleCollectionGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FirewallPolicyRuleCollectionGroupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// ruleCollectionGroupName - The name of the FirewallPolicyRuleCollectionGroup.
// options - FirewallPolicyRuleCollectionGroupsClientGetOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.Get
// method.
func (client *FirewallPolicyRuleCollectionGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsClientGetOptions) (FirewallPolicyRuleCollectionGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallPolicyRuleCollectionGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
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

// NewListPager - Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// options - FirewallPolicyRuleCollectionGroupsClientListOptions contains the optional parameters for the FirewallPolicyRuleCollectionGroupsClient.List
// method.
func (client *FirewallPolicyRuleCollectionGroupsClient) NewListPager(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsClientListOptions) *runtime.Pager[FirewallPolicyRuleCollectionGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FirewallPolicyRuleCollectionGroupsClientListResponse]{
		More: func(page FirewallPolicyRuleCollectionGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallPolicyRuleCollectionGroupsClientListResponse) (FirewallPolicyRuleCollectionGroupsClientListResponse, error) {
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
			resp, err := client.pl.Do(req)
			if err != nil {
				return FirewallPolicyRuleCollectionGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FirewallPolicyRuleCollectionGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
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
