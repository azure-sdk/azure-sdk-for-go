//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommunitytraining

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

// CommunityTrainingsClient contains the methods for the CommunityTrainings group.
// Don't use this type directly, use NewCommunityTrainingsClient() instead.
type CommunityTrainingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCommunityTrainingsClient creates a new instance of CommunityTrainingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCommunityTrainingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CommunityTrainingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CommunityTrainingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a CommunityTraining
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communityTrainingName - The name of the Community Training Resource
//   - resource - Resource create parameters.
//   - options - CommunityTrainingsClientBeginCreateOptions contains the optional parameters for the CommunityTrainingsClient.BeginCreate
//     method.
func (client *CommunityTrainingsClient) BeginCreate(ctx context.Context, resourceGroupName string, communityTrainingName string, resource CommunityTraining, options *CommunityTrainingsClientBeginCreateOptions) (*runtime.Poller[CommunityTrainingsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, communityTrainingName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CommunityTrainingsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CommunityTrainingsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a CommunityTraining
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *CommunityTrainingsClient) create(ctx context.Context, resourceGroupName string, communityTrainingName string, resource CommunityTraining, options *CommunityTrainingsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "CommunityTrainingsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, communityTrainingName, resource, options)
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

// createCreateRequest creates the Create request.
func (client *CommunityTrainingsClient) createCreateRequest(ctx context.Context, resourceGroupName string, communityTrainingName string, resource CommunityTraining, options *CommunityTrainingsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Community/communityTrainings/{communityTrainingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communityTrainingName == "" {
		return nil, errors.New("parameter communityTrainingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communityTrainingName}", url.PathEscape(communityTrainingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a CommunityTraining
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communityTrainingName - The name of the Community Training Resource
//   - options - CommunityTrainingsClientBeginDeleteOptions contains the optional parameters for the CommunityTrainingsClient.BeginDelete
//     method.
func (client *CommunityTrainingsClient) BeginDelete(ctx context.Context, resourceGroupName string, communityTrainingName string, options *CommunityTrainingsClientBeginDeleteOptions) (*runtime.Poller[CommunityTrainingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, communityTrainingName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CommunityTrainingsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CommunityTrainingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a CommunityTraining
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *CommunityTrainingsClient) deleteOperation(ctx context.Context, resourceGroupName string, communityTrainingName string, options *CommunityTrainingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CommunityTrainingsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, communityTrainingName, options)
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
func (client *CommunityTrainingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, communityTrainingName string, options *CommunityTrainingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Community/communityTrainings/{communityTrainingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communityTrainingName == "" {
		return nil, errors.New("parameter communityTrainingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communityTrainingName}", url.PathEscape(communityTrainingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a CommunityTraining
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communityTrainingName - The name of the Community Training Resource
//   - options - CommunityTrainingsClientGetOptions contains the optional parameters for the CommunityTrainingsClient.Get method.
func (client *CommunityTrainingsClient) Get(ctx context.Context, resourceGroupName string, communityTrainingName string, options *CommunityTrainingsClientGetOptions) (CommunityTrainingsClientGetResponse, error) {
	var err error
	const operationName = "CommunityTrainingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, communityTrainingName, options)
	if err != nil {
		return CommunityTrainingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CommunityTrainingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CommunityTrainingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CommunityTrainingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, communityTrainingName string, options *CommunityTrainingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Community/communityTrainings/{communityTrainingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communityTrainingName == "" {
		return nil, errors.New("parameter communityTrainingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communityTrainingName}", url.PathEscape(communityTrainingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommunityTrainingsClient) getHandleResponse(resp *http.Response) (CommunityTrainingsClientGetResponse, error) {
	result := CommunityTrainingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunityTraining); err != nil {
		return CommunityTrainingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List CommunityTraining resources by resource group
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CommunityTrainingsClientListByResourceGroupOptions contains the optional parameters for the CommunityTrainingsClient.NewListByResourceGroupPager
//     method.
func (client *CommunityTrainingsClient) NewListByResourceGroupPager(resourceGroupName string, options *CommunityTrainingsClientListByResourceGroupOptions) *runtime.Pager[CommunityTrainingsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CommunityTrainingsClientListByResourceGroupResponse]{
		More: func(page CommunityTrainingsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CommunityTrainingsClientListByResourceGroupResponse) (CommunityTrainingsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CommunityTrainingsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CommunityTrainingsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CommunityTrainingsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CommunityTrainingsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Community/communityTrainings"
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
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CommunityTrainingsClient) listByResourceGroupHandleResponse(resp *http.Response) (CommunityTrainingsClientListByResourceGroupResponse, error) {
	result := CommunityTrainingsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return CommunityTrainingsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List CommunityTraining resources by subscription ID
//
// Generated from API version 2023-11-01
//   - options - CommunityTrainingsClientListBySubscriptionOptions contains the optional parameters for the CommunityTrainingsClient.NewListBySubscriptionPager
//     method.
func (client *CommunityTrainingsClient) NewListBySubscriptionPager(options *CommunityTrainingsClientListBySubscriptionOptions) *runtime.Pager[CommunityTrainingsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CommunityTrainingsClientListBySubscriptionResponse]{
		More: func(page CommunityTrainingsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CommunityTrainingsClientListBySubscriptionResponse) (CommunityTrainingsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CommunityTrainingsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CommunityTrainingsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CommunityTrainingsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CommunityTrainingsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Community/communityTrainings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CommunityTrainingsClient) listBySubscriptionHandleResponse(resp *http.Response) (CommunityTrainingsClientListBySubscriptionResponse, error) {
	result := CommunityTrainingsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return CommunityTrainingsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a CommunityTraining
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communityTrainingName - The name of the Community Training Resource
//   - properties - The resource properties to be updated.
//   - options - CommunityTrainingsClientBeginUpdateOptions contains the optional parameters for the CommunityTrainingsClient.BeginUpdate
//     method.
func (client *CommunityTrainingsClient) BeginUpdate(ctx context.Context, resourceGroupName string, communityTrainingName string, properties Update, options *CommunityTrainingsClientBeginUpdateOptions) (*runtime.Poller[CommunityTrainingsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, communityTrainingName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CommunityTrainingsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CommunityTrainingsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a CommunityTraining
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *CommunityTrainingsClient) update(ctx context.Context, resourceGroupName string, communityTrainingName string, properties Update, options *CommunityTrainingsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CommunityTrainingsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, communityTrainingName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *CommunityTrainingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, communityTrainingName string, properties Update, options *CommunityTrainingsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Community/communityTrainings/{communityTrainingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communityTrainingName == "" {
		return nil, errors.New("parameter communityTrainingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communityTrainingName}", url.PathEscape(communityTrainingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
