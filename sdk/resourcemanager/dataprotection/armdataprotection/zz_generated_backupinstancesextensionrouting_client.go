//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BackupInstancesExtensionRoutingClient contains the methods for the BackupInstancesExtensionRouting group.
// Don't use this type directly, use NewBackupInstancesExtensionRoutingClient() instead.
type BackupInstancesExtensionRoutingClient struct {
	host string
	pl   runtime.Pipeline
}

// NewBackupInstancesExtensionRoutingClient creates a new instance of BackupInstancesExtensionRoutingClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupInstancesExtensionRoutingClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupInstancesExtensionRoutingClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &BackupInstancesExtensionRoutingClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// List - Gets a list backup instances associated with a tracked resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - ARM path of the resource to be protected using Microsoft.DataProtection
// options - BackupInstancesExtensionRoutingClientListOptions contains the optional parameters for the BackupInstancesExtensionRoutingClient.List
// method.
func (client *BackupInstancesExtensionRoutingClient) List(resourceID string, options *BackupInstancesExtensionRoutingClientListOptions) *runtime.Pager[BackupInstancesExtensionRoutingClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[BackupInstancesExtensionRoutingClientListResponse]{
		More: func(page BackupInstancesExtensionRoutingClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupInstancesExtensionRoutingClientListResponse) (BackupInstancesExtensionRoutingClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BackupInstancesExtensionRoutingClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BackupInstancesExtensionRoutingClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupInstancesExtensionRoutingClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BackupInstancesExtensionRoutingClient) listCreateRequest(ctx context.Context, resourceID string, options *BackupInstancesExtensionRoutingClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.DataProtection/backupInstances"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", url.PathEscape(resourceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupInstancesExtensionRoutingClient) listHandleResponse(resp *http.Response) (BackupInstancesExtensionRoutingClientListResponse, error) {
	result := BackupInstancesExtensionRoutingClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupInstanceResourceList); err != nil {
		return BackupInstancesExtensionRoutingClientListResponse{}, err
	}
	return result, nil
}
