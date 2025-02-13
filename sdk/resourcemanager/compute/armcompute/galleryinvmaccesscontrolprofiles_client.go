//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// GalleryInVMAccessControlProfilesClient contains the methods for the GalleryInVMAccessControlProfiles group.
// Don't use this type directly, use NewGalleryInVMAccessControlProfilesClient() instead.
type GalleryInVMAccessControlProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGalleryInVMAccessControlProfilesClient creates a new instance of GalleryInVMAccessControlProfilesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGalleryInVMAccessControlProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GalleryInVMAccessControlProfilesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GalleryInVMAccessControlProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a gallery inVMAccessControlProfile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - galleryName - The name of the Shared Image Gallery.
//   - inVMAccessControlProfileName - The name of the gallery inVMAccessControlProfile to be retrieved.
//   - galleryInVMAccessControlProfile - Parameters supplied to the create or update gallery inVMAccessControlProfile operation.
//   - options - GalleryInVMAccessControlProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the GalleryInVMAccessControlProfilesClient.BeginCreateOrUpdate
//     method.
func (client *GalleryInVMAccessControlProfilesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, galleryInVMAccessControlProfile GalleryInVMAccessControlProfile, options *GalleryInVMAccessControlProfilesClientBeginCreateOrUpdateOptions) (*runtime.Poller[GalleryInVMAccessControlProfilesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, inVMAccessControlProfileName, galleryInVMAccessControlProfile, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GalleryInVMAccessControlProfilesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GalleryInVMAccessControlProfilesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a gallery inVMAccessControlProfile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-03
func (client *GalleryInVMAccessControlProfilesClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, galleryInVMAccessControlProfile GalleryInVMAccessControlProfile, options *GalleryInVMAccessControlProfilesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GalleryInVMAccessControlProfilesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, inVMAccessControlProfileName, galleryInVMAccessControlProfile, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryInVMAccessControlProfilesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, galleryInVMAccessControlProfile GalleryInVMAccessControlProfile, options *GalleryInVMAccessControlProfilesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/inVMAccessControlProfiles/{inVMAccessControlProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if inVMAccessControlProfileName == "" {
		return nil, errors.New("parameter inVMAccessControlProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inVMAccessControlProfileName}", url.PathEscape(inVMAccessControlProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, galleryInVMAccessControlProfile); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a gallery inVMAccessControlProfile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - galleryName - The name of the Shared Image Gallery.
//   - inVMAccessControlProfileName - The name of the gallery inVMAccessControlProfile to be retrieved.
//   - options - GalleryInVMAccessControlProfilesClientBeginDeleteOptions contains the optional parameters for the GalleryInVMAccessControlProfilesClient.BeginDelete
//     method.
func (client *GalleryInVMAccessControlProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, options *GalleryInVMAccessControlProfilesClientBeginDeleteOptions) (*runtime.Poller[GalleryInVMAccessControlProfilesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, inVMAccessControlProfileName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GalleryInVMAccessControlProfilesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GalleryInVMAccessControlProfilesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a gallery inVMAccessControlProfile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-03
func (client *GalleryInVMAccessControlProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, options *GalleryInVMAccessControlProfilesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "GalleryInVMAccessControlProfilesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, inVMAccessControlProfileName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GalleryInVMAccessControlProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, options *GalleryInVMAccessControlProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/inVMAccessControlProfiles/{inVMAccessControlProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if inVMAccessControlProfileName == "" {
		return nil, errors.New("parameter inVMAccessControlProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inVMAccessControlProfileName}", url.PathEscape(inVMAccessControlProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a gallery inVMAccessControlProfile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - galleryName - The name of the Shared Image Gallery.
//   - inVMAccessControlProfileName - The name of the gallery inVMAccessControlProfile to be retrieved.
//   - options - GalleryInVMAccessControlProfilesClientGetOptions contains the optional parameters for the GalleryInVMAccessControlProfilesClient.Get
//     method.
func (client *GalleryInVMAccessControlProfilesClient) Get(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, options *GalleryInVMAccessControlProfilesClientGetOptions) (GalleryInVMAccessControlProfilesClientGetResponse, error) {
	var err error
	const operationName = "GalleryInVMAccessControlProfilesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, inVMAccessControlProfileName, options)
	if err != nil {
		return GalleryInVMAccessControlProfilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GalleryInVMAccessControlProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GalleryInVMAccessControlProfilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GalleryInVMAccessControlProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, options *GalleryInVMAccessControlProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/inVMAccessControlProfiles/{inVMAccessControlProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if inVMAccessControlProfileName == "" {
		return nil, errors.New("parameter inVMAccessControlProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inVMAccessControlProfileName}", url.PathEscape(inVMAccessControlProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryInVMAccessControlProfilesClient) getHandleResponse(resp *http.Response) (GalleryInVMAccessControlProfilesClientGetResponse, error) {
	result := GalleryInVMAccessControlProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryInVMAccessControlProfile); err != nil {
		return GalleryInVMAccessControlProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByGalleryPager - List gallery inVMAccessControlProfiles in a gallery.
//
// Generated from API version 2024-03-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - galleryName - The name of the Shared Image Gallery.
//   - options - GalleryInVMAccessControlProfilesClientListByGalleryOptions contains the optional parameters for the GalleryInVMAccessControlProfilesClient.NewListByGalleryPager
//     method.
func (client *GalleryInVMAccessControlProfilesClient) NewListByGalleryPager(resourceGroupName string, galleryName string, options *GalleryInVMAccessControlProfilesClientListByGalleryOptions) *runtime.Pager[GalleryInVMAccessControlProfilesClientListByGalleryResponse] {
	return runtime.NewPager(runtime.PagingHandler[GalleryInVMAccessControlProfilesClientListByGalleryResponse]{
		More: func(page GalleryInVMAccessControlProfilesClientListByGalleryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleryInVMAccessControlProfilesClientListByGalleryResponse) (GalleryInVMAccessControlProfilesClientListByGalleryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GalleryInVMAccessControlProfilesClient.NewListByGalleryPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByGalleryCreateRequest(ctx, resourceGroupName, galleryName, options)
			}, nil)
			if err != nil {
				return GalleryInVMAccessControlProfilesClientListByGalleryResponse{}, err
			}
			return client.listByGalleryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByGalleryCreateRequest creates the ListByGallery request.
func (client *GalleryInVMAccessControlProfilesClient) listByGalleryCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleryInVMAccessControlProfilesClientListByGalleryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/inVMAccessControlProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByGalleryHandleResponse handles the ListByGallery response.
func (client *GalleryInVMAccessControlProfilesClient) listByGalleryHandleResponse(resp *http.Response) (GalleryInVMAccessControlProfilesClientListByGalleryResponse, error) {
	result := GalleryInVMAccessControlProfilesClientListByGalleryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryInVMAccessControlProfileList); err != nil {
		return GalleryInVMAccessControlProfilesClientListByGalleryResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a gallery inVMAccessControlProfile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - galleryName - The name of the Shared Image Gallery.
//   - inVMAccessControlProfileName - The name of the gallery inVMAccessControlProfile to be retrieved.
//   - galleryInVMAccessControlProfile - Parameters supplied to the update gallery inVMAccessControlProfile operation.
//   - options - GalleryInVMAccessControlProfilesClientBeginUpdateOptions contains the optional parameters for the GalleryInVMAccessControlProfilesClient.BeginUpdate
//     method.
func (client *GalleryInVMAccessControlProfilesClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, galleryInVMAccessControlProfile GalleryInVMAccessControlProfileUpdate, options *GalleryInVMAccessControlProfilesClientBeginUpdateOptions) (*runtime.Poller[GalleryInVMAccessControlProfilesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, galleryName, inVMAccessControlProfileName, galleryInVMAccessControlProfile, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GalleryInVMAccessControlProfilesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GalleryInVMAccessControlProfilesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a gallery inVMAccessControlProfile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-03
func (client *GalleryInVMAccessControlProfilesClient) update(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, galleryInVMAccessControlProfile GalleryInVMAccessControlProfileUpdate, options *GalleryInVMAccessControlProfilesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GalleryInVMAccessControlProfilesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, inVMAccessControlProfileName, galleryInVMAccessControlProfile, options)
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
func (client *GalleryInVMAccessControlProfilesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, inVMAccessControlProfileName string, galleryInVMAccessControlProfile GalleryInVMAccessControlProfileUpdate, options *GalleryInVMAccessControlProfilesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/inVMAccessControlProfiles/{inVMAccessControlProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if inVMAccessControlProfileName == "" {
		return nil, errors.New("parameter inVMAccessControlProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inVMAccessControlProfileName}", url.PathEscape(inVMAccessControlProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, galleryInVMAccessControlProfile); err != nil {
		return nil, err
	}
	return req, nil
}
