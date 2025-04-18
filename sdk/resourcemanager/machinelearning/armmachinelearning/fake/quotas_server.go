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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
)

// QuotasServer is a fake server for instances of the armmachinelearning.QuotasClient type.
type QuotasServer struct {
	// NewListPager is the fake for method QuotasClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armmachinelearning.QuotasClientListOptions) (resp azfake.PagerResponder[armmachinelearning.QuotasClientListResponse])

	// Update is the fake for method QuotasClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, location string, parameters armmachinelearning.QuotaUpdateParameters, options *armmachinelearning.QuotasClientUpdateOptions) (resp azfake.Responder[armmachinelearning.QuotasClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewQuotasServerTransport creates a new instance of QuotasServerTransport with the provided implementation.
// The returned QuotasServerTransport instance is connected to an instance of armmachinelearning.QuotasClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQuotasServerTransport(srv *QuotasServer) *QuotasServerTransport {
	return &QuotasServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.QuotasClientListResponse]](),
	}
}

// QuotasServerTransport connects instances of armmachinelearning.QuotasClient to instances of QuotasServer.
// Don't use this type directly, use NewQuotasServerTransport instead.
type QuotasServerTransport struct {
	srv          *QuotasServer
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.QuotasClientListResponse]]
}

// Do implements the policy.Transporter interface for QuotasServerTransport.
func (q *QuotasServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return q.dispatchToMethodFake(req, method)
}

func (q *QuotasServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if quotasServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = quotasServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "QuotasClient.NewListPager":
				res.resp, res.err = q.dispatchNewListPager(req)
			case "QuotasClient.Update":
				res.resp, res.err = q.dispatchUpdate(req)
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

func (q *QuotasServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if q.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := q.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotas`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := q.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		q.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.QuotasClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		q.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		q.newListPager.remove(req)
	}
	return resp, nil
}

func (q *QuotasServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if q.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateQuotas`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.QuotaUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Update(req.Context(), locationParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UpdateWorkspaceQuotasResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to QuotasServerTransport
var quotasServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
