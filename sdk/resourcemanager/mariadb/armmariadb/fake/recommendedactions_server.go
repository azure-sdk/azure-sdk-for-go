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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"net/http"
	"net/url"
	"regexp"
)

// RecommendedActionsServer is a fake server for instances of the armmariadb.RecommendedActionsClient type.
type RecommendedActionsServer struct {
	// Get is the fake for method RecommendedActionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, advisorName string, recommendedActionName string, options *armmariadb.RecommendedActionsClientGetOptions) (resp azfake.Responder[armmariadb.RecommendedActionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method RecommendedActionsClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, advisorName string, options *armmariadb.RecommendedActionsClientListByServerOptions) (resp azfake.PagerResponder[armmariadb.RecommendedActionsClientListByServerResponse])
}

// NewRecommendedActionsServerTransport creates a new instance of RecommendedActionsServerTransport with the provided implementation.
// The returned RecommendedActionsServerTransport instance is connected to an instance of armmariadb.RecommendedActionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRecommendedActionsServerTransport(srv *RecommendedActionsServer) *RecommendedActionsServerTransport {
	return &RecommendedActionsServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armmariadb.RecommendedActionsClientListByServerResponse]](),
	}
}

// RecommendedActionsServerTransport connects instances of armmariadb.RecommendedActionsClient to instances of RecommendedActionsServer.
// Don't use this type directly, use NewRecommendedActionsServerTransport instead.
type RecommendedActionsServerTransport struct {
	srv                  *RecommendedActionsServer
	newListByServerPager *tracker[azfake.PagerResponder[armmariadb.RecommendedActionsClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for RecommendedActionsServerTransport.
func (r *RecommendedActionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RecommendedActionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if recommendedActionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = recommendedActionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RecommendedActionsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RecommendedActionsClient.NewListByServerPager":
				res.resp, res.err = r.dispatchNewListByServerPager(req)
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

func (r *RecommendedActionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMariaDB/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/advisors/(?P<advisorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recommendedActions/(?P<recommendedActionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	advisorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("advisorName")])
	if err != nil {
		return nil, err
	}
	recommendedActionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recommendedActionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, advisorNameParam, recommendedActionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RecommendationAction, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecommendedActionsServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := r.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMariaDB/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/advisors/(?P<advisorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recommendedActions`
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
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		advisorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("advisorName")])
		if err != nil {
			return nil, err
		}
		sessionIDUnescaped, err := url.QueryUnescape(qp.Get("sessionId"))
		if err != nil {
			return nil, err
		}
		sessionIDParam := getOptional(sessionIDUnescaped)
		var options *armmariadb.RecommendedActionsClientListByServerOptions
		if sessionIDParam != nil {
			options = &armmariadb.RecommendedActionsClientListByServerOptions{
				SessionID: sessionIDParam,
			}
		}
		resp := r.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, advisorNameParam, options)
		newListByServerPager = &resp
		r.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armmariadb.RecommendedActionsClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		r.newListByServerPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to RecommendedActionsServerTransport
var recommendedActionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
