// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

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

// SimsClient contains the methods for the Sims group.
// Don't use this type directly, use NewSimsClient() instead.
type SimsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSimsClient creates a new instance of SimsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSimsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SimsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SimsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginBulkDelete - Bulk delete SIMs from a SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - parameters - Parameters supplied to the bulk SIM delete operation.
//   - options - SimsClientBeginBulkDeleteOptions contains the optional parameters for the SimsClient.BeginBulkDelete method.
func (client *SimsClient) BeginBulkDelete(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimDeleteList, options *SimsClientBeginBulkDeleteOptions) (*runtime.Poller[SimsClientBulkDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.bulkDelete(ctx, resourceGroupName, simGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimsClientBulkDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimsClientBulkDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// BulkDelete - Bulk delete SIMs from a SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SimsClient) bulkDelete(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimDeleteList, options *SimsClientBeginBulkDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SimsClient.BeginBulkDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.bulkDeleteCreateRequest(ctx, resourceGroupName, simGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// bulkDeleteCreateRequest creates the BulkDelete request.
func (client *SimsClient) bulkDeleteCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimDeleteList, _ *SimsClientBeginBulkDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/deleteSims"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginBulkUpload - Bulk upload SIMs to a SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - parameters - Parameters supplied to the bulk SIM upload operation.
//   - options - SimsClientBeginBulkUploadOptions contains the optional parameters for the SimsClient.BeginBulkUpload method.
func (client *SimsClient) BeginBulkUpload(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimUploadList, options *SimsClientBeginBulkUploadOptions) (*runtime.Poller[SimsClientBulkUploadResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.bulkUpload(ctx, resourceGroupName, simGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimsClientBulkUploadResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimsClientBulkUploadResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// BulkUpload - Bulk upload SIMs to a SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SimsClient) bulkUpload(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimUploadList, options *SimsClientBeginBulkUploadOptions) (*http.Response, error) {
	var err error
	const operationName = "SimsClient.BeginBulkUpload"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.bulkUploadCreateRequest(ctx, resourceGroupName, simGroupName, parameters, options)
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

// bulkUploadCreateRequest creates the BulkUpload request.
func (client *SimsClient) bulkUploadCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimUploadList, _ *SimsClientBeginBulkUploadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/uploadSims"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginBulkUploadEncrypted - Bulk upload SIMs in encrypted form to a SIM group. The SIM credentials must be encrypted.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - parameters - Parameters supplied to the encrypted SIMs upload operation.
//   - options - SimsClientBeginBulkUploadEncryptedOptions contains the optional parameters for the SimsClient.BeginBulkUploadEncrypted
//     method.
func (client *SimsClient) BeginBulkUploadEncrypted(ctx context.Context, resourceGroupName string, simGroupName string, parameters EncryptedSimUploadList, options *SimsClientBeginBulkUploadEncryptedOptions) (*runtime.Poller[SimsClientBulkUploadEncryptedResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.bulkUploadEncrypted(ctx, resourceGroupName, simGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimsClientBulkUploadEncryptedResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimsClientBulkUploadEncryptedResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// BulkUploadEncrypted - Bulk upload SIMs in encrypted form to a SIM group. The SIM credentials must be encrypted.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SimsClient) bulkUploadEncrypted(ctx context.Context, resourceGroupName string, simGroupName string, parameters EncryptedSimUploadList, options *SimsClientBeginBulkUploadEncryptedOptions) (*http.Response, error) {
	var err error
	const operationName = "SimsClient.BeginBulkUploadEncrypted"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.bulkUploadEncryptedCreateRequest(ctx, resourceGroupName, simGroupName, parameters, options)
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

// bulkUploadEncryptedCreateRequest creates the BulkUploadEncrypted request.
func (client *SimsClient) bulkUploadEncryptedCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, parameters EncryptedSimUploadList, _ *SimsClientBeginBulkUploadEncryptedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/uploadEncryptedSims"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginClone - Clone SIMs to another SIM Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - parameters - Parameters supplied to clone the SIMs.
//   - options - SimsClientBeginCloneOptions contains the optional parameters for the SimsClient.BeginClone method.
func (client *SimsClient) BeginClone(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimClone, options *SimsClientBeginCloneOptions) (*runtime.Poller[SimsClientCloneResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.clone(ctx, resourceGroupName, simGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimsClientCloneResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimsClientCloneResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Clone - Clone SIMs to another SIM Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SimsClient) clone(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimClone, options *SimsClientBeginCloneOptions) (*http.Response, error) {
	var err error
	const operationName = "SimsClient.BeginClone"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cloneCreateRequest(ctx, resourceGroupName, simGroupName, parameters, options)
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

// cloneCreateRequest creates the Clone request.
func (client *SimsClient) cloneCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimClone, _ *SimsClientBeginCloneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/cloneSims"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - Creates or updates a SIM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - simName - The name of the SIM.
//   - parameters - Parameters supplied to the create or update SIM operation.
//   - options - SimsClientBeginCreateOrUpdateOptions contains the optional parameters for the SimsClient.BeginCreateOrUpdate
//     method.
func (client *SimsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, simGroupName string, simName string, parameters Sim, options *SimsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SimsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, simGroupName, simName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a SIM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SimsClient) createOrUpdate(ctx context.Context, resourceGroupName string, simGroupName string, simName string, parameters Sim, options *SimsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SimsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, simGroupName, simName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SimsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, simName string, parameters Sim, _ *SimsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/sims/{simName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	if simName == "" {
		return nil, errors.New("parameter simName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simName}", url.PathEscape(simName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified SIM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - simName - The name of the SIM.
//   - options - SimsClientBeginDeleteOptions contains the optional parameters for the SimsClient.BeginDelete method.
func (client *SimsClient) BeginDelete(ctx context.Context, resourceGroupName string, simGroupName string, simName string, options *SimsClientBeginDeleteOptions) (*runtime.Poller[SimsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, simGroupName, simName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified SIM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SimsClient) deleteOperation(ctx context.Context, resourceGroupName string, simGroupName string, simName string, options *SimsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SimsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, simGroupName, simName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SimsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, simName string, _ *SimsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/sims/{simName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	if simName == "" {
		return nil, errors.New("parameter simName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simName}", url.PathEscape(simName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified SIM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - simName - The name of the SIM.
//   - options - SimsClientGetOptions contains the optional parameters for the SimsClient.Get method.
func (client *SimsClient) Get(ctx context.Context, resourceGroupName string, simGroupName string, simName string, options *SimsClientGetOptions) (SimsClientGetResponse, error) {
	var err error
	const operationName = "SimsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, simGroupName, simName, options)
	if err != nil {
		return SimsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SimsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SimsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SimsClient) getCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, simName string, _ *SimsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/sims/{simName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	if simName == "" {
		return nil, errors.New("parameter simName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simName}", url.PathEscape(simName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SimsClient) getHandleResponse(resp *http.Response) (SimsClientGetResponse, error) {
	result := SimsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Sim); err != nil {
		return SimsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByGroupPager - Gets all the SIMs in a SIM group.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - options - SimsClientListByGroupOptions contains the optional parameters for the SimsClient.NewListByGroupPager method.
func (client *SimsClient) NewListByGroupPager(resourceGroupName string, simGroupName string, options *SimsClientListByGroupOptions) *runtime.Pager[SimsClientListByGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SimsClientListByGroupResponse]{
		More: func(page SimsClientListByGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SimsClientListByGroupResponse) (SimsClientListByGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SimsClient.NewListByGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByGroupCreateRequest(ctx, resourceGroupName, simGroupName, options)
			}, nil)
			if err != nil {
				return SimsClientListByGroupResponse{}, err
			}
			return client.listByGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByGroupCreateRequest creates the ListByGroup request.
func (client *SimsClient) listByGroupCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, _ *SimsClientListByGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/sims"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByGroupHandleResponse handles the ListByGroup response.
func (client *SimsClient) listByGroupHandleResponse(resp *http.Response) (SimsClientListByGroupResponse, error) {
	result := SimsClientListByGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimListResult); err != nil {
		return SimsClientListByGroupResponse{}, err
	}
	return result, nil
}

// BeginMove - Move SIMs to another SIM Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - parameters - Parameters supplied to move the SIMs.
//   - options - SimsClientBeginMoveOptions contains the optional parameters for the SimsClient.BeginMove method.
func (client *SimsClient) BeginMove(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimMove, options *SimsClientBeginMoveOptions) (*runtime.Poller[SimsClientMoveResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.move(ctx, resourceGroupName, simGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimsClientMoveResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SimsClientMoveResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Move - Move SIMs to another SIM Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SimsClient) move(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimMove, options *SimsClientBeginMoveOptions) (*http.Response, error) {
	var err error
	const operationName = "SimsClient.BeginMove"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.moveCreateRequest(ctx, resourceGroupName, simGroupName, parameters, options)
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

// moveCreateRequest creates the Move request.
func (client *SimsClient) moveCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimMove, _ *SimsClientBeginMoveOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}/moveSims"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
