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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DeploymentInfoClient contains the methods for the DeploymentInfo group.
// Don't use this type directly, use NewDeploymentInfoClient() instead.
type DeploymentInfoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeploymentInfoClient creates a new instance of DeploymentInfoClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeploymentInfoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeploymentInfoClient, error) {
	cl, err := arm.NewClient(moduleName+".DeploymentInfoClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeploymentInfoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Fetch information regarding Elastic cloud deployment corresponding to the Elastic monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15-preview
//   - resourceGroupName - The name of the resource group to which the Elastic resource belongs.
//   - monitorName - Monitor resource name
//   - options - DeploymentInfoClientListOptions contains the optional parameters for the DeploymentInfoClient.List method.
func (client *DeploymentInfoClient) List(ctx context.Context, resourceGroupName string, monitorName string, options *DeploymentInfoClientListOptions) (DeploymentInfoClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return DeploymentInfoClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentInfoClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentInfoClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DeploymentInfoClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *DeploymentInfoClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/listDeploymentInfo"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeploymentInfoClient) listHandleResponse(resp *http.Response) (DeploymentInfoClientListResponse, error) {
	result := DeploymentInfoClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentInfoResponse); err != nil {
		return DeploymentInfoClientListResponse{}, err
	}
	return result, nil
}
