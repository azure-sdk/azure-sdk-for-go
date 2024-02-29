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

// ThreatIntelligenceClient contains the methods for the ThreatIntelligence group.
// Don't use this type directly, use NewThreatIntelligenceClient() instead.
type ThreatIntelligenceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewThreatIntelligenceClient creates a new instance of ThreatIntelligenceClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewThreatIntelligenceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ThreatIntelligenceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ThreatIntelligenceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Count - Gets the count of all TI objects for the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - tiType - TI type
//   - options - ThreatIntelligenceClientCountOptions contains the optional parameters for the ThreatIntelligenceClient.Count
//     method.
func (client *ThreatIntelligenceClient) Count(ctx context.Context, resourceGroupName string, workspaceName string, tiType TiType, options *ThreatIntelligenceClientCountOptions) (ThreatIntelligenceClientCountResponse, error) {
	var err error
	const operationName = "ThreatIntelligenceClient.Count"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.countCreateRequest(ctx, resourceGroupName, workspaceName, tiType, options)
	if err != nil {
		return ThreatIntelligenceClientCountResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThreatIntelligenceClientCountResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ThreatIntelligenceClientCountResponse{}, err
	}
	resp, err := client.countHandleResponse(httpResp)
	return resp, err
}

// countCreateRequest creates the Count request.
func (client *ThreatIntelligenceClient) countCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, tiType TiType, options *ThreatIntelligenceClientCountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/{tiType}/count"
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
	if tiType == "" {
		return nil, errors.New("parameter tiType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tiType}", url.PathEscape(string(tiType)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Query != nil {
		if err := runtime.MarshalAsJSON(req, *options.Query); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// countHandleResponse handles the Count response.
func (client *ThreatIntelligenceClient) countHandleResponse(resp *http.Response) (ThreatIntelligenceClientCountResponse, error) {
	result := ThreatIntelligenceClientCountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThreatIntelligenceCount); err != nil {
		return ThreatIntelligenceClientCountResponse{}, err
	}
	return result, nil
}

// NewQueryPager - Gets all TI objects for the workspace.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - tiType - TI type
//   - options - ThreatIntelligenceClientQueryOptions contains the optional parameters for the ThreatIntelligenceClient.NewQueryPager
//     method.
func (client *ThreatIntelligenceClient) NewQueryPager(resourceGroupName string, workspaceName string, tiType TiType, options *ThreatIntelligenceClientQueryOptions) *runtime.Pager[ThreatIntelligenceClientQueryResponse] {
	return runtime.NewPager(runtime.PagingHandler[ThreatIntelligenceClientQueryResponse]{
		More: func(page ThreatIntelligenceClientQueryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ThreatIntelligenceClientQueryResponse) (ThreatIntelligenceClientQueryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ThreatIntelligenceClient.NewQueryPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.queryCreateRequest(ctx, resourceGroupName, workspaceName, tiType, options)
			}, nil)
			if err != nil {
				return ThreatIntelligenceClientQueryResponse{}, err
			}
			return client.queryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// queryCreateRequest creates the Query request.
func (client *ThreatIntelligenceClient) queryCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, tiType TiType, options *ThreatIntelligenceClientQueryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/{tiType}/query"
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
	if tiType == "" {
		return nil, errors.New("parameter tiType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tiType}", url.PathEscape(string(tiType)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Query != nil {
		if err := runtime.MarshalAsJSON(req, *options.Query); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// queryHandleResponse handles the Query response.
func (client *ThreatIntelligenceClient) queryHandleResponse(resp *http.Response) (ThreatIntelligenceClientQueryResponse, error) {
	result := ThreatIntelligenceClientQueryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThreatIntelligenceList); err != nil {
		return ThreatIntelligenceClientQueryResponse{}, err
	}
	return result, nil
}
