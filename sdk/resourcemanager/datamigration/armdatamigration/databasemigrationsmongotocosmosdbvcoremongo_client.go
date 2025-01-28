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
	"strconv"
	"strings"
)

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClient contains the methods for the DatabaseMigrationsMongoToCosmosDbvCoreMongo group.
// Don't use this type directly, use NewDatabaseMigrationsMongoToCosmosDbvCoreMongoClient() instead.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseMigrationsMongoToCosmosDbvCoreMongoClient creates a new instance of DatabaseMigrationsMongoToCosmosDbvCoreMongoClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseMigrationsMongoToCosmosDbvCoreMongoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseMigrationsMongoToCosmosDbvCoreMongoClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseMigrationsMongoToCosmosDbvCoreMongoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create or Update Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetResourceName - The name of the target resource/account.
//   - migrationName - Name of the migration.
//   - parameters - Details of CosmosDB for Mongo API Migration resource.
//   - options - DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginCreateOptions contains the optional parameters for the
//     DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginCreate method.
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) BeginCreate(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, parameters DatabaseMigrationCosmosDbMongo, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginCreateOptions) (*runtime.Poller[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, targetResourceName, migrationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create or Update Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) create(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, parameters DatabaseMigrationCosmosDbMongo, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, targetResourceName, migrationName, parameters, options)
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
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) createCreateRequest(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, parameters DatabaseMigrationCosmosDbMongo, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{targetResourceName}/providers/Microsoft.DataMigration/databaseMigrations/{migrationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if targetResourceName == "" {
		return nil, errors.New("parameter targetResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetResourceName}", url.PathEscape(targetResourceName))
	if migrationName == "" {
		return nil, errors.New("parameter migrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationName}", url.PathEscape(migrationName))
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

// BeginDelete - Delete Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetResourceName - The name of the target resource/account.
//   - migrationName - Name of the migration.
//   - options - DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginDeleteOptions contains the optional parameters for the
//     DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginDelete method.
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) BeginDelete(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginDeleteOptions) (*runtime.Poller[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, targetResourceName, migrationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) deleteOperation(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, targetResourceName, migrationName, options)
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
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{targetResourceName}/providers/Microsoft.DataMigration/databaseMigrations/{migrationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if targetResourceName == "" {
		return nil, errors.New("parameter targetResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetResourceName}", url.PathEscape(targetResourceName))
	if migrationName == "" {
		return nil, errors.New("parameter migrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationName}", url.PathEscape(migrationName))
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
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetResourceName - The name of the target resource/account.
//   - migrationName - Name of the migration.
//   - options - DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.Get
//     method.
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) Get(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetOptions) (DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse, error) {
	var err error
	const operationName = "DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, targetResourceName, migrationName, options)
	if err != nil {
		return DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) getCreateRequest(ctx context.Context, resourceGroupName string, targetResourceName string, migrationName string, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{targetResourceName}/providers/Microsoft.DataMigration/databaseMigrations/{migrationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if targetResourceName == "" {
		return nil, errors.New("parameter targetResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetResourceName}", url.PathEscape(targetResourceName))
	if migrationName == "" {
		return nil, errors.New("parameter migrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationName}", url.PathEscape(migrationName))
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
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) getHandleResponse(resp *http.Response) (DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse, error) {
	result := DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseMigrationCosmosDbMongo); err != nil {
		return DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse{}, err
	}
	return result, nil
}

// NewGetForScopePager - Get Database Migration resources for the scope.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetResourceName - The name of the target resource/account.
//   - options - DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeOptions contains the optional parameters for the
//     DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.NewGetForScopePager method.
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) NewGetForScopePager(resourceGroupName string, targetResourceName string, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeOptions) *runtime.Pager[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse]{
		More: func(page DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse) (DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.NewGetForScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getForScopeCreateRequest(ctx, resourceGroupName, targetResourceName, options)
			}, nil)
			if err != nil {
				return DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse{}, err
			}
			return client.getForScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// getForScopeCreateRequest creates the GetForScope request.
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) getForScopeCreateRequest(ctx context.Context, resourceGroupName string, targetResourceName string, options *DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{targetResourceName}/providers/Microsoft.DataMigration/databaseMigrations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if targetResourceName == "" {
		return nil, errors.New("parameter targetResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetResourceName}", url.PathEscape(targetResourceName))
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

// getForScopeHandleResponse handles the GetForScope response.
func (client *DatabaseMigrationsMongoToCosmosDbvCoreMongoClient) getForScopeHandleResponse(resp *http.Response) (DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse, error) {
	result := DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseMigrationCosmosDbMongoListResult); err != nil {
		return DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse{}, err
	}
	return result, nil
}
