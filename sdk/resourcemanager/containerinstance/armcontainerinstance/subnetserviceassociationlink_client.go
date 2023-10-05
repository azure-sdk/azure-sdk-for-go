//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerinstance

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

// SubnetServiceAssociationLinkClient contains the methods for the SubnetServiceAssociationLink group.
// Don't use this type directly, use NewSubnetServiceAssociationLinkClient() instead.
type SubnetServiceAssociationLinkClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubnetServiceAssociationLinkClient creates a new instance of SubnetServiceAssociationLinkClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubnetServiceAssociationLinkClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubnetServiceAssociationLinkClient, error) {
	cl, err := arm.NewClient(moduleName+".SubnetServiceAssociationLinkClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubnetServiceAssociationLinkClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Delete container group virtual network association links. The operation does not delete other resources provided
// by the user.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - The name of the virtual network.
//   - subnetName - The name of the subnet.
//   - options - SubnetServiceAssociationLinkClientBeginDeleteOptions contains the optional parameters for the SubnetServiceAssociationLinkClient.BeginDelete
//     method.
func (client *SubnetServiceAssociationLinkClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetServiceAssociationLinkClientBeginDeleteOptions) (*runtime.Poller[SubnetServiceAssociationLinkClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SubnetServiceAssociationLinkClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SubnetServiceAssociationLinkClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete container group virtual network association links. The operation does not delete other resources provided
// by the user.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *SubnetServiceAssociationLinkClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetServiceAssociationLinkClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SubnetServiceAssociationLinkClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetServiceAssociationLinkClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/providers/Microsoft.ContainerInstance/serviceAssociationLinks/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
