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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// WorkloadsServer is a fake server for instances of the armworkloadssapvirtualinstance.WorkloadsClient type.
type WorkloadsServer struct {
	// SAPAvailabilityZoneDetails is the fake for method WorkloadsClient.SAPAvailabilityZoneDetails
	// HTTP status codes to indicate success: http.StatusOK
	SAPAvailabilityZoneDetails func(ctx context.Context, location string, options *armworkloadssapvirtualinstance.WorkloadsClientSAPAvailabilityZoneDetailsOptions) (resp azfake.Responder[armworkloadssapvirtualinstance.WorkloadsClientSAPAvailabilityZoneDetailsResponse], errResp azfake.ErrorResponder)

	// SAPDiskConfigurations is the fake for method WorkloadsClient.SAPDiskConfigurations
	// HTTP status codes to indicate success: http.StatusOK
	SAPDiskConfigurations func(ctx context.Context, location string, options *armworkloadssapvirtualinstance.WorkloadsClientSAPDiskConfigurationsOptions) (resp azfake.Responder[armworkloadssapvirtualinstance.WorkloadsClientSAPDiskConfigurationsResponse], errResp azfake.ErrorResponder)

	// SAPSizingRecommendations is the fake for method WorkloadsClient.SAPSizingRecommendations
	// HTTP status codes to indicate success: http.StatusOK
	SAPSizingRecommendations func(ctx context.Context, location string, options *armworkloadssapvirtualinstance.WorkloadsClientSAPSizingRecommendationsOptions) (resp azfake.Responder[armworkloadssapvirtualinstance.WorkloadsClientSAPSizingRecommendationsResponse], errResp azfake.ErrorResponder)

	// SAPSupportedSKU is the fake for method WorkloadsClient.SAPSupportedSKU
	// HTTP status codes to indicate success: http.StatusOK
	SAPSupportedSKU func(ctx context.Context, location string, options *armworkloadssapvirtualinstance.WorkloadsClientSAPSupportedSKUOptions) (resp azfake.Responder[armworkloadssapvirtualinstance.WorkloadsClientSAPSupportedSKUResponse], errResp azfake.ErrorResponder)
}

// NewWorkloadsServerTransport creates a new instance of WorkloadsServerTransport with the provided implementation.
// The returned WorkloadsServerTransport instance is connected to an instance of armworkloadssapvirtualinstance.WorkloadsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkloadsServerTransport(srv *WorkloadsServer) *WorkloadsServerTransport {
	return &WorkloadsServerTransport{srv: srv}
}

// WorkloadsServerTransport connects instances of armworkloadssapvirtualinstance.WorkloadsClient to instances of WorkloadsServer.
// Don't use this type directly, use NewWorkloadsServerTransport instead.
type WorkloadsServerTransport struct {
	srv *WorkloadsServer
}

// Do implements the policy.Transporter interface for WorkloadsServerTransport.
func (w *WorkloadsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkloadsClient.SAPAvailabilityZoneDetails":
		resp, err = w.dispatchSAPAvailabilityZoneDetails(req)
	case "WorkloadsClient.SAPDiskConfigurations":
		resp, err = w.dispatchSAPDiskConfigurations(req)
	case "WorkloadsClient.SAPSizingRecommendations":
		resp, err = w.dispatchSAPSizingRecommendations(req)
	case "WorkloadsClient.SAPSupportedSKU":
		resp, err = w.dispatchSAPSupportedSKU(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkloadsServerTransport) dispatchSAPAvailabilityZoneDetails(req *http.Request) (*http.Response, error) {
	if w.srv.SAPAvailabilityZoneDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPAvailabilityZoneDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getAvailabilityZoneDetails`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.SAPAvailabilityZoneDetailsRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloadssapvirtualinstance.WorkloadsClientSAPAvailabilityZoneDetailsOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloadssapvirtualinstance.WorkloadsClientSAPAvailabilityZoneDetailsOptions{
			SAPAvailabilityZoneDetails: &body,
		}
	}
	respr, errRespr := w.srv.SAPAvailabilityZoneDetails(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPAvailabilityZoneDetailsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkloadsServerTransport) dispatchSAPDiskConfigurations(req *http.Request) (*http.Response, error) {
	if w.srv.SAPDiskConfigurations == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPDiskConfigurations not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getDiskConfigurations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.SAPDiskConfigurationsRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloadssapvirtualinstance.WorkloadsClientSAPDiskConfigurationsOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloadssapvirtualinstance.WorkloadsClientSAPDiskConfigurationsOptions{
			SAPDiskConfigurations: &body,
		}
	}
	respr, errRespr := w.srv.SAPDiskConfigurations(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPDiskConfigurationsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkloadsServerTransport) dispatchSAPSizingRecommendations(req *http.Request) (*http.Response, error) {
	if w.srv.SAPSizingRecommendations == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPSizingRecommendations not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getSizingRecommendations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.SAPSizingRecommendationRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloadssapvirtualinstance.WorkloadsClientSAPSizingRecommendationsOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloadssapvirtualinstance.WorkloadsClientSAPSizingRecommendationsOptions{
			SAPSizingRecommendation: &body,
		}
	}
	respr, errRespr := w.srv.SAPSizingRecommendations(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPSizingRecommendationResultClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkloadsServerTransport) dispatchSAPSupportedSKU(req *http.Request) (*http.Response, error) {
	if w.srv.SAPSupportedSKU == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPSupportedSKU not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getSapSupportedSku`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.SAPSupportedSKUsRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloadssapvirtualinstance.WorkloadsClientSAPSupportedSKUOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloadssapvirtualinstance.WorkloadsClientSAPSupportedSKUOptions{
			SAPSupportedSKU: &body,
		}
	}
	respr, errRespr := w.srv.SAPSupportedSKU(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPSupportedResourceSKUsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
