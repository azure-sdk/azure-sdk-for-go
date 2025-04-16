// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

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

// MicrosoftStorageSyncClient contains the methods for the MicrosoftStorageSync group.
// Don't use this type directly, use NewMicrosoftStorageSyncClient() instead.
type MicrosoftStorageSyncClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMicrosoftStorageSyncClient creates a new instance of MicrosoftStorageSyncClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMicrosoftStorageSyncClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MicrosoftStorageSyncClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MicrosoftStorageSyncClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// LocationOperationStatus - Get Operation status
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - locationName - The desired region to obtain information from.
//   - operationID - operation Id
//   - options - MicrosoftStorageSyncClientLocationOperationStatusOptions contains the optional parameters for the MicrosoftStorageSyncClient.LocationOperationStatus
//     method.
func (client *MicrosoftStorageSyncClient) LocationOperationStatus(ctx context.Context, locationName string, operationID string, options *MicrosoftStorageSyncClientLocationOperationStatusOptions) (MicrosoftStorageSyncClientLocationOperationStatusResponse, error) {
	var err error
	const operationName = "MicrosoftStorageSyncClient.LocationOperationStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.locationOperationStatusCreateRequest(ctx, locationName, operationID, options)
	if err != nil {
		return MicrosoftStorageSyncClientLocationOperationStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MicrosoftStorageSyncClientLocationOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MicrosoftStorageSyncClientLocationOperationStatusResponse{}, err
	}
	resp, err := client.locationOperationStatusHandleResponse(httpResp)
	return resp, err
}

// locationOperationStatusCreateRequest creates the LocationOperationStatus request.
func (client *MicrosoftStorageSyncClient) locationOperationStatusCreateRequest(ctx context.Context, locationName string, operationID string, _ *MicrosoftStorageSyncClientLocationOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/locations/{locationName}/operations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// locationOperationStatusHandleResponse handles the LocationOperationStatus response.
func (client *MicrosoftStorageSyncClient) locationOperationStatusHandleResponse(resp *http.Response) (MicrosoftStorageSyncClientLocationOperationStatusResponse, error) {
	result := MicrosoftStorageSyncClientLocationOperationStatusResponse{}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationOperationStatus); err != nil {
		return MicrosoftStorageSyncClientLocationOperationStatusResponse{}, err
	}
	return result, nil
}
