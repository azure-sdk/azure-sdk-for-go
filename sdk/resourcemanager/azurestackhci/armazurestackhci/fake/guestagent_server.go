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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// GuestAgentServer is a fake server for instances of the armazurestackhci.GuestAgentClient type.
type GuestAgentServer struct {
	// BeginCreate is the fake for method GuestAgentClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceURI string, options *armazurestackhci.GuestAgentClientBeginCreateOptions) (resp azfake.PollerResponder[armazurestackhci.GuestAgentClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GuestAgentClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceURI string, options *armazurestackhci.GuestAgentClientBeginDeleteOptions) (resp azfake.PollerResponder[armazurestackhci.GuestAgentClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GuestAgentClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, options *armazurestackhci.GuestAgentClientGetOptions) (resp azfake.Responder[armazurestackhci.GuestAgentClientGetResponse], errResp azfake.ErrorResponder)
}

// NewGuestAgentServerTransport creates a new instance of GuestAgentServerTransport with the provided implementation.
// The returned GuestAgentServerTransport instance is connected to an instance of armazurestackhci.GuestAgentClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGuestAgentServerTransport(srv *GuestAgentServer) *GuestAgentServerTransport {
	return &GuestAgentServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armazurestackhci.GuestAgentClientCreateResponse]](),
		beginDelete: newTracker[azfake.PollerResponder[armazurestackhci.GuestAgentClientDeleteResponse]](),
	}
}

// GuestAgentServerTransport connects instances of armazurestackhci.GuestAgentClient to instances of GuestAgentServer.
// Don't use this type directly, use NewGuestAgentServerTransport instead.
type GuestAgentServerTransport struct {
	srv         *GuestAgentServer
	beginCreate *tracker[azfake.PollerResponder[armazurestackhci.GuestAgentClientCreateResponse]]
	beginDelete *tracker[azfake.PollerResponder[armazurestackhci.GuestAgentClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for GuestAgentServerTransport.
func (g *GuestAgentServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GuestAgentClient.BeginCreate":
		resp, err = g.dispatchBeginCreate(req)
	case "GuestAgentClient.BeginDelete":
		resp, err = g.dispatchBeginDelete(req)
	case "GuestAgentClient.Get":
		resp, err = g.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GuestAgentServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := g.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/virtualMachineInstances/default/guestAgents/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurestackhci.GuestAgent](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		var options *armazurestackhci.GuestAgentClientBeginCreateOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armazurestackhci.GuestAgentClientBeginCreateOptions{
				Body: &body,
			}
		}
		respr, errRespr := g.srv.BeginCreate(req.Context(), resourceURIParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		g.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		g.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		g.beginCreate.remove(req)
	}

	return resp, nil
}

func (g *GuestAgentServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := g.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/virtualMachineInstances/default/guestAgents/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), resourceURIParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		g.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		g.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		g.beginDelete.remove(req)
	}

	return resp, nil
}

func (g *GuestAgentServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/virtualMachineInstances/default/guestAgents/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceURIParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GuestAgent, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
