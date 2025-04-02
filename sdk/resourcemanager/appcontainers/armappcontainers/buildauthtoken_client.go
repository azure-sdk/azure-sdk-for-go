// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// BuildAuthTokenClient contains the methods for the BuildAuthToken group.
// Don't use this type directly, use NewBuildAuthTokenClient() instead.
type BuildAuthTokenClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBuildAuthTokenClient creates a new instance of BuildAuthTokenClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBuildAuthTokenClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BuildAuthTokenClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BuildAuthTokenClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Gets the token used to connect to the endpoint where source code can be uploaded for a build.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - builderName - The name of the builder.
//   - buildName - The name of a build.
//   - options - BuildAuthTokenClientListOptions contains the optional parameters for the BuildAuthTokenClient.List method.
func (client *BuildAuthTokenClient) List(ctx context.Context, resourceGroupName string, builderName string, buildName string, options *BuildAuthTokenClientListOptions) (BuildAuthTokenClientListResponse, error) {
	var err error
	const operationName = "BuildAuthTokenClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, builderName, buildName, options)
	if err != nil {
		return BuildAuthTokenClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BuildAuthTokenClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BuildAuthTokenClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *BuildAuthTokenClient) listCreateRequest(ctx context.Context, resourceGroupName string, builderName string, buildName string, _ *BuildAuthTokenClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/builders/{builderName}/builds/{buildName}/listAuthToken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	if buildName == "" {
		return nil, errors.New("parameter buildName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildName}", url.PathEscape(buildName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BuildAuthTokenClient) listHandleResponse(resp *http.Response) (BuildAuthTokenClientListResponse, error) {
	result := BuildAuthTokenClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuildToken); err != nil {
		return BuildAuthTokenClientListResponse{}, err
	}
	return result, nil
}
