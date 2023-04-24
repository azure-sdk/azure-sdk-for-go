//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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
	"strconv"
	"strings"
)

// OnlineEndpointsClient contains the methods for the OnlineEndpoints group.
// Don't use this type directly, use NewOnlineEndpointsClient() instead.
type OnlineEndpointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOnlineEndpointsClient creates a new instance of OnlineEndpointsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOnlineEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OnlineEndpointsClient, error) {
	cl, err := arm.NewClient(moduleName+".OnlineEndpointsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OnlineEndpointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update Online Endpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - body - Online Endpoint entity to apply during operation.
//   - options - OnlineEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the OnlineEndpointsClient.BeginCreateOrUpdate
//     method.
func (client *OnlineEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body OnlineEndpoint, options *OnlineEndpointsClientBeginCreateOrUpdateOptions) (*runtime.Poller[OnlineEndpointsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, endpointName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OnlineEndpointsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
	} else {
		return runtime.NewPollerFromResumeToken[OnlineEndpointsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update Online Endpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *OnlineEndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body OnlineEndpoint, options *OnlineEndpointsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, body, options)
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
func (client *OnlineEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body OnlineEndpoint, options *OnlineEndpointsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete Online Endpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - options - OnlineEndpointsClientBeginDeleteOptions contains the optional parameters for the OnlineEndpointsClient.BeginDelete
//     method.
func (client *OnlineEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientBeginDeleteOptions) (*runtime.Poller[OnlineEndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, endpointName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OnlineEndpointsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[OnlineEndpointsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete Online Endpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *OnlineEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, options)
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
func (client *OnlineEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
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

// Get - Get Online Endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - options - OnlineEndpointsClientGetOptions contains the optional parameters for the OnlineEndpointsClient.Get method.
func (client *OnlineEndpointsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientGetOptions) (OnlineEndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, options)
	if err != nil {
		return OnlineEndpointsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OnlineEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnlineEndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OnlineEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
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
func (client *OnlineEndpointsClient) getHandleResponse(resp *http.Response) (OnlineEndpointsClientGetResponse, error) {
	result := OnlineEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OnlineEndpoint); err != nil {
		return OnlineEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// GetToken - Retrieve a valid AML token for an Endpoint using AMLToken-based authentication.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - options - OnlineEndpointsClientGetTokenOptions contains the optional parameters for the OnlineEndpointsClient.GetToken
//     method.
func (client *OnlineEndpointsClient) GetToken(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientGetTokenOptions) (OnlineEndpointsClientGetTokenResponse, error) {
	req, err := client.getTokenCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, options)
	if err != nil {
		return OnlineEndpointsClientGetTokenResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OnlineEndpointsClientGetTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnlineEndpointsClientGetTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.getTokenHandleResponse(resp)
}

// getTokenCreateRequest creates the GetToken request.
func (client *OnlineEndpointsClient) getTokenCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientGetTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/token"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTokenHandleResponse handles the GetToken response.
func (client *OnlineEndpointsClient) getTokenHandleResponse(resp *http.Response) (OnlineEndpointsClientGetTokenResponse, error) {
	result := OnlineEndpointsClientGetTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointAuthToken); err != nil {
		return OnlineEndpointsClientGetTokenResponse{}, err
	}
	return result, nil
}

// NewListPager - List Online Endpoints.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - OnlineEndpointsClientListOptions contains the optional parameters for the OnlineEndpointsClient.NewListPager
//     method.
func (client *OnlineEndpointsClient) NewListPager(resourceGroupName string, workspaceName string, options *OnlineEndpointsClientListOptions) *runtime.Pager[OnlineEndpointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OnlineEndpointsClientListResponse]{
		More: func(page OnlineEndpointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OnlineEndpointsClientListResponse) (OnlineEndpointsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OnlineEndpointsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OnlineEndpointsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OnlineEndpointsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OnlineEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *OnlineEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints"
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
	reqQP.Set("api-version", "2023-04-01")
	if options != nil && options.Name != nil {
		reqQP.Set("name", *options.Name)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatInt(int64(*options.Count), 10))
	}
	if options != nil && options.ComputeType != nil {
		reqQP.Set("computeType", string(*options.ComputeType))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	if options != nil && options.Properties != nil {
		reqQP.Set("properties", *options.Properties)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", string(*options.OrderBy))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OnlineEndpointsClient) listHandleResponse(resp *http.Response) (OnlineEndpointsClientListResponse, error) {
	result := OnlineEndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OnlineEndpointTrackedResourceArmPaginatedResult); err != nil {
		return OnlineEndpointsClientListResponse{}, err
	}
	return result, nil
}

// ListKeys - List EndpointAuthKeys for an Endpoint using Key-based authentication.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - options - OnlineEndpointsClientListKeysOptions contains the optional parameters for the OnlineEndpointsClient.ListKeys
//     method.
func (client *OnlineEndpointsClient) ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientListKeysOptions) (OnlineEndpointsClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, options)
	if err != nil {
		return OnlineEndpointsClientListKeysResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OnlineEndpointsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnlineEndpointsClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *OnlineEndpointsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineEndpointsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/listKeys"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *OnlineEndpointsClient) listKeysHandleResponse(resp *http.Response) (OnlineEndpointsClientListKeysResponse, error) {
	result := OnlineEndpointsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointAuthKeys); err != nil {
		return OnlineEndpointsClientListKeysResponse{}, err
	}
	return result, nil
}

// BeginRegenerateKeys - Regenerate EndpointAuthKeys for an Endpoint using Key-based authentication (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - body - RegenerateKeys request .
//   - options - OnlineEndpointsClientBeginRegenerateKeysOptions contains the optional parameters for the OnlineEndpointsClient.BeginRegenerateKeys
//     method.
func (client *OnlineEndpointsClient) BeginRegenerateKeys(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body RegenerateEndpointKeysRequest, options *OnlineEndpointsClientBeginRegenerateKeysOptions) (*runtime.Poller[OnlineEndpointsClientRegenerateKeysResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.regenerateKeys(ctx, resourceGroupName, workspaceName, endpointName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OnlineEndpointsClientRegenerateKeysResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[OnlineEndpointsClientRegenerateKeysResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RegenerateKeys - Regenerate EndpointAuthKeys for an Endpoint using Key-based authentication (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *OnlineEndpointsClient) regenerateKeys(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body RegenerateEndpointKeysRequest, options *OnlineEndpointsClientBeginRegenerateKeysOptions) (*http.Response, error) {
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *OnlineEndpointsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body RegenerateEndpointKeysRequest, options *OnlineEndpointsClientBeginRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/regenerateKeys"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginUpdate - Update Online Endpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - body - Online Endpoint entity to apply during operation.
//   - options - OnlineEndpointsClientBeginUpdateOptions contains the optional parameters for the OnlineEndpointsClient.BeginUpdate
//     method.
func (client *OnlineEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body PartialMinimalTrackedResourceWithIdentity, options *OnlineEndpointsClientBeginUpdateOptions) (*runtime.Poller[OnlineEndpointsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, endpointName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OnlineEndpointsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OnlineEndpointsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update Online Endpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *OnlineEndpointsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body PartialMinimalTrackedResourceWithIdentity, options *OnlineEndpointsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OnlineEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body PartialMinimalTrackedResourceWithIdentity, options *OnlineEndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
