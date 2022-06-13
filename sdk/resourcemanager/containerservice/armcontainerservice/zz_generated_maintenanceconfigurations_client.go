//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// MaintenanceConfigurationsClient contains the methods for the MaintenanceConfigurations group.
// Don't use this type directly, use NewMaintenanceConfigurationsClient() instead.
type MaintenanceConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMaintenanceConfigurationsClient creates a new instance of MaintenanceConfigurationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMaintenanceConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MaintenanceConfigurationsClient, error) {
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
	client := &MaintenanceConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a maintenance configuration in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// configName - The name of the maintenance configuration.
// parameters - The maintenance configuration to create or update.
// options - MaintenanceConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the MaintenanceConfigurationsClient.CreateOrUpdate
// method.
func (client *MaintenanceConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, configName string, parameters MaintenanceConfiguration, options *MaintenanceConfigurationsClientCreateOrUpdateOptions) (MaintenanceConfigurationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, configName, parameters, options)
	if err != nil {
		return MaintenanceConfigurationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceConfigurationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MaintenanceConfigurationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MaintenanceConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configName string, parameters MaintenanceConfiguration, options *MaintenanceConfigurationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(configName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MaintenanceConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (MaintenanceConfigurationsClientCreateOrUpdateResponse, error) {
	result := MaintenanceConfigurationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceConfiguration); err != nil {
		return MaintenanceConfigurationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a maintenance configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// configName - The name of the maintenance configuration.
// options - MaintenanceConfigurationsClientDeleteOptions contains the optional parameters for the MaintenanceConfigurationsClient.Delete
// method.
func (client *MaintenanceConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsClientDeleteOptions) (MaintenanceConfigurationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, configName, options)
	if err != nil {
		return MaintenanceConfigurationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MaintenanceConfigurationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return MaintenanceConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MaintenanceConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(configName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified maintenance configuration of a managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// configName - The name of the maintenance configuration.
// options - MaintenanceConfigurationsClientGetOptions contains the optional parameters for the MaintenanceConfigurationsClient.Get
// method.
func (client *MaintenanceConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsClientGetOptions) (MaintenanceConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, configName, options)
	if err != nil {
		return MaintenanceConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MaintenanceConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MaintenanceConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MaintenanceConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *MaintenanceConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(configName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MaintenanceConfigurationsClient) getHandleResponse(resp *http.Response) (MaintenanceConfigurationsClientGetResponse, error) {
	result := MaintenanceConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceConfiguration); err != nil {
		return MaintenanceConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedClusterPager - Gets a list of maintenance configurations in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// options - MaintenanceConfigurationsClientListByManagedClusterOptions contains the optional parameters for the MaintenanceConfigurationsClient.ListByManagedCluster
// method.
func (client *MaintenanceConfigurationsClient) NewListByManagedClusterPager(resourceGroupName string, resourceName string, options *MaintenanceConfigurationsClientListByManagedClusterOptions) *runtime.Pager[MaintenanceConfigurationsClientListByManagedClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[MaintenanceConfigurationsClientListByManagedClusterResponse]{
		More: func(page MaintenanceConfigurationsClientListByManagedClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MaintenanceConfigurationsClientListByManagedClusterResponse) (MaintenanceConfigurationsClientListByManagedClusterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagedClusterCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MaintenanceConfigurationsClientListByManagedClusterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MaintenanceConfigurationsClientListByManagedClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MaintenanceConfigurationsClientListByManagedClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagedClusterHandleResponse(resp)
		},
	})
}

// listByManagedClusterCreateRequest creates the ListByManagedCluster request.
func (client *MaintenanceConfigurationsClient) listByManagedClusterCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *MaintenanceConfigurationsClientListByManagedClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedClusterHandleResponse handles the ListByManagedCluster response.
func (client *MaintenanceConfigurationsClient) listByManagedClusterHandleResponse(resp *http.Response) (MaintenanceConfigurationsClientListByManagedClusterResponse, error) {
	result := MaintenanceConfigurationsClientListByManagedClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceConfigurationListResult); err != nil {
		return MaintenanceConfigurationsClientListByManagedClusterResponse{}, err
	}
	return result, nil
}
