//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhealthcareapis

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

// AnalyticsConnectorsClient contains the methods for the AnalyticsConnectors group.
// Don't use this type directly, use NewAnalyticsConnectorsClient() instead.
type AnalyticsConnectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAnalyticsConnectorsClient creates a new instance of AnalyticsConnectorsClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAnalyticsConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AnalyticsConnectorsClient, error) {
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
	client := &AnalyticsConnectorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Analytics Connector resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// analyticsConnectorName - The name of Analytics Connector resource.
// analyticsConnector - The parameters for creating or updating a Analytics Connector resource.
// options - AnalyticsConnectorsClientBeginCreateOrUpdateOptions contains the optional parameters for the AnalyticsConnectorsClient.BeginCreateOrUpdate
// method.
func (client *AnalyticsConnectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, analyticsConnector AnalyticsConnector, options *AnalyticsConnectorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AnalyticsConnectorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, analyticsConnectorName, analyticsConnector, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AnalyticsConnectorsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AnalyticsConnectorsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Analytics Connector resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
func (client *AnalyticsConnectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, analyticsConnector AnalyticsConnector, options *AnalyticsConnectorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, analyticsConnectorName, analyticsConnector, options)
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
func (client *AnalyticsConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, analyticsConnector AnalyticsConnector, options *AnalyticsConnectorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/analyticsconnectors/{analyticsConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if analyticsConnectorName == "" {
		return nil, errors.New("parameter analyticsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{analyticsConnectorName}", url.PathEscape(analyticsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, analyticsConnector)
}

// BeginDelete - Deletes a Analytics Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// analyticsConnectorName - The name of Analytics Connector resource.
// options - AnalyticsConnectorsClientBeginDeleteOptions contains the optional parameters for the AnalyticsConnectorsClient.BeginDelete
// method.
func (client *AnalyticsConnectorsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, options *AnalyticsConnectorsClientBeginDeleteOptions) (*runtime.Poller[AnalyticsConnectorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, analyticsConnectorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AnalyticsConnectorsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AnalyticsConnectorsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Analytics Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
func (client *AnalyticsConnectorsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, options *AnalyticsConnectorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, analyticsConnectorName, options)
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
func (client *AnalyticsConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, options *AnalyticsConnectorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/analyticsconnectors/{analyticsConnectorName}"
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
	if analyticsConnectorName == "" {
		return nil, errors.New("parameter analyticsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{analyticsConnectorName}", url.PathEscape(analyticsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified Analytics Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// analyticsConnectorName - The name of Analytics Connector resource.
// options - AnalyticsConnectorsClientGetOptions contains the optional parameters for the AnalyticsConnectorsClient.Get method.
func (client *AnalyticsConnectorsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, options *AnalyticsConnectorsClientGetOptions) (AnalyticsConnectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, analyticsConnectorName, options)
	if err != nil {
		return AnalyticsConnectorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AnalyticsConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AnalyticsConnectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AnalyticsConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, options *AnalyticsConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/analyticsconnectors/{analyticsConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if analyticsConnectorName == "" {
		return nil, errors.New("parameter analyticsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{analyticsConnectorName}", url.PathEscape(analyticsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AnalyticsConnectorsClient) getHandleResponse(resp *http.Response) (AnalyticsConnectorsClientGetResponse, error) {
	result := AnalyticsConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AnalyticsConnector); err != nil {
		return AnalyticsConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists all Analytics Connectors for the given workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// options - AnalyticsConnectorsClientListByWorkspaceOptions contains the optional parameters for the AnalyticsConnectorsClient.ListByWorkspace
// method.
func (client *AnalyticsConnectorsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *AnalyticsConnectorsClientListByWorkspaceOptions) *runtime.Pager[AnalyticsConnectorsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[AnalyticsConnectorsClientListByWorkspaceResponse]{
		More: func(page AnalyticsConnectorsClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AnalyticsConnectorsClientListByWorkspaceResponse) (AnalyticsConnectorsClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AnalyticsConnectorsClientListByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AnalyticsConnectorsClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AnalyticsConnectorsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *AnalyticsConnectorsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AnalyticsConnectorsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/analyticsconnectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *AnalyticsConnectorsClient) listByWorkspaceHandleResponse(resp *http.Response) (AnalyticsConnectorsClientListByWorkspaceResponse, error) {
	result := AnalyticsConnectorsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AnalyticsConnectorCollection); err != nil {
		return AnalyticsConnectorsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
