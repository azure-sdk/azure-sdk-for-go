// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// SecretClient contains the methods for the Secret group.
// Don't use this type directly, use NewSecretClient() instead.
type SecretClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecretClient creates a new instance of SecretClient with the specified values.
//   - subscriptionID - The customer subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecretClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecretClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecretClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a secret resource with the specified name, description and properties. If a secret resource with the same
// name exists, then it is updated with the specified description and properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - secretResourceDescription - Description for creating a secret resource.
//   - options - SecretClientCreateOptions contains the optional parameters for the SecretClient.Create method.
func (client *SecretClient) Create(ctx context.Context, resourceGroupName string, secretResourceName string, secretResourceDescription SecretResourceDescription, options *SecretClientCreateOptions) (SecretClientCreateResponse, error) {
	var err error
	const operationName = "SecretClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, secretResourceName, secretResourceDescription, options)
	if err != nil {
		return SecretClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return SecretClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SecretClient) createCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, secretResourceDescription SecretResourceDescription, _ *SecretClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, secretResourceDescription); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SecretClient) createHandleResponse(resp *http.Response) (SecretClientCreateResponse, error) {
	result := SecretClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretResourceDescription); err != nil {
		return SecretClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the secret resource identified by the name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - options - SecretClientDeleteOptions contains the optional parameters for the SecretClient.Delete method.
func (client *SecretClient) Delete(ctx context.Context, resourceGroupName string, secretResourceName string, options *SecretClientDeleteOptions) (SecretClientDeleteResponse, error) {
	var err error
	const operationName = "SecretClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, secretResourceName, options)
	if err != nil {
		return SecretClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SecretClientDeleteResponse{}, err
	}
	return SecretClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecretClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, _ *SecretClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the information about the secret resource with the given name. The information include the description and other
// properties of the secret.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - options - SecretClientGetOptions contains the optional parameters for the SecretClient.Get method.
func (client *SecretClient) Get(ctx context.Context, resourceGroupName string, secretResourceName string, options *SecretClientGetOptions) (SecretClientGetResponse, error) {
	var err error
	const operationName = "SecretClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, secretResourceName, options)
	if err != nil {
		return SecretClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecretClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecretClient) getCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, _ *SecretClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecretClient) getHandleResponse(resp *http.Response) (SecretClientGetResponse, error) {
	result := SecretClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretResourceDescription); err != nil {
		return SecretClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets the information about all secret resources in a given resource group. The information
// include the description and other properties of the Secret.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - options - SecretClientListByResourceGroupOptions contains the optional parameters for the SecretClient.NewListByResourceGroupPager
//     method.
func (client *SecretClient) NewListByResourceGroupPager(resourceGroupName string, options *SecretClientListByResourceGroupOptions) *runtime.Pager[SecretClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecretClientListByResourceGroupResponse]{
		More: func(page SecretClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecretClientListByResourceGroupResponse) (SecretClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecretClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SecretClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SecretClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *SecretClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets"
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
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SecretClient) listByResourceGroupHandleResponse(resp *http.Response) (SecretClientListByResourceGroupResponse, error) {
	result := SecretClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretResourceDescriptionList); err != nil {
		return SecretClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets the information about all secret resources in a given resource group. The information
// include the description and other properties of the secret.
//
// Generated from API version 2018-09-01-preview
//   - options - SecretClientListBySubscriptionOptions contains the optional parameters for the SecretClient.NewListBySubscriptionPager
//     method.
func (client *SecretClient) NewListBySubscriptionPager(options *SecretClientListBySubscriptionOptions) *runtime.Pager[SecretClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecretClientListBySubscriptionResponse]{
		More: func(page SecretClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecretClientListBySubscriptionResponse) (SecretClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecretClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SecretClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SecretClient) listBySubscriptionCreateRequest(ctx context.Context, _ *SecretClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/secrets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SecretClient) listBySubscriptionHandleResponse(resp *http.Response) (SecretClientListBySubscriptionResponse, error) {
	result := SecretClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretResourceDescriptionList); err != nil {
		return SecretClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
