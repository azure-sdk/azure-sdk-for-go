//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

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

// ClustersClient contains the methods for the Clusters group.
// Don't use this type directly, use NewClustersClient() instead.
type ClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClustersClient creates a new instance of ClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClustersClient, error) {
	cl, err := arm.NewClient(moduleName+".ClustersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create an HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - cluster - Details of the HCI cluster.
//   - options - ClustersClientCreateOptions contains the optional parameters for the ClustersClient.Create method.
func (client *ClustersClient) Create(ctx context.Context, resourceGroupName string, clusterName string, cluster Cluster, options *ClustersClientCreateOptions) (ClustersClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, cluster, options)
	if err != nil {
		return ClustersClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClustersClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClustersClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ClustersClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, cluster Cluster, options *ClustersClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, cluster); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ClustersClient) createHandleResponse(resp *http.Response) (ClustersClientCreateResponse, error) {
	result := ClustersClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cluster); err != nil {
		return ClustersClientCreateResponse{}, err
	}
	return result, nil
}

// BeginCreateIdentity - Create cluster identity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ClustersClientBeginCreateIdentityOptions contains the optional parameters for the ClustersClient.BeginCreateIdentity
//     method.
func (client *ClustersClient) BeginCreateIdentity(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginCreateIdentityOptions) (*runtime.Poller[ClustersClientCreateIdentityResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createIdentity(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientCreateIdentityResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientCreateIdentityResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateIdentity - Create cluster identity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ClustersClient) createIdentity(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginCreateIdentityOptions) (*http.Response, error) {
	var err error
	req, err := client.createIdentityCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createIdentityCreateRequest creates the CreateIdentity request.
func (client *ClustersClient) createIdentityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginCreateIdentityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/createClusterIdentity"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Delete an HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
func (client *ClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*runtime.Poller[ClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ClustersClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete an HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExtendSoftwareAssuranceBenefit - Extends Software Assurance Benefit to a cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - softwareAssuranceChangeRequest - Software Assurance Change Request Payload
//   - options - ClustersClientBeginExtendSoftwareAssuranceBenefitOptions contains the optional parameters for the ClustersClient.BeginExtendSoftwareAssuranceBenefit
//     method.
func (client *ClustersClient) BeginExtendSoftwareAssuranceBenefit(ctx context.Context, resourceGroupName string, clusterName string, softwareAssuranceChangeRequest SoftwareAssuranceChangeRequest, options *ClustersClientBeginExtendSoftwareAssuranceBenefitOptions) (*runtime.Poller[ClustersClientExtendSoftwareAssuranceBenefitResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.extendSoftwareAssuranceBenefit(ctx, resourceGroupName, clusterName, softwareAssuranceChangeRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientExtendSoftwareAssuranceBenefitResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientExtendSoftwareAssuranceBenefitResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ExtendSoftwareAssuranceBenefit - Extends Software Assurance Benefit to a cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ClustersClient) extendSoftwareAssuranceBenefit(ctx context.Context, resourceGroupName string, clusterName string, softwareAssuranceChangeRequest SoftwareAssuranceChangeRequest, options *ClustersClientBeginExtendSoftwareAssuranceBenefitOptions) (*http.Response, error) {
	var err error
	req, err := client.extendSoftwareAssuranceBenefitCreateRequest(ctx, resourceGroupName, clusterName, softwareAssuranceChangeRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// extendSoftwareAssuranceBenefitCreateRequest creates the ExtendSoftwareAssuranceBenefit request.
func (client *ClustersClient) extendSoftwareAssuranceBenefitCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, softwareAssuranceChangeRequest SoftwareAssuranceChangeRequest, options *ClustersClientBeginExtendSoftwareAssuranceBenefitOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/extendSoftwareAssuranceBenefit"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, softwareAssuranceChangeRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
func (client *ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientGetOptions) (ClustersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClustersClient) getHandleResponse(resp *http.Response) (ClustersClientGetResponse, error) {
	result := ClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cluster); err != nil {
		return ClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all HCI clusters in a resource group.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ClustersClientListByResourceGroupOptions contains the optional parameters for the ClustersClient.NewListByResourceGroupPager
//     method.
func (client *ClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *ClustersClientListByResourceGroupOptions) *runtime.Pager[ClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListByResourceGroupResponse]{
		More: func(page ClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListByResourceGroupResponse) (ClustersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClustersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClustersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClustersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (ClustersClientListByResourceGroupResponse, error) {
	result := ClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterList); err != nil {
		return ClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all HCI clusters in a subscription.
//
// Generated from API version 2023-03-01
//   - options - ClustersClientListBySubscriptionOptions contains the optional parameters for the ClustersClient.NewListBySubscriptionPager
//     method.
func (client *ClustersClient) NewListBySubscriptionPager(options *ClustersClientListBySubscriptionOptions) *runtime.Pager[ClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListBySubscriptionResponse]{
		More: func(page ClustersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListBySubscriptionResponse) (ClustersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClustersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClustersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClustersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *ClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureStackHCI/clusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (ClustersClientListBySubscriptionResponse, error) {
	result := ClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterList); err != nil {
		return ClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - cluster - Details of the HCI cluster.
//   - options - ClustersClientUpdateOptions contains the optional parameters for the ClustersClient.Update method.
func (client *ClustersClient) Update(ctx context.Context, resourceGroupName string, clusterName string, cluster ClusterPatch, options *ClustersClientUpdateOptions) (ClustersClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, cluster, options)
	if err != nil {
		return ClustersClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClustersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClustersClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, cluster ClusterPatch, options *ClustersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, cluster); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ClustersClient) updateHandleResponse(resp *http.Response) (ClustersClientUpdateResponse, error) {
	result := ClustersClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cluster); err != nil {
		return ClustersClientUpdateResponse{}, err
	}
	return result, nil
}

// BeginUploadCertificate - Upload certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - uploadCertificateRequest - Upload certificate request.
//   - options - ClustersClientBeginUploadCertificateOptions contains the optional parameters for the ClustersClient.BeginUploadCertificate
//     method.
func (client *ClustersClient) BeginUploadCertificate(ctx context.Context, resourceGroupName string, clusterName string, uploadCertificateRequest UploadCertificateRequest, options *ClustersClientBeginUploadCertificateOptions) (*runtime.Poller[ClustersClientUploadCertificateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.uploadCertificate(ctx, resourceGroupName, clusterName, uploadCertificateRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientUploadCertificateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientUploadCertificateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UploadCertificate - Upload certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ClustersClient) uploadCertificate(ctx context.Context, resourceGroupName string, clusterName string, uploadCertificateRequest UploadCertificateRequest, options *ClustersClientBeginUploadCertificateOptions) (*http.Response, error) {
	var err error
	req, err := client.uploadCertificateCreateRequest(ctx, resourceGroupName, clusterName, uploadCertificateRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// uploadCertificateCreateRequest creates the UploadCertificate request.
func (client *ClustersClient) uploadCertificateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, uploadCertificateRequest UploadCertificateRequest, options *ClustersClientBeginUploadCertificateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/uploadCertificate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, uploadCertificateRequest); err != nil {
		return nil, err
	}
	return req, nil
}
