//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory

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

// DataFlowDebugSessionClient contains the methods for the DataFlowDebugSession group.
// Don't use this type directly, use NewDataFlowDebugSessionClient() instead.
type DataFlowDebugSessionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataFlowDebugSessionClient creates a new instance of DataFlowDebugSessionClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataFlowDebugSessionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataFlowDebugSessionClient, error) {
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
	client := &DataFlowDebugSessionClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// AddDataFlow - Add a data flow into debug session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - request - Data flow debug session definition with debug content.
//   - options - DataFlowDebugSessionClientAddDataFlowOptions contains the optional parameters for the DataFlowDebugSessionClient.AddDataFlow
//     method.
func (client *DataFlowDebugSessionClient) AddDataFlow(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugPackage, options *DataFlowDebugSessionClientAddDataFlowOptions) (DataFlowDebugSessionClientAddDataFlowResponse, error) {
	req, err := client.addDataFlowCreateRequest(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, runtime.NewResponseError(resp)
	}
	return client.addDataFlowHandleResponse(resp)
}

// addDataFlowCreateRequest creates the AddDataFlow request.
func (client *DataFlowDebugSessionClient) addDataFlowCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugPackage, options *DataFlowDebugSessionClientAddDataFlowOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/addDataFlowToDebugSession"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// addDataFlowHandleResponse handles the AddDataFlow response.
func (client *DataFlowDebugSessionClient) addDataFlowHandleResponse(resp *http.Response) (DataFlowDebugSessionClientAddDataFlowResponse, error) {
	result := DataFlowDebugSessionClientAddDataFlowResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AddDataFlowToDebugSessionResponse); err != nil {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, err
	}
	return result, nil
}

// BeginCreate - Creates a data flow debug session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - request - Data flow debug session definition
//   - options - DataFlowDebugSessionClientBeginCreateOptions contains the optional parameters for the DataFlowDebugSessionClient.BeginCreate
//     method.
func (client *DataFlowDebugSessionClient) BeginCreate(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientBeginCreateOptions) (*runtime.Poller[DataFlowDebugSessionClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, factoryName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DataFlowDebugSessionClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowDebugSessionClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a data flow debug session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *DataFlowDebugSessionClient) create(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, factoryName, request, options)
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

// createCreateRequest creates the Create request.
func (client *DataFlowDebugSessionClient) createCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/createDataFlowDebugSession"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// Delete - Deletes a data flow debug session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - request - Data flow debug session definition for deletion
//   - options - DataFlowDebugSessionClientDeleteOptions contains the optional parameters for the DataFlowDebugSessionClient.Delete
//     method.
func (client *DataFlowDebugSessionClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientDeleteOptions) (DataFlowDebugSessionClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return DataFlowDebugSessionClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataFlowDebugSessionClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataFlowDebugSessionClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataFlowDebugSessionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/deleteDataFlowDebugSession"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// BeginExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - request - Data flow debug command definition.
//   - options - DataFlowDebugSessionClientBeginExecuteCommandOptions contains the optional parameters for the DataFlowDebugSessionClient.BeginExecuteCommand
//     method.
func (client *DataFlowDebugSessionClient) BeginExecuteCommand(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionClientBeginExecuteCommandOptions) (*runtime.Poller[DataFlowDebugSessionClientExecuteCommandResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.executeCommand(ctx, resourceGroupName, factoryName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DataFlowDebugSessionClientExecuteCommandResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowDebugSessionClientExecuteCommandResponse](options.ResumeToken, client.pl, nil)
	}
}

// ExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *DataFlowDebugSessionClient) executeCommand(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionClientBeginExecuteCommandOptions) (*http.Response, error) {
	req, err := client.executeCommandCreateRequest(ctx, resourceGroupName, factoryName, request, options)
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

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *DataFlowDebugSessionClient) executeCommandCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionClientBeginExecuteCommandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/executeDataFlowDebugCommand"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// NewQueryByFactoryPager - Query all active data flow debug sessions.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - options - DataFlowDebugSessionClientQueryByFactoryOptions contains the optional parameters for the DataFlowDebugSessionClient.NewQueryByFactoryPager
//     method.
func (client *DataFlowDebugSessionClient) NewQueryByFactoryPager(resourceGroupName string, factoryName string, options *DataFlowDebugSessionClientQueryByFactoryOptions) *runtime.Pager[DataFlowDebugSessionClientQueryByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataFlowDebugSessionClientQueryByFactoryResponse]{
		More: func(page DataFlowDebugSessionClientQueryByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataFlowDebugSessionClientQueryByFactoryResponse) (DataFlowDebugSessionClientQueryByFactoryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.queryByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataFlowDebugSessionClientQueryByFactoryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataFlowDebugSessionClientQueryByFactoryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataFlowDebugSessionClientQueryByFactoryResponse{}, runtime.NewResponseError(resp)
			}
			return client.queryByFactoryHandleResponse(resp)
		},
	})
}

// queryByFactoryCreateRequest creates the QueryByFactory request.
func (client *DataFlowDebugSessionClient) queryByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *DataFlowDebugSessionClientQueryByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryDataFlowDebugSessions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// queryByFactoryHandleResponse handles the QueryByFactory response.
func (client *DataFlowDebugSessionClient) queryByFactoryHandleResponse(resp *http.Response) (DataFlowDebugSessionClientQueryByFactoryResponse, error) {
	result := DataFlowDebugSessionClientQueryByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryDataFlowDebugSessionsResponse); err != nil {
		return DataFlowDebugSessionClientQueryByFactoryResponse{}, err
	}
	return result, nil
}
