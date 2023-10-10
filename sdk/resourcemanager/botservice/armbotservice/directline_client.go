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

// DirectLineClient contains the methods for the DirectLine group.
// Don't use this type directly, use NewDirectLineClient() instead.
type DirectLineClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDirectLineClient creates a new instance of DirectLineClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDirectLineClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DirectLineClient, error) {
	cl, err := arm.NewClient(moduleName+".DirectLineClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DirectLineClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// RegenerateKeys - Regenerates secret keys and returns them for the DirectLine Channel of a particular BotService resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - channelName - The name of the Channel resource for which keys are to be regenerated.
//   - parameters - The parameters to provide for the created bot.
//   - options - DirectLineClientRegenerateKeysOptions contains the optional parameters for the DirectLineClient.RegenerateKeys
//     method.
func (client *DirectLineClient) RegenerateKeys(ctx context.Context, resourceGroupName string, resourceName string, channelName RegenerateKeysChannelName, parameters SiteInfo, options *DirectLineClientRegenerateKeysOptions) (DirectLineClientRegenerateKeysResponse, error) {
	var err error
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, resourceName, channelName, parameters, options)
	if err != nil {
		return DirectLineClientRegenerateKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DirectLineClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DirectLineClientRegenerateKeysResponse{}, err
	}
	resp, err := client.regenerateKeysHandleResponse(httpResp)
	return resp, err
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *DirectLineClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, channelName RegenerateKeysChannelName, parameters SiteInfo, options *DirectLineClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}/regeneratekeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(string(channelName)))
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

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *DirectLineClient) regenerateKeysHandleResponse(resp *http.Response) (DirectLineClientRegenerateKeysResponse, error) {
	result := DirectLineClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotChannel); err != nil {
		return DirectLineClientRegenerateKeysResponse{}, err
	}
	return result, nil
}
