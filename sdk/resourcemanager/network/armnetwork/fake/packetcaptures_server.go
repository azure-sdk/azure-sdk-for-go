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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
)

// PacketCapturesServer is a fake server for instances of the armnetwork.PacketCapturesClient type.
type PacketCapturesServer struct {
	// BeginCreate is the fake for method PacketCapturesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters armnetwork.PacketCapture, options *armnetwork.PacketCapturesClientBeginCreateOptions) (resp azfake.PollerResponder[armnetwork.PacketCapturesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PacketCapturesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *armnetwork.PacketCapturesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.PacketCapturesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PacketCapturesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *armnetwork.PacketCapturesClientGetOptions) (resp azfake.Responder[armnetwork.PacketCapturesClientGetResponse], errResp azfake.ErrorResponder)

	// BeginGetStatus is the fake for method PacketCapturesClient.BeginGetStatus
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetStatus func(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *armnetwork.PacketCapturesClientBeginGetStatusOptions) (resp azfake.PollerResponder[armnetwork.PacketCapturesClientGetStatusResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PacketCapturesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkWatcherName string, options *armnetwork.PacketCapturesClientListOptions) (resp azfake.PagerResponder[armnetwork.PacketCapturesClientListResponse])

	// BeginStop is the fake for method PacketCapturesClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *armnetwork.PacketCapturesClientBeginStopOptions) (resp azfake.PollerResponder[armnetwork.PacketCapturesClientStopResponse], errResp azfake.ErrorResponder)
}

// NewPacketCapturesServerTransport creates a new instance of PacketCapturesServerTransport with the provided implementation.
// The returned PacketCapturesServerTransport instance is connected to an instance of armnetwork.PacketCapturesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPacketCapturesServerTransport(srv *PacketCapturesServer) *PacketCapturesServerTransport {
	return &PacketCapturesServerTransport{
		srv:            srv,
		beginCreate:    newTracker[azfake.PollerResponder[armnetwork.PacketCapturesClientCreateResponse]](),
		beginDelete:    newTracker[azfake.PollerResponder[armnetwork.PacketCapturesClientDeleteResponse]](),
		beginGetStatus: newTracker[azfake.PollerResponder[armnetwork.PacketCapturesClientGetStatusResponse]](),
		newListPager:   newTracker[azfake.PagerResponder[armnetwork.PacketCapturesClientListResponse]](),
		beginStop:      newTracker[azfake.PollerResponder[armnetwork.PacketCapturesClientStopResponse]](),
	}
}

// PacketCapturesServerTransport connects instances of armnetwork.PacketCapturesClient to instances of PacketCapturesServer.
// Don't use this type directly, use NewPacketCapturesServerTransport instead.
type PacketCapturesServerTransport struct {
	srv            *PacketCapturesServer
	beginCreate    *tracker[azfake.PollerResponder[armnetwork.PacketCapturesClientCreateResponse]]
	beginDelete    *tracker[azfake.PollerResponder[armnetwork.PacketCapturesClientDeleteResponse]]
	beginGetStatus *tracker[azfake.PollerResponder[armnetwork.PacketCapturesClientGetStatusResponse]]
	newListPager   *tracker[azfake.PagerResponder[armnetwork.PacketCapturesClientListResponse]]
	beginStop      *tracker[azfake.PollerResponder[armnetwork.PacketCapturesClientStopResponse]]
}

// Do implements the policy.Transporter interface for PacketCapturesServerTransport.
func (p *PacketCapturesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PacketCapturesClient.BeginCreate":
		resp, err = p.dispatchBeginCreate(req)
	case "PacketCapturesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PacketCapturesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PacketCapturesClient.BeginGetStatus":
		resp, err = p.dispatchBeginGetStatus(req)
	case "PacketCapturesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PacketCapturesClient.BeginStop":
		resp, err = p.dispatchBeginStop(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PacketCapturesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := p.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packetCaptures/(?P<packetCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PacketCapture](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		packetCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packetCaptureName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreate(req.Context(), resourceGroupNameParam, networkWatcherNameParam, packetCaptureNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		p.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		p.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		p.beginCreate.remove(req)
	}

	return resp, nil
}

func (p *PacketCapturesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packetCaptures/(?P<packetCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		packetCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packetCaptureName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkWatcherNameParam, packetCaptureNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PacketCapturesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packetCaptures/(?P<packetCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
	if err != nil {
		return nil, err
	}
	packetCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packetCaptureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, networkWatcherNameParam, packetCaptureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PacketCaptureResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PacketCapturesServerTransport) dispatchBeginGetStatus(req *http.Request) (*http.Response, error) {
	if p.srv.BeginGetStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetStatus not implemented")}
	}
	beginGetStatus := p.beginGetStatus.get(req)
	if beginGetStatus == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packetCaptures/(?P<packetCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryStatus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		packetCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packetCaptureName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginGetStatus(req.Context(), resourceGroupNameParam, networkWatcherNameParam, packetCaptureNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetStatus = &respr
		p.beginGetStatus.add(req, beginGetStatus)
	}

	resp, err := server.PollerResponderNext(beginGetStatus, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginGetStatus.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetStatus) {
		p.beginGetStatus.remove(req)
	}

	return resp, nil
}

func (p *PacketCapturesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packetCaptures`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, networkWatcherNameParam, nil)
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

func (p *PacketCapturesServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if p.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := p.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkWatchers/(?P<networkWatcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packetCaptures/(?P<packetCaptureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkWatcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkWatcherName")])
		if err != nil {
			return nil, err
		}
		packetCaptureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packetCaptureName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginStop(req.Context(), resourceGroupNameParam, networkWatcherNameParam, packetCaptureNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		p.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		p.beginStop.remove(req)
	}

	return resp, nil
}
