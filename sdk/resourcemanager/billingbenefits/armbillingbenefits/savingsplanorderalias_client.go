// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbillingbenefits

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

// SavingsPlanOrderAliasClient contains the methods for the SavingsPlanOrderAlias group.
// Don't use this type directly, use NewSavingsPlanOrderAliasClient() instead.
type SavingsPlanOrderAliasClient struct {
	internal *arm.Client
}

// NewSavingsPlanOrderAliasClient creates a new instance of SavingsPlanOrderAliasClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSavingsPlanOrderAliasClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SavingsPlanOrderAliasClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SavingsPlanOrderAliasClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreate - Create a savings plan. Learn more about permissions needed at https://go.microsoft.com/fwlink/?linkid=2215851
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - savingsPlanOrderAliasName - Name of the savings plan order alias
//   - body - Request body for creating a savings plan order alias
//   - options - SavingsPlanOrderAliasClientBeginCreateOptions contains the optional parameters for the SavingsPlanOrderAliasClient.BeginCreate
//     method.
func (client *SavingsPlanOrderAliasClient) BeginCreate(ctx context.Context, savingsPlanOrderAliasName string, body SavingsPlanOrderAliasModel, options *SavingsPlanOrderAliasClientBeginCreateOptions) (*runtime.Poller[SavingsPlanOrderAliasClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, savingsPlanOrderAliasName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SavingsPlanOrderAliasClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SavingsPlanOrderAliasClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a savings plan. Learn more about permissions needed at https://go.microsoft.com/fwlink/?linkid=2215851
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
func (client *SavingsPlanOrderAliasClient) create(ctx context.Context, savingsPlanOrderAliasName string, body SavingsPlanOrderAliasModel, options *SavingsPlanOrderAliasClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SavingsPlanOrderAliasClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, savingsPlanOrderAliasName, body, options)
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
func (client *SavingsPlanOrderAliasClient) createCreateRequest(ctx context.Context, savingsPlanOrderAliasName string, body SavingsPlanOrderAliasModel, _ *SavingsPlanOrderAliasClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrderAliases/{savingsPlanOrderAliasName}"
	if savingsPlanOrderAliasName == "" {
		return nil, errors.New("parameter savingsPlanOrderAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderAliasName}", url.PathEscape(savingsPlanOrderAliasName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a savings plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - savingsPlanOrderAliasName - Name of the savings plan order alias
//   - options - SavingsPlanOrderAliasClientGetOptions contains the optional parameters for the SavingsPlanOrderAliasClient.Get
//     method.
func (client *SavingsPlanOrderAliasClient) Get(ctx context.Context, savingsPlanOrderAliasName string, options *SavingsPlanOrderAliasClientGetOptions) (SavingsPlanOrderAliasClientGetResponse, error) {
	var err error
	const operationName = "SavingsPlanOrderAliasClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, savingsPlanOrderAliasName, options)
	if err != nil {
		return SavingsPlanOrderAliasClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SavingsPlanOrderAliasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SavingsPlanOrderAliasClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SavingsPlanOrderAliasClient) getCreateRequest(ctx context.Context, savingsPlanOrderAliasName string, _ *SavingsPlanOrderAliasClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrderAliases/{savingsPlanOrderAliasName}"
	if savingsPlanOrderAliasName == "" {
		return nil, errors.New("parameter savingsPlanOrderAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderAliasName}", url.PathEscape(savingsPlanOrderAliasName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SavingsPlanOrderAliasClient) getHandleResponse(resp *http.Response) (SavingsPlanOrderAliasClientGetResponse, error) {
	result := SavingsPlanOrderAliasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanOrderAliasModel); err != nil {
		return SavingsPlanOrderAliasClientGetResponse{}, err
	}
	return result, nil
}
