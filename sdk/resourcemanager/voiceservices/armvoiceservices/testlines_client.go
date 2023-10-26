//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvoiceservices

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

// TestLinesClient contains the methods for the TestLines group.
// Don't use this type directly, use NewTestLinesClient() instead.
type TestLinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTestLinesClient creates a new instance of TestLinesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTestLinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TestLinesClient, error) {
	cl, err := arm.NewClient(moduleName+".TestLinesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TestLinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a TestLine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - testLineName - Unique identifier for this test line
//   - resource - Resource create parameters.
//   - options - TestLinesClientBeginCreateOrUpdateOptions contains the optional parameters for the TestLinesClient.BeginCreateOrUpdate
//     method.
func (client *TestLinesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, resource TestLine, options *TestLinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[TestLinesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, communicationsGatewayName, testLineName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TestLinesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TestLinesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a TestLine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *TestLinesClient) createOrUpdate(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, resource TestLine, options *TestLinesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, communicationsGatewayName, testLineName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TestLinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, resource TestLine, options *TestLinesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}/testLines/{testLineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationsGatewayName == "" {
		return nil, errors.New("parameter communicationsGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationsGatewayName}", url.PathEscape(communicationsGatewayName))
	if testLineName == "" {
		return nil, errors.New("parameter testLineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testLineName}", url.PathEscape(testLineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a TestLine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - testLineName - Unique identifier for this test line
//   - options - TestLinesClientBeginDeleteOptions contains the optional parameters for the TestLinesClient.BeginDelete method.
func (client *TestLinesClient) BeginDelete(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, options *TestLinesClientBeginDeleteOptions) (*runtime.Poller[TestLinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, communicationsGatewayName, testLineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TestLinesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TestLinesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a TestLine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *TestLinesClient) deleteOperation(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, options *TestLinesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, communicationsGatewayName, testLineName, options)
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
func (client *TestLinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, options *TestLinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}/testLines/{testLineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationsGatewayName == "" {
		return nil, errors.New("parameter communicationsGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationsGatewayName}", url.PathEscape(communicationsGatewayName))
	if testLineName == "" {
		return nil, errors.New("parameter testLineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testLineName}", url.PathEscape(testLineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a TestLine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - testLineName - Unique identifier for this test line
//   - options - TestLinesClientGetOptions contains the optional parameters for the TestLinesClient.Get method.
func (client *TestLinesClient) Get(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, options *TestLinesClientGetOptions) (TestLinesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, communicationsGatewayName, testLineName, options)
	if err != nil {
		return TestLinesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TestLinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TestLinesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TestLinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, options *TestLinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}/testLines/{testLineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationsGatewayName == "" {
		return nil, errors.New("parameter communicationsGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationsGatewayName}", url.PathEscape(communicationsGatewayName))
	if testLineName == "" {
		return nil, errors.New("parameter testLineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testLineName}", url.PathEscape(testLineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TestLinesClient) getHandleResponse(resp *http.Response) (TestLinesClientGetResponse, error) {
	result := TestLinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestLine); err != nil {
		return TestLinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByCommunicationsGatewayPager - List TestLine resources by CommunicationsGateway
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - options - TestLinesClientListByCommunicationsGatewayOptions contains the optional parameters for the TestLinesClient.NewListByCommunicationsGatewayPager
//     method.
func (client *TestLinesClient) NewListByCommunicationsGatewayPager(resourceGroupName string, communicationsGatewayName string, options *TestLinesClientListByCommunicationsGatewayOptions) *runtime.Pager[TestLinesClientListByCommunicationsGatewayResponse] {
	return runtime.NewPager(runtime.PagingHandler[TestLinesClientListByCommunicationsGatewayResponse]{
		More: func(page TestLinesClientListByCommunicationsGatewayResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TestLinesClientListByCommunicationsGatewayResponse) (TestLinesClientListByCommunicationsGatewayResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByCommunicationsGatewayCreateRequest(ctx, resourceGroupName, communicationsGatewayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TestLinesClientListByCommunicationsGatewayResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TestLinesClientListByCommunicationsGatewayResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TestLinesClientListByCommunicationsGatewayResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByCommunicationsGatewayHandleResponse(resp)
		},
	})
}

// listByCommunicationsGatewayCreateRequest creates the ListByCommunicationsGateway request.
func (client *TestLinesClient) listByCommunicationsGatewayCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *TestLinesClientListByCommunicationsGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}/testLines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationsGatewayName == "" {
		return nil, errors.New("parameter communicationsGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationsGatewayName}", url.PathEscape(communicationsGatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCommunicationsGatewayHandleResponse handles the ListByCommunicationsGateway response.
func (client *TestLinesClient) listByCommunicationsGatewayHandleResponse(resp *http.Response) (TestLinesClientListByCommunicationsGatewayResponse, error) {
	result := TestLinesClientListByCommunicationsGatewayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestLineListResult); err != nil {
		return TestLinesClientListByCommunicationsGatewayResponse{}, err
	}
	return result, nil
}

// Update - Update a TestLine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - testLineName - Unique identifier for this test line
//   - properties - The resource properties to be updated.
//   - options - TestLinesClientUpdateOptions contains the optional parameters for the TestLinesClient.Update method.
func (client *TestLinesClient) Update(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, properties TestLineUpdate, options *TestLinesClientUpdateOptions) (TestLinesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, communicationsGatewayName, testLineName, properties, options)
	if err != nil {
		return TestLinesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TestLinesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TestLinesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TestLinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, testLineName string, properties TestLineUpdate, options *TestLinesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}/testLines/{testLineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationsGatewayName == "" {
		return nil, errors.New("parameter communicationsGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationsGatewayName}", url.PathEscape(communicationsGatewayName))
	if testLineName == "" {
		return nil, errors.New("parameter testLineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testLineName}", url.PathEscape(testLineName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TestLinesClient) updateHandleResponse(resp *http.Response) (TestLinesClientUpdateResponse, error) {
	result := TestLinesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestLine); err != nil {
		return TestLinesClientUpdateResponse{}, err
	}
	return result, nil
}
