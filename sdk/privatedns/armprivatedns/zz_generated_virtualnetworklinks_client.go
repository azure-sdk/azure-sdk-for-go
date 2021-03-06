// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// VirtualNetworkLinksClient contains the methods for the VirtualNetworkLinks group.
// Don't use this type directly, use NewVirtualNetworkLinksClient() instead.
type VirtualNetworkLinksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualNetworkLinksClient creates a new instance of VirtualNetworkLinksClient with the specified values.
func NewVirtualNetworkLinksClient(con *armcore.Connection, subscriptionID string) *VirtualNetworkLinksClient {
	return &VirtualNetworkLinksClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksBeginCreateOrUpdateOptions) (VirtualNetworkLinkPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	result := VirtualNetworkLinkPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualNetworkLinksClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	poller := &virtualNetworkLinkPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkLinkResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualNetworkLinkPoller from the specified resume token.
// token - The value must come from a previous call to VirtualNetworkLinkPoller.ResumeToken().
func (client *VirtualNetworkLinksClient) ResumeCreateOrUpdate(ctx context.Context, token string) (VirtualNetworkLinkPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualNetworkLinksClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	poller := &virtualNetworkLinkPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	result := VirtualNetworkLinkPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkLinkResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworkLinksClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes a virtual network link to the specified Private DNS zone. WARNING: In case of a registration virtual network, all auto-registered
// DNS records in the zone for the virtual network will also be
// deleted. This operation cannot be undone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) BeginDelete(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualNetworkLinksClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualNetworkLinksClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualNetworkLinksClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a virtual network link to the specified Private DNS zone. WARNING: In case of a registration virtual network, all auto-registered DNS
// records in the zone for the virtual network will also be
// deleted. This operation cannot be undone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) deleteOperation(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualNetworkLinksClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets a virtual network link to the specified Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) Get(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksGetOptions) (VirtualNetworkLinkResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, options)
	if err != nil {
		return VirtualNetworkLinkResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkLinkResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualNetworkLinkResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkLinksClient) getHandleResponse(resp *azcore.Response) (VirtualNetworkLinkResponse, error) {
	var val *VirtualNetworkLink
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkLinkResponse{}, err
	}
	return VirtualNetworkLinkResponse{RawResponse: resp.Response, VirtualNetworkLink: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualNetworkLinksClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Lists the virtual network links to the specified Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) List(resourceGroupName string, privateZoneName string, options *VirtualNetworkLinksListOptions) VirtualNetworkLinkListResultPager {
	return &virtualNetworkLinkListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateZoneName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp VirtualNetworkLinkListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkLinkListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualNetworkLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, options *VirtualNetworkLinksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworkLinksClient) listHandleResponse(resp *azcore.Response) (VirtualNetworkLinkListResultResponse, error) {
	var val *VirtualNetworkLinkListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkLinkListResultResponse{}, err
	}
	return VirtualNetworkLinkListResultResponse{RawResponse: resp.Response, VirtualNetworkLinkListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualNetworkLinksClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - Updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) BeginUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksBeginUpdateOptions) (VirtualNetworkLinkPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	result := VirtualNetworkLinkPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualNetworkLinksClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	poller := &virtualNetworkLinkPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkLinkResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new VirtualNetworkLinkPoller from the specified resume token.
// token - The value must come from a previous call to VirtualNetworkLinkPoller.ResumeToken().
func (client *VirtualNetworkLinksClient) ResumeUpdate(ctx context.Context, token string) (VirtualNetworkLinkPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualNetworkLinksClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	poller := &virtualNetworkLinkPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualNetworkLinkPollerResponse{}, err
	}
	result := VirtualNetworkLinkPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkLinkResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkLinksClient) update(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualNetworkLinksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleError handles the Update error response.
func (client *VirtualNetworkLinksClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
