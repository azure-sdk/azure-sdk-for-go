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

// GraphQLAPIResolverServer is a fake server for instances of the armapimanagement.GraphQLAPIResolverClient type.
type GraphQLAPIResolverServer struct {
	// CreateOrUpdate is the fake for method GraphQLAPIResolverClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, parameters armapimanagement.ResolverContract, options *armapimanagement.GraphQLAPIResolverClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.GraphQLAPIResolverClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method GraphQLAPIResolverClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, ifMatch string, options *armapimanagement.GraphQLAPIResolverClientDeleteOptions) (resp azfake.Responder[armapimanagement.GraphQLAPIResolverClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GraphQLAPIResolverClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, options *armapimanagement.GraphQLAPIResolverClientGetOptions) (resp azfake.Responder[armapimanagement.GraphQLAPIResolverClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method GraphQLAPIResolverClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, options *armapimanagement.GraphQLAPIResolverClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.GraphQLAPIResolverClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListByAPIPager is the fake for method GraphQLAPIResolverClient.NewListByAPIPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAPIPager func(resourceGroupName string, serviceName string, apiID string, options *armapimanagement.GraphQLAPIResolverClientListByAPIOptions) (resp azfake.PagerResponder[armapimanagement.GraphQLAPIResolverClientListByAPIResponse])

	// Update is the fake for method GraphQLAPIResolverClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, ifMatch string, parameters armapimanagement.ResolverUpdateContract, options *armapimanagement.GraphQLAPIResolverClientUpdateOptions) (resp azfake.Responder[armapimanagement.GraphQLAPIResolverClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGraphQLAPIResolverServerTransport creates a new instance of GraphQLAPIResolverServerTransport with the provided implementation.
// The returned GraphQLAPIResolverServerTransport instance is connected to an instance of armapimanagement.GraphQLAPIResolverClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGraphQLAPIResolverServerTransport(srv *GraphQLAPIResolverServer) *GraphQLAPIResolverServerTransport {
	return &GraphQLAPIResolverServerTransport{
		srv:               srv,
		newListByAPIPager: newTracker[azfake.PagerResponder[armapimanagement.GraphQLAPIResolverClientListByAPIResponse]](),
	}
}

// GraphQLAPIResolverServerTransport connects instances of armapimanagement.GraphQLAPIResolverClient to instances of GraphQLAPIResolverServer.
// Don't use this type directly, use NewGraphQLAPIResolverServerTransport instead.
type GraphQLAPIResolverServerTransport struct {
	srv               *GraphQLAPIResolverServer
	newListByAPIPager *tracker[azfake.PagerResponder[armapimanagement.GraphQLAPIResolverClientListByAPIResponse]]
}

// Do implements the policy.Transporter interface for GraphQLAPIResolverServerTransport.
func (g *GraphQLAPIResolverServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GraphQLAPIResolverClient.CreateOrUpdate":
		resp, err = g.dispatchCreateOrUpdate(req)
	case "GraphQLAPIResolverClient.Delete":
		resp, err = g.dispatchDelete(req)
	case "GraphQLAPIResolverClient.Get":
		resp, err = g.dispatchGet(req)
	case "GraphQLAPIResolverClient.GetEntityTag":
		resp, err = g.dispatchGetEntityTag(req)
	case "GraphQLAPIResolverClient.NewListByAPIPager":
		resp, err = g.dispatchNewListByAPIPager(req)
	case "GraphQLAPIResolverClient.Update":
		resp, err = g.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GraphQLAPIResolverServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resolvers/(?P<resolverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.ResolverContract](req)
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	resolverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resolverId")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.GraphQLAPIResolverClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.GraphQLAPIResolverClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := g.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, resolverIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResolverContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (g *GraphQLAPIResolverServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if g.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resolvers/(?P<resolverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	resolverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resolverId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, resolverIDParam, getHeaderValue(req.Header, "If-Match"), nil)
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

func (g *GraphQLAPIResolverServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resolvers/(?P<resolverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	resolverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resolverId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, resolverIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResolverContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (g *GraphQLAPIResolverServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if g.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resolvers/(?P<resolverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	resolverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resolverId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, resolverIDParam, nil)
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

func (g *GraphQLAPIResolverServerTransport) dispatchNewListByAPIPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByAPIPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAPIPager not implemented")}
	}
	newListByAPIPager := g.newListByAPIPager.get(req)
	if newListByAPIPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resolvers`
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
		apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
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
		var options *armapimanagement.GraphQLAPIResolverClientListByAPIOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.GraphQLAPIResolverClientListByAPIOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := g.srv.NewListByAPIPager(resourceGroupNameParam, serviceNameParam, apiIDParam, options)
		newListByAPIPager = &resp
		g.newListByAPIPager.add(req, newListByAPIPager)
		server.PagerResponderInjectNextLinks(newListByAPIPager, req, func(page *armapimanagement.GraphQLAPIResolverClientListByAPIResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAPIPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByAPIPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAPIPager) {
		g.newListByAPIPager.remove(req)
	}
	return resp, nil
}

func (g *GraphQLAPIResolverServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resolvers/(?P<resolverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.ResolverUpdateContract](req)
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	resolverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resolverId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Update(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, resolverIDParam, getHeaderValue(req.Header, "If-Match"), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResolverContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}
