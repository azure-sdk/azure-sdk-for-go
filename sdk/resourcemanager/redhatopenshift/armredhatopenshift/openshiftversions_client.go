//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredhatopenshift

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

// OpenShiftVersionsClient contains the methods for the OpenShiftVersions group.
// Don't use this type directly, use NewOpenShiftVersionsClient() instead.
type OpenShiftVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOpenShiftVersionsClient creates a new instance of OpenShiftVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOpenShiftVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OpenShiftVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".OpenShiftVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OpenShiftVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - The operation returns the installable OpenShift versions as strings.
//
// Generated from API version 2023-04-01
//   - location - The name of Azure region.
//   - options - OpenShiftVersionsClientListOptions contains the optional parameters for the OpenShiftVersionsClient.NewListPager
//     method.
func (client *OpenShiftVersionsClient) NewListPager(location string, options *OpenShiftVersionsClientListOptions) *runtime.Pager[OpenShiftVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OpenShiftVersionsClientListResponse]{
		More: func(page OpenShiftVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OpenShiftVersionsClientListResponse) (OpenShiftVersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OpenShiftVersionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OpenShiftVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OpenShiftVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OpenShiftVersionsClient) listCreateRequest(ctx context.Context, location string, options *OpenShiftVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.RedHatOpenShift/locations/{location}/openshiftversions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OpenShiftVersionsClient) listHandleResponse(resp *http.Response) (OpenShiftVersionsClientListResponse, error) {
	result := OpenShiftVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftVersionList); err != nil {
		return OpenShiftVersionsClientListResponse{}, err
	}
	return result, nil
}
