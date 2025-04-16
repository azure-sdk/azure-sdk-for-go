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

// OpenIDConnectProviderServer is a fake server for instances of the armapimanagement.OpenIDConnectProviderClient type.
type OpenIDConnectProviderServer struct {
	// CreateOrUpdate is the fake for method OpenIDConnectProviderClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, opid string, parameters armapimanagement.OpenidConnectProviderContract, options *armapimanagement.OpenIDConnectProviderClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.OpenIDConnectProviderClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method OpenIDConnectProviderClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, opid string, ifMatch string, options *armapimanagement.OpenIDConnectProviderClientDeleteOptions) (resp azfake.Responder[armapimanagement.OpenIDConnectProviderClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OpenIDConnectProviderClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *armapimanagement.OpenIDConnectProviderClientGetOptions) (resp azfake.Responder[armapimanagement.OpenIDConnectProviderClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method OpenIDConnectProviderClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *armapimanagement.OpenIDConnectProviderClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.OpenIDConnectProviderClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method OpenIDConnectProviderClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, options *armapimanagement.OpenIDConnectProviderClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.OpenIDConnectProviderClientListByServiceResponse])

	// ListSecrets is the fake for method OpenIDConnectProviderClient.ListSecrets
	// HTTP status codes to indicate success: http.StatusOK
	ListSecrets func(ctx context.Context, resourceGroupName string, serviceName string, opid string, options *armapimanagement.OpenIDConnectProviderClientListSecretsOptions) (resp azfake.Responder[armapimanagement.OpenIDConnectProviderClientListSecretsResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method OpenIDConnectProviderClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, serviceName string, opid string, ifMatch string, parameters armapimanagement.OpenidConnectProviderUpdateContract, options *armapimanagement.OpenIDConnectProviderClientUpdateOptions) (resp azfake.Responder[armapimanagement.OpenIDConnectProviderClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewOpenIDConnectProviderServerTransport creates a new instance of OpenIDConnectProviderServerTransport with the provided implementation.
// The returned OpenIDConnectProviderServerTransport instance is connected to an instance of armapimanagement.OpenIDConnectProviderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOpenIDConnectProviderServerTransport(srv *OpenIDConnectProviderServer) *OpenIDConnectProviderServerTransport {
	return &OpenIDConnectProviderServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.OpenIDConnectProviderClientListByServiceResponse]](),
	}
}

// OpenIDConnectProviderServerTransport connects instances of armapimanagement.OpenIDConnectProviderClient to instances of OpenIDConnectProviderServer.
// Don't use this type directly, use NewOpenIDConnectProviderServerTransport instead.
type OpenIDConnectProviderServerTransport struct {
	srv                   *OpenIDConnectProviderServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.OpenIDConnectProviderClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for OpenIDConnectProviderServerTransport.
func (o *OpenIDConnectProviderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OpenIDConnectProviderServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if openIdConnectProviderServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = openIdConnectProviderServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OpenIDConnectProviderClient.CreateOrUpdate":
				res.resp, res.err = o.dispatchCreateOrUpdate(req)
			case "OpenIDConnectProviderClient.Delete":
				res.resp, res.err = o.dispatchDelete(req)
			case "OpenIDConnectProviderClient.Get":
				res.resp, res.err = o.dispatchGet(req)
			case "OpenIDConnectProviderClient.GetEntityTag":
				res.resp, res.err = o.dispatchGetEntityTag(req)
			case "OpenIDConnectProviderClient.NewListByServicePager":
				res.resp, res.err = o.dispatchNewListByServicePager(req)
			case "OpenIDConnectProviderClient.ListSecrets":
				res.resp, res.err = o.dispatchListSecrets(req)
			case "OpenIDConnectProviderClient.Update":
				res.resp, res.err = o.dispatchUpdate(req)
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

func (o *OpenIDConnectProviderServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/openidConnectProviders/(?P<opid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.OpenidConnectProviderContract](req)
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
	opidParam, err := url.PathUnescape(matches[regex.SubexpIndex("opid")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.OpenIDConnectProviderClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.OpenIDConnectProviderClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := o.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, opidParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OpenidConnectProviderContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (o *OpenIDConnectProviderServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if o.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/openidConnectProviders/(?P<opid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	opidParam, err := url.PathUnescape(matches[regex.SubexpIndex("opid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, opidParam, getHeaderValue(req.Header, "If-Match"), nil)
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

func (o *OpenIDConnectProviderServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/openidConnectProviders/(?P<opid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	opidParam, err := url.PathUnescape(matches[regex.SubexpIndex("opid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, opidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OpenidConnectProviderContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (o *OpenIDConnectProviderServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if o.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/openidConnectProviders/(?P<opid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	opidParam, err := url.PathUnescape(matches[regex.SubexpIndex("opid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, opidParam, nil)
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

func (o *OpenIDConnectProviderServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := o.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/openidConnectProviders`
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
		var options *armapimanagement.OpenIDConnectProviderClientListByServiceOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.OpenIDConnectProviderClientListByServiceOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := o.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, options)
		newListByServicePager = &resp
		o.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.OpenIDConnectProviderClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		o.newListByServicePager.remove(req)
	}
	return resp, nil
}

func (o *OpenIDConnectProviderServerTransport) dispatchListSecrets(req *http.Request) (*http.Response, error) {
	if o.srv.ListSecrets == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSecrets not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/openidConnectProviders/(?P<opid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSecrets`
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
	opidParam, err := url.PathUnescape(matches[regex.SubexpIndex("opid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.ListSecrets(req.Context(), resourceGroupNameParam, serviceNameParam, opidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ClientSecretContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (o *OpenIDConnectProviderServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/openidConnectProviders/(?P<opid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.OpenidConnectProviderUpdateContract](req)
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
	opidParam, err := url.PathUnescape(matches[regex.SubexpIndex("opid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Update(req.Context(), resourceGroupNameParam, serviceNameParam, opidParam, getHeaderValue(req.Header, "If-Match"), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OpenidConnectProviderContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to OpenIDConnectProviderServerTransport
var openIdConnectProviderServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
