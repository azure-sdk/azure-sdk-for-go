//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

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
	"strconv"
	"strings"
)

// DscNodeClient contains the methods for the DscNode group.
// Don't use this type directly, use NewDscNodeClient() instead.
type DscNodeClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDscNodeClient creates a new instance of DscNodeClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDscNodeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscNodeClient, error) {
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
	client := &DscNodeClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Delete - Delete the dsc node identified by node id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeID - The node id.
//   - options - DscNodeClientDeleteOptions contains the optional parameters for the DscNodeClient.Delete method.
func (client *DscNodeClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeClientDeleteOptions) (DscNodeClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, options)
	if err != nil {
		return DscNodeClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscNodeClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscNodeClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DscNodeClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DscNodeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the dsc node identified by node id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeID - The node id.
//   - options - DscNodeClientGetOptions contains the optional parameters for the DscNodeClient.Get method.
func (client *DscNodeClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeClientGetOptions) (DscNodeClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, options)
	if err != nil {
		return DscNodeClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscNodeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscNodeClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DscNodeClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscNodeClient) getHandleResponse(resp *http.Response) (DscNodeClientGetResponse, error) {
	result := DscNodeClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNode); err != nil {
		return DscNodeClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of dsc nodes.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - DscNodeClientListByAutomationAccountOptions contains the optional parameters for the DscNodeClient.NewListByAutomationAccountPager
//     method.
func (client *DscNodeClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *DscNodeClientListByAutomationAccountOptions) *runtime.Pager[DscNodeClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscNodeClientListByAutomationAccountResponse]{
		More: func(page DscNodeClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscNodeClientListByAutomationAccountResponse) (DscNodeClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DscNodeClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DscNodeClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DscNodeClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscNodeClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscNodeClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Inlinecount != nil {
		reqQP.Set("$inlinecount", *options.Inlinecount)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscNodeClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscNodeClientListByAutomationAccountResponse, error) {
	result := DscNodeClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNodeListResult); err != nil {
		return DscNodeClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update the dsc node.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeID - Parameters supplied to the update dsc node.
//   - dscNodeUpdateParameters - Parameters supplied to the update dsc node.
//   - options - DscNodeClientUpdateOptions contains the optional parameters for the DscNodeClient.Update method.
func (client *DscNodeClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, dscNodeUpdateParameters DscNodeUpdateParameters, options *DscNodeClientUpdateOptions) (DscNodeClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, dscNodeUpdateParameters, options)
	if err != nil {
		return DscNodeClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscNodeClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscNodeClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DscNodeClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, dscNodeUpdateParameters DscNodeUpdateParameters, options *DscNodeClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dscNodeUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *DscNodeClient) updateHandleResponse(resp *http.Response) (DscNodeClientUpdateResponse, error) {
	result := DscNodeClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNode); err != nil {
		return DscNodeClientUpdateResponse{}, err
	}
	return result, nil
}
