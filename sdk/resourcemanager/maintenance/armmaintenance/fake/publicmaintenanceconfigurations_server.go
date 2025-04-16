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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"net/http"
	"net/url"
	"regexp"
)

// PublicMaintenanceConfigurationsServer is a fake server for instances of the armmaintenance.PublicMaintenanceConfigurationsClient type.
type PublicMaintenanceConfigurationsServer struct {
	// Get is the fake for method PublicMaintenanceConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceName string, options *armmaintenance.PublicMaintenanceConfigurationsClientGetOptions) (resp azfake.Responder[armmaintenance.PublicMaintenanceConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PublicMaintenanceConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armmaintenance.PublicMaintenanceConfigurationsClientListOptions) (resp azfake.PagerResponder[armmaintenance.PublicMaintenanceConfigurationsClientListResponse])
}

// NewPublicMaintenanceConfigurationsServerTransport creates a new instance of PublicMaintenanceConfigurationsServerTransport with the provided implementation.
// The returned PublicMaintenanceConfigurationsServerTransport instance is connected to an instance of armmaintenance.PublicMaintenanceConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPublicMaintenanceConfigurationsServerTransport(srv *PublicMaintenanceConfigurationsServer) *PublicMaintenanceConfigurationsServerTransport {
	return &PublicMaintenanceConfigurationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmaintenance.PublicMaintenanceConfigurationsClientListResponse]](),
	}
}

// PublicMaintenanceConfigurationsServerTransport connects instances of armmaintenance.PublicMaintenanceConfigurationsClient to instances of PublicMaintenanceConfigurationsServer.
// Don't use this type directly, use NewPublicMaintenanceConfigurationsServerTransport instead.
type PublicMaintenanceConfigurationsServerTransport struct {
	srv          *PublicMaintenanceConfigurationsServer
	newListPager *tracker[azfake.PagerResponder[armmaintenance.PublicMaintenanceConfigurationsClientListResponse]]
}

// Do implements the policy.Transporter interface for PublicMaintenanceConfigurationsServerTransport.
func (p *PublicMaintenanceConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PublicMaintenanceConfigurationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if publicMaintenanceConfigurationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = publicMaintenanceConfigurationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PublicMaintenanceConfigurationsClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PublicMaintenanceConfigurationsClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
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

func (p *PublicMaintenanceConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/publicMaintenanceConfigurations/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Configuration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PublicMaintenanceConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/publicMaintenanceConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListPager(nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PublicMaintenanceConfigurationsServerTransport
var publicMaintenanceConfigurationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
