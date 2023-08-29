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

// DomainTopicsClient contains the methods for the DomainTopics group.
// Don't use this type directly, use NewDomainTopicsClient() instead.
type DomainTopicsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDomainTopicsClient creates a new instance of DomainTopicsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDomainTopicsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DomainTopicsClient, error) {
	cl, err := arm.NewClient(moduleName+".DomainTopicsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DomainTopicsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates or updates a new domain topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - domainName - Name of the domain.
//   - domainTopicName - Name of the domain topic.
//   - options - DomainTopicsClientBeginCreateOrUpdateOptions contains the optional parameters for the DomainTopicsClient.BeginCreateOrUpdate
//     method.
func (client *DomainTopicsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DomainTopicsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, domainName, domainTopicName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DomainTopicsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DomainTopicsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Asynchronously creates or updates a new domain topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *DomainTopicsClient) createOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, domainName, domainTopicName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DomainTopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if domainTopicName == "" {
		return nil, errors.New("parameter domainTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainTopicName}", url.PathEscape(domainTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Delete existing domain topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - domainName - Name of the domain.
//   - domainTopicName - Name of the domain topic.
//   - options - DomainTopicsClientBeginDeleteOptions contains the optional parameters for the DomainTopicsClient.BeginDelete
//     method.
func (client *DomainTopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientBeginDeleteOptions) (*runtime.Poller[DomainTopicsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, domainName, domainTopicName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DomainTopicsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DomainTopicsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete existing domain topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *DomainTopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainName, domainTopicName, options)
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
func (client *DomainTopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if domainTopicName == "" {
		return nil, errors.New("parameter domainTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainTopicName}", url.PathEscape(domainTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of a domain topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - domainName - Name of the domain.
//   - domainTopicName - Name of the topic.
//   - options - DomainTopicsClientGetOptions contains the optional parameters for the DomainTopicsClient.Get method.
func (client *DomainTopicsClient) Get(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientGetOptions) (DomainTopicsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainName, domainTopicName, options)
	if err != nil {
		return DomainTopicsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DomainTopicsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DomainTopicsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DomainTopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string, options *DomainTopicsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if domainTopicName == "" {
		return nil, errors.New("parameter domainTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainTopicName}", url.PathEscape(domainTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainTopicsClient) getHandleResponse(resp *http.Response) (DomainTopicsClientGetResponse, error) {
	result := DomainTopicsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainTopic); err != nil {
		return DomainTopicsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDomainPager - List all the topics in a domain.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - domainName - Domain name.
//   - options - DomainTopicsClientListByDomainOptions contains the optional parameters for the DomainTopicsClient.NewListByDomainPager
//     method.
func (client *DomainTopicsClient) NewListByDomainPager(resourceGroupName string, domainName string, options *DomainTopicsClientListByDomainOptions) *runtime.Pager[DomainTopicsClientListByDomainResponse] {
	return runtime.NewPager(runtime.PagingHandler[DomainTopicsClientListByDomainResponse]{
		More: func(page DomainTopicsClientListByDomainResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DomainTopicsClientListByDomainResponse) (DomainTopicsClientListByDomainResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDomainCreateRequest(ctx, resourceGroupName, domainName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DomainTopicsClientListByDomainResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DomainTopicsClientListByDomainResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DomainTopicsClientListByDomainResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDomainHandleResponse(resp)
		},
	})
}

// listByDomainCreateRequest creates the ListByDomain request.
func (client *DomainTopicsClient) listByDomainCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainTopicsClientListByDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
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

// listByDomainHandleResponse handles the ListByDomain response.
func (client *DomainTopicsClient) listByDomainHandleResponse(resp *http.Response) (DomainTopicsClientListByDomainResponse, error) {
	result := DomainTopicsClientListByDomainResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainTopicsListResult); err != nil {
		return DomainTopicsClientListByDomainResponse{}, err
	}
	return result, nil
}
