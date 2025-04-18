// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package arminformaticadatamgmt

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

// ServerlessRuntimesClient contains the methods for the ServerlessRuntimes group.
// Don't use this type directly, use NewServerlessRuntimesClient() instead.
type ServerlessRuntimesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerlessRuntimesClient creates a new instance of ServerlessRuntimesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerlessRuntimesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerlessRuntimesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerlessRuntimesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckDependencies - Checks all dependencies for a serverless runtime resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - serverlessRuntimeName - Name of the Serverless Runtime resource
//   - options - ServerlessRuntimesClientCheckDependenciesOptions contains the optional parameters for the ServerlessRuntimesClient.CheckDependencies
//     method.
func (client *ServerlessRuntimesClient) CheckDependencies(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, options *ServerlessRuntimesClientCheckDependenciesOptions) (ServerlessRuntimesClientCheckDependenciesResponse, error) {
	var err error
	const operationName = "ServerlessRuntimesClient.CheckDependencies"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkDependenciesCreateRequest(ctx, resourceGroupName, organizationName, serverlessRuntimeName, options)
	if err != nil {
		return ServerlessRuntimesClientCheckDependenciesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerlessRuntimesClientCheckDependenciesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerlessRuntimesClientCheckDependenciesResponse{}, err
	}
	resp, err := client.checkDependenciesHandleResponse(httpResp)
	return resp, err
}

// checkDependenciesCreateRequest creates the CheckDependencies request.
func (client *ServerlessRuntimesClient) checkDependenciesCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, _ *ServerlessRuntimesClientCheckDependenciesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes/{serverlessRuntimeName}/checkDependencies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if serverlessRuntimeName == "" {
		return nil, errors.New("parameter serverlessRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverlessRuntimeName}", url.PathEscape(serverlessRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// checkDependenciesHandleResponse handles the CheckDependencies response.
func (client *ServerlessRuntimesClient) checkDependenciesHandleResponse(resp *http.Response) (ServerlessRuntimesClientCheckDependenciesResponse, error) {
	result := ServerlessRuntimesClientCheckDependenciesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckDependenciesResponse); err != nil {
		return ServerlessRuntimesClientCheckDependenciesResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create a InformaticaServerlessRuntimeResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - serverlessRuntimeName - Name of the Serverless Runtime resource
//   - resource - Resource create parameters.
//   - options - ServerlessRuntimesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerlessRuntimesClient.BeginCreateOrUpdate
//     method.
func (client *ServerlessRuntimesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, resource InformaticaServerlessRuntimeResource, options *ServerlessRuntimesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerlessRuntimesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, organizationName, serverlessRuntimeName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerlessRuntimesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerlessRuntimesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a InformaticaServerlessRuntimeResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
func (client *ServerlessRuntimesClient) createOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, resource InformaticaServerlessRuntimeResource, options *ServerlessRuntimesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerlessRuntimesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, organizationName, serverlessRuntimeName, resource, options)
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
func (client *ServerlessRuntimesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, resource InformaticaServerlessRuntimeResource, _ *ServerlessRuntimesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes/{serverlessRuntimeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if serverlessRuntimeName == "" {
		return nil, errors.New("parameter serverlessRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverlessRuntimeName}", url.PathEscape(serverlessRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a InformaticaServerlessRuntimeResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - serverlessRuntimeName - Name of the Serverless Runtime resource
//   - options - ServerlessRuntimesClientBeginDeleteOptions contains the optional parameters for the ServerlessRuntimesClient.BeginDelete
//     method.
func (client *ServerlessRuntimesClient) BeginDelete(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, options *ServerlessRuntimesClientBeginDeleteOptions) (*runtime.Poller[ServerlessRuntimesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, organizationName, serverlessRuntimeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerlessRuntimesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerlessRuntimesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a InformaticaServerlessRuntimeResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
func (client *ServerlessRuntimesClient) deleteOperation(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, options *ServerlessRuntimesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerlessRuntimesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, serverlessRuntimeName, options)
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
func (client *ServerlessRuntimesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, _ *ServerlessRuntimesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes/{serverlessRuntimeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if serverlessRuntimeName == "" {
		return nil, errors.New("parameter serverlessRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverlessRuntimeName}", url.PathEscape(serverlessRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a InformaticaServerlessRuntimeResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - serverlessRuntimeName - Name of the Serverless Runtime resource
//   - options - ServerlessRuntimesClientGetOptions contains the optional parameters for the ServerlessRuntimesClient.Get method.
func (client *ServerlessRuntimesClient) Get(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, options *ServerlessRuntimesClientGetOptions) (ServerlessRuntimesClientGetResponse, error) {
	var err error
	const operationName = "ServerlessRuntimesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, organizationName, serverlessRuntimeName, options)
	if err != nil {
		return ServerlessRuntimesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerlessRuntimesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerlessRuntimesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerlessRuntimesClient) getCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, _ *ServerlessRuntimesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes/{serverlessRuntimeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if serverlessRuntimeName == "" {
		return nil, errors.New("parameter serverlessRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverlessRuntimeName}", url.PathEscape(serverlessRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerlessRuntimesClient) getHandleResponse(resp *http.Response) (ServerlessRuntimesClientGetResponse, error) {
	result := ServerlessRuntimesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaServerlessRuntimeResource); err != nil {
		return ServerlessRuntimesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInformaticaOrganizationResourcePager - List InformaticaServerlessRuntimeResource resources by InformaticaOrganizationResource
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - options - ServerlessRuntimesClientListByInformaticaOrganizationResourceOptions contains the optional parameters for the
//     ServerlessRuntimesClient.NewListByInformaticaOrganizationResourcePager method.
func (client *ServerlessRuntimesClient) NewListByInformaticaOrganizationResourcePager(resourceGroupName string, organizationName string, options *ServerlessRuntimesClientListByInformaticaOrganizationResourceOptions) *runtime.Pager[ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse]{
		More: func(page ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse) (ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerlessRuntimesClient.NewListByInformaticaOrganizationResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInformaticaOrganizationResourceCreateRequest(ctx, resourceGroupName, organizationName, options)
			}, nil)
			if err != nil {
				return ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse{}, err
			}
			return client.listByInformaticaOrganizationResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInformaticaOrganizationResourceCreateRequest creates the ListByInformaticaOrganizationResource request.
func (client *ServerlessRuntimesClient) listByInformaticaOrganizationResourceCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, _ *ServerlessRuntimesClientListByInformaticaOrganizationResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInformaticaOrganizationResourceHandleResponse handles the ListByInformaticaOrganizationResource response.
func (client *ServerlessRuntimesClient) listByInformaticaOrganizationResourceHandleResponse(resp *http.Response) (ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse, error) {
	result := ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaServerlessRuntimeResourceListResult); err != nil {
		return ServerlessRuntimesClientListByInformaticaOrganizationResourceResponse{}, err
	}
	return result, nil
}

// ServerlessResourceByID - Returns a serverless runtime resource by ID
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - serverlessRuntimeName - Name of the Serverless Runtime resource
//   - options - ServerlessRuntimesClientServerlessResourceByIDOptions contains the optional parameters for the ServerlessRuntimesClient.ServerlessResourceByID
//     method.
func (client *ServerlessRuntimesClient) ServerlessResourceByID(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, options *ServerlessRuntimesClientServerlessResourceByIDOptions) (ServerlessRuntimesClientServerlessResourceByIDResponse, error) {
	var err error
	const operationName = "ServerlessRuntimesClient.ServerlessResourceByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.serverlessResourceByIDCreateRequest(ctx, resourceGroupName, organizationName, serverlessRuntimeName, options)
	if err != nil {
		return ServerlessRuntimesClientServerlessResourceByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerlessRuntimesClientServerlessResourceByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerlessRuntimesClientServerlessResourceByIDResponse{}, err
	}
	resp, err := client.serverlessResourceByIDHandleResponse(httpResp)
	return resp, err
}

// serverlessResourceByIDCreateRequest creates the ServerlessResourceByID request.
func (client *ServerlessRuntimesClient) serverlessResourceByIDCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, _ *ServerlessRuntimesClientServerlessResourceByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes/{serverlessRuntimeName}/serverlessResourceById"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if serverlessRuntimeName == "" {
		return nil, errors.New("parameter serverlessRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverlessRuntimeName}", url.PathEscape(serverlessRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// serverlessResourceByIDHandleResponse handles the ServerlessResourceByID response.
func (client *ServerlessRuntimesClient) serverlessResourceByIDHandleResponse(resp *http.Response) (ServerlessRuntimesClientServerlessResourceByIDResponse, error) {
	result := ServerlessRuntimesClientServerlessResourceByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaServerlessRuntimeResource); err != nil {
		return ServerlessRuntimesClientServerlessResourceByIDResponse{}, err
	}
	return result, nil
}

// StartFailedServerlessRuntime - Starts a failed runtime resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - serverlessRuntimeName - Name of the Serverless Runtime resource
//   - options - ServerlessRuntimesClientStartFailedServerlessRuntimeOptions contains the optional parameters for the ServerlessRuntimesClient.StartFailedServerlessRuntime
//     method.
func (client *ServerlessRuntimesClient) StartFailedServerlessRuntime(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, options *ServerlessRuntimesClientStartFailedServerlessRuntimeOptions) (ServerlessRuntimesClientStartFailedServerlessRuntimeResponse, error) {
	var err error
	const operationName = "ServerlessRuntimesClient.StartFailedServerlessRuntime"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startFailedServerlessRuntimeCreateRequest(ctx, resourceGroupName, organizationName, serverlessRuntimeName, options)
	if err != nil {
		return ServerlessRuntimesClientStartFailedServerlessRuntimeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerlessRuntimesClientStartFailedServerlessRuntimeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServerlessRuntimesClientStartFailedServerlessRuntimeResponse{}, err
	}
	return ServerlessRuntimesClientStartFailedServerlessRuntimeResponse{}, nil
}

// startFailedServerlessRuntimeCreateRequest creates the StartFailedServerlessRuntime request.
func (client *ServerlessRuntimesClient) startFailedServerlessRuntimeCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, _ *ServerlessRuntimesClientStartFailedServerlessRuntimeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes/{serverlessRuntimeName}/startFailedServerlessRuntime"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if serverlessRuntimeName == "" {
		return nil, errors.New("parameter serverlessRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverlessRuntimeName}", url.PathEscape(serverlessRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Update a InformaticaServerlessRuntimeResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - serverlessRuntimeName - Name of the Serverless Runtime resource
//   - properties - The resource properties to be updated.
//   - options - ServerlessRuntimesClientUpdateOptions contains the optional parameters for the ServerlessRuntimesClient.Update
//     method.
func (client *ServerlessRuntimesClient) Update(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, properties InformaticaServerlessRuntimeResourceUpdate, options *ServerlessRuntimesClientUpdateOptions) (ServerlessRuntimesClientUpdateResponse, error) {
	var err error
	const operationName = "ServerlessRuntimesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, organizationName, serverlessRuntimeName, properties, options)
	if err != nil {
		return ServerlessRuntimesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerlessRuntimesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerlessRuntimesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ServerlessRuntimesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, serverlessRuntimeName string, properties InformaticaServerlessRuntimeResourceUpdate, _ *ServerlessRuntimesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/serverlessRuntimes/{serverlessRuntimeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if serverlessRuntimeName == "" {
		return nil, errors.New("parameter serverlessRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverlessRuntimeName}", url.PathEscape(serverlessRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ServerlessRuntimesClient) updateHandleResponse(resp *http.Response) (ServerlessRuntimesClientUpdateResponse, error) {
	result := ServerlessRuntimesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaServerlessRuntimeResource); err != nil {
		return ServerlessRuntimesClientUpdateResponse{}, err
	}
	return result, nil
}
