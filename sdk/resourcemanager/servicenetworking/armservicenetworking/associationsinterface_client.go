//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicenetworking

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

// AssociationsInterfaceClient contains the methods for the AssociationsInterface group.
// Don't use this type directly, use NewAssociationsInterfaceClient() instead.
type AssociationsInterfaceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssociationsInterfaceClient creates a new instance of AssociationsInterfaceClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssociationsInterfaceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssociationsInterfaceClient, error) {
	cl, err := arm.NewClient(moduleName+".AssociationsInterfaceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssociationsInterfaceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - A sequence of textual characters.
//   - associationName - A sequence of textual characters.
//   - resource - Resource create parameters.
//   - options - AssociationsInterfaceClientBeginCreateOrUpdateOptions contains the optional parameters for the AssociationsInterfaceClient.BeginCreateOrUpdate
//     method.
func (client *AssociationsInterfaceClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, resource Association, options *AssociationsInterfaceClientBeginCreateOrUpdateOptions) (*runtime.Poller[AssociationsInterfaceClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, trafficControllerName, associationName, resource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AssociationsInterfaceClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AssociationsInterfaceClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *AssociationsInterfaceClient) createOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, resource Association, options *AssociationsInterfaceClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, resource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssociationsInterfaceClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, resource Association, options *AssociationsInterfaceClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// BeginDelete - Delete a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - A sequence of textual characters.
//   - associationName - A sequence of textual characters.
//   - options - AssociationsInterfaceClientBeginDeleteOptions contains the optional parameters for the AssociationsInterfaceClient.BeginDelete
//     method.
func (client *AssociationsInterfaceClient) BeginDelete(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientBeginDeleteOptions) (*runtime.Poller[AssociationsInterfaceClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, trafficControllerName, associationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AssociationsInterfaceClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AssociationsInterfaceClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *AssociationsInterfaceClient) deleteOperation(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssociationsInterfaceClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - A sequence of textual characters.
//   - associationName - A sequence of textual characters.
//   - options - AssociationsInterfaceClientGetOptions contains the optional parameters for the AssociationsInterfaceClient.Get
//     method.
func (client *AssociationsInterfaceClient) Get(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientGetOptions) (AssociationsInterfaceClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, options)
	if err != nil {
		return AssociationsInterfaceClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssociationsInterfaceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssociationsInterfaceClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssociationsInterfaceClient) getCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssociationsInterfaceClient) getHandleResponse(resp *http.Response) (AssociationsInterfaceClientGetResponse, error) {
	result := AssociationsInterfaceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Association); err != nil {
		return AssociationsInterfaceClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTrafficControllerPager - List Association resources by TrafficController
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - A sequence of textual characters.
//   - options - AssociationsInterfaceClientListByTrafficControllerOptions contains the optional parameters for the AssociationsInterfaceClient.NewListByTrafficControllerPager
//     method.
func (client *AssociationsInterfaceClient) NewListByTrafficControllerPager(resourceGroupName string, trafficControllerName string, options *AssociationsInterfaceClientListByTrafficControllerOptions) *runtime.Pager[AssociationsInterfaceClientListByTrafficControllerResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssociationsInterfaceClientListByTrafficControllerResponse]{
		More: func(page AssociationsInterfaceClientListByTrafficControllerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssociationsInterfaceClientListByTrafficControllerResponse) (AssociationsInterfaceClientListByTrafficControllerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByTrafficControllerCreateRequest(ctx, resourceGroupName, trafficControllerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AssociationsInterfaceClientListByTrafficControllerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AssociationsInterfaceClientListByTrafficControllerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssociationsInterfaceClientListByTrafficControllerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTrafficControllerHandleResponse(resp)
		},
	})
}

// listByTrafficControllerCreateRequest creates the ListByTrafficController request.
func (client *AssociationsInterfaceClient) listByTrafficControllerCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, options *AssociationsInterfaceClientListByTrafficControllerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTrafficControllerHandleResponse handles the ListByTrafficController response.
func (client *AssociationsInterfaceClient) listByTrafficControllerHandleResponse(resp *http.Response) (AssociationsInterfaceClientListByTrafficControllerResponse, error) {
	result := AssociationsInterfaceClientListByTrafficControllerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssociationListResult); err != nil {
		return AssociationsInterfaceClientListByTrafficControllerResponse{}, err
	}
	return result, nil
}

// Update - Update a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - A sequence of textual characters.
//   - associationName - A sequence of textual characters.
//   - properties - The resource properties to be updated.
//   - options - AssociationsInterfaceClientUpdateOptions contains the optional parameters for the AssociationsInterfaceClient.Update
//     method.
func (client *AssociationsInterfaceClient) Update(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, properties AssociationUpdate, options *AssociationsInterfaceClientUpdateOptions) (AssociationsInterfaceClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, properties, options)
	if err != nil {
		return AssociationsInterfaceClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssociationsInterfaceClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssociationsInterfaceClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AssociationsInterfaceClient) updateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, properties AssociationUpdate, options *AssociationsInterfaceClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// updateHandleResponse handles the Update response.
func (client *AssociationsInterfaceClient) updateHandleResponse(resp *http.Response) (AssociationsInterfaceClientUpdateResponse, error) {
	result := AssociationsInterfaceClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Association); err != nil {
		return AssociationsInterfaceClientUpdateResponse{}, err
	}
	return result, nil
}
