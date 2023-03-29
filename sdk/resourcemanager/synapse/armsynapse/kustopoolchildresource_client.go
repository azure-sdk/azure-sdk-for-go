//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

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

// KustoPoolChildResourceClient contains the methods for the KustoPoolChildResource group.
// Don't use this type directly, use NewKustoPoolChildResourceClient() instead.
type KustoPoolChildResourceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKustoPoolChildResourceClient creates a new instance of KustoPoolChildResourceClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKustoPoolChildResourceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KustoPoolChildResourceClient, error) {
	cl, err := arm.NewClient(moduleName+".KustoPoolChildResourceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KustoPoolChildResourceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks that the Kusto Pool child resource name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Kusto Pool child resource.
//   - options - KustoPoolChildResourceClientCheckNameAvailabilityOptions contains the optional parameters for the KustoPoolChildResourceClient.CheckNameAvailability
//     method.
func (client *KustoPoolChildResourceClient) CheckNameAvailability(ctx context.Context, workspaceName string, kustoPoolName string, resourceGroupName string, resourceName DatabaseCheckNameRequest, options *KustoPoolChildResourceClientCheckNameAvailabilityOptions) (KustoPoolChildResourceClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, workspaceName, kustoPoolName, resourceGroupName, resourceName, options)
	if err != nil {
		return KustoPoolChildResourceClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KustoPoolChildResourceClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolChildResourceClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *KustoPoolChildResourceClient) checkNameAvailabilityCreateRequest(ctx context.Context, workspaceName string, kustoPoolName string, resourceGroupName string, resourceName DatabaseCheckNameRequest, options *KustoPoolChildResourceClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/checkNameAvailability"
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resourceName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *KustoPoolChildResourceClient) checkNameAvailabilityHandleResponse(resp *http.Response) (KustoPoolChildResourceClientCheckNameAvailabilityResponse, error) {
	result := KustoPoolChildResourceClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return KustoPoolChildResourceClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}
