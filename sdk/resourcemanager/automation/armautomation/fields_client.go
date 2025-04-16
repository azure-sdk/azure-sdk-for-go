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

// FieldsClient contains the methods for the Fields group.
// Don't use this type directly, use NewFieldsClient() instead.
type FieldsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFieldsClient creates a new instance of FieldsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFieldsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FieldsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FieldsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByTypePager - Retrieve a list of fields of a given type identified by module name.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - moduleName - The name of module.
//   - typeName - The name of type.
//   - options - FieldsClientListByTypeOptions contains the optional parameters for the FieldsClient.NewListByTypePager method.
func (client *FieldsClient) NewListByTypePager(resourceGroupName string, automationAccountName string, moduleName string, typeName string, options *FieldsClientListByTypeOptions) *runtime.Pager[FieldsClientListByTypeResponse] {
	return runtime.NewPager(runtime.PagingHandler[FieldsClientListByTypeResponse]{
		More: func(page FieldsClientListByTypeResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *FieldsClientListByTypeResponse) (FieldsClientListByTypeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FieldsClient.NewListByTypePager")
			req, err := client.listByTypeCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, typeName, options)
			if err != nil {
				return FieldsClientListByTypeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FieldsClientListByTypeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FieldsClientListByTypeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTypeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTypeCreateRequest creates the ListByType request.
func (client *FieldsClient) listByTypeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string, _ *FieldsClientListByTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/types/{typeName}/fields"
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
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTypeHandleResponse handles the ListByType response.
func (client *FieldsClient) listByTypeHandleResponse(resp *http.Response) (FieldsClientListByTypeResponse, error) {
	result := FieldsClientListByTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TypeFieldListResult); err != nil {
		return FieldsClientListByTypeResponse{}, err
	}
	return result, nil
}
