//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationhub

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

// MachinesControllerClient contains the methods for the MachinesController group.
// Don't use this type directly, use NewMachinesControllerClient() instead.
type MachinesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMachinesControllerClient creates a new instance of MachinesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMachinesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MachinesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MachinesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetMachine - Gets a machine in the migrate project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - machineName - Unique name of a machine in Azure migration hub.
//   - options - MachinesControllerClientGetMachineOptions contains the optional parameters for the MachinesControllerClient.GetMachine
//     method.
func (client *MachinesControllerClient) GetMachine(ctx context.Context, resourceGroupName string, migrateProjectName string, machineName string, options *MachinesControllerClientGetMachineOptions) (MachinesControllerClientGetMachineResponse, error) {
	var err error
	const operationName = "MachinesControllerClient.GetMachine"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMachineCreateRequest(ctx, resourceGroupName, migrateProjectName, machineName, options)
	if err != nil {
		return MachinesControllerClientGetMachineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinesControllerClientGetMachineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachinesControllerClientGetMachineResponse{}, err
	}
	resp, err := client.getMachineHandleResponse(httpResp)
	return resp, err
}

// getMachineCreateRequest creates the GetMachine request.
func (client *MachinesControllerClient) getMachineCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, machineName string, options *MachinesControllerClientGetMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMachineHandleResponse handles the GetMachine response.
func (client *MachinesControllerClient) getMachineHandleResponse(resp *http.Response) (MachinesControllerClientGetMachineResponse, error) {
	result := MachinesControllerClientGetMachineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Machine); err != nil {
		return MachinesControllerClientGetMachineResponse{}, err
	}
	return result, nil
}

// NewListMachinesPager - Gets a list of machines in the migrate project.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - options - MachinesControllerClientListMachinesOptions contains the optional parameters for the MachinesControllerClient.NewListMachinesPager
//     method.
func (client *MachinesControllerClient) NewListMachinesPager(resourceGroupName string, migrateProjectName string, options *MachinesControllerClientListMachinesOptions) *runtime.Pager[MachinesControllerClientListMachinesResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachinesControllerClientListMachinesResponse]{
		More: func(page MachinesControllerClientListMachinesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachinesControllerClientListMachinesResponse) (MachinesControllerClientListMachinesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MachinesControllerClient.NewListMachinesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listMachinesCreateRequest(ctx, resourceGroupName, migrateProjectName, options)
			}, nil)
			if err != nil {
				return MachinesControllerClientListMachinesResponse{}, err
			}
			return client.listMachinesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listMachinesCreateRequest creates the ListMachines request.
func (client *MachinesControllerClient) listMachinesCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, options *MachinesControllerClientListMachinesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/machines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMachinesHandleResponse handles the ListMachines response.
func (client *MachinesControllerClient) listMachinesHandleResponse(resp *http.Response) (MachinesControllerClientListMachinesResponse, error) {
	result := MachinesControllerClientListMachinesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineCollection); err != nil {
		return MachinesControllerClientListMachinesResponse{}, err
	}
	return result, nil
}
