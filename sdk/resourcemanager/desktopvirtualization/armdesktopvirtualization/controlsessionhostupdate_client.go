// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// ControlSessionHostUpdateClient contains the methods for the ControlSessionHostUpdate group.
// Don't use this type directly, use NewControlSessionHostUpdateClient() instead.
type ControlSessionHostUpdateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewControlSessionHostUpdateClient creates a new instance of ControlSessionHostUpdateClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewControlSessionHostUpdateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ControlSessionHostUpdateClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ControlSessionHostUpdateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginPost - Control update of a hostpool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - hostPoolUpdateControlParameter - Object containing control action for hostpool update.
//   - options - ControlSessionHostUpdateClientBeginPostOptions contains the optional parameters for the ControlSessionHostUpdateClient.BeginPost
//     method.
func (client *ControlSessionHostUpdateClient) BeginPost(ctx context.Context, resourceGroupName string, hostPoolName string, hostPoolUpdateControlParameter HostPoolUpdateControlParameter, options *ControlSessionHostUpdateClientBeginPostOptions) (*runtime.Poller[ControlSessionHostUpdateClientPostResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.post(ctx, resourceGroupName, hostPoolName, hostPoolUpdateControlParameter, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ControlSessionHostUpdateClientPostResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ControlSessionHostUpdateClientPostResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Post - Control update of a hostpool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
func (client *ControlSessionHostUpdateClient) post(ctx context.Context, resourceGroupName string, hostPoolName string, hostPoolUpdateControlParameter HostPoolUpdateControlParameter, options *ControlSessionHostUpdateClientBeginPostOptions) (*http.Response, error) {
	var err error
	const operationName = "ControlSessionHostUpdateClient.BeginPost"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postCreateRequest(ctx, resourceGroupName, hostPoolName, hostPoolUpdateControlParameter, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// postCreateRequest creates the Post request.
func (client *ControlSessionHostUpdateClient) postCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, hostPoolUpdateControlParameter HostPoolUpdateControlParameter, _ *ControlSessionHostUpdateClientBeginPostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostManagements/default/controlSessionHostUpdate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, hostPoolUpdateControlParameter); err != nil {
		return nil, err
	}
	return req, nil
}
