//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// PartnerDestinationsClient contains the methods for the PartnerDestinations group.
// Don't use this type directly, use NewPartnerDestinationsClient() instead.
type PartnerDestinationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPartnerDestinationsClient creates a new instance of PartnerDestinationsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPartnerDestinationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PartnerDestinationsClient, error) {
	cl, err := arm.NewClient(moduleName+".PartnerDestinationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PartnerDestinationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Activate - Activate a newly created partner destination.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerDestinationName - Name of the partner destination.
//   - options - PartnerDestinationsClientActivateOptions contains the optional parameters for the PartnerDestinationsClient.Activate
//     method.
func (client *PartnerDestinationsClient) Activate(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *PartnerDestinationsClientActivateOptions) (PartnerDestinationsClientActivateResponse, error) {
	var err error
	req, err := client.activateCreateRequest(ctx, resourceGroupName, partnerDestinationName, options)
	if err != nil {
		return PartnerDestinationsClientActivateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerDestinationsClientActivateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerDestinationsClientActivateResponse{}, err
	}
	resp, err := client.activateHandleResponse(httpResp)
	return resp, err
}

// activateCreateRequest creates the Activate request.
func (client *PartnerDestinationsClient) activateCreateRequest(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *PartnerDestinationsClientActivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerDestinations/{partnerDestinationName}/activate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerDestinationName == "" {
		return nil, errors.New("parameter partnerDestinationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerDestinationName}", url.PathEscape(partnerDestinationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// activateHandleResponse handles the Activate response.
func (client *PartnerDestinationsClient) activateHandleResponse(resp *http.Response) (PartnerDestinationsClientActivateResponse, error) {
	result := PartnerDestinationsClientActivateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerDestination); err != nil {
		return PartnerDestinationsClientActivateResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Asynchronously creates a new partner destination with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerDestinationName - Name of the partner destination.
//   - partnerDestination - Partner destination create information.
//   - options - PartnerDestinationsClientBeginCreateOrUpdateOptions contains the optional parameters for the PartnerDestinationsClient.BeginCreateOrUpdate
//     method.
func (client *PartnerDestinationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestination PartnerDestination, options *PartnerDestinationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PartnerDestinationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, partnerDestinationName, partnerDestination, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PartnerDestinationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PartnerDestinationsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Asynchronously creates a new partner destination with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *PartnerDestinationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestination PartnerDestination, options *PartnerDestinationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerDestinationName, partnerDestination, options)
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
func (client *PartnerDestinationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestination PartnerDestination, options *PartnerDestinationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerDestinations/{partnerDestinationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerDestinationName == "" {
		return nil, errors.New("parameter partnerDestinationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerDestinationName}", url.PathEscape(partnerDestinationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, partnerDestination); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete existing partner destination.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerDestinationName - Name of the partner destination.
//   - options - PartnerDestinationsClientBeginDeleteOptions contains the optional parameters for the PartnerDestinationsClient.BeginDelete
//     method.
func (client *PartnerDestinationsClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *PartnerDestinationsClientBeginDeleteOptions) (*runtime.Poller[PartnerDestinationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, partnerDestinationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PartnerDestinationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PartnerDestinationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete existing partner destination.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *PartnerDestinationsClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *PartnerDestinationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerDestinationName, options)
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
func (client *PartnerDestinationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *PartnerDestinationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerDestinations/{partnerDestinationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerDestinationName == "" {
		return nil, errors.New("parameter partnerDestinationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerDestinationName}", url.PathEscape(partnerDestinationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of a partner destination.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerDestinationName - Name of the partner destination.
//   - options - PartnerDestinationsClientGetOptions contains the optional parameters for the PartnerDestinationsClient.Get method.
func (client *PartnerDestinationsClient) Get(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *PartnerDestinationsClientGetOptions) (PartnerDestinationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerDestinationName, options)
	if err != nil {
		return PartnerDestinationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerDestinationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerDestinationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PartnerDestinationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerDestinationName string, options *PartnerDestinationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerDestinations/{partnerDestinationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerDestinationName == "" {
		return nil, errors.New("parameter partnerDestinationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerDestinationName}", url.PathEscape(partnerDestinationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerDestinationsClient) getHandleResponse(resp *http.Response) (PartnerDestinationsClientGetResponse, error) {
	result := PartnerDestinationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerDestination); err != nil {
		return PartnerDestinationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the partner destinations under a resource group.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - options - PartnerDestinationsClientListByResourceGroupOptions contains the optional parameters for the PartnerDestinationsClient.NewListByResourceGroupPager
//     method.
func (client *PartnerDestinationsClient) NewListByResourceGroupPager(resourceGroupName string, options *PartnerDestinationsClientListByResourceGroupOptions) *runtime.Pager[PartnerDestinationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerDestinationsClientListByResourceGroupResponse]{
		More: func(page PartnerDestinationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerDestinationsClientListByResourceGroupResponse) (PartnerDestinationsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PartnerDestinationsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PartnerDestinationsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PartnerDestinationsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PartnerDestinationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PartnerDestinationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerDestinations"
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
	reqQP.Set("api-version", "2023-12-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PartnerDestinationsClient) listByResourceGroupHandleResponse(resp *http.Response) (PartnerDestinationsClientListByResourceGroupResponse, error) {
	result := PartnerDestinationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerDestinationsListResult); err != nil {
		return PartnerDestinationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the partner destinations under an Azure subscription.
//
// Generated from API version 2023-12-15-preview
//   - options - PartnerDestinationsClientListBySubscriptionOptions contains the optional parameters for the PartnerDestinationsClient.NewListBySubscriptionPager
//     method.
func (client *PartnerDestinationsClient) NewListBySubscriptionPager(options *PartnerDestinationsClientListBySubscriptionOptions) *runtime.Pager[PartnerDestinationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerDestinationsClientListBySubscriptionResponse]{
		More: func(page PartnerDestinationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerDestinationsClientListBySubscriptionResponse) (PartnerDestinationsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PartnerDestinationsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PartnerDestinationsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PartnerDestinationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PartnerDestinationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *PartnerDestinationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/partnerDestinations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PartnerDestinationsClient) listBySubscriptionHandleResponse(resp *http.Response) (PartnerDestinationsClientListBySubscriptionResponse, error) {
	result := PartnerDestinationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerDestinationsListResult); err != nil {
		return PartnerDestinationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Asynchronously updates a partner destination with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerDestinationName - Name of the partner destination.
//   - partnerDestinationUpdateParameters - Partner destination update information.
//   - options - PartnerDestinationsClientBeginUpdateOptions contains the optional parameters for the PartnerDestinationsClient.BeginUpdate
//     method.
func (client *PartnerDestinationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestinationUpdateParameters PartnerDestinationUpdateParameters, options *PartnerDestinationsClientBeginUpdateOptions) (*runtime.Poller[PartnerDestinationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, partnerDestinationName, partnerDestinationUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PartnerDestinationsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PartnerDestinationsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Asynchronously updates a partner destination with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *PartnerDestinationsClient) update(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestinationUpdateParameters PartnerDestinationUpdateParameters, options *PartnerDestinationsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, partnerDestinationName, partnerDestinationUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *PartnerDestinationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, partnerDestinationName string, partnerDestinationUpdateParameters PartnerDestinationUpdateParameters, options *PartnerDestinationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerDestinations/{partnerDestinationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerDestinationName == "" {
		return nil, errors.New("parameter partnerDestinationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerDestinationName}", url.PathEscape(partnerDestinationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, partnerDestinationUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
