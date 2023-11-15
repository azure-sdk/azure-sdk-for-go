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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// SchemaRegistryServer is a fake server for instances of the armeventhub.SchemaRegistryClient type.
type SchemaRegistryServer struct {
	// CreateOrUpdate is the fake for method SchemaRegistryClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, parameters armeventhub.SchemaGroup, options *armeventhub.SchemaRegistryClientCreateOrUpdateOptions) (resp azfake.Responder[armeventhub.SchemaRegistryClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SchemaRegistryClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, options *armeventhub.SchemaRegistryClientDeleteOptions) (resp azfake.Responder[armeventhub.SchemaRegistryClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SchemaRegistryClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, schemaGroupName string, options *armeventhub.SchemaRegistryClientGetOptions) (resp azfake.Responder[armeventhub.SchemaRegistryClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByNamespacePager is the fake for method SchemaRegistryClient.NewListByNamespacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNamespacePager func(resourceGroupName string, namespaceName string, options *armeventhub.SchemaRegistryClientListByNamespaceOptions) (resp azfake.PagerResponder[armeventhub.SchemaRegistryClientListByNamespaceResponse])
}

// NewSchemaRegistryServerTransport creates a new instance of SchemaRegistryServerTransport with the provided implementation.
// The returned SchemaRegistryServerTransport instance is connected to an instance of armeventhub.SchemaRegistryClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSchemaRegistryServerTransport(srv *SchemaRegistryServer) *SchemaRegistryServerTransport {
	return &SchemaRegistryServerTransport{
		srv:                     srv,
		newListByNamespacePager: newTracker[azfake.PagerResponder[armeventhub.SchemaRegistryClientListByNamespaceResponse]](),
	}
}

// SchemaRegistryServerTransport connects instances of armeventhub.SchemaRegistryClient to instances of SchemaRegistryServer.
// Don't use this type directly, use NewSchemaRegistryServerTransport instead.
type SchemaRegistryServerTransport struct {
	srv                     *SchemaRegistryServer
	newListByNamespacePager *tracker[azfake.PagerResponder[armeventhub.SchemaRegistryClientListByNamespaceResponse]]
}

// Do implements the policy.Transporter interface for SchemaRegistryServerTransport.
func (s *SchemaRegistryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SchemaRegistryClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "SchemaRegistryClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "SchemaRegistryClient.Get":
		resp, err = s.dispatchGet(req)
	case "SchemaRegistryClient.NewListByNamespacePager":
		resp, err = s.dispatchNewListByNamespacePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SchemaRegistryServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemagroups/(?P<schemaGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armeventhub.SchemaGroup](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	schemaGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, schemaGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SchemaGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SchemaRegistryServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemagroups/(?P<schemaGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	schemaGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, namespaceNameParam, schemaGroupNameParam, nil)
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

func (s *SchemaRegistryServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemagroups/(?P<schemaGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	schemaGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, schemaGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SchemaGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SchemaRegistryServerTransport) dispatchNewListByNamespacePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByNamespacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNamespacePager not implemented")}
	}
	newListByNamespacePager := s.newListByNamespacePager.get(req)
	if newListByNamespacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemagroups`
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
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
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
		var options *armeventhub.SchemaRegistryClientListByNamespaceOptions
		if skipParam != nil || topParam != nil {
			options = &armeventhub.SchemaRegistryClientListByNamespaceOptions{
				Skip: skipParam,
				Top:  topParam,
			}
		}
		resp := s.srv.NewListByNamespacePager(resourceGroupNameParam, namespaceNameParam, options)
		newListByNamespacePager = &resp
		s.newListByNamespacePager.add(req, newListByNamespacePager)
		server.PagerResponderInjectNextLinks(newListByNamespacePager, req, func(page *armeventhub.SchemaRegistryClientListByNamespaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNamespacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByNamespacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNamespacePager) {
		s.newListByNamespacePager.remove(req)
	}
	return resp, nil
}
