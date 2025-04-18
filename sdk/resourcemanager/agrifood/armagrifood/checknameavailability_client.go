// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

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

// CheckNameAvailabilityClient contains the methods for the CheckNameAvailability group.
// Don't use this type directly, use NewCheckNameAvailabilityClient() instead.
type CheckNameAvailabilityClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCheckNameAvailabilityClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CheckNameAvailabilityClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CheckNameAvailabilityClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks the name availability of the resource with requested resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - nameAvailabilityRequest - NameAvailabilityRequest object.
//   - options - CheckNameAvailabilityClientCheckNameAvailabilityOptions contains the optional parameters for the CheckNameAvailabilityClient.CheckNameAvailability
//     method.
func (client *CheckNameAvailabilityClient) CheckNameAvailability(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, options *CheckNameAvailabilityClientCheckNameAvailabilityOptions) (CheckNameAvailabilityClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "CheckNameAvailabilityClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return CheckNameAvailabilityClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CheckNameAvailabilityClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CheckNameAvailabilityClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *CheckNameAvailabilityClient) checkNameAvailabilityCreateRequest(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, _ *CheckNameAvailabilityClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AgFoodPlatform/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, nameAvailabilityRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *CheckNameAvailabilityClient) checkNameAvailabilityHandleResponse(resp *http.Response) (CheckNameAvailabilityClientCheckNameAvailabilityResponse, error) {
	result := CheckNameAvailabilityClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return CheckNameAvailabilityClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}
