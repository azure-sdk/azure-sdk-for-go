//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// ServerTrustCertificatesClient contains the methods for the ServerTrustCertificates group.
// Don't use this type directly, use NewServerTrustCertificatesClient() instead.
type ServerTrustCertificatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerTrustCertificatesClient creates a new instance of ServerTrustCertificatesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerTrustCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerTrustCertificatesClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerTrustCertificatesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerTrustCertificatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Uploads a server trust certificate from box to Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - certificateName - Name of of the certificate to upload.
//   - parameters - The server trust certificate info.
//   - options - ServerTrustCertificatesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerTrustCertificatesClient.BeginCreateOrUpdate
//     method.
func (client *ServerTrustCertificatesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, parameters ServerTrustCertificate, options *ServerTrustCertificatesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerTrustCertificatesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, certificateName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerTrustCertificatesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerTrustCertificatesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Uploads a server trust certificate from box to Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *ServerTrustCertificatesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, parameters ServerTrustCertificate, options *ServerTrustCertificatesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, certificateName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerTrustCertificatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, parameters ServerTrustCertificate, options *ServerTrustCertificatesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverTrustCertificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a server trust certificate that was uploaded from box to Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - certificateName - Name of of the certificate to delete.
//   - options - ServerTrustCertificatesClientBeginDeleteOptions contains the optional parameters for the ServerTrustCertificatesClient.BeginDelete
//     method.
func (client *ServerTrustCertificatesClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, options *ServerTrustCertificatesClientBeginDeleteOptions) (*runtime.Poller[ServerTrustCertificatesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, certificateName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerTrustCertificatesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerTrustCertificatesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a server trust certificate that was uploaded from box to Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *ServerTrustCertificatesClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, options *ServerTrustCertificatesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, certificateName, options)
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
func (client *ServerTrustCertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, options *ServerTrustCertificatesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverTrustCertificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a server trust certificate that was uploaded from box to Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - certificateName - Name of of the certificate to get.
//   - options - ServerTrustCertificatesClientGetOptions contains the optional parameters for the ServerTrustCertificatesClient.Get
//     method.
func (client *ServerTrustCertificatesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, options *ServerTrustCertificatesClientGetOptions) (ServerTrustCertificatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, certificateName, options)
	if err != nil {
		return ServerTrustCertificatesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerTrustCertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerTrustCertificatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerTrustCertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, options *ServerTrustCertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverTrustCertificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerTrustCertificatesClient) getHandleResponse(resp *http.Response) (ServerTrustCertificatesClientGetResponse, error) {
	result := ServerTrustCertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerTrustCertificate); err != nil {
		return ServerTrustCertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Gets a list of server trust certificates that were uploaded from box to the given Sql Managed
// Instance.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ServerTrustCertificatesClientListByInstanceOptions contains the optional parameters for the ServerTrustCertificatesClient.NewListByInstancePager
//     method.
func (client *ServerTrustCertificatesClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *ServerTrustCertificatesClientListByInstanceOptions) *runtime.Pager[ServerTrustCertificatesClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerTrustCertificatesClientListByInstanceResponse]{
		More: func(page ServerTrustCertificatesClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerTrustCertificatesClientListByInstanceResponse) (ServerTrustCertificatesClientListByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerTrustCertificatesClientListByInstanceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerTrustCertificatesClientListByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerTrustCertificatesClientListByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstanceHandleResponse(resp)
		},
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ServerTrustCertificatesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ServerTrustCertificatesClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverTrustCertificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ServerTrustCertificatesClient) listByInstanceHandleResponse(resp *http.Response) (ServerTrustCertificatesClientListByInstanceResponse, error) {
	result := ServerTrustCertificatesClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerTrustCertificatesListResult); err != nil {
		return ServerTrustCertificatesClientListByInstanceResponse{}, err
	}
	return result, nil
}
