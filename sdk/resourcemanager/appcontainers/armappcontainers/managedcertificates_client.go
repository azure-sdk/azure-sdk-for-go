// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ManagedCertificatesClient contains the methods for the ManagedCertificates group.
// Don't use this type directly, use NewManagedCertificatesClient() instead.
type ManagedCertificatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedCertificatesClient creates a new instance of ManagedCertificatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedCertificatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedCertificatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update a Managed Certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - managedCertificateName - Name of the Managed Certificate.
//   - options - ManagedCertificatesClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedCertificatesClient.BeginCreateOrUpdate
//     method.
func (client *ManagedCertificatesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *ManagedCertificatesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedCertificatesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, environmentName, managedCertificateName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedCertificatesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedCertificatesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or Update a Managed Certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-02-preview
func (client *ManagedCertificatesClient) createOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *ManagedCertificatesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedCertificatesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, managedCertificateName, options)
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
func (client *ManagedCertificatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *ManagedCertificatesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/managedCertificates/{managedCertificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if managedCertificateName == "" {
		return nil, errors.New("parameter managedCertificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedCertificateName}", url.PathEscape(managedCertificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ManagedCertificateEnvelope != nil {
		if err := runtime.MarshalAsJSON(req, *options.ManagedCertificateEnvelope); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Delete - Deletes the specified Managed Certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - managedCertificateName - Name of the Managed Certificate.
//   - options - ManagedCertificatesClientDeleteOptions contains the optional parameters for the ManagedCertificatesClient.Delete
//     method.
func (client *ManagedCertificatesClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *ManagedCertificatesClientDeleteOptions) (ManagedCertificatesClientDeleteResponse, error) {
	var err error
	const operationName = "ManagedCertificatesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, managedCertificateName, options)
	if err != nil {
		return ManagedCertificatesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedCertificatesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ManagedCertificatesClientDeleteResponse{}, err
	}
	return ManagedCertificatesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedCertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, _ *ManagedCertificatesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/managedCertificates/{managedCertificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if managedCertificateName == "" {
		return nil, errors.New("parameter managedCertificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedCertificateName}", url.PathEscape(managedCertificateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified Managed Certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - managedCertificateName - Name of the Managed Certificate.
//   - options - ManagedCertificatesClientGetOptions contains the optional parameters for the ManagedCertificatesClient.Get method.
func (client *ManagedCertificatesClient) Get(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *ManagedCertificatesClientGetOptions) (ManagedCertificatesClientGetResponse, error) {
	var err error
	const operationName = "ManagedCertificatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, managedCertificateName, options)
	if err != nil {
		return ManagedCertificatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedCertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedCertificatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedCertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, _ *ManagedCertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/managedCertificates/{managedCertificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if managedCertificateName == "" {
		return nil, errors.New("parameter managedCertificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedCertificateName}", url.PathEscape(managedCertificateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedCertificatesClient) getHandleResponse(resp *http.Response) (ManagedCertificatesClientGetResponse, error) {
	result := ManagedCertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedCertificate); err != nil {
		return ManagedCertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Managed Certificates in a given managed environment.
//
// Generated from API version 2024-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - options - ManagedCertificatesClientListOptions contains the optional parameters for the ManagedCertificatesClient.NewListPager
//     method.
func (client *ManagedCertificatesClient) NewListPager(resourceGroupName string, environmentName string, options *ManagedCertificatesClientListOptions) *runtime.Pager[ManagedCertificatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedCertificatesClientListResponse]{
		More: func(page ManagedCertificatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedCertificatesClientListResponse) (ManagedCertificatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedCertificatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, environmentName, options)
			}, nil)
			if err != nil {
				return ManagedCertificatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedCertificatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, _ *ManagedCertificatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/managedCertificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedCertificatesClient) listHandleResponse(resp *http.Response) (ManagedCertificatesClientListResponse, error) {
	result := ManagedCertificatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedCertificateCollection); err != nil {
		return ManagedCertificatesClientListResponse{}, err
	}
	return result, nil
}

// Update - Patches a managed certificate. Oly patching of tags is supported
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - managedCertificateName - Name of the Managed Certificate.
//   - managedCertificateEnvelope - Properties of a managed certificate that need to be updated
//   - options - ManagedCertificatesClientUpdateOptions contains the optional parameters for the ManagedCertificatesClient.Update
//     method.
func (client *ManagedCertificatesClient) Update(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, managedCertificateEnvelope ManagedCertificatePatch, options *ManagedCertificatesClientUpdateOptions) (ManagedCertificatesClientUpdateResponse, error) {
	var err error
	const operationName = "ManagedCertificatesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, environmentName, managedCertificateName, managedCertificateEnvelope, options)
	if err != nil {
		return ManagedCertificatesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedCertificatesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedCertificatesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ManagedCertificatesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, managedCertificateEnvelope ManagedCertificatePatch, _ *ManagedCertificatesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/managedCertificates/{managedCertificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if managedCertificateName == "" {
		return nil, errors.New("parameter managedCertificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedCertificateName}", url.PathEscape(managedCertificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, managedCertificateEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ManagedCertificatesClient) updateHandleResponse(resp *http.Response) (ManagedCertificatesClientUpdateResponse, error) {
	result := ManagedCertificatesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedCertificate); err != nil {
		return ManagedCertificatesClientUpdateResponse{}, err
	}
	return result, nil
}
