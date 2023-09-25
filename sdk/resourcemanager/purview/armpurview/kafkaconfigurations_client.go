//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

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

// KafkaConfigurationsClient contains the methods for the KafkaConfigurations group.
// Don't use this type directly, use NewKafkaConfigurationsClient() instead.
type KafkaConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKafkaConfigurationsClient creates a new instance of KafkaConfigurationsClient with the specified values.
//   - subscriptionID - The subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKafkaConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KafkaConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".KafkaConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KafkaConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update Kafka Configuration
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - kafkaConfigurationName - The kafka configuration name.
//   - kafkaConfiguration - The kafka configuration of the account.
//   - options - KafkaConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the KafkaConfigurationsClient.CreateOrUpdate
//     method.
func (client *KafkaConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, kafkaConfigurationName string, kafkaConfiguration KafkaConfiguration, options *KafkaConfigurationsClientCreateOrUpdateOptions) (KafkaConfigurationsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, kafkaConfigurationName, kafkaConfiguration, options)
	if err != nil {
		return KafkaConfigurationsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KafkaConfigurationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return KafkaConfigurationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *KafkaConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, kafkaConfigurationName string, kafkaConfiguration KafkaConfiguration, options *KafkaConfigurationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/kafkaConfigurations/{kafkaConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if kafkaConfigurationName == "" {
		return nil, errors.New("parameter kafkaConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kafkaConfigurationName}", url.PathEscape(kafkaConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, kafkaConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *KafkaConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (KafkaConfigurationsClientCreateOrUpdateResponse, error) {
	result := KafkaConfigurationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KafkaConfiguration); err != nil {
		return KafkaConfigurationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a KafkaConfiguration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - kafkaConfigurationName - Name of kafka configuration.
//   - options - KafkaConfigurationsClientDeleteOptions contains the optional parameters for the KafkaConfigurationsClient.Delete
//     method.
func (client *KafkaConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, kafkaConfigurationName string, options *KafkaConfigurationsClientDeleteOptions) (KafkaConfigurationsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, kafkaConfigurationName, options)
	if err != nil {
		return KafkaConfigurationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KafkaConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return KafkaConfigurationsClientDeleteResponse{}, err
	}
	return KafkaConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *KafkaConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, kafkaConfigurationName string, options *KafkaConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/kafkaConfigurations/{kafkaConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if kafkaConfigurationName == "" {
		return nil, errors.New("parameter kafkaConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kafkaConfigurationName}", url.PathEscape(kafkaConfigurationName))
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

// Get - Gets the kafka configuration for the account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - kafkaConfigurationName - Name of kafka configuration.
//   - options - KafkaConfigurationsClientGetOptions contains the optional parameters for the KafkaConfigurationsClient.Get method.
func (client *KafkaConfigurationsClient) Get(ctx context.Context, resourceGroupName string, accountName string, kafkaConfigurationName string, options *KafkaConfigurationsClientGetOptions) (KafkaConfigurationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, kafkaConfigurationName, options)
	if err != nil {
		return KafkaConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KafkaConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return KafkaConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *KafkaConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, kafkaConfigurationName string, options *KafkaConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/kafkaConfigurations/{kafkaConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if kafkaConfigurationName == "" {
		return nil, errors.New("parameter kafkaConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kafkaConfigurationName}", url.PathEscape(kafkaConfigurationName))
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
func (client *KafkaConfigurationsClient) getHandleResponse(resp *http.Response) (KafkaConfigurationsClientGetResponse, error) {
	result := KafkaConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KafkaConfiguration); err != nil {
		return KafkaConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - Lists the Kafka configurations in the Account
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the account.
//   - options - KafkaConfigurationsClientListByAccountOptions contains the optional parameters for the KafkaConfigurationsClient.NewListByAccountPager
//     method.
func (client *KafkaConfigurationsClient) NewListByAccountPager(resourceGroupName string, accountName string, options *KafkaConfigurationsClientListByAccountOptions) *runtime.Pager[KafkaConfigurationsClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[KafkaConfigurationsClientListByAccountResponse]{
		More: func(page KafkaConfigurationsClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KafkaConfigurationsClientListByAccountResponse) (KafkaConfigurationsClientListByAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return KafkaConfigurationsClientListByAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return KafkaConfigurationsClientListByAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return KafkaConfigurationsClientListByAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAccountHandleResponse(resp)
		},
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *KafkaConfigurationsClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *KafkaConfigurationsClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/kafkaConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *KafkaConfigurationsClient) listByAccountHandleResponse(resp *http.Response) (KafkaConfigurationsClientListByAccountResponse, error) {
	result := KafkaConfigurationsClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KafkaConfigurationList); err != nil {
		return KafkaConfigurationsClientListByAccountResponse{}, err
	}
	return result, nil
}
