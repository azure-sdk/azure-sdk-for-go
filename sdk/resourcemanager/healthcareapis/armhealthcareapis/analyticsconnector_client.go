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

// AnalyticsConnectorClient contains the methods for the AnalyticsConnector group.
// Don't use this type directly, use NewAnalyticsConnectorClient() instead.
type AnalyticsConnectorClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAnalyticsConnectorClient creates a new instance of AnalyticsConnectorClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAnalyticsConnectorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AnalyticsConnectorClient, error) {
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
	client := &AnalyticsConnectorClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginUpdate - Patch Analytics Connector Service details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// analyticsConnectorName - The name of Analytics Connector resource.
// analyticsConnectorPatchResource - The parameters for updating a Analytics Connector.
// options - AnalyticsConnectorClientBeginUpdateOptions contains the optional parameters for the AnalyticsConnectorClient.BeginUpdate
// method.
func (client *AnalyticsConnectorClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, analyticsConnectorPatchResource AnalyticsConnectorPatchResource, options *AnalyticsConnectorClientBeginUpdateOptions) (*runtime.Poller[AnalyticsConnectorClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, analyticsConnectorName, analyticsConnectorPatchResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AnalyticsConnectorClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AnalyticsConnectorClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch Analytics Connector Service details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
func (client *AnalyticsConnectorClient) update(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, analyticsConnectorPatchResource AnalyticsConnectorPatchResource, options *AnalyticsConnectorClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, analyticsConnectorName, analyticsConnectorPatchResource, options)
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
func (client *AnalyticsConnectorClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, analyticsConnectorName string, analyticsConnectorPatchResource AnalyticsConnectorPatchResource, options *AnalyticsConnectorClientBeginUpdateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, analyticsConnectorPatchResource)
}
