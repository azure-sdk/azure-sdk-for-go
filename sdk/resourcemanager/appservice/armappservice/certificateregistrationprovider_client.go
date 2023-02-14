//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// CertificateRegistrationProviderClient contains the methods for the CertificateRegistrationProvider group.
// Don't use this type directly, use NewCertificateRegistrationProviderClient() instead.
type CertificateRegistrationProviderClient struct {
	host string
	pl   runtime.Pipeline
}

// NewCertificateRegistrationProviderClient creates a new instance of CertificateRegistrationProviderClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCertificateRegistrationProviderClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificateRegistrationProviderClient, error) {
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
	client := &CertificateRegistrationProviderClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListOperationsPager - Description for Implements Csm operations Api to exposes the list of available Csm Apis under
// the resource provider
//
// Generated from API version 2022-03-01
//   - options - CertificateRegistrationProviderClientListOperationsOptions contains the optional parameters for the CertificateRegistrationProviderClient.NewListOperationsPager
//     method.
func (client *CertificateRegistrationProviderClient) NewListOperationsPager(options *CertificateRegistrationProviderClientListOperationsOptions) *runtime.Pager[CertificateRegistrationProviderClientListOperationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificateRegistrationProviderClientListOperationsResponse]{
		More: func(page CertificateRegistrationProviderClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificateRegistrationProviderClientListOperationsResponse) (CertificateRegistrationProviderClientListOperationsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listOperationsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CertificateRegistrationProviderClientListOperationsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CertificateRegistrationProviderClientListOperationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CertificateRegistrationProviderClientListOperationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listOperationsHandleResponse(resp)
		},
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *CertificateRegistrationProviderClient) listOperationsCreateRequest(ctx context.Context, options *CertificateRegistrationProviderClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CertificateRegistration/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *CertificateRegistrationProviderClient) listOperationsHandleResponse(resp *http.Response) (CertificateRegistrationProviderClientListOperationsResponse, error) {
	result := CertificateRegistrationProviderClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CsmOperationCollection); err != nil {
		return CertificateRegistrationProviderClientListOperationsResponse{}, err
	}
	return result, nil
}
