//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

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

// ConfigurationsClient contains the methods for the Configurations group.
// Don't use this type directly, use NewConfigurationsClient() instead.
type ConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateInResourceGroup - Create/Overwrite Azure Advisor configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - configurationName - Advisor configuration name. Value must be 'default'
//   - resourceGroup - The name of the Azure resource group.
//   - configContract - The Azure Advisor configuration data structure.
//   - options - ConfigurationsClientCreateInResourceGroupOptions contains the optional parameters for the ConfigurationsClient.CreateInResourceGroup
//     method.
func (client *ConfigurationsClient) CreateInResourceGroup(ctx context.Context, configurationName ConfigurationName, resourceGroup string, configContract ConfigData, options *ConfigurationsClientCreateInResourceGroupOptions) (ConfigurationsClientCreateInResourceGroupResponse, error) {
	var err error
	req, err := client.createInResourceGroupCreateRequest(ctx, configurationName, resourceGroup, configContract, options)
	if err != nil {
		return ConfigurationsClientCreateInResourceGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsClientCreateInResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationsClientCreateInResourceGroupResponse{}, err
	}
	resp, err := client.createInResourceGroupHandleResponse(httpResp)
	return resp, err
}

// createInResourceGroupCreateRequest creates the CreateInResourceGroup request.
func (client *ConfigurationsClient) createInResourceGroupCreateRequest(ctx context.Context, configurationName ConfigurationName, resourceGroup string, configContract ConfigData, options *ConfigurationsClientCreateInResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, configContract); err != nil {
		return nil, err
	}
	return req, nil
}

// createInResourceGroupHandleResponse handles the CreateInResourceGroup response.
func (client *ConfigurationsClient) createInResourceGroupHandleResponse(resp *http.Response) (ConfigurationsClientCreateInResourceGroupResponse, error) {
	result := ConfigurationsClientCreateInResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigData); err != nil {
		return ConfigurationsClientCreateInResourceGroupResponse{}, err
	}
	return result, nil
}

// CreateInSubscription - Create/Overwrite Azure Advisor configuration and also delete all configurations of contained resource
// groups.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - configurationName - Advisor configuration name. Value must be 'default'
//   - configContract - The Azure Advisor configuration data structure.
//   - options - ConfigurationsClientCreateInSubscriptionOptions contains the optional parameters for the ConfigurationsClient.CreateInSubscription
//     method.
func (client *ConfigurationsClient) CreateInSubscription(ctx context.Context, configurationName ConfigurationName, configContract ConfigData, options *ConfigurationsClientCreateInSubscriptionOptions) (ConfigurationsClientCreateInSubscriptionResponse, error) {
	var err error
	req, err := client.createInSubscriptionCreateRequest(ctx, configurationName, configContract, options)
	if err != nil {
		return ConfigurationsClientCreateInSubscriptionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationsClientCreateInSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationsClientCreateInSubscriptionResponse{}, err
	}
	resp, err := client.createInSubscriptionHandleResponse(httpResp)
	return resp, err
}

// createInSubscriptionCreateRequest creates the CreateInSubscription request.
func (client *ConfigurationsClient) createInSubscriptionCreateRequest(ctx context.Context, configurationName ConfigurationName, configContract ConfigData, options *ConfigurationsClientCreateInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, configContract); err != nil {
		return nil, err
	}
	return req, nil
}

// createInSubscriptionHandleResponse handles the CreateInSubscription response.
func (client *ConfigurationsClient) createInSubscriptionHandleResponse(resp *http.Response) (ConfigurationsClientCreateInSubscriptionResponse, error) {
	result := ConfigurationsClientCreateInSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigData); err != nil {
		return ConfigurationsClientCreateInSubscriptionResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve Azure Advisor configurations.
//
// Generated from API version 2023-01-01
//   - resourceGroup - The name of the Azure resource group.
//   - options - ConfigurationsClientListByResourceGroupOptions contains the optional parameters for the ConfigurationsClient.NewListByResourceGroupPager
//     method.
func (client *ConfigurationsClient) NewListByResourceGroupPager(resourceGroup string, options *ConfigurationsClientListByResourceGroupOptions) *runtime.Pager[ConfigurationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationsClientListByResourceGroupResponse]{
		More: func(page ConfigurationsClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsClientListByResourceGroupResponse) (ConfigurationsClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroup, options)
			if err != nil {
				return ConfigurationsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroup string, options *ConfigurationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationsClient) listByResourceGroupHandleResponse(resp *http.Response) (ConfigurationsClientListByResourceGroupResponse, error) {
	result := ConfigurationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationListResult); err != nil {
		return ConfigurationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieve Azure Advisor configurations and also retrieve configurations of contained resource
// groups.
//
// Generated from API version 2023-01-01
//   - options - ConfigurationsClientListBySubscriptionOptions contains the optional parameters for the ConfigurationsClient.NewListBySubscriptionPager
//     method.
func (client *ConfigurationsClient) NewListBySubscriptionPager(options *ConfigurationsClientListBySubscriptionOptions) *runtime.Pager[ConfigurationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationsClientListBySubscriptionResponse]{
		More: func(page ConfigurationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsClientListBySubscriptionResponse) (ConfigurationsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConfigurationsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConfigurationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConfigurationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConfigurationsClient) listBySubscriptionHandleResponse(resp *http.Response) (ConfigurationsClientListBySubscriptionResponse, error) {
	result := ConfigurationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationListResult); err != nil {
		return ConfigurationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
