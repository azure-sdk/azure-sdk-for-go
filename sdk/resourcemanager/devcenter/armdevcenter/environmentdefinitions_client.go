//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

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

// EnvironmentDefinitionsClient contains the methods for the EnvironmentDefinitions group.
// Don't use this type directly, use NewEnvironmentDefinitionsClient() instead.
type EnvironmentDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEnvironmentDefinitionsClient creates a new instance of EnvironmentDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEnvironmentDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnvironmentDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName+".EnvironmentDefinitionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EnvironmentDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an environment definition from the catalog.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - environmentDefinitionName - The name of the Environment Definition.
//   - options - EnvironmentDefinitionsClientGetOptions contains the optional parameters for the EnvironmentDefinitionsClient.Get
//     method.
func (client *EnvironmentDefinitionsClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, environmentDefinitionName string, options *EnvironmentDefinitionsClientGetOptions) (EnvironmentDefinitionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, environmentDefinitionName, options)
	if err != nil {
		return EnvironmentDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EnvironmentDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, environmentDefinitionName string, options *EnvironmentDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/environmentDefinitions/{environmentDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if environmentDefinitionName == "" {
		return nil, errors.New("parameter environmentDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentDefinitionName}", url.PathEscape(environmentDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EnvironmentDefinitionsClient) getHandleResponse(resp *http.Response) (EnvironmentDefinitionsClientGetResponse, error) {
	result := EnvironmentDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentDefinition); err != nil {
		return EnvironmentDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetErrorDetails - Gets Environment Definition error details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - environmentDefinitionName - The name of the Environment Definition.
//   - options - EnvironmentDefinitionsClientGetErrorDetailsOptions contains the optional parameters for the EnvironmentDefinitionsClient.GetErrorDetails
//     method.
func (client *EnvironmentDefinitionsClient) GetErrorDetails(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, environmentDefinitionName string, options *EnvironmentDefinitionsClientGetErrorDetailsOptions) (EnvironmentDefinitionsClientGetErrorDetailsResponse, error) {
	var err error
	req, err := client.getErrorDetailsCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, environmentDefinitionName, options)
	if err != nil {
		return EnvironmentDefinitionsClientGetErrorDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentDefinitionsClientGetErrorDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentDefinitionsClientGetErrorDetailsResponse{}, err
	}
	resp, err := client.getErrorDetailsHandleResponse(httpResp)
	return resp, err
}

// getErrorDetailsCreateRequest creates the GetErrorDetails request.
func (client *EnvironmentDefinitionsClient) getErrorDetailsCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, environmentDefinitionName string, options *EnvironmentDefinitionsClientGetErrorDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/environmentDefinitions/{environmentDefinitionName}/getErrorDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if environmentDefinitionName == "" {
		return nil, errors.New("parameter environmentDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentDefinitionName}", url.PathEscape(environmentDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getErrorDetailsHandleResponse handles the GetErrorDetails response.
func (client *EnvironmentDefinitionsClient) getErrorDetailsHandleResponse(resp *http.Response) (EnvironmentDefinitionsClientGetErrorDetailsResponse, error) {
	result := EnvironmentDefinitionsClientGetErrorDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CatalogResourceValidationErrorDetails); err != nil {
		return EnvironmentDefinitionsClientGetErrorDetailsResponse{}, err
	}
	return result, nil
}

// NewListByCatalogPager - List environment definitions in the catalog.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - EnvironmentDefinitionsClientListByCatalogOptions contains the optional parameters for the EnvironmentDefinitionsClient.NewListByCatalogPager
//     method.
func (client *EnvironmentDefinitionsClient) NewListByCatalogPager(resourceGroupName string, devCenterName string, catalogName string, options *EnvironmentDefinitionsClientListByCatalogOptions) *runtime.Pager[EnvironmentDefinitionsClientListByCatalogResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnvironmentDefinitionsClientListByCatalogResponse]{
		More: func(page EnvironmentDefinitionsClientListByCatalogResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnvironmentDefinitionsClientListByCatalogResponse) (EnvironmentDefinitionsClientListByCatalogResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByCatalogCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EnvironmentDefinitionsClientListByCatalogResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EnvironmentDefinitionsClientListByCatalogResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EnvironmentDefinitionsClientListByCatalogResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByCatalogHandleResponse(resp)
		},
	})
}

// listByCatalogCreateRequest creates the ListByCatalog request.
func (client *EnvironmentDefinitionsClient) listByCatalogCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *EnvironmentDefinitionsClientListByCatalogOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/environmentDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCatalogHandleResponse handles the ListByCatalog response.
func (client *EnvironmentDefinitionsClient) listByCatalogHandleResponse(resp *http.Response) (EnvironmentDefinitionsClientListByCatalogResponse, error) {
	result := EnvironmentDefinitionsClientListByCatalogResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentDefinitionListResult); err != nil {
		return EnvironmentDefinitionsClientListByCatalogResponse{}, err
	}
	return result, nil
}
