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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
	"net/http"
	"net/url"
	"regexp"
)

// MetadataServer is a fake server for instances of the armresourcehealth.MetadataClient type.
type MetadataServer struct {
	// GetEntity is the fake for method MetadataClient.GetEntity
	// HTTP status codes to indicate success: http.StatusOK
	GetEntity func(ctx context.Context, name string, options *armresourcehealth.MetadataClientGetEntityOptions) (resp azfake.Responder[armresourcehealth.MetadataClientGetEntityResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method MetadataClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armresourcehealth.MetadataClientListOptions) (resp azfake.PagerResponder[armresourcehealth.MetadataClientListResponse])
}

// NewMetadataServerTransport creates a new instance of MetadataServerTransport with the provided implementation.
// The returned MetadataServerTransport instance is connected to an instance of armresourcehealth.MetadataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMetadataServerTransport(srv *MetadataServer) *MetadataServerTransport {
	return &MetadataServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armresourcehealth.MetadataClientListResponse]](),
	}
}

// MetadataServerTransport connects instances of armresourcehealth.MetadataClient to instances of MetadataServer.
// Don't use this type directly, use NewMetadataServerTransport instead.
type MetadataServerTransport struct {
	srv          *MetadataServer
	newListPager *tracker[azfake.PagerResponder[armresourcehealth.MetadataClientListResponse]]
}

// Do implements the policy.Transporter interface for MetadataServerTransport.
func (m *MetadataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MetadataServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if metadataServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = metadataServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MetadataClient.GetEntity":
				res.resp, res.err = m.dispatchGetEntity(req)
			case "MetadataClient.NewListPager":
				res.resp, res.err = m.dispatchNewListPager(req)
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

func (m *MetadataServerTransport) dispatchGetEntity(req *http.Request) (*http.Response, error) {
	if m.srv.GetEntity == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntity not implemented")}
	}
	const regexStr = `/providers/Microsoft\.ResourceHealth/metadata/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.GetEntity(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MetadataEntity, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MetadataServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		resp := m.srv.NewListPager(nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armresourcehealth.MetadataClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to MetadataServerTransport
var metadataServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
