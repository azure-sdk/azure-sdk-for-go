//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// ContainerAppsClient contains the methods for the ContainerApps group.
// Don't use this type directly, use NewContainerAppsClient() instead.
type ContainerAppsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContainerAppsClient creates a new instance of ContainerAppsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContainerAppsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsClient, error) {
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
	client := &ContainerAppsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerAppName - Name of the Container App.
// containerAppEnvelope - Properties used to create a container app
// options - ContainerAppsClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainerAppsClient.BeginCreateOrUpdate
// method.
func (client *ContainerAppsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ContainerAppsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, containerAppName, containerAppEnvelope, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainerAppsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainerAppsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
func (client *ContainerAppsClient) createOrUpdate(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerAppName, containerAppEnvelope, options)
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
func (client *ContainerAppsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, containerAppEnvelope)
}

// BeginDelete - Delete a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerAppName - Name of the Container App.
// options - ContainerAppsClientBeginDeleteOptions contains the optional parameters for the ContainerAppsClient.BeginDelete
// method.
func (client *ContainerAppsClient) BeginDelete(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientBeginDeleteOptions) (*runtime.Poller[ContainerAppsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, containerAppName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainerAppsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainerAppsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
func (client *ContainerAppsClient) deleteOperation(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerAppName, options)
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
func (client *ContainerAppsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerAppName - Name of the Container App.
// options - ContainerAppsClientGetOptions contains the optional parameters for the ContainerAppsClient.Get method.
func (client *ContainerAppsClient) Get(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientGetOptions) (ContainerAppsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerAppName, options)
	if err != nil {
		return ContainerAppsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerAppsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerAppsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContainerAppsClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerAppsClient) getHandleResponse(resp *http.Response) (ContainerAppsClientGetResponse, error) {
	result := ContainerAppsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerApp); err != nil {
		return ContainerAppsClientGetResponse{}, err
	}
	return result, nil
}

// GetAuthToken - Get auth token for a container app
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerAppName - Name of the Container App.
// options - ContainerAppsClientGetAuthTokenOptions contains the optional parameters for the ContainerAppsClient.GetAuthToken
// method.
func (client *ContainerAppsClient) GetAuthToken(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientGetAuthTokenOptions) (ContainerAppsClientGetAuthTokenResponse, error) {
	req, err := client.getAuthTokenCreateRequest(ctx, resourceGroupName, containerAppName, options)
	if err != nil {
		return ContainerAppsClientGetAuthTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerAppsClientGetAuthTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerAppsClientGetAuthTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAuthTokenHandleResponse(resp)
}

// getAuthTokenCreateRequest creates the GetAuthToken request.
func (client *ContainerAppsClient) getAuthTokenCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientGetAuthTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/getAuthtoken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAuthTokenHandleResponse handles the GetAuthToken response.
func (client *ContainerAppsClient) getAuthTokenHandleResponse(resp *http.Response) (ContainerAppsClientGetAuthTokenResponse, error) {
	result := ContainerAppsClientGetAuthTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppAuthToken); err != nil {
		return ContainerAppsClientGetAuthTokenResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get the Container Apps in a given resource group.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ContainerAppsClientListByResourceGroupOptions contains the optional parameters for the ContainerAppsClient.ListByResourceGroup
// method.
func (client *ContainerAppsClient) NewListByResourceGroupPager(resourceGroupName string, options *ContainerAppsClientListByResourceGroupOptions) *runtime.Pager[ContainerAppsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsClientListByResourceGroupResponse]{
		More: func(page ContainerAppsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsClientListByResourceGroupResponse) (ContainerAppsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerAppsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContainerAppsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerAppsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ContainerAppsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ContainerAppsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ContainerAppsClient) listByResourceGroupHandleResponse(resp *http.Response) (ContainerAppsClientListByResourceGroupResponse, error) {
	result := ContainerAppsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppCollection); err != nil {
		return ContainerAppsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get the Container Apps in a given subscription.
// Generated from API version 2022-10-01
// options - ContainerAppsClientListBySubscriptionOptions contains the optional parameters for the ContainerAppsClient.ListBySubscription
// method.
func (client *ContainerAppsClient) NewListBySubscriptionPager(options *ContainerAppsClientListBySubscriptionOptions) *runtime.Pager[ContainerAppsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsClientListBySubscriptionResponse]{
		More: func(page ContainerAppsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsClientListBySubscriptionResponse) (ContainerAppsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerAppsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContainerAppsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerAppsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ContainerAppsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ContainerAppsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.App/containerApps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ContainerAppsClient) listBySubscriptionHandleResponse(resp *http.Response) (ContainerAppsClientListBySubscriptionResponse, error) {
	result := ContainerAppsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppCollection); err != nil {
		return ContainerAppsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListCustomHostNameAnalysis - Analyzes a custom hostname for a Container App
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerAppName - Name of the Container App.
// options - ContainerAppsClientListCustomHostNameAnalysisOptions contains the optional parameters for the ContainerAppsClient.ListCustomHostNameAnalysis
// method.
func (client *ContainerAppsClient) ListCustomHostNameAnalysis(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientListCustomHostNameAnalysisOptions) (ContainerAppsClientListCustomHostNameAnalysisResponse, error) {
	req, err := client.listCustomHostNameAnalysisCreateRequest(ctx, resourceGroupName, containerAppName, options)
	if err != nil {
		return ContainerAppsClientListCustomHostNameAnalysisResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerAppsClientListCustomHostNameAnalysisResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerAppsClientListCustomHostNameAnalysisResponse{}, runtime.NewResponseError(resp)
	}
	return client.listCustomHostNameAnalysisHandleResponse(resp)
}

// listCustomHostNameAnalysisCreateRequest creates the ListCustomHostNameAnalysis request.
func (client *ContainerAppsClient) listCustomHostNameAnalysisCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientListCustomHostNameAnalysisOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/listCustomHostNameAnalysis"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CustomHostname != nil {
		reqQP.Set("customHostname", *options.CustomHostname)
	}
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listCustomHostNameAnalysisHandleResponse handles the ListCustomHostNameAnalysis response.
func (client *ContainerAppsClient) listCustomHostNameAnalysisHandleResponse(resp *http.Response) (ContainerAppsClientListCustomHostNameAnalysisResponse, error) {
	result := ContainerAppsClientListCustomHostNameAnalysisResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomHostnameAnalysisResult); err != nil {
		return ContainerAppsClientListCustomHostNameAnalysisResponse{}, err
	}
	return result, nil
}

// ListSecrets - List secrets for a container app
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerAppName - Name of the Container App.
// options - ContainerAppsClientListSecretsOptions contains the optional parameters for the ContainerAppsClient.ListSecrets
// method.
func (client *ContainerAppsClient) ListSecrets(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientListSecretsOptions) (ContainerAppsClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, containerAppName, options)
	if err != nil {
		return ContainerAppsClientListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerAppsClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerAppsClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *ContainerAppsClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/listSecrets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *ContainerAppsClient) listSecretsHandleResponse(resp *http.Response) (ContainerAppsClientListSecretsResponse, error) {
	result := ContainerAppsClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretsCollection); err != nil {
		return ContainerAppsClientListSecretsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patches a Container App using JSON Merge Patch
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerAppName - Name of the Container App.
// containerAppEnvelope - Properties of a Container App that need to be updated
// options - ContainerAppsClientBeginUpdateOptions contains the optional parameters for the ContainerAppsClient.BeginUpdate
// method.
func (client *ContainerAppsClient) BeginUpdate(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginUpdateOptions) (*runtime.Poller[ContainerAppsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, containerAppName, containerAppEnvelope, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainerAppsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainerAppsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patches a Container App using JSON Merge Patch
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
func (client *ContainerAppsClient) update(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, containerAppName, containerAppEnvelope, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ContainerAppsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, containerAppEnvelope)
}
