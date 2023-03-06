//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armelastic

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

// TrafficFiltersClient contains the methods for the TrafficFilters group.
// Don't use this type directly, use NewTrafficFiltersClient() instead.
type TrafficFiltersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTrafficFiltersClient creates a new instance of TrafficFiltersClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTrafficFiltersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TrafficFiltersClient, error) {
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
	client := &TrafficFiltersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Delete - Delete traffic filter from the account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group to which the Elastic resource belongs.
//   - monitorName - Monitor resource name
//   - options - TrafficFiltersClientDeleteOptions contains the optional parameters for the TrafficFiltersClient.Delete method.
func (client *TrafficFiltersClient) Delete(ctx context.Context, resourceGroupName string, monitorName string, options *TrafficFiltersClientDeleteOptions) (TrafficFiltersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return TrafficFiltersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrafficFiltersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TrafficFiltersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return TrafficFiltersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TrafficFiltersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *TrafficFiltersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/deleteTrafficFilter"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	if options != nil && options.RulesetID != nil {
		reqQP.Set("rulesetId", *options.RulesetID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
