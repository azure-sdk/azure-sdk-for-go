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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ProductServer is a fake server for instances of the armapimanagement.ProductClient type.
type ProductServer struct {
	// CreateOrUpdate is the fake for method ProductClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, productID string, parameters armapimanagement.ProductContract, options *armapimanagement.ProductClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.ProductClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProductClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, productID string, ifMatch string, options *armapimanagement.ProductClientDeleteOptions) (resp azfake.Responder[armapimanagement.ProductClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProductClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *armapimanagement.ProductClientGetOptions) (resp azfake.Responder[armapimanagement.ProductClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method ProductClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *armapimanagement.ProductClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.ProductClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method ProductClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, options *armapimanagement.ProductClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.ProductClientListByServiceResponse])

	// NewListByTagsPager is the fake for method ProductClient.NewListByTagsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByTagsPager func(resourceGroupName string, serviceName string, options *armapimanagement.ProductClientListByTagsOptions) (resp azfake.PagerResponder[armapimanagement.ProductClientListByTagsResponse])

	// Update is the fake for method ProductClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, serviceName string, productID string, ifMatch string, parameters armapimanagement.ProductUpdateParameters, options *armapimanagement.ProductClientUpdateOptions) (resp azfake.Responder[armapimanagement.ProductClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewProductServerTransport creates a new instance of ProductServerTransport with the provided implementation.
// The returned ProductServerTransport instance is connected to an instance of armapimanagement.ProductClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProductServerTransport(srv *ProductServer) *ProductServerTransport {
	return &ProductServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.ProductClientListByServiceResponse]](),
		newListByTagsPager:    newTracker[azfake.PagerResponder[armapimanagement.ProductClientListByTagsResponse]](),
	}
}

// ProductServerTransport connects instances of armapimanagement.ProductClient to instances of ProductServer.
// Don't use this type directly, use NewProductServerTransport instead.
type ProductServerTransport struct {
	srv                   *ProductServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.ProductClientListByServiceResponse]]
	newListByTagsPager    *tracker[azfake.PagerResponder[armapimanagement.ProductClientListByTagsResponse]]
}

// Do implements the policy.Transporter interface for ProductServerTransport.
func (p *ProductServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProductClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "ProductClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "ProductClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProductClient.GetEntityTag":
		resp, err = p.dispatchGetEntityTag(req)
	case "ProductClient.NewListByServicePager":
		resp, err = p.dispatchNewListByServicePager(req)
	case "ProductClient.NewListByTagsPager":
		resp, err = p.dispatchNewListByTagsPager(req)
	case "ProductClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProductServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.ProductContract](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.ProductClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.ProductClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProductContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *ProductServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	deleteSubscriptionsUnescaped, err := url.QueryUnescape(qp.Get("deleteSubscriptions"))
	if err != nil {
		return nil, err
	}
	deleteSubscriptionsParam, err := parseOptional(deleteSubscriptionsUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armapimanagement.ProductClientDeleteOptions
	if deleteSubscriptionsParam != nil {
		options = &armapimanagement.ProductClientDeleteOptions{
			DeleteSubscriptions: deleteSubscriptionsParam,
		}
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, getHeaderValue(req.Header, "If-Match"), options)
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

func (p *ProductServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProductContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *ProductServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if p.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *ProductServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := p.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products`
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
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		expandGroupsUnescaped, err := url.QueryUnescape(qp.Get("expandGroups"))
		if err != nil {
			return nil, err
		}
		expandGroupsParam, err := parseOptional(expandGroupsUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		tagsUnescaped, err := url.QueryUnescape(qp.Get("tags"))
		if err != nil {
			return nil, err
		}
		tagsParam := getOptional(tagsUnescaped)
		var options *armapimanagement.ProductClientListByServiceOptions
		if filterParam != nil || topParam != nil || skipParam != nil || expandGroupsParam != nil || tagsParam != nil {
			options = &armapimanagement.ProductClientListByServiceOptions{
				Filter:       filterParam,
				Top:          topParam,
				Skip:         skipParam,
				ExpandGroups: expandGroupsParam,
				Tags:         tagsParam,
			}
		}
		resp := p.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, options)
		newListByServicePager = &resp
		p.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.ProductClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		p.newListByServicePager.remove(req)
	}
	return resp, nil
}

func (p *ProductServerTransport) dispatchNewListByTagsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByTagsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByTagsPager not implemented")}
	}
	newListByTagsPager := p.newListByTagsPager.get(req)
	if newListByTagsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/productsByTags`
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
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		includeNotTaggedProductsUnescaped, err := url.QueryUnescape(qp.Get("includeNotTaggedProducts"))
		if err != nil {
			return nil, err
		}
		includeNotTaggedProductsParam, err := parseOptional(includeNotTaggedProductsUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armapimanagement.ProductClientListByTagsOptions
		if filterParam != nil || topParam != nil || skipParam != nil || includeNotTaggedProductsParam != nil {
			options = &armapimanagement.ProductClientListByTagsOptions{
				Filter:                   filterParam,
				Top:                      topParam,
				Skip:                     skipParam,
				IncludeNotTaggedProducts: includeNotTaggedProductsParam,
			}
		}
		resp := p.srv.NewListByTagsPager(resourceGroupNameParam, serviceNameParam, options)
		newListByTagsPager = &resp
		p.newListByTagsPager.add(req, newListByTagsPager)
		server.PagerResponderInjectNextLinks(newListByTagsPager, req, func(page *armapimanagement.ProductClientListByTagsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByTagsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByTagsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByTagsPager) {
		p.newListByTagsPager.remove(req)
	}
	return resp, nil
}

func (p *ProductServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.ProductUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Update(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, getHeaderValue(req.Header, "If-Match"), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProductContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}
