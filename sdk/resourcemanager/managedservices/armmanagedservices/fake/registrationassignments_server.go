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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// RegistrationAssignmentsServer is a fake server for instances of the armmanagedservices.RegistrationAssignmentsClient type.
type RegistrationAssignmentsServer struct {
	// BeginCreateOrUpdate is the fake for method RegistrationAssignmentsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, scope string, registrationAssignmentID string, requestBody armmanagedservices.RegistrationAssignment, options *armmanagedservices.RegistrationAssignmentsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmanagedservices.RegistrationAssignmentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RegistrationAssignmentsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, scope string, registrationAssignmentID string, options *armmanagedservices.RegistrationAssignmentsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmanagedservices.RegistrationAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RegistrationAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, registrationAssignmentID string, options *armmanagedservices.RegistrationAssignmentsClientGetOptions) (resp azfake.Responder[armmanagedservices.RegistrationAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RegistrationAssignmentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, options *armmanagedservices.RegistrationAssignmentsClientListOptions) (resp azfake.PagerResponder[armmanagedservices.RegistrationAssignmentsClientListResponse])
}

// NewRegistrationAssignmentsServerTransport creates a new instance of RegistrationAssignmentsServerTransport with the provided implementation.
// The returned RegistrationAssignmentsServerTransport instance is connected to an instance of armmanagedservices.RegistrationAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegistrationAssignmentsServerTransport(srv *RegistrationAssignmentsServer) *RegistrationAssignmentsServerTransport {
	return &RegistrationAssignmentsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmanagedservices.RegistrationAssignmentsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armmanagedservices.RegistrationAssignmentsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armmanagedservices.RegistrationAssignmentsClientListResponse]](),
	}
}

// RegistrationAssignmentsServerTransport connects instances of armmanagedservices.RegistrationAssignmentsClient to instances of RegistrationAssignmentsServer.
// Don't use this type directly, use NewRegistrationAssignmentsServerTransport instead.
type RegistrationAssignmentsServerTransport struct {
	srv                 *RegistrationAssignmentsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmanagedservices.RegistrationAssignmentsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armmanagedservices.RegistrationAssignmentsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armmanagedservices.RegistrationAssignmentsClientListResponse]]
}

// Do implements the policy.Transporter interface for RegistrationAssignmentsServerTransport.
func (r *RegistrationAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RegistrationAssignmentsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if registrationAssignmentsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = registrationAssignmentsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RegistrationAssignmentsClient.BeginCreateOrUpdate":
				res.resp, res.err = r.dispatchBeginCreateOrUpdate(req)
			case "RegistrationAssignmentsClient.BeginDelete":
				res.resp, res.err = r.dispatchBeginDelete(req)
			case "RegistrationAssignmentsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RegistrationAssignmentsClient.NewListPager":
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

func (r *RegistrationAssignmentsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationAssignments/(?P<registrationAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagedservices.RegistrationAssignment](req)
		if err != nil {
			return nil, err
		}
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		registrationAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("registrationAssignmentId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), scopeParam, registrationAssignmentIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *RegistrationAssignmentsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationAssignments/(?P<registrationAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		registrationAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("registrationAssignmentId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), scopeParam, registrationAssignmentIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RegistrationAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationAssignments/(?P<registrationAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	registrationAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("registrationAssignmentId")])
	if err != nil {
		return nil, err
	}
	expandRegistrationDefinitionUnescaped, err := url.QueryUnescape(qp.Get("$expandRegistrationDefinition"))
	if err != nil {
		return nil, err
	}
	expandRegistrationDefinitionParam, err := parseOptional(expandRegistrationDefinitionUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armmanagedservices.RegistrationAssignmentsClientGetOptions
	if expandRegistrationDefinitionParam != nil {
		options = &armmanagedservices.RegistrationAssignmentsClientGetOptions{
			ExpandRegistrationDefinition: expandRegistrationDefinitionParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), scopeParam, registrationAssignmentIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RegistrationAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RegistrationAssignmentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		expandRegistrationDefinitionUnescaped, err := url.QueryUnescape(qp.Get("$expandRegistrationDefinition"))
		if err != nil {
			return nil, err
		}
		expandRegistrationDefinitionParam, err := parseOptional(expandRegistrationDefinitionUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armmanagedservices.RegistrationAssignmentsClientListOptions
		if expandRegistrationDefinitionParam != nil || filterParam != nil {
			options = &armmanagedservices.RegistrationAssignmentsClientListOptions{
				ExpandRegistrationDefinition: expandRegistrationDefinitionParam,
				Filter:                       filterParam,
			}
		}
		resp := r.srv.NewListPager(scopeParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmanagedservices.RegistrationAssignmentsClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to RegistrationAssignmentsServerTransport
var registrationAssignmentsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
