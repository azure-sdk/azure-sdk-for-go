//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement

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

// SubscriptionClient contains the methods for the Subscription group.
// Don't use this type directly, use NewSubscriptionClient() instead.
type SubscriptionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubscriptionClient creates a new instance of SubscriptionClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubscriptionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionClient, error) {
	cl, err := arm.NewClient(moduleName+".SubscriptionClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the subscription of specified user to the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - parameters - Create parameters.
//   - options - SubscriptionClientCreateOrUpdateOptions contains the optional parameters for the SubscriptionClient.CreateOrUpdate
//     method.
func (client *SubscriptionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters, options *SubscriptionClientCreateOrUpdateOptions) (SubscriptionClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, sid, parameters, options)
	if err != nil {
		return SubscriptionClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SubscriptionClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SubscriptionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters, options *SubscriptionClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Notify != nil {
		reqQP.Set("notify", strconv.FormatBool(*options.Notify))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SubscriptionClient) createOrUpdateHandleResponse(resp *http.Response) (SubscriptionClientCreateOrUpdateResponse, error) {
	result := SubscriptionClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionContract); err != nil {
		return SubscriptionClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - SubscriptionClientDeleteOptions contains the optional parameters for the SubscriptionClient.Delete method.
func (client *SubscriptionClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, options *SubscriptionClientDeleteOptions) (SubscriptionClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, sid, ifMatch, options)
	if err != nil {
		return SubscriptionClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SubscriptionClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SubscriptionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, options *SubscriptionClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Subscription entity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - options - SubscriptionClientGetOptions contains the optional parameters for the SubscriptionClient.Get method.
func (client *SubscriptionClient) Get(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientGetOptions) (SubscriptionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubscriptionClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *SubscriptionClient) getHandleResponse(resp *http.Response) (SubscriptionClientGetResponse, error) {
	result := SubscriptionClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionContract); err != nil {
		return SubscriptionClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the apimanagement subscription specified by its identifier.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - options - SubscriptionClientGetEntityTagOptions contains the optional parameters for the SubscriptionClient.GetEntityTag
//     method.
func (client *SubscriptionClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientGetEntityTagOptions) (SubscriptionClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionClientGetEntityTagResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *SubscriptionClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *SubscriptionClient) getEntityTagHandleResponse(resp *http.Response) (SubscriptionClientGetEntityTagResponse, error) {
	result := SubscriptionClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListPager - Lists all subscriptions of the API Management service instance.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - SubscriptionClientListOptions contains the optional parameters for the SubscriptionClient.NewListPager method.
func (client *SubscriptionClient) NewListPager(resourceGroupName string, serviceName string, options *SubscriptionClientListOptions) *runtime.Pager[SubscriptionClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubscriptionClientListResponse]{
		More: func(page SubscriptionClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionClientListResponse) (SubscriptionClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubscriptionClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SubscriptionClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubscriptionClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SubscriptionClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *SubscriptionClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubscriptionClient) listHandleResponse(resp *http.Response) (SubscriptionClientListResponse, error) {
	result := SubscriptionClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionCollection); err != nil {
		return SubscriptionClientListResponse{}, err
	}
	return result, nil
}

// ListSecrets - Gets the specified Subscription keys.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - options - SubscriptionClientListSecretsOptions contains the optional parameters for the SubscriptionClient.ListSecrets
//     method.
func (client *SubscriptionClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientListSecretsOptions) (SubscriptionClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionClientListSecretsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *SubscriptionClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/listSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *SubscriptionClient) listSecretsHandleResponse(resp *http.Response) (SubscriptionClientListSecretsResponse, error) {
	result := SubscriptionClientListSecretsResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionKeysContract); err != nil {
		return SubscriptionClientListSecretsResponse{}, err
	}
	return result, nil
}

// RegeneratePrimaryKey - Regenerates primary key of existing subscription of the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - options - SubscriptionClientRegeneratePrimaryKeyOptions contains the optional parameters for the SubscriptionClient.RegeneratePrimaryKey
//     method.
func (client *SubscriptionClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientRegeneratePrimaryKeyOptions) (SubscriptionClientRegeneratePrimaryKeyResponse, error) {
	req, err := client.regeneratePrimaryKeyCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionClientRegeneratePrimaryKeyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientRegeneratePrimaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return SubscriptionClientRegeneratePrimaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionClientRegeneratePrimaryKeyResponse{}, nil
}

// regeneratePrimaryKeyCreateRequest creates the RegeneratePrimaryKey request.
func (client *SubscriptionClient) regeneratePrimaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientRegeneratePrimaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regeneratePrimaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// RegenerateSecondaryKey - Regenerates secondary key of existing subscription of the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - options - SubscriptionClientRegenerateSecondaryKeyOptions contains the optional parameters for the SubscriptionClient.RegenerateSecondaryKey
//     method.
func (client *SubscriptionClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientRegenerateSecondaryKeyOptions) (SubscriptionClientRegenerateSecondaryKeyResponse, error) {
	req, err := client.regenerateSecondaryKeyCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionClientRegenerateSecondaryKeyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientRegenerateSecondaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return SubscriptionClientRegenerateSecondaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionClientRegenerateSecondaryKeyResponse{}, nil
}

// regenerateSecondaryKeyCreateRequest creates the RegenerateSecondaryKey request.
func (client *SubscriptionClient) regenerateSecondaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionClientRegenerateSecondaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regenerateSecondaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Updates the details of a subscription specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - sid - Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - Update parameters.
//   - options - SubscriptionClientUpdateOptions contains the optional parameters for the SubscriptionClient.Update method.
func (client *SubscriptionClient) Update(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, parameters SubscriptionUpdateParameters, options *SubscriptionClientUpdateOptions) (SubscriptionClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, sid, ifMatch, parameters, options)
	if err != nil {
		return SubscriptionClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SubscriptionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, parameters SubscriptionUpdateParameters, options *SubscriptionClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Notify != nil {
		reqQP.Set("notify", strconv.FormatBool(*options.Notify))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SubscriptionClient) updateHandleResponse(resp *http.Response) (SubscriptionClientUpdateResponse, error) {
	result := SubscriptionClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionContract); err != nil {
		return SubscriptionClientUpdateResponse{}, err
	}
	return result, nil
}
