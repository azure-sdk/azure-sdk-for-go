//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

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

// SnapshotsClient contains the methods for the Snapshots group.
// Don't use this type directly, use NewSnapshotsClient() instead.
type SnapshotsClient struct {
	internal *arm.Client
}

// NewSnapshotsClient creates a new instance of SnapshotsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSnapshotsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SnapshotsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SnapshotsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Get the AppComplianceAutomation snapshot list.
//
// Generated from API version 2022-11-16-preview
//   - reportName - Report Name.
//   - options - SnapshotsClientListOptions contains the optional parameters for the SnapshotsClient.NewListPager method.
func (client *SnapshotsClient) NewListPager(reportName string, options *SnapshotsClientListOptions) *runtime.Pager[SnapshotsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SnapshotsClientListResponse]{
		More: func(page SnapshotsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SnapshotsClientListResponse) (SnapshotsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SnapshotsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, reportName, options)
			}, nil)
			if err != nil {
				return SnapshotsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SnapshotsClient) listCreateRequest(ctx context.Context, reportName string, options *SnapshotsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/snapshots"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-16-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.ReportCreatorTenantID != nil {
		reqQP.Set("reportCreatorTenantId", *options.ReportCreatorTenantID)
	}
	if options != nil && options.OfferGUID != nil {
		reqQP.Set("offerGuid", *options.OfferGUID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SnapshotsClient) listHandleResponse(resp *http.Response) (SnapshotsClientListResponse, error) {
	result := SnapshotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotResourceList); err != nil {
		return SnapshotsClientListResponse{}, err
	}
	return result, nil
}