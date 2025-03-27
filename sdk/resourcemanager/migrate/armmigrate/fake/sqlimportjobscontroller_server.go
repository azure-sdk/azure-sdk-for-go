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
)

// SQLImportJobsControllerServer is a fake server for instances of the armmigrate.SQLImportJobsControllerClient type.
type SQLImportJobsControllerServer struct {
	// GetImportjob is the fake for method SQLImportJobsControllerClient.GetImportjob
	// HTTP status codes to indicate success: http.StatusOK
	GetImportjob func(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, jobName string, options *armmigrate.SQLImportJobsControllerClientGetImportjobOptions) (resp azfake.Responder[armmigrate.SQLImportJobsControllerClientGetImportjobResponse], errResp azfake.ErrorResponder)

	// NewListImportjobsPager is the fake for method SQLImportJobsControllerClient.NewListImportjobsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListImportjobsPager func(resourceGroupName string, siteName string, sqlSiteName string, options *armmigrate.SQLImportJobsControllerClientListImportjobsOptions) (resp azfake.PagerResponder[armmigrate.SQLImportJobsControllerClientListImportjobsResponse])
}

// NewSQLImportJobsControllerServerTransport creates a new instance of SQLImportJobsControllerServerTransport with the provided implementation.
// The returned SQLImportJobsControllerServerTransport instance is connected to an instance of armmigrate.SQLImportJobsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLImportJobsControllerServerTransport(srv *SQLImportJobsControllerServer) *SQLImportJobsControllerServerTransport {
	return &SQLImportJobsControllerServerTransport{
		srv:                    srv,
		newListImportjobsPager: newTracker[azfake.PagerResponder[armmigrate.SQLImportJobsControllerClientListImportjobsResponse]](),
	}
}

// SQLImportJobsControllerServerTransport connects instances of armmigrate.SQLImportJobsControllerClient to instances of SQLImportJobsControllerServer.
// Don't use this type directly, use NewSQLImportJobsControllerServerTransport instead.
type SQLImportJobsControllerServerTransport struct {
	srv                    *SQLImportJobsControllerServer
	newListImportjobsPager *tracker[azfake.PagerResponder[armmigrate.SQLImportJobsControllerClientListImportjobsResponse]]
}

// Do implements the policy.Transporter interface for SQLImportJobsControllerServerTransport.
func (s *SQLImportJobsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLImportJobsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlImportJobsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlImportJobsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLImportJobsControllerClient.GetImportjob":
				res.resp, res.err = s.dispatchGetImportjob(req)
			case "SQLImportJobsControllerClient.NewListImportjobsPager":
				res.resp, res.err = s.dispatchNewListImportjobsPager(req)
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

func (s *SQLImportJobsControllerServerTransport) dispatchGetImportjob(req *http.Request) (*http.Response, error) {
	if s.srv.GetImportjob == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetImportjob not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importJobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetImportjob(req.Context(), resourceGroupNameParam, siteNameParam, sqlSiteNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImportSQLInventoryJob, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLImportJobsControllerServerTransport) dispatchNewListImportjobsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListImportjobsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListImportjobsPager not implemented")}
	}
	newListImportjobsPager := s.newListImportjobsPager.get(req)
	if newListImportjobsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlSites/(?P<sqlSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importJobs`
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
		resp := s.srv.NewListImportjobsPager(resourceGroupNameParam, siteNameParam, sqlSiteNameParam, nil)
		newListImportjobsPager = &resp
		s.newListImportjobsPager.add(req, newListImportjobsPager)
		server.PagerResponderInjectNextLinks(newListImportjobsPager, req, func(page *armmigrate.SQLImportJobsControllerClientListImportjobsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListImportjobsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListImportjobsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListImportjobsPager) {
		s.newListImportjobsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SQLImportJobsControllerServerTransport
var sqlImportJobsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
