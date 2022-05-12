//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

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

// FluxConfigurationsClient contains the methods for the FluxConfigurations group.
// Don't use this type directly, use NewFluxConfigurationsClient() instead.
type FluxConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFluxConfigurationsClient creates a new instance of FluxConfigurationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFluxConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FluxConfigurationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &FluxConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new Kubernetes Flux Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// fluxConfigurationName - Name of the Flux Configuration.
// fluxConfiguration - Properties necessary to Create a FluxConfiguration.
// options - FluxConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the FluxConfigurationsClient.BeginCreateOrUpdate
// method.
func (client *FluxConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfiguration FluxConfiguration, options *FluxConfigurationsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[FluxConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, fluxConfiguration, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[FluxConfigurationsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[FluxConfigurationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create a new Kubernetes Flux Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FluxConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfiguration FluxConfiguration, options *FluxConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, fluxConfiguration, options)
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
func (client *FluxConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfiguration FluxConfiguration, options *FluxConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if fluxConfigurationName == "" {
		return nil, errors.New("parameter fluxConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluxConfigurationName}", url.PathEscape(fluxConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, fluxConfiguration)
}

// BeginDelete - This will delete the YAML file used to set up the Flux Configuration, thus stopping future sync from the
// source repo.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// fluxConfigurationName - Name of the Flux Configuration.
// options - FluxConfigurationsClientBeginDeleteOptions contains the optional parameters for the FluxConfigurationsClient.BeginDelete
// method.
func (client *FluxConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, options *FluxConfigurationsClientBeginDeleteOptions) (*armruntime.Poller[FluxConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[FluxConfigurationsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[FluxConfigurationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - This will delete the YAML file used to set up the Flux Configuration, thus stopping future sync from the source
// repo.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FluxConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, options *FluxConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, options)
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
func (client *FluxConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, options *FluxConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if fluxConfigurationName == "" {
		return nil, errors.New("parameter fluxConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluxConfigurationName}", url.PathEscape(fluxConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	if options != nil && options.ForceDelete != nil {
		reqQP.Set("forceDelete", strconv.FormatBool(*options.ForceDelete))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets details of the Flux Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// fluxConfigurationName - Name of the Flux Configuration.
// options - FluxConfigurationsClientGetOptions contains the optional parameters for the FluxConfigurationsClient.Get method.
func (client *FluxConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, options *FluxConfigurationsClientGetOptions) (FluxConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, options)
	if err != nil {
		return FluxConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FluxConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FluxConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FluxConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, options *FluxConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if fluxConfigurationName == "" {
		return nil, errors.New("parameter fluxConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluxConfigurationName}", url.PathEscape(fluxConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FluxConfigurationsClient) getHandleResponse(resp *http.Response) (FluxConfigurationsClientGetResponse, error) {
	result := FluxConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FluxConfiguration); err != nil {
		return FluxConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Flux Configurations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// options - FluxConfigurationsClientListOptions contains the optional parameters for the FluxConfigurationsClient.List method.
func (client *FluxConfigurationsClient) NewListPager(resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *FluxConfigurationsClientListOptions) *runtime.Pager[FluxConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[FluxConfigurationsClientListResponse]{
		More: func(page FluxConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FluxConfigurationsClientListResponse) (FluxConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FluxConfigurationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FluxConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FluxConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FluxConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *FluxConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FluxConfigurationsClient) listHandleResponse(resp *http.Response) (FluxConfigurationsClientListResponse, error) {
	result := FluxConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FluxConfigurationsList); err != nil {
		return FluxConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an existing Kubernetes Flux Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
// clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
// clusterName - The name of the kubernetes cluster.
// fluxConfigurationName - Name of the Flux Configuration.
// fluxConfigurationPatch - Properties to Patch in an existing Flux Configuration.
// options - FluxConfigurationsClientBeginUpdateOptions contains the optional parameters for the FluxConfigurationsClient.BeginUpdate
// method.
func (client *FluxConfigurationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfigurationPatch FluxConfigurationPatch, options *FluxConfigurationsClientBeginUpdateOptions) (*armruntime.Poller[FluxConfigurationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, fluxConfigurationPatch, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[FluxConfigurationsClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[FluxConfigurationsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update an existing Kubernetes Flux Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FluxConfigurationsClient) update(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfigurationPatch FluxConfigurationPatch, options *FluxConfigurationsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, fluxConfigurationPatch, options)
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

// updateCreateRequest creates the Update request.
func (client *FluxConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfigurationPatch FluxConfigurationPatch, options *FluxConfigurationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if fluxConfigurationName == "" {
		return nil, errors.New("parameter fluxConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluxConfigurationName}", url.PathEscape(fluxConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, fluxConfigurationPatch)
}
