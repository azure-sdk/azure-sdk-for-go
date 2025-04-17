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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// BenefitUtilizationSummariesServer is a fake server for instances of the armcostmanagement.BenefitUtilizationSummariesClient type.
type BenefitUtilizationSummariesServer struct {
	// NewListByBillingAccountIDPager is the fake for method BenefitUtilizationSummariesClient.NewListByBillingAccountIDPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingAccountIDPager func(billingAccountID string, options *armcostmanagement.BenefitUtilizationSummariesClientListByBillingAccountIDOptions) (resp azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListByBillingAccountIDResponse])

	// NewListByBillingProfileIDPager is the fake for method BenefitUtilizationSummariesClient.NewListByBillingProfileIDPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingProfileIDPager func(billingAccountID string, billingProfileID string, options *armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDOptions) (resp azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDResponse])

	// NewListBySavingsPlanIDPager is the fake for method BenefitUtilizationSummariesClient.NewListBySavingsPlanIDPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySavingsPlanIDPager func(savingsPlanOrderID string, savingsPlanID string, options *armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanIDOptions) (resp azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanIDResponse])

	// NewListBySavingsPlanOrderPager is the fake for method BenefitUtilizationSummariesClient.NewListBySavingsPlanOrderPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySavingsPlanOrderPager func(savingsPlanOrderID string, options *armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions) (resp azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse])
}

// NewBenefitUtilizationSummariesServerTransport creates a new instance of BenefitUtilizationSummariesServerTransport with the provided implementation.
// The returned BenefitUtilizationSummariesServerTransport instance is connected to an instance of armcostmanagement.BenefitUtilizationSummariesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBenefitUtilizationSummariesServerTransport(srv *BenefitUtilizationSummariesServer) *BenefitUtilizationSummariesServerTransport {
	return &BenefitUtilizationSummariesServerTransport{
		srv:                            srv,
		newListByBillingAccountIDPager: newTracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListByBillingAccountIDResponse]](),
		newListByBillingProfileIDPager: newTracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDResponse]](),
		newListBySavingsPlanIDPager:    newTracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanIDResponse]](),
		newListBySavingsPlanOrderPager: newTracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse]](),
	}
}

// BenefitUtilizationSummariesServerTransport connects instances of armcostmanagement.BenefitUtilizationSummariesClient to instances of BenefitUtilizationSummariesServer.
// Don't use this type directly, use NewBenefitUtilizationSummariesServerTransport instead.
type BenefitUtilizationSummariesServerTransport struct {
	srv                            *BenefitUtilizationSummariesServer
	newListByBillingAccountIDPager *tracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListByBillingAccountIDResponse]]
	newListByBillingProfileIDPager *tracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDResponse]]
	newListBySavingsPlanIDPager    *tracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanIDResponse]]
	newListBySavingsPlanOrderPager *tracker[azfake.PagerResponder[armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse]]
}

// Do implements the policy.Transporter interface for BenefitUtilizationSummariesServerTransport.
func (b *BenefitUtilizationSummariesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BenefitUtilizationSummariesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if benefitUtilizationSummariesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = benefitUtilizationSummariesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BenefitUtilizationSummariesClient.NewListByBillingAccountIDPager":
				res.resp, res.err = b.dispatchNewListByBillingAccountIDPager(req)
			case "BenefitUtilizationSummariesClient.NewListByBillingProfileIDPager":
				res.resp, res.err = b.dispatchNewListByBillingProfileIDPager(req)
			case "BenefitUtilizationSummariesClient.NewListBySavingsPlanIDPager":
				res.resp, res.err = b.dispatchNewListBySavingsPlanIDPager(req)
			case "BenefitUtilizationSummariesClient.NewListBySavingsPlanOrderPager":
				res.resp, res.err = b.dispatchNewListBySavingsPlanOrderPager(req)
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

func (b *BenefitUtilizationSummariesServerTransport) dispatchNewListByBillingAccountIDPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByBillingAccountIDPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingAccountIDPager not implemented")}
	}
	newListByBillingAccountIDPager := b.newListByBillingAccountIDPager.get(req)
	if newListByBillingAccountIDPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/benefitUtilizationSummaries`
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
		grainParameterUnescaped, err := url.QueryUnescape(qp.Get("grainParameter"))
		if err != nil {
			return nil, err
		}
		grainParameterParam := getOptional(armcostmanagement.GrainParameter(grainParameterUnescaped))
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armcostmanagement.BenefitUtilizationSummariesClientListByBillingAccountIDOptions
		if grainParameterParam != nil || filterParam != nil {
			options = &armcostmanagement.BenefitUtilizationSummariesClientListByBillingAccountIDOptions{
				GrainParameter: grainParameterParam,
				Filter:         filterParam,
			}
		}
		resp := b.srv.NewListByBillingAccountIDPager(billingAccountIDParam, options)
		newListByBillingAccountIDPager = &resp
		b.newListByBillingAccountIDPager.add(req, newListByBillingAccountIDPager)
		server.PagerResponderInjectNextLinks(newListByBillingAccountIDPager, req, func(page *armcostmanagement.BenefitUtilizationSummariesClientListByBillingAccountIDResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingAccountIDPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByBillingAccountIDPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingAccountIDPager) {
		b.newListByBillingAccountIDPager.remove(req)
	}
	return resp, nil
}

func (b *BenefitUtilizationSummariesServerTransport) dispatchNewListByBillingProfileIDPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByBillingProfileIDPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingProfileIDPager not implemented")}
	}
	newListByBillingProfileIDPager := b.newListByBillingProfileIDPager.get(req)
	if newListByBillingProfileIDPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/benefitUtilizationSummaries`
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
		billingProfileIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileId")])
		if err != nil {
			return nil, err
		}
		grainParameterUnescaped, err := url.QueryUnescape(qp.Get("grainParameter"))
		if err != nil {
			return nil, err
		}
		grainParameterParam := getOptional(armcostmanagement.GrainParameter(grainParameterUnescaped))
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDOptions
		if grainParameterParam != nil || filterParam != nil {
			options = &armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDOptions{
				GrainParameter: grainParameterParam,
				Filter:         filterParam,
			}
		}
		resp := b.srv.NewListByBillingProfileIDPager(billingAccountIDParam, billingProfileIDParam, options)
		newListByBillingProfileIDPager = &resp
		b.newListByBillingProfileIDPager.add(req, newListByBillingProfileIDPager)
		server.PagerResponderInjectNextLinks(newListByBillingProfileIDPager, req, func(page *armcostmanagement.BenefitUtilizationSummariesClientListByBillingProfileIDResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingProfileIDPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByBillingProfileIDPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingProfileIDPager) {
		b.newListByBillingProfileIDPager.remove(req)
	}
	return resp, nil
}

func (b *BenefitUtilizationSummariesServerTransport) dispatchNewListBySavingsPlanIDPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListBySavingsPlanIDPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySavingsPlanIDPager not implemented")}
	}
	newListBySavingsPlanIDPager := b.newListBySavingsPlanIDPager.get(req)
	if newListBySavingsPlanIDPager == nil {
		const regexStr = `/providers/Microsoft\.BillingBenefits/savingsPlanOrders/(?P<savingsPlanOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/savingsPlans/(?P<savingsPlanId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/benefitUtilizationSummaries`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		grainParameterUnescaped, err := url.QueryUnescape(qp.Get("grainParameter"))
		if err != nil {
			return nil, err
		}
		grainParameterParam := getOptional(armcostmanagement.GrainParameter(grainParameterUnescaped))
		savingsPlanOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savingsPlanOrderId")])
		if err != nil {
			return nil, err
		}
		savingsPlanIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savingsPlanId")])
		if err != nil {
			return nil, err
		}
		var options *armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanIDOptions
		if filterParam != nil || grainParameterParam != nil {
			options = &armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanIDOptions{
				Filter:         filterParam,
				GrainParameter: grainParameterParam,
			}
		}
		resp := b.srv.NewListBySavingsPlanIDPager(savingsPlanOrderIDParam, savingsPlanIDParam, options)
		newListBySavingsPlanIDPager = &resp
		b.newListBySavingsPlanIDPager.add(req, newListBySavingsPlanIDPager)
		server.PagerResponderInjectNextLinks(newListBySavingsPlanIDPager, req, func(page *armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanIDResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySavingsPlanIDPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListBySavingsPlanIDPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySavingsPlanIDPager) {
		b.newListBySavingsPlanIDPager.remove(req)
	}
	return resp, nil
}

func (b *BenefitUtilizationSummariesServerTransport) dispatchNewListBySavingsPlanOrderPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListBySavingsPlanOrderPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySavingsPlanOrderPager not implemented")}
	}
	newListBySavingsPlanOrderPager := b.newListBySavingsPlanOrderPager.get(req)
	if newListBySavingsPlanOrderPager == nil {
		const regexStr = `/providers/Microsoft\.BillingBenefits/savingsPlanOrders/(?P<savingsPlanOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/benefitUtilizationSummaries`
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
		grainParameterUnescaped, err := url.QueryUnescape(qp.Get("grainParameter"))
		if err != nil {
			return nil, err
		}
		grainParameterParam := getOptional(armcostmanagement.GrainParameter(grainParameterUnescaped))
		savingsPlanOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savingsPlanOrderId")])
		if err != nil {
			return nil, err
		}
		var options *armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions
		if filterParam != nil || grainParameterParam != nil {
			options = &armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions{
				Filter:         filterParam,
				GrainParameter: grainParameterParam,
			}
		}
		resp := b.srv.NewListBySavingsPlanOrderPager(savingsPlanOrderIDParam, options)
		newListBySavingsPlanOrderPager = &resp
		b.newListBySavingsPlanOrderPager.add(req, newListBySavingsPlanOrderPager)
		server.PagerResponderInjectNextLinks(newListBySavingsPlanOrderPager, req, func(page *armcostmanagement.BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySavingsPlanOrderPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListBySavingsPlanOrderPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySavingsPlanOrderPager) {
		b.newListBySavingsPlanOrderPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to BenefitUtilizationSummariesServerTransport
var benefitUtilizationSummariesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
