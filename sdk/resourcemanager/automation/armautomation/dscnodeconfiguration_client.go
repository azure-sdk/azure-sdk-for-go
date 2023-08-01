//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DscNodeConfigurationClient contains the methods for the DscNodeConfiguration group.
// Don't use this type directly, use NewDscNodeConfigurationClient() instead.
type DscNodeConfigurationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDscNodeConfigurationClient creates a new instance of DscNodeConfigurationClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDscNodeConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscNodeConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName+".DscNodeConfigurationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DscNodeConfigurationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create the node configuration identified by node configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeConfigurationName - The Dsc node configuration name.
//   - parameters - The create or update parameters for configuration.
//   - options - DscNodeConfigurationClientBeginCreateOrUpdateOptions contains the optional parameters for the DscNodeConfigurationClient.BeginCreateOrUpdate
//     method.
func (client *DscNodeConfigurationClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, parameters DscNodeConfigurationCreateOrUpdateParameters, options *DscNodeConfigurationClientBeginCreateOrUpdateOptions) (*runtime.Poller[DscNodeConfigurationClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, automationAccountName, nodeConfigurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DscNodeConfigurationClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DscNodeConfigurationClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create the node configuration identified by node configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
func (client *DscNodeConfigurationClient) createOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, parameters DscNodeConfigurationCreateOrUpdateParameters, options *DscNodeConfigurationClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, nodeConfigurationName, parameters, options)
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
func (client *DscNodeConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, parameters DscNodeConfigurationCreateOrUpdateParameters, options *DscNodeConfigurationClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeConfigurationName == "" {
		return nil, errors.New("parameter nodeConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeConfigurationName}", url.PathEscape(nodeConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete the Dsc node configurations by node configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeConfigurationName - The Dsc node configuration name.
//   - options - DscNodeConfigurationClientDeleteOptions contains the optional parameters for the DscNodeConfigurationClient.Delete
//     method.
func (client *DscNodeConfigurationClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, options *DscNodeConfigurationClientDeleteOptions) (DscNodeConfigurationClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, nodeConfigurationName, options)
	if err != nil {
		return DscNodeConfigurationClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscNodeConfigurationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscNodeConfigurationClientDeleteResponse{}, err
	}
	return DscNodeConfigurationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DscNodeConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, options *DscNodeConfigurationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeConfigurationName == "" {
		return nil, errors.New("parameter nodeConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeConfigurationName}", url.PathEscape(nodeConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the Dsc node configurations by node configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeConfigurationName - The Dsc node configuration name.
//   - options - DscNodeConfigurationClientGetOptions contains the optional parameters for the DscNodeConfigurationClient.Get
//     method.
func (client *DscNodeConfigurationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, options *DscNodeConfigurationClientGetOptions) (DscNodeConfigurationClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, nodeConfigurationName, options)
	if err != nil {
		return DscNodeConfigurationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscNodeConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscNodeConfigurationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DscNodeConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, options *DscNodeConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeConfigurationName == "" {
		return nil, errors.New("parameter nodeConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeConfigurationName}", url.PathEscape(nodeConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscNodeConfigurationClient) getHandleResponse(resp *http.Response) (DscNodeConfigurationClientGetResponse, error) {
	result := DscNodeConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNodeConfiguration); err != nil {
		return DscNodeConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of dsc node configurations.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - DscNodeConfigurationClientListByAutomationAccountOptions contains the optional parameters for the DscNodeConfigurationClient.NewListByAutomationAccountPager
//     method.
func (client *DscNodeConfigurationClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *DscNodeConfigurationClientListByAutomationAccountOptions) *runtime.Pager[DscNodeConfigurationClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscNodeConfigurationClientListByAutomationAccountResponse]{
		More: func(page DscNodeConfigurationClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscNodeConfigurationClientListByAutomationAccountResponse) (DscNodeConfigurationClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DscNodeConfigurationClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DscNodeConfigurationClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DscNodeConfigurationClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscNodeConfigurationClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscNodeConfigurationClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscNodeConfigurationClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscNodeConfigurationClientListByAutomationAccountResponse, error) {
	result := DscNodeConfigurationClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNodeConfigurationListResult); err != nil {
		return DscNodeConfigurationClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}
