// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

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

// RegistryEnvironmentContainersClient contains the methods for the RegistryEnvironmentContainers group.
// Don't use this type directly, use NewRegistryEnvironmentContainersClient() instead.
type RegistryEnvironmentContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegistryEnvironmentContainersClient creates a new instance of RegistryEnvironmentContainersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistryEnvironmentContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryEnvironmentContainersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryEnvironmentContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - environmentName - Container name.
//   - body - Container entity to create or update.
//   - options - RegistryEnvironmentContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistryEnvironmentContainersClient.BeginCreateOrUpdate
//     method.
func (client *RegistryEnvironmentContainersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, environmentName string, body EnvironmentContainer, options *RegistryEnvironmentContainersClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistryEnvironmentContainersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, registryName, environmentName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryEnvironmentContainersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RegistryEnvironmentContainersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *RegistryEnvironmentContainersClient) createOrUpdate(ctx context.Context, resourceGroupName string, registryName string, environmentName string, body EnvironmentContainer, options *RegistryEnvironmentContainersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RegistryEnvironmentContainersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, registryName, environmentName, body, options)
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
func (client *RegistryEnvironmentContainersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, environmentName string, body EnvironmentContainer, _ *RegistryEnvironmentContainersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/environments/{environmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - environmentName - Container name.
//   - options - RegistryEnvironmentContainersClientBeginDeleteOptions contains the optional parameters for the RegistryEnvironmentContainersClient.BeginDelete
//     method.
func (client *RegistryEnvironmentContainersClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, environmentName string, options *RegistryEnvironmentContainersClientBeginDeleteOptions) (*runtime.Poller[RegistryEnvironmentContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, environmentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryEnvironmentContainersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RegistryEnvironmentContainersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *RegistryEnvironmentContainersClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, environmentName string, options *RegistryEnvironmentContainersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RegistryEnvironmentContainersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, environmentName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistryEnvironmentContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, environmentName string, _ *RegistryEnvironmentContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/environments/{environmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - environmentName - Container name. This is case-sensitive.
//   - options - RegistryEnvironmentContainersClientGetOptions contains the optional parameters for the RegistryEnvironmentContainersClient.Get
//     method.
func (client *RegistryEnvironmentContainersClient) Get(ctx context.Context, resourceGroupName string, registryName string, environmentName string, options *RegistryEnvironmentContainersClientGetOptions) (RegistryEnvironmentContainersClientGetResponse, error) {
	var err error
	const operationName = "RegistryEnvironmentContainersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, environmentName, options)
	if err != nil {
		return RegistryEnvironmentContainersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryEnvironmentContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RegistryEnvironmentContainersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RegistryEnvironmentContainersClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, environmentName string, _ *RegistryEnvironmentContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/environments/{environmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistryEnvironmentContainersClient) getHandleResponse(resp *http.Response) (RegistryEnvironmentContainersClientGetResponse, error) {
	result := RegistryEnvironmentContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentContainer); err != nil {
		return RegistryEnvironmentContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List environment containers.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - options - RegistryEnvironmentContainersClientListOptions contains the optional parameters for the RegistryEnvironmentContainersClient.NewListPager
//     method.
func (client *RegistryEnvironmentContainersClient) NewListPager(resourceGroupName string, registryName string, options *RegistryEnvironmentContainersClientListOptions) *runtime.Pager[RegistryEnvironmentContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistryEnvironmentContainersClientListResponse]{
		More: func(page RegistryEnvironmentContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistryEnvironmentContainersClientListResponse) (RegistryEnvironmentContainersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RegistryEnvironmentContainersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			}, nil)
			if err != nil {
				return RegistryEnvironmentContainersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RegistryEnvironmentContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistryEnvironmentContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/environments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2025-01-01-preview")
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistryEnvironmentContainersClient) listHandleResponse(resp *http.Response) (RegistryEnvironmentContainersClientListResponse, error) {
	result := RegistryEnvironmentContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentContainerResourceArmPaginatedResult); err != nil {
		return RegistryEnvironmentContainersClientListResponse{}, err
	}
	return result, nil
}
