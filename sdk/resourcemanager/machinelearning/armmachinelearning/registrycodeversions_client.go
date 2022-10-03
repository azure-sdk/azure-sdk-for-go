//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// RegistryCodeVersionsClient contains the methods for the RegistryCodeVersions group.
// Don't use this type directly, use NewRegistryCodeVersionsClient() instead.
type RegistryCodeVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRegistryCodeVersionsClient creates a new instance of RegistryCodeVersionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegistryCodeVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryCodeVersionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryCodeVersionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// registryName - Name of Azure Machine Learning registry.
// codeName - Container name.
// version - Version identifier.
// body - Version entity to create or update.
// options - RegistryCodeVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistryCodeVersionsClient.BeginCreateOrUpdate
// method.
func (client *RegistryCodeVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, body CodeVersion, options *RegistryCodeVersionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistryCodeVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, registryName, codeName, version, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[RegistryCodeVersionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryCodeVersionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
func (client *RegistryCodeVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, body CodeVersion, options *RegistryCodeVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, registryName, codeName, version, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegistryCodeVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, body CodeVersion, options *RegistryCodeVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes/{codeName}/versions/{version}"
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
	if codeName == "" {
		return nil, errors.New("parameter codeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeName}", url.PathEscape(codeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// registryName - Name of Azure Machine Learning registry.
// codeName - Container name.
// version - Version identifier.
// options - RegistryCodeVersionsClientBeginDeleteOptions contains the optional parameters for the RegistryCodeVersionsClient.BeginDelete
// method.
func (client *RegistryCodeVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, options *RegistryCodeVersionsClientBeginDeleteOptions) (*runtime.Poller[RegistryCodeVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, codeName, version, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[RegistryCodeVersionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryCodeVersionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
func (client *RegistryCodeVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, options *RegistryCodeVersionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, codeName, version, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistryCodeVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, options *RegistryCodeVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes/{codeName}/versions/{version}"
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
	if codeName == "" {
		return nil, errors.New("parameter codeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeName}", url.PathEscape(codeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// registryName - Name of Azure Machine Learning registry.
// codeName - Container name.
// version - Version identifier.
// options - RegistryCodeVersionsClientGetOptions contains the optional parameters for the RegistryCodeVersionsClient.Get
// method.
func (client *RegistryCodeVersionsClient) Get(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, options *RegistryCodeVersionsClientGetOptions) (RegistryCodeVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, codeName, version, options)
	if err != nil {
		return RegistryCodeVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistryCodeVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistryCodeVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistryCodeVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, codeName string, version string, options *RegistryCodeVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes/{codeName}/versions/{version}"
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
	if codeName == "" {
		return nil, errors.New("parameter codeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeName}", url.PathEscape(codeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistryCodeVersionsClient) getHandleResponse(resp *http.Response) (RegistryCodeVersionsClientGetResponse, error) {
	result := RegistryCodeVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CodeVersion); err != nil {
		return RegistryCodeVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List versions.
// Generated from API version 2022-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// registryName - Name of Azure Machine Learning registry.
// codeName - Container name.
// options - RegistryCodeVersionsClientListOptions contains the optional parameters for the RegistryCodeVersionsClient.List
// method.
func (client *RegistryCodeVersionsClient) NewListPager(resourceGroupName string, registryName string, codeName string, options *RegistryCodeVersionsClientListOptions) *runtime.Pager[RegistryCodeVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistryCodeVersionsClientListResponse]{
		More: func(page RegistryCodeVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistryCodeVersionsClientListResponse) (RegistryCodeVersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, codeName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistryCodeVersionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RegistryCodeVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistryCodeVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegistryCodeVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, codeName string, options *RegistryCodeVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes/{codeName}/versions"
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
	if codeName == "" {
		return nil, errors.New("parameter codeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeName}", url.PathEscape(codeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderBy", *options.OrderBy)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistryCodeVersionsClient) listHandleResponse(resp *http.Response) (RegistryCodeVersionsClientListResponse, error) {
	result := RegistryCodeVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CodeVersionResourceArmPaginatedResult); err != nil {
		return RegistryCodeVersionsClientListResponse{}, err
	}
	return result, nil
}
