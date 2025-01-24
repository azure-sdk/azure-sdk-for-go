//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice

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

// QnAMakerEndpointKeysClient contains the methods for the QnAMakerEndpointKeys group.
// Don't use this type directly, use NewQnAMakerEndpointKeysClient() instead.
type QnAMakerEndpointKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQnAMakerEndpointKeysClient creates a new instance of QnAMakerEndpointKeysClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQnAMakerEndpointKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QnAMakerEndpointKeysClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QnAMakerEndpointKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Lists the QnA Maker endpoint keys
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - parameters - The request body parameters to provide for the check name availability request
//   - options - QnAMakerEndpointKeysClientGetOptions contains the optional parameters for the QnAMakerEndpointKeysClient.Get
//     method.
func (client *QnAMakerEndpointKeysClient) Get(ctx context.Context, parameters QnAMakerEndpointKeysRequestBody, options *QnAMakerEndpointKeysClientGetOptions) (QnAMakerEndpointKeysClientGetResponse, error) {
	var err error
	const operationName = "QnAMakerEndpointKeysClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, parameters, options)
	if err != nil {
		return QnAMakerEndpointKeysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QnAMakerEndpointKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QnAMakerEndpointKeysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *QnAMakerEndpointKeysClient) getCreateRequest(ctx context.Context, parameters QnAMakerEndpointKeysRequestBody, options *QnAMakerEndpointKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BotService/listQnAMakerEndpointKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QnAMakerEndpointKeysClient) getHandleResponse(resp *http.Response) (QnAMakerEndpointKeysClientGetResponse, error) {
	result := QnAMakerEndpointKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QnAMakerEndpointKeysResponse); err != nil {
		return QnAMakerEndpointKeysClientGetResponse{}, err
	}
	return result, nil
}
