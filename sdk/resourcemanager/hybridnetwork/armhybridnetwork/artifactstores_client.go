//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// ArtifactStoresClient contains the methods for the ArtifactStores group.
// Don't use this type directly, use NewArtifactStoresClient() instead.
type ArtifactStoresClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewArtifactStoresClient creates a new instance of ArtifactStoresClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewArtifactStoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArtifactStoresClient, error) {
	cl, err := arm.NewClient(moduleName+".ArtifactStoresClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ArtifactStoresClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a artifact store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - artifactStoreName - The name of the artifact store.
//   - parameters - Parameters supplied to the create or update application group operation.
//   - options - ArtifactStoresClientBeginCreateOrUpdateOptions contains the optional parameters for the ArtifactStoresClient.BeginCreateOrUpdate
//     method.
func (client *ArtifactStoresClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, parameters ArtifactStore, options *ArtifactStoresClientBeginCreateOrUpdateOptions) (*runtime.Poller[ArtifactStoresClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, publisherName, artifactStoreName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ArtifactStoresClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ArtifactStoresClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a artifact store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *ArtifactStoresClient) createOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, parameters ArtifactStore, options *ArtifactStoresClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publisherName, artifactStoreName, parameters, options)
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
func (client *ArtifactStoresClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, parameters ArtifactStore, options *ArtifactStoresClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores/{artifactStoreName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if artifactStoreName == "" {
		return nil, errors.New("parameter artifactStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactStoreName}", url.PathEscape(artifactStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified artifact store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - artifactStoreName - The name of the artifact store.
//   - options - ArtifactStoresClientBeginDeleteOptions contains the optional parameters for the ArtifactStoresClient.BeginDelete
//     method.
func (client *ArtifactStoresClient) BeginDelete(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, options *ArtifactStoresClientBeginDeleteOptions) (*runtime.Poller[ArtifactStoresClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, publisherName, artifactStoreName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ArtifactStoresClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ArtifactStoresClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified artifact store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *ArtifactStoresClient) deleteOperation(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, options *ArtifactStoresClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publisherName, artifactStoreName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ArtifactStoresClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, options *ArtifactStoresClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores/{artifactStoreName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if artifactStoreName == "" {
		return nil, errors.New("parameter artifactStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactStoreName}", url.PathEscape(artifactStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// Get - Gets information about the specified artifact store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - artifactStoreName - The name of the artifact store.
//   - options - ArtifactStoresClientGetOptions contains the optional parameters for the ArtifactStoresClient.Get method.
func (client *ArtifactStoresClient) Get(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, options *ArtifactStoresClientGetOptions) (ArtifactStoresClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, publisherName, artifactStoreName, options)
	if err != nil {
		return ArtifactStoresClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactStoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArtifactStoresClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ArtifactStoresClient) getCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, options *ArtifactStoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores/{artifactStoreName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if artifactStoreName == "" {
		return nil, errors.New("parameter artifactStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactStoreName}", url.PathEscape(artifactStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *ArtifactStoresClient) getHandleResponse(resp *http.Response) (ArtifactStoresClientGetResponse, error) {
	result := ArtifactStoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactStore); err != nil {
		return ArtifactStoresClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPublisherPager - Gets information of the ArtifactStores under publisher.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - options - ArtifactStoresClientListByPublisherOptions contains the optional parameters for the ArtifactStoresClient.NewListByPublisherPager
//     method.
func (client *ArtifactStoresClient) NewListByPublisherPager(resourceGroupName string, publisherName string, options *ArtifactStoresClientListByPublisherOptions) *runtime.Pager[ArtifactStoresClientListByPublisherResponse] {
	return runtime.NewPager(runtime.PagingHandler[ArtifactStoresClientListByPublisherResponse]{
		More: func(page ArtifactStoresClientListByPublisherResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArtifactStoresClientListByPublisherResponse) (ArtifactStoresClientListByPublisherResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPublisherCreateRequest(ctx, resourceGroupName, publisherName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ArtifactStoresClientListByPublisherResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ArtifactStoresClientListByPublisherResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ArtifactStoresClientListByPublisherResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPublisherHandleResponse(resp)
		},
	})
}

// listByPublisherCreateRequest creates the ListByPublisher request.
func (client *ArtifactStoresClient) listByPublisherCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, options *ArtifactStoresClientListByPublisherOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listByPublisherHandleResponse handles the ListByPublisher response.
func (client *ArtifactStoresClient) listByPublisherHandleResponse(resp *http.Response) (ArtifactStoresClientListByPublisherResponse, error) {
	result := ArtifactStoresClientListByPublisherResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactStoreListResult); err != nil {
		return ArtifactStoresClientListByPublisherResponse{}, err
	}
	return result, nil
}

// Update - Update artifact store resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - artifactStoreName - The name of the artifact store.
//   - parameters - Parameters supplied to the create or update application group operation.
//   - options - ArtifactStoresClientUpdateOptions contains the optional parameters for the ArtifactStoresClient.Update method.
func (client *ArtifactStoresClient) Update(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, parameters TagsObject, options *ArtifactStoresClientUpdateOptions) (ArtifactStoresClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, publisherName, artifactStoreName, parameters, options)
	if err != nil {
		return ArtifactStoresClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactStoresClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArtifactStoresClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ArtifactStoresClient) updateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, parameters TagsObject, options *ArtifactStoresClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores/{artifactStoreName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if artifactStoreName == "" {
		return nil, errors.New("parameter artifactStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactStoreName}", url.PathEscape(artifactStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ArtifactStoresClient) updateHandleResponse(resp *http.Response) (ArtifactStoresClientUpdateResponse, error) {
	result := ArtifactStoresClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactStore); err != nil {
		return ArtifactStoresClientUpdateResponse{}, err
	}
	return result, nil
}
