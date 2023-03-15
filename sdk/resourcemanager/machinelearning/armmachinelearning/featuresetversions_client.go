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
	"strings"
)

// FeaturesetVersionsClient contains the methods for the FeaturesetVersions group.
// Don't use this type directly, use NewFeaturesetVersionsClient() instead.
type FeaturesetVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFeaturesetVersionsClient creates a new instance of FeaturesetVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFeaturesetVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FeaturesetVersionsClient, error) {
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
	client := &FeaturesetVersionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginBackfill - Backfill.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - body - Feature set version backfill request entity.
//   - options - FeaturesetVersionsClientBeginBackfillOptions contains the optional parameters for the FeaturesetVersionsClient.BeginBackfill
//     method.
func (client *FeaturesetVersionsClient) BeginBackfill(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body FeaturesetVersionBackfillRequest, options *FeaturesetVersionsClientBeginBackfillOptions) (*runtime.Poller[FeaturesetVersionsClientBackfillResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.backfill(ctx, resourceGroupName, workspaceName, name, version, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FeaturesetVersionsClientBackfillResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FeaturesetVersionsClientBackfillResponse](options.ResumeToken, client.pl, nil)
	}
}

// Backfill - Backfill.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *FeaturesetVersionsClient) backfill(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body FeaturesetVersionBackfillRequest, options *FeaturesetVersionsClientBeginBackfillOptions) (*http.Response, error) {
	req, err := client.backfillCreateRequest(ctx, resourceGroupName, workspaceName, name, version, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// backfillCreateRequest creates the Backfill request.
func (client *FeaturesetVersionsClient) backfillCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body FeaturesetVersionBackfillRequest, options *FeaturesetVersionsClientBeginBackfillOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions/{version}/backfill"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginCreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - body - Version entity to create or update.
//   - options - FeaturesetVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the FeaturesetVersionsClient.BeginCreateOrUpdate
//     method.
func (client *FeaturesetVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body FeaturesetVersion, options *FeaturesetVersionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FeaturesetVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, name, version, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FeaturesetVersionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FeaturesetVersionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *FeaturesetVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body FeaturesetVersion, options *FeaturesetVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, name, version, body, options)
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
func (client *FeaturesetVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body FeaturesetVersion, options *FeaturesetVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - options - FeaturesetVersionsClientBeginDeleteOptions contains the optional parameters for the FeaturesetVersionsClient.BeginDelete
//     method.
func (client *FeaturesetVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientBeginDeleteOptions) (*runtime.Poller[FeaturesetVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, name, version, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FeaturesetVersionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FeaturesetVersionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *FeaturesetVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
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
func (client *FeaturesetVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - options - FeaturesetVersionsClientGetOptions contains the optional parameters for the FeaturesetVersionsClient.Get method.
func (client *FeaturesetVersionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientGetOptions) (FeaturesetVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
	if err != nil {
		return FeaturesetVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FeaturesetVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FeaturesetVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FeaturesetVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FeaturesetVersionsClient) getHandleResponse(resp *http.Response) (FeaturesetVersionsClientGetResponse, error) {
	result := FeaturesetVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeaturesetVersion); err != nil {
		return FeaturesetVersionsClientGetResponse{}, err
	}
	return result, nil
}

// GetFeature - Get feature.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Feature set name. This is case-sensitive.
//   - version - Feature set version identifier. This is case-sensitive.
//   - body - Feature Name request. This is case-sensitive.
//   - options - FeaturesetVersionsClientGetFeatureOptions contains the optional parameters for the FeaturesetVersionsClient.GetFeature
//     method.
func (client *FeaturesetVersionsClient) GetFeature(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body GetFeatureRequest, options *FeaturesetVersionsClientGetFeatureOptions) (FeaturesetVersionsClientGetFeatureResponse, error) {
	req, err := client.getFeatureCreateRequest(ctx, resourceGroupName, workspaceName, name, version, body, options)
	if err != nil {
		return FeaturesetVersionsClientGetFeatureResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FeaturesetVersionsClientGetFeatureResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FeaturesetVersionsClientGetFeatureResponse{}, runtime.NewResponseError(resp)
	}
	return client.getFeatureHandleResponse(resp)
}

// getFeatureCreateRequest creates the GetFeature request.
func (client *FeaturesetVersionsClient) getFeatureCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body GetFeatureRequest, options *FeaturesetVersionsClientGetFeatureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions/{version}/getFeature"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// getFeatureHandleResponse handles the GetFeature response.
func (client *FeaturesetVersionsClient) getFeatureHandleResponse(resp *http.Response) (FeaturesetVersionsClientGetFeatureResponse, error) {
	result := FeaturesetVersionsClientGetFeatureResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Feature); err != nil {
		return FeaturesetVersionsClientGetFeatureResponse{}, err
	}
	return result, nil
}

// NewListPager - List versions.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Featureset name. This is case-sensitive.
//   - options - FeaturesetVersionsClientListOptions contains the optional parameters for the FeaturesetVersionsClient.NewListPager
//     method.
func (client *FeaturesetVersionsClient) NewListPager(resourceGroupName string, workspaceName string, name string, options *FeaturesetVersionsClientListOptions) *runtime.Pager[FeaturesetVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FeaturesetVersionsClientListResponse]{
		More: func(page FeaturesetVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FeaturesetVersionsClientListResponse) (FeaturesetVersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FeaturesetVersionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FeaturesetVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FeaturesetVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FeaturesetVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *FeaturesetVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FeaturesetVersionsClient) listHandleResponse(resp *http.Response) (FeaturesetVersionsClientListResponse, error) {
	result := FeaturesetVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeaturesetVersionResourceArmPaginatedResult); err != nil {
		return FeaturesetVersionsClientListResponse{}, err
	}
	return result, nil
}

// NewListFeaturesPager - List Features.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Featureset name. This is case-sensitive.
//   - version - Featureset Version identifier. This is case-sensitive.
//   - options - FeaturesetVersionsClientListFeaturesOptions contains the optional parameters for the FeaturesetVersionsClient.NewListFeaturesPager
//     method.
func (client *FeaturesetVersionsClient) NewListFeaturesPager(resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientListFeaturesOptions) *runtime.Pager[FeaturesetVersionsClientListFeaturesResponse] {
	return runtime.NewPager(runtime.PagingHandler[FeaturesetVersionsClientListFeaturesResponse]{
		More: func(page FeaturesetVersionsClientListFeaturesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FeaturesetVersionsClientListFeaturesResponse) (FeaturesetVersionsClientListFeaturesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listFeaturesCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FeaturesetVersionsClientListFeaturesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FeaturesetVersionsClientListFeaturesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FeaturesetVersionsClientListFeaturesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listFeaturesHandleResponse(resp)
		},
	})
}

// listFeaturesCreateRequest creates the ListFeatures request.
func (client *FeaturesetVersionsClient) listFeaturesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientListFeaturesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions/{version}/listFeatures"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("$tags", *options.Tags)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listFeaturesHandleResponse handles the ListFeatures response.
func (client *FeaturesetVersionsClient) listFeaturesHandleResponse(resp *http.Response) (FeaturesetVersionsClientListFeaturesResponse, error) {
	result := FeaturesetVersionsClientListFeaturesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeatureArmPaginatedResult); err != nil {
		return FeaturesetVersionsClientListFeaturesResponse{}, err
	}
	return result, nil
}

// NewListMaterializationJobsPager - List materialization Jobs.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - options - FeaturesetVersionsClientListMaterializationJobsOptions contains the optional parameters for the FeaturesetVersionsClient.NewListMaterializationJobsPager
//     method.
func (client *FeaturesetVersionsClient) NewListMaterializationJobsPager(resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientListMaterializationJobsOptions) *runtime.Pager[FeaturesetVersionsClientListMaterializationJobsResponse] {
	return runtime.NewPager(runtime.PagingHandler[FeaturesetVersionsClientListMaterializationJobsResponse]{
		More: func(page FeaturesetVersionsClientListMaterializationJobsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FeaturesetVersionsClientListMaterializationJobsResponse) (FeaturesetVersionsClientListMaterializationJobsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listMaterializationJobsCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FeaturesetVersionsClientListMaterializationJobsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FeaturesetVersionsClientListMaterializationJobsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FeaturesetVersionsClientListMaterializationJobsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMaterializationJobsHandleResponse(resp)
		},
	})
}

// listMaterializationJobsCreateRequest creates the ListMaterializationJobs request.
func (client *FeaturesetVersionsClient) listMaterializationJobsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *FeaturesetVersionsClientListMaterializationJobsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}/versions/{version}/listMaterializationJobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Filters != nil {
		reqQP.Set("filters", *options.Filters)
	}
	if options != nil && options.FeatureWindowStart != nil {
		reqQP.Set("featureWindowStart", *options.FeatureWindowStart)
	}
	if options != nil && options.FeatureWindowEnd != nil {
		reqQP.Set("featureWindowEnd", *options.FeatureWindowEnd)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMaterializationJobsHandleResponse handles the ListMaterializationJobs response.
func (client *FeaturesetVersionsClient) listMaterializationJobsHandleResponse(resp *http.Response) (FeaturesetVersionsClientListMaterializationJobsResponse, error) {
	result := FeaturesetVersionsClientListMaterializationJobsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeaturesetJobArmPaginatedResult); err != nil {
		return FeaturesetVersionsClientListMaterializationJobsResponse{}, err
	}
	return result, nil
}
