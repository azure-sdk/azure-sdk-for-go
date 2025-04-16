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

// ProductGroupServer is a fake server for instances of the armapimanagement.ProductGroupClient type.
type ProductGroupServer struct {
	// CheckEntityExists is the fake for method ProductGroupClient.CheckEntityExists
	// HTTP status codes to indicate success: http.StatusNoContent
	CheckEntityExists func(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *armapimanagement.ProductGroupClientCheckEntityExistsOptions) (resp azfake.Responder[armapimanagement.ProductGroupClientCheckEntityExistsResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method ProductGroupClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *armapimanagement.ProductGroupClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.ProductGroupClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProductGroupClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string, options *armapimanagement.ProductGroupClientDeleteOptions) (resp azfake.Responder[armapimanagement.ProductGroupClientDeleteResponse], errResp azfake.ErrorResponder)

	// NewListByProductPager is the fake for method ProductGroupClient.NewListByProductPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProductPager func(resourceGroupName string, serviceName string, productID string, options *armapimanagement.ProductGroupClientListByProductOptions) (resp azfake.PagerResponder[armapimanagement.ProductGroupClientListByProductResponse])
}

// NewProductGroupServerTransport creates a new instance of ProductGroupServerTransport with the provided implementation.
// The returned ProductGroupServerTransport instance is connected to an instance of armapimanagement.ProductGroupClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProductGroupServerTransport(srv *ProductGroupServer) *ProductGroupServerTransport {
	return &ProductGroupServerTransport{
		srv:                   srv,
		newListByProductPager: newTracker[azfake.PagerResponder[armapimanagement.ProductGroupClientListByProductResponse]](),
	}
}

// ProductGroupServerTransport connects instances of armapimanagement.ProductGroupClient to instances of ProductGroupServer.
// Don't use this type directly, use NewProductGroupServerTransport instead.
type ProductGroupServerTransport struct {
	srv                   *ProductGroupServer
	newListByProductPager *tracker[azfake.PagerResponder[armapimanagement.ProductGroupClientListByProductResponse]]
}

// Do implements the policy.Transporter interface for ProductGroupServerTransport.
func (p *ProductGroupServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProductGroupServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if productGroupServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = productGroupServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProductGroupClient.CheckEntityExists":
				res.resp, res.err = p.dispatchCheckEntityExists(req)
			case "ProductGroupClient.CreateOrUpdate":
				res.resp, res.err = p.dispatchCreateOrUpdate(req)
			case "ProductGroupClient.Delete":
				res.resp, res.err = p.dispatchDelete(req)
			case "ProductGroupClient.NewListByProductPager":
				res.resp, res.err = p.dispatchNewListByProductPager(req)
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

func (p *ProductGroupServerTransport) dispatchCheckEntityExists(req *http.Request) (*http.Response, error) {
	if p.srv.CheckEntityExists == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckEntityExists not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CheckEntityExists(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, groupIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProductGroupServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, groupIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GroupContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProductGroupServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, groupIDParam, nil)
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

func (p *ProductGroupServerTransport) dispatchNewListByProductPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByProductPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProductPager not implemented")}
	}
	newListByProductPager := p.newListByProductPager.get(req)
	if newListByProductPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups`
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
		var options *armapimanagement.ProductGroupClientListByProductOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.ProductGroupClientListByProductOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := p.srv.NewListByProductPager(resourceGroupNameParam, serviceNameParam, productIDParam, options)
		newListByProductPager = &resp
		p.newListByProductPager.add(req, newListByProductPager)
		server.PagerResponderInjectNextLinks(newListByProductPager, req, func(page *armapimanagement.ProductGroupClientListByProductResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProductPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByProductPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProductPager) {
		p.newListByProductPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ProductGroupServerTransport
var productGroupServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
