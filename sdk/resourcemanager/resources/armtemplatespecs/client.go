//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

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

// Client contains the methods for the TemplateSpecs group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - Subscription Id which forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName+".Client", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Template Spec.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - templateSpec - Template Spec supplied to the operation.
//   - options - ClientCreateOrUpdateOptions contains the optional parameters for the Client.CreateOrUpdate method.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpec TemplateSpec, options *ClientCreateOrUpdateOptions) (ClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, templateSpecName, templateSpec, options)
	if err != nil {
		return ClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *Client) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpec TemplateSpec, options *ClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, templateSpec); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *Client) createOrUpdateHandleResponse(resp *http.Response) (ClientCreateOrUpdateResponse, error) {
	result := ClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpec); err != nil {
		return ClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Template Spec by name. When operation completes, status code 200 returned without content.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - options - ClientDeleteOptions contains the optional parameters for the Client.Delete method.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, templateSpecName string, options *ClientDeleteOptions) (ClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, templateSpecName, options)
	if err != nil {
		return ClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientDeleteResponse{}, err
	}
	return ClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *Client) deleteCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, options *ClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Template Spec with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - options - ClientGetOptions contains the optional parameters for the Client.Get method.
func (client *Client) Get(ctx context.Context, resourceGroupName string, templateSpecName string, options *ClientGetOptions) (ClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, templateSpecName, options)
	if err != nil {
		return ClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *Client) getCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, options *ClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Client) getHandleResponse(resp *http.Response) (ClientGetResponse, error) {
	result := ClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpec); err != nil {
		return ClientGetResponse{}, err
	}
	return result, nil
}

// GetBuiltIn - Gets a built-in Template Spec with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - templateSpecName - Name of the Template Spec.
//   - options - ClientGetBuiltInOptions contains the optional parameters for the Client.GetBuiltIn method.
func (client *Client) GetBuiltIn(ctx context.Context, templateSpecName string, options *ClientGetBuiltInOptions) (ClientGetBuiltInResponse, error) {
	var err error
	req, err := client.getBuiltInCreateRequest(ctx, templateSpecName, options)
	if err != nil {
		return ClientGetBuiltInResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetBuiltInResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetBuiltInResponse{}, err
	}
	resp, err := client.getBuiltInHandleResponse(httpResp)
	return resp, err
}

// getBuiltInCreateRequest creates the GetBuiltIn request.
func (client *Client) getBuiltInCreateRequest(ctx context.Context, templateSpecName string, options *ClientGetBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/builtInTemplateSpecs/{templateSpecName}"
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBuiltInHandleResponse handles the GetBuiltIn response.
func (client *Client) getBuiltInHandleResponse(resp *http.Response) (ClientGetBuiltInResponse, error) {
	result := ClientGetBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpec); err != nil {
		return ClientGetBuiltInResponse{}, err
	}
	return result, nil
}

// NewListBuiltInsPager - Lists built-in Template Specs.
//
// Generated from API version 2022-02-01
//   - options - ClientListBuiltInsOptions contains the optional parameters for the Client.NewListBuiltInsPager method.
func (client *Client) NewListBuiltInsPager(options *ClientListBuiltInsOptions) *runtime.Pager[ClientListBuiltInsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListBuiltInsResponse]{
		More: func(page ClientListBuiltInsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListBuiltInsResponse) (ClientListBuiltInsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBuiltInsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListBuiltInsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientListBuiltInsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListBuiltInsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBuiltInsHandleResponse(resp)
		},
	})
}

// listBuiltInsCreateRequest creates the ListBuiltIns request.
func (client *Client) listBuiltInsCreateRequest(ctx context.Context, options *ClientListBuiltInsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/builtInTemplateSpecs/"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBuiltInsHandleResponse handles the ListBuiltIns response.
func (client *Client) listBuiltInsHandleResponse(resp *http.Response) (ClientListBuiltInsResponse, error) {
	result := ClientListBuiltInsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ClientListBuiltInsResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the Template Specs within the specified resource group.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ClientListByResourceGroupOptions contains the optional parameters for the Client.NewListByResourceGroupPager
//     method.
func (client *Client) NewListByResourceGroupPager(resourceGroupName string, options *ClientListByResourceGroupOptions) *runtime.Pager[ClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListByResourceGroupResponse]{
		More: func(page ClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListByResourceGroupResponse) (ClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *Client) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *Client) listByResourceGroupHandleResponse(resp *http.Response) (ClientListByResourceGroupResponse, error) {
	result := ClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all the Template Specs within the specified subscriptions.
//
// Generated from API version 2022-02-01
//   - options - ClientListBySubscriptionOptions contains the optional parameters for the Client.NewListBySubscriptionPager method.
func (client *Client) NewListBySubscriptionPager(options *ClientListBySubscriptionOptions) *runtime.Pager[ClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListBySubscriptionResponse]{
		More: func(page ClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListBySubscriptionResponse) (ClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *Client) listBySubscriptionCreateRequest(ctx context.Context, options *ClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Resources/templateSpecs/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *Client) listBySubscriptionHandleResponse(resp *http.Response) (ClientListBySubscriptionResponse, error) {
	result := ClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates Template Spec tags with specified values.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - options - ClientUpdateOptions contains the optional parameters for the Client.Update method.
func (client *Client) Update(ctx context.Context, resourceGroupName string, templateSpecName string, options *ClientUpdateOptions) (ClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, templateSpecName, options)
	if err != nil {
		return ClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *Client) updateCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, options *ClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.TemplateSpec != nil {
		if err := runtime.MarshalAsJSON(req, *options.TemplateSpec); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *Client) updateHandleResponse(resp *http.Response) (ClientUpdateResponse, error) {
	result := ClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpec); err != nil {
		return ClientUpdateResponse{}, err
	}
	return result, nil
}
