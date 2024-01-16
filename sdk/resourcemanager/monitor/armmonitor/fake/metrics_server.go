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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// MetricsServer is a fake server for instances of the armmonitor.MetricsClient type.
type MetricsServer struct {
	// List is the fake for method MetricsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceURI string, options *armmonitor.MetricsClientListOptions) (resp azfake.Responder[armmonitor.MetricsClientListResponse], errResp azfake.ErrorResponder)

	// ListAtSubscriptionScope is the fake for method MetricsClient.ListAtSubscriptionScope
	// HTTP status codes to indicate success: http.StatusOK
	ListAtSubscriptionScope func(ctx context.Context, region string, options *armmonitor.MetricsClientListAtSubscriptionScopeOptions) (resp azfake.Responder[armmonitor.MetricsClientListAtSubscriptionScopeResponse], errResp azfake.ErrorResponder)

	// ListAtSubscriptionScopePost is the fake for method MetricsClient.ListAtSubscriptionScopePost
	// HTTP status codes to indicate success: http.StatusOK
	ListAtSubscriptionScopePost func(ctx context.Context, region string, options *armmonitor.MetricsClientListAtSubscriptionScopePostOptions) (resp azfake.Responder[armmonitor.MetricsClientListAtSubscriptionScopePostResponse], errResp azfake.ErrorResponder)
}

// NewMetricsServerTransport creates a new instance of MetricsServerTransport with the provided implementation.
// The returned MetricsServerTransport instance is connected to an instance of armmonitor.MetricsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMetricsServerTransport(srv *MetricsServer) *MetricsServerTransport {
	return &MetricsServerTransport{srv: srv}
}

// MetricsServerTransport connects instances of armmonitor.MetricsClient to instances of MetricsServer.
// Don't use this type directly, use NewMetricsServerTransport instead.
type MetricsServerTransport struct {
	srv *MetricsServer
}

// Do implements the policy.Transporter interface for MetricsServerTransport.
func (m *MetricsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MetricsClient.List":
		resp, err = m.dispatchList(req)
	case "MetricsClient.ListAtSubscriptionScope":
		resp, err = m.dispatchListAtSubscriptionScope(req)
	case "MetricsClient.ListAtSubscriptionScopePost":
		resp, err = m.dispatchListAtSubscriptionScopePost(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MetricsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if m.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/metrics`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	timespanUnescaped, err := url.QueryUnescape(qp.Get("timespan"))
	if err != nil {
		return nil, err
	}
	timespanParam := getOptional(timespanUnescaped)
	intervalUnescaped, err := url.QueryUnescape(qp.Get("interval"))
	if err != nil {
		return nil, err
	}
	intervalParam := getOptional(intervalUnescaped)
	metricnamesUnescaped, err := url.QueryUnescape(qp.Get("metricnames"))
	if err != nil {
		return nil, err
	}
	metricnamesParam := getOptional(metricnamesUnescaped)
	aggregationUnescaped, err := url.QueryUnescape(qp.Get("aggregation"))
	if err != nil {
		return nil, err
	}
	aggregationParam := getOptional(aggregationUnescaped)
	topUnescaped, err := url.QueryUnescape(qp.Get("top"))
	if err != nil {
		return nil, err
	}
	topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	orderbyUnescaped, err := url.QueryUnescape(qp.Get("orderby"))
	if err != nil {
		return nil, err
	}
	orderbyParam := getOptional(orderbyUnescaped)
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	resultTypeUnescaped, err := url.QueryUnescape(qp.Get("resultType"))
	if err != nil {
		return nil, err
	}
	resultTypeParam := getOptional(armmonitor.ResultType(resultTypeUnescaped))
	metricnamespaceUnescaped, err := url.QueryUnescape(qp.Get("metricnamespace"))
	if err != nil {
		return nil, err
	}
	metricnamespaceParam := getOptional(metricnamespaceUnescaped)
	autoAdjustTimegrainUnescaped, err := url.QueryUnescape(qp.Get("AutoAdjustTimegrain"))
	if err != nil {
		return nil, err
	}
	autoAdjustTimegrainParam, err := parseOptional(autoAdjustTimegrainUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	validateDimensionsUnescaped, err := url.QueryUnescape(qp.Get("ValidateDimensions"))
	if err != nil {
		return nil, err
	}
	validateDimensionsParam, err := parseOptional(validateDimensionsUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	rollupbyUnescaped, err := url.QueryUnescape(qp.Get("rollupby"))
	if err != nil {
		return nil, err
	}
	rollupbyParam := getOptional(rollupbyUnescaped)
	var options *armmonitor.MetricsClientListOptions
	if timespanParam != nil || intervalParam != nil || metricnamesParam != nil || aggregationParam != nil || topParam != nil || orderbyParam != nil || filterParam != nil || resultTypeParam != nil || metricnamespaceParam != nil || autoAdjustTimegrainParam != nil || validateDimensionsParam != nil || rollupbyParam != nil {
		options = &armmonitor.MetricsClientListOptions{
			Timespan:            timespanParam,
			Interval:            intervalParam,
			Metricnames:         metricnamesParam,
			Aggregation:         aggregationParam,
			Top:                 topParam,
			Orderby:             orderbyParam,
			Filter:              filterParam,
			ResultType:          resultTypeParam,
			Metricnamespace:     metricnamespaceParam,
			AutoAdjustTimegrain: autoAdjustTimegrainParam,
			ValidateDimensions:  validateDimensionsParam,
			Rollupby:            rollupbyParam,
		}
	}
	respr, errRespr := m.srv.List(req.Context(), resourceURIParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Response, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MetricsServerTransport) dispatchListAtSubscriptionScope(req *http.Request) (*http.Response, error) {
	if m.srv.ListAtSubscriptionScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAtSubscriptionScope not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/metrics`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	regionParam, err := url.QueryUnescape(qp.Get("region"))
	if err != nil {
		return nil, err
	}
	timespanUnescaped, err := url.QueryUnescape(qp.Get("timespan"))
	if err != nil {
		return nil, err
	}
	timespanParam := getOptional(timespanUnescaped)
	intervalUnescaped, err := url.QueryUnescape(qp.Get("interval"))
	if err != nil {
		return nil, err
	}
	intervalParam := getOptional(intervalUnescaped)
	metricnamesUnescaped, err := url.QueryUnescape(qp.Get("metricnames"))
	if err != nil {
		return nil, err
	}
	metricnamesParam := getOptional(metricnamesUnescaped)
	aggregationUnescaped, err := url.QueryUnescape(qp.Get("aggregation"))
	if err != nil {
		return nil, err
	}
	aggregationParam := getOptional(aggregationUnescaped)
	topUnescaped, err := url.QueryUnescape(qp.Get("top"))
	if err != nil {
		return nil, err
	}
	topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	orderbyUnescaped, err := url.QueryUnescape(qp.Get("orderby"))
	if err != nil {
		return nil, err
	}
	orderbyParam := getOptional(orderbyUnescaped)
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	resultTypeUnescaped, err := url.QueryUnescape(qp.Get("resultType"))
	if err != nil {
		return nil, err
	}
	resultTypeParam := getOptional(armmonitor.MetricResultType(resultTypeUnescaped))
	metricnamespaceUnescaped, err := url.QueryUnescape(qp.Get("metricnamespace"))
	if err != nil {
		return nil, err
	}
	metricnamespaceParam := getOptional(metricnamespaceUnescaped)
	autoAdjustTimegrainUnescaped, err := url.QueryUnescape(qp.Get("AutoAdjustTimegrain"))
	if err != nil {
		return nil, err
	}
	autoAdjustTimegrainParam, err := parseOptional(autoAdjustTimegrainUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	validateDimensionsUnescaped, err := url.QueryUnescape(qp.Get("ValidateDimensions"))
	if err != nil {
		return nil, err
	}
	validateDimensionsParam, err := parseOptional(validateDimensionsUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	rollupbyUnescaped, err := url.QueryUnescape(qp.Get("rollupby"))
	if err != nil {
		return nil, err
	}
	rollupbyParam := getOptional(rollupbyUnescaped)
	var options *armmonitor.MetricsClientListAtSubscriptionScopeOptions
	if timespanParam != nil || intervalParam != nil || metricnamesParam != nil || aggregationParam != nil || topParam != nil || orderbyParam != nil || filterParam != nil || resultTypeParam != nil || metricnamespaceParam != nil || autoAdjustTimegrainParam != nil || validateDimensionsParam != nil || rollupbyParam != nil {
		options = &armmonitor.MetricsClientListAtSubscriptionScopeOptions{
			Timespan:            timespanParam,
			Interval:            intervalParam,
			Metricnames:         metricnamesParam,
			Aggregation:         aggregationParam,
			Top:                 topParam,
			Orderby:             orderbyParam,
			Filter:              filterParam,
			ResultType:          resultTypeParam,
			Metricnamespace:     metricnamespaceParam,
			AutoAdjustTimegrain: autoAdjustTimegrainParam,
			ValidateDimensions:  validateDimensionsParam,
			Rollupby:            rollupbyParam,
		}
	}
	respr, errRespr := m.srv.ListAtSubscriptionScope(req.Context(), regionParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Response, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MetricsServerTransport) dispatchListAtSubscriptionScopePost(req *http.Request) (*http.Response, error) {
	if m.srv.ListAtSubscriptionScopePost == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAtSubscriptionScopePost not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/metrics`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armmonitor.SubscriptionScopeMetricsRequestBodyParameters](req)
	if err != nil {
		return nil, err
	}
	regionParam, err := url.QueryUnescape(qp.Get("region"))
	if err != nil {
		return nil, err
	}
	var options *armmonitor.MetricsClientListAtSubscriptionScopePostOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmonitor.MetricsClientListAtSubscriptionScopePostOptions{
			Body: &body,
		}
	}
	respr, errRespr := m.srv.ListAtSubscriptionScopePost(req.Context(), regionParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Response, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
