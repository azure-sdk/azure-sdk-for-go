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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AccessReviewHistoryDefinitionServer is a fake server for instances of the armauthorization.AccessReviewHistoryDefinitionClient type.
type AccessReviewHistoryDefinitionServer struct {
	// Create is the fake for method AccessReviewHistoryDefinitionClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, historyDefinitionID string, properties armauthorization.AccessReviewHistoryDefinitionProperties, options *armauthorization.AccessReviewHistoryDefinitionClientCreateOptions) (resp azfake.Responder[armauthorization.AccessReviewHistoryDefinitionClientCreateResponse], errResp azfake.ErrorResponder)

	// DeleteByID is the fake for method AccessReviewHistoryDefinitionClient.DeleteByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteByID func(ctx context.Context, historyDefinitionID string, options *armauthorization.AccessReviewHistoryDefinitionClientDeleteByIDOptions) (resp azfake.Responder[armauthorization.AccessReviewHistoryDefinitionClientDeleteByIDResponse], errResp azfake.ErrorResponder)
}

// NewAccessReviewHistoryDefinitionServerTransport creates a new instance of AccessReviewHistoryDefinitionServerTransport with the provided implementation.
// The returned AccessReviewHistoryDefinitionServerTransport instance is connected to an instance of armauthorization.AccessReviewHistoryDefinitionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessReviewHistoryDefinitionServerTransport(srv *AccessReviewHistoryDefinitionServer) *AccessReviewHistoryDefinitionServerTransport {
	return &AccessReviewHistoryDefinitionServerTransport{srv: srv}
}

// AccessReviewHistoryDefinitionServerTransport connects instances of armauthorization.AccessReviewHistoryDefinitionClient to instances of AccessReviewHistoryDefinitionServer.
// Don't use this type directly, use NewAccessReviewHistoryDefinitionServerTransport instead.
type AccessReviewHistoryDefinitionServerTransport struct {
	srv *AccessReviewHistoryDefinitionServer
}

// Do implements the policy.Transporter interface for AccessReviewHistoryDefinitionServerTransport.
func (a *AccessReviewHistoryDefinitionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AccessReviewHistoryDefinitionServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if accessReviewHistoryDefinitionServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = accessReviewHistoryDefinitionServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AccessReviewHistoryDefinitionClient.Create":
				res.resp, res.err = a.dispatchCreate(req)
			case "AccessReviewHistoryDefinitionClient.DeleteByID":
				res.resp, res.err = a.dispatchDeleteByID(req)
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

func (a *AccessReviewHistoryDefinitionServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if a.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewHistoryDefinitions/(?P<historyDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.AccessReviewHistoryDefinitionProperties](req)
	if err != nil {
		return nil, err
	}
	historyDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("historyDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Create(req.Context(), historyDefinitionIDParam, body, nil)
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

func (a *AccessReviewHistoryDefinitionServerTransport) dispatchDeleteByID(req *http.Request) (*http.Response, error) {
	if a.srv.DeleteByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteByID not implemented")}
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
	respr, errRespr := a.srv.DeleteByID(req.Context(), historyDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AccessReviewHistoryDefinitionServerTransport
var accessReviewHistoryDefinitionServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
