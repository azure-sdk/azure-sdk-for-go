//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FeaturesClient contains the methods for the Features group.
// Don't use this type directly, use NewFeaturesClient() instead.
type FeaturesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFeaturesClient creates a new instance of FeaturesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFeaturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FeaturesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FeaturesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get feature.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - featuresetName - Feature set name. This is case-sensitive.
//   - featuresetVersion - Feature set version identifier. This is case-sensitive.
//   - featureName - Feature Name. This is case-sensitive.
//   - options - FeaturesClientGetOptions contains the optional parameters for the FeaturesClient.Get method.
func (client *FeaturesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, featuresetName string, featuresetVersion string, featureName string, options *FeaturesClientGetOptions) (FeaturesClientGetResponse, error) {
	var err error
	const operationName = "FeaturesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, featuresetName, featuresetVersion, featureName, options)
	if err != nil {
		return FeaturesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FeaturesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FeaturesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FeaturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, featuresetName string, featuresetVersion string, featureName string, options *FeaturesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{featuresetName}/versions/{featuresetVersion}/features/{featureName}"
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
	if featuresetName == "" {
		return nil, errors.New("parameter featuresetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featuresetName}", url.PathEscape(featuresetName))
	if featuresetVersion == "" {
		return nil, errors.New("parameter featuresetVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featuresetVersion}", url.PathEscape(featuresetVersion))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FeaturesClient) getHandleResponse(resp *http.Response) (FeaturesClientGetResponse, error) {
	result := FeaturesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Feature); err != nil {
		return FeaturesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Features.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - featuresetName - Featureset name. This is case-sensitive.
//   - featuresetVersion - Featureset Version identifier. This is case-sensitive.
//   - options - FeaturesClientListOptions contains the optional parameters for the FeaturesClient.NewListPager method.
func (client *FeaturesClient) NewListPager(resourceGroupName string, workspaceName string, featuresetName string, featuresetVersion string, options *FeaturesClientListOptions) *runtime.Pager[FeaturesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FeaturesClientListResponse]{
		More: func(page FeaturesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FeaturesClientListResponse) (FeaturesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FeaturesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, featuresetName, featuresetVersion, options)
			}, nil)
			if err != nil {
				return FeaturesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *FeaturesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, featuresetName string, featuresetVersion string, options *FeaturesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{featuresetName}/versions/{featuresetVersion}/features"
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
	if featuresetName == "" {
		return nil, errors.New("parameter featuresetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featuresetName}", url.PathEscape(featuresetName))
	if featuresetVersion == "" {
		return nil, errors.New("parameter featuresetVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featuresetVersion}", url.PathEscape(featuresetVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2024-10-01-preview")
	if options != nil && options.Description != nil {
		reqQP.Set("description", *options.Description)
	}
	if options != nil && options.FeatureName != nil {
		reqQP.Set("featureName", *options.FeatureName)
	}
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FeaturesClient) listHandleResponse(resp *http.Response) (FeaturesClientListResponse, error) {
	result := FeaturesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeatureResourceArmPaginatedResult); err != nil {
		return FeaturesClientListResponse{}, err
	}
	return result, nil
}
