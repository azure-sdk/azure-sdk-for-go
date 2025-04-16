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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
	"net/http"
	"net/url"
	"regexp"
)

// OuContainerServer is a fake server for instances of the armdomainservices.OuContainerClient type.
type OuContainerServer struct {
	// BeginCreate is the fake for method OuContainerClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount armdomainservices.ContainerAccount, options *armdomainservices.OuContainerClientBeginCreateOptions) (resp azfake.PollerResponder[armdomainservices.OuContainerClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method OuContainerClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *armdomainservices.OuContainerClientBeginDeleteOptions) (resp azfake.PollerResponder[armdomainservices.OuContainerClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OuContainerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *armdomainservices.OuContainerClientGetOptions) (resp azfake.Responder[armdomainservices.OuContainerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method OuContainerClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, domainServiceName string, options *armdomainservices.OuContainerClientListOptions) (resp azfake.PagerResponder[armdomainservices.OuContainerClientListResponse])

	// BeginUpdate is the fake for method OuContainerClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount armdomainservices.ContainerAccount, options *armdomainservices.OuContainerClientBeginUpdateOptions) (resp azfake.PollerResponder[armdomainservices.OuContainerClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewOuContainerServerTransport creates a new instance of OuContainerServerTransport with the provided implementation.
// The returned OuContainerServerTransport instance is connected to an instance of armdomainservices.OuContainerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOuContainerServerTransport(srv *OuContainerServer) *OuContainerServerTransport {
	return &OuContainerServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armdomainservices.OuContainerClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armdomainservices.OuContainerClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armdomainservices.OuContainerClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armdomainservices.OuContainerClientUpdateResponse]](),
	}
}

// OuContainerServerTransport connects instances of armdomainservices.OuContainerClient to instances of OuContainerServer.
// Don't use this type directly, use NewOuContainerServerTransport instead.
type OuContainerServerTransport struct {
	srv          *OuContainerServer
	beginCreate  *tracker[azfake.PollerResponder[armdomainservices.OuContainerClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armdomainservices.OuContainerClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armdomainservices.OuContainerClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armdomainservices.OuContainerClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for OuContainerServerTransport.
func (o *OuContainerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OuContainerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if ouContainerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = ouContainerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OuContainerClient.BeginCreate":
				res.resp, res.err = o.dispatchBeginCreate(req)
			case "OuContainerClient.BeginDelete":
				res.resp, res.err = o.dispatchBeginDelete(req)
			case "OuContainerClient.Get":
				res.resp, res.err = o.dispatchGet(req)
			case "OuContainerClient.NewListPager":
				res.resp, res.err = o.dispatchNewListPager(req)
			case "OuContainerClient.BeginUpdate":
				res.resp, res.err = o.dispatchBeginUpdate(req)
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

func (o *OuContainerServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := o.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Aad/domainServices/(?P<domainServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ouContainer/(?P<ouContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdomainservices.ContainerAccount](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainServiceName")])
		if err != nil {
			return nil, err
		}
		ouContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ouContainerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginCreate(req.Context(), resourceGroupNameParam, domainServiceNameParam, ouContainerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		o.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		o.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		o.beginCreate.remove(req)
	}

	return resp, nil
}

func (o *OuContainerServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if o.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := o.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Aad/domainServices/(?P<domainServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ouContainer/(?P<ouContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainServiceName")])
		if err != nil {
			return nil, err
		}
		ouContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ouContainerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginDelete(req.Context(), resourceGroupNameParam, domainServiceNameParam, ouContainerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		o.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		o.beginDelete.remove(req)
	}

	return resp, nil
}

func (o *OuContainerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Aad/domainServices/(?P<domainServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ouContainer/(?P<ouContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	domainServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainServiceName")])
	if err != nil {
		return nil, err
	}
	ouContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ouContainerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, domainServiceNameParam, ouContainerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OuContainer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OuContainerServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := o.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Aad/domainServices/(?P<domainServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ouContainer`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainServiceName")])
		if err != nil {
			return nil, err
		}
		resp := o.srv.NewListPager(resourceGroupNameParam, domainServiceNameParam, nil)
		newListPager = &resp
		o.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdomainservices.OuContainerClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		o.newListPager.remove(req)
	}
	return resp, nil
}

func (o *OuContainerServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := o.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Aad/domainServices/(?P<domainServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ouContainer/(?P<ouContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdomainservices.ContainerAccount](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		domainServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainServiceName")])
		if err != nil {
			return nil, err
		}
		ouContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ouContainerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginUpdate(req.Context(), resourceGroupNameParam, domainServiceNameParam, ouContainerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		o.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		o.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		o.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to OuContainerServerTransport
var ouContainerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
