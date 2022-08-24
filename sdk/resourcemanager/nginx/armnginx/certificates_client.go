//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnginx

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

// CertificatesClient contains the methods for the Certificates group.
// Don't use this type directly, use NewCertificatesClient() instead.
type CertificatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCertificatesClient creates a new instance of CertificatesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificatesClient, error) {
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
	client := &CertificatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create or update the Nginx certificates for given Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// deploymentName - The name of targeted Nginx deployment
// certificateName - The name of certificate
// options - CertificatesClientBeginCreateOptions contains the optional parameters for the CertificatesClient.BeginCreate
// method.
func (client *CertificatesClient) BeginCreate(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientBeginCreateOptions) (*runtime.Poller[CertificatesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, deploymentName, certificateName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CertificatesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CertificatesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create or update the Nginx certificates for given Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
func (client *CertificatesClient) create(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, deploymentName, certificateName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *CertificatesClient) createCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginDelete - Deletes a certificate from the nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// deploymentName - The name of targeted Nginx deployment
// certificateName - The name of certificate
// options - CertificatesClientBeginDeleteOptions contains the optional parameters for the CertificatesClient.BeginDelete
// method.
func (client *CertificatesClient) BeginDelete(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientBeginDeleteOptions) (*runtime.Poller[CertificatesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, deploymentName, certificateName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CertificatesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[CertificatesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a certificate from the nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
func (client *CertificatesClient) deleteOperation(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, deploymentName, certificateName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a certificate of given Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// deploymentName - The name of targeted Nginx deployment
// certificateName - The name of certificate
// options - CertificatesClientGetOptions contains the optional parameters for the CertificatesClient.Get method.
func (client *CertificatesClient) Get(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientGetOptions) (CertificatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, deploymentName, certificateName, options)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, certificateName string, options *CertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CertificatesClient) getHandleResponse(resp *http.Response) (CertificatesClientGetResponse, error) {
	result := CertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all certificates of given Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// deploymentName - The name of targeted Nginx deployment
// options - CertificatesClientListOptions contains the optional parameters for the CertificatesClient.List method.
func (client *CertificatesClient) NewListPager(resourceGroupName string, deploymentName string, options *CertificatesClientListOptions) *runtime.Pager[CertificatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificatesClientListResponse]{
		More: func(page CertificatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificatesClientListResponse) (CertificatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, deploymentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CertificatesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CertificatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CertificatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CertificatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, options *CertificatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CertificatesClient) listHandleResponse(resp *http.Response) (CertificatesClientListResponse, error) {
	result := CertificatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateListResponse); err != nil {
		return CertificatesClientListResponse{}, err
	}
	return result, nil
}
