// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
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
	"strconv"
)

// WebServersControllerServer is a fake server for instances of the armmigrate.WebServersControllerClient type.
type WebServersControllerServer struct {
	// NewListByWebAppSitePager is the fake for method WebServersControllerClient.NewListByWebAppSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWebAppSitePager func(resourceGroupName string, siteName string, webAppSiteName string, options *armmigrate.WebServersControllerClientListByWebAppSiteOptions) (resp azfake.PagerResponder[armmigrate.WebServersControllerClientListByWebAppSiteResponse])
}

// NewWebServersControllerServerTransport creates a new instance of WebServersControllerServerTransport with the provided implementation.
// The returned WebServersControllerServerTransport instance is connected to an instance of armmigrate.WebServersControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWebServersControllerServerTransport(srv *WebServersControllerServer) *WebServersControllerServerTransport {
	return &WebServersControllerServerTransport{
		srv:                      srv,
		newListByWebAppSitePager: newTracker[azfake.PagerResponder[armmigrate.WebServersControllerClientListByWebAppSiteResponse]](),
	}
}

// WebServersControllerServerTransport connects instances of armmigrate.WebServersControllerClient to instances of WebServersControllerServer.
// Don't use this type directly, use NewWebServersControllerServerTransport instead.
type WebServersControllerServerTransport struct {
	srv                      *WebServersControllerServer
	newListByWebAppSitePager *tracker[azfake.PagerResponder[armmigrate.WebServersControllerClientListByWebAppSiteResponse]]
}

// Do implements the policy.Transporter interface for WebServersControllerServerTransport.
func (w *WebServersControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WebServersControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if webServersControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = webServersControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WebServersControllerClient.NewListByWebAppSitePager":
				res.resp, res.err = w.dispatchNewListByWebAppSitePager(req)
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

func (w *WebServersControllerServerTransport) dispatchNewListByWebAppSitePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByWebAppSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWebAppSitePager not implemented")}
	}
	newListByWebAppSitePager := w.newListByWebAppSitePager.get(req)
	if newListByWebAppSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webServers`
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
		var options *armmigrate.WebServersControllerClientListByWebAppSiteOptions
		if filterParam != nil || topParam != nil || continuationTokenParam != nil || totalRecordCountParam != nil {
			options = &armmigrate.WebServersControllerClientListByWebAppSiteOptions{
				Filter:            filterParam,
				Top:               topParam,
				ContinuationToken: continuationTokenParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := w.srv.NewListByWebAppSitePager(resourceGroupNameParam, siteNameParam, webAppSiteNameParam, options)
		newListByWebAppSitePager = &resp
		w.newListByWebAppSitePager.add(req, newListByWebAppSitePager)
		server.PagerResponderInjectNextLinks(newListByWebAppSitePager, req, func(page *armmigrate.WebServersControllerClientListByWebAppSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWebAppSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByWebAppSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWebAppSitePager) {
		w.newListByWebAppSitePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to WebServersControllerServerTransport
var webServersControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
