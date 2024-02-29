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

// ReevaluateClient contains the methods for the Reevaluate group.
// Don't use this type directly, use NewReevaluateClient() instead.
type ReevaluateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReevaluateClient creates a new instance of ReevaluateClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReevaluateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReevaluateClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReevaluateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Recommendation - Reevaluate a recommendation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - recommendationID - Recommendation Id.
//   - options - ReevaluateClientRecommendationOptions contains the optional parameters for the ReevaluateClient.Recommendation
//     method.
func (client *ReevaluateClient) Recommendation(ctx context.Context, resourceGroupName string, workspaceName string, recommendationID string, options *ReevaluateClientRecommendationOptions) (ReevaluateClientRecommendationResponse, error) {
	var err error
	const operationName = "ReevaluateClient.Recommendation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.recommendationCreateRequest(ctx, resourceGroupName, workspaceName, recommendationID, options)
	if err != nil {
		return ReevaluateClientRecommendationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReevaluateClientRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReevaluateClientRecommendationResponse{}, err
	}
	resp, err := client.recommendationHandleResponse(httpResp)
	return resp, err
}

// recommendationCreateRequest creates the Recommendation request.
func (client *ReevaluateClient) recommendationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, recommendationID string, options *ReevaluateClientRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/recommendations/{recommendationId}/triggerEvaluation"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// recommendationHandleResponse handles the Recommendation response.
func (client *ReevaluateClient) recommendationHandleResponse(resp *http.Response) (ReevaluateClientRecommendationResponse, error) {
	result := ReevaluateClientRecommendationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReevaluateResponse); err != nil {
		return ReevaluateClientRecommendationResponse{}, err
	}
	return result, nil
}
