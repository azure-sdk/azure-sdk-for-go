//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// ObjectDataTypesClient contains the methods for the ObjectDataTypes group.
// Don't use this type directly, use NewObjectDataTypesClient() instead.
type ObjectDataTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewObjectDataTypesClient creates a new instance of ObjectDataTypesClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewObjectDataTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ObjectDataTypesClient, error) {
	cl, err := arm.NewClient(moduleName+".ObjectDataTypesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ObjectDataTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListFieldsByModuleAndTypePager - Retrieve a list of fields of a given type identified by module name.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - moduleName - The name of module.
//   - typeName - The name of type.
//   - options - ObjectDataTypesClientListFieldsByModuleAndTypeOptions contains the optional parameters for the ObjectDataTypesClient.NewListFieldsByModuleAndTypePager
//     method.
func (client *ObjectDataTypesClient) NewListFieldsByModuleAndTypePager(resourceGroupName string, automationAccountName string, moduleName string, typeName string, options *ObjectDataTypesClientListFieldsByModuleAndTypeOptions) *runtime.Pager[ObjectDataTypesClientListFieldsByModuleAndTypeResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObjectDataTypesClientListFieldsByModuleAndTypeResponse]{
		More: func(page ObjectDataTypesClientListFieldsByModuleAndTypeResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ObjectDataTypesClientListFieldsByModuleAndTypeResponse) (ObjectDataTypesClientListFieldsByModuleAndTypeResponse, error) {
			req, err := client.listFieldsByModuleAndTypeCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, typeName, options)
			if err != nil {
				return ObjectDataTypesClientListFieldsByModuleAndTypeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ObjectDataTypesClientListFieldsByModuleAndTypeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ObjectDataTypesClientListFieldsByModuleAndTypeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listFieldsByModuleAndTypeHandleResponse(resp)
		},
	})
}

// listFieldsByModuleAndTypeCreateRequest creates the ListFieldsByModuleAndType request.
func (client *ObjectDataTypesClient) listFieldsByModuleAndTypeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string, options *ObjectDataTypesClientListFieldsByModuleAndTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/objectDataTypes/{typeName}/fields"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if typeName == "" {
		return nil, errors.New("parameter typeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{typeName}", url.PathEscape(typeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listFieldsByModuleAndTypeHandleResponse handles the ListFieldsByModuleAndType response.
func (client *ObjectDataTypesClient) listFieldsByModuleAndTypeHandleResponse(resp *http.Response) (ObjectDataTypesClientListFieldsByModuleAndTypeResponse, error) {
	result := ObjectDataTypesClientListFieldsByModuleAndTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TypeFieldListResult); err != nil {
		return ObjectDataTypesClientListFieldsByModuleAndTypeResponse{}, err
	}
	return result, nil
}

// NewListFieldsByTypePager - Retrieve a list of fields of a given type across all accessible modules.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - typeName - The name of type.
//   - options - ObjectDataTypesClientListFieldsByTypeOptions contains the optional parameters for the ObjectDataTypesClient.NewListFieldsByTypePager
//     method.
func (client *ObjectDataTypesClient) NewListFieldsByTypePager(resourceGroupName string, automationAccountName string, typeName string, options *ObjectDataTypesClientListFieldsByTypeOptions) *runtime.Pager[ObjectDataTypesClientListFieldsByTypeResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObjectDataTypesClientListFieldsByTypeResponse]{
		More: func(page ObjectDataTypesClientListFieldsByTypeResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ObjectDataTypesClientListFieldsByTypeResponse) (ObjectDataTypesClientListFieldsByTypeResponse, error) {
			req, err := client.listFieldsByTypeCreateRequest(ctx, resourceGroupName, automationAccountName, typeName, options)
			if err != nil {
				return ObjectDataTypesClientListFieldsByTypeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ObjectDataTypesClientListFieldsByTypeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ObjectDataTypesClientListFieldsByTypeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listFieldsByTypeHandleResponse(resp)
		},
	})
}

// listFieldsByTypeCreateRequest creates the ListFieldsByType request.
func (client *ObjectDataTypesClient) listFieldsByTypeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, typeName string, options *ObjectDataTypesClientListFieldsByTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/objectDataTypes/{typeName}/fields"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if typeName == "" {
		return nil, errors.New("parameter typeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{typeName}", url.PathEscape(typeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listFieldsByTypeHandleResponse handles the ListFieldsByType response.
func (client *ObjectDataTypesClient) listFieldsByTypeHandleResponse(resp *http.Response) (ObjectDataTypesClientListFieldsByTypeResponse, error) {
	result := ObjectDataTypesClientListFieldsByTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TypeFieldListResult); err != nil {
		return ObjectDataTypesClientListFieldsByTypeResponse{}, err
	}
	return result, nil
}
