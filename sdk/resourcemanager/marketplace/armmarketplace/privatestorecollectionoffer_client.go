//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace

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

// PrivateStoreCollectionOfferClient contains the methods for the PrivateStoreCollectionOffer group.
// Don't use this type directly, use NewPrivateStoreCollectionOfferClient() instead.
type PrivateStoreCollectionOfferClient struct {
	internal *arm.Client
}

// NewPrivateStoreCollectionOfferClient creates a new instance of PrivateStoreCollectionOfferClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateStoreCollectionOfferClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateStoreCollectionOfferClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateStoreCollectionOfferClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Update or add an offer to a specific collection of the private store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - privateStoreID - The store ID - must use the tenant ID
//   - collectionID - The collection ID
//   - offerID - The offer ID to update or delete
//   - options - PrivateStoreCollectionOfferClientCreateOrUpdateOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.CreateOrUpdate
//     method.
func (client *PrivateStoreCollectionOfferClient) CreateOrUpdate(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientCreateOrUpdateOptions) (PrivateStoreCollectionOfferClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PrivateStoreCollectionOfferClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, privateStoreID, collectionID, offerID, options)
	if err != nil {
		return PrivateStoreCollectionOfferClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateStoreCollectionOfferClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateStoreCollectionOfferClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateStoreCollectionOfferClient) createOrUpdateCreateRequest(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Marketplace/privateStores/{privateStoreId}/collections/{collectionId}/offers/{offerId}"
	if privateStoreID == "" {
		return nil, errors.New("parameter privateStoreID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateStoreId}", url.PathEscape(privateStoreID))
	if collectionID == "" {
		return nil, errors.New("parameter collectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionId}", url.PathEscape(collectionID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Payload != nil {
		if err := runtime.MarshalAsJSON(req, *options.Payload); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrivateStoreCollectionOfferClient) createOrUpdateHandleResponse(resp *http.Response) (PrivateStoreCollectionOfferClientCreateOrUpdateResponse, error) {
	result := PrivateStoreCollectionOfferClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Offer); err != nil {
		return PrivateStoreCollectionOfferClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an offer from the given collection of private store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - privateStoreID - The store ID - must use the tenant ID
//   - collectionID - The collection ID
//   - offerID - The offer ID to update or delete
//   - options - PrivateStoreCollectionOfferClientDeleteOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Delete
//     method.
func (client *PrivateStoreCollectionOfferClient) Delete(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientDeleteOptions) (PrivateStoreCollectionOfferClientDeleteResponse, error) {
	var err error
	const operationName = "PrivateStoreCollectionOfferClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, privateStoreID, collectionID, offerID, options)
	if err != nil {
		return PrivateStoreCollectionOfferClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateStoreCollectionOfferClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PrivateStoreCollectionOfferClientDeleteResponse{}, err
	}
	return PrivateStoreCollectionOfferClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateStoreCollectionOfferClient) deleteCreateRequest(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Marketplace/privateStores/{privateStoreId}/collections/{collectionId}/offers/{offerId}"
	if privateStoreID == "" {
		return nil, errors.New("parameter privateStoreID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateStoreId}", url.PathEscape(privateStoreID))
	if collectionID == "" {
		return nil, errors.New("parameter collectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionId}", url.PathEscape(collectionID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about a specific offer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - privateStoreID - The store ID - must use the tenant ID
//   - collectionID - The collection ID
//   - offerID - The offer ID to update or delete
//   - options - PrivateStoreCollectionOfferClientGetOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Get
//     method.
func (client *PrivateStoreCollectionOfferClient) Get(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientGetOptions) (PrivateStoreCollectionOfferClientGetResponse, error) {
	var err error
	const operationName = "PrivateStoreCollectionOfferClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, privateStoreID, collectionID, offerID, options)
	if err != nil {
		return PrivateStoreCollectionOfferClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateStoreCollectionOfferClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateStoreCollectionOfferClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateStoreCollectionOfferClient) getCreateRequest(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Marketplace/privateStores/{privateStoreId}/collections/{collectionId}/offers/{offerId}"
	if privateStoreID == "" {
		return nil, errors.New("parameter privateStoreID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateStoreId}", url.PathEscape(privateStoreID))
	if collectionID == "" {
		return nil, errors.New("parameter collectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionId}", url.PathEscape(collectionID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateStoreCollectionOfferClient) getHandleResponse(resp *http.Response) (PrivateStoreCollectionOfferClientGetResponse, error) {
	result := PrivateStoreCollectionOfferClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Offer); err != nil {
		return PrivateStoreCollectionOfferClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of all private offers in the given private store and collection
//
// Generated from API version 2023-01-01
//   - privateStoreID - The store ID - must use the tenant ID
//   - collectionID - The collection ID
//   - options - PrivateStoreCollectionOfferClientListOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.NewListPager
//     method.
func (client *PrivateStoreCollectionOfferClient) NewListPager(privateStoreID string, collectionID string, options *PrivateStoreCollectionOfferClientListOptions) *runtime.Pager[PrivateStoreCollectionOfferClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateStoreCollectionOfferClientListResponse]{
		More: func(page PrivateStoreCollectionOfferClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateStoreCollectionOfferClientListResponse) (PrivateStoreCollectionOfferClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateStoreCollectionOfferClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, privateStoreID, collectionID, options)
			}, nil)
			if err != nil {
				return PrivateStoreCollectionOfferClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PrivateStoreCollectionOfferClient) listCreateRequest(ctx context.Context, privateStoreID string, collectionID string, options *PrivateStoreCollectionOfferClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Marketplace/privateStores/{privateStoreId}/collections/{collectionId}/offers"
	if privateStoreID == "" {
		return nil, errors.New("parameter privateStoreID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateStoreId}", url.PathEscape(privateStoreID))
	if collectionID == "" {
		return nil, errors.New("parameter collectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionId}", url.PathEscape(collectionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateStoreCollectionOfferClient) listHandleResponse(resp *http.Response) (PrivateStoreCollectionOfferClientListResponse, error) {
	result := PrivateStoreCollectionOfferClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OfferListResponse); err != nil {
		return PrivateStoreCollectionOfferClientListResponse{}, err
	}
	return result, nil
}

// NewListByContextsPager - Get a list of all offers in the given collection according to the required contexts.
//
// Generated from API version 2023-01-01
//   - privateStoreID - The store ID - must use the tenant ID
//   - collectionID - The collection ID
//   - options - PrivateStoreCollectionOfferClientListByContextsOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.NewListByContextsPager
//     method.
func (client *PrivateStoreCollectionOfferClient) NewListByContextsPager(privateStoreID string, collectionID string, options *PrivateStoreCollectionOfferClientListByContextsOptions) *runtime.Pager[PrivateStoreCollectionOfferClientListByContextsResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateStoreCollectionOfferClientListByContextsResponse]{
		More: func(page PrivateStoreCollectionOfferClientListByContextsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateStoreCollectionOfferClientListByContextsResponse) (PrivateStoreCollectionOfferClientListByContextsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateStoreCollectionOfferClient.NewListByContextsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByContextsCreateRequest(ctx, privateStoreID, collectionID, options)
			}, nil)
			if err != nil {
				return PrivateStoreCollectionOfferClientListByContextsResponse{}, err
			}
			return client.listByContextsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByContextsCreateRequest creates the ListByContexts request.
func (client *PrivateStoreCollectionOfferClient) listByContextsCreateRequest(ctx context.Context, privateStoreID string, collectionID string, options *PrivateStoreCollectionOfferClientListByContextsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Marketplace/privateStores/{privateStoreId}/collections/{collectionId}/mapOffersToContexts"
	if privateStoreID == "" {
		return nil, errors.New("parameter privateStoreID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateStoreId}", url.PathEscape(privateStoreID))
	if collectionID == "" {
		return nil, errors.New("parameter collectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionId}", url.PathEscape(collectionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Payload != nil {
		if err := runtime.MarshalAsJSON(req, *options.Payload); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// listByContextsHandleResponse handles the ListByContexts response.
func (client *PrivateStoreCollectionOfferClient) listByContextsHandleResponse(resp *http.Response) (PrivateStoreCollectionOfferClientListByContextsResponse, error) {
	result := PrivateStoreCollectionOfferClientListByContextsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CollectionOffersByContextList); err != nil {
		return PrivateStoreCollectionOfferClientListByContextsResponse{}, err
	}
	return result, nil
}

// Post - Delete Private store offer. This is a workaround.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - privateStoreID - The store ID - must use the tenant ID
//   - collectionID - The collection ID
//   - offerID - The offer ID to update or delete
//   - options - PrivateStoreCollectionOfferClientPostOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Post
//     method.
func (client *PrivateStoreCollectionOfferClient) Post(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientPostOptions) (PrivateStoreCollectionOfferClientPostResponse, error) {
	var err error
	const operationName = "PrivateStoreCollectionOfferClient.Post"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postCreateRequest(ctx, privateStoreID, collectionID, offerID, options)
	if err != nil {
		return PrivateStoreCollectionOfferClientPostResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateStoreCollectionOfferClientPostResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateStoreCollectionOfferClientPostResponse{}, err
	}
	return PrivateStoreCollectionOfferClientPostResponse{}, nil
}

// postCreateRequest creates the Post request.
func (client *PrivateStoreCollectionOfferClient) postCreateRequest(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientPostOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Marketplace/privateStores/{privateStoreId}/collections/{collectionId}/offers/{offerId}"
	if privateStoreID == "" {
		return nil, errors.New("parameter privateStoreID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateStoreId}", url.PathEscape(privateStoreID))
	if collectionID == "" {
		return nil, errors.New("parameter collectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionId}", url.PathEscape(collectionID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Payload != nil {
		if err := runtime.MarshalAsJSON(req, *options.Payload); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// UpsertOfferWithMultiContext - Upsert an offer with multiple context details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - privateStoreID - The store ID - must use the tenant ID
//   - collectionID - The collection ID
//   - offerID - The offer ID to update or delete
//   - options - PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextOptions contains the optional parameters for the
//     PrivateStoreCollectionOfferClient.UpsertOfferWithMultiContext method.
func (client *PrivateStoreCollectionOfferClient) UpsertOfferWithMultiContext(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextOptions) (PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse, error) {
	var err error
	const operationName = "PrivateStoreCollectionOfferClient.UpsertOfferWithMultiContext"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.upsertOfferWithMultiContextCreateRequest(ctx, privateStoreID, collectionID, offerID, options)
	if err != nil {
		return PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse{}, err
	}
	resp, err := client.upsertOfferWithMultiContextHandleResponse(httpResp)
	return resp, err
}

// upsertOfferWithMultiContextCreateRequest creates the UpsertOfferWithMultiContext request.
func (client *PrivateStoreCollectionOfferClient) upsertOfferWithMultiContextCreateRequest(ctx context.Context, privateStoreID string, collectionID string, offerID string, options *PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Marketplace/privateStores/{privateStoreId}/collections/{collectionId}/offers/{offerId}/upsertOfferWithMultiContext"
	if privateStoreID == "" {
		return nil, errors.New("parameter privateStoreID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateStoreId}", url.PathEscape(privateStoreID))
	if collectionID == "" {
		return nil, errors.New("parameter collectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionId}", url.PathEscape(collectionID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Payload != nil {
		if err := runtime.MarshalAsJSON(req, *options.Payload); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// upsertOfferWithMultiContextHandleResponse handles the UpsertOfferWithMultiContext response.
func (client *PrivateStoreCollectionOfferClient) upsertOfferWithMultiContextHandleResponse(resp *http.Response) (PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse, error) {
	result := PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Offer); err != nil {
		return PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse{}, err
	}
	return result, nil
}
