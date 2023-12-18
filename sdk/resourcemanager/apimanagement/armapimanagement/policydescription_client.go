//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// PolicyDescriptionClient contains the methods for the PolicyDescription group.
// Don't use this type directly, use NewPolicyDescriptionClient() instead.
type PolicyDescriptionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPolicyDescriptionClient creates a new instance of PolicyDescriptionClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPolicyDescriptionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicyDescriptionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PolicyDescriptionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// ListByService - Lists all policy descriptions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - PolicyDescriptionClientListByServiceOptions contains the optional parameters for the PolicyDescriptionClient.ListByService
//     method.
func (client *PolicyDescriptionClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, options *PolicyDescriptionClientListByServiceOptions) (PolicyDescriptionClientListByServiceResponse, error) {
	var err error
	const operationName = "PolicyDescriptionClient.ListByService"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return PolicyDescriptionClientListByServiceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicyDescriptionClientListByServiceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicyDescriptionClientListByServiceResponse{}, err
	}
	resp, err := client.listByServiceHandleResponse(httpResp)
	return resp, err
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PolicyDescriptionClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PolicyDescriptionClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policyDescriptions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Scope != nil {
		reqQP.Set("scope", string(*options.Scope))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PolicyDescriptionClient) listByServiceHandleResponse(resp *http.Response) (PolicyDescriptionClientListByServiceResponse, error) {
	result := PolicyDescriptionClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDescriptionCollection); err != nil {
		return PolicyDescriptionClientListByServiceResponse{}, err
	}
	return result, nil
}
