//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsapvirtualinstance

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

// SAPApplicationServerInstancesClient contains the methods for the SAPApplicationServerInstances group.
// Don't use this type directly, use NewSAPApplicationServerInstancesClient() instead.
type SAPApplicationServerInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSAPApplicationServerInstancesClient creates a new instance of SAPApplicationServerInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSAPApplicationServerInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SAPApplicationServerInstancesClient, error) {
	cl, err := arm.NewClient(moduleName+".SAPApplicationServerInstancesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SAPApplicationServerInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Puts the SAP Application Server Instance resource.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - applicationInstanceName - The name of SAP Application Server instance resource.
//   - body - The SAP Application Server Instance resource request body.
//   - options - SAPApplicationServerInstancesClientBeginCreateOptions contains the optional parameters for the SAPApplicationServerInstancesClient.BeginCreate
//     method.
func (client *SAPApplicationServerInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body SAPApplicationServerInstance, options *SAPApplicationServerInstancesClientBeginCreateOptions) (*runtime.Poller[SAPApplicationServerInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SAPApplicationServerInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SAPApplicationServerInstancesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Puts the SAP Application Server Instance resource.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *SAPApplicationServerInstancesClient) create(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body SAPApplicationServerInstance, options *SAPApplicationServerInstancesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *SAPApplicationServerInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body SAPApplicationServerInstance, options *SAPApplicationServerInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the SAP Application Server Instance resource.
// This operation will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - applicationInstanceName - The name of SAP Application Server instance resource.
//   - options - SAPApplicationServerInstancesClientBeginDeleteOptions contains the optional parameters for the SAPApplicationServerInstancesClient.BeginDelete
//     method.
func (client *SAPApplicationServerInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginDeleteOptions) (*runtime.Poller[SAPApplicationServerInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SAPApplicationServerInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SAPApplicationServerInstancesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the SAP Application Server Instance resource.
// This operation will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *SAPApplicationServerInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
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
func (client *SAPApplicationServerInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the SAP Application Server Instance corresponding to the Virtual Instance for SAP solutions resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - applicationInstanceName - The name of SAP Application Server instance resource.
//   - options - SAPApplicationServerInstancesClientGetOptions contains the optional parameters for the SAPApplicationServerInstancesClient.Get
//     method.
func (client *SAPApplicationServerInstancesClient) Get(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientGetOptions) (SAPApplicationServerInstancesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
	if err != nil {
		return SAPApplicationServerInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SAPApplicationServerInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SAPApplicationServerInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SAPApplicationServerInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SAPApplicationServerInstancesClient) getHandleResponse(resp *http.Response) (SAPApplicationServerInstancesClientGetResponse, error) {
	result := SAPApplicationServerInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPApplicationServerInstance); err != nil {
		return SAPApplicationServerInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the SAP Application Server Instance resources for a given Virtual Instance for SAP solutions resource.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - options - SAPApplicationServerInstancesClientListOptions contains the optional parameters for the SAPApplicationServerInstancesClient.NewListPager
//     method.
func (client *SAPApplicationServerInstancesClient) NewListPager(resourceGroupName string, sapVirtualInstanceName string, options *SAPApplicationServerInstancesClientListOptions) *runtime.Pager[SAPApplicationServerInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SAPApplicationServerInstancesClientListResponse]{
		More: func(page SAPApplicationServerInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SAPApplicationServerInstancesClientListResponse) (SAPApplicationServerInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SAPApplicationServerInstancesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SAPApplicationServerInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SAPApplicationServerInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SAPApplicationServerInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPApplicationServerInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SAPApplicationServerInstancesClient) listHandleResponse(resp *http.Response) (SAPApplicationServerInstancesClientListResponse, error) {
	result := SAPApplicationServerInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPApplicationServerInstanceList); err != nil {
		return SAPApplicationServerInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginStartInstance - Starts the SAP Application Server Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - applicationInstanceName - The name of SAP Application Server instance resource.
//   - options - SAPApplicationServerInstancesClientBeginStartInstanceOptions contains the optional parameters for the SAPApplicationServerInstancesClient.BeginStartInstance
//     method.
func (client *SAPApplicationServerInstancesClient) BeginStartInstance(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginStartInstanceOptions) (*runtime.Poller[SAPApplicationServerInstancesClientStartInstanceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.startInstance(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SAPApplicationServerInstancesClientStartInstanceResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SAPApplicationServerInstancesClientStartInstanceResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// StartInstance - Starts the SAP Application Server Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *SAPApplicationServerInstancesClient) startInstance(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginStartInstanceOptions) (*http.Response, error) {
	var err error
	req, err := client.startInstanceCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startInstanceCreateRequest creates the StartInstance request.
func (client *SAPApplicationServerInstancesClient) startInstanceCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginStartInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
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

// BeginStopInstance - Stops the SAP Application Server Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - applicationInstanceName - The name of SAP Application Server instance resource.
//   - options - SAPApplicationServerInstancesClientBeginStopInstanceOptions contains the optional parameters for the SAPApplicationServerInstancesClient.BeginStopInstance
//     method.
func (client *SAPApplicationServerInstancesClient) BeginStopInstance(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginStopInstanceOptions) (*runtime.Poller[SAPApplicationServerInstancesClientStopInstanceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stopInstance(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SAPApplicationServerInstancesClientStopInstanceResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SAPApplicationServerInstancesClientStopInstanceResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// StopInstance - Stops the SAP Application Server Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *SAPApplicationServerInstancesClient) stopInstance(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginStopInstanceOptions) (*http.Response, error) {
	var err error
	req, err := client.stopInstanceCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// stopInstanceCreateRequest creates the StopInstance request.
func (client *SAPApplicationServerInstancesClient) stopInstanceCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginStopInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
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

// Update - Updates the SAP Application server instance resource. This can be used to update tags on the resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - applicationInstanceName - The name of SAP Application Server instance resource.
//   - body - The SAP Application Server Instance resource request body.
//   - options - SAPApplicationServerInstancesClientUpdateOptions contains the optional parameters for the SAPApplicationServerInstancesClient.Update
//     method.
func (client *SAPApplicationServerInstancesClient) Update(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body UpdateSAPApplicationInstanceRequest, options *SAPApplicationServerInstancesClientUpdateOptions) (SAPApplicationServerInstancesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, body, options)
	if err != nil {
		return SAPApplicationServerInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SAPApplicationServerInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SAPApplicationServerInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SAPApplicationServerInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body UpdateSAPApplicationInstanceRequest, options *SAPApplicationServerInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SAPApplicationServerInstancesClient) updateHandleResponse(resp *http.Response) (SAPApplicationServerInstancesClientUpdateResponse, error) {
	result := SAPApplicationServerInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPApplicationServerInstance); err != nil {
		return SAPApplicationServerInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
