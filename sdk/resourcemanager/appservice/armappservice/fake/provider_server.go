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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ProviderServer is a fake server for instances of the armappservice.ProviderClient type.
type ProviderServer struct {
	// NewGetAvailableStacksPager is the fake for method ProviderClient.NewGetAvailableStacksPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetAvailableStacksPager func(options *armappservice.ProviderClientGetAvailableStacksOptions) (resp azfake.PagerResponder[armappservice.ProviderClientGetAvailableStacksResponse])

	// NewGetAvailableStacksOnPremPager is the fake for method ProviderClient.NewGetAvailableStacksOnPremPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetAvailableStacksOnPremPager func(options *armappservice.ProviderClientGetAvailableStacksOnPremOptions) (resp azfake.PagerResponder[armappservice.ProviderClientGetAvailableStacksOnPremResponse])

	// NewGetFunctionAppStacksPager is the fake for method ProviderClient.NewGetFunctionAppStacksPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetFunctionAppStacksPager func(options *armappservice.ProviderClientGetFunctionAppStacksOptions) (resp azfake.PagerResponder[armappservice.ProviderClientGetFunctionAppStacksResponse])

	// NewGetFunctionAppStacksForLocationPager is the fake for method ProviderClient.NewGetFunctionAppStacksForLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetFunctionAppStacksForLocationPager func(location string, options *armappservice.ProviderClientGetFunctionAppStacksForLocationOptions) (resp azfake.PagerResponder[armappservice.ProviderClientGetFunctionAppStacksForLocationResponse])

	// NewGetWebAppStacksPager is the fake for method ProviderClient.NewGetWebAppStacksPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetWebAppStacksPager func(options *armappservice.ProviderClientGetWebAppStacksOptions) (resp azfake.PagerResponder[armappservice.ProviderClientGetWebAppStacksResponse])

	// NewGetWebAppStacksForLocationPager is the fake for method ProviderClient.NewGetWebAppStacksForLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetWebAppStacksForLocationPager func(location string, options *armappservice.ProviderClientGetWebAppStacksForLocationOptions) (resp azfake.PagerResponder[armappservice.ProviderClientGetWebAppStacksForLocationResponse])

	// NewListOperationsPager is the fake for method ProviderClient.NewListOperationsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListOperationsPager func(options *armappservice.ProviderClientListOperationsOptions) (resp azfake.PagerResponder[armappservice.ProviderClientListOperationsResponse])
}

// NewProviderServerTransport creates a new instance of ProviderServerTransport with the provided implementation.
// The returned ProviderServerTransport instance is connected to an instance of armappservice.ProviderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProviderServerTransport(srv *ProviderServer) *ProviderServerTransport {
	return &ProviderServerTransport{
		srv:                                     srv,
		newGetAvailableStacksPager:              newTracker[azfake.PagerResponder[armappservice.ProviderClientGetAvailableStacksResponse]](),
		newGetAvailableStacksOnPremPager:        newTracker[azfake.PagerResponder[armappservice.ProviderClientGetAvailableStacksOnPremResponse]](),
		newGetFunctionAppStacksPager:            newTracker[azfake.PagerResponder[armappservice.ProviderClientGetFunctionAppStacksResponse]](),
		newGetFunctionAppStacksForLocationPager: newTracker[azfake.PagerResponder[armappservice.ProviderClientGetFunctionAppStacksForLocationResponse]](),
		newGetWebAppStacksPager:                 newTracker[azfake.PagerResponder[armappservice.ProviderClientGetWebAppStacksResponse]](),
		newGetWebAppStacksForLocationPager:      newTracker[azfake.PagerResponder[armappservice.ProviderClientGetWebAppStacksForLocationResponse]](),
		newListOperationsPager:                  newTracker[azfake.PagerResponder[armappservice.ProviderClientListOperationsResponse]](),
	}
}

// ProviderServerTransport connects instances of armappservice.ProviderClient to instances of ProviderServer.
// Don't use this type directly, use NewProviderServerTransport instead.
type ProviderServerTransport struct {
	srv                                     *ProviderServer
	newGetAvailableStacksPager              *tracker[azfake.PagerResponder[armappservice.ProviderClientGetAvailableStacksResponse]]
	newGetAvailableStacksOnPremPager        *tracker[azfake.PagerResponder[armappservice.ProviderClientGetAvailableStacksOnPremResponse]]
	newGetFunctionAppStacksPager            *tracker[azfake.PagerResponder[armappservice.ProviderClientGetFunctionAppStacksResponse]]
	newGetFunctionAppStacksForLocationPager *tracker[azfake.PagerResponder[armappservice.ProviderClientGetFunctionAppStacksForLocationResponse]]
	newGetWebAppStacksPager                 *tracker[azfake.PagerResponder[armappservice.ProviderClientGetWebAppStacksResponse]]
	newGetWebAppStacksForLocationPager      *tracker[azfake.PagerResponder[armappservice.ProviderClientGetWebAppStacksForLocationResponse]]
	newListOperationsPager                  *tracker[azfake.PagerResponder[armappservice.ProviderClientListOperationsResponse]]
}

// Do implements the policy.Transporter interface for ProviderServerTransport.
func (p *ProviderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProviderServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if providerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = providerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProviderClient.NewGetAvailableStacksPager":
				res.resp, res.err = p.dispatchNewGetAvailableStacksPager(req)
			case "ProviderClient.NewGetAvailableStacksOnPremPager":
				res.resp, res.err = p.dispatchNewGetAvailableStacksOnPremPager(req)
			case "ProviderClient.NewGetFunctionAppStacksPager":
				res.resp, res.err = p.dispatchNewGetFunctionAppStacksPager(req)
			case "ProviderClient.NewGetFunctionAppStacksForLocationPager":
				res.resp, res.err = p.dispatchNewGetFunctionAppStacksForLocationPager(req)
			case "ProviderClient.NewGetWebAppStacksPager":
				res.resp, res.err = p.dispatchNewGetWebAppStacksPager(req)
			case "ProviderClient.NewGetWebAppStacksForLocationPager":
				res.resp, res.err = p.dispatchNewGetWebAppStacksForLocationPager(req)
			case "ProviderClient.NewListOperationsPager":
				res.resp, res.err = p.dispatchNewListOperationsPager(req)
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

func (p *ProviderServerTransport) dispatchNewGetAvailableStacksPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetAvailableStacksPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetAvailableStacksPager not implemented")}
	}
	newGetAvailableStacksPager := p.newGetAvailableStacksPager.get(req)
	if newGetAvailableStacksPager == nil {
		qp := req.URL.Query()
		oSTypeSelectedUnescaped, err := url.QueryUnescape(qp.Get("osTypeSelected"))
		if err != nil {
			return nil, err
		}
		oSTypeSelectedParam := getOptional(armappservice.ProviderOsTypeSelected(oSTypeSelectedUnescaped))
		var options *armappservice.ProviderClientGetAvailableStacksOptions
		if oSTypeSelectedParam != nil {
			options = &armappservice.ProviderClientGetAvailableStacksOptions{
				OSTypeSelected: oSTypeSelectedParam,
			}
		}
		resp := p.srv.NewGetAvailableStacksPager(options)
		newGetAvailableStacksPager = &resp
		p.newGetAvailableStacksPager.add(req, newGetAvailableStacksPager)
		server.PagerResponderInjectNextLinks(newGetAvailableStacksPager, req, func(page *armappservice.ProviderClientGetAvailableStacksResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetAvailableStacksPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetAvailableStacksPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetAvailableStacksPager) {
		p.newGetAvailableStacksPager.remove(req)
	}
	return resp, nil
}

func (p *ProviderServerTransport) dispatchNewGetAvailableStacksOnPremPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetAvailableStacksOnPremPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetAvailableStacksOnPremPager not implemented")}
	}
	newGetAvailableStacksOnPremPager := p.newGetAvailableStacksOnPremPager.get(req)
	if newGetAvailableStacksOnPremPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/availableStacks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		oSTypeSelectedUnescaped, err := url.QueryUnescape(qp.Get("osTypeSelected"))
		if err != nil {
			return nil, err
		}
		oSTypeSelectedParam := getOptional(armappservice.ProviderOsTypeSelected(oSTypeSelectedUnescaped))
		var options *armappservice.ProviderClientGetAvailableStacksOnPremOptions
		if oSTypeSelectedParam != nil {
			options = &armappservice.ProviderClientGetAvailableStacksOnPremOptions{
				OSTypeSelected: oSTypeSelectedParam,
			}
		}
		resp := p.srv.NewGetAvailableStacksOnPremPager(options)
		newGetAvailableStacksOnPremPager = &resp
		p.newGetAvailableStacksOnPremPager.add(req, newGetAvailableStacksOnPremPager)
		server.PagerResponderInjectNextLinks(newGetAvailableStacksOnPremPager, req, func(page *armappservice.ProviderClientGetAvailableStacksOnPremResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetAvailableStacksOnPremPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetAvailableStacksOnPremPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetAvailableStacksOnPremPager) {
		p.newGetAvailableStacksOnPremPager.remove(req)
	}
	return resp, nil
}

func (p *ProviderServerTransport) dispatchNewGetFunctionAppStacksPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetFunctionAppStacksPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetFunctionAppStacksPager not implemented")}
	}
	newGetFunctionAppStacksPager := p.newGetFunctionAppStacksPager.get(req)
	if newGetFunctionAppStacksPager == nil {
		qp := req.URL.Query()
		stackOsTypeUnescaped, err := url.QueryUnescape(qp.Get("stackOsType"))
		if err != nil {
			return nil, err
		}
		stackOsTypeParam := getOptional(armappservice.ProviderStackOsType(stackOsTypeUnescaped))
		var options *armappservice.ProviderClientGetFunctionAppStacksOptions
		if stackOsTypeParam != nil {
			options = &armappservice.ProviderClientGetFunctionAppStacksOptions{
				StackOsType: stackOsTypeParam,
			}
		}
		resp := p.srv.NewGetFunctionAppStacksPager(options)
		newGetFunctionAppStacksPager = &resp
		p.newGetFunctionAppStacksPager.add(req, newGetFunctionAppStacksPager)
		server.PagerResponderInjectNextLinks(newGetFunctionAppStacksPager, req, func(page *armappservice.ProviderClientGetFunctionAppStacksResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetFunctionAppStacksPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetFunctionAppStacksPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetFunctionAppStacksPager) {
		p.newGetFunctionAppStacksPager.remove(req)
	}
	return resp, nil
}

func (p *ProviderServerTransport) dispatchNewGetFunctionAppStacksForLocationPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetFunctionAppStacksForLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetFunctionAppStacksForLocationPager not implemented")}
	}
	newGetFunctionAppStacksForLocationPager := p.newGetFunctionAppStacksForLocationPager.get(req)
	if newGetFunctionAppStacksForLocationPager == nil {
		const regexStr = `/providers/Microsoft\.Web/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/functionAppStacks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		stackOsTypeUnescaped, err := url.QueryUnescape(qp.Get("stackOsType"))
		if err != nil {
			return nil, err
		}
		stackOsTypeParam := getOptional(armappservice.ProviderStackOsType(stackOsTypeUnescaped))
		var options *armappservice.ProviderClientGetFunctionAppStacksForLocationOptions
		if stackOsTypeParam != nil {
			options = &armappservice.ProviderClientGetFunctionAppStacksForLocationOptions{
				StackOsType: stackOsTypeParam,
			}
		}
		resp := p.srv.NewGetFunctionAppStacksForLocationPager(locationParam, options)
		newGetFunctionAppStacksForLocationPager = &resp
		p.newGetFunctionAppStacksForLocationPager.add(req, newGetFunctionAppStacksForLocationPager)
		server.PagerResponderInjectNextLinks(newGetFunctionAppStacksForLocationPager, req, func(page *armappservice.ProviderClientGetFunctionAppStacksForLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetFunctionAppStacksForLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetFunctionAppStacksForLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetFunctionAppStacksForLocationPager) {
		p.newGetFunctionAppStacksForLocationPager.remove(req)
	}
	return resp, nil
}

func (p *ProviderServerTransport) dispatchNewGetWebAppStacksPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetWebAppStacksPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetWebAppStacksPager not implemented")}
	}
	newGetWebAppStacksPager := p.newGetWebAppStacksPager.get(req)
	if newGetWebAppStacksPager == nil {
		qp := req.URL.Query()
		stackOsTypeUnescaped, err := url.QueryUnescape(qp.Get("stackOsType"))
		if err != nil {
			return nil, err
		}
		stackOsTypeParam := getOptional(armappservice.ProviderStackOsType(stackOsTypeUnescaped))
		var options *armappservice.ProviderClientGetWebAppStacksOptions
		if stackOsTypeParam != nil {
			options = &armappservice.ProviderClientGetWebAppStacksOptions{
				StackOsType: stackOsTypeParam,
			}
		}
		resp := p.srv.NewGetWebAppStacksPager(options)
		newGetWebAppStacksPager = &resp
		p.newGetWebAppStacksPager.add(req, newGetWebAppStacksPager)
		server.PagerResponderInjectNextLinks(newGetWebAppStacksPager, req, func(page *armappservice.ProviderClientGetWebAppStacksResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetWebAppStacksPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetWebAppStacksPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetWebAppStacksPager) {
		p.newGetWebAppStacksPager.remove(req)
	}
	return resp, nil
}

func (p *ProviderServerTransport) dispatchNewGetWebAppStacksForLocationPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetWebAppStacksForLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetWebAppStacksForLocationPager not implemented")}
	}
	newGetWebAppStacksForLocationPager := p.newGetWebAppStacksForLocationPager.get(req)
	if newGetWebAppStacksForLocationPager == nil {
		const regexStr = `/providers/Microsoft\.Web/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppStacks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		stackOsTypeUnescaped, err := url.QueryUnescape(qp.Get("stackOsType"))
		if err != nil {
			return nil, err
		}
		stackOsTypeParam := getOptional(armappservice.ProviderStackOsType(stackOsTypeUnescaped))
		var options *armappservice.ProviderClientGetWebAppStacksForLocationOptions
		if stackOsTypeParam != nil {
			options = &armappservice.ProviderClientGetWebAppStacksForLocationOptions{
				StackOsType: stackOsTypeParam,
			}
		}
		resp := p.srv.NewGetWebAppStacksForLocationPager(locationParam, options)
		newGetWebAppStacksForLocationPager = &resp
		p.newGetWebAppStacksForLocationPager.add(req, newGetWebAppStacksForLocationPager)
		server.PagerResponderInjectNextLinks(newGetWebAppStacksForLocationPager, req, func(page *armappservice.ProviderClientGetWebAppStacksForLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetWebAppStacksForLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetWebAppStacksForLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetWebAppStacksForLocationPager) {
		p.newGetWebAppStacksForLocationPager.remove(req)
	}
	return resp, nil
}

func (p *ProviderServerTransport) dispatchNewListOperationsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListOperationsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListOperationsPager not implemented")}
	}
	newListOperationsPager := p.newListOperationsPager.get(req)
	if newListOperationsPager == nil {
		resp := p.srv.NewListOperationsPager(nil)
		newListOperationsPager = &resp
		p.newListOperationsPager.add(req, newListOperationsPager)
		server.PagerResponderInjectNextLinks(newListOperationsPager, req, func(page *armappservice.ProviderClientListOperationsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListOperationsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListOperationsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListOperationsPager) {
		p.newListOperationsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ProviderServerTransport
var providerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
