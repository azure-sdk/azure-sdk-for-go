//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ContainerAppsRevisionReplicasClient contains the methods for the ContainerAppsRevisionReplicas group.
// Don't use this type directly, use NewContainerAppsRevisionReplicasClient() instead.
type ContainerAppsRevisionReplicasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsRevisionReplicasClient creates a new instance of ContainerAppsRevisionReplicasClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsRevisionReplicasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsRevisionReplicasClient, error) {
	cl, err := arm.NewClient(moduleName+".ContainerAppsRevisionReplicasClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsRevisionReplicasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetReplica - Get a replica for a Container App Revision.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - revisionName - Name of the Container App Revision.
//   - replicaName - Name of the Container App Revision Replica.
//   - options - ContainerAppsRevisionReplicasClientGetReplicaOptions contains the optional parameters for the ContainerAppsRevisionReplicasClient.GetReplica
//     method.
func (client *ContainerAppsRevisionReplicasClient) GetReplica(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, replicaName string, options *ContainerAppsRevisionReplicasClientGetReplicaOptions) (ContainerAppsRevisionReplicasClientGetReplicaResponse, error) {
	var err error
	req, err := client.getReplicaCreateRequest(ctx, resourceGroupName, containerAppName, revisionName, replicaName, options)
	if err != nil {
		return ContainerAppsRevisionReplicasClientGetReplicaResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsRevisionReplicasClientGetReplicaResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsRevisionReplicasClientGetReplicaResponse{}, err
	}
	resp, err := client.getReplicaHandleResponse(httpResp)
	return resp, err
}

// getReplicaCreateRequest creates the GetReplica request.
func (client *ContainerAppsRevisionReplicasClient) getReplicaCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, replicaName string, options *ContainerAppsRevisionReplicasClientGetReplicaOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/revisions/{revisionName}/replicas/{replicaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if revisionName == "" {
		return nil, errors.New("parameter revisionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{revisionName}", url.PathEscape(revisionName))
	if replicaName == "" {
		return nil, errors.New("parameter replicaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicaName}", url.PathEscape(replicaName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getReplicaHandleResponse handles the GetReplica response.
func (client *ContainerAppsRevisionReplicasClient) getReplicaHandleResponse(resp *http.Response) (ContainerAppsRevisionReplicasClientGetReplicaResponse, error) {
	result := ContainerAppsRevisionReplicasClientGetReplicaResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Replica); err != nil {
		return ContainerAppsRevisionReplicasClientGetReplicaResponse{}, err
	}
	return result, nil
}

// ListReplicas - List replicas for a Container App Revision.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - revisionName - Name of the Container App Revision.
//   - options - ContainerAppsRevisionReplicasClientListReplicasOptions contains the optional parameters for the ContainerAppsRevisionReplicasClient.ListReplicas
//     method.
func (client *ContainerAppsRevisionReplicasClient) ListReplicas(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, options *ContainerAppsRevisionReplicasClientListReplicasOptions) (ContainerAppsRevisionReplicasClientListReplicasResponse, error) {
	var err error
	req, err := client.listReplicasCreateRequest(ctx, resourceGroupName, containerAppName, revisionName, options)
	if err != nil {
		return ContainerAppsRevisionReplicasClientListReplicasResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsRevisionReplicasClientListReplicasResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsRevisionReplicasClientListReplicasResponse{}, err
	}
	resp, err := client.listReplicasHandleResponse(httpResp)
	return resp, err
}

// listReplicasCreateRequest creates the ListReplicas request.
func (client *ContainerAppsRevisionReplicasClient) listReplicasCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, options *ContainerAppsRevisionReplicasClientListReplicasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/revisions/{revisionName}/replicas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if revisionName == "" {
		return nil, errors.New("parameter revisionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{revisionName}", url.PathEscape(revisionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listReplicasHandleResponse handles the ListReplicas response.
func (client *ContainerAppsRevisionReplicasClient) listReplicasHandleResponse(resp *http.Response) (ContainerAppsRevisionReplicasClientListReplicasResponse, error) {
	result := ContainerAppsRevisionReplicasClientListReplicasResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicaCollection); err != nil {
		return ContainerAppsRevisionReplicasClientListReplicasResponse{}, err
	}
	return result, nil
}
