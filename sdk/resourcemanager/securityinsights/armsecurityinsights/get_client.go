//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// GetClient contains the methods for the Get group.
// Don't use this type directly, use NewGetClient() instead.
type GetClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGetClient creates a new instance of GetClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GetClient, error) {
	cl, err := arm.NewClient(moduleName+".GetClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GetClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// SingleRecommendation - Gets a recommendation by its id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - recommendationID - Recommendation Id.
//   - options - GetClientSingleRecommendationOptions contains the optional parameters for the GetClient.SingleRecommendation
//     method.
func (client *GetClient) SingleRecommendation(ctx context.Context, resourceGroupName string, workspaceName string, recommendationID string, options *GetClientSingleRecommendationOptions) (GetClientSingleRecommendationResponse, error) {
	var err error
	req, err := client.singleRecommendationCreateRequest(ctx, resourceGroupName, workspaceName, recommendationID, options)
	if err != nil {
		return GetClientSingleRecommendationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetClientSingleRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GetClientSingleRecommendationResponse{}, err
	}
	resp, err := client.singleRecommendationHandleResponse(httpResp)
	return resp, err
}

// singleRecommendationCreateRequest creates the SingleRecommendation request.
func (client *GetClient) singleRecommendationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, recommendationID string, options *GetClientSingleRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/recommendations/{recommendationId}"
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
	urlPath = strings.ReplaceAll(urlPath, "{recommendationId}", url.PathEscape(recommendationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// singleRecommendationHandleResponse handles the SingleRecommendation response.
func (client *GetClient) singleRecommendationHandleResponse(resp *http.Response) (GetClientSingleRecommendationResponse, error) {
	result := GetClientSingleRecommendationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Recommendation); err != nil {
		return GetClientSingleRecommendationResponse{}, err
	}
	return result, nil
}
