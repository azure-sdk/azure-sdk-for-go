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

// SnapshotResourcesServer is a fake server for instances of the armappcomplianceautomation.SnapshotResourcesClient type.
type SnapshotResourcesServer struct {
	// BeginDownload is the fake for method SnapshotResourcesClient.BeginDownload
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDownload func(ctx context.Context, reportName string, snapshotName string, body armappcomplianceautomation.SnapshotDownloadRequest, options *armappcomplianceautomation.SnapshotResourcesClientBeginDownloadOptions) (resp azfake.PollerResponder[armappcomplianceautomation.SnapshotResourcesClientDownloadResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SnapshotResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, reportName string, snapshotName string, options *armappcomplianceautomation.SnapshotResourcesClientGetOptions) (resp azfake.Responder[armappcomplianceautomation.SnapshotResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByReportResourcePager is the fake for method SnapshotResourcesClient.NewListByReportResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReportResourcePager func(reportName string, options *armappcomplianceautomation.SnapshotResourcesClientListByReportResourceOptions) (resp azfake.PagerResponder[armappcomplianceautomation.SnapshotResourcesClientListByReportResourceResponse])
}

// NewSnapshotResourcesServerTransport creates a new instance of SnapshotResourcesServerTransport with the provided implementation.
// The returned SnapshotResourcesServerTransport instance is connected to an instance of armappcomplianceautomation.SnapshotResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSnapshotResourcesServerTransport(srv *SnapshotResourcesServer) *SnapshotResourcesServerTransport {
	return &SnapshotResourcesServerTransport{
		srv:                          srv,
		beginDownload:                newTracker[azfake.PollerResponder[armappcomplianceautomation.SnapshotResourcesClientDownloadResponse]](),
		newListByReportResourcePager: newTracker[azfake.PagerResponder[armappcomplianceautomation.SnapshotResourcesClientListByReportResourceResponse]](),
	}
}

// SnapshotResourcesServerTransport connects instances of armappcomplianceautomation.SnapshotResourcesClient to instances of SnapshotResourcesServer.
// Don't use this type directly, use NewSnapshotResourcesServerTransport instead.
type SnapshotResourcesServerTransport struct {
	srv                          *SnapshotResourcesServer
	beginDownload                *tracker[azfake.PollerResponder[armappcomplianceautomation.SnapshotResourcesClientDownloadResponse]]
	newListByReportResourcePager *tracker[azfake.PagerResponder[armappcomplianceautomation.SnapshotResourcesClientListByReportResourceResponse]]
}

// Do implements the policy.Transporter interface for SnapshotResourcesServerTransport.
func (s *SnapshotResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SnapshotResourcesClient.BeginDownload":
		resp, err = s.dispatchBeginDownload(req)
	case "SnapshotResourcesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SnapshotResourcesClient.NewListByReportResourcePager":
		resp, err = s.dispatchNewListByReportResourcePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SnapshotResourcesServerTransport) dispatchBeginDownload(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDownload == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDownload not implemented")}
	}
	beginDownload := s.beginDownload.get(req)
	if beginDownload == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/download`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.SnapshotDownloadRequest](req)
		if err != nil {
			return nil, err
		}
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDownload(req.Context(), reportNameParam, snapshotNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDownload = &respr
		s.beginDownload.add(req, beginDownload)
	}

	resp, err := server.PollerResponderNext(beginDownload, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginDownload.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDownload) {
		s.beginDownload.remove(req)
	}

	return resp, nil
}

func (s *SnapshotResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), reportNameParam, snapshotNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnapshotResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SnapshotResourcesServerTransport) dispatchNewListByReportResourcePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByReportResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReportResourcePager not implemented")}
	}
	newListByReportResourcePager := s.newListByReportResourcePager.get(req)
	if newListByReportResourcePager == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots`
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
		reportCreatorTenantIDUnescaped, err := url.QueryUnescape(qp.Get("reportCreatorTenantId"))
		if err != nil {
			return nil, err
		}
		reportCreatorTenantIDParam := getOptional(reportCreatorTenantIDUnescaped)
		offerGUIDUnescaped, err := url.QueryUnescape(qp.Get("offerGuid"))
		if err != nil {
			return nil, err
		}
		offerGUIDParam := getOptional(offerGUIDUnescaped)
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		var options *armappcomplianceautomation.SnapshotResourcesClientListByReportResourceOptions
		if skipTokenParam != nil || topParam != nil || selectParam != nil || reportCreatorTenantIDParam != nil || offerGUIDParam != nil {
			options = &armappcomplianceautomation.SnapshotResourcesClientListByReportResourceOptions{
				SkipToken:             skipTokenParam,
				Top:                   topParam,
				Select:                selectParam,
				ReportCreatorTenantID: reportCreatorTenantIDParam,
				OfferGUID:             offerGUIDParam,
			}
		}
		resp := s.srv.NewListByReportResourcePager(reportNameParam, options)
		newListByReportResourcePager = &resp
		s.newListByReportResourcePager.add(req, newListByReportResourcePager)
		server.PagerResponderInjectNextLinks(newListByReportResourcePager, req, func(page *armappcomplianceautomation.SnapshotResourcesClientListByReportResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReportResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByReportResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReportResourcePager) {
		s.newListByReportResourcePager.remove(req)
	}
	return resp, nil
}
