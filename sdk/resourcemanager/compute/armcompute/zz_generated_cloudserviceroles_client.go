//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// CloudServiceRolesClient contains the methods for the CloudServiceRoles group.
// Don't use this type directly, use NewCloudServiceRolesClient() instead.
type CloudServiceRolesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCloudServiceRolesClient creates a new instance of CloudServiceRolesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCloudServiceRolesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudServiceRolesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CloudServiceRolesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a role from a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04
// roleName - Name of the role.
// resourceGroupName - Name of the resource group.
// cloudServiceName - Name of the cloud service.
// options - CloudServiceRolesClientGetOptions contains the optional parameters for the CloudServiceRolesClient.Get method.
func (client *CloudServiceRolesClient) Get(ctx context.Context, roleName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRolesClientGetOptions) (CloudServiceRolesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, roleName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return CloudServiceRolesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceRolesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceRolesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CloudServiceRolesClient) getCreateRequest(ctx context.Context, roleName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRolesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roles/{roleName}"
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudServiceRolesClient) getHandleResponse(resp *http.Response) (CloudServiceRolesClientGetResponse, error) {
	result := CloudServiceRolesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudServiceRole); err != nil {
		return CloudServiceRolesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all roles in a cloud service. Use nextLink property in the response to get the next page
// of roles. Do this till nextLink is null to fetch all the roles.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04
// resourceGroupName - Name of the resource group.
// cloudServiceName - Name of the cloud service.
// options - CloudServiceRolesClientListOptions contains the optional parameters for the CloudServiceRolesClient.List method.
func (client *CloudServiceRolesClient) NewListPager(resourceGroupName string, cloudServiceName string, options *CloudServiceRolesClientListOptions) *runtime.Pager[CloudServiceRolesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudServiceRolesClientListResponse]{
		More: func(page CloudServiceRolesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudServiceRolesClientListResponse) (CloudServiceRolesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, cloudServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CloudServiceRolesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CloudServiceRolesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CloudServiceRolesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CloudServiceRolesClient) listCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, options *CloudServiceRolesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roles"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CloudServiceRolesClient) listHandleResponse(resp *http.Response) (CloudServiceRolesClientListResponse, error) {
	result := CloudServiceRolesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudServiceRoleListResult); err != nil {
		return CloudServiceRolesClientListResponse{}, err
	}
	return result, nil
}
