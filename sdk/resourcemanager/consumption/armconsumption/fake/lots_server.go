//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
	"net/http"
	"net/url"
	"regexp"
)

// LotsServer is a fake server for instances of the armconsumption.LotsClient type.
type LotsServer struct {
	// NewListByBillingAccountPager is the fake for method LotsClient.NewListByBillingAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingAccountPager func(billingAccountID string, options *armconsumption.LotsClientListByBillingAccountOptions) (resp azfake.PagerResponder[armconsumption.LotsClientListByBillingAccountResponse])

	// NewListByBillingProfilePager is the fake for method LotsClient.NewListByBillingProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingProfilePager func(billingAccountID string, billingProfileID string, options *armconsumption.LotsClientListByBillingProfileOptions) (resp azfake.PagerResponder[armconsumption.LotsClientListByBillingProfileResponse])

	// NewListByCustomerPager is the fake for method LotsClient.NewListByCustomerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCustomerPager func(billingAccountID string, customerID string, options *armconsumption.LotsClientListByCustomerOptions) (resp azfake.PagerResponder[armconsumption.LotsClientListByCustomerResponse])
}

// NewLotsServerTransport creates a new instance of LotsServerTransport with the provided implementation.
// The returned LotsServerTransport instance is connected to an instance of armconsumption.LotsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLotsServerTransport(srv *LotsServer) *LotsServerTransport {
	return &LotsServerTransport{
		srv:                          srv,
		newListByBillingAccountPager: newTracker[azfake.PagerResponder[armconsumption.LotsClientListByBillingAccountResponse]](),
		newListByBillingProfilePager: newTracker[azfake.PagerResponder[armconsumption.LotsClientListByBillingProfileResponse]](),
		newListByCustomerPager:       newTracker[azfake.PagerResponder[armconsumption.LotsClientListByCustomerResponse]](),
	}
}

// LotsServerTransport connects instances of armconsumption.LotsClient to instances of LotsServer.
// Don't use this type directly, use NewLotsServerTransport instead.
type LotsServerTransport struct {
	srv                          *LotsServer
	newListByBillingAccountPager *tracker[azfake.PagerResponder[armconsumption.LotsClientListByBillingAccountResponse]]
	newListByBillingProfilePager *tracker[azfake.PagerResponder[armconsumption.LotsClientListByBillingProfileResponse]]
	newListByCustomerPager       *tracker[azfake.PagerResponder[armconsumption.LotsClientListByCustomerResponse]]
}

// Do implements the policy.Transporter interface for LotsServerTransport.
func (l *LotsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LotsClient.NewListByBillingAccountPager":
		resp, err = l.dispatchNewListByBillingAccountPager(req)
	case "LotsClient.NewListByBillingProfilePager":
		resp, err = l.dispatchNewListByBillingProfilePager(req)
	case "LotsClient.NewListByCustomerPager":
		resp, err = l.dispatchNewListByCustomerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LotsServerTransport) dispatchNewListByBillingAccountPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByBillingAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingAccountPager not implemented")}
	}
	newListByBillingAccountPager := l.newListByBillingAccountPager.get(req)
	if newListByBillingAccountPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/lots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armconsumption.LotsClientListByBillingAccountOptions
		if filterParam != nil {
			options = &armconsumption.LotsClientListByBillingAccountOptions{
				Filter: filterParam,
			}
		}
		resp := l.srv.NewListByBillingAccountPager(billingAccountIDParam, options)
		newListByBillingAccountPager = &resp
		l.newListByBillingAccountPager.add(req, newListByBillingAccountPager)
		server.PagerResponderInjectNextLinks(newListByBillingAccountPager, req, func(page *armconsumption.LotsClientListByBillingAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByBillingAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingAccountPager) {
		l.newListByBillingAccountPager.remove(req)
	}
	return resp, nil
}

func (l *LotsServerTransport) dispatchNewListByBillingProfilePager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByBillingProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingProfilePager not implemented")}
	}
	newListByBillingProfilePager := l.newListByBillingProfilePager.get(req)
	if newListByBillingProfilePager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/lots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
		if err != nil {
			return nil, err
		}
		billingProfileIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileId")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByBillingProfilePager(billingAccountIDParam, billingProfileIDParam, nil)
		newListByBillingProfilePager = &resp
		l.newListByBillingProfilePager.add(req, newListByBillingProfilePager)
		server.PagerResponderInjectNextLinks(newListByBillingProfilePager, req, func(page *armconsumption.LotsClientListByBillingProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByBillingProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingProfilePager) {
		l.newListByBillingProfilePager.remove(req)
	}
	return resp, nil
}

func (l *LotsServerTransport) dispatchNewListByCustomerPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByCustomerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCustomerPager not implemented")}
	}
	newListByCustomerPager := l.newListByCustomerPager.get(req)
	if newListByCustomerPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customers/(?P<customerId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/lots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
		if err != nil {
			return nil, err
		}
		customerIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armconsumption.LotsClientListByCustomerOptions
		if filterParam != nil {
			options = &armconsumption.LotsClientListByCustomerOptions{
				Filter: filterParam,
			}
		}
		resp := l.srv.NewListByCustomerPager(billingAccountIDParam, customerIDParam, options)
		newListByCustomerPager = &resp
		l.newListByCustomerPager.add(req, newListByCustomerPager)
		server.PagerResponderInjectNextLinks(newListByCustomerPager, req, func(page *armconsumption.LotsClientListByCustomerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCustomerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByCustomerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCustomerPager) {
		l.newListByCustomerPager.remove(req)
	}
	return resp, nil
}
