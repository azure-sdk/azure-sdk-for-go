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

// UpdateClient contains the methods for the Update group.
// Don't use this type directly, use NewUpdateClient() instead.
type UpdateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUpdateClient creates a new instance of UpdateClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUpdateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UpdateClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UpdateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Recommendation - Patch a recommendation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - recommendationID - Recommendation Id.
//   - recommendationPatch - Recommendation Fields to Update.
//   - options - UpdateClientRecommendationOptions contains the optional parameters for the UpdateClient.Recommendation method.
func (client *UpdateClient) Recommendation(ctx context.Context, resourceGroupName string, workspaceName string, recommendationID string, recommendationPatch RecommendationPatch, options *UpdateClientRecommendationOptions) (UpdateClientRecommendationResponse, error) {
	var err error
	const operationName = "UpdateClient.Recommendation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.recommendationCreateRequest(ctx, resourceGroupName, workspaceName, recommendationID, recommendationPatch, options)
	if err != nil {
		return UpdateClientRecommendationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UpdateClientRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UpdateClientRecommendationResponse{}, err
	}
	resp, err := client.recommendationHandleResponse(httpResp)
	return resp, err
}

// recommendationCreateRequest creates the Recommendation request.
func (client *UpdateClient) recommendationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, recommendationID string, recommendationPatch RecommendationPatch, _ *UpdateClientRecommendationOptions) (*policy.Request, error) {
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
	if recommendationID == "" {
		return nil, errors.New("parameter recommendationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendationId}", url.PathEscape(recommendationID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, recommendationPatch); err != nil {
		return nil, err
	}
	return req, nil
}

// recommendationHandleResponse handles the Recommendation response.
func (client *UpdateClient) recommendationHandleResponse(resp *http.Response) (UpdateClientRecommendationResponse, error) {
	result := UpdateClientRecommendationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Recommendation); err != nil {
		return UpdateClientRecommendationResponse{}, err
	}
	return result, nil
}
