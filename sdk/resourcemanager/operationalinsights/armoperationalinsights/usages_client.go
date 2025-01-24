//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets a list of usage metrics for a workspace.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - UsagesClientListOptions contains the optional parameters for the UsagesClient.NewListPager method.
func (client *UsagesClient) NewListPager(resourceGroupName string, workspaceName string, options *UsagesClientListOptions) *runtime.Pager[UsagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[UsagesClientListResponse]{
		More: func(page UsagesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *UsagesClientListResponse) (UsagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UsagesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			if err != nil {
				return UsagesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UsagesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UsagesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *UsagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *UsagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsagesClient) listHandleResponse(resp *http.Response) (UsagesClientListResponse, error) {
	result := UsagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListUsagesResult); err != nil {
		return UsagesClientListResponse{}, err
	}
	return result, nil
}
