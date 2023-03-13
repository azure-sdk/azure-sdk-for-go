//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurestackhci

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SKUsClient contains the methods for the SKUs group.
// Don't use this type directly, use NewSKUsClient() instead.
type SKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSKUsClient creates a new instance of SKUsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SKUsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get SKU resource details within a offer of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - publisherName - The name of the publisher available within HCI cluster.
//   - offerName - The name of the offer available within HCI cluster.
//   - skuName - The name of the SKU available within HCI cluster.
//   - options - SKUsClientGetOptions contains the optional parameters for the SKUsClient.Get method.
func (client *SKUsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, publisherName string, offerName string, skuName string, options *SKUsClientGetOptions) (SKUsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, publisherName, offerName, skuName, options)
	if err != nil {
		return SKUsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SKUsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SKUsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SKUsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, publisherName string, offerName string, skuName string, options *SKUsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/publishers/{publisherName}/offers/{offerName}/skus/{skuName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offerName == "" {
		return nil, errors.New("parameter offerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerName}", url.PathEscape(offerName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SKUsClient) getHandleResponse(resp *http.Response) (SKUsClientGetResponse, error) {
	result := SKUsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKU); err != nil {
		return SKUsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByOfferPager - List Skus available for a offer within the HCI Cluster.
//
// Generated from API version 2022-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - publisherName - The name of the publisher available within HCI cluster.
//   - offerName - The name of the offer available within HCI cluster.
//   - options - SKUsClientListByOfferOptions contains the optional parameters for the SKUsClient.NewListByOfferPager method.
func (client *SKUsClient) NewListByOfferPager(resourceGroupName string, clusterName string, publisherName string, offerName string, options *SKUsClientListByOfferOptions) *runtime.Pager[SKUsClientListByOfferResponse] {
	return runtime.NewPager(runtime.PagingHandler[SKUsClientListByOfferResponse]{
		More: func(page SKUsClientListByOfferResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SKUsClientListByOfferResponse) (SKUsClientListByOfferResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByOfferCreateRequest(ctx, resourceGroupName, clusterName, publisherName, offerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SKUsClientListByOfferResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SKUsClientListByOfferResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SKUsClientListByOfferResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByOfferHandleResponse(resp)
		},
	})
}

// listByOfferCreateRequest creates the ListByOffer request.
func (client *SKUsClient) listByOfferCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, publisherName string, offerName string, options *SKUsClientListByOfferOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/publishers/{publisherName}/offers/{offerName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offerName == "" {
		return nil, errors.New("parameter offerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerName}", url.PathEscape(offerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByOfferHandleResponse handles the ListByOffer response.
func (client *SKUsClient) listByOfferHandleResponse(resp *http.Response) (SKUsClientListByOfferResponse, error) {
	result := SKUsClientListByOfferResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUList); err != nil {
		return SKUsClientListByOfferResponse{}, err
	}
	return result, nil
}
