//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// DiskEncryptionSetsClient contains the methods for the DiskEncryptionSets group.
// Don't use this type directly, use NewDiskEncryptionSetsClient() instead.
type DiskEncryptionSetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiskEncryptionSetsClient creates a new instance of DiskEncryptionSetsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiskEncryptionSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiskEncryptionSetsClient, error) {
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
	client := &DiskEncryptionSetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a disk encryption set
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum
// name length is 80 characters.
// diskEncryptionSet - disk encryption set object supplied in the body of the Put disk encryption set operation.
// options - DiskEncryptionSetsClientBeginCreateOrUpdateOptions contains the optional parameters for the DiskEncryptionSetsClient.BeginCreateOrUpdate
// method.
func (client *DiskEncryptionSetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[DiskEncryptionSetsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DiskEncryptionSetsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DiskEncryptionSetsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a disk encryption set
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskEncryptionSetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiskEncryptionSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskEncryptionSet)
}

// BeginDelete - Deletes a disk encryption set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum
// name length is 80 characters.
// options - DiskEncryptionSetsClientBeginDeleteOptions contains the optional parameters for the DiskEncryptionSetsClient.BeginDelete
// method.
func (client *DiskEncryptionSetsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsClientBeginDeleteOptions) (*armruntime.Poller[DiskEncryptionSetsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, diskEncryptionSetName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DiskEncryptionSetsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DiskEncryptionSetsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a disk encryption set.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskEncryptionSetsClient) deleteOperation(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
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
func (client *DiskEncryptionSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about a disk encryption set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum
// name length is 80 characters.
// options - DiskEncryptionSetsClientGetOptions contains the optional parameters for the DiskEncryptionSetsClient.Get method.
func (client *DiskEncryptionSetsClient) Get(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsClientGetOptions) (DiskEncryptionSetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
	if err != nil {
		return DiskEncryptionSetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskEncryptionSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskEncryptionSetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskEncryptionSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskEncryptionSetsClient) getHandleResponse(resp *http.Response) (DiskEncryptionSetsClientGetResponse, error) {
	result := DiskEncryptionSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskEncryptionSet); err != nil {
		return DiskEncryptionSetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the disk encryption sets under a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DiskEncryptionSetsClientListOptions contains the optional parameters for the DiskEncryptionSetsClient.List method.
func (client *DiskEncryptionSetsClient) NewListPager(options *DiskEncryptionSetsClientListOptions) *runtime.Pager[DiskEncryptionSetsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiskEncryptionSetsClientListResponse]{
		More: func(page DiskEncryptionSetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskEncryptionSetsClientListResponse) (DiskEncryptionSetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiskEncryptionSetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiskEncryptionSetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiskEncryptionSetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DiskEncryptionSetsClient) listCreateRequest(ctx context.Context, options *DiskEncryptionSetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiskEncryptionSetsClient) listHandleResponse(resp *http.Response) (DiskEncryptionSetsClientListResponse, error) {
	result := DiskEncryptionSetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskEncryptionSetList); err != nil {
		return DiskEncryptionSetsClientListResponse{}, err
	}
	return result, nil
}

// NewListAssociatedResourcesPager - Lists all resources that are encrypted with this disk encryption set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum
// name length is 80 characters.
// options - DiskEncryptionSetsClientListAssociatedResourcesOptions contains the optional parameters for the DiskEncryptionSetsClient.ListAssociatedResources
// method.
func (client *DiskEncryptionSetsClient) NewListAssociatedResourcesPager(resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsClientListAssociatedResourcesOptions) *runtime.Pager[DiskEncryptionSetsClientListAssociatedResourcesResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiskEncryptionSetsClientListAssociatedResourcesResponse]{
		More: func(page DiskEncryptionSetsClientListAssociatedResourcesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskEncryptionSetsClientListAssociatedResourcesResponse) (DiskEncryptionSetsClientListAssociatedResourcesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAssociatedResourcesCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiskEncryptionSetsClientListAssociatedResourcesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiskEncryptionSetsClientListAssociatedResourcesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiskEncryptionSetsClientListAssociatedResourcesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAssociatedResourcesHandleResponse(resp)
		},
	})
}

// listAssociatedResourcesCreateRequest creates the ListAssociatedResources request.
func (client *DiskEncryptionSetsClient) listAssociatedResourcesCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsClientListAssociatedResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}/associatedResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAssociatedResourcesHandleResponse handles the ListAssociatedResources response.
func (client *DiskEncryptionSetsClient) listAssociatedResourcesHandleResponse(resp *http.Response) (DiskEncryptionSetsClientListAssociatedResourcesResponse, error) {
	result := DiskEncryptionSetsClientListAssociatedResourcesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceURIList); err != nil {
		return DiskEncryptionSetsClientListAssociatedResourcesResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the disk encryption sets under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - DiskEncryptionSetsClientListByResourceGroupOptions contains the optional parameters for the DiskEncryptionSetsClient.ListByResourceGroup
// method.
func (client *DiskEncryptionSetsClient) NewListByResourceGroupPager(resourceGroupName string, options *DiskEncryptionSetsClientListByResourceGroupOptions) *runtime.Pager[DiskEncryptionSetsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiskEncryptionSetsClientListByResourceGroupResponse]{
		More: func(page DiskEncryptionSetsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskEncryptionSetsClientListByResourceGroupResponse) (DiskEncryptionSetsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiskEncryptionSetsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiskEncryptionSetsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiskEncryptionSetsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DiskEncryptionSetsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DiskEncryptionSetsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DiskEncryptionSetsClient) listByResourceGroupHandleResponse(resp *http.Response) (DiskEncryptionSetsClientListByResourceGroupResponse, error) {
	result := DiskEncryptionSetsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskEncryptionSetList); err != nil {
		return DiskEncryptionSetsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates (patches) a disk encryption set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskEncryptionSetName - The name of the disk encryption set that is being created. The name can't be changed after the
// disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum
// name length is 80 characters.
// diskEncryptionSet - disk encryption set object supplied in the body of the Patch disk encryption set operation.
// options - DiskEncryptionSetsClientBeginUpdateOptions contains the optional parameters for the DiskEncryptionSetsClient.BeginUpdate
// method.
func (client *DiskEncryptionSetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsClientBeginUpdateOptions) (*armruntime.Poller[DiskEncryptionSetsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DiskEncryptionSetsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DiskEncryptionSetsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates (patches) a disk encryption set.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskEncryptionSetsClient) update(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
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
func (client *DiskEncryptionSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskEncryptionSet)
}
