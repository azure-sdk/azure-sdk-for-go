//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// CommunicationsGatewaysClient contains the methods for the CommunicationsGateways group.
// Don't use this type directly, use NewCommunicationsGatewaysClient() instead.
type CommunicationsGatewaysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCommunicationsGatewaysClient creates a new instance of CommunicationsGatewaysClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCommunicationsGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CommunicationsGatewaysClient, error) {
	cl, err := arm.NewClient(moduleName+".CommunicationsGatewaysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CommunicationsGatewaysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckLocal - Implements global CheckNameAvailability operations
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-03
//   - location - The location name.
//   - body - The CheckAvailability request
//   - options - CommunicationsGatewaysClientCheckLocalOptions contains the optional parameters for the CommunicationsGatewaysClient.CheckLocal
//     method.
func (client *CommunicationsGatewaysClient) CheckLocal(ctx context.Context, location string, body CheckNameAvailabilityRequest, options *CommunicationsGatewaysClientCheckLocalOptions) (CommunicationsGatewaysClientCheckLocalResponse, error) {
	req, err := client.checkLocalCreateRequest(ctx, location, body, options)
	if err != nil {
		return CommunicationsGatewaysClientCheckLocalResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CommunicationsGatewaysClientCheckLocalResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunicationsGatewaysClientCheckLocalResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkLocalHandleResponse(resp)
}

// checkLocalCreateRequest creates the CheckLocal request.
func (client *CommunicationsGatewaysClient) checkLocalCreateRequest(ctx context.Context, location string, body CheckNameAvailabilityRequest, options *CommunicationsGatewaysClientCheckLocalOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/locations/{location}/providers/Microsoft.VoiceServices/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// checkLocalHandleResponse handles the CheckLocal response.
func (client *CommunicationsGatewaysClient) checkLocalHandleResponse(resp *http.Response) (CommunicationsGatewaysClientCheckLocalResponse, error) {
	result := CommunicationsGatewaysClientCheckLocalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return CommunicationsGatewaysClientCheckLocalResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create a CommunicationsGateway
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - resource - Resource create parameters.
//   - options - CommunicationsGatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the CommunicationsGatewaysClient.BeginCreateOrUpdate
//     method.
func (client *CommunicationsGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, communicationsGatewayName string, resource CommunicationsGateway, options *CommunicationsGatewaysClientBeginCreateOrUpdateOptions) (*runtime.Poller[CommunicationsGatewaysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, communicationsGatewayName, resource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CommunicationsGatewaysClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CommunicationsGatewaysClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a CommunicationsGateway
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-03
func (client *CommunicationsGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, communicationsGatewayName string, resource CommunicationsGateway, options *CommunicationsGatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, communicationsGatewayName, resource, options)
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
func (client *CommunicationsGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, resource CommunicationsGateway, options *CommunicationsGatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// BeginDelete - Delete a CommunicationsGateway
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - options - CommunicationsGatewaysClientBeginDeleteOptions contains the optional parameters for the CommunicationsGatewaysClient.BeginDelete
//     method.
func (client *CommunicationsGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *CommunicationsGatewaysClientBeginDeleteOptions) (*runtime.Poller[CommunicationsGatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, communicationsGatewayName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CommunicationsGatewaysClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CommunicationsGatewaysClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a CommunicationsGateway
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-03
func (client *CommunicationsGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *CommunicationsGatewaysClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, communicationsGatewayName, options)
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
func (client *CommunicationsGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *CommunicationsGatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a CommunicationsGateway
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - options - CommunicationsGatewaysClientGetOptions contains the optional parameters for the CommunicationsGatewaysClient.Get
//     method.
func (client *CommunicationsGatewaysClient) Get(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *CommunicationsGatewaysClientGetOptions) (CommunicationsGatewaysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, communicationsGatewayName, options)
	if err != nil {
		return CommunicationsGatewaysClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CommunicationsGatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunicationsGatewaysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CommunicationsGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *CommunicationsGatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}"
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
	reqQP.Set("api-version", "2023-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommunicationsGatewaysClient) getHandleResponse(resp *http.Response) (CommunicationsGatewaysClientGetResponse, error) {
	result := CommunicationsGatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunicationsGateway); err != nil {
		return CommunicationsGatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List CommunicationsGateway resources by resource group
//
// Generated from API version 2023-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CommunicationsGatewaysClientListByResourceGroupOptions contains the optional parameters for the CommunicationsGatewaysClient.NewListByResourceGroupPager
//     method.
func (client *CommunicationsGatewaysClient) NewListByResourceGroupPager(resourceGroupName string, options *CommunicationsGatewaysClientListByResourceGroupOptions) *runtime.Pager[CommunicationsGatewaysClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CommunicationsGatewaysClientListByResourceGroupResponse]{
		More: func(page CommunicationsGatewaysClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CommunicationsGatewaysClientListByResourceGroupResponse) (CommunicationsGatewaysClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CommunicationsGatewaysClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CommunicationsGatewaysClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CommunicationsGatewaysClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CommunicationsGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CommunicationsGatewaysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CommunicationsGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (CommunicationsGatewaysClientListByResourceGroupResponse, error) {
	result := CommunicationsGatewaysClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunicationsGatewayListResult); err != nil {
		return CommunicationsGatewaysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List CommunicationsGateway resources by subscription ID
//
// Generated from API version 2023-04-03
//   - options - CommunicationsGatewaysClientListBySubscriptionOptions contains the optional parameters for the CommunicationsGatewaysClient.NewListBySubscriptionPager
//     method.
func (client *CommunicationsGatewaysClient) NewListBySubscriptionPager(options *CommunicationsGatewaysClientListBySubscriptionOptions) *runtime.Pager[CommunicationsGatewaysClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CommunicationsGatewaysClientListBySubscriptionResponse]{
		More: func(page CommunicationsGatewaysClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CommunicationsGatewaysClientListBySubscriptionResponse) (CommunicationsGatewaysClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CommunicationsGatewaysClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CommunicationsGatewaysClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CommunicationsGatewaysClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CommunicationsGatewaysClient) listBySubscriptionCreateRequest(ctx context.Context, options *CommunicationsGatewaysClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VoiceServices/communicationsGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CommunicationsGatewaysClient) listBySubscriptionHandleResponse(resp *http.Response) (CommunicationsGatewaysClientListBySubscriptionResponse, error) {
	result := CommunicationsGatewaysClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunicationsGatewayListResult); err != nil {
		return CommunicationsGatewaysClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a CommunicationsGateway
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationsGatewayName - Unique identifier for this deployment
//   - properties - The resource properties to be updated.
//   - options - CommunicationsGatewaysClientUpdateOptions contains the optional parameters for the CommunicationsGatewaysClient.Update
//     method.
func (client *CommunicationsGatewaysClient) Update(ctx context.Context, resourceGroupName string, communicationsGatewayName string, properties CommunicationsGatewayUpdate, options *CommunicationsGatewaysClientUpdateOptions) (CommunicationsGatewaysClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, communicationsGatewayName, properties, options)
	if err != nil {
		return CommunicationsGatewaysClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CommunicationsGatewaysClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunicationsGatewaysClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CommunicationsGatewaysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, communicationsGatewayName string, properties CommunicationsGatewayUpdate, options *CommunicationsGatewaysClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VoiceServices/communicationsGateways/{communicationsGatewayName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// updateHandleResponse handles the Update response.
func (client *CommunicationsGatewaysClient) updateHandleResponse(resp *http.Response) (CommunicationsGatewaysClientUpdateResponse, error) {
	result := CommunicationsGatewaysClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunicationsGateway); err != nil {
		return CommunicationsGatewaysClientUpdateResponse{}, err
	}
	return result, nil
}
