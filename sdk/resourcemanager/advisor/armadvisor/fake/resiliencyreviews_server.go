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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ResiliencyReviewsServer is a fake server for instances of the armadvisor.ResiliencyReviewsClient type.
type ResiliencyReviewsServer struct {
	// Get is the fake for method ResiliencyReviewsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, reviewID string, options *armadvisor.ResiliencyReviewsClientGetOptions) (resp azfake.Responder[armadvisor.ResiliencyReviewsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ResiliencyReviewsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armadvisor.ResiliencyReviewsClientListOptions) (resp azfake.PagerResponder[armadvisor.ResiliencyReviewsClientListResponse])
}

// NewResiliencyReviewsServerTransport creates a new instance of ResiliencyReviewsServerTransport with the provided implementation.
// The returned ResiliencyReviewsServerTransport instance is connected to an instance of armadvisor.ResiliencyReviewsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResiliencyReviewsServerTransport(srv *ResiliencyReviewsServer) *ResiliencyReviewsServerTransport {
	return &ResiliencyReviewsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armadvisor.ResiliencyReviewsClientListResponse]](),
	}
}

// ResiliencyReviewsServerTransport connects instances of armadvisor.ResiliencyReviewsClient to instances of ResiliencyReviewsServer.
// Don't use this type directly, use NewResiliencyReviewsServerTransport instead.
type ResiliencyReviewsServerTransport struct {
	srv          *ResiliencyReviewsServer
	newListPager *tracker[azfake.PagerResponder[armadvisor.ResiliencyReviewsClientListResponse]]
}

// Do implements the policy.Transporter interface for ResiliencyReviewsServerTransport.
func (r *ResiliencyReviewsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ResiliencyReviewsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if resiliencyReviewsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = resiliencyReviewsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ResiliencyReviewsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ResiliencyReviewsClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
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

func (r *ResiliencyReviewsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/resiliencyReviews/(?P<reviewId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reviewIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reviewId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), reviewIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResiliencyReview, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResiliencyReviewsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/resiliencyReviews`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armadvisor.ResiliencyReviewsClientListOptions
		if topParam != nil || skipParam != nil || filterParam != nil {
			options = &armadvisor.ResiliencyReviewsClientListOptions{
				Top:    topParam,
				Skip:   skipParam,
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListPager(options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armadvisor.ResiliencyReviewsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ResiliencyReviewsServerTransport
var resiliencyReviewsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
