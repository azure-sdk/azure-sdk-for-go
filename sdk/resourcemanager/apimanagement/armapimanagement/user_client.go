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

// UserClient contains the methods for the User group.
// Don't use this type directly, use NewUserClient() instead.
type UserClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUserClient creates a new instance of UserClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUserClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UserClient, error) {
	cl, err := arm.NewClient(moduleName+".UserClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UserClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates a user.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - parameters - Create or update parameters.
//   - options - UserClientCreateOrUpdateOptions contains the optional parameters for the UserClient.CreateOrUpdate method.
func (client *UserClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, userID string, parameters UserCreateParameters, options *UserClientCreateOrUpdateOptions) (UserClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, userID, parameters, options)
	if err != nil {
		return UserClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return UserClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *UserClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, parameters UserCreateParameters, options *UserClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Notify != nil {
		reqQP.Set("notify", strconv.FormatBool(*options.Notify))
	}
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *UserClient) createOrUpdateHandleResponse(resp *http.Response) (UserClientCreateOrUpdateResponse, error) {
	result := UserClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserContract); err != nil {
		return UserClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specific user.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - UserClientDeleteOptions contains the optional parameters for the UserClient.Delete method.
func (client *UserClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, userID string, ifMatch string, options *UserClientDeleteOptions) (UserClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, userID, ifMatch, options)
	if err != nil {
		return UserClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return UserClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return UserClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *UserClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, ifMatch string, options *UserClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.DeleteSubscriptions != nil {
		reqQP.Set("deleteSubscriptions", strconv.FormatBool(*options.DeleteSubscriptions))
	}
	if options != nil && options.Notify != nil {
		reqQP.Set("notify", strconv.FormatBool(*options.Notify))
	}
	reqQP.Set("api-version", "2023-03-01-preview")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GenerateSsoURL - Retrieves a redirection URL containing an authentication token for signing a given user into the developer
// portal.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - UserClientGenerateSsoURLOptions contains the optional parameters for the UserClient.GenerateSsoURL method.
func (client *UserClient) GenerateSsoURL(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserClientGenerateSsoURLOptions) (UserClientGenerateSsoURLResponse, error) {
	req, err := client.generateSsoURLCreateRequest(ctx, resourceGroupName, serviceName, userID, options)
	if err != nil {
		return UserClientGenerateSsoURLResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserClientGenerateSsoURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UserClientGenerateSsoURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateSsoURLHandleResponse(resp)
}

// generateSsoURLCreateRequest creates the GenerateSsoURL request.
func (client *UserClient) generateSsoURLCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserClientGenerateSsoURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/generateSsoUrl"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateSsoURLHandleResponse handles the GenerateSsoURL response.
func (client *UserClient) generateSsoURLHandleResponse(resp *http.Response) (UserClientGenerateSsoURLResponse, error) {
	result := UserClientGenerateSsoURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GenerateSsoURLResult); err != nil {
		return UserClientGenerateSsoURLResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of the user specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - UserClientGetOptions contains the optional parameters for the UserClient.Get method.
func (client *UserClient) Get(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserClientGetOptions) (UserClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, userID, options)
	if err != nil {
		return UserClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UserClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *UserClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UserClient) getHandleResponse(resp *http.Response) (UserClientGetResponse, error) {
	result := UserClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserContract); err != nil {
		return UserClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the user specified by its identifier.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - UserClientGetEntityTagOptions contains the optional parameters for the UserClient.GetEntityTag method.
func (client *UserClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserClientGetEntityTagOptions) (UserClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, userID, options)
	if err != nil {
		return UserClientGetEntityTagResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UserClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *UserClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *UserClient) getEntityTagHandleResponse(resp *http.Response) (UserClientGetEntityTagResponse, error) {
	result := UserClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// GetSharedAccessToken - Gets the Shared Access Authorization Token for the User.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - parameters - Create Authorization Token parameters.
//   - options - UserClientGetSharedAccessTokenOptions contains the optional parameters for the UserClient.GetSharedAccessToken
//     method.
func (client *UserClient) GetSharedAccessToken(ctx context.Context, resourceGroupName string, serviceName string, userID string, parameters UserTokenParameters, options *UserClientGetSharedAccessTokenOptions) (UserClientGetSharedAccessTokenResponse, error) {
	req, err := client.getSharedAccessTokenCreateRequest(ctx, resourceGroupName, serviceName, userID, parameters, options)
	if err != nil {
		return UserClientGetSharedAccessTokenResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserClientGetSharedAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UserClientGetSharedAccessTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSharedAccessTokenHandleResponse(resp)
}

// getSharedAccessTokenCreateRequest creates the GetSharedAccessToken request.
func (client *UserClient) getSharedAccessTokenCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, parameters UserTokenParameters, options *UserClientGetSharedAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/token"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// getSharedAccessTokenHandleResponse handles the GetSharedAccessToken response.
func (client *UserClient) getSharedAccessTokenHandleResponse(resp *http.Response) (UserClientGetSharedAccessTokenResponse, error) {
	result := UserClientGetSharedAccessTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserTokenResult); err != nil {
		return UserClientGetSharedAccessTokenResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Lists a collection of registered users in the specified service instance.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - UserClientListByServiceOptions contains the optional parameters for the UserClient.NewListByServicePager method.
func (client *UserClient) NewListByServicePager(resourceGroupName string, serviceName string, options *UserClientListByServiceOptions) *runtime.Pager[UserClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[UserClientListByServiceResponse]{
		More: func(page UserClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UserClientListByServiceResponse) (UserClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return UserClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UserClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UserClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *UserClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *UserClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users"
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
	if options != nil && options.ExpandGroups != nil {
		reqQP.Set("expandGroups", strconv.FormatBool(*options.ExpandGroups))
	}
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *UserClient) listByServiceHandleResponse(resp *http.Response) (UserClientListByServiceResponse, error) {
	result := UserClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserCollection); err != nil {
		return UserClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the user specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - Update parameters.
//   - options - UserClientUpdateOptions contains the optional parameters for the UserClient.Update method.
func (client *UserClient) Update(ctx context.Context, resourceGroupName string, serviceName string, userID string, ifMatch string, parameters UserUpdateParameters, options *UserClientUpdateOptions) (UserClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, userID, ifMatch, parameters, options)
	if err != nil {
		return UserClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UserClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UserClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *UserClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, ifMatch string, parameters UserUpdateParameters, options *UserClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *UserClient) updateHandleResponse(resp *http.Response) (UserClientUpdateResponse, error) {
	result := UserClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserContract); err != nil {
		return UserClientUpdateResponse{}, err
	}
	return result, nil
}
