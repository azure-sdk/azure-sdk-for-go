// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
	"net/http"
	"net/url"
	"regexp"
)

// ImpactCategoriesServer is a fake server for instances of the armimpactreporting.ImpactCategoriesClient type.
type ImpactCategoriesServer struct {
	// Get is the fake for method ImpactCategoriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, impactCategoryName string, options *armimpactreporting.ImpactCategoriesClientGetOptions) (resp azfake.Responder[armimpactreporting.ImpactCategoriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySubscriptionPager is the fake for method ImpactCategoriesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armimpactreporting.ImpactCategoriesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armimpactreporting.ImpactCategoriesClientListBySubscriptionResponse])
}

// NewImpactCategoriesServerTransport creates a new instance of ImpactCategoriesServerTransport with the provided implementation.
// The returned ImpactCategoriesServerTransport instance is connected to an instance of armimpactreporting.ImpactCategoriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewImpactCategoriesServerTransport(srv *ImpactCategoriesServer) *ImpactCategoriesServerTransport {
	return &ImpactCategoriesServerTransport{
		srv:                        srv,
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armimpactreporting.ImpactCategoriesClientListBySubscriptionResponse]](),
	}
}

// ImpactCategoriesServerTransport connects instances of armimpactreporting.ImpactCategoriesClient to instances of ImpactCategoriesServer.
// Don't use this type directly, use NewImpactCategoriesServerTransport instead.
type ImpactCategoriesServerTransport struct {
	srv                        *ImpactCategoriesServer
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armimpactreporting.ImpactCategoriesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for ImpactCategoriesServerTransport.
func (i *ImpactCategoriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *ImpactCategoriesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if impactCategoriesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = impactCategoriesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ImpactCategoriesClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "ImpactCategoriesClient.NewListBySubscriptionPager":
				res.resp, res.err = i.dispatchNewListBySubscriptionPager(req)
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

func (i *ImpactCategoriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Impact/impactCategories/(?P<impactCategoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	impactCategoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("impactCategoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), impactCategoryNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImpactCategory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImpactCategoriesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := i.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Impact/impactCategories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		categoryNameUnescaped, err := url.QueryUnescape(qp.Get("categoryName"))
		if err != nil {
			return nil, err
		}
		categoryNameParam := getOptional(categoryNameUnescaped)
		resourceTypeUnescaped, err := url.QueryUnescape(qp.Get("resourceType"))
		if err != nil {
			return nil, err
		}
		resourceTypeParam := getOptional(resourceTypeUnescaped)
		var options *armimpactreporting.ImpactCategoriesClientListBySubscriptionOptions
		if categoryNameParam != nil || resourceTypeParam != nil {
			options = &armimpactreporting.ImpactCategoriesClientListBySubscriptionOptions{
				CategoryName: categoryNameParam,
				ResourceType: resourceTypeParam,
			}
		}
		resp := i.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		i.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armimpactreporting.ImpactCategoriesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		i.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ImpactCategoriesServerTransport
var impactCategoriesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
