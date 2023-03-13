//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection

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

// BackupVaultsClient contains the methods for the BackupVaults group.
// Don't use this type directly, use NewBackupVaultsClient() instead.
type BackupVaultsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupVaultsClient creates a new instance of BackupVaultsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupVaultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupVaultsClient, error) {
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
	client := &BackupVaultsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - API to check for resource name availability
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - location - The location in which uniqueness will be verified.
//   - parameters - Check name availability request
//   - options - BackupVaultsClientCheckNameAvailabilityOptions contains the optional parameters for the BackupVaultsClient.CheckNameAvailability
//     method.
func (client *BackupVaultsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, location string, parameters CheckNameAvailabilityRequest, options *BackupVaultsClientCheckNameAvailabilityOptions) (BackupVaultsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, location, parameters, options)
	if err != nil {
		return BackupVaultsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupVaultsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupVaultsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *BackupVaultsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, location string, parameters CheckNameAvailabilityRequest, options *BackupVaultsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/locations/{location}/checkNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *BackupVaultsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (BackupVaultsClientCheckNameAvailabilityResponse, error) {
	result := BackupVaultsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return BackupVaultsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates a BackupVault resource belonging to a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - vaultName - The name of the backup vault.
//   - parameters - Request body for operation
//   - options - BackupVaultsClientBeginCreateOrUpdateOptions contains the optional parameters for the BackupVaultsClient.BeginCreateOrUpdate
//     method.
func (client *BackupVaultsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters BackupVaultResource, options *BackupVaultsClientBeginCreateOrUpdateOptions) (*runtime.Poller[BackupVaultsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vaultName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BackupVaultsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[BackupVaultsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a BackupVault resource belonging to a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
func (client *BackupVaultsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters BackupVaultResource, options *BackupVaultsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, parameters, options)
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
func (client *BackupVaultsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, parameters BackupVaultResource, options *BackupVaultsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes a BackupVault resource from the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - vaultName - The name of the backup vault.
//   - options - BackupVaultsClientDeleteOptions contains the optional parameters for the BackupVaultsClient.Delete method.
func (client *BackupVaultsClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, options *BackupVaultsClientDeleteOptions) (BackupVaultsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return BackupVaultsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupVaultsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return BackupVaultsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return BackupVaultsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BackupVaultsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *BackupVaultsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
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

// Get - Returns a resource belonging to a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - vaultName - The name of the backup vault.
//   - options - BackupVaultsClientGetOptions contains the optional parameters for the BackupVaultsClient.Get method.
func (client *BackupVaultsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, options *BackupVaultsClientGetOptions) (BackupVaultsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return BackupVaultsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupVaultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupVaultsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupVaultsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *BackupVaultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
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
func (client *BackupVaultsClient) getHandleResponse(resp *http.Response) (BackupVaultsClientGetResponse, error) {
	result := BackupVaultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupVaultResource); err != nil {
		return BackupVaultsClientGetResponse{}, err
	}
	return result, nil
}

// NewGetInResourceGroupPager - Returns resource collection belonging to a resource group.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - options - BackupVaultsClientGetInResourceGroupOptions contains the optional parameters for the BackupVaultsClient.NewGetInResourceGroupPager
//     method.
func (client *BackupVaultsClient) NewGetInResourceGroupPager(resourceGroupName string, options *BackupVaultsClientGetInResourceGroupOptions) *runtime.Pager[BackupVaultsClientGetInResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupVaultsClientGetInResourceGroupResponse]{
		More: func(page BackupVaultsClientGetInResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupVaultsClientGetInResourceGroupResponse) (BackupVaultsClientGetInResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getInResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BackupVaultsClientGetInResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BackupVaultsClientGetInResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupVaultsClientGetInResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.getInResourceGroupHandleResponse(resp)
		},
	})
}

// getInResourceGroupCreateRequest creates the GetInResourceGroup request.
func (client *BackupVaultsClient) getInResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *BackupVaultsClientGetInResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// getInResourceGroupHandleResponse handles the GetInResourceGroup response.
func (client *BackupVaultsClient) getInResourceGroupHandleResponse(resp *http.Response) (BackupVaultsClientGetInResourceGroupResponse, error) {
	result := BackupVaultsClientGetInResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupVaultResourceList); err != nil {
		return BackupVaultsClientGetInResourceGroupResponse{}, err
	}
	return result, nil
}

// NewGetInSubscriptionPager - Returns resource collection belonging to a subscription.
//
// Generated from API version 2022-10-01-preview
//   - options - BackupVaultsClientGetInSubscriptionOptions contains the optional parameters for the BackupVaultsClient.NewGetInSubscriptionPager
//     method.
func (client *BackupVaultsClient) NewGetInSubscriptionPager(options *BackupVaultsClientGetInSubscriptionOptions) *runtime.Pager[BackupVaultsClientGetInSubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupVaultsClientGetInSubscriptionResponse]{
		More: func(page BackupVaultsClientGetInSubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupVaultsClientGetInSubscriptionResponse) (BackupVaultsClientGetInSubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getInSubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BackupVaultsClientGetInSubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BackupVaultsClientGetInSubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupVaultsClientGetInSubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.getInSubscriptionHandleResponse(resp)
		},
	})
}

// getInSubscriptionCreateRequest creates the GetInSubscription request.
func (client *BackupVaultsClient) getInSubscriptionCreateRequest(ctx context.Context, options *BackupVaultsClientGetInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataProtection/backupVaults"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// getInSubscriptionHandleResponse handles the GetInSubscription response.
func (client *BackupVaultsClient) getInSubscriptionHandleResponse(resp *http.Response) (BackupVaultsClientGetInSubscriptionResponse, error) {
	result := BackupVaultsClientGetInSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupVaultResourceList); err != nil {
		return BackupVaultsClientGetInSubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a BackupVault resource belonging to a resource group. For example, updating tags for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - vaultName - The name of the backup vault.
//   - parameters - Request body for operation
//   - options - BackupVaultsClientBeginUpdateOptions contains the optional parameters for the BackupVaultsClient.BeginUpdate
//     method.
func (client *BackupVaultsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters PatchResourceRequestInput, options *BackupVaultsClientBeginUpdateOptions) (*runtime.Poller[BackupVaultsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vaultName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BackupVaultsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[BackupVaultsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a BackupVault resource belonging to a resource group. For example, updating tags for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
func (client *BackupVaultsClient) update(ctx context.Context, resourceGroupName string, vaultName string, parameters PatchResourceRequestInput, options *BackupVaultsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, parameters, options)
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
func (client *BackupVaultsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, parameters PatchResourceRequestInput, options *BackupVaultsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
