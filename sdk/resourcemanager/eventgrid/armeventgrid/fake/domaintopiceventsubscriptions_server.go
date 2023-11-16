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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DomainTopicEventSubscriptionsServer is a fake server for instances of the armeventgrid.DomainTopicEventSubscriptionsClient type.
type DomainTopicEventSubscriptionsServer struct {
	// BeginCreateOrUpdate is the fake for method DomainTopicEventSubscriptionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, domainName string, topicName string, eventSubscriptionName string, eventSubscriptionInfo armeventgrid.EventSubscription, options *armeventgrid.DomainTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DomainTopicEventSubscriptionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, domainName string, topicName string, eventSubscriptionName string, options *armeventgrid.DomainTopicEventSubscriptionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DomainTopicEventSubscriptionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, domainName string, topicName string, eventSubscriptionName string, options *armeventgrid.DomainTopicEventSubscriptionsClientGetOptions) (resp azfake.Responder[armeventgrid.DomainTopicEventSubscriptionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetDeliveryAttributes is the fake for method DomainTopicEventSubscriptionsClient.GetDeliveryAttributes
	// HTTP status codes to indicate success: http.StatusOK
	GetDeliveryAttributes func(ctx context.Context, resourceGroupName string, domainName string, topicName string, eventSubscriptionName string, options *armeventgrid.DomainTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (resp azfake.Responder[armeventgrid.DomainTopicEventSubscriptionsClientGetDeliveryAttributesResponse], errResp azfake.ErrorResponder)

	// GetFullURL is the fake for method DomainTopicEventSubscriptionsClient.GetFullURL
	// HTTP status codes to indicate success: http.StatusOK
	GetFullURL func(ctx context.Context, resourceGroupName string, domainName string, topicName string, eventSubscriptionName string, options *armeventgrid.DomainTopicEventSubscriptionsClientGetFullURLOptions) (resp azfake.Responder[armeventgrid.DomainTopicEventSubscriptionsClientGetFullURLResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DomainTopicEventSubscriptionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, domainName string, topicName string, options *armeventgrid.DomainTopicEventSubscriptionsClientListOptions) (resp azfake.PagerResponder[armeventgrid.DomainTopicEventSubscriptionsClientListResponse])

	// BeginUpdate is the fake for method DomainTopicEventSubscriptionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, domainName string, topicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters armeventgrid.EventSubscriptionUpdateParameters, options *armeventgrid.DomainTopicEventSubscriptionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDomainTopicEventSubscriptionsServerTransport creates a new instance of DomainTopicEventSubscriptionsServerTransport with the provided implementation.
// The returned DomainTopicEventSubscriptionsServerTransport instance is connected to an instance of armeventgrid.DomainTopicEventSubscriptionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDomainTopicEventSubscriptionsServerTransport(srv *DomainTopicEventSubscriptionsServer) *DomainTopicEventSubscriptionsServerTransport {
	return &DomainTopicEventSubscriptionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armeventgrid.DomainTopicEventSubscriptionsClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientUpdateResponse]](),
	}
}

// DomainTopicEventSubscriptionsServerTransport connects instances of armeventgrid.DomainTopicEventSubscriptionsClient to instances of DomainTopicEventSubscriptionsServer.
// Don't use this type directly, use NewDomainTopicEventSubscriptionsServerTransport instead.
type DomainTopicEventSubscriptionsServerTransport struct {
	srv                 *DomainTopicEventSubscriptionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armeventgrid.DomainTopicEventSubscriptionsClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armeventgrid.DomainTopicEventSubscriptionsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for DomainTopicEventSubscriptionsServerTransport.
func (d *DomainTopicEventSubscriptionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DomainTopicEventSubscriptionsClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DomainTopicEventSubscriptionsClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DomainTopicEventSubscriptionsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DomainTopicEventSubscriptionsClient.GetDeliveryAttributes":
		resp, err = d.dispatchGetDeliveryAttributes(req)
	case "DomainTopicEventSubscriptionsClient.GetFullURL":
		resp, err = d.dispatchGetFullURL(req)
	case "DomainTopicEventSubscriptionsClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DomainTopicEventSubscriptionsClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DomainTopicEventSubscriptionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventSubscriptions/(?P<eventSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.EventSubscription](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		eventSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventSubscriptionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, domainNameParam, topicNameParam, eventSubscriptionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DomainTopicEventSubscriptionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventSubscriptions/(?P<eventSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		eventSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventSubscriptionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, domainNameParam, topicNameParam, eventSubscriptionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DomainTopicEventSubscriptionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventSubscriptions/(?P<eventSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
	if err != nil {
		return nil, err
	}
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	eventSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventSubscriptionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, domainNameParam, topicNameParam, eventSubscriptionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EventSubscription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DomainTopicEventSubscriptionsServerTransport) dispatchGetDeliveryAttributes(req *http.Request) (*http.Response, error) {
	if d.srv.GetDeliveryAttributes == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDeliveryAttributes not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventSubscriptions/(?P<eventSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDeliveryAttributes`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
	if err != nil {
		return nil, err
	}
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	eventSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventSubscriptionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetDeliveryAttributes(req.Context(), resourceGroupNameParam, domainNameParam, topicNameParam, eventSubscriptionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeliveryAttributeListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DomainTopicEventSubscriptionsServerTransport) dispatchGetFullURL(req *http.Request) (*http.Response, error) {
	if d.srv.GetFullURL == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFullURL not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventSubscriptions/(?P<eventSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getFullUrl`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
	if err != nil {
		return nil, err
	}
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	eventSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventSubscriptionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetFullURL(req.Context(), resourceGroupNameParam, domainNameParam, topicNameParam, eventSubscriptionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EventSubscriptionFullURL, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DomainTopicEventSubscriptionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventSubscriptions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
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
		var options *armeventgrid.DomainTopicEventSubscriptionsClientListOptions
		if filterParam != nil || topParam != nil {
			options = &armeventgrid.DomainTopicEventSubscriptionsClientListOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, domainNameParam, topicNameParam, options)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armeventgrid.DomainTopicEventSubscriptionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DomainTopicEventSubscriptionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := d.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventSubscriptions/(?P<eventSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.EventSubscriptionUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		eventSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventSubscriptionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameParam, domainNameParam, topicNameParam, eventSubscriptionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		d.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		d.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		d.beginUpdate.remove(req)
	}

	return resp, nil
}
