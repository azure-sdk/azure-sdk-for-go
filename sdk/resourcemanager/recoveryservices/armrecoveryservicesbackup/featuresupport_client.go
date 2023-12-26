//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// FeatureSupportClient contains the methods for the FeatureSupport group.
// Don't use this type directly, use NewFeatureSupportClient() instead.
type FeatureSupportClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFeatureSupportClient creates a new instance of FeatureSupportClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFeatureSupportClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FeatureSupportClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FeatureSupportClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Validate - It will validate if given feature with resource properties is supported in service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - azureRegion - Azure region to hit Api
//   - parameters - Feature support request object
//   - options - FeatureSupportClientValidateOptions contains the optional parameters for the FeatureSupportClient.Validate method.
func (client *FeatureSupportClient) Validate(ctx context.Context, azureRegion string, parameters FeatureSupportRequestClassification, options *FeatureSupportClientValidateOptions) (FeatureSupportClientValidateResponse, error) {
	var err error
	const operationName = "FeatureSupportClient.Validate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateCreateRequest(ctx, azureRegion, parameters, options)
	if err != nil {
		return FeatureSupportClientValidateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FeatureSupportClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FeatureSupportClientValidateResponse{}, err
	}
	resp, err := client.validateHandleResponse(httpResp)
	return resp, err
}

// validateCreateRequest creates the Validate request.
func (client *FeatureSupportClient) validateCreateRequest(ctx context.Context, azureRegion string, parameters FeatureSupportRequestClassification, options *FeatureSupportClientValidateOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/locations/{azureRegion}/backupValidateFeatures"
	if azureRegion == "" {
		return nil, errors.New("parameter azureRegion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureRegion}", url.PathEscape(azureRegion))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// validateHandleResponse handles the Validate response.
func (client *FeatureSupportClient) validateHandleResponse(resp *http.Response) (FeatureSupportClientValidateResponse, error) {
	result := FeatureSupportClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureVMResourceFeatureSupportResponse); err != nil {
		return FeatureSupportClientValidateResponse{}, err
	}
	return result, nil
}
