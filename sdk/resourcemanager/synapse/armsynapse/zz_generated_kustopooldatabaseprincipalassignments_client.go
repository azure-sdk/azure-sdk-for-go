//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// KustoPoolDatabasePrincipalAssignmentsClient contains the methods for the KustoPoolDatabasePrincipalAssignments group.
// Don't use this type directly, use NewKustoPoolDatabasePrincipalAssignmentsClient() instead.
type KustoPoolDatabasePrincipalAssignmentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewKustoPoolDatabasePrincipalAssignmentsClient creates a new instance of KustoPoolDatabasePrincipalAssignmentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewKustoPoolDatabasePrincipalAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KustoPoolDatabasePrincipalAssignmentsClient, error) {
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
	client := &KustoPoolDatabasePrincipalAssignmentsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Checks that the database principal assignment is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// workspaceName - The name of the workspace.
// kustoPoolName - The name of the Kusto pool.
// databaseName - The name of the database in the Kusto pool.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// principalAssignmentName - The name of the resource.
// options - KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityOptions contains the optional parameters for
// the KustoPoolDatabasePrincipalAssignmentsClient.CheckNameAvailability method.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) CheckNameAvailability(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, resourceGroupName string, principalAssignmentName DatabasePrincipalAssignmentCheckNameRequest, options *KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityOptions) (KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, workspaceName, kustoPoolName, databaseName, resourceGroupName, principalAssignmentName, options)
	if err != nil {
		return KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) checkNameAvailabilityCreateRequest(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, resourceGroupName string, principalAssignmentName DatabasePrincipalAssignmentCheckNameRequest, options *KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/checkPrincipalAssignmentNameAvailability"
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, principalAssignmentName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityResponse, error) {
	result := KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return KustoPoolDatabasePrincipalAssignmentsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates a Kusto pool database principalAssignment.
// If the operation fails it returns an *azcore.ResponseError type.
// workspaceName - The name of the workspace.
// kustoPoolName - The name of the Kusto pool.
// databaseName - The name of the database in the Kusto pool.
// principalAssignmentName - The name of the Kusto principalAssignment.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// parameters - The Kusto principalAssignments parameters supplied for the operation.
// options - KustoPoolDatabasePrincipalAssignmentsClientBeginCreateOrUpdateOptions contains the optional parameters for the
// KustoPoolDatabasePrincipalAssignmentsClient.BeginCreateOrUpdate method.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) BeginCreateOrUpdate(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, parameters DatabasePrincipalAssignment, options *KustoPoolDatabasePrincipalAssignmentsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[KustoPoolDatabasePrincipalAssignmentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, workspaceName, kustoPoolName, databaseName, principalAssignmentName, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[KustoPoolDatabasePrincipalAssignmentsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[KustoPoolDatabasePrincipalAssignmentsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a Kusto pool database principalAssignment.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) createOrUpdate(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, parameters DatabasePrincipalAssignment, options *KustoPoolDatabasePrincipalAssignmentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, workspaceName, kustoPoolName, databaseName, principalAssignmentName, resourceGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, parameters DatabasePrincipalAssignment, options *KustoPoolDatabasePrincipalAssignmentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/principalAssignments/{principalAssignmentName}"
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if principalAssignmentName == "" {
		return nil, errors.New("parameter principalAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{principalAssignmentName}", url.PathEscape(principalAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a Kusto pool principalAssignment.
// If the operation fails it returns an *azcore.ResponseError type.
// workspaceName - The name of the workspace.
// kustoPoolName - The name of the Kusto pool.
// databaseName - The name of the database in the Kusto pool.
// principalAssignmentName - The name of the Kusto principalAssignment.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - KustoPoolDatabasePrincipalAssignmentsClientBeginDeleteOptions contains the optional parameters for the KustoPoolDatabasePrincipalAssignmentsClient.BeginDelete
// method.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) BeginDelete(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, options *KustoPoolDatabasePrincipalAssignmentsClientBeginDeleteOptions) (*armruntime.Poller[KustoPoolDatabasePrincipalAssignmentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, workspaceName, kustoPoolName, databaseName, principalAssignmentName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[KustoPoolDatabasePrincipalAssignmentsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[KustoPoolDatabasePrincipalAssignmentsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Kusto pool principalAssignment.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) deleteOperation(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, options *KustoPoolDatabasePrincipalAssignmentsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, workspaceName, kustoPoolName, databaseName, principalAssignmentName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) deleteCreateRequest(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, options *KustoPoolDatabasePrincipalAssignmentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/principalAssignments/{principalAssignmentName}"
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if principalAssignmentName == "" {
		return nil, errors.New("parameter principalAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{principalAssignmentName}", url.PathEscape(principalAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a Kusto pool database principalAssignment.
// If the operation fails it returns an *azcore.ResponseError type.
// workspaceName - The name of the workspace.
// kustoPoolName - The name of the Kusto pool.
// databaseName - The name of the database in the Kusto pool.
// principalAssignmentName - The name of the Kusto principalAssignment.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - KustoPoolDatabasePrincipalAssignmentsClientGetOptions contains the optional parameters for the KustoPoolDatabasePrincipalAssignmentsClient.Get
// method.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) Get(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, options *KustoPoolDatabasePrincipalAssignmentsClientGetOptions) (KustoPoolDatabasePrincipalAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, workspaceName, kustoPoolName, databaseName, principalAssignmentName, resourceGroupName, options)
	if err != nil {
		return KustoPoolDatabasePrincipalAssignmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KustoPoolDatabasePrincipalAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDatabasePrincipalAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) getCreateRequest(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, principalAssignmentName string, resourceGroupName string, options *KustoPoolDatabasePrincipalAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/principalAssignments/{principalAssignmentName}"
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if principalAssignmentName == "" {
		return nil, errors.New("parameter principalAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{principalAssignmentName}", url.PathEscape(principalAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) getHandleResponse(resp *http.Response) (KustoPoolDatabasePrincipalAssignmentsClientGetResponse, error) {
	result := KustoPoolDatabasePrincipalAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalAssignment); err != nil {
		return KustoPoolDatabasePrincipalAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all Kusto pool database principalAssignments.
// If the operation fails it returns an *azcore.ResponseError type.
// workspaceName - The name of the workspace.
// kustoPoolName - The name of the Kusto pool.
// databaseName - The name of the database in the Kusto pool.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - KustoPoolDatabasePrincipalAssignmentsClientListOptions contains the optional parameters for the KustoPoolDatabasePrincipalAssignmentsClient.List
// method.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) List(workspaceName string, kustoPoolName string, databaseName string, resourceGroupName string, options *KustoPoolDatabasePrincipalAssignmentsClientListOptions) *runtime.Pager[KustoPoolDatabasePrincipalAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[KustoPoolDatabasePrincipalAssignmentsClientListResponse]{
		More: func(page KustoPoolDatabasePrincipalAssignmentsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *KustoPoolDatabasePrincipalAssignmentsClientListResponse) (KustoPoolDatabasePrincipalAssignmentsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, workspaceName, kustoPoolName, databaseName, resourceGroupName, options)
			if err != nil {
				return KustoPoolDatabasePrincipalAssignmentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return KustoPoolDatabasePrincipalAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return KustoPoolDatabasePrincipalAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) listCreateRequest(ctx context.Context, workspaceName string, kustoPoolName string, databaseName string, resourceGroupName string, options *KustoPoolDatabasePrincipalAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/principalAssignments"
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *KustoPoolDatabasePrincipalAssignmentsClient) listHandleResponse(resp *http.Response) (KustoPoolDatabasePrincipalAssignmentsClientListResponse, error) {
	result := KustoPoolDatabasePrincipalAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalAssignmentListResult); err != nil {
		return KustoPoolDatabasePrincipalAssignmentsClientListResponse{}, err
	}
	return result, nil
}
