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
	"time"
)

// CertificateOrdersDiagnosticsClient contains the methods for the CertificateOrdersDiagnostics group.
// Don't use this type directly, use NewCertificateOrdersDiagnosticsClient() instead.
type CertificateOrdersDiagnosticsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCertificateOrdersDiagnosticsClient creates a new instance of CertificateOrdersDiagnosticsClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCertificateOrdersDiagnosticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificateOrdersDiagnosticsClient, error) {
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
	client := &CertificateOrdersDiagnosticsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetAppServiceCertificateOrderDetectorResponse - Description for Microsoft.CertificateRegistration call to get a detector
// response from App Lens.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// certificateOrderName - The certificate order name for which the response is needed.
// detectorName - The detector name which needs to be run.
// options - CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseOptions contains the optional
// parameters for the CertificateOrdersDiagnosticsClient.GetAppServiceCertificateOrderDetectorResponse method.
func (client *CertificateOrdersDiagnosticsClient) GetAppServiceCertificateOrderDetectorResponse(ctx context.Context, resourceGroupName string, certificateOrderName string, detectorName string, options *CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseOptions) (CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse, error) {
	req, err := client.getAppServiceCertificateOrderDetectorResponseCreateRequest(ctx, resourceGroupName, certificateOrderName, detectorName, options)
	if err != nil {
		return CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAppServiceCertificateOrderDetectorResponseHandleResponse(resp)
}

// getAppServiceCertificateOrderDetectorResponseCreateRequest creates the GetAppServiceCertificateOrderDetectorResponse request.
func (client *CertificateOrdersDiagnosticsClient) getAppServiceCertificateOrderDetectorResponseCreateRequest(ctx context.Context, resourceGroupName string, certificateOrderName string, detectorName string, options *CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/detectors/{detectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if certificateOrderName == "" {
		return nil, errors.New("parameter certificateOrderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateOrderName}", url.PathEscape(certificateOrderName))
	if detectorName == "" {
		return nil, errors.New("parameter detectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{detectorName}", url.PathEscape(detectorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", options.StartTime.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", options.EndTime.Format(time.RFC3339Nano))
	}
	if options != nil && options.TimeGrain != nil {
		reqQP.Set("timeGrain", *options.TimeGrain)
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAppServiceCertificateOrderDetectorResponseHandleResponse handles the GetAppServiceCertificateOrderDetectorResponse response.
func (client *CertificateOrdersDiagnosticsClient) getAppServiceCertificateOrderDetectorResponseHandleResponse(resp *http.Response) (CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse, error) {
	result := CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DetectorResponse); err != nil {
		return CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse{}, err
	}
	return result, nil
}

// NewListAppServiceCertificateOrderDetectorResponsePager - Description for Microsoft.CertificateRegistration to get the list
// of detectors for this RP.
// Generated from API version 2022-09-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// certificateOrderName - The certificate order name for which the response is needed.
// options - CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseOptions contains the optional
// parameters for the CertificateOrdersDiagnosticsClient.ListAppServiceCertificateOrderDetectorResponse method.
func (client *CertificateOrdersDiagnosticsClient) NewListAppServiceCertificateOrderDetectorResponsePager(resourceGroupName string, certificateOrderName string, options *CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseOptions) *runtime.Pager[CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse]{
		More: func(page CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse) (CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAppServiceCertificateOrderDetectorResponseCreateRequest(ctx, resourceGroupName, certificateOrderName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAppServiceCertificateOrderDetectorResponseHandleResponse(resp)
		},
	})
}

// listAppServiceCertificateOrderDetectorResponseCreateRequest creates the ListAppServiceCertificateOrderDetectorResponse request.
func (client *CertificateOrdersDiagnosticsClient) listAppServiceCertificateOrderDetectorResponseCreateRequest(ctx context.Context, resourceGroupName string, certificateOrderName string, options *CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/detectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if certificateOrderName == "" {
		return nil, errors.New("parameter certificateOrderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateOrderName}", url.PathEscape(certificateOrderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAppServiceCertificateOrderDetectorResponseHandleResponse handles the ListAppServiceCertificateOrderDetectorResponse response.
func (client *CertificateOrdersDiagnosticsClient) listAppServiceCertificateOrderDetectorResponseHandleResponse(resp *http.Response) (CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse, error) {
	result := CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DetectorResponseCollection); err != nil {
		return CertificateOrdersDiagnosticsClientListAppServiceCertificateOrderDetectorResponseResponse{}, err
	}
	return result, nil
}
