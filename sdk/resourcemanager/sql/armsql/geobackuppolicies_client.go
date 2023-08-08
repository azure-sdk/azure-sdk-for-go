//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// GeoBackupPoliciesClient contains the methods for the GeoBackupPolicies group.
// Don't use this type directly, use NewGeoBackupPoliciesClient() instead.
type GeoBackupPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGeoBackupPoliciesClient creates a new instance of GeoBackupPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGeoBackupPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GeoBackupPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".GeoBackupPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GeoBackupPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a database default Geo backup policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - geoBackupPolicyName - The name of the Geo backup policy. This should always be 'Default'.
//   - parameters - The required parameters for creating or updating the geo backup policy.
//   - options - GeoBackupPoliciesClientCreateOrUpdateOptions contains the optional parameters for the GeoBackupPoliciesClient.CreateOrUpdate
//     method.
func (client *GeoBackupPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName GeoBackupPolicyName, parameters GeoBackupPolicy, options *GeoBackupPoliciesClientCreateOrUpdateOptions) (GeoBackupPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, geoBackupPolicyName, parameters, options)
	if err != nil {
		return GeoBackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GeoBackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return GeoBackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GeoBackupPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName GeoBackupPolicyName, parameters GeoBackupPolicy, options *GeoBackupPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if geoBackupPolicyName == "" {
		return nil, errors.New("parameter geoBackupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{geoBackupPolicyName}", url.PathEscape(string(geoBackupPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GeoBackupPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (GeoBackupPoliciesClientCreateOrUpdateResponse, error) {
	result := GeoBackupPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeoBackupPolicy); err != nil {
		return GeoBackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a Geo backup policy for the given database resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - geoBackupPolicyName - The name of the Geo backup policy. This should always be 'Default'.
//   - options - GeoBackupPoliciesClientGetOptions contains the optional parameters for the GeoBackupPoliciesClient.Get method.
func (client *GeoBackupPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName GeoBackupPolicyName, options *GeoBackupPoliciesClientGetOptions) (GeoBackupPoliciesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, geoBackupPolicyName, options)
	if err != nil {
		return GeoBackupPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GeoBackupPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GeoBackupPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GeoBackupPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, geoBackupPolicyName GeoBackupPolicyName, options *GeoBackupPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if geoBackupPolicyName == "" {
		return nil, errors.New("parameter geoBackupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{geoBackupPolicyName}", url.PathEscape(string(geoBackupPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GeoBackupPoliciesClient) getHandleResponse(resp *http.Response) (GeoBackupPoliciesClientGetResponse, error) {
	result := GeoBackupPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeoBackupPolicy); err != nil {
		return GeoBackupPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of Geo backup policies for the given database resource.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - GeoBackupPoliciesClientListOptions contains the optional parameters for the GeoBackupPoliciesClient.NewListPager
//     method.
func (client *GeoBackupPoliciesClient) NewListPager(resourceGroupName string, serverName string, databaseName string, options *GeoBackupPoliciesClientListOptions) *runtime.Pager[GeoBackupPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GeoBackupPoliciesClientListResponse]{
		More: func(page GeoBackupPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GeoBackupPoliciesClientListResponse) (GeoBackupPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GeoBackupPoliciesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return GeoBackupPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GeoBackupPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GeoBackupPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *GeoBackupPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GeoBackupPoliciesClient) listHandleResponse(resp *http.Response) (GeoBackupPoliciesClientListResponse, error) {
	result := GeoBackupPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeoBackupPolicyListResult); err != nil {
		return GeoBackupPoliciesClientListResponse{}, err
	}
	return result, nil
}
