//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ConfigurationPolicyGroupsClient contains the methods for the ConfigurationPolicyGroups group.
// Don't use this type directly, use NewConfigurationPolicyGroupsClient() instead.
type ConfigurationPolicyGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigurationPolicyGroupsClient creates a new instance of ConfigurationPolicyGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigurationPolicyGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationPolicyGroupsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationPolicyGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a ConfigurationPolicyGroup if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the ConfigurationPolicyGroup.
// vpnServerConfigurationName - The name of the VpnServerConfiguration.
// configurationPolicyGroupName - The name of the ConfigurationPolicyGroup.
// vpnServerConfigurationPolicyGroupParameters - Parameters supplied to create or update a VpnServerConfiguration PolicyGroup.
// options - ConfigurationPolicyGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the ConfigurationPolicyGroupsClient.BeginCreateOrUpdate
// method.
func (client *ConfigurationPolicyGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, vpnServerConfigurationPolicyGroupParameters VPNServerConfigurationPolicyGroup, options *ConfigurationPolicyGroupsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ConfigurationPolicyGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName, vpnServerConfigurationPolicyGroupParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ConfigurationPolicyGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ConfigurationPolicyGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a ConfigurationPolicyGroup if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConfigurationPolicyGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, vpnServerConfigurationPolicyGroupParameters VPNServerConfigurationPolicyGroup, options *ConfigurationPolicyGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName, vpnServerConfigurationPolicyGroupParameters, options)
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
func (client *ConfigurationPolicyGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, vpnServerConfigurationPolicyGroupParameters VPNServerConfigurationPolicyGroup, options *ConfigurationPolicyGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups/{configurationPolicyGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	if configurationPolicyGroupName == "" {
		return nil, errors.New("parameter configurationPolicyGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationPolicyGroupName}", url.PathEscape(configurationPolicyGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnServerConfigurationPolicyGroupParameters)
}

// BeginDelete - Deletes a ConfigurationPolicyGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the ConfigurationPolicyGroup.
// vpnServerConfigurationName - The name of the VpnServerConfiguration.
// configurationPolicyGroupName - The name of the ConfigurationPolicyGroup.
// options - ConfigurationPolicyGroupsClientBeginDeleteOptions contains the optional parameters for the ConfigurationPolicyGroupsClient.BeginDelete
// method.
func (client *ConfigurationPolicyGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, options *ConfigurationPolicyGroupsClientBeginDeleteOptions) (*armruntime.Poller[ConfigurationPolicyGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ConfigurationPolicyGroupsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ConfigurationPolicyGroupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a ConfigurationPolicyGroup.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConfigurationPolicyGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, options *ConfigurationPolicyGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName, options)
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
func (client *ConfigurationPolicyGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, options *ConfigurationPolicyGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups/{configurationPolicyGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	if configurationPolicyGroupName == "" {
		return nil, errors.New("parameter configurationPolicyGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationPolicyGroupName}", url.PathEscape(configurationPolicyGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the details of a ConfigurationPolicyGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration.
// configurationPolicyGroupName - The name of the ConfigurationPolicyGroup being retrieved.
// options - ConfigurationPolicyGroupsClientGetOptions contains the optional parameters for the ConfigurationPolicyGroupsClient.Get
// method.
func (client *ConfigurationPolicyGroupsClient) Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, options *ConfigurationPolicyGroupsClientGetOptions) (ConfigurationPolicyGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName, options)
	if err != nil {
		return ConfigurationPolicyGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationPolicyGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationPolicyGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationPolicyGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, options *ConfigurationPolicyGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups/{configurationPolicyGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	if configurationPolicyGroupName == "" {
		return nil, errors.New("parameter configurationPolicyGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationPolicyGroupName}", url.PathEscape(configurationPolicyGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationPolicyGroupsClient) getHandleResponse(resp *http.Response) (ConfigurationPolicyGroupsClientGetResponse, error) {
	result := ConfigurationPolicyGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNServerConfigurationPolicyGroup); err != nil {
		return ConfigurationPolicyGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVPNServerConfigurationPager - Lists all the configurationPolicyGroups in a resource group for a vpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration.
// options - ConfigurationPolicyGroupsClientListByVPNServerConfigurationOptions contains the optional parameters for the ConfigurationPolicyGroupsClient.ListByVPNServerConfiguration
// method.
func (client *ConfigurationPolicyGroupsClient) NewListByVPNServerConfigurationPager(resourceGroupName string, vpnServerConfigurationName string, options *ConfigurationPolicyGroupsClientListByVPNServerConfigurationOptions) *runtime.Pager[ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse] {
	return runtime.NewPager(runtime.PageProcessor[ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse]{
		More: func(page ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse) (ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVPNServerConfigurationCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVPNServerConfigurationHandleResponse(resp)
		},
	})
}

// listByVPNServerConfigurationCreateRequest creates the ListByVPNServerConfiguration request.
func (client *ConfigurationPolicyGroupsClient) listByVPNServerConfigurationCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *ConfigurationPolicyGroupsClientListByVPNServerConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVPNServerConfigurationHandleResponse handles the ListByVPNServerConfiguration response.
func (client *ConfigurationPolicyGroupsClient) listByVPNServerConfigurationHandleResponse(resp *http.Response) (ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse, error) {
	result := ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNServerConfigurationPolicyGroupsResult); err != nil {
		return ConfigurationPolicyGroupsClientListByVPNServerConfigurationResponse{}, err
	}
	return result, nil
}
