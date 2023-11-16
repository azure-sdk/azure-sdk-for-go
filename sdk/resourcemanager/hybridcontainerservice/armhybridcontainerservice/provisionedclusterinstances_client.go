//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcontainerservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ProvisionedClusterInstancesClient contains the methods for the ProvisionedClusterInstances group.
// Don't use this type directly, use NewProvisionedClusterInstancesClient() instead.
type ProvisionedClusterInstancesClient struct {
	internal *arm.Client
}

// NewProvisionedClusterInstancesClient creates a new instance of ProvisionedClusterInstancesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProvisionedClusterInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ProvisionedClusterInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProvisionedClusterInstancesClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates the Hybrid AKS provisioned cluster instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientBeginCreateOrUpdateOptions contains the optional parameters for the ProvisionedClusterInstancesClient.BeginCreateOrUpdate
//     method.
func (client *ProvisionedClusterInstancesClient) BeginCreateOrUpdate(ctx context.Context, connectedClusterResourceURI string, provisionedClusterInstance ProvisionedClusters, options *ProvisionedClusterInstancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ProvisionedClusterInstancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, connectedClusterResourceURI, provisionedClusterInstance, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProvisionedClusterInstancesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProvisionedClusterInstancesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates the Hybrid AKS provisioned cluster instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
func (client *ProvisionedClusterInstancesClient) createOrUpdate(ctx context.Context, connectedClusterResourceURI string, provisionedClusterInstance ProvisionedClusters, options *ProvisionedClusterInstancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ProvisionedClusterInstancesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, connectedClusterResourceURI, provisionedClusterInstance, options)
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
func (client *ProvisionedClusterInstancesClient) createOrUpdateCreateRequest(ctx context.Context, connectedClusterResourceURI string, provisionedClusterInstance ProvisionedClusters, options *ProvisionedClusterInstancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, provisionedClusterInstance); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the Hybrid AKS provisioned cluster instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientBeginDeleteOptions contains the optional parameters for the ProvisionedClusterInstancesClient.BeginDelete
//     method.
func (client *ProvisionedClusterInstancesClient) BeginDelete(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginDeleteOptions) (*runtime.Poller[ProvisionedClusterInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, connectedClusterResourceURI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProvisionedClusterInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProvisionedClusterInstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the Hybrid AKS provisioned cluster instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
func (client *ProvisionedClusterInstancesClient) deleteOperation(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ProvisionedClusterInstancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, connectedClusterResourceURI, options)
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
func (client *ProvisionedClusterInstancesClient) deleteCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the Hybrid AKS provisioned cluster instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientGetOptions contains the optional parameters for the ProvisionedClusterInstancesClient.Get
//     method.
func (client *ProvisionedClusterInstancesClient) Get(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientGetOptions) (ProvisionedClusterInstancesClientGetResponse, error) {
	var err error
	const operationName = "ProvisionedClusterInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, connectedClusterResourceURI, options)
	if err != nil {
		return ProvisionedClusterInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProvisionedClusterInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProvisionedClusterInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProvisionedClusterInstancesClient) getCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProvisionedClusterInstancesClient) getHandleResponse(resp *http.Response) (ProvisionedClusterInstancesClientGetResponse, error) {
	result := ProvisionedClusterInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProvisionedClusters); err != nil {
		return ProvisionedClusterInstancesClientGetResponse{}, err
	}
	return result, nil
}

// GetUpgradeProfile - Gets the upgrade profile of a provisioned cluster instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientGetUpgradeProfileOptions contains the optional parameters for the ProvisionedClusterInstancesClient.GetUpgradeProfile
//     method.
func (client *ProvisionedClusterInstancesClient) GetUpgradeProfile(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientGetUpgradeProfileOptions) (ProvisionedClusterInstancesClientGetUpgradeProfileResponse, error) {
	var err error
	const operationName = "ProvisionedClusterInstancesClient.GetUpgradeProfile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getUpgradeProfileCreateRequest(ctx, connectedClusterResourceURI, options)
	if err != nil {
		return ProvisionedClusterInstancesClientGetUpgradeProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProvisionedClusterInstancesClientGetUpgradeProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProvisionedClusterInstancesClientGetUpgradeProfileResponse{}, err
	}
	resp, err := client.getUpgradeProfileHandleResponse(httpResp)
	return resp, err
}

// getUpgradeProfileCreateRequest creates the GetUpgradeProfile request.
func (client *ProvisionedClusterInstancesClient) getUpgradeProfileCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientGetUpgradeProfileOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/upgradeProfiles/default"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUpgradeProfileHandleResponse handles the GetUpgradeProfile response.
func (client *ProvisionedClusterInstancesClient) getUpgradeProfileHandleResponse(resp *http.Response) (ProvisionedClusterInstancesClientGetUpgradeProfileResponse, error) {
	result := ProvisionedClusterInstancesClientGetUpgradeProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProvisionedClusterUpgradeProfile); err != nil {
		return ProvisionedClusterInstancesClientGetUpgradeProfileResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the Hybrid AKS provisioned cluster instances associated with the connected cluster
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientListOptions contains the optional parameters for the ProvisionedClusterInstancesClient.NewListPager
//     method.
func (client *ProvisionedClusterInstancesClient) NewListPager(connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientListOptions) *runtime.Pager[ProvisionedClusterInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProvisionedClusterInstancesClientListResponse]{
		More: func(page ProvisionedClusterInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProvisionedClusterInstancesClientListResponse) (ProvisionedClusterInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProvisionedClusterInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, connectedClusterResourceURI, options)
			}, nil)
			if err != nil {
				return ProvisionedClusterInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ProvisionedClusterInstancesClient) listCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProvisionedClusterInstancesClient) listHandleResponse(resp *http.Response) (ProvisionedClusterInstancesClientListResponse, error) {
	result := ProvisionedClusterInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProvisionedClustersListResult); err != nil {
		return ProvisionedClusterInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginListAdminKubeconfig - Lists the admin credentials of a provisioned cluster instance used only in direct mode.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientBeginListAdminKubeconfigOptions contains the optional parameters for the ProvisionedClusterInstancesClient.BeginListAdminKubeconfig
//     method.
func (client *ProvisionedClusterInstancesClient) BeginListAdminKubeconfig(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginListAdminKubeconfigOptions) (*runtime.Poller[ProvisionedClusterInstancesClientListAdminKubeconfigResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listAdminKubeconfig(ctx, connectedClusterResourceURI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProvisionedClusterInstancesClientListAdminKubeconfigResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProvisionedClusterInstancesClientListAdminKubeconfigResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListAdminKubeconfig - Lists the admin credentials of a provisioned cluster instance used only in direct mode.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
func (client *ProvisionedClusterInstancesClient) listAdminKubeconfig(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginListAdminKubeconfigOptions) (*http.Response, error) {
	var err error
	const operationName = "ProvisionedClusterInstancesClient.BeginListAdminKubeconfig"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAdminKubeconfigCreateRequest(ctx, connectedClusterResourceURI, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// listAdminKubeconfigCreateRequest creates the ListAdminKubeconfig request.
func (client *ProvisionedClusterInstancesClient) listAdminKubeconfigCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginListAdminKubeconfigOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/listAdminKubeconfig"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ListUpgradeProfile - Lists the upgrade profiles of a provisioned cluster instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientListUpgradeProfileOptions contains the optional parameters for the ProvisionedClusterInstancesClient.ListUpgradeProfile
//     method.
func (client *ProvisionedClusterInstancesClient) ListUpgradeProfile(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientListUpgradeProfileOptions) (ProvisionedClusterInstancesClientListUpgradeProfileResponse, error) {
	var err error
	const operationName = "ProvisionedClusterInstancesClient.ListUpgradeProfile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listUpgradeProfileCreateRequest(ctx, connectedClusterResourceURI, options)
	if err != nil {
		return ProvisionedClusterInstancesClientListUpgradeProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProvisionedClusterInstancesClientListUpgradeProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProvisionedClusterInstancesClientListUpgradeProfileResponse{}, err
	}
	resp, err := client.listUpgradeProfileHandleResponse(httpResp)
	return resp, err
}

// listUpgradeProfileCreateRequest creates the ListUpgradeProfile request.
func (client *ProvisionedClusterInstancesClient) listUpgradeProfileCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientListUpgradeProfileOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/upgradeProfiles"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUpgradeProfileHandleResponse handles the ListUpgradeProfile response.
func (client *ProvisionedClusterInstancesClient) listUpgradeProfileHandleResponse(resp *http.Response) (ProvisionedClusterInstancesClientListUpgradeProfileResponse, error) {
	result := ProvisionedClusterInstancesClientListUpgradeProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProvisionedClusterUpgradeProfileList); err != nil {
		return ProvisionedClusterInstancesClientListUpgradeProfileResponse{}, err
	}
	return result, nil
}

// BeginListUserKubeconfig - Lists the AAD user credentials of a provisioned cluster instance used only in direct mode.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - ProvisionedClusterInstancesClientBeginListUserKubeconfigOptions contains the optional parameters for the ProvisionedClusterInstancesClient.BeginListUserKubeconfig
//     method.
func (client *ProvisionedClusterInstancesClient) BeginListUserKubeconfig(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginListUserKubeconfigOptions) (*runtime.Poller[ProvisionedClusterInstancesClientListUserKubeconfigResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listUserKubeconfig(ctx, connectedClusterResourceURI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProvisionedClusterInstancesClientListUserKubeconfigResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProvisionedClusterInstancesClientListUserKubeconfigResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListUserKubeconfig - Lists the AAD user credentials of a provisioned cluster instance used only in direct mode.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
func (client *ProvisionedClusterInstancesClient) listUserKubeconfig(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginListUserKubeconfigOptions) (*http.Response, error) {
	var err error
	const operationName = "ProvisionedClusterInstancesClient.BeginListUserKubeconfig"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listUserKubeconfigCreateRequest(ctx, connectedClusterResourceURI, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// listUserKubeconfigCreateRequest creates the ListUserKubeconfig request.
func (client *ProvisionedClusterInstancesClient) listUserKubeconfigCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *ProvisionedClusterInstancesClientBeginListUserKubeconfigOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/listUserKubeconfig"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
