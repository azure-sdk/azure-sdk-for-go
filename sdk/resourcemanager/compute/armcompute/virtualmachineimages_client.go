//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualMachineImagesClient contains the methods for the VirtualMachineImages group.
// Don't use this type directly, use NewVirtualMachineImagesClient() instead.
type VirtualMachineImagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachineImagesClient creates a new instance of VirtualMachineImagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineImagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineImagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a virtual machine image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of Azure region.
//   - publisherName - A valid image publisher.
//   - offer - A valid image publisher offer.
//   - skus - A valid image SKU.
//   - version - A valid image SKU version.
//   - options - VirtualMachineImagesClientGetOptions contains the optional parameters for the VirtualMachineImagesClient.Get
//     method.
func (client *VirtualMachineImagesClient) Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesClientGetOptions) (VirtualMachineImagesClientGetResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, publisherName, offer, skus, version, options)
	if err != nil {
		return VirtualMachineImagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineImagesClient) getCreateRequest(ctx context.Context, location string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineImagesClient) getHandleResponse(resp *http.Response) (VirtualMachineImagesClientGetResponse, error) {
	result := VirtualMachineImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImage); err != nil {
		return VirtualMachineImagesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of Azure region.
//   - publisherName - A valid image publisher.
//   - offer - A valid image publisher offer.
//   - skus - A valid image SKU.
//   - options - VirtualMachineImagesClientListOptions contains the optional parameters for the VirtualMachineImagesClient.List
//     method.
func (client *VirtualMachineImagesClient) List(ctx context.Context, location string, publisherName string, offer string, skus string, options *VirtualMachineImagesClientListOptions) (VirtualMachineImagesClientListResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, location, publisherName, offer, skus, options)
	if err != nil {
		return VirtualMachineImagesClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *VirtualMachineImagesClient) listCreateRequest(ctx context.Context, location string, publisherName string, offer string, skus string, options *VirtualMachineImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineImagesClient) listHandleResponse(resp *http.Response) (VirtualMachineImagesClientListResponse, error) {
	result := VirtualMachineImagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesClientListResponse{}, err
	}
	return result, nil
}

// ListByEdgeZone - Gets a list of all virtual machine image versions for the specified edge zone
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of Azure region.
//   - edgeZone - The name of the edge zone.
//   - options - VirtualMachineImagesClientListByEdgeZoneOptions contains the optional parameters for the VirtualMachineImagesClient.ListByEdgeZone
//     method.
func (client *VirtualMachineImagesClient) ListByEdgeZone(ctx context.Context, location string, edgeZone string, options *VirtualMachineImagesClientListByEdgeZoneOptions) (VirtualMachineImagesClientListByEdgeZoneResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesClient.ListByEdgeZone"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByEdgeZoneCreateRequest(ctx, location, edgeZone, options)
	if err != nil {
		return VirtualMachineImagesClientListByEdgeZoneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesClientListByEdgeZoneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesClientListByEdgeZoneResponse{}, err
	}
	resp, err := client.listByEdgeZoneHandleResponse(httpResp)
	return resp, err
}

// listByEdgeZoneCreateRequest creates the ListByEdgeZone request.
func (client *VirtualMachineImagesClient) listByEdgeZoneCreateRequest(ctx context.Context, location string, edgeZone string, options *VirtualMachineImagesClientListByEdgeZoneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/vmimages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEdgeZoneHandleResponse handles the ListByEdgeZone response.
func (client *VirtualMachineImagesClient) listByEdgeZoneHandleResponse(resp *http.Response) (VirtualMachineImagesClientListByEdgeZoneResponse, error) {
	result := VirtualMachineImagesClientListByEdgeZoneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMImagesInEdgeZoneListResult); err != nil {
		return VirtualMachineImagesClientListByEdgeZoneResponse{}, err
	}
	return result, nil
}

// ListOffers - Gets a list of virtual machine image offers for the specified location and publisher.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of Azure region.
//   - publisherName - A valid image publisher.
//   - options - VirtualMachineImagesClientListOffersOptions contains the optional parameters for the VirtualMachineImagesClient.ListOffers
//     method.
func (client *VirtualMachineImagesClient) ListOffers(ctx context.Context, location string, publisherName string, options *VirtualMachineImagesClientListOffersOptions) (VirtualMachineImagesClientListOffersResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesClient.ListOffers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listOffersCreateRequest(ctx, location, publisherName, options)
	if err != nil {
		return VirtualMachineImagesClientListOffersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesClientListOffersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesClientListOffersResponse{}, err
	}
	resp, err := client.listOffersHandleResponse(httpResp)
	return resp, err
}

// listOffersCreateRequest creates the ListOffers request.
func (client *VirtualMachineImagesClient) listOffersCreateRequest(ctx context.Context, location string, publisherName string, options *VirtualMachineImagesClientListOffersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOffersHandleResponse handles the ListOffers response.
func (client *VirtualMachineImagesClient) listOffersHandleResponse(resp *http.Response) (VirtualMachineImagesClientListOffersResponse, error) {
	result := VirtualMachineImagesClientListOffersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesClientListOffersResponse{}, err
	}
	return result, nil
}

// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of Azure region.
//   - options - VirtualMachineImagesClientListPublishersOptions contains the optional parameters for the VirtualMachineImagesClient.ListPublishers
//     method.
func (client *VirtualMachineImagesClient) ListPublishers(ctx context.Context, location string, options *VirtualMachineImagesClientListPublishersOptions) (VirtualMachineImagesClientListPublishersResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesClient.ListPublishers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listPublishersCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineImagesClientListPublishersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesClientListPublishersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesClientListPublishersResponse{}, err
	}
	resp, err := client.listPublishersHandleResponse(httpResp)
	return resp, err
}

// listPublishersCreateRequest creates the ListPublishers request.
func (client *VirtualMachineImagesClient) listPublishersCreateRequest(ctx context.Context, location string, options *VirtualMachineImagesClientListPublishersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listPublishersHandleResponse handles the ListPublishers response.
func (client *VirtualMachineImagesClient) listPublishersHandleResponse(resp *http.Response) (VirtualMachineImagesClientListPublishersResponse, error) {
	result := VirtualMachineImagesClientListPublishersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesClientListPublishersResponse{}, err
	}
	return result, nil
}

// ListSKUs - Gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of Azure region.
//   - publisherName - A valid image publisher.
//   - offer - A valid image publisher offer.
//   - options - VirtualMachineImagesClientListSKUsOptions contains the optional parameters for the VirtualMachineImagesClient.ListSKUs
//     method.
func (client *VirtualMachineImagesClient) ListSKUs(ctx context.Context, location string, publisherName string, offer string, options *VirtualMachineImagesClientListSKUsOptions) (VirtualMachineImagesClientListSKUsResponse, error) {
	var err error
	const operationName = "VirtualMachineImagesClient.ListSKUs"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSKUsCreateRequest(ctx, location, publisherName, offer, options)
	if err != nil {
		return VirtualMachineImagesClientListSKUsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImagesClientListSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineImagesClientListSKUsResponse{}, err
	}
	resp, err := client.listSKUsHandleResponse(httpResp)
	return resp, err
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *VirtualMachineImagesClient) listSKUsCreateRequest(ctx context.Context, location string, publisherName string, offer string, options *VirtualMachineImagesClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *VirtualMachineImagesClient) listSKUsHandleResponse(resp *http.Response) (VirtualMachineImagesClientListSKUsResponse, error) {
	result := VirtualMachineImagesClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesClientListSKUsResponse{}, err
	}
	return result, nil
}
