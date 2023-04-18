//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ActivityClient contains the methods for the Activity group.
// Don't use this type directly, use NewActivityClient() instead.
type ActivityClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewActivityClient creates a new instance of ActivityClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewActivityClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ActivityClient, error) {
	cl, err := arm.NewClient(moduleName+".ActivityClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ActivityClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve the activity in the module identified by module name and activity name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - moduleName - The name of module.
//   - activityName - The name of activity.
//   - options - ActivityClientGetOptions contains the optional parameters for the ActivityClient.Get method.
func (client *ActivityClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, activityName string, options *ActivityClientGetOptions) (ActivityClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, activityName, options)
	if err != nil {
		return ActivityClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ActivityClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ActivityClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, activityName string, options *ActivityClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities/{activityName}"
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
	if activityName == "" {
		return nil, errors.New("parameter activityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityName}", url.PathEscape(activityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActivityClient) getHandleResponse(resp *http.Response) (ActivityClientGetResponse, error) {
	result := ActivityClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Activity); err != nil {
		return ActivityClientGetResponse{}, err
	}
	return result, nil
}

// NewListByModulePager - Retrieve a list of activities in the module identified by module name.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - moduleName - The name of module.
//   - options - ActivityClientListByModuleOptions contains the optional parameters for the ActivityClient.NewListByModulePager
//     method.
func (client *ActivityClient) NewListByModulePager(resourceGroupName string, automationAccountName string, moduleName string, options *ActivityClientListByModuleOptions) *runtime.Pager[ActivityClientListByModuleResponse] {
	return runtime.NewPager(runtime.PagingHandler[ActivityClientListByModuleResponse]{
		More: func(page ActivityClientListByModuleResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ActivityClientListByModuleResponse) (ActivityClientListByModuleResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByModuleCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ActivityClientListByModuleResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ActivityClientListByModuleResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ActivityClientListByModuleResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByModuleHandleResponse(resp)
		},
	})
}

// listByModuleCreateRequest creates the ListByModule request.
func (client *ActivityClient) listByModuleCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, options *ActivityClientListByModuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByModuleHandleResponse handles the ListByModule response.
func (client *ActivityClient) listByModuleHandleResponse(resp *http.Response) (ActivityClientListByModuleResponse, error) {
	result := ActivityClientListByModuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityListResult); err != nil {
		return ActivityClientListByModuleResponse{}, err
	}
	return result, nil
}
