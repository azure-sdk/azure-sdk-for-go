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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscovery"
	"net/http"
	"net/url"
	"regexp"
)

// SQLRunAsAccountsControllerServer is a fake server for instances of the armmigrationdiscovery.SQLRunAsAccountsControllerClient type.
type SQLRunAsAccountsControllerServer struct {
	// Get is the fake for method SQLRunAsAccountsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, accountName string, options *armmigrationdiscovery.SQLRunAsAccountsControllerClientGetOptions) (resp azfake.Responder[armmigrationdiscovery.SQLRunAsAccountsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySQLSitePager is the fake for method SQLRunAsAccountsControllerClient.NewListBySQLSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySQLSitePager func(resourceGroupName string, siteName string, sqlSiteName string, options *armmigrationdiscovery.SQLRunAsAccountsControllerClientListBySQLSiteOptions) (resp azfake.PagerResponder[armmigrationdiscovery.SQLRunAsAccountsControllerClientListBySQLSiteResponse])
}

// NewSQLRunAsAccountsControllerServerTransport creates a new instance of SQLRunAsAccountsControllerServerTransport with the provided implementation.
// The returned SQLRunAsAccountsControllerServerTransport instance is connected to an instance of armmigrationdiscovery.SQLRunAsAccountsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLRunAsAccountsControllerServerTransport(srv *SQLRunAsAccountsControllerServer) *SQLRunAsAccountsControllerServerTransport {
	return &SQLRunAsAccountsControllerServerTransport{
		srv:                   srv,
		newListBySQLSitePager: newTracker[azfake.PagerResponder[armmigrationdiscovery.SQLRunAsAccountsControllerClientListBySQLSiteResponse]](),
	}
}

// SQLRunAsAccountsControllerServerTransport connects instances of armmigrationdiscovery.SQLRunAsAccountsControllerClient to instances of SQLRunAsAccountsControllerServer.
// Don't use this type directly, use NewSQLRunAsAccountsControllerServerTransport instead.
type SQLRunAsAccountsControllerServerTransport struct {
	srv                   *SQLRunAsAccountsControllerServer
	newListBySQLSitePager *tracker[azfake.PagerResponder[armmigrationdiscovery.SQLRunAsAccountsControllerClientListBySQLSiteResponse]]
}

// Do implements the policy.Transporter interface for SQLRunAsAccountsControllerServerTransport.
func (s *SQLRunAsAccountsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLRunAsAccountsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlRunAsAccountsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlRunAsAccountsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLRunAsAccountsControllerClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SQLRunAsAccountsControllerClient.NewListBySQLSitePager":
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

func (s *SQLRunAsAccountsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, sqlSiteNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLRunAsAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLRunAsAccountsControllerServerTransport) dispatchNewListBySQLSitePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySQLSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySQLSitePager not implemented")}
	}
	newListBySQLSitePager := s.newListBySQLSitePager.get(req)
	if newListBySQLSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts`
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
		server.PagerResponderInjectNextLinks(newListBySQLSitePager, req, func(page *armmigrationdiscovery.SQLRunAsAccountsControllerClientListBySQLSiteResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to SQLRunAsAccountsControllerServerTransport
var sqlRunAsAccountsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
