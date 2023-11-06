//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// ConfigurationGroupValuesClient contains the methods for the ConfigurationGroupValues group.
// Don't use this type directly, use NewConfigurationGroupValuesClient() instead.
type ConfigurationGroupValuesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationGroupValuesClient creates a new instance of ConfigurationGroupValuesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationGroupValuesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationGroupValuesClient, error) {
	cl, err := arm.NewClient(moduleName+".ConfigurationGroupValuesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationGroupValuesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a hybrid configuration group value.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - configurationGroupValueName - The name of the configuration group value.
//   - parameters - Parameters supplied to the create or update configuration group value resource.
//   - options - ConfigurationGroupValuesClientBeginCreateOrUpdateOptions contains the optional parameters for the ConfigurationGroupValuesClient.BeginCreateOrUpdate
//     method.
func (client *ConfigurationGroupValuesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, configurationGroupValueName string, parameters ConfigurationGroupValue, options *ConfigurationGroupValuesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ConfigurationGroupValuesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, configurationGroupValueName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConfigurationGroupValuesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConfigurationGroupValuesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a hybrid configuration group value.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *ConfigurationGroupValuesClient) createOrUpdate(ctx context.Context, resourceGroupName string, configurationGroupValueName string, parameters ConfigurationGroupValue, options *ConfigurationGroupValuesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, configurationGroupValueName, parameters, options)
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
func (client *ConfigurationGroupValuesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, configurationGroupValueName string, parameters ConfigurationGroupValue, options *ConfigurationGroupValuesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/configurationGroupValues/{configurationGroupValueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationGroupValueName == "" {
		return nil, errors.New("parameter configurationGroupValueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationGroupValueName}", url.PathEscape(configurationGroupValueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified hybrid configuration group value.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - configurationGroupValueName - The name of the configuration group value.
//   - options - ConfigurationGroupValuesClientBeginDeleteOptions contains the optional parameters for the ConfigurationGroupValuesClient.BeginDelete
//     method.
func (client *ConfigurationGroupValuesClient) BeginDelete(ctx context.Context, resourceGroupName string, configurationGroupValueName string, options *ConfigurationGroupValuesClientBeginDeleteOptions) (*runtime.Poller[ConfigurationGroupValuesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, configurationGroupValueName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConfigurationGroupValuesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConfigurationGroupValuesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified hybrid configuration group value.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *ConfigurationGroupValuesClient) deleteOperation(ctx context.Context, resourceGroupName string, configurationGroupValueName string, options *ConfigurationGroupValuesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configurationGroupValueName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationGroupValuesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configurationGroupValueName string, options *ConfigurationGroupValuesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/configurationGroupValues/{configurationGroupValueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationGroupValueName == "" {
		return nil, errors.New("parameter configurationGroupValueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationGroupValueName}", url.PathEscape(configurationGroupValueName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified hybrid configuration group values.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - configurationGroupValueName - The name of the configuration group value.
//   - options - ConfigurationGroupValuesClientGetOptions contains the optional parameters for the ConfigurationGroupValuesClient.Get
//     method.
func (client *ConfigurationGroupValuesClient) Get(ctx context.Context, resourceGroupName string, configurationGroupValueName string, options *ConfigurationGroupValuesClientGetOptions) (ConfigurationGroupValuesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, configurationGroupValueName, options)
	if err != nil {
		return ConfigurationGroupValuesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationGroupValuesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationGroupValuesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationGroupValuesClient) getCreateRequest(ctx context.Context, resourceGroupName string, configurationGroupValueName string, options *ConfigurationGroupValuesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/configurationGroupValues/{configurationGroupValueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationGroupValueName == "" {
		return nil, errors.New("parameter configurationGroupValueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationGroupValueName}", url.PathEscape(configurationGroupValueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationGroupValuesClient) getHandleResponse(resp *http.Response) (ConfigurationGroupValuesClientGetResponse, error) {
	result := ConfigurationGroupValuesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationGroupValue); err != nil {
		return ConfigurationGroupValuesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the hybrid network configurationGroupValues in a resource group.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ConfigurationGroupValuesClientListByResourceGroupOptions contains the optional parameters for the ConfigurationGroupValuesClient.NewListByResourceGroupPager
//     method.
func (client *ConfigurationGroupValuesClient) NewListByResourceGroupPager(resourceGroupName string, options *ConfigurationGroupValuesClientListByResourceGroupOptions) *runtime.Pager[ConfigurationGroupValuesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationGroupValuesClientListByResourceGroupResponse]{
		More: func(page ConfigurationGroupValuesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationGroupValuesClientListByResourceGroupResponse) (ConfigurationGroupValuesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConfigurationGroupValuesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationGroupValuesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationGroupValuesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationGroupValuesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationGroupValuesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/configurationGroupValues"
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
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationGroupValuesClient) listByResourceGroupHandleResponse(resp *http.Response) (ConfigurationGroupValuesClientListByResourceGroupResponse, error) {
	result := ConfigurationGroupValuesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationGroupValueListResult); err != nil {
		return ConfigurationGroupValuesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all sites in the configuration group value in a subscription.
//
// Generated from API version 2023-09-01
//   - options - ConfigurationGroupValuesClientListBySubscriptionOptions contains the optional parameters for the ConfigurationGroupValuesClient.NewListBySubscriptionPager
//     method.
func (client *ConfigurationGroupValuesClient) NewListBySubscriptionPager(options *ConfigurationGroupValuesClientListBySubscriptionOptions) *runtime.Pager[ConfigurationGroupValuesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationGroupValuesClientListBySubscriptionResponse]{
		More: func(page ConfigurationGroupValuesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationGroupValuesClientListBySubscriptionResponse) (ConfigurationGroupValuesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConfigurationGroupValuesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationGroupValuesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationGroupValuesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConfigurationGroupValuesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConfigurationGroupValuesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/configurationGroupValues"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConfigurationGroupValuesClient) listBySubscriptionHandleResponse(resp *http.Response) (ConfigurationGroupValuesClientListBySubscriptionResponse, error) {
	result := ConfigurationGroupValuesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationGroupValueListResult); err != nil {
		return ConfigurationGroupValuesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a hybrid configuration group tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - configurationGroupValueName - The name of the configuration group value.
//   - parameters - Parameters supplied to update configuration group values tags.
//   - options - ConfigurationGroupValuesClientUpdateTagsOptions contains the optional parameters for the ConfigurationGroupValuesClient.UpdateTags
//     method.
func (client *ConfigurationGroupValuesClient) UpdateTags(ctx context.Context, resourceGroupName string, configurationGroupValueName string, parameters TagsObject, options *ConfigurationGroupValuesClientUpdateTagsOptions) (ConfigurationGroupValuesClientUpdateTagsResponse, error) {
	var err error
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, configurationGroupValueName, parameters, options)
	if err != nil {
		return ConfigurationGroupValuesClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationGroupValuesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationGroupValuesClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ConfigurationGroupValuesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, configurationGroupValueName string, parameters TagsObject, options *ConfigurationGroupValuesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/configurationGroupValues/{configurationGroupValueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationGroupValueName == "" {
		return nil, errors.New("parameter configurationGroupValueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationGroupValueName}", url.PathEscape(configurationGroupValueName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ConfigurationGroupValuesClient) updateTagsHandleResponse(resp *http.Response) (ConfigurationGroupValuesClientUpdateTagsResponse, error) {
	result := ConfigurationGroupValuesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationGroupValue); err != nil {
		return ConfigurationGroupValuesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
