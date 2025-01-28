//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// RegisteredAsnsClient contains the methods for the RegisteredAsns group.
// Don't use this type directly, use NewRegisteredAsnsClient() instead.
type RegisteredAsnsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegisteredAsnsClient creates a new instance of RegisteredAsnsClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegisteredAsnsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegisteredAsnsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegisteredAsnsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new registered ASN with the specified name under the given subscription, resource group and
// peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - registeredAsnName - The name of the ASN.
//   - registeredAsn - The properties needed to create a registered ASN.
//   - options - RegisteredAsnsClientCreateOrUpdateOptions contains the optional parameters for the RegisteredAsnsClient.CreateOrUpdate
//     method.
func (client *RegisteredAsnsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, registeredAsn RegisteredAsn, options *RegisteredAsnsClientCreateOrUpdateOptions) (RegisteredAsnsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "RegisteredAsnsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringName, registeredAsnName, registeredAsn, options)
	if err != nil {
		return RegisteredAsnsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegisteredAsnsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RegisteredAsnsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegisteredAsnsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, registeredAsn RegisteredAsn, options *RegisteredAsnsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredAsnName == "" {
		return nil, errors.New("parameter registeredAsnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredAsnName}", url.PathEscape(registeredAsnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, registeredAsn); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RegisteredAsnsClient) createOrUpdateHandleResponse(resp *http.Response) (RegisteredAsnsClientCreateOrUpdateResponse, error) {
	result := RegisteredAsnsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredAsn); err != nil {
		return RegisteredAsnsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing registered ASN with the specified name under the given subscription, resource group and peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - registeredAsnName - The name of the registered ASN.
//   - options - RegisteredAsnsClientDeleteOptions contains the optional parameters for the RegisteredAsnsClient.Delete method.
func (client *RegisteredAsnsClient) Delete(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, options *RegisteredAsnsClientDeleteOptions) (RegisteredAsnsClientDeleteResponse, error) {
	var err error
	const operationName = "RegisteredAsnsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringName, registeredAsnName, options)
	if err != nil {
		return RegisteredAsnsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegisteredAsnsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RegisteredAsnsClientDeleteResponse{}, err
	}
	return RegisteredAsnsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegisteredAsnsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, options *RegisteredAsnsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredAsnName == "" {
		return nil, errors.New("parameter registeredAsnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredAsnName}", url.PathEscape(registeredAsnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing registered ASN with the specified name under the given subscription, resource group and peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - registeredAsnName - The name of the registered ASN.
//   - options - RegisteredAsnsClientGetOptions contains the optional parameters for the RegisteredAsnsClient.Get method.
func (client *RegisteredAsnsClient) Get(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, options *RegisteredAsnsClientGetOptions) (RegisteredAsnsClientGetResponse, error) {
	var err error
	const operationName = "RegisteredAsnsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringName, registeredAsnName, options)
	if err != nil {
		return RegisteredAsnsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegisteredAsnsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RegisteredAsnsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RegisteredAsnsClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, options *RegisteredAsnsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredAsnName == "" {
		return nil, errors.New("parameter registeredAsnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredAsnName}", url.PathEscape(registeredAsnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegisteredAsnsClient) getHandleResponse(resp *http.Response) (RegisteredAsnsClientGetResponse, error) {
	result := RegisteredAsnsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredAsn); err != nil {
		return RegisteredAsnsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPeeringPager - Lists all registered ASNs under the given subscription, resource group and peering.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - options - RegisteredAsnsClientListByPeeringOptions contains the optional parameters for the RegisteredAsnsClient.NewListByPeeringPager
//     method.
func (client *RegisteredAsnsClient) NewListByPeeringPager(resourceGroupName string, peeringName string, options *RegisteredAsnsClientListByPeeringOptions) *runtime.Pager[RegisteredAsnsClientListByPeeringResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegisteredAsnsClientListByPeeringResponse]{
		More: func(page RegisteredAsnsClientListByPeeringResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegisteredAsnsClientListByPeeringResponse) (RegisteredAsnsClientListByPeeringResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RegisteredAsnsClient.NewListByPeeringPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPeeringCreateRequest(ctx, resourceGroupName, peeringName, options)
			}, nil)
			if err != nil {
				return RegisteredAsnsClientListByPeeringResponse{}, err
			}
			return client.listByPeeringHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPeeringCreateRequest creates the ListByPeering request.
func (client *RegisteredAsnsClient) listByPeeringCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, options *RegisteredAsnsClientListByPeeringOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPeeringHandleResponse handles the ListByPeering response.
func (client *RegisteredAsnsClient) listByPeeringHandleResponse(resp *http.Response) (RegisteredAsnsClientListByPeeringResponse, error) {
	result := RegisteredAsnsClientListByPeeringResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredAsnListResult); err != nil {
		return RegisteredAsnsClientListByPeeringResponse{}, err
	}
	return result, nil
}
