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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"net/http"
	"net/url"
	"regexp"
)

// LogFilesServer is a fake server for instances of the armmysql.LogFilesClient type.
type LogFilesServer struct {
	// NewListByServerPager is the fake for method LogFilesClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armmysql.LogFilesClientListByServerOptions) (resp azfake.PagerResponder[armmysql.LogFilesClientListByServerResponse])
}

// NewLogFilesServerTransport creates a new instance of LogFilesServerTransport with the provided implementation.
// The returned LogFilesServerTransport instance is connected to an instance of armmysql.LogFilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLogFilesServerTransport(srv *LogFilesServer) *LogFilesServerTransport {
	return &LogFilesServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armmysql.LogFilesClientListByServerResponse]](),
	}
}

// LogFilesServerTransport connects instances of armmysql.LogFilesClient to instances of LogFilesServer.
// Don't use this type directly, use NewLogFilesServerTransport instead.
type LogFilesServerTransport struct {
	srv                  *LogFilesServer
	newListByServerPager *tracker[azfake.PagerResponder[armmysql.LogFilesClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for LogFilesServerTransport.
func (l *LogFilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LogFilesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if logFilesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = logFilesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LogFilesClient.NewListByServerPager":
				res.resp, res.err = l.dispatchNewListByServerPager(req)
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

func (l *LogFilesServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := l.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/logFiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		l.newListByServerPager.add(req, newListByServerPager)
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		l.newListByServerPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to LogFilesServerTransport
var logFilesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
