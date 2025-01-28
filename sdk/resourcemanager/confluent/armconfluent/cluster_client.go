//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

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

// ClusterClient contains the methods for the Cluster group.
// Don't use this type directly, use NewClusterClient() instead.
type ClusterClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClusterClient creates a new instance of ClusterClient with the specified values.
//   - subscriptionID - Microsoft Azure subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClusterClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClusterClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClusterClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create confluent clusters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Organization resource name
//   - environmentID - Confluent environment id
//   - clusterID - Confluent kafka or schema registry cluster id
//   - options - ClusterClientCreateOrUpdateOptions contains the optional parameters for the ClusterClient.CreateOrUpdate method.
func (client *ClusterClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, environmentID string, clusterID string, options *ClusterClientCreateOrUpdateOptions) (ClusterClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ClusterClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, organizationName, environmentID, clusterID, options)
	if err != nil {
		return ClusterClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClusterClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ClusterClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ClusterClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, environmentID string, clusterID string, options *ClusterClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}/environments/{environmentId}/clusters/{clusterId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	if clusterID == "" {
		return nil, errors.New("parameter clusterID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterId}", url.PathEscape(clusterID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ClusterClient) createOrUpdateHandleResponse(resp *http.Response) (ClusterClientCreateOrUpdateResponse, error) {
	result := ClusterClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SCClusterRecord); err != nil {
		return ClusterClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete confluent cluster by id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Organization resource name
//   - environmentID - Confluent environment id
//   - clusterID - Confluent kafka or schema registry cluster id
//   - options - ClusterClientBeginDeleteOptions contains the optional parameters for the ClusterClient.BeginDelete method.
func (client *ClusterClient) BeginDelete(ctx context.Context, resourceGroupName string, organizationName string, environmentID string, clusterID string, options *ClusterClientBeginDeleteOptions) (*runtime.Poller[ClusterClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, organizationName, environmentID, clusterID, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClusterClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClusterClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete confluent cluster by id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *ClusterClient) deleteOperation(ctx context.Context, resourceGroupName string, organizationName string, environmentID string, clusterID string, options *ClusterClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ClusterClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, environmentID, clusterID, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ClusterClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, environmentID string, clusterID string, options *ClusterClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}/environments/{environmentId}/clusters/{clusterId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	if clusterID == "" {
		return nil, errors.New("parameter clusterID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterId}", url.PathEscape(clusterID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
