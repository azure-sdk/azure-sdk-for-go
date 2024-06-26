//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EvidenceClient contains the methods for the Evidence group.
// Don't use this type directly, use NewEvidenceClient() instead.
type EvidenceClient struct {
	internal *arm.Client
}

// NewEvidenceClient creates a new instance of EvidenceClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEvidenceClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EvidenceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EvidenceClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update an evidence a specified report
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - evidenceName - The evidence name.
//   - properties - Parameters for the create or update operation
//   - options - EvidenceClientCreateOrUpdateOptions contains the optional parameters for the EvidenceClient.CreateOrUpdate method.
func (client *EvidenceClient) CreateOrUpdate(ctx context.Context, reportName string, evidenceName string, properties EvidenceResource, options *EvidenceClientCreateOrUpdateOptions) (EvidenceClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "EvidenceClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, reportName, evidenceName, properties, options)
	if err != nil {
		return EvidenceClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EvidenceClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return EvidenceClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EvidenceClient) createOrUpdateCreateRequest(ctx context.Context, reportName string, evidenceName string, properties EvidenceResource, options *EvidenceClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/evidences/{evidenceName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if evidenceName == "" {
		return nil, errors.New("parameter evidenceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{evidenceName}", url.PathEscape(evidenceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	if options != nil && options.OfferGUID != nil {
		reqQP.Set("offerGuid", *options.OfferGUID)
	}
	if options != nil && options.ReportCreatorTenantID != nil {
		reqQP.Set("reportCreatorTenantId", *options.ReportCreatorTenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EvidenceClient) createOrUpdateHandleResponse(resp *http.Response) (EvidenceClientCreateOrUpdateResponse, error) {
	result := EvidenceClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvidenceResource); err != nil {
		return EvidenceClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an existent evidence from a specified report
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - evidenceName - The evidence name.
//   - options - EvidenceClientDeleteOptions contains the optional parameters for the EvidenceClient.Delete method.
func (client *EvidenceClient) Delete(ctx context.Context, reportName string, evidenceName string, options *EvidenceClientDeleteOptions) (EvidenceClientDeleteResponse, error) {
	var err error
	const operationName = "EvidenceClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, reportName, evidenceName, options)
	if err != nil {
		return EvidenceClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EvidenceClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EvidenceClientDeleteResponse{}, err
	}
	return EvidenceClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EvidenceClient) deleteCreateRequest(ctx context.Context, reportName string, evidenceName string, options *EvidenceClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/evidences/{evidenceName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if evidenceName == "" {
		return nil, errors.New("parameter evidenceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{evidenceName}", url.PathEscape(evidenceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Download - Download evidence file.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - evidenceName - The evidence name.
//   - body - Parameters for the query operation
//   - options - EvidenceClientDownloadOptions contains the optional parameters for the EvidenceClient.Download method.
func (client *EvidenceClient) Download(ctx context.Context, reportName string, evidenceName string, body EvidenceFileDownloadRequest, options *EvidenceClientDownloadOptions) (EvidenceClientDownloadResponse, error) {
	var err error
	const operationName = "EvidenceClient.Download"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.downloadCreateRequest(ctx, reportName, evidenceName, body, options)
	if err != nil {
		return EvidenceClientDownloadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EvidenceClientDownloadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EvidenceClientDownloadResponse{}, err
	}
	resp, err := client.downloadHandleResponse(httpResp)
	return resp, err
}

// downloadCreateRequest creates the Download request.
func (client *EvidenceClient) downloadCreateRequest(ctx context.Context, reportName string, evidenceName string, body EvidenceFileDownloadRequest, options *EvidenceClientDownloadOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/evidences/{evidenceName}/download"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if evidenceName == "" {
		return nil, errors.New("parameter evidenceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{evidenceName}", url.PathEscape(evidenceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// downloadHandleResponse handles the Download response.
func (client *EvidenceClient) downloadHandleResponse(resp *http.Response) (EvidenceClientDownloadResponse, error) {
	result := EvidenceClientDownloadResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvidenceFileDownloadResponse); err != nil {
		return EvidenceClientDownloadResponse{}, err
	}
	return result, nil
}

// Get - Get the evidence metadata
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - evidenceName - The evidence name.
//   - options - EvidenceClientGetOptions contains the optional parameters for the EvidenceClient.Get method.
func (client *EvidenceClient) Get(ctx context.Context, reportName string, evidenceName string, options *EvidenceClientGetOptions) (EvidenceClientGetResponse, error) {
	var err error
	const operationName = "EvidenceClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, reportName, evidenceName, options)
	if err != nil {
		return EvidenceClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EvidenceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EvidenceClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EvidenceClient) getCreateRequest(ctx context.Context, reportName string, evidenceName string, options *EvidenceClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/evidences/{evidenceName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if evidenceName == "" {
		return nil, errors.New("parameter evidenceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{evidenceName}", url.PathEscape(evidenceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EvidenceClient) getHandleResponse(resp *http.Response) (EvidenceClientGetResponse, error) {
	result := EvidenceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvidenceResource); err != nil {
		return EvidenceClientGetResponse{}, err
	}
	return result, nil
}

// NewListByReportPager - Returns a paginated list of evidences for a specified report.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - options - EvidenceClientListByReportOptions contains the optional parameters for the EvidenceClient.NewListByReportPager
//     method.
func (client *EvidenceClient) NewListByReportPager(reportName string, options *EvidenceClientListByReportOptions) *runtime.Pager[EvidenceClientListByReportResponse] {
	return runtime.NewPager(runtime.PagingHandler[EvidenceClientListByReportResponse]{
		More: func(page EvidenceClientListByReportResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EvidenceClientListByReportResponse) (EvidenceClientListByReportResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EvidenceClient.NewListByReportPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByReportCreateRequest(ctx, reportName, options)
			}, nil)
			if err != nil {
				return EvidenceClientListByReportResponse{}, err
			}
			return client.listByReportHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByReportCreateRequest creates the ListByReport request.
func (client *EvidenceClient) listByReportCreateRequest(ctx context.Context, reportName string, options *EvidenceClientListByReportOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/evidences"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-27")
	if options != nil && options.OfferGUID != nil {
		reqQP.Set("offerGuid", *options.OfferGUID)
	}
	if options != nil && options.ReportCreatorTenantID != nil {
		reqQP.Set("reportCreatorTenantId", *options.ReportCreatorTenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReportHandleResponse handles the ListByReport response.
func (client *EvidenceClient) listByReportHandleResponse(resp *http.Response) (EvidenceClientListByReportResponse, error) {
	result := EvidenceClientListByReportResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvidenceResourceListResult); err != nil {
		return EvidenceClientListByReportResponse{}, err
	}
	return result, nil
}
