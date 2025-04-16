// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// EndpointsClient contains the methods for the Endpoints group.
// Don't use this type directly, use NewEndpointsClient() instead.
type EndpointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEndpointsClient creates a new instance of EndpointsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EndpointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginPurgeContent - Removes a content from Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - frontDoorName - Name of the Front Door which is globally unique.
//   - contentFilePaths - The path to the content to be purged. Path can be a full URL, e.g. '/pictures/city.png' which removes
//     a single file, or a directory with a wildcard, e.g. '/pictures/*' which removes all folders and
//     files in the directory.
//   - options - EndpointsClientBeginPurgeContentOptions contains the optional parameters for the EndpointsClient.BeginPurgeContent
//     method.
func (client *EndpointsClient) BeginPurgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters, options *EndpointsClientBeginPurgeContentOptions) (*runtime.Poller[EndpointsClientPurgeContentResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purgeContent(ctx, resourceGroupName, frontDoorName, contentFilePaths, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EndpointsClientPurgeContentResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EndpointsClientPurgeContentResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// PurgeContent - Removes a content from Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *EndpointsClient) purgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters, options *EndpointsClientBeginPurgeContentOptions) (*http.Response, error) {
	var err error
	const operationName = "EndpointsClient.BeginPurgeContent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.purgeContentCreateRequest(ctx, resourceGroupName, frontDoorName, contentFilePaths, options)
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

// purgeContentCreateRequest creates the PurgeContent request.
func (client *EndpointsClient) purgeContentCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters, _ *EndpointsClientBeginPurgeContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/purge"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, contentFilePaths); err != nil {
		return nil, err
	}
	return req, nil
}
