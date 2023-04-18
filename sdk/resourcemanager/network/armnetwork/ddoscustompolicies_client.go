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

// DdosCustomPoliciesClient contains the methods for the DdosCustomPolicies group.
// Don't use this type directly, use NewDdosCustomPoliciesClient() instead.
type DdosCustomPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDdosCustomPoliciesClient creates a new instance of DdosCustomPoliciesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDdosCustomPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DdosCustomPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".DdosCustomPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DdosCustomPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DDoS custom policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - ddosCustomPolicyName - The name of the DDoS custom policy.
//   - parameters - Parameters supplied to the create or update operation.
//   - options - DdosCustomPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the DdosCustomPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *DdosCustomPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy, options *DdosCustomPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DdosCustomPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, ddosCustomPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DdosCustomPoliciesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DdosCustomPoliciesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a DDoS custom policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *DdosCustomPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy, options *DdosCustomPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, parameters, options)
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
func (client *DdosCustomPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy, options *DdosCustomPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified DDoS custom policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - ddosCustomPolicyName - The name of the DDoS custom policy.
//   - options - DdosCustomPoliciesClientBeginDeleteOptions contains the optional parameters for the DdosCustomPoliciesClient.BeginDelete
//     method.
func (client *DdosCustomPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesClientBeginDeleteOptions) (*runtime.Poller[DdosCustomPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ddosCustomPolicyName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DdosCustomPoliciesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DdosCustomPoliciesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified DDoS custom policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *DdosCustomPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, options)
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
func (client *DdosCustomPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified DDoS custom policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - ddosCustomPolicyName - The name of the DDoS custom policy.
//   - options - DdosCustomPoliciesClientGetOptions contains the optional parameters for the DdosCustomPoliciesClient.Get method.
func (client *DdosCustomPoliciesClient) Get(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesClientGetOptions) (DdosCustomPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, options)
	if err != nil {
		return DdosCustomPoliciesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DdosCustomPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DdosCustomPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DdosCustomPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DdosCustomPoliciesClient) getHandleResponse(resp *http.Response) (DdosCustomPoliciesClientGetResponse, error) {
	result := DdosCustomPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosCustomPolicy); err != nil {
		return DdosCustomPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update a DDoS custom policy tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - ddosCustomPolicyName - The name of the DDoS custom policy.
//   - parameters - Parameters supplied to update DDoS custom policy resource tags.
//   - options - DdosCustomPoliciesClientUpdateTagsOptions contains the optional parameters for the DdosCustomPoliciesClient.UpdateTags
//     method.
func (client *DdosCustomPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject, options *DdosCustomPoliciesClientUpdateTagsOptions) (DdosCustomPoliciesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, parameters, options)
	if err != nil {
		return DdosCustomPoliciesClientUpdateTagsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DdosCustomPoliciesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DdosCustomPoliciesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DdosCustomPoliciesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject, options *DdosCustomPoliciesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DdosCustomPoliciesClient) updateTagsHandleResponse(resp *http.Response) (DdosCustomPoliciesClientUpdateTagsResponse, error) {
	result := DdosCustomPoliciesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosCustomPolicy); err != nil {
		return DdosCustomPoliciesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
