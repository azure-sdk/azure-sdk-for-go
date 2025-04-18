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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AccessReviewHistoryDefinitionsServer is a fake server for instances of the armauthorization.AccessReviewHistoryDefinitionsClient type.
type AccessReviewHistoryDefinitionsServer struct {
	// GetByID is the fake for method AccessReviewHistoryDefinitionsClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK
	GetByID func(ctx context.Context, historyDefinitionID string, options *armauthorization.AccessReviewHistoryDefinitionsClientGetByIDOptions) (resp azfake.Responder[armauthorization.AccessReviewHistoryDefinitionsClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AccessReviewHistoryDefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armauthorization.AccessReviewHistoryDefinitionsClientListOptions) (resp azfake.PagerResponder[armauthorization.AccessReviewHistoryDefinitionsClientListResponse])
}

// NewAccessReviewHistoryDefinitionsServerTransport creates a new instance of AccessReviewHistoryDefinitionsServerTransport with the provided implementation.
// The returned AccessReviewHistoryDefinitionsServerTransport instance is connected to an instance of armauthorization.AccessReviewHistoryDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessReviewHistoryDefinitionsServerTransport(srv *AccessReviewHistoryDefinitionsServer) *AccessReviewHistoryDefinitionsServerTransport {
	return &AccessReviewHistoryDefinitionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.AccessReviewHistoryDefinitionsClientListResponse]](),
	}
}

// AccessReviewHistoryDefinitionsServerTransport connects instances of armauthorization.AccessReviewHistoryDefinitionsClient to instances of AccessReviewHistoryDefinitionsServer.
// Don't use this type directly, use NewAccessReviewHistoryDefinitionsServerTransport instead.
type AccessReviewHistoryDefinitionsServerTransport struct {
	srv          *AccessReviewHistoryDefinitionsServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.AccessReviewHistoryDefinitionsClientListResponse]]
}

// Do implements the policy.Transporter interface for AccessReviewHistoryDefinitionsServerTransport.
func (a *AccessReviewHistoryDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AccessReviewHistoryDefinitionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if accessReviewHistoryDefinitionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = accessReviewHistoryDefinitionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AccessReviewHistoryDefinitionsClient.GetByID":
				res.resp, res.err = a.dispatchGetByID(req)
			case "AccessReviewHistoryDefinitionsClient.NewListPager":
				res.resp, res.err = a.dispatchNewListPager(req)
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

func (a *AccessReviewHistoryDefinitionsServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if a.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewHistoryDefinitions/(?P<historyDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	historyDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("historyDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetByID(req.Context(), historyDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewHistoryDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessReviewHistoryDefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewHistoryDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.AccessReviewHistoryDefinitionsClientListOptions
		if filterParam != nil {
			options = &armauthorization.AccessReviewHistoryDefinitionsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := a.srv.NewListPager(options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.AccessReviewHistoryDefinitionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AccessReviewHistoryDefinitionsServerTransport
var accessReviewHistoryDefinitionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
