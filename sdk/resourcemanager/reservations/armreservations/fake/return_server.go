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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ReturnServer is a fake server for instances of the armreservations.ReturnClient type.
type ReturnServer struct {
	// BeginPost is the fake for method ReturnClient.BeginPost
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPost func(ctx context.Context, reservationOrderID string, body armreservations.RefundRequest, options *armreservations.ReturnClientBeginPostOptions) (resp azfake.PollerResponder[armreservations.ReturnClientPostResponse], errResp azfake.ErrorResponder)
}

// NewReturnServerTransport creates a new instance of ReturnServerTransport with the provided implementation.
// The returned ReturnServerTransport instance is connected to an instance of armreservations.ReturnClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReturnServerTransport(srv *ReturnServer) *ReturnServerTransport {
	return &ReturnServerTransport{
		srv:       srv,
		beginPost: newTracker[azfake.PollerResponder[armreservations.ReturnClientPostResponse]](),
	}
}

// ReturnServerTransport connects instances of armreservations.ReturnClient to instances of ReturnServer.
// Don't use this type directly, use NewReturnServerTransport instead.
type ReturnServerTransport struct {
	srv       *ReturnServer
	beginPost *tracker[azfake.PollerResponder[armreservations.ReturnClientPostResponse]]
}

// Do implements the policy.Transporter interface for ReturnServerTransport.
func (r *ReturnServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ReturnServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if returnServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = returnServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ReturnClient.BeginPost":
				res.resp, res.err = r.dispatchBeginPost(req)
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

func (r *ReturnServerTransport) dispatchBeginPost(req *http.Request) (*http.Response, error) {
	if r.srv.BeginPost == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPost not implemented")}
	}
	beginPost := r.beginPost.get(req)
	if beginPost == nil {
		const regexStr = `/providers/Microsoft\.Capacity/reservationOrders/(?P<reservationOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/return`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armreservations.RefundRequest](req)
		if err != nil {
			return nil, err
		}
		reservationOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginPost(req.Context(), reservationOrderIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPost = &respr
		r.beginPost.add(req, beginPost)
	}

	resp, err := server.PollerResponderNext(beginPost, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginPost.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPost) {
		r.beginPost.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ReturnServerTransport
var returnServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
