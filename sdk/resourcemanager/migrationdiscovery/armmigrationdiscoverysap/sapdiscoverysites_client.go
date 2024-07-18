// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmigrationdiscoverysap

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

// SAPDiscoverySitesClient contains the methods for the SAPDiscoverySites group.
// Don't use this type directly, use NewSAPDiscoverySitesClient() instead.
type SAPDiscoverySitesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSAPDiscoverySitesClient creates a new instance of SAPDiscoverySitesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSAPDiscoverySitesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SAPDiscoverySitesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SAPDiscoverySitesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a discovery site resource for SAP Migration. This resource will be used to run system discovery and
// assessment with Azure Migrate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapDiscoverySiteName - The name of the discovery site resource for SAP Migration.
//   - resource - Resource create parameters.
//   - options - SAPDiscoverySitesClientBeginCreateOptions contains the optional parameters for the SAPDiscoverySitesClient.BeginCreate
//     method.
func (client *SAPDiscoverySitesClient) BeginCreate(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, resource SAPDiscoverySite, options *SAPDiscoverySitesClientBeginCreateOptions) (*runtime.Poller[SAPDiscoverySitesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sapDiscoverySiteName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SAPDiscoverySitesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SAPDiscoverySitesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a discovery site resource for SAP Migration. This resource will be used to run system discovery and assessment
// with Azure Migrate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *SAPDiscoverySitesClient) create(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, resource SAPDiscoverySite, options *SAPDiscoverySitesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SAPDiscoverySitesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, sapDiscoverySiteName, resource, options)
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

// createCreateRequest creates the Create request.
func (client *SAPDiscoverySitesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, resource SAPDiscoverySite, _ *SAPDiscoverySitesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapDiscoverySites/{sapDiscoverySiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapDiscoverySiteName == "" {
		return nil, errors.New("parameter sapDiscoverySiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapDiscoverySiteName}", url.PathEscape(sapDiscoverySiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a SAP Migration discovery site resource and its child resources, that is the associated SAP Instances
// and Server Instances.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapDiscoverySiteName - The name of the discovery site resource for SAP Migration.
//   - options - SAPDiscoverySitesClientBeginDeleteOptions contains the optional parameters for the SAPDiscoverySitesClient.BeginDelete
//     method.
func (client *SAPDiscoverySitesClient) BeginDelete(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *SAPDiscoverySitesClientBeginDeleteOptions) (*runtime.Poller[SAPDiscoverySitesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sapDiscoverySiteName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SAPDiscoverySitesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SAPDiscoverySitesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a SAP Migration discovery site resource and its child resources, that is the associated SAP Instances
// and Server Instances.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *SAPDiscoverySitesClient) deleteOperation(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *SAPDiscoverySitesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SAPDiscoverySitesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sapDiscoverySiteName, options)
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
func (client *SAPDiscoverySitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, _ *SAPDiscoverySitesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapDiscoverySites/{sapDiscoverySiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapDiscoverySiteName == "" {
		return nil, errors.New("parameter sapDiscoverySiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapDiscoverySiteName}", url.PathEscape(sapDiscoverySiteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a SAP Migration discovery site resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapDiscoverySiteName - The name of the discovery site resource for SAP Migration.
//   - options - SAPDiscoverySitesClientGetOptions contains the optional parameters for the SAPDiscoverySitesClient.Get method.
func (client *SAPDiscoverySitesClient) Get(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *SAPDiscoverySitesClientGetOptions) (SAPDiscoverySitesClientGetResponse, error) {
	var err error
	const operationName = "SAPDiscoverySitesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, sapDiscoverySiteName, options)
	if err != nil {
		return SAPDiscoverySitesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SAPDiscoverySitesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SAPDiscoverySitesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SAPDiscoverySitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, _ *SAPDiscoverySitesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapDiscoverySites/{sapDiscoverySiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapDiscoverySiteName == "" {
		return nil, errors.New("parameter sapDiscoverySiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapDiscoverySiteName}", url.PathEscape(sapDiscoverySiteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SAPDiscoverySitesClient) getHandleResponse(resp *http.Response) (SAPDiscoverySitesClientGetResponse, error) {
	result := SAPDiscoverySitesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDiscoverySite); err != nil {
		return SAPDiscoverySitesClientGetResponse{}, err
	}
	return result, nil
}

// BeginImportEntities - Import your SAP systems' inventory using the [Discovery template](https://go.microsoft.com/fwlink/?linkid=2249111)
// into your SAP Migration discovery site resource and it's child resources, the SAP instances and Server instances.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapDiscoverySiteName - The name of the discovery site resource for SAP Migration.
//   - options - SAPDiscoverySitesClientBeginImportEntitiesOptions contains the optional parameters for the SAPDiscoverySitesClient.BeginImportEntities
//     method.
func (client *SAPDiscoverySitesClient) BeginImportEntities(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *SAPDiscoverySitesClientBeginImportEntitiesOptions) (*runtime.Poller[SAPDiscoverySitesClientImportEntitiesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.importEntities(ctx, resourceGroupName, sapDiscoverySiteName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SAPDiscoverySitesClientImportEntitiesResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SAPDiscoverySitesClientImportEntitiesResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ImportEntities - Import your SAP systems' inventory using the [Discovery template](https://go.microsoft.com/fwlink/?linkid=2249111)
// into your SAP Migration discovery site resource and it's child resources, the SAP instances and Server instances.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *SAPDiscoverySitesClient) importEntities(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *SAPDiscoverySitesClientBeginImportEntitiesOptions) (*http.Response, error) {
	var err error
	const operationName = "SAPDiscoverySitesClient.BeginImportEntities"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.importEntitiesCreateRequest(ctx, resourceGroupName, sapDiscoverySiteName, options)
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

// importEntitiesCreateRequest creates the ImportEntities request.
func (client *SAPDiscoverySitesClient) importEntitiesCreateRequest(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, _ *SAPDiscoverySitesClientBeginImportEntitiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapDiscoverySites/{sapDiscoverySiteName}/importEntities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapDiscoverySiteName == "" {
		return nil, errors.New("parameter sapDiscoverySiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapDiscoverySiteName}", url.PathEscape(sapDiscoverySiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByResourceGroupPager - Gets all SAP Migration discovery site resources in a Resource Group.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SAPDiscoverySitesClientListByResourceGroupOptions contains the optional parameters for the SAPDiscoverySitesClient.NewListByResourceGroupPager
//     method.
func (client *SAPDiscoverySitesClient) NewListByResourceGroupPager(resourceGroupName string, options *SAPDiscoverySitesClientListByResourceGroupOptions) *runtime.Pager[SAPDiscoverySitesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SAPDiscoverySitesClientListByResourceGroupResponse]{
		More: func(page SAPDiscoverySitesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SAPDiscoverySitesClientListByResourceGroupResponse) (SAPDiscoverySitesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SAPDiscoverySitesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SAPDiscoverySitesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SAPDiscoverySitesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *SAPDiscoverySitesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapDiscoverySites"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SAPDiscoverySitesClient) listByResourceGroupHandleResponse(resp *http.Response) (SAPDiscoverySitesClientListByResourceGroupResponse, error) {
	result := SAPDiscoverySitesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDiscoverySiteListResult); err != nil {
		return SAPDiscoverySitesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all SAP Migration discovery site resources in a Subscription.
//
// Generated from API version 2023-10-01-preview
//   - options - SAPDiscoverySitesClientListBySubscriptionOptions contains the optional parameters for the SAPDiscoverySitesClient.NewListBySubscriptionPager
//     method.
func (client *SAPDiscoverySitesClient) NewListBySubscriptionPager(options *SAPDiscoverySitesClientListBySubscriptionOptions) *runtime.Pager[SAPDiscoverySitesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SAPDiscoverySitesClientListBySubscriptionResponse]{
		More: func(page SAPDiscoverySitesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SAPDiscoverySitesClientListBySubscriptionResponse) (SAPDiscoverySitesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SAPDiscoverySitesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SAPDiscoverySitesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SAPDiscoverySitesClient) listBySubscriptionCreateRequest(ctx context.Context, _ *SAPDiscoverySitesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/sapDiscoverySites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SAPDiscoverySitesClient) listBySubscriptionHandleResponse(resp *http.Response) (SAPDiscoverySitesClientListBySubscriptionResponse, error) {
	result := SAPDiscoverySitesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDiscoverySiteListResult); err != nil {
		return SAPDiscoverySitesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates a SAP Migration discovery site resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapDiscoverySiteName - The name of the discovery site resource for SAP Migration.
//   - properties - The resource properties to be updated.
//   - options - SAPDiscoverySitesClientUpdateOptions contains the optional parameters for the SAPDiscoverySitesClient.Update
//     method.
func (client *SAPDiscoverySitesClient) Update(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, properties SAPDiscoverySiteTagsUpdate, options *SAPDiscoverySitesClientUpdateOptions) (SAPDiscoverySitesClientUpdateResponse, error) {
	var err error
	const operationName = "SAPDiscoverySitesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sapDiscoverySiteName, properties, options)
	if err != nil {
		return SAPDiscoverySitesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SAPDiscoverySitesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SAPDiscoverySitesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SAPDiscoverySitesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, properties SAPDiscoverySiteTagsUpdate, _ *SAPDiscoverySitesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapDiscoverySites/{sapDiscoverySiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapDiscoverySiteName == "" {
		return nil, errors.New("parameter sapDiscoverySiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapDiscoverySiteName}", url.PathEscape(sapDiscoverySiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SAPDiscoverySitesClient) updateHandleResponse(resp *http.Response) (SAPDiscoverySitesClientUpdateResponse, error) {
	result := SAPDiscoverySitesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDiscoverySite); err != nil {
		return SAPDiscoverySitesClientUpdateResponse{}, err
	}
	return result, nil
}
