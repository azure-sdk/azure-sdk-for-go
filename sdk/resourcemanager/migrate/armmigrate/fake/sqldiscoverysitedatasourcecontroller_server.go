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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// SQLDiscoverySiteDataSourceControllerServer is a fake server for instances of the armmigrate.SQLDiscoverySiteDataSourceControllerClient type.
type SQLDiscoverySiteDataSourceControllerServer struct {
	// BeginCreate is the fake for method SQLDiscoverySiteDataSourceControllerClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, discoverySiteDataSourceName string, body armmigrate.SQLDiscoverySiteDataSource, options *armmigrate.SQLDiscoverySiteDataSourceControllerClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrate.SQLDiscoverySiteDataSourceControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SQLDiscoverySiteDataSourceControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, discoverySiteDataSourceName string, options *armmigrate.SQLDiscoverySiteDataSourceControllerClientDeleteOptions) (resp azfake.Responder[armmigrate.SQLDiscoverySiteDataSourceControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SQLDiscoverySiteDataSourceControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, discoverySiteDataSourceName string, options *armmigrate.SQLDiscoverySiteDataSourceControllerClientGetOptions) (resp azfake.Responder[armmigrate.SQLDiscoverySiteDataSourceControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySQLSitePager is the fake for method SQLDiscoverySiteDataSourceControllerClient.NewListBySQLSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySQLSitePager func(resourceGroupName string, siteName string, sqlSiteName string, options *armmigrate.SQLDiscoverySiteDataSourceControllerClientListBySQLSiteOptions) (resp azfake.PagerResponder[armmigrate.SQLDiscoverySiteDataSourceControllerClientListBySQLSiteResponse])
}

// NewSQLDiscoverySiteDataSourceControllerServerTransport creates a new instance of SQLDiscoverySiteDataSourceControllerServerTransport with the provided implementation.
// The returned SQLDiscoverySiteDataSourceControllerServerTransport instance is connected to an instance of armmigrate.SQLDiscoverySiteDataSourceControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLDiscoverySiteDataSourceControllerServerTransport(srv *SQLDiscoverySiteDataSourceControllerServer) *SQLDiscoverySiteDataSourceControllerServerTransport {
	return &SQLDiscoverySiteDataSourceControllerServerTransport{
		srv:                   srv,
		beginCreate:           newTracker[azfake.PollerResponder[armmigrate.SQLDiscoverySiteDataSourceControllerClientCreateResponse]](),
		newListBySQLSitePager: newTracker[azfake.PagerResponder[armmigrate.SQLDiscoverySiteDataSourceControllerClientListBySQLSiteResponse]](),
	}
}

// SQLDiscoverySiteDataSourceControllerServerTransport connects instances of armmigrate.SQLDiscoverySiteDataSourceControllerClient to instances of SQLDiscoverySiteDataSourceControllerServer.
// Don't use this type directly, use NewSQLDiscoverySiteDataSourceControllerServerTransport instead.
type SQLDiscoverySiteDataSourceControllerServerTransport struct {
	srv                   *SQLDiscoverySiteDataSourceControllerServer
	beginCreate           *tracker[azfake.PollerResponder[armmigrate.SQLDiscoverySiteDataSourceControllerClientCreateResponse]]
	newListBySQLSitePager *tracker[azfake.PagerResponder[armmigrate.SQLDiscoverySiteDataSourceControllerClientListBySQLSiteResponse]]
}

// Do implements the policy.Transporter interface for SQLDiscoverySiteDataSourceControllerServerTransport.
func (s *SQLDiscoverySiteDataSourceControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLDiscoverySiteDataSourceControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlDiscoverySiteDataSourceControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlDiscoverySiteDataSourceControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLDiscoverySiteDataSourceControllerClient.BeginCreate":
				res.resp, res.err = s.dispatchBeginCreate(req)
			case "SQLDiscoverySiteDataSourceControllerClient.Delete":
				res.resp, res.err = s.dispatchDelete(req)
			case "SQLDiscoverySiteDataSourceControllerClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SQLDiscoverySiteDataSourceControllerClient.NewListBySQLSitePager":
				res.resp, res.err = s.dispatchNewListBySQLSitePager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (s *SQLDiscoverySiteDataSourceControllerServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/discoverySiteDataSources/(?P<discoverySiteDataSourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.SQLDiscoverySiteDataSource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		sqlSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlSiteName")])
		if err != nil {
			return nil, err
		}
		discoverySiteDataSourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("discoverySiteDataSourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreate(req.Context(), resourceGroupNameParam, siteNameParam, sqlSiteNameParam, discoverySiteDataSourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		s.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		s.beginCreate.remove(req)
	}

	return resp, nil
}

func (s *SQLDiscoverySiteDataSourceControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/discoverySiteDataSources/(?P<discoverySiteDataSourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	sqlSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlSiteName")])
	if err != nil {
		return nil, err
	}
	discoverySiteDataSourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("discoverySiteDataSourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, siteNameParam, sqlSiteNameParam, discoverySiteDataSourceNameParam, nil)
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

func (s *SQLDiscoverySiteDataSourceControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/discoverySiteDataSources/(?P<discoverySiteDataSourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	sqlSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlSiteName")])
	if err != nil {
		return nil, err
	}
	discoverySiteDataSourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("discoverySiteDataSourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, sqlSiteNameParam, discoverySiteDataSourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLDiscoverySiteDataSource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLDiscoverySiteDataSourceControllerServerTransport) dispatchNewListBySQLSitePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySQLSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySQLSitePager not implemented")}
	}
	newListBySQLSitePager := s.newListBySQLSitePager.get(req)
	if newListBySQLSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/discoverySiteDataSources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		sqlSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlSiteName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListBySQLSitePager(resourceGroupNameParam, siteNameParam, sqlSiteNameParam, nil)
		newListBySQLSitePager = &resp
		s.newListBySQLSitePager.add(req, newListBySQLSitePager)
		server.PagerResponderInjectNextLinks(newListBySQLSitePager, req, func(page *armmigrate.SQLDiscoverySiteDataSourceControllerClientListBySQLSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySQLSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListBySQLSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySQLSitePager) {
		s.newListBySQLSitePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SQLDiscoverySiteDataSourceControllerServerTransport
var sqlDiscoverySiteDataSourceControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
