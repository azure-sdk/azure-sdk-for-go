//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcommunication

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

// SenderUsernamesClient contains the methods for the SenderUsernames group.
// Don't use this type directly, use NewSenderUsernamesClient() instead.
type SenderUsernamesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSenderUsernamesClient creates a new instance of SenderUsernamesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSenderUsernamesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SenderUsernamesClient, error) {
	cl, err := arm.NewClient(moduleName+".SenderUsernamesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SenderUsernamesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Add a new SenderUsername resource under the parent Domains resource or update an existing SenderUsername
// resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - emailServiceName - The name of the EmailService resource.
//   - domainName - The name of the Domains resource.
//   - senderUsername - The valid sender Username.
//   - parameters - Parameters for the create or update operation
//   - options - SenderUsernamesClientCreateOrUpdateOptions contains the optional parameters for the SenderUsernamesClient.CreateOrUpdate
//     method.
func (client *SenderUsernamesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, senderUsername string, parameters SenderUsernameResource, options *SenderUsernamesClientCreateOrUpdateOptions) (SenderUsernamesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, senderUsername, parameters, options)
	if err != nil {
		return SenderUsernamesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SenderUsernamesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SenderUsernamesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SenderUsernamesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, senderUsername string, parameters SenderUsernameResource, options *SenderUsernamesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}/senderUsernames/{senderUsername}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if senderUsername == "" {
		return nil, errors.New("parameter senderUsername cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{senderUsername}", url.PathEscape(senderUsername))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SenderUsernamesClient) createOrUpdateHandleResponse(resp *http.Response) (SenderUsernamesClientCreateOrUpdateResponse, error) {
	result := SenderUsernamesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SenderUsernameResource); err != nil {
		return SenderUsernamesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Operation to delete a SenderUsernames resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - emailServiceName - The name of the EmailService resource.
//   - domainName - The name of the Domains resource.
//   - senderUsername - The valid sender Username.
//   - options - SenderUsernamesClientDeleteOptions contains the optional parameters for the SenderUsernamesClient.Delete method.
func (client *SenderUsernamesClient) Delete(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, senderUsername string, options *SenderUsernamesClientDeleteOptions) (SenderUsernamesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, senderUsername, options)
	if err != nil {
		return SenderUsernamesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SenderUsernamesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SenderUsernamesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SenderUsernamesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SenderUsernamesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, senderUsername string, options *SenderUsernamesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}/senderUsernames/{senderUsername}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if senderUsername == "" {
		return nil, errors.New("parameter senderUsername cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{senderUsername}", url.PathEscape(senderUsername))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a valid sender username for a domains resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - emailServiceName - The name of the EmailService resource.
//   - domainName - The name of the Domains resource.
//   - senderUsername - The valid sender Username.
//   - options - SenderUsernamesClientGetOptions contains the optional parameters for the SenderUsernamesClient.Get method.
func (client *SenderUsernamesClient) Get(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, senderUsername string, options *SenderUsernamesClientGetOptions) (SenderUsernamesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, senderUsername, options)
	if err != nil {
		return SenderUsernamesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SenderUsernamesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SenderUsernamesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SenderUsernamesClient) getCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, senderUsername string, options *SenderUsernamesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}/senderUsernames/{senderUsername}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if senderUsername == "" {
		return nil, errors.New("parameter senderUsername cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{senderUsername}", url.PathEscape(senderUsername))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SenderUsernamesClient) getHandleResponse(resp *http.Response) (SenderUsernamesClientGetResponse, error) {
	result := SenderUsernamesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SenderUsernameResource); err != nil {
		return SenderUsernamesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDomainsPager - List all valid sender usernames for a domains resource.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - emailServiceName - The name of the EmailService resource.
//   - domainName - The name of the Domains resource.
//   - options - SenderUsernamesClientListByDomainsOptions contains the optional parameters for the SenderUsernamesClient.NewListByDomainsPager
//     method.
func (client *SenderUsernamesClient) NewListByDomainsPager(resourceGroupName string, emailServiceName string, domainName string, options *SenderUsernamesClientListByDomainsOptions) *runtime.Pager[SenderUsernamesClientListByDomainsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SenderUsernamesClientListByDomainsResponse]{
		More: func(page SenderUsernamesClientListByDomainsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SenderUsernamesClientListByDomainsResponse) (SenderUsernamesClientListByDomainsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDomainsCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SenderUsernamesClientListByDomainsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SenderUsernamesClientListByDomainsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SenderUsernamesClientListByDomainsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDomainsHandleResponse(resp)
		},
	})
}

// listByDomainsCreateRequest creates the ListByDomains request.
func (client *SenderUsernamesClient) listByDomainsCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *SenderUsernamesClientListByDomainsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}/senderUsernames"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDomainsHandleResponse handles the ListByDomains response.
func (client *SenderUsernamesClient) listByDomainsHandleResponse(resp *http.Response) (SenderUsernamesClientListByDomainsResponse, error) {
	result := SenderUsernamesClientListByDomainsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SenderUsernameResourceCollection); err != nil {
		return SenderUsernamesClientListByDomainsResponse{}, err
	}
	return result, nil
}
