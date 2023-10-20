//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

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

// ManagedNetworkSettingsRuleClient contains the methods for the ManagedNetworkSettingsRule group.
// Don't use this type directly, use NewManagedNetworkSettingsRuleClient() instead.
type ManagedNetworkSettingsRuleClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedNetworkSettingsRuleClient creates a new instance of ManagedNetworkSettingsRuleClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedNetworkSettingsRuleClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedNetworkSettingsRuleClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedNetworkSettingsRuleClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedNetworkSettingsRuleClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an outbound rule in the managed network of a machine learning workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - ruleName - Name of the workspace managed network outbound rule
//   - body - Outbound Rule to be created or updated in the managed network of a machine learning workspace.
//   - options - ManagedNetworkSettingsRuleClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedNetworkSettingsRuleClient.BeginCreateOrUpdate
//     method.
func (client *ManagedNetworkSettingsRuleClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, body OutboundRuleBasicResource, options *ManagedNetworkSettingsRuleClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedNetworkSettingsRuleClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, ruleName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedNetworkSettingsRuleClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagedNetworkSettingsRuleClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an outbound rule in the managed network of a machine learning workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *ManagedNetworkSettingsRuleClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, body OutboundRuleBasicResource, options *ManagedNetworkSettingsRuleClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, ruleName, body, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedNetworkSettingsRuleClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, body OutboundRuleBasicResource, options *ManagedNetworkSettingsRuleClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/outboundRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an outbound rule from the managed network of a machine learning workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - ruleName - Name of the workspace managed network outbound rule
//   - options - ManagedNetworkSettingsRuleClientBeginDeleteOptions contains the optional parameters for the ManagedNetworkSettingsRuleClient.BeginDelete
//     method.
func (client *ManagedNetworkSettingsRuleClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, options *ManagedNetworkSettingsRuleClientBeginDeleteOptions) (*runtime.Poller[ManagedNetworkSettingsRuleClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, ruleName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedNetworkSettingsRuleClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagedNetworkSettingsRuleClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an outbound rule from the managed network of a machine learning workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *ManagedNetworkSettingsRuleClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, options *ManagedNetworkSettingsRuleClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, ruleName, options)
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
func (client *ManagedNetworkSettingsRuleClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, options *ManagedNetworkSettingsRuleClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/outboundRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an outbound rule from the managed network of a machine learning workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - ruleName - Name of the workspace managed network outbound rule
//   - options - ManagedNetworkSettingsRuleClientGetOptions contains the optional parameters for the ManagedNetworkSettingsRuleClient.Get
//     method.
func (client *ManagedNetworkSettingsRuleClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, options *ManagedNetworkSettingsRuleClientGetOptions) (ManagedNetworkSettingsRuleClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, ruleName, options)
	if err != nil {
		return ManagedNetworkSettingsRuleClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedNetworkSettingsRuleClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedNetworkSettingsRuleClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedNetworkSettingsRuleClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, options *ManagedNetworkSettingsRuleClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/outboundRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedNetworkSettingsRuleClient) getHandleResponse(resp *http.Response) (ManagedNetworkSettingsRuleClientGetResponse, error) {
	result := ManagedNetworkSettingsRuleClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundRuleBasicResource); err != nil {
		return ManagedNetworkSettingsRuleClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the managed network outbound rules for a machine learning workspace.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - options - ManagedNetworkSettingsRuleClientListOptions contains the optional parameters for the ManagedNetworkSettingsRuleClient.NewListPager
//     method.
func (client *ManagedNetworkSettingsRuleClient) NewListPager(resourceGroupName string, workspaceName string, options *ManagedNetworkSettingsRuleClientListOptions) *runtime.Pager[ManagedNetworkSettingsRuleClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedNetworkSettingsRuleClientListResponse]{
		More: func(page ManagedNetworkSettingsRuleClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedNetworkSettingsRuleClientListResponse) (ManagedNetworkSettingsRuleClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedNetworkSettingsRuleClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagedNetworkSettingsRuleClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedNetworkSettingsRuleClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ManagedNetworkSettingsRuleClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ManagedNetworkSettingsRuleClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/outboundRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedNetworkSettingsRuleClient) listHandleResponse(resp *http.Response) (ManagedNetworkSettingsRuleClientListResponse, error) {
	result := ManagedNetworkSettingsRuleClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundRuleListResult); err != nil {
		return ManagedNetworkSettingsRuleClientListResponse{}, err
	}
	return result, nil
}
