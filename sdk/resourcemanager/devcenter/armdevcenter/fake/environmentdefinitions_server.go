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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// EnvironmentDefinitionsServer is a fake server for instances of the armdevcenter.EnvironmentDefinitionsClient type.
type EnvironmentDefinitionsServer struct {
	// Get is the fake for method EnvironmentDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, environmentDefinitionName string, options *armdevcenter.EnvironmentDefinitionsClientGetOptions) (resp azfake.Responder[armdevcenter.EnvironmentDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetByProjectCatalog is the fake for method EnvironmentDefinitionsClient.GetByProjectCatalog
	// HTTP status codes to indicate success: http.StatusOK
	GetByProjectCatalog func(ctx context.Context, resourceGroupName string, projectName string, catalogName string, environmentDefinitionName string, options *armdevcenter.EnvironmentDefinitionsClientGetByProjectCatalogOptions) (resp azfake.Responder[armdevcenter.EnvironmentDefinitionsClientGetByProjectCatalogResponse], errResp azfake.ErrorResponder)

	// GetErrorDetails is the fake for method EnvironmentDefinitionsClient.GetErrorDetails
	// HTTP status codes to indicate success: http.StatusOK
	GetErrorDetails func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, environmentDefinitionName string, options *armdevcenter.EnvironmentDefinitionsClientGetErrorDetailsOptions) (resp azfake.Responder[armdevcenter.EnvironmentDefinitionsClientGetErrorDetailsResponse], errResp azfake.ErrorResponder)

	// NewListByCatalogPager is the fake for method EnvironmentDefinitionsClient.NewListByCatalogPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCatalogPager func(resourceGroupName string, devCenterName string, catalogName string, options *armdevcenter.EnvironmentDefinitionsClientListByCatalogOptions) (resp azfake.PagerResponder[armdevcenter.EnvironmentDefinitionsClientListByCatalogResponse])

	// NewListByProjectCatalogPager is the fake for method EnvironmentDefinitionsClient.NewListByProjectCatalogPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProjectCatalogPager func(resourceGroupName string, projectName string, catalogName string, options *armdevcenter.EnvironmentDefinitionsClientListByProjectCatalogOptions) (resp azfake.PagerResponder[armdevcenter.EnvironmentDefinitionsClientListByProjectCatalogResponse])
}

// NewEnvironmentDefinitionsServerTransport creates a new instance of EnvironmentDefinitionsServerTransport with the provided implementation.
// The returned EnvironmentDefinitionsServerTransport instance is connected to an instance of armdevcenter.EnvironmentDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEnvironmentDefinitionsServerTransport(srv *EnvironmentDefinitionsServer) *EnvironmentDefinitionsServerTransport {
	return &EnvironmentDefinitionsServerTransport{
		srv:                          srv,
		newListByCatalogPager:        newTracker[azfake.PagerResponder[armdevcenter.EnvironmentDefinitionsClientListByCatalogResponse]](),
		newListByProjectCatalogPager: newTracker[azfake.PagerResponder[armdevcenter.EnvironmentDefinitionsClientListByProjectCatalogResponse]](),
	}
}

// EnvironmentDefinitionsServerTransport connects instances of armdevcenter.EnvironmentDefinitionsClient to instances of EnvironmentDefinitionsServer.
// Don't use this type directly, use NewEnvironmentDefinitionsServerTransport instead.
type EnvironmentDefinitionsServerTransport struct {
	srv                          *EnvironmentDefinitionsServer
	newListByCatalogPager        *tracker[azfake.PagerResponder[armdevcenter.EnvironmentDefinitionsClientListByCatalogResponse]]
	newListByProjectCatalogPager *tracker[azfake.PagerResponder[armdevcenter.EnvironmentDefinitionsClientListByProjectCatalogResponse]]
}

// Do implements the policy.Transporter interface for EnvironmentDefinitionsServerTransport.
func (e *EnvironmentDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EnvironmentDefinitionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if environmentDefinitionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = environmentDefinitionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "EnvironmentDefinitionsClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "EnvironmentDefinitionsClient.GetByProjectCatalog":
				res.resp, res.err = e.dispatchGetByProjectCatalog(req)
			case "EnvironmentDefinitionsClient.GetErrorDetails":
				res.resp, res.err = e.dispatchGetErrorDetails(req)
			case "EnvironmentDefinitionsClient.NewListByCatalogPager":
				res.resp, res.err = e.dispatchNewListByCatalogPager(req)
			case "EnvironmentDefinitionsClient.NewListByProjectCatalogPager":
				res.resp, res.err = e.dispatchNewListByProjectCatalogPager(req)
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

func (e *EnvironmentDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environmentDefinitions/(?P<environmentDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	environmentDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, environmentDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnvironmentDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnvironmentDefinitionsServerTransport) dispatchGetByProjectCatalog(req *http.Request) (*http.Response, error) {
	if e.srv.GetByProjectCatalog == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByProjectCatalog not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environmentDefinitions/(?P<environmentDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	environmentDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetByProjectCatalog(req.Context(), resourceGroupNameParam, projectNameParam, catalogNameParam, environmentDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnvironmentDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnvironmentDefinitionsServerTransport) dispatchGetErrorDetails(req *http.Request) (*http.Response, error) {
	if e.srv.GetErrorDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetErrorDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environmentDefinitions/(?P<environmentDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getErrorDetails`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	environmentDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetErrorDetails(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, environmentDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CatalogResourceValidationErrorDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnvironmentDefinitionsServerTransport) dispatchNewListByCatalogPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByCatalogPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCatalogPager not implemented")}
	}
	newListByCatalogPager := e.newListByCatalogPager.get(req)
	if newListByCatalogPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environmentDefinitions`
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
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
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
		var options *armdevcenter.EnvironmentDefinitionsClientListByCatalogOptions
		if topParam != nil {
			options = &armdevcenter.EnvironmentDefinitionsClientListByCatalogOptions{
				Top: topParam,
			}
		}
		resp := e.srv.NewListByCatalogPager(resourceGroupNameParam, devCenterNameParam, catalogNameParam, options)
		newListByCatalogPager = &resp
		e.newListByCatalogPager.add(req, newListByCatalogPager)
		server.PagerResponderInjectNextLinks(newListByCatalogPager, req, func(page *armdevcenter.EnvironmentDefinitionsClientListByCatalogResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCatalogPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByCatalogPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCatalogPager) {
		e.newListByCatalogPager.remove(req)
	}
	return resp, nil
}

func (e *EnvironmentDefinitionsServerTransport) dispatchNewListByProjectCatalogPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByProjectCatalogPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProjectCatalogPager not implemented")}
	}
	newListByProjectCatalogPager := e.newListByProjectCatalogPager.get(req)
	if newListByProjectCatalogPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environmentDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByProjectCatalogPager(resourceGroupNameParam, projectNameParam, catalogNameParam, nil)
		newListByProjectCatalogPager = &resp
		e.newListByProjectCatalogPager.add(req, newListByProjectCatalogPager)
		server.PagerResponderInjectNextLinks(newListByProjectCatalogPager, req, func(page *armdevcenter.EnvironmentDefinitionsClientListByProjectCatalogResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProjectCatalogPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByProjectCatalogPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProjectCatalogPager) {
		e.newListByProjectCatalogPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to EnvironmentDefinitionsServerTransport
var environmentDefinitionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
