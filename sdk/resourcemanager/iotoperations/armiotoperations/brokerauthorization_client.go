// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armiotoperations

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

// BrokerAuthorizationClient contains the methods for the BrokerAuthorization group.
// Don't use this type directly, use NewBrokerAuthorizationClient() instead.
type BrokerAuthorizationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBrokerAuthorizationClient creates a new instance of BrokerAuthorizationClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBrokerAuthorizationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BrokerAuthorizationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BrokerAuthorizationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a BrokerAuthorizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - brokerName - Name of broker.
//   - authorizationName - Name of Instance broker authorization resource
//   - resource - Resource create parameters.
//   - options - BrokerAuthorizationClientBeginCreateOrUpdateOptions contains the optional parameters for the BrokerAuthorizationClient.BeginCreateOrUpdate
//     method.
func (client *BrokerAuthorizationClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, resource BrokerAuthorizationResource, options *BrokerAuthorizationClientBeginCreateOrUpdateOptions) (*runtime.Poller[BrokerAuthorizationClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, instanceName, brokerName, authorizationName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BrokerAuthorizationClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BrokerAuthorizationClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a BrokerAuthorizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *BrokerAuthorizationClient) createOrUpdate(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, resource BrokerAuthorizationResource, options *BrokerAuthorizationClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BrokerAuthorizationClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, instanceName, brokerName, authorizationName, resource, options)
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
func (client *BrokerAuthorizationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, resource BrokerAuthorizationResource, _ *BrokerAuthorizationClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/brokers/{brokerName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if brokerName == "" {
		return nil, errors.New("parameter brokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{brokerName}", url.PathEscape(brokerName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a BrokerAuthorizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - brokerName - Name of broker.
//   - authorizationName - Name of Instance broker authorization resource
//   - options - BrokerAuthorizationClientBeginDeleteOptions contains the optional parameters for the BrokerAuthorizationClient.BeginDelete
//     method.
func (client *BrokerAuthorizationClient) BeginDelete(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, options *BrokerAuthorizationClientBeginDeleteOptions) (*runtime.Poller[BrokerAuthorizationClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, instanceName, brokerName, authorizationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BrokerAuthorizationClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BrokerAuthorizationClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a BrokerAuthorizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *BrokerAuthorizationClient) deleteOperation(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, options *BrokerAuthorizationClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "BrokerAuthorizationClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, instanceName, brokerName, authorizationName, options)
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
func (client *BrokerAuthorizationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, _ *BrokerAuthorizationClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/brokers/{brokerName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if brokerName == "" {
		return nil, errors.New("parameter brokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{brokerName}", url.PathEscape(brokerName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a BrokerAuthorizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - brokerName - Name of broker.
//   - authorizationName - Name of Instance broker authorization resource
//   - options - BrokerAuthorizationClientGetOptions contains the optional parameters for the BrokerAuthorizationClient.Get method.
func (client *BrokerAuthorizationClient) Get(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, options *BrokerAuthorizationClientGetOptions) (BrokerAuthorizationClientGetResponse, error) {
	var err error
	const operationName = "BrokerAuthorizationClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, instanceName, brokerName, authorizationName, options)
	if err != nil {
		return BrokerAuthorizationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BrokerAuthorizationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BrokerAuthorizationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BrokerAuthorizationClient) getCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, authorizationName string, _ *BrokerAuthorizationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/brokers/{brokerName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if brokerName == "" {
		return nil, errors.New("parameter brokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{brokerName}", url.PathEscape(brokerName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BrokerAuthorizationClient) getHandleResponse(resp *http.Response) (BrokerAuthorizationClientGetResponse, error) {
	result := BrokerAuthorizationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BrokerAuthorizationResource); err != nil {
		return BrokerAuthorizationClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List BrokerAuthorizationResource resources by BrokerResource
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - brokerName - Name of broker.
//   - options - BrokerAuthorizationClientListByResourceGroupOptions contains the optional parameters for the BrokerAuthorizationClient.NewListByResourceGroupPager
//     method.
func (client *BrokerAuthorizationClient) NewListByResourceGroupPager(resourceGroupName string, instanceName string, brokerName string, options *BrokerAuthorizationClientListByResourceGroupOptions) *runtime.Pager[BrokerAuthorizationClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[BrokerAuthorizationClientListByResourceGroupResponse]{
		More: func(page BrokerAuthorizationClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BrokerAuthorizationClientListByResourceGroupResponse) (BrokerAuthorizationClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BrokerAuthorizationClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, instanceName, brokerName, options)
			}, nil)
			if err != nil {
				return BrokerAuthorizationClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *BrokerAuthorizationClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, brokerName string, _ *BrokerAuthorizationClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/brokers/{brokerName}/authorizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if brokerName == "" {
		return nil, errors.New("parameter brokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{brokerName}", url.PathEscape(brokerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *BrokerAuthorizationClient) listByResourceGroupHandleResponse(resp *http.Response) (BrokerAuthorizationClientListByResourceGroupResponse, error) {
	result := BrokerAuthorizationClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BrokerAuthorizationResourceListResult); err != nil {
		return BrokerAuthorizationClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
