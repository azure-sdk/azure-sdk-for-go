// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// SecurityPerimeterAssociationsClient contains the methods for the NetworkSecurityPerimeterAssociations group.
// Don't use this type directly, use NewSecurityPerimeterAssociationsClient() instead.
type SecurityPerimeterAssociationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecurityPerimeterAssociationsClient creates a new instance of SecurityPerimeterAssociationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecurityPerimeterAssociationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityPerimeterAssociationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecurityPerimeterAssociationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a NSP resource association.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityPerimeterName - The name of the network security perimeter.
//   - associationName - The name of the NSP association.
//   - parameters - Parameters that hold the NspAssociation resource to be created/updated.
//   - options - SecurityPerimeterAssociationsClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityPerimeterAssociationsClient.BeginCreateOrUpdate
//     method.
func (client *SecurityPerimeterAssociationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, parameters NspAssociation, options *SecurityPerimeterAssociationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecurityPerimeterAssociationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityPerimeterName, associationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecurityPerimeterAssociationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SecurityPerimeterAssociationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a NSP resource association.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *SecurityPerimeterAssociationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, parameters NspAssociation, options *SecurityPerimeterAssociationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SecurityPerimeterAssociationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityPerimeterName, associationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityPerimeterAssociationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, parameters NspAssociation, _ *SecurityPerimeterAssociationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/resourceAssociations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityPerimeterName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterName}", url.PathEscape(networkSecurityPerimeterName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an NSP association resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityPerimeterName - The name of the network security perimeter.
//   - associationName - The name of the NSP association.
//   - options - SecurityPerimeterAssociationsClientBeginDeleteOptions contains the optional parameters for the SecurityPerimeterAssociationsClient.BeginDelete
//     method.
func (client *SecurityPerimeterAssociationsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, options *SecurityPerimeterAssociationsClientBeginDeleteOptions) (*runtime.Poller[SecurityPerimeterAssociationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityPerimeterName, associationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecurityPerimeterAssociationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SecurityPerimeterAssociationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an NSP association resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *SecurityPerimeterAssociationsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, options *SecurityPerimeterAssociationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SecurityPerimeterAssociationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityPerimeterName, associationName, options)
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
func (client *SecurityPerimeterAssociationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, _ *SecurityPerimeterAssociationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/resourceAssociations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityPerimeterName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterName}", url.PathEscape(networkSecurityPerimeterName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified NSP association by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityPerimeterName - The name of the network security perimeter.
//   - associationName - The name of the NSP association.
//   - options - SecurityPerimeterAssociationsClientGetOptions contains the optional parameters for the SecurityPerimeterAssociationsClient.Get
//     method.
func (client *SecurityPerimeterAssociationsClient) Get(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, options *SecurityPerimeterAssociationsClientGetOptions) (SecurityPerimeterAssociationsClientGetResponse, error) {
	var err error
	const operationName = "SecurityPerimeterAssociationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityPerimeterName, associationName, options)
	if err != nil {
		return SecurityPerimeterAssociationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityPerimeterAssociationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecurityPerimeterAssociationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecurityPerimeterAssociationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, _ *SecurityPerimeterAssociationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/resourceAssociations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityPerimeterName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterName}", url.PathEscape(networkSecurityPerimeterName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
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
func (client *SecurityPerimeterAssociationsClient) getHandleResponse(resp *http.Response) (SecurityPerimeterAssociationsClientGetResponse, error) {
	result := SecurityPerimeterAssociationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NspAssociation); err != nil {
		return SecurityPerimeterAssociationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the NSP resource associations.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityPerimeterName - The name of the network security perimeter.
//   - options - SecurityPerimeterAssociationsClientListOptions contains the optional parameters for the SecurityPerimeterAssociationsClient.NewListPager
//     method.
func (client *SecurityPerimeterAssociationsClient) NewListPager(resourceGroupName string, networkSecurityPerimeterName string, options *SecurityPerimeterAssociationsClientListOptions) *runtime.Pager[SecurityPerimeterAssociationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityPerimeterAssociationsClientListResponse]{
		More: func(page SecurityPerimeterAssociationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityPerimeterAssociationsClientListResponse) (SecurityPerimeterAssociationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecurityPerimeterAssociationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, networkSecurityPerimeterName, options)
			}, nil)
			if err != nil {
				return SecurityPerimeterAssociationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SecurityPerimeterAssociationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, options *SecurityPerimeterAssociationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/resourceAssociations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityPerimeterName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterName}", url.PathEscape(networkSecurityPerimeterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
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
func (client *SecurityPerimeterAssociationsClient) listHandleResponse(resp *http.Response) (SecurityPerimeterAssociationsClientListResponse, error) {
	result := SecurityPerimeterAssociationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NspAssociationsListResult); err != nil {
		return SecurityPerimeterAssociationsClientListResponse{}, err
	}
	return result, nil
}
