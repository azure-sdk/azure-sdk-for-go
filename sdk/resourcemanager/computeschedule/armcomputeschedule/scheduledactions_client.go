//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcomputeschedule

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

// ScheduledActionsClient contains the methods for the ScheduledActions group.
// Don't use this type directly, use NewScheduledActionsClient() instead.
type ScheduledActionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewScheduledActionsClient creates a new instance of ScheduledActionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScheduledActionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScheduledActionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScheduledActionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// VirtualMachinesCancelOperations - virtualMachinesCancelOperations: cancelOperations for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesCancelOperationsOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesCancelOperations
//     method.
func (client *ScheduledActionsClient) VirtualMachinesCancelOperations(ctx context.Context, locationparameter string, requestBody CancelOperationsRequest, options *ScheduledActionsClientVirtualMachinesCancelOperationsOptions) (ScheduledActionsClientVirtualMachinesCancelOperationsResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesCancelOperations"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesCancelOperationsCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesCancelOperationsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesCancelOperationsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesCancelOperationsResponse{}, err
	}
	resp, err := client.virtualMachinesCancelOperationsHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesCancelOperationsCreateRequest creates the VirtualMachinesCancelOperations request.
func (client *ScheduledActionsClient) virtualMachinesCancelOperationsCreateRequest(ctx context.Context, locationparameter string, requestBody CancelOperationsRequest, options *ScheduledActionsClientVirtualMachinesCancelOperationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesCancelOperations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesCancelOperationsHandleResponse handles the VirtualMachinesCancelOperations response.
func (client *ScheduledActionsClient) virtualMachinesCancelOperationsHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesCancelOperationsResponse, error) {
	result := ScheduledActionsClientVirtualMachinesCancelOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CancelOperationsResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesCancelOperationsResponse{}, err
	}
	return result, nil
}

// VirtualMachinesExecuteDeallocate - virtualMachinesExecuteDeallocate: executeDeallocate for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesExecuteDeallocateOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesExecuteDeallocate
//     method.
func (client *ScheduledActionsClient) VirtualMachinesExecuteDeallocate(ctx context.Context, locationparameter string, requestBody ExecuteDeallocateRequest, options *ScheduledActionsClientVirtualMachinesExecuteDeallocateOptions) (ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesExecuteDeallocate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesExecuteDeallocateCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse{}, err
	}
	resp, err := client.virtualMachinesExecuteDeallocateHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesExecuteDeallocateCreateRequest creates the VirtualMachinesExecuteDeallocate request.
func (client *ScheduledActionsClient) virtualMachinesExecuteDeallocateCreateRequest(ctx context.Context, locationparameter string, requestBody ExecuteDeallocateRequest, options *ScheduledActionsClientVirtualMachinesExecuteDeallocateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesExecuteDeallocate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesExecuteDeallocateHandleResponse handles the VirtualMachinesExecuteDeallocate response.
func (client *ScheduledActionsClient) virtualMachinesExecuteDeallocateHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse, error) {
	result := ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeallocateResourceOperationResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse{}, err
	}
	return result, nil
}

// VirtualMachinesExecuteHibernate - virtualMachinesExecuteHibernate: executeHibernate for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesExecuteHibernateOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesExecuteHibernate
//     method.
func (client *ScheduledActionsClient) VirtualMachinesExecuteHibernate(ctx context.Context, locationparameter string, requestBody ExecuteHibernateRequest, options *ScheduledActionsClientVirtualMachinesExecuteHibernateOptions) (ScheduledActionsClientVirtualMachinesExecuteHibernateResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesExecuteHibernate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesExecuteHibernateCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteHibernateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteHibernateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesExecuteHibernateResponse{}, err
	}
	resp, err := client.virtualMachinesExecuteHibernateHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesExecuteHibernateCreateRequest creates the VirtualMachinesExecuteHibernate request.
func (client *ScheduledActionsClient) virtualMachinesExecuteHibernateCreateRequest(ctx context.Context, locationparameter string, requestBody ExecuteHibernateRequest, options *ScheduledActionsClientVirtualMachinesExecuteHibernateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesExecuteHibernate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesExecuteHibernateHandleResponse handles the VirtualMachinesExecuteHibernate response.
func (client *ScheduledActionsClient) virtualMachinesExecuteHibernateHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesExecuteHibernateResponse, error) {
	result := ScheduledActionsClientVirtualMachinesExecuteHibernateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HibernateResourceOperationResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteHibernateResponse{}, err
	}
	return result, nil
}

// VirtualMachinesExecuteStart - virtualMachinesExecuteStart: executeStart for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesExecuteStartOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesExecuteStart
//     method.
func (client *ScheduledActionsClient) VirtualMachinesExecuteStart(ctx context.Context, locationparameter string, requestBody ExecuteStartRequest, options *ScheduledActionsClientVirtualMachinesExecuteStartOptions) (ScheduledActionsClientVirtualMachinesExecuteStartResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesExecuteStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesExecuteStartCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteStartResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteStartResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesExecuteStartResponse{}, err
	}
	resp, err := client.virtualMachinesExecuteStartHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesExecuteStartCreateRequest creates the VirtualMachinesExecuteStart request.
func (client *ScheduledActionsClient) virtualMachinesExecuteStartCreateRequest(ctx context.Context, locationparameter string, requestBody ExecuteStartRequest, options *ScheduledActionsClientVirtualMachinesExecuteStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesExecuteStart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesExecuteStartHandleResponse handles the VirtualMachinesExecuteStart response.
func (client *ScheduledActionsClient) virtualMachinesExecuteStartHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesExecuteStartResponse, error) {
	result := ScheduledActionsClientVirtualMachinesExecuteStartResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StartResourceOperationResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesExecuteStartResponse{}, err
	}
	return result, nil
}

// VirtualMachinesGetOperationStatus - virtualMachinesGetOperationStatus: getOperationStatus for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesGetOperationStatusOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesGetOperationStatus
//     method.
func (client *ScheduledActionsClient) VirtualMachinesGetOperationStatus(ctx context.Context, locationparameter string, requestBody GetOperationStatusRequest, options *ScheduledActionsClientVirtualMachinesGetOperationStatusOptions) (ScheduledActionsClientVirtualMachinesGetOperationStatusResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesGetOperationStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesGetOperationStatusCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesGetOperationStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesGetOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesGetOperationStatusResponse{}, err
	}
	resp, err := client.virtualMachinesGetOperationStatusHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesGetOperationStatusCreateRequest creates the VirtualMachinesGetOperationStatus request.
func (client *ScheduledActionsClient) virtualMachinesGetOperationStatusCreateRequest(ctx context.Context, locationparameter string, requestBody GetOperationStatusRequest, options *ScheduledActionsClientVirtualMachinesGetOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesGetOperationStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesGetOperationStatusHandleResponse handles the VirtualMachinesGetOperationStatus response.
func (client *ScheduledActionsClient) virtualMachinesGetOperationStatusHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesGetOperationStatusResponse, error) {
	result := ScheduledActionsClientVirtualMachinesGetOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetOperationStatusResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesGetOperationStatusResponse{}, err
	}
	return result, nil
}

// VirtualMachinesSubmitDeallocate - virtualMachinesSubmitDeallocate: submitDeallocate for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesSubmitDeallocateOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesSubmitDeallocate
//     method.
func (client *ScheduledActionsClient) VirtualMachinesSubmitDeallocate(ctx context.Context, locationparameter string, requestBody SubmitDeallocateRequest, options *ScheduledActionsClientVirtualMachinesSubmitDeallocateOptions) (ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesSubmitDeallocate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesSubmitDeallocateCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse{}, err
	}
	resp, err := client.virtualMachinesSubmitDeallocateHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesSubmitDeallocateCreateRequest creates the VirtualMachinesSubmitDeallocate request.
func (client *ScheduledActionsClient) virtualMachinesSubmitDeallocateCreateRequest(ctx context.Context, locationparameter string, requestBody SubmitDeallocateRequest, options *ScheduledActionsClientVirtualMachinesSubmitDeallocateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesSubmitDeallocate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesSubmitDeallocateHandleResponse handles the VirtualMachinesSubmitDeallocate response.
func (client *ScheduledActionsClient) virtualMachinesSubmitDeallocateHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse, error) {
	result := ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeallocateResourceOperationResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse{}, err
	}
	return result, nil
}

// VirtualMachinesSubmitHibernate - virtualMachinesSubmitHibernate: submitHibernate for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesSubmitHibernateOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesSubmitHibernate
//     method.
func (client *ScheduledActionsClient) VirtualMachinesSubmitHibernate(ctx context.Context, locationparameter string, requestBody SubmitHibernateRequest, options *ScheduledActionsClientVirtualMachinesSubmitHibernateOptions) (ScheduledActionsClientVirtualMachinesSubmitHibernateResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesSubmitHibernate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesSubmitHibernateCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitHibernateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitHibernateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesSubmitHibernateResponse{}, err
	}
	resp, err := client.virtualMachinesSubmitHibernateHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesSubmitHibernateCreateRequest creates the VirtualMachinesSubmitHibernate request.
func (client *ScheduledActionsClient) virtualMachinesSubmitHibernateCreateRequest(ctx context.Context, locationparameter string, requestBody SubmitHibernateRequest, options *ScheduledActionsClientVirtualMachinesSubmitHibernateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesSubmitHibernate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesSubmitHibernateHandleResponse handles the VirtualMachinesSubmitHibernate response.
func (client *ScheduledActionsClient) virtualMachinesSubmitHibernateHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesSubmitHibernateResponse, error) {
	result := ScheduledActionsClientVirtualMachinesSubmitHibernateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HibernateResourceOperationResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitHibernateResponse{}, err
	}
	return result, nil
}

// VirtualMachinesSubmitStart - virtualMachinesSubmitStart: submitStart for a virtual machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - locationparameter - The location name.
//   - requestBody - The request body
//   - options - ScheduledActionsClientVirtualMachinesSubmitStartOptions contains the optional parameters for the ScheduledActionsClient.VirtualMachinesSubmitStart
//     method.
func (client *ScheduledActionsClient) VirtualMachinesSubmitStart(ctx context.Context, locationparameter string, requestBody SubmitStartRequest, options *ScheduledActionsClientVirtualMachinesSubmitStartOptions) (ScheduledActionsClientVirtualMachinesSubmitStartResponse, error) {
	var err error
	const operationName = "ScheduledActionsClient.VirtualMachinesSubmitStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.virtualMachinesSubmitStartCreateRequest(ctx, locationparameter, requestBody, options)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitStartResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitStartResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledActionsClientVirtualMachinesSubmitStartResponse{}, err
	}
	resp, err := client.virtualMachinesSubmitStartHandleResponse(httpResp)
	return resp, err
}

// virtualMachinesSubmitStartCreateRequest creates the VirtualMachinesSubmitStart request.
func (client *ScheduledActionsClient) virtualMachinesSubmitStartCreateRequest(ctx context.Context, locationparameter string, requestBody SubmitStartRequest, options *ScheduledActionsClientVirtualMachinesSubmitStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ComputeSchedule/locations/{locationparameter}/virtualMachinesSubmitStart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationparameter == "" {
		return nil, errors.New("parameter locationparameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationparameter}", url.PathEscape(locationparameter))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// virtualMachinesSubmitStartHandleResponse handles the VirtualMachinesSubmitStart response.
func (client *ScheduledActionsClient) virtualMachinesSubmitStartHandleResponse(resp *http.Response) (ScheduledActionsClientVirtualMachinesSubmitStartResponse, error) {
	result := ScheduledActionsClientVirtualMachinesSubmitStartResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StartResourceOperationResponse); err != nil {
		return ScheduledActionsClientVirtualMachinesSubmitStartResponse{}, err
	}
	return result, nil
}
