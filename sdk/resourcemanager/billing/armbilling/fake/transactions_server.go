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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// TransactionsServer is a fake server for instances of the armbilling.TransactionsClient type.
type TransactionsServer struct {
	// GetTransactionSummaryByInvoice is the fake for method TransactionsClient.GetTransactionSummaryByInvoice
	// HTTP status codes to indicate success: http.StatusOK
	GetTransactionSummaryByInvoice func(ctx context.Context, billingAccountName string, invoiceName string, options *armbilling.TransactionsClientGetTransactionSummaryByInvoiceOptions) (resp azfake.Responder[armbilling.TransactionsClientGetTransactionSummaryByInvoiceResponse], errResp azfake.ErrorResponder)

	// NewListByBillingProfilePager is the fake for method TransactionsClient.NewListByBillingProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingProfilePager func(billingAccountName string, billingProfileName string, periodStartDate time.Time, periodEndDate time.Time, typeParam armbilling.TransactionType, options *armbilling.TransactionsClientListByBillingProfileOptions) (resp azfake.PagerResponder[armbilling.TransactionsClientListByBillingProfileResponse])

	// NewListByCustomerPager is the fake for method TransactionsClient.NewListByCustomerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCustomerPager func(billingAccountName string, billingProfileName string, customerName string, periodStartDate time.Time, periodEndDate time.Time, typeParam armbilling.TransactionType, options *armbilling.TransactionsClientListByCustomerOptions) (resp azfake.PagerResponder[armbilling.TransactionsClientListByCustomerResponse])

	// NewListByInvoicePager is the fake for method TransactionsClient.NewListByInvoicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInvoicePager func(billingAccountName string, invoiceName string, options *armbilling.TransactionsClientListByInvoiceOptions) (resp azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceResponse])

	// NewListByInvoiceSectionPager is the fake for method TransactionsClient.NewListByInvoiceSectionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInvoiceSectionPager func(billingAccountName string, billingProfileName string, invoiceSectionName string, periodStartDate time.Time, periodEndDate time.Time, typeParam armbilling.TransactionType, options *armbilling.TransactionsClientListByInvoiceSectionOptions) (resp azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceSectionResponse])

	// BeginTransactionsDownloadByInvoice is the fake for method TransactionsClient.BeginTransactionsDownloadByInvoice
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginTransactionsDownloadByInvoice func(ctx context.Context, billingAccountName string, invoiceName string, options *armbilling.TransactionsClientBeginTransactionsDownloadByInvoiceOptions) (resp azfake.PollerResponder[armbilling.TransactionsClientTransactionsDownloadByInvoiceResponse], errResp azfake.ErrorResponder)
}

// NewTransactionsServerTransport creates a new instance of TransactionsServerTransport with the provided implementation.
// The returned TransactionsServerTransport instance is connected to an instance of armbilling.TransactionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTransactionsServerTransport(srv *TransactionsServer) *TransactionsServerTransport {
	return &TransactionsServerTransport{
		srv:                                srv,
		newListByBillingProfilePager:       newTracker[azfake.PagerResponder[armbilling.TransactionsClientListByBillingProfileResponse]](),
		newListByCustomerPager:             newTracker[azfake.PagerResponder[armbilling.TransactionsClientListByCustomerResponse]](),
		newListByInvoicePager:              newTracker[azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceResponse]](),
		newListByInvoiceSectionPager:       newTracker[azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceSectionResponse]](),
		beginTransactionsDownloadByInvoice: newTracker[azfake.PollerResponder[armbilling.TransactionsClientTransactionsDownloadByInvoiceResponse]](),
	}
}

// TransactionsServerTransport connects instances of armbilling.TransactionsClient to instances of TransactionsServer.
// Don't use this type directly, use NewTransactionsServerTransport instead.
type TransactionsServerTransport struct {
	srv                                *TransactionsServer
	newListByBillingProfilePager       *tracker[azfake.PagerResponder[armbilling.TransactionsClientListByBillingProfileResponse]]
	newListByCustomerPager             *tracker[azfake.PagerResponder[armbilling.TransactionsClientListByCustomerResponse]]
	newListByInvoicePager              *tracker[azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceResponse]]
	newListByInvoiceSectionPager       *tracker[azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceSectionResponse]]
	beginTransactionsDownloadByInvoice *tracker[azfake.PollerResponder[armbilling.TransactionsClientTransactionsDownloadByInvoiceResponse]]
}

// Do implements the policy.Transporter interface for TransactionsServerTransport.
func (t *TransactionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TransactionsClient.GetTransactionSummaryByInvoice":
		resp, err = t.dispatchGetTransactionSummaryByInvoice(req)
	case "TransactionsClient.NewListByBillingProfilePager":
		resp, err = t.dispatchNewListByBillingProfilePager(req)
	case "TransactionsClient.NewListByCustomerPager":
		resp, err = t.dispatchNewListByCustomerPager(req)
	case "TransactionsClient.NewListByInvoicePager":
		resp, err = t.dispatchNewListByInvoicePager(req)
	case "TransactionsClient.NewListByInvoiceSectionPager":
		resp, err = t.dispatchNewListByInvoiceSectionPager(req)
	case "TransactionsClient.BeginTransactionsDownloadByInvoice":
		resp, err = t.dispatchBeginTransactionsDownloadByInvoice(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransactionsServerTransport) dispatchGetTransactionSummaryByInvoice(req *http.Request) (*http.Response, error) {
	if t.srv.GetTransactionSummaryByInvoice == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetTransactionSummaryByInvoice not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoices/(?P<invoiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transactionSummary`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	invoiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceName")])
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	searchUnescaped, err := url.QueryUnescape(qp.Get("search"))
	if err != nil {
		return nil, err
	}
	searchParam := getOptional(searchUnescaped)
	var options *armbilling.TransactionsClientGetTransactionSummaryByInvoiceOptions
	if filterParam != nil || searchParam != nil {
		options = &armbilling.TransactionsClientGetTransactionSummaryByInvoiceOptions{
			Filter: filterParam,
			Search: searchParam,
		}
	}
	respr, errRespr := t.srv.GetTransactionSummaryByInvoice(req.Context(), billingAccountNameParam, invoiceNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TransactionSummary, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TransactionsServerTransport) dispatchNewListByBillingProfilePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByBillingProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingProfilePager not implemented")}
	}
	newListByBillingProfilePager := t.newListByBillingProfilePager.get(req)
	if newListByBillingProfilePager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transactions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		periodStartDateUnescaped, err := url.QueryUnescape(qp.Get("periodStartDate"))
		if err != nil {
			return nil, err
		}
		periodStartDateParam, err := time.Parse("2006-01-02", periodStartDateUnescaped)
		if err != nil {
			return nil, err
		}
		periodEndDateUnescaped, err := url.QueryUnescape(qp.Get("periodEndDate"))
		if err != nil {
			return nil, err
		}
		periodEndDateParam, err := time.Parse("2006-01-02", periodEndDateUnescaped)
		if err != nil {
			return nil, err
		}
		typeParamParam, err := parseWithCast(qp.Get("type"), func(v string) (armbilling.TransactionType, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armbilling.TransactionType(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		countUnescaped, err := url.QueryUnescape(qp.Get("count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		searchUnescaped, err := url.QueryUnescape(qp.Get("search"))
		if err != nil {
			return nil, err
		}
		searchParam := getOptional(searchUnescaped)
		var options *armbilling.TransactionsClientListByBillingProfileOptions
		if filterParam != nil || orderByParam != nil || topParam != nil || skipParam != nil || countParam != nil || searchParam != nil {
			options = &armbilling.TransactionsClientListByBillingProfileOptions{
				Filter:  filterParam,
				OrderBy: orderByParam,
				Top:     topParam,
				Skip:    skipParam,
				Count:   countParam,
				Search:  searchParam,
			}
		}
		resp := t.srv.NewListByBillingProfilePager(billingAccountNameParam, billingProfileNameParam, periodStartDateParam, periodEndDateParam, typeParamParam, options)
		newListByBillingProfilePager = &resp
		t.newListByBillingProfilePager.add(req, newListByBillingProfilePager)
		server.PagerResponderInjectNextLinks(newListByBillingProfilePager, req, func(page *armbilling.TransactionsClientListByBillingProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByBillingProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingProfilePager) {
		t.newListByBillingProfilePager.remove(req)
	}
	return resp, nil
}

func (t *TransactionsServerTransport) dispatchNewListByCustomerPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByCustomerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCustomerPager not implemented")}
	}
	newListByCustomerPager := t.newListByCustomerPager.get(req)
	if newListByCustomerPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customers/(?P<customerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transactions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		customerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerName")])
		if err != nil {
			return nil, err
		}
		periodStartDateUnescaped, err := url.QueryUnescape(qp.Get("periodStartDate"))
		if err != nil {
			return nil, err
		}
		periodStartDateParam, err := time.Parse("2006-01-02", periodStartDateUnescaped)
		if err != nil {
			return nil, err
		}
		periodEndDateUnescaped, err := url.QueryUnescape(qp.Get("periodEndDate"))
		if err != nil {
			return nil, err
		}
		periodEndDateParam, err := time.Parse("2006-01-02", periodEndDateUnescaped)
		if err != nil {
			return nil, err
		}
		typeParamParam, err := parseWithCast(qp.Get("type"), func(v string) (armbilling.TransactionType, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armbilling.TransactionType(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		countUnescaped, err := url.QueryUnescape(qp.Get("count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		searchUnescaped, err := url.QueryUnescape(qp.Get("search"))
		if err != nil {
			return nil, err
		}
		searchParam := getOptional(searchUnescaped)
		var options *armbilling.TransactionsClientListByCustomerOptions
		if filterParam != nil || orderByParam != nil || topParam != nil || skipParam != nil || countParam != nil || searchParam != nil {
			options = &armbilling.TransactionsClientListByCustomerOptions{
				Filter:  filterParam,
				OrderBy: orderByParam,
				Top:     topParam,
				Skip:    skipParam,
				Count:   countParam,
				Search:  searchParam,
			}
		}
		resp := t.srv.NewListByCustomerPager(billingAccountNameParam, billingProfileNameParam, customerNameParam, periodStartDateParam, periodEndDateParam, typeParamParam, options)
		newListByCustomerPager = &resp
		t.newListByCustomerPager.add(req, newListByCustomerPager)
		server.PagerResponderInjectNextLinks(newListByCustomerPager, req, func(page *armbilling.TransactionsClientListByCustomerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCustomerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByCustomerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCustomerPager) {
		t.newListByCustomerPager.remove(req)
	}
	return resp, nil
}

func (t *TransactionsServerTransport) dispatchNewListByInvoicePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByInvoicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInvoicePager not implemented")}
	}
	newListByInvoicePager := t.newListByInvoicePager.get(req)
	if newListByInvoicePager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoices/(?P<invoiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transactions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		invoiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		countUnescaped, err := url.QueryUnescape(qp.Get("count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		searchUnescaped, err := url.QueryUnescape(qp.Get("search"))
		if err != nil {
			return nil, err
		}
		searchParam := getOptional(searchUnescaped)
		var options *armbilling.TransactionsClientListByInvoiceOptions
		if filterParam != nil || orderByParam != nil || topParam != nil || skipParam != nil || countParam != nil || searchParam != nil {
			options = &armbilling.TransactionsClientListByInvoiceOptions{
				Filter:  filterParam,
				OrderBy: orderByParam,
				Top:     topParam,
				Skip:    skipParam,
				Count:   countParam,
				Search:  searchParam,
			}
		}
		resp := t.srv.NewListByInvoicePager(billingAccountNameParam, invoiceNameParam, options)
		newListByInvoicePager = &resp
		t.newListByInvoicePager.add(req, newListByInvoicePager)
		server.PagerResponderInjectNextLinks(newListByInvoicePager, req, func(page *armbilling.TransactionsClientListByInvoiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInvoicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByInvoicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInvoicePager) {
		t.newListByInvoicePager.remove(req)
	}
	return resp, nil
}

func (t *TransactionsServerTransport) dispatchNewListByInvoiceSectionPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByInvoiceSectionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInvoiceSectionPager not implemented")}
	}
	newListByInvoiceSectionPager := t.newListByInvoiceSectionPager.get(req)
	if newListByInvoiceSectionPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoiceSections/(?P<invoiceSectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transactions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		invoiceSectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceSectionName")])
		if err != nil {
			return nil, err
		}
		periodStartDateUnescaped, err := url.QueryUnescape(qp.Get("periodStartDate"))
		if err != nil {
			return nil, err
		}
		periodStartDateParam, err := time.Parse("2006-01-02", periodStartDateUnescaped)
		if err != nil {
			return nil, err
		}
		periodEndDateUnescaped, err := url.QueryUnescape(qp.Get("periodEndDate"))
		if err != nil {
			return nil, err
		}
		periodEndDateParam, err := time.Parse("2006-01-02", periodEndDateUnescaped)
		if err != nil {
			return nil, err
		}
		typeParamParam, err := parseWithCast(qp.Get("type"), func(v string) (armbilling.TransactionType, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armbilling.TransactionType(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		countUnescaped, err := url.QueryUnescape(qp.Get("count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		searchUnescaped, err := url.QueryUnescape(qp.Get("search"))
		if err != nil {
			return nil, err
		}
		searchParam := getOptional(searchUnescaped)
		var options *armbilling.TransactionsClientListByInvoiceSectionOptions
		if filterParam != nil || orderByParam != nil || topParam != nil || skipParam != nil || countParam != nil || searchParam != nil {
			options = &armbilling.TransactionsClientListByInvoiceSectionOptions{
				Filter:  filterParam,
				OrderBy: orderByParam,
				Top:     topParam,
				Skip:    skipParam,
				Count:   countParam,
				Search:  searchParam,
			}
		}
		resp := t.srv.NewListByInvoiceSectionPager(billingAccountNameParam, billingProfileNameParam, invoiceSectionNameParam, periodStartDateParam, periodEndDateParam, typeParamParam, options)
		newListByInvoiceSectionPager = &resp
		t.newListByInvoiceSectionPager.add(req, newListByInvoiceSectionPager)
		server.PagerResponderInjectNextLinks(newListByInvoiceSectionPager, req, func(page *armbilling.TransactionsClientListByInvoiceSectionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInvoiceSectionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByInvoiceSectionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInvoiceSectionPager) {
		t.newListByInvoiceSectionPager.remove(req)
	}
	return resp, nil
}

func (t *TransactionsServerTransport) dispatchBeginTransactionsDownloadByInvoice(req *http.Request) (*http.Response, error) {
	if t.srv.BeginTransactionsDownloadByInvoice == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTransactionsDownloadByInvoice not implemented")}
	}
	beginTransactionsDownloadByInvoice := t.beginTransactionsDownloadByInvoice.get(req)
	if beginTransactionsDownloadByInvoice == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoices/(?P<invoiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transactionsDownload`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		invoiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginTransactionsDownloadByInvoice(req.Context(), billingAccountNameParam, invoiceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTransactionsDownloadByInvoice = &respr
		t.beginTransactionsDownloadByInvoice.add(req, beginTransactionsDownloadByInvoice)
	}

	resp, err := server.PollerResponderNext(beginTransactionsDownloadByInvoice, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginTransactionsDownloadByInvoice.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTransactionsDownloadByInvoice) {
		t.beginTransactionsDownloadByInvoice.remove(req)
	}

	return resp, nil
}
