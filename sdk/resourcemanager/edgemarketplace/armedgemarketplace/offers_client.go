//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgemarketplace

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

// OffersClient contains the methods for the Offers group.
// Don't use this type directly, use NewOffersClient() instead.
type OffersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOffersClient creates a new instance of OffersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOffersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OffersClient, error) {
	cl, err := arm.NewClient(moduleName+".OffersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OffersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginGenerateAccessToken - A long-running resource action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - offerID - Id of the offer
//   - body - The content of the action request
//   - options - OffersClientBeginGenerateAccessTokenOptions contains the optional parameters for the OffersClient.BeginGenerateAccessToken
//     method.
func (client *OffersClient) BeginGenerateAccessToken(ctx context.Context, resourceURI string, offerID string, body AccessTokenRequest, options *OffersClientBeginGenerateAccessTokenOptions) (*runtime.Poller[OffersClientGenerateAccessTokenResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.generateAccessToken(ctx, resourceURI, offerID, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OffersClientGenerateAccessTokenResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[OffersClientGenerateAccessTokenResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// GenerateAccessToken - A long-running resource action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *OffersClient) generateAccessToken(ctx context.Context, resourceURI string, offerID string, body AccessTokenRequest, options *OffersClientBeginGenerateAccessTokenOptions) (*http.Response, error) {
	var err error
	req, err := client.generateAccessTokenCreateRequest(ctx, resourceURI, offerID, body, options)
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

// generateAccessTokenCreateRequest creates the GenerateAccessToken request.
func (client *OffersClient) generateAccessTokenCreateRequest(ctx context.Context, resourceURI string, offerID string, body AccessTokenRequest, options *OffersClientBeginGenerateAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.EdgeMarketplace/offers/{offerId}/generateAccessToken"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a Offer
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - offerID - Id of the offer
//   - options - OffersClientGetOptions contains the optional parameters for the OffersClient.Get method.
func (client *OffersClient) Get(ctx context.Context, resourceURI string, offerID string, options *OffersClientGetOptions) (OffersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceURI, offerID, options)
	if err != nil {
		return OffersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OffersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OffersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OffersClient) getCreateRequest(ctx context.Context, resourceURI string, offerID string, options *OffersClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.EdgeMarketplace/offers/{offerId}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OffersClient) getHandleResponse(resp *http.Response) (OffersClientGetResponse, error) {
	result := OffersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Offer); err != nil {
		return OffersClientGetResponse{}, err
	}
	return result, nil
}

// GetAccessToken - get access token.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - offerID - Id of the offer
//   - body - The content of the action request
//   - options - OffersClientGetAccessTokenOptions contains the optional parameters for the OffersClient.GetAccessToken method.
func (client *OffersClient) GetAccessToken(ctx context.Context, resourceURI string, offerID string, body AccessTokenReadRequest, options *OffersClientGetAccessTokenOptions) (OffersClientGetAccessTokenResponse, error) {
	var err error
	req, err := client.getAccessTokenCreateRequest(ctx, resourceURI, offerID, body, options)
	if err != nil {
		return OffersClientGetAccessTokenResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OffersClientGetAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OffersClientGetAccessTokenResponse{}, err
	}
	resp, err := client.getAccessTokenHandleResponse(httpResp)
	return resp, err
}

// getAccessTokenCreateRequest creates the GetAccessToken request.
func (client *OffersClient) getAccessTokenCreateRequest(ctx context.Context, resourceURI string, offerID string, body AccessTokenReadRequest, options *OffersClientGetAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.EdgeMarketplace/offers/{offerId}/getAccessToken"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getAccessTokenHandleResponse handles the GetAccessToken response.
func (client *OffersClient) getAccessTokenHandleResponse(resp *http.Response) (OffersClientGetAccessTokenResponse, error) {
	result := OffersClientGetAccessTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskAccessToken); err != nil {
		return OffersClientGetAccessTokenResponse{}, err
	}
	return result, nil
}

// NewListPager - List Offer resources by parent
//
// Generated from API version 2023-08-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - OffersClientListOptions contains the optional parameters for the OffersClient.NewListPager method.
func (client *OffersClient) NewListPager(resourceURI string, options *OffersClientListOptions) *runtime.Pager[OffersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OffersClientListResponse]{
		More: func(page OffersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OffersClientListResponse) (OffersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceURI, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OffersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OffersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OffersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OffersClient) listCreateRequest(ctx context.Context, resourceURI string, options *OffersClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.EdgeMarketplace/offers"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OffersClient) listHandleResponse(resp *http.Response) (OffersClientListResponse, error) {
	result := OffersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OfferListResult); err != nil {
		return OffersClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List Offer resources by subscription
//
// Generated from API version 2023-08-01-preview
//   - options - OffersClientListBySubscriptionOptions contains the optional parameters for the OffersClient.NewListBySubscriptionPager
//     method.
func (client *OffersClient) NewListBySubscriptionPager(options *OffersClientListBySubscriptionOptions) *runtime.Pager[OffersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[OffersClientListBySubscriptionResponse]{
		More: func(page OffersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OffersClientListBySubscriptionResponse) (OffersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OffersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OffersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OffersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OffersClient) listBySubscriptionCreateRequest(ctx context.Context, options *OffersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeMarketplace/offers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *OffersClient) listBySubscriptionHandleResponse(resp *http.Response) (OffersClientListBySubscriptionResponse, error) {
	result := OffersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OfferListResult); err != nil {
		return OffersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
