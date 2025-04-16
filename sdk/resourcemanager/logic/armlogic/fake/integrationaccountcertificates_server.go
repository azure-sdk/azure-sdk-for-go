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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// IntegrationAccountCertificatesServer is a fake server for instances of the armlogic.IntegrationAccountCertificatesClient type.
type IntegrationAccountCertificatesServer struct {
	// CreateOrUpdate is the fake for method IntegrationAccountCertificatesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, certificate armlogic.IntegrationAccountCertificate, options *armlogic.IntegrationAccountCertificatesClientCreateOrUpdateOptions) (resp azfake.Responder[armlogic.IntegrationAccountCertificatesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method IntegrationAccountCertificatesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, options *armlogic.IntegrationAccountCertificatesClientDeleteOptions) (resp azfake.Responder[armlogic.IntegrationAccountCertificatesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IntegrationAccountCertificatesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, options *armlogic.IntegrationAccountCertificatesClientGetOptions) (resp azfake.Responder[armlogic.IntegrationAccountCertificatesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IntegrationAccountCertificatesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, integrationAccountName string, options *armlogic.IntegrationAccountCertificatesClientListOptions) (resp azfake.PagerResponder[armlogic.IntegrationAccountCertificatesClientListResponse])
}

// NewIntegrationAccountCertificatesServerTransport creates a new instance of IntegrationAccountCertificatesServerTransport with the provided implementation.
// The returned IntegrationAccountCertificatesServerTransport instance is connected to an instance of armlogic.IntegrationAccountCertificatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntegrationAccountCertificatesServerTransport(srv *IntegrationAccountCertificatesServer) *IntegrationAccountCertificatesServerTransport {
	return &IntegrationAccountCertificatesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armlogic.IntegrationAccountCertificatesClientListResponse]](),
	}
}

// IntegrationAccountCertificatesServerTransport connects instances of armlogic.IntegrationAccountCertificatesClient to instances of IntegrationAccountCertificatesServer.
// Don't use this type directly, use NewIntegrationAccountCertificatesServerTransport instead.
type IntegrationAccountCertificatesServerTransport struct {
	srv          *IntegrationAccountCertificatesServer
	newListPager *tracker[azfake.PagerResponder[armlogic.IntegrationAccountCertificatesClientListResponse]]
}

// Do implements the policy.Transporter interface for IntegrationAccountCertificatesServerTransport.
func (i *IntegrationAccountCertificatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IntegrationAccountCertificatesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if integrationAccountCertificatesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = integrationAccountCertificatesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "IntegrationAccountCertificatesClient.CreateOrUpdate":
				res.resp, res.err = i.dispatchCreateOrUpdate(req)
			case "IntegrationAccountCertificatesClient.Delete":
				res.resp, res.err = i.dispatchDelete(req)
			case "IntegrationAccountCertificatesClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "IntegrationAccountCertificatesClient.NewListPager":
				res.resp, res.err = i.dispatchNewListPager(req)
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

func (i *IntegrationAccountCertificatesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlogic.IntegrationAccountCertificate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, integrationAccountNameParam, certificateNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationAccountCertificate, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationAccountCertificatesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, integrationAccountNameParam, certificateNameParam, nil)
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

func (i *IntegrationAccountCertificatesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, integrationAccountNameParam, certificateNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationAccountCertificate, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationAccountCertificatesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/certificates`
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
		integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
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
		var options *armlogic.IntegrationAccountCertificatesClientListOptions
		if topParam != nil {
			options = &armlogic.IntegrationAccountCertificatesClientListOptions{
				Top: topParam,
			}
		}
		resp := i.srv.NewListPager(resourceGroupNameParam, integrationAccountNameParam, options)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlogic.IntegrationAccountCertificatesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to IntegrationAccountCertificatesServerTransport
var integrationAccountCertificatesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
