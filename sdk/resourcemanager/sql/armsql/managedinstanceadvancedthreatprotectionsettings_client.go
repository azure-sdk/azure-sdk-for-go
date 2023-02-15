//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// ManagedInstanceAdvancedThreatProtectionSettingsClient contains the methods for the ManagedInstanceAdvancedThreatProtectionSettings group.
// Don't use this type directly, use NewManagedInstanceAdvancedThreatProtectionSettingsClient() instead.
type ManagedInstanceAdvancedThreatProtectionSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedInstanceAdvancedThreatProtectionSettingsClient creates a new instance of ManagedInstanceAdvancedThreatProtectionSettingsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedInstanceAdvancedThreatProtectionSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstanceAdvancedThreatProtectionSettingsClient, error) {
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
	client := &ManagedInstanceAdvancedThreatProtectionSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates Advanced Threat Protection settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - advancedThreatProtectionName - The name of the Advanced Threat Protection state.
//   - parameters - The managed instance Advanced Threat Protection state.
//   - options - ManagedInstanceAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions contains the optional parameters
//     for the ManagedInstanceAdvancedThreatProtectionSettingsClient.BeginCreateOrUpdate method.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters ManagedInstanceAdvancedThreatProtection, options *ManagedInstanceAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstanceAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, advancedThreatProtectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstanceAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstanceAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates Advanced Threat Protection settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters ManagedInstanceAdvancedThreatProtection, options *ManagedInstanceAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, advancedThreatProtectionName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters ManagedInstanceAdvancedThreatProtection, options *ManagedInstanceAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if advancedThreatProtectionName == "" {
		return nil, errors.New("parameter advancedThreatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advancedThreatProtectionName}", url.PathEscape(string(advancedThreatProtectionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Get a managed instance's Advanced Threat Protection state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - advancedThreatProtectionName - The name of the Advanced Threat Protection state.
//   - options - ManagedInstanceAdvancedThreatProtectionSettingsClientGetOptions contains the optional parameters for the ManagedInstanceAdvancedThreatProtectionSettingsClient.Get
//     method.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, advancedThreatProtectionName AdvancedThreatProtectionName, options *ManagedInstanceAdvancedThreatProtectionSettingsClientGetOptions) (ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, advancedThreatProtectionName, options)
	if err != nil {
		return ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, advancedThreatProtectionName AdvancedThreatProtectionName, options *ManagedInstanceAdvancedThreatProtectionSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if advancedThreatProtectionName == "" {
		return nil, errors.New("parameter advancedThreatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advancedThreatProtectionName}", url.PathEscape(string(advancedThreatProtectionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) getHandleResponse(resp *http.Response) (ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse, error) {
	result := ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceAdvancedThreatProtection); err != nil {
		return ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Get the managed instance's Advanced Threat Protection settings.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceOptions contains the optional parameters for
//     the ManagedInstanceAdvancedThreatProtectionSettingsClient.NewListByInstancePager method.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceOptions) *runtime.Pager[ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse]{
		More: func(page ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse) (ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstanceHandleResponse(resp)
		},
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/advancedThreatProtectionSettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ManagedInstanceAdvancedThreatProtectionSettingsClient) listByInstanceHandleResponse(resp *http.Response) (ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse, error) {
	result := ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceAdvancedThreatProtectionListResult); err != nil {
		return ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse{}, err
	}
	return result, nil
}
