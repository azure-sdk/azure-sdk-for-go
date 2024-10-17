//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// EvidenceServer is a fake server for instances of the armappcomplianceautomation.EvidenceClient type.
type EvidenceServer struct {
	// CreateOrUpdate is the fake for method EvidenceClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, reportName string, evidenceName string, resource armappcomplianceautomation.EvidenceResource, options *armappcomplianceautomation.EvidenceClientCreateOrUpdateOptions) (resp azfake.Responder[armappcomplianceautomation.EvidenceClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method EvidenceClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, reportName string, evidenceName string, options *armappcomplianceautomation.EvidenceClientDeleteOptions) (resp azfake.Responder[armappcomplianceautomation.EvidenceClientDeleteResponse], errResp azfake.ErrorResponder)

	// Download is the fake for method EvidenceClient.Download
	// HTTP status codes to indicate success: http.StatusOK
	Download func(ctx context.Context, reportName string, evidenceName string, body armappcomplianceautomation.EvidenceFileDownloadRequest, options *armappcomplianceautomation.EvidenceClientDownloadOptions) (resp azfake.Responder[armappcomplianceautomation.EvidenceClientDownloadResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EvidenceClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, reportName string, evidenceName string, options *armappcomplianceautomation.EvidenceClientGetOptions) (resp azfake.Responder[armappcomplianceautomation.EvidenceClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByReportPager is the fake for method EvidenceClient.NewListByReportPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReportPager func(reportName string, options *armappcomplianceautomation.EvidenceClientListByReportOptions) (resp azfake.PagerResponder[armappcomplianceautomation.EvidenceClientListByReportResponse])
}

// NewEvidenceServerTransport creates a new instance of EvidenceServerTransport with the provided implementation.
// The returned EvidenceServerTransport instance is connected to an instance of armappcomplianceautomation.EvidenceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEvidenceServerTransport(srv *EvidenceServer) *EvidenceServerTransport {
	return &EvidenceServerTransport{
		srv:                  srv,
		newListByReportPager: newTracker[azfake.PagerResponder[armappcomplianceautomation.EvidenceClientListByReportResponse]](),
	}
}

// EvidenceServerTransport connects instances of armappcomplianceautomation.EvidenceClient to instances of EvidenceServer.
// Don't use this type directly, use NewEvidenceServerTransport instead.
type EvidenceServerTransport struct {
	srv                  *EvidenceServer
	newListByReportPager *tracker[azfake.PagerResponder[armappcomplianceautomation.EvidenceClientListByReportResponse]]
}

// Do implements the policy.Transporter interface for EvidenceServerTransport.
func (e *EvidenceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EvidenceClient.CreateOrUpdate":
		resp, err = e.dispatchCreateOrUpdate(req)
	case "EvidenceClient.Delete":
		resp, err = e.dispatchDelete(req)
	case "EvidenceClient.Download":
		resp, err = e.dispatchDownload(req)
	case "EvidenceClient.Get":
		resp, err = e.dispatchGet(req)
	case "EvidenceClient.NewListByReportPager":
		resp, err = e.dispatchNewListByReportPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EvidenceServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/evidences/(?P<evidenceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.EvidenceResource](req)
	if err != nil {
		return nil, err
	}
	offerGUIDUnescaped, err := url.QueryUnescape(qp.Get("offerGuid"))
	if err != nil {
		return nil, err
	}
	offerGUIDParam := getOptional(offerGUIDUnescaped)
	reportCreatorTenantIDUnescaped, err := url.QueryUnescape(qp.Get("reportCreatorTenantId"))
	if err != nil {
		return nil, err
	}
	reportCreatorTenantIDParam := getOptional(reportCreatorTenantIDUnescaped)
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	evidenceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("evidenceName")])
	if err != nil {
		return nil, err
	}
	var options *armappcomplianceautomation.EvidenceClientCreateOrUpdateOptions
	if offerGUIDParam != nil || reportCreatorTenantIDParam != nil {
		options = &armappcomplianceautomation.EvidenceClientCreateOrUpdateOptions{
			OfferGUID:             offerGUIDParam,
			ReportCreatorTenantID: reportCreatorTenantIDParam,
		}
	}
	respr, errRespr := e.srv.CreateOrUpdate(req.Context(), reportNameParam, evidenceNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EvidenceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EvidenceServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if e.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/evidences/(?P<evidenceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	evidenceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("evidenceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Delete(req.Context(), reportNameParam, evidenceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EvidenceServerTransport) dispatchDownload(req *http.Request) (*http.Response, error) {
	if e.srv.Download == nil {
		return nil, &nonRetriableError{errors.New("fake for method Download not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/evidences/(?P<evidenceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/download`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.EvidenceFileDownloadRequest](req)
	if err != nil {
		return nil, err
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	evidenceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("evidenceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Download(req.Context(), reportNameParam, evidenceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EvidenceFileDownloadResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EvidenceServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/evidences/(?P<evidenceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	evidenceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("evidenceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), reportNameParam, evidenceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EvidenceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EvidenceServerTransport) dispatchNewListByReportPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByReportPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReportPager not implemented")}
	}
	newListByReportPager := e.newListByReportPager.get(req)
	if newListByReportPager == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/evidences`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
		if err != nil {
			return nil, err
		}
		selectParam := getOptional(selectUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		offerGUIDUnescaped, err := url.QueryUnescape(qp.Get("offerGuid"))
		if err != nil {
			return nil, err
		}
		offerGUIDParam := getOptional(offerGUIDUnescaped)
		reportCreatorTenantIDUnescaped, err := url.QueryUnescape(qp.Get("reportCreatorTenantId"))
		if err != nil {
			return nil, err
		}
		reportCreatorTenantIDParam := getOptional(reportCreatorTenantIDUnescaped)
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		var options *armappcomplianceautomation.EvidenceClientListByReportOptions
		if skipTokenParam != nil || topParam != nil || selectParam != nil || filterParam != nil || orderbyParam != nil || offerGUIDParam != nil || reportCreatorTenantIDParam != nil {
			options = &armappcomplianceautomation.EvidenceClientListByReportOptions{
				SkipToken:             skipTokenParam,
				Top:                   topParam,
				Select:                selectParam,
				Filter:                filterParam,
				Orderby:               orderbyParam,
				OfferGUID:             offerGUIDParam,
				ReportCreatorTenantID: reportCreatorTenantIDParam,
			}
		}
		resp := e.srv.NewListByReportPager(reportNameParam, options)
		newListByReportPager = &resp
		e.newListByReportPager.add(req, newListByReportPager)
		server.PagerResponderInjectNextLinks(newListByReportPager, req, func(page *armappcomplianceautomation.EvidenceClientListByReportResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReportPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByReportPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReportPager) {
		e.newListByReportPager.remove(req)
	}
	return resp, nil
}
