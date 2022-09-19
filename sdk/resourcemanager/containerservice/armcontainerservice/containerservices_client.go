//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ContainerServicesClient contains the methods for the ContainerServices group.
// Don't use this type directly, use NewContainerServicesClient() instead.
type ContainerServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContainerServicesClient creates a new instance of ContainerServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContainerServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerServicesClient, error) {
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
	client := &ContainerServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListOrchestrators - Gets a list of supported orchestrators in the specified subscription. The operation returns properties
// of each orchestrator including version, available upgrades and whether that version or upgrades
// are in preview.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-03-preview
// location - The name of Azure region.
// options - ContainerServicesClientListOrchestratorsOptions contains the optional parameters for the ContainerServicesClient.ListOrchestrators
// method.
func (client *ContainerServicesClient) ListOrchestrators(ctx context.Context, location string, options *ContainerServicesClientListOrchestratorsOptions) (ContainerServicesClientListOrchestratorsResponse, error) {
	req, err := client.listOrchestratorsCreateRequest(ctx, location, options)
	if err != nil {
		return ContainerServicesClientListOrchestratorsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerServicesClientListOrchestratorsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerServicesClientListOrchestratorsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listOrchestratorsHandleResponse(resp)
}

// listOrchestratorsCreateRequest creates the ListOrchestrators request.
func (client *ContainerServicesClient) listOrchestratorsCreateRequest(ctx context.Context, location string, options *ContainerServicesClientListOrchestratorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/orchestrators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-03-preview")
	if options != nil && options.ResourceType != nil {
		reqQP.Set("resource-type", *options.ResourceType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOrchestratorsHandleResponse handles the ListOrchestrators response.
func (client *ContainerServicesClient) listOrchestratorsHandleResponse(resp *http.Response) (ContainerServicesClientListOrchestratorsResponse, error) {
	result := ContainerServicesClientListOrchestratorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrchestratorVersionProfileListResult); err != nil {
		return ContainerServicesClientListOrchestratorsResponse{}, err
	}
	return result, nil
}
