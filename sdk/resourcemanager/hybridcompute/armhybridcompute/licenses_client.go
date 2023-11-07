//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// LicensesClient contains the methods for the Licenses group.
// Don't use this type directly, use NewLicensesClient() instead.
type LicensesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLicensesClient creates a new instance of LicensesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLicensesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LicensesClient, error) {
	cl, err := arm.NewClient(moduleName+".LicensesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LicensesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - licenseName - The name of the license.
//   - parameters - Parameters supplied to the Create license operation.
//   - options - LicensesClientBeginCreateOrUpdateOptions contains the optional parameters for the LicensesClient.BeginCreateOrUpdate
//     method.
func (client *LicensesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, licenseName string, parameters License, options *LicensesClientBeginCreateOrUpdateOptions) (*runtime.Poller[LicensesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, licenseName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LicensesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LicensesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - The operation to create or update a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
func (client *LicensesClient) createOrUpdate(ctx context.Context, resourceGroupName string, licenseName string, parameters License, options *LicensesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, licenseName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LicensesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, licenseName string, parameters License, options *LicensesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/licenses/{licenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if licenseName == "" {
		return nil, errors.New("parameter licenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{licenseName}", url.PathEscape(licenseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - licenseName - The name of the license.
//   - options - LicensesClientBeginDeleteOptions contains the optional parameters for the LicensesClient.BeginDelete method.
func (client *LicensesClient) BeginDelete(ctx context.Context, resourceGroupName string, licenseName string, options *LicensesClientBeginDeleteOptions) (*runtime.Poller[LicensesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, licenseName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LicensesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LicensesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
func (client *LicensesClient) deleteOperation(ctx context.Context, resourceGroupName string, licenseName string, options *LicensesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, licenseName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LicensesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, licenseName string, options *LicensesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/licenses/{licenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if licenseName == "" {
		return nil, errors.New("parameter licenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{licenseName}", url.PathEscape(licenseName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about the view of a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - licenseName - The name of the license.
//   - options - LicensesClientGetOptions contains the optional parameters for the LicensesClient.Get method.
func (client *LicensesClient) Get(ctx context.Context, resourceGroupName string, licenseName string, options *LicensesClientGetOptions) (LicensesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, licenseName, options)
	if err != nil {
		return LicensesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LicensesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LicensesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LicensesClient) getCreateRequest(ctx context.Context, resourceGroupName string, licenseName string, options *LicensesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/licenses/{licenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if licenseName == "" {
		return nil, errors.New("parameter licenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{licenseName}", url.PathEscape(licenseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LicensesClient) getHandleResponse(resp *http.Response) (LicensesClientGetResponse, error) {
	result := LicensesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.License); err != nil {
		return LicensesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The operation to get all licenses of a non-Azure machine
//
// Generated from API version 2023-06-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - LicensesClientListByResourceGroupOptions contains the optional parameters for the LicensesClient.NewListByResourceGroupPager
//     method.
func (client *LicensesClient) NewListByResourceGroupPager(resourceGroupName string, options *LicensesClientListByResourceGroupOptions) *runtime.Pager[LicensesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[LicensesClientListByResourceGroupResponse]{
		More: func(page LicensesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LicensesClientListByResourceGroupResponse) (LicensesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LicensesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LicensesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LicensesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LicensesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *LicensesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/licenses"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LicensesClient) listByResourceGroupHandleResponse(resp *http.Response) (LicensesClientListByResourceGroupResponse, error) {
	result := LicensesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LicensesListResult); err != nil {
		return LicensesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - The operation to get all licenses of a non-Azure machine
//
// Generated from API version 2023-06-20-preview
//   - options - LicensesClientListBySubscriptionOptions contains the optional parameters for the LicensesClient.NewListBySubscriptionPager
//     method.
func (client *LicensesClient) NewListBySubscriptionPager(options *LicensesClientListBySubscriptionOptions) *runtime.Pager[LicensesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[LicensesClientListBySubscriptionResponse]{
		More: func(page LicensesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LicensesClientListBySubscriptionResponse) (LicensesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LicensesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LicensesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LicensesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LicensesClient) listBySubscriptionCreateRequest(ctx context.Context, options *LicensesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/licenses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LicensesClient) listBySubscriptionHandleResponse(resp *http.Response) (LicensesClientListBySubscriptionResponse, error) {
	result := LicensesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LicensesListResult); err != nil {
		return LicensesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - licenseName - The name of the license.
//   - parameters - Parameters supplied to the Update license operation.
//   - options - LicensesClientBeginUpdateOptions contains the optional parameters for the LicensesClient.BeginUpdate method.
func (client *LicensesClient) BeginUpdate(ctx context.Context, resourceGroupName string, licenseName string, parameters LicenseUpdate, options *LicensesClientBeginUpdateOptions) (*runtime.Poller[LicensesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, licenseName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LicensesClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LicensesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
func (client *LicensesClient) update(ctx context.Context, resourceGroupName string, licenseName string, parameters LicenseUpdate, options *LicensesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, licenseName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *LicensesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, licenseName string, parameters LicenseUpdate, options *LicensesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/licenses/{licenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if licenseName == "" {
		return nil, errors.New("parameter licenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{licenseName}", url.PathEscape(licenseName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginValidateLicense - The operation to validate a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
//   - parameters - Parameters supplied to the license validation operation.
//   - options - LicensesClientBeginValidateLicenseOptions contains the optional parameters for the LicensesClient.BeginValidateLicense
//     method.
func (client *LicensesClient) BeginValidateLicense(ctx context.Context, parameters License, options *LicensesClientBeginValidateLicenseOptions) (*runtime.Poller[LicensesClientValidateLicenseResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.validateLicense(ctx, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LicensesClientValidateLicenseResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LicensesClientValidateLicenseResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ValidateLicense - The operation to validate a license.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-20-preview
func (client *LicensesClient) validateLicense(ctx context.Context, parameters License, options *LicensesClientBeginValidateLicenseOptions) (*http.Response, error) {
	var err error
	req, err := client.validateLicenseCreateRequest(ctx, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// validateLicenseCreateRequest creates the ValidateLicense request.
func (client *LicensesClient) validateLicenseCreateRequest(ctx context.Context, parameters License, options *LicensesClientBeginValidateLicenseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/validateLicense"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
