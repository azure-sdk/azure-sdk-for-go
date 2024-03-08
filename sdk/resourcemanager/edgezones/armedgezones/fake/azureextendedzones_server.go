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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgezones/armedgezones"
	"net/http"
	"net/url"
	"regexp"
)

// AzureExtendedZonesServer is a fake server for instances of the armedgezones.AzureExtendedZonesClient type.
type AzureExtendedZonesServer struct {
	// Get is the fake for method AzureExtendedZonesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, azureExtendedZoneName string, options *armedgezones.AzureExtendedZonesClientGetOptions) (resp azfake.Responder[armedgezones.AzureExtendedZonesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySubscriptionPager is the fake for method AzureExtendedZonesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armedgezones.AzureExtendedZonesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armedgezones.AzureExtendedZonesClientListBySubscriptionResponse])

	// Register is the fake for method AzureExtendedZonesClient.Register
	// HTTP status codes to indicate success: http.StatusOK
	Register func(ctx context.Context, azureExtendedZoneName string, options *armedgezones.AzureExtendedZonesClientRegisterOptions) (resp azfake.Responder[armedgezones.AzureExtendedZonesClientRegisterResponse], errResp azfake.ErrorResponder)

	// Unregister is the fake for method AzureExtendedZonesClient.Unregister
	// HTTP status codes to indicate success: http.StatusOK
	Unregister func(ctx context.Context, azureExtendedZoneName string, options *armedgezones.AzureExtendedZonesClientUnregisterOptions) (resp azfake.Responder[armedgezones.AzureExtendedZonesClientUnregisterResponse], errResp azfake.ErrorResponder)
}

// NewAzureExtendedZonesServerTransport creates a new instance of AzureExtendedZonesServerTransport with the provided implementation.
// The returned AzureExtendedZonesServerTransport instance is connected to an instance of armedgezones.AzureExtendedZonesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureExtendedZonesServerTransport(srv *AzureExtendedZonesServer) *AzureExtendedZonesServerTransport {
	return &AzureExtendedZonesServerTransport{
		srv:                        srv,
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armedgezones.AzureExtendedZonesClientListBySubscriptionResponse]](),
	}
}

// AzureExtendedZonesServerTransport connects instances of armedgezones.AzureExtendedZonesClient to instances of AzureExtendedZonesServer.
// Don't use this type directly, use NewAzureExtendedZonesServerTransport instead.
type AzureExtendedZonesServerTransport struct {
	srv                        *AzureExtendedZonesServer
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armedgezones.AzureExtendedZonesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for AzureExtendedZonesServerTransport.
func (a *AzureExtendedZonesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureExtendedZonesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureExtendedZonesClient.NewListBySubscriptionPager":
		resp, err = a.dispatchNewListBySubscriptionPager(req)
	case "AzureExtendedZonesClient.Register":
		resp, err = a.dispatchRegister(req)
	case "AzureExtendedZonesClient.Unregister":
		resp, err = a.dispatchUnregister(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureExtendedZonesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeZones/azureExtendedZones/(?P<azureExtendedZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	azureExtendedZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureExtendedZoneName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), azureExtendedZoneNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureExtendedZone, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureExtendedZonesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := a.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeZones/azureExtendedZones`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		a.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armedgezones.AzureExtendedZonesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		a.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (a *AzureExtendedZonesServerTransport) dispatchRegister(req *http.Request) (*http.Response, error) {
	if a.srv.Register == nil {
		return nil, &nonRetriableError{errors.New("fake for method Register not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeZones/azureExtendedZones/(?P<azureExtendedZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/register`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	azureExtendedZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureExtendedZoneName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Register(req.Context(), azureExtendedZoneNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureExtendedZone, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureExtendedZonesServerTransport) dispatchUnregister(req *http.Request) (*http.Response, error) {
	if a.srv.Unregister == nil {
		return nil, &nonRetriableError{errors.New("fake for method Unregister not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeZones/azureExtendedZones/(?P<azureExtendedZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unregister`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	azureExtendedZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureExtendedZoneName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Unregister(req.Context(), azureExtendedZoneNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureExtendedZone, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
