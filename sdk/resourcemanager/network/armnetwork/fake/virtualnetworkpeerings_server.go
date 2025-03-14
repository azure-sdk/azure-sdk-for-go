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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualNetworkPeeringsServer is a fake server for instances of the armnetwork.VirtualNetworkPeeringsClient type.
type VirtualNetworkPeeringsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualNetworkPeeringsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters armnetwork.VirtualNetworkPeering, options *armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualNetworkPeeringsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *armnetwork.VirtualNetworkPeeringsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualNetworkPeeringsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *armnetwork.VirtualNetworkPeeringsClientGetOptions) (resp azfake.Responder[armnetwork.VirtualNetworkPeeringsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualNetworkPeeringsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualNetworkName string, options *armnetwork.VirtualNetworkPeeringsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualNetworkPeeringsClientListResponse])
}

// NewVirtualNetworkPeeringsServerTransport creates a new instance of VirtualNetworkPeeringsServerTransport with the provided implementation.
// The returned VirtualNetworkPeeringsServerTransport instance is connected to an instance of armnetwork.VirtualNetworkPeeringsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualNetworkPeeringsServerTransport(srv *VirtualNetworkPeeringsServer) *VirtualNetworkPeeringsServerTransport {
	return &VirtualNetworkPeeringsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.VirtualNetworkPeeringsClientListResponse]](),
	}
}

// VirtualNetworkPeeringsServerTransport connects instances of armnetwork.VirtualNetworkPeeringsClient to instances of VirtualNetworkPeeringsServer.
// Don't use this type directly, use NewVirtualNetworkPeeringsServerTransport instead.
type VirtualNetworkPeeringsServerTransport struct {
	srv                 *VirtualNetworkPeeringsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.VirtualNetworkPeeringsClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualNetworkPeeringsServerTransport.
func (v *VirtualNetworkPeeringsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if virtualNetworkPeeringsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = virtualNetworkPeeringsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VirtualNetworkPeeringsClient.BeginCreateOrUpdate":
				res.resp, res.err = v.dispatchBeginCreateOrUpdate(req)
			case "VirtualNetworkPeeringsClient.BeginDelete":
				res.resp, res.err = v.dispatchBeginDelete(req)
			case "VirtualNetworkPeeringsClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VirtualNetworkPeeringsClient.NewListPager":
				res.resp, res.err = v.dispatchNewListPager(req)
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

func (v *VirtualNetworkPeeringsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings/(?P<virtualNetworkPeeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualNetworkPeering](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkPeeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkPeeringName")])
		if err != nil {
			return nil, err
		}
		syncRemoteAddressSpaceUnescaped, err := url.QueryUnescape(qp.Get("syncRemoteAddressSpace"))
		if err != nil {
			return nil, err
		}
		syncRemoteAddressSpaceParam := getOptional(armnetwork.SyncRemoteAddressSpace(syncRemoteAddressSpaceUnescaped))
		var options *armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions
		if syncRemoteAddressSpaceParam != nil {
			options = &armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions{
				SyncRemoteAddressSpace: syncRemoteAddressSpaceParam,
			}
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, virtualNetworkNameParam, virtualNetworkPeeringNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings/(?P<virtualNetworkPeeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkPeeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkPeeringName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, virtualNetworkNameParam, virtualNetworkPeeringNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings/(?P<virtualNetworkPeeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkPeeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkPeeringName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, virtualNetworkNameParam, virtualNetworkPeeringNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetworkPeering, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, virtualNetworkNameParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualNetworkPeeringsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to VirtualNetworkPeeringsServerTransport
var virtualNetworkPeeringsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
