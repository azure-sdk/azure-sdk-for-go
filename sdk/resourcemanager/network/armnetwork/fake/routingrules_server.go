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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// RoutingRulesServer is a fake server for instances of the armnetwork.RoutingRulesClient type.
type RoutingRulesServer struct {
	// CreateOrUpdate is the fake for method RoutingRulesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, routingRule armnetwork.RoutingRule, options *armnetwork.RoutingRulesClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.RoutingRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RoutingRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *armnetwork.RoutingRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.RoutingRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RoutingRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *armnetwork.RoutingRulesClientGetOptions) (resp azfake.Responder[armnetwork.RoutingRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RoutingRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *armnetwork.RoutingRulesClientListOptions) (resp azfake.PagerResponder[armnetwork.RoutingRulesClientListResponse])
}

// NewRoutingRulesServerTransport creates a new instance of RoutingRulesServerTransport with the provided implementation.
// The returned RoutingRulesServerTransport instance is connected to an instance of armnetwork.RoutingRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRoutingRulesServerTransport(srv *RoutingRulesServer) *RoutingRulesServerTransport {
	return &RoutingRulesServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armnetwork.RoutingRulesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetwork.RoutingRulesClientListResponse]](),
	}
}

// RoutingRulesServerTransport connects instances of armnetwork.RoutingRulesClient to instances of RoutingRulesServer.
// Don't use this type directly, use NewRoutingRulesServerTransport instead.
type RoutingRulesServerTransport struct {
	srv          *RoutingRulesServer
	beginDelete  *tracker[azfake.PollerResponder[armnetwork.RoutingRulesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetwork.RoutingRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for RoutingRulesServerTransport.
func (r *RoutingRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RoutingRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if routingRulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = routingRulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RoutingRulesClient.CreateOrUpdate":
				res.resp, res.err = r.dispatchCreateOrUpdate(req)
			case "RoutingRulesClient.BeginDelete":
				res.resp, res.err = r.dispatchBeginDelete(req)
			case "RoutingRulesClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RoutingRulesClient.NewListPager":
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

func (r *RoutingRulesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routingConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.RoutingRule](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
	if err != nil {
		return nil, err
	}
	ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
	if err != nil {
		return nil, err
	}
	ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoutingRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoutingRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routingConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
		if err != nil {
			return nil, err
		}
		ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
		if err != nil {
			return nil, err
		}
		ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
		if err != nil {
			return nil, err
		}
		forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
		if err != nil {
			return nil, err
		}
		forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armnetwork.RoutingRulesClientBeginDeleteOptions
		if forceParam != nil {
			options = &armnetwork.RoutingRulesClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, options)
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

func (r *RoutingRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routingConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
	if err != nil {
		return nil, err
	}
	ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
	if err != nil {
		return nil, err
	}
	ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoutingRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoutingRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routingConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
		if err != nil {
			return nil, err
		}
		ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
		if err != nil {
			return nil, err
		}
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
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.RoutingRulesClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.RoutingRulesClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.RoutingRulesClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to RoutingRulesServerTransport
var routingRulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
