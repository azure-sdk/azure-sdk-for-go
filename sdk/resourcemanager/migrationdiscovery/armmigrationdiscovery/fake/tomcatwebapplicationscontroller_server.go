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
	"strconv"
)

// TomcatWebApplicationsControllerServer is a fake server for instances of the armmigrationdiscovery.TomcatWebApplicationsControllerClient type.
type TomcatWebApplicationsControllerServer struct {
	// Get is the fake for method TomcatWebApplicationsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, webApplicationName string, options *armmigrationdiscovery.TomcatWebApplicationsControllerClientGetOptions) (resp azfake.Responder[armmigrationdiscovery.TomcatWebApplicationsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByWebAppSitePager is the fake for method TomcatWebApplicationsControllerClient.NewListByWebAppSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWebAppSitePager func(resourceGroupName string, siteName string, webAppSiteName string, options *armmigrationdiscovery.TomcatWebApplicationsControllerClientListByWebAppSiteOptions) (resp azfake.PagerResponder[armmigrationdiscovery.TomcatWebApplicationsControllerClientListByWebAppSiteResponse])

	// Update is the fake for method TomcatWebApplicationsControllerClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, webApplicationName string, body any, options *armmigrationdiscovery.TomcatWebApplicationsControllerClientUpdateOptions) (resp azfake.Responder[armmigrationdiscovery.TomcatWebApplicationsControllerClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTomcatWebApplicationsControllerServerTransport creates a new instance of TomcatWebApplicationsControllerServerTransport with the provided implementation.
// The returned TomcatWebApplicationsControllerServerTransport instance is connected to an instance of armmigrationdiscovery.TomcatWebApplicationsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTomcatWebApplicationsControllerServerTransport(srv *TomcatWebApplicationsControllerServer) *TomcatWebApplicationsControllerServerTransport {
	return &TomcatWebApplicationsControllerServerTransport{
		srv:                      srv,
		newListByWebAppSitePager: newTracker[azfake.PagerResponder[armmigrationdiscovery.TomcatWebApplicationsControllerClientListByWebAppSiteResponse]](),
	}
}

// TomcatWebApplicationsControllerServerTransport connects instances of armmigrationdiscovery.TomcatWebApplicationsControllerClient to instances of TomcatWebApplicationsControllerServer.
// Don't use this type directly, use NewTomcatWebApplicationsControllerServerTransport instead.
type TomcatWebApplicationsControllerServerTransport struct {
	srv                      *TomcatWebApplicationsControllerServer
	newListByWebAppSitePager *tracker[azfake.PagerResponder[armmigrationdiscovery.TomcatWebApplicationsControllerClientListByWebAppSiteResponse]]
}

// Do implements the policy.Transporter interface for TomcatWebApplicationsControllerServerTransport.
func (t *TomcatWebApplicationsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TomcatWebApplicationsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if tomcatWebApplicationsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = tomcatWebApplicationsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TomcatWebApplicationsControllerClient.Get":
				res.resp, res.err = t.dispatchGet(req)
			case "TomcatWebApplicationsControllerClient.NewListByWebAppSitePager":
				res.resp, res.err = t.dispatchNewListByWebAppSitePager(req)
			case "TomcatWebApplicationsControllerClient.Update":
				res.resp, res.err = t.dispatchUpdate(req)
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

func (t *TomcatWebApplicationsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tomcatWebApplications/(?P<webApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
	if err != nil {
		return nil, err
	}
	webApplicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webApplicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, webApplicationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TomcatWebApplications, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TomcatWebApplicationsControllerServerTransport) dispatchNewListByWebAppSitePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByWebAppSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWebAppSitePager not implemented")}
	}
	newListByWebAppSitePager := t.newListByWebAppSitePager.get(req)
	if newListByWebAppSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tomcatWebApplications`
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
		webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrationdiscovery.TomcatWebApplicationsControllerClientListByWebAppSiteOptions
		if filterParam != nil || topParam != nil || continuationTokenParam != nil || totalRecordCountParam != nil {
			options = &armmigrationdiscovery.TomcatWebApplicationsControllerClientListByWebAppSiteOptions{
				Filter:            filterParam,
				Top:               topParam,
				ContinuationToken: continuationTokenParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := t.srv.NewListByWebAppSitePager(resourceGroupNameParam, siteNameParam, webAppSiteNameParam, options)
		newListByWebAppSitePager = &resp
		t.newListByWebAppSitePager.add(req, newListByWebAppSitePager)
		server.PagerResponderInjectNextLinks(newListByWebAppSitePager, req, func(page *armmigrationdiscovery.TomcatWebApplicationsControllerClientListByWebAppSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWebAppSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByWebAppSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWebAppSitePager) {
		t.newListByWebAppSitePager.remove(req)
	}
	return resp, nil
}

func (t *TomcatWebApplicationsControllerServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tomcatWebApplications/(?P<webApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[any](req)
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
	webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
	if err != nil {
		return nil, err
	}
	webApplicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webApplicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Update(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, webApplicationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TomcatWebApplications, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to TomcatWebApplicationsControllerServerTransport
var tomcatWebApplicationsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
