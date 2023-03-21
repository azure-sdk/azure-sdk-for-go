//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcommunication

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

// ServicesClient contains the methods for the CommunicationServices group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
	cl, err := arm.NewClient(moduleName+".ServicesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks that the CommunicationService name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - nameAvailabilityParameters - Parameters supplied to the operation.
//   - options - ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
//     method.
func (client *ServicesClient) CheckNameAvailability(ctx context.Context, nameAvailabilityParameters NameAvailabilityParameters, options *ServicesClientCheckNameAvailabilityOptions) (ServicesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, nameAvailabilityParameters, options)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ServicesClient) checkNameAvailabilityCreateRequest(ctx context.Context, nameAvailabilityParameters NameAvailabilityParameters, options *ServicesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Communication/checkNameAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, nameAvailabilityParameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ServicesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ServicesClientCheckNameAvailabilityResponse, error) {
	result := ServicesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create a new CommunicationService or update an existing CommunicationService.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationServiceName - The name of the CommunicationService resource.
//   - parameters - Parameters for the create or update operation
//   - options - ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate
//     method.
func (client *ServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, communicationServiceName string, parameters ServiceResource, options *ServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, communicationServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServicesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new CommunicationService or update an existing CommunicationService.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
func (client *ServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, communicationServiceName string, parameters ServiceResource, options *ServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, communicationServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, communicationServiceName string, parameters ServiceResource, options *ServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices/{communicationServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationServiceName == "" {
		return nil, errors.New("parameter communicationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationServiceName}", url.PathEscape(communicationServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Operation to delete a CommunicationService.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationServiceName - The name of the CommunicationService resource.
//   - options - ServicesClientBeginDeleteOptions contains the optional parameters for the ServicesClient.BeginDelete method.
func (client *ServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientBeginDeleteOptions) (*runtime.Poller[ServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, communicationServiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServicesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Operation to delete a CommunicationService.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
func (client *ServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, communicationServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices/{communicationServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationServiceName == "" {
		return nil, errors.New("parameter communicationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationServiceName}", url.PathEscape(communicationServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the CommunicationService and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationServiceName - The name of the CommunicationService resource.
//   - options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, communicationServiceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices/{communicationServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationServiceName == "" {
		return nil, errors.New("parameter communicationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationServiceName}", url.PathEscape(communicationServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// LinkNotificationHub - Links an Azure Notification Hub to this communication service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationServiceName - The name of the CommunicationService resource.
//   - options - ServicesClientLinkNotificationHubOptions contains the optional parameters for the ServicesClient.LinkNotificationHub
//     method.
func (client *ServicesClient) LinkNotificationHub(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientLinkNotificationHubOptions) (ServicesClientLinkNotificationHubResponse, error) {
	req, err := client.linkNotificationHubCreateRequest(ctx, resourceGroupName, communicationServiceName, options)
	if err != nil {
		return ServicesClientLinkNotificationHubResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientLinkNotificationHubResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientLinkNotificationHubResponse{}, runtime.NewResponseError(resp)
	}
	return client.linkNotificationHubHandleResponse(resp)
}

// linkNotificationHubCreateRequest creates the LinkNotificationHub request.
func (client *ServicesClient) linkNotificationHubCreateRequest(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientLinkNotificationHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices/{communicationServiceName}/linkNotificationHub"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationServiceName == "" {
		return nil, errors.New("parameter communicationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationServiceName}", url.PathEscape(communicationServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.LinkNotificationHubParameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.LinkNotificationHubParameters)
	}
	return req, nil
}

// linkNotificationHubHandleResponse handles the LinkNotificationHub response.
func (client *ServicesClient) linkNotificationHubHandleResponse(resp *http.Response) (ServicesClientLinkNotificationHubResponse, error) {
	result := ServicesClientLinkNotificationHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedNotificationHub); err != nil {
		return ServicesClientLinkNotificationHubResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Handles requests to list all resources in a resource group.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.NewListByResourceGroupPager
//     method.
func (client *ServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *ServicesClientListByResourceGroupOptions) *runtime.Pager[ServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListByResourceGroupResponse]{
		More: func(page ServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListByResourceGroupResponse) (ServicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ServicesClientListByResourceGroupResponse, error) {
	result := ServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceList); err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Handles requests to list all resources in a subscription.
//
// Generated from API version 2023-03-31
//   - options - ServicesClientListBySubscriptionOptions contains the optional parameters for the ServicesClient.NewListBySubscriptionPager
//     method.
func (client *ServicesClient) NewListBySubscriptionPager(options *ServicesClientListBySubscriptionOptions) *runtime.Pager[ServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListBySubscriptionResponse]{
		More: func(page ServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListBySubscriptionResponse) (ServicesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Communication/communicationServices"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (ServicesClientListBySubscriptionResponse, error) {
	result := ServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceList); err != nil {
		return ServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListKeys - Get the access keys of the CommunicationService resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationServiceName - The name of the CommunicationService resource.
//   - options - ServicesClientListKeysOptions contains the optional parameters for the ServicesClient.ListKeys method.
func (client *ServicesClient) ListKeys(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientListKeysOptions) (ServicesClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, communicationServiceName, options)
	if err != nil {
		return ServicesClientListKeysResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ServicesClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, communicationServiceName string, options *ServicesClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices/{communicationServiceName}/listKeys"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationServiceName == "" {
		return nil, errors.New("parameter communicationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationServiceName}", url.PathEscape(communicationServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ServicesClient) listKeysHandleResponse(resp *http.Response) (ServicesClientListKeysResponse, error) {
	result := ServicesClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceKeys); err != nil {
		return ServicesClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKey - Regenerate CommunicationService access key. PrimaryKey and SecondaryKey cannot be regenerated at the same
// time.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationServiceName - The name of the CommunicationService resource.
//   - parameters - Parameter that describes the Regenerate Key Operation.
//   - options - ServicesClientRegenerateKeyOptions contains the optional parameters for the ServicesClient.RegenerateKey method.
func (client *ServicesClient) RegenerateKey(ctx context.Context, resourceGroupName string, communicationServiceName string, parameters RegenerateKeyParameters, options *ServicesClientRegenerateKeyOptions) (ServicesClientRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, communicationServiceName, parameters, options)
	if err != nil {
		return ServicesClientRegenerateKeyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientRegenerateKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *ServicesClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, communicationServiceName string, parameters RegenerateKeyParameters, options *ServicesClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices/{communicationServiceName}/regenerateKey"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationServiceName == "" {
		return nil, errors.New("parameter communicationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationServiceName}", url.PathEscape(communicationServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *ServicesClient) regenerateKeyHandleResponse(resp *http.Response) (ServicesClientRegenerateKeyResponse, error) {
	result := ServicesClientRegenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceKeys); err != nil {
		return ServicesClientRegenerateKeyResponse{}, err
	}
	return result, nil
}

// Update - Operation to update an existing CommunicationService.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-31
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - communicationServiceName - The name of the CommunicationService resource.
//   - parameters - Parameters for the update operation
//   - options - ServicesClientUpdateOptions contains the optional parameters for the ServicesClient.Update method.
func (client *ServicesClient) Update(ctx context.Context, resourceGroupName string, communicationServiceName string, parameters ServiceResourceUpdate, options *ServicesClientUpdateOptions) (ServicesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, communicationServiceName, parameters, options)
	if err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, communicationServiceName string, parameters ServiceResourceUpdate, options *ServicesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/communicationServices/{communicationServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if communicationServiceName == "" {
		return nil, errors.New("parameter communicationServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationServiceName}", url.PathEscape(communicationServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ServicesClient) updateHandleResponse(resp *http.Response) (ServicesClientUpdateResponse, error) {
	result := ServicesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	return result, nil
}
