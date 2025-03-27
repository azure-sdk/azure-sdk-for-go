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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// SQLAvailabilityGroupsControllerServer is a fake server for instances of the armmigrate.SQLAvailabilityGroupsControllerClient type.
type SQLAvailabilityGroupsControllerServer struct {
	// Get is the fake for method SQLAvailabilityGroupsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, sqlAvailabilityGroupName string, options *armmigrate.SQLAvailabilityGroupsControllerClientGetOptions) (resp azfake.Responder[armmigrate.SQLAvailabilityGroupsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySQLSitePager is the fake for method SQLAvailabilityGroupsControllerClient.NewListBySQLSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySQLSitePager func(resourceGroupName string, siteName string, sqlSiteName string, options *armmigrate.SQLAvailabilityGroupsControllerClientListBySQLSiteOptions) (resp azfake.PagerResponder[armmigrate.SQLAvailabilityGroupsControllerClientListBySQLSiteResponse])
}

// NewSQLAvailabilityGroupsControllerServerTransport creates a new instance of SQLAvailabilityGroupsControllerServerTransport with the provided implementation.
// The returned SQLAvailabilityGroupsControllerServerTransport instance is connected to an instance of armmigrate.SQLAvailabilityGroupsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLAvailabilityGroupsControllerServerTransport(srv *SQLAvailabilityGroupsControllerServer) *SQLAvailabilityGroupsControllerServerTransport {
	return &SQLAvailabilityGroupsControllerServerTransport{
		srv:                   srv,
		newListBySQLSitePager: newTracker[azfake.PagerResponder[armmigrate.SQLAvailabilityGroupsControllerClientListBySQLSiteResponse]](),
	}
}

// SQLAvailabilityGroupsControllerServerTransport connects instances of armmigrate.SQLAvailabilityGroupsControllerClient to instances of SQLAvailabilityGroupsControllerServer.
// Don't use this type directly, use NewSQLAvailabilityGroupsControllerServerTransport instead.
type SQLAvailabilityGroupsControllerServerTransport struct {
	srv                   *SQLAvailabilityGroupsControllerServer
	newListBySQLSitePager *tracker[azfake.PagerResponder[armmigrate.SQLAvailabilityGroupsControllerClientListBySQLSiteResponse]]
}

// Do implements the policy.Transporter interface for SQLAvailabilityGroupsControllerServerTransport.
func (s *SQLAvailabilityGroupsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLAvailabilityGroupsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlAvailabilityGroupsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlAvailabilityGroupsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLAvailabilityGroupsControllerClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SQLAvailabilityGroupsControllerClient.NewListBySQLSitePager":
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

func (s *SQLAvailabilityGroupsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlAvailabilityGroups/(?P<sqlAvailabilityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	sqlAvailabilityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlAvailabilityGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, sqlSiteNameParam, sqlAvailabilityGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLAvailabilityGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLAvailabilityGroupsControllerServerTransport) dispatchNewListBySQLSitePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySQLSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySQLSitePager not implemented")}
	}
	newListBySQLSitePager := s.newListBySQLSitePager.get(req)
	if newListBySQLSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlAvailabilityGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam := getOptional(topUnescaped)
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		totalRecordCountUnescaped, err := url.QueryUnescape(qp.Get("totalRecordCount"))
		if err != nil {
			return nil, err
		}
		totalRecordCountParam, err := parseOptional(totalRecordCountUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
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
		var options *armmigrate.SQLAvailabilityGroupsControllerClientListBySQLSiteOptions
		if filterParam != nil || topParam != nil || continuationTokenParam != nil || totalRecordCountParam != nil {
			options = &armmigrate.SQLAvailabilityGroupsControllerClientListBySQLSiteOptions{
				Filter:            filterParam,
				Top:               topParam,
				ContinuationToken: continuationTokenParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := s.srv.NewListBySQLSitePager(resourceGroupNameParam, siteNameParam, sqlSiteNameParam, options)
		newListBySQLSitePager = &resp
		s.newListBySQLSitePager.add(req, newListBySQLSitePager)
		server.PagerResponderInjectNextLinks(newListBySQLSitePager, req, func(page *armmigrate.SQLAvailabilityGroupsControllerClientListBySQLSiteResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to SQLAvailabilityGroupsControllerServerTransport
var sqlAvailabilityGroupsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
