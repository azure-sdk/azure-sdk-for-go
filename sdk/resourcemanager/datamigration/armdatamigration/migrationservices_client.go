//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// MigrationServicesClient contains the methods for the MigrationServices group.
// Don't use this type directly, use NewMigrationServicesClient() instead.
type MigrationServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMigrationServicesClient creates a new instance of MigrationServicesClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMigrationServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MigrationServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MigrationServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - migrationServiceName - Name of the Migration Service.
//   - parameters - Details of MigrationService resource.
//   - options - MigrationServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the MigrationServicesClient.BeginCreateOrUpdate
//     method.
func (client *MigrationServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters MigrationService, options *MigrationServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[MigrationServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, migrationServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MigrationServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MigrationServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *MigrationServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters MigrationService, options *MigrationServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MigrationServicesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, migrationServiceName, parameters, options)
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
func (client *MigrationServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters MigrationService, options *MigrationServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/migrationServices/{migrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrationServiceName == "" {
		return nil, errors.New("parameter migrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationServiceName}", url.PathEscape(migrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - migrationServiceName - Name of the Migration Service.
//   - options - MigrationServicesClientBeginDeleteOptions contains the optional parameters for the MigrationServicesClient.BeginDelete
//     method.
func (client *MigrationServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, migrationServiceName string, options *MigrationServicesClientBeginDeleteOptions) (*runtime.Poller[MigrationServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, migrationServiceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MigrationServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MigrationServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *MigrationServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, migrationServiceName string, options *MigrationServicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MigrationServicesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, migrationServiceName, options)
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
func (client *MigrationServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, migrationServiceName string, options *MigrationServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/migrationServices/{migrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrationServiceName == "" {
		return nil, errors.New("parameter migrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationServiceName}", url.PathEscape(migrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the Database Migration Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - migrationServiceName - Name of the Migration Service.
//   - options - MigrationServicesClientGetOptions contains the optional parameters for the MigrationServicesClient.Get method.
func (client *MigrationServicesClient) Get(ctx context.Context, resourceGroupName string, migrationServiceName string, options *MigrationServicesClientGetOptions) (MigrationServicesClientGetResponse, error) {
	var err error
	const operationName = "MigrationServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, migrationServiceName, options)
	if err != nil {
		return MigrationServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MigrationServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MigrationServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MigrationServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, migrationServiceName string, options *MigrationServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/migrationServices/{migrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrationServiceName == "" {
		return nil, errors.New("parameter migrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationServiceName}", url.PathEscape(migrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MigrationServicesClient) getHandleResponse(resp *http.Response) (MigrationServicesClientGetResponse, error) {
	result := MigrationServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationService); err != nil {
		return MigrationServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve all migration services in the resource group.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - options - MigrationServicesClientListByResourceGroupOptions contains the optional parameters for the MigrationServicesClient.NewListByResourceGroupPager
//     method.
func (client *MigrationServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *MigrationServicesClientListByResourceGroupOptions) *runtime.Pager[MigrationServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[MigrationServicesClientListByResourceGroupResponse]{
		More: func(page MigrationServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MigrationServicesClientListByResourceGroupResponse) (MigrationServicesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MigrationServicesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return MigrationServicesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MigrationServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *MigrationServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/migrationServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MigrationServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (MigrationServicesClientListByResourceGroupResponse, error) {
	result := MigrationServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationServiceListResult); err != nil {
		return MigrationServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieve all migration services in the subscriptions.
//
// Generated from API version 2023-07-15-preview
//   - options - MigrationServicesClientListBySubscriptionOptions contains the optional parameters for the MigrationServicesClient.NewListBySubscriptionPager
//     method.
func (client *MigrationServicesClient) NewListBySubscriptionPager(options *MigrationServicesClientListBySubscriptionOptions) *runtime.Pager[MigrationServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[MigrationServicesClientListBySubscriptionResponse]{
		More: func(page MigrationServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MigrationServicesClientListBySubscriptionResponse) (MigrationServicesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MigrationServicesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return MigrationServicesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MigrationServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *MigrationServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/migrationServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MigrationServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (MigrationServicesClientListBySubscriptionResponse, error) {
	result := MigrationServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationServiceListResult); err != nil {
		return MigrationServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListMigrationsPager - Retrieve the List of database migrations attached to the service.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - migrationServiceName - Name of the Migration Service.
//   - options - MigrationServicesClientListMigrationsOptions contains the optional parameters for the MigrationServicesClient.NewListMigrationsPager
//     method.
func (client *MigrationServicesClient) NewListMigrationsPager(resourceGroupName string, migrationServiceName string, options *MigrationServicesClientListMigrationsOptions) *runtime.Pager[MigrationServicesClientListMigrationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[MigrationServicesClientListMigrationsResponse]{
		More: func(page MigrationServicesClientListMigrationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MigrationServicesClientListMigrationsResponse) (MigrationServicesClientListMigrationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MigrationServicesClient.NewListMigrationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listMigrationsCreateRequest(ctx, resourceGroupName, migrationServiceName, options)
			}, nil)
			if err != nil {
				return MigrationServicesClientListMigrationsResponse{}, err
			}
			return client.listMigrationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listMigrationsCreateRequest creates the ListMigrations request.
func (client *MigrationServicesClient) listMigrationsCreateRequest(ctx context.Context, resourceGroupName string, migrationServiceName string, options *MigrationServicesClientListMigrationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/migrationServices/{migrationServiceName}/listMigrations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrationServiceName == "" {
		return nil, errors.New("parameter migrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationServiceName}", url.PathEscape(migrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMigrationsHandleResponse handles the ListMigrations response.
func (client *MigrationServicesClient) listMigrationsHandleResponse(resp *http.Response) (MigrationServicesClientListMigrationsResponse, error) {
	result := MigrationServicesClientListMigrationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseMigrationBaseListResult); err != nil {
		return MigrationServicesClientListMigrationsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - migrationServiceName - Name of the Migration Service.
//   - parameters - Details of MigrationService resource.
//   - options - MigrationServicesClientBeginUpdateOptions contains the optional parameters for the MigrationServicesClient.BeginUpdate
//     method.
func (client *MigrationServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters MigrationServiceUpdate, options *MigrationServicesClientBeginUpdateOptions) (*runtime.Poller[MigrationServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, migrationServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MigrationServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MigrationServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update Database Migration Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *MigrationServicesClient) update(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters MigrationServiceUpdate, options *MigrationServicesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MigrationServicesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, migrationServiceName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *MigrationServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters MigrationServiceUpdate, options *MigrationServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/migrationServices/{migrationServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrationServiceName == "" {
		return nil, errors.New("parameter migrationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationServiceName}", url.PathEscape(migrationServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
