//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
	"net/http"
	"net/url"
	"regexp"
)

// ScheduledActionsServer is a fake server for instances of the armcomputeschedule.ScheduledActionsClient type.
type ScheduledActionsServer struct {
	// VirtualMachinesCancelOperations is the fake for method ScheduledActionsClient.VirtualMachinesCancelOperations
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesCancelOperations func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.CancelOperationsRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesCancelOperationsOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesCancelOperationsResponse], errResp azfake.ErrorResponder)

	// VirtualMachinesExecuteDeallocate is the fake for method ScheduledActionsClient.VirtualMachinesExecuteDeallocate
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesExecuteDeallocate func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.ExecuteDeallocateRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteDeallocateOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteDeallocateResponse], errResp azfake.ErrorResponder)

	// VirtualMachinesExecuteHibernate is the fake for method ScheduledActionsClient.VirtualMachinesExecuteHibernate
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesExecuteHibernate func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.ExecuteHibernateRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteHibernateOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteHibernateResponse], errResp azfake.ErrorResponder)

	// VirtualMachinesExecuteStart is the fake for method ScheduledActionsClient.VirtualMachinesExecuteStart
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesExecuteStart func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.ExecuteStartRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteStartOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteStartResponse], errResp azfake.ErrorResponder)

	// VirtualMachinesGetOperationStatus is the fake for method ScheduledActionsClient.VirtualMachinesGetOperationStatus
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesGetOperationStatus func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.GetOperationStatusRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesGetOperationStatusOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesGetOperationStatusResponse], errResp azfake.ErrorResponder)

	// VirtualMachinesSubmitDeallocate is the fake for method ScheduledActionsClient.VirtualMachinesSubmitDeallocate
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesSubmitDeallocate func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.SubmitDeallocateRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitDeallocateOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitDeallocateResponse], errResp azfake.ErrorResponder)

	// VirtualMachinesSubmitHibernate is the fake for method ScheduledActionsClient.VirtualMachinesSubmitHibernate
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesSubmitHibernate func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.SubmitHibernateRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitHibernateOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitHibernateResponse], errResp azfake.ErrorResponder)

	// VirtualMachinesSubmitStart is the fake for method ScheduledActionsClient.VirtualMachinesSubmitStart
	// HTTP status codes to indicate success: http.StatusOK
	VirtualMachinesSubmitStart func(ctx context.Context, locationparameter string, requestBody armcomputeschedule.SubmitStartRequest, options *armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitStartOptions) (resp azfake.Responder[armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitStartResponse], errResp azfake.ErrorResponder)
}

// NewScheduledActionsServerTransport creates a new instance of ScheduledActionsServerTransport with the provided implementation.
// The returned ScheduledActionsServerTransport instance is connected to an instance of armcomputeschedule.ScheduledActionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScheduledActionsServerTransport(srv *ScheduledActionsServer) *ScheduledActionsServerTransport {
	return &ScheduledActionsServerTransport{srv: srv}
}

// ScheduledActionsServerTransport connects instances of armcomputeschedule.ScheduledActionsClient to instances of ScheduledActionsServer.
// Don't use this type directly, use NewScheduledActionsServerTransport instead.
type ScheduledActionsServerTransport struct {
	srv *ScheduledActionsServer
}

// Do implements the policy.Transporter interface for ScheduledActionsServerTransport.
func (s *ScheduledActionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ScheduledActionsClient.VirtualMachinesCancelOperations":
		resp, err = s.dispatchVirtualMachinesCancelOperations(req)
	case "ScheduledActionsClient.VirtualMachinesExecuteDeallocate":
		resp, err = s.dispatchVirtualMachinesExecuteDeallocate(req)
	case "ScheduledActionsClient.VirtualMachinesExecuteHibernate":
		resp, err = s.dispatchVirtualMachinesExecuteHibernate(req)
	case "ScheduledActionsClient.VirtualMachinesExecuteStart":
		resp, err = s.dispatchVirtualMachinesExecuteStart(req)
	case "ScheduledActionsClient.VirtualMachinesGetOperationStatus":
		resp, err = s.dispatchVirtualMachinesGetOperationStatus(req)
	case "ScheduledActionsClient.VirtualMachinesSubmitDeallocate":
		resp, err = s.dispatchVirtualMachinesSubmitDeallocate(req)
	case "ScheduledActionsClient.VirtualMachinesSubmitHibernate":
		resp, err = s.dispatchVirtualMachinesSubmitHibernate(req)
	case "ScheduledActionsClient.VirtualMachinesSubmitStart":
		resp, err = s.dispatchVirtualMachinesSubmitStart(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesCancelOperations(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesCancelOperations == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesCancelOperations not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesCancelOperations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.CancelOperationsRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesCancelOperations(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CancelOperationsResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesExecuteDeallocate(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesExecuteDeallocate == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesExecuteDeallocate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesExecuteDeallocate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.ExecuteDeallocateRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesExecuteDeallocate(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeallocateResourceOperationResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesExecuteHibernate(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesExecuteHibernate == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesExecuteHibernate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesExecuteHibernate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.ExecuteHibernateRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesExecuteHibernate(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HibernateResourceOperationResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesExecuteStart(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesExecuteStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesExecuteStart not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesExecuteStart`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.ExecuteStartRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesExecuteStart(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StartResourceOperationResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesGetOperationStatus(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesGetOperationStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesGetOperationStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesGetOperationStatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.GetOperationStatusRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesGetOperationStatus(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GetOperationStatusResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesSubmitDeallocate(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesSubmitDeallocate == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesSubmitDeallocate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesSubmitDeallocate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.SubmitDeallocateRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesSubmitDeallocate(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeallocateResourceOperationResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesSubmitHibernate(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesSubmitHibernate == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesSubmitHibernate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesSubmitHibernate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.SubmitHibernateRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesSubmitHibernate(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HibernateResourceOperationResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScheduledActionsServerTransport) dispatchVirtualMachinesSubmitStart(req *http.Request) (*http.Response, error) {
	if s.srv.VirtualMachinesSubmitStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method VirtualMachinesSubmitStart not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ComputeSchedule/locations/(?P<locationparameter>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachinesSubmitStart`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcomputeschedule.SubmitStartRequest](req)
	if err != nil {
		return nil, err
	}
	locationparameterParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationparameter")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.VirtualMachinesSubmitStart(req.Context(), locationparameterParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StartResourceOperationResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
