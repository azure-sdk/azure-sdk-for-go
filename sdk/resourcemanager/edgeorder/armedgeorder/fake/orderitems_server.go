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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// OrderItemsServer is a fake server for instances of the armedgeorder.OrderItemsClient type.
type OrderItemsServer struct {
	// Cancel is the fake for method OrderItemsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Cancel func(ctx context.Context, resourceGroupName string, orderItemName string, cancellationReason armedgeorder.CancellationReason, options *armedgeorder.OrderItemsClientCancelOptions) (resp azfake.Responder[armedgeorder.OrderItemsClientCancelResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method OrderItemsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource armedgeorder.OrderItemResource, options *armedgeorder.OrderItemsClientBeginCreateOptions) (resp azfake.PollerResponder[armedgeorder.OrderItemsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method OrderItemsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, orderItemName string, options *armedgeorder.OrderItemsClientBeginDeleteOptions) (resp azfake.PollerResponder[armedgeorder.OrderItemsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OrderItemsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, orderItemName string, options *armedgeorder.OrderItemsClientGetOptions) (resp azfake.Responder[armedgeorder.OrderItemsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method OrderItemsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armedgeorder.OrderItemsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armedgeorder.OrderItemsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method OrderItemsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armedgeorder.OrderItemsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armedgeorder.OrderItemsClientListBySubscriptionResponse])

	// BeginReturn is the fake for method OrderItemsClient.BeginReturn
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginReturn func(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails armedgeorder.ReturnOrderItemDetails, options *armedgeorder.OrderItemsClientBeginReturnOptions) (resp azfake.PollerResponder[armedgeorder.OrderItemsClientReturnResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method OrderItemsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter armedgeorder.OrderItemUpdateParameter, options *armedgeorder.OrderItemsClientBeginUpdateOptions) (resp azfake.PollerResponder[armedgeorder.OrderItemsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewOrderItemsServerTransport creates a new instance of OrderItemsServerTransport with the provided implementation.
// The returned OrderItemsServerTransport instance is connected to an instance of armedgeorder.OrderItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOrderItemsServerTransport(srv *OrderItemsServer) *OrderItemsServerTransport {
	return &OrderItemsServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armedgeorder.OrderItemsClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armedgeorder.OrderItemsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armedgeorder.OrderItemsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armedgeorder.OrderItemsClientListBySubscriptionResponse]](),
		beginReturn:                 newTracker[azfake.PollerResponder[armedgeorder.OrderItemsClientReturnResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armedgeorder.OrderItemsClientUpdateResponse]](),
	}
}

// OrderItemsServerTransport connects instances of armedgeorder.OrderItemsClient to instances of OrderItemsServer.
// Don't use this type directly, use NewOrderItemsServerTransport instead.
type OrderItemsServerTransport struct {
	srv                         *OrderItemsServer
	beginCreate                 *tracker[azfake.PollerResponder[armedgeorder.OrderItemsClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armedgeorder.OrderItemsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armedgeorder.OrderItemsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armedgeorder.OrderItemsClientListBySubscriptionResponse]]
	beginReturn                 *tracker[azfake.PollerResponder[armedgeorder.OrderItemsClientReturnResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armedgeorder.OrderItemsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for OrderItemsServerTransport.
func (o *OrderItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OrderItemsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if orderItemsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = orderItemsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OrderItemsClient.Cancel":
				res.resp, res.err = o.dispatchCancel(req)
			case "OrderItemsClient.BeginCreate":
				res.resp, res.err = o.dispatchBeginCreate(req)
			case "OrderItemsClient.BeginDelete":
				res.resp, res.err = o.dispatchBeginDelete(req)
			case "OrderItemsClient.Get":
				res.resp, res.err = o.dispatchGet(req)
			case "OrderItemsClient.NewListByResourceGroupPager":
				res.resp, res.err = o.dispatchNewListByResourceGroupPager(req)
			case "OrderItemsClient.NewListBySubscriptionPager":
				res.resp, res.err = o.dispatchNewListBySubscriptionPager(req)
			case "OrderItemsClient.BeginReturn":
				res.resp, res.err = o.dispatchBeginReturn(req)
			case "OrderItemsClient.BeginUpdate":
				res.resp, res.err = o.dispatchBeginUpdate(req)
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

func (o *OrderItemsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if o.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems/(?P<orderItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armedgeorder.CancellationReason](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	orderItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("orderItemName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Cancel(req.Context(), resourceGroupNameParam, orderItemNameParam, body, nil)
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

func (o *OrderItemsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := o.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems/(?P<orderItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armedgeorder.OrderItemResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		orderItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("orderItemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginCreate(req.Context(), resourceGroupNameParam, orderItemNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		o.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		o.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		o.beginCreate.remove(req)
	}

	return resp, nil
}

func (o *OrderItemsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if o.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := o.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems/(?P<orderItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		orderItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("orderItemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginDelete(req.Context(), resourceGroupNameParam, orderItemNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		o.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		o.beginDelete.remove(req)
	}

	return resp, nil
}

func (o *OrderItemsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems/(?P<orderItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	orderItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("orderItemName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armedgeorder.OrderItemsClientGetOptions
	if expandParam != nil {
		options = &armedgeorder.OrderItemsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, orderItemNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OrderItemResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OrderItemsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := o.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
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
		var options *armedgeorder.OrderItemsClientListByResourceGroupOptions
		if filterParam != nil || expandParam != nil || skipTokenParam != nil || topParam != nil {
			options = &armedgeorder.OrderItemsClientListByResourceGroupOptions{
				Filter:    filterParam,
				Expand:    expandParam,
				SkipToken: skipTokenParam,
				Top:       topParam,
			}
		}
		resp := o.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		o.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armedgeorder.OrderItemsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		o.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (o *OrderItemsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := o.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems`
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
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
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
		var options *armedgeorder.OrderItemsClientListBySubscriptionOptions
		if filterParam != nil || expandParam != nil || skipTokenParam != nil || topParam != nil {
			options = &armedgeorder.OrderItemsClientListBySubscriptionOptions{
				Filter:    filterParam,
				Expand:    expandParam,
				SkipToken: skipTokenParam,
				Top:       topParam,
			}
		}
		resp := o.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		o.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armedgeorder.OrderItemsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		o.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (o *OrderItemsServerTransport) dispatchBeginReturn(req *http.Request) (*http.Response, error) {
	if o.srv.BeginReturn == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReturn not implemented")}
	}
	beginReturn := o.beginReturn.get(req)
	if beginReturn == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems/(?P<orderItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/return`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armedgeorder.ReturnOrderItemDetails](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		orderItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("orderItemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginReturn(req.Context(), resourceGroupNameParam, orderItemNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReturn = &respr
		o.beginReturn.add(req, beginReturn)
	}

	resp, err := server.PollerResponderNext(beginReturn, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginReturn.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReturn) {
		o.beginReturn.remove(req)
	}

	return resp, nil
}

func (o *OrderItemsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := o.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EdgeOrder/orderItems/(?P<orderItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armedgeorder.OrderItemUpdateParameter](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		orderItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("orderItemName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *armedgeorder.OrderItemsClientBeginUpdateOptions
		if ifMatchParam != nil {
			options = &armedgeorder.OrderItemsClientBeginUpdateOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := o.srv.BeginUpdate(req.Context(), resourceGroupNameParam, orderItemNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		o.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		o.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		o.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to OrderItemsServerTransport
var orderItemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
