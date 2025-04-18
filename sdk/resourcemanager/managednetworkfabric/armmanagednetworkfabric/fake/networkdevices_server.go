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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkDevicesServer is a fake server for instances of the armmanagednetworkfabric.NetworkDevicesClient type.
type NetworkDevicesServer struct {
	// BeginCreate is the fake for method NetworkDevicesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, networkDeviceName string, body armmanagednetworkfabric.NetworkDevice, options *armmanagednetworkfabric.NetworkDevicesClientBeginCreateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NetworkDevicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkDeviceName string, options *armmanagednetworkfabric.NetworkDevicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkDevicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkDeviceName string, options *armmanagednetworkfabric.NetworkDevicesClientGetOptions) (resp azfake.Responder[armmanagednetworkfabric.NetworkDevicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method NetworkDevicesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmanagednetworkfabric.NetworkDevicesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmanagednetworkfabric.NetworkDevicesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method NetworkDevicesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmanagednetworkfabric.NetworkDevicesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmanagednetworkfabric.NetworkDevicesClientListBySubscriptionResponse])

	// BeginReboot is the fake for method NetworkDevicesClient.BeginReboot
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReboot func(ctx context.Context, resourceGroupName string, networkDeviceName string, body armmanagednetworkfabric.RebootProperties, options *armmanagednetworkfabric.NetworkDevicesClientBeginRebootOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientRebootResponse], errResp azfake.ErrorResponder)

	// BeginRefreshConfiguration is the fake for method NetworkDevicesClient.BeginRefreshConfiguration
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRefreshConfiguration func(ctx context.Context, resourceGroupName string, networkDeviceName string, options *armmanagednetworkfabric.NetworkDevicesClientBeginRefreshConfigurationOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientRefreshConfigurationResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method NetworkDevicesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, networkDeviceName string, body armmanagednetworkfabric.NetworkDevicePatchParameters, options *armmanagednetworkfabric.NetworkDevicesClientBeginUpdateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateAdministrativeState is the fake for method NetworkDevicesClient.BeginUpdateAdministrativeState
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateAdministrativeState func(ctx context.Context, resourceGroupName string, networkDeviceName string, body armmanagednetworkfabric.UpdateDeviceAdministrativeState, options *armmanagednetworkfabric.NetworkDevicesClientBeginUpdateAdministrativeStateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpdateAdministrativeStateResponse], errResp azfake.ErrorResponder)

	// BeginUpgrade is the fake for method NetworkDevicesClient.BeginUpgrade
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpgrade func(ctx context.Context, resourceGroupName string, networkDeviceName string, body armmanagednetworkfabric.UpdateVersion, options *armmanagednetworkfabric.NetworkDevicesClientBeginUpgradeOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpgradeResponse], errResp azfake.ErrorResponder)
}

// NewNetworkDevicesServerTransport creates a new instance of NetworkDevicesServerTransport with the provided implementation.
// The returned NetworkDevicesServerTransport instance is connected to an instance of armmanagednetworkfabric.NetworkDevicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkDevicesServerTransport(srv *NetworkDevicesServer) *NetworkDevicesServerTransport {
	return &NetworkDevicesServerTransport{
		srv:                            srv,
		beginCreate:                    newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientCreateResponse]](),
		beginDelete:                    newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientDeleteResponse]](),
		newListByResourceGroupPager:    newTracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkDevicesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:     newTracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkDevicesClientListBySubscriptionResponse]](),
		beginReboot:                    newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientRebootResponse]](),
		beginRefreshConfiguration:      newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientRefreshConfigurationResponse]](),
		beginUpdate:                    newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpdateResponse]](),
		beginUpdateAdministrativeState: newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpdateAdministrativeStateResponse]](),
		beginUpgrade:                   newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpgradeResponse]](),
	}
}

// NetworkDevicesServerTransport connects instances of armmanagednetworkfabric.NetworkDevicesClient to instances of NetworkDevicesServer.
// Don't use this type directly, use NewNetworkDevicesServerTransport instead.
type NetworkDevicesServerTransport struct {
	srv                            *NetworkDevicesServer
	beginCreate                    *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientCreateResponse]]
	beginDelete                    *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientDeleteResponse]]
	newListByResourceGroupPager    *tracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkDevicesClientListByResourceGroupResponse]]
	newListBySubscriptionPager     *tracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkDevicesClientListBySubscriptionResponse]]
	beginReboot                    *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientRebootResponse]]
	beginRefreshConfiguration      *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientRefreshConfigurationResponse]]
	beginUpdate                    *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpdateResponse]]
	beginUpdateAdministrativeState *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpdateAdministrativeStateResponse]]
	beginUpgrade                   *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkDevicesClientUpgradeResponse]]
}

// Do implements the policy.Transporter interface for NetworkDevicesServerTransport.
func (n *NetworkDevicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NetworkDevicesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if networkDevicesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = networkDevicesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "NetworkDevicesClient.BeginCreate":
				res.resp, res.err = n.dispatchBeginCreate(req)
			case "NetworkDevicesClient.BeginDelete":
				res.resp, res.err = n.dispatchBeginDelete(req)
			case "NetworkDevicesClient.Get":
				res.resp, res.err = n.dispatchGet(req)
			case "NetworkDevicesClient.NewListByResourceGroupPager":
				res.resp, res.err = n.dispatchNewListByResourceGroupPager(req)
			case "NetworkDevicesClient.NewListBySubscriptionPager":
				res.resp, res.err = n.dispatchNewListBySubscriptionPager(req)
			case "NetworkDevicesClient.BeginReboot":
				res.resp, res.err = n.dispatchBeginReboot(req)
			case "NetworkDevicesClient.BeginRefreshConfiguration":
				res.resp, res.err = n.dispatchBeginRefreshConfiguration(req)
			case "NetworkDevicesClient.BeginUpdate":
				res.resp, res.err = n.dispatchBeginUpdate(req)
			case "NetworkDevicesClient.BeginUpdateAdministrativeState":
				res.resp, res.err = n.dispatchBeginUpdateAdministrativeState(req)
			case "NetworkDevicesClient.BeginUpgrade":
				res.resp, res.err = n.dispatchBeginUpgrade(req)
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

func (n *NetworkDevicesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := n.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.NetworkDevice](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreate(req.Context(), resourceGroupNameParam, networkDeviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		n.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		n.beginCreate.remove(req)
	}

	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkDeviceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, networkDeviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkDevice, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := n.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		n.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmanagednetworkfabric.NetworkDevicesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		n.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := n.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := n.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		n.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmanagednetworkfabric.NetworkDevicesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		n.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchBeginReboot(req *http.Request) (*http.Response, error) {
	if n.srv.BeginReboot == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReboot not implemented")}
	}
	beginReboot := n.beginReboot.get(req)
	if beginReboot == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reboot`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.RebootProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginReboot(req.Context(), resourceGroupNameParam, networkDeviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReboot = &respr
		n.beginReboot.add(req, beginReboot)
	}

	resp, err := server.PollerResponderNext(beginReboot, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginReboot.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReboot) {
		n.beginReboot.remove(req)
	}

	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchBeginRefreshConfiguration(req *http.Request) (*http.Response, error) {
	if n.srv.BeginRefreshConfiguration == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRefreshConfiguration not implemented")}
	}
	beginRefreshConfiguration := n.beginRefreshConfiguration.get(req)
	if beginRefreshConfiguration == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshConfiguration`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginRefreshConfiguration(req.Context(), resourceGroupNameParam, networkDeviceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRefreshConfiguration = &respr
		n.beginRefreshConfiguration.add(req, beginRefreshConfiguration)
	}

	resp, err := server.PollerResponderNext(beginRefreshConfiguration, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginRefreshConfiguration.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRefreshConfiguration) {
		n.beginRefreshConfiguration.remove(req)
	}

	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := n.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.NetworkDevicePatchParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdate(req.Context(), resourceGroupNameParam, networkDeviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		n.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		n.beginUpdate.remove(req)
	}

	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchBeginUpdateAdministrativeState(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdateAdministrativeState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateAdministrativeState not implemented")}
	}
	beginUpdateAdministrativeState := n.beginUpdateAdministrativeState.get(req)
	if beginUpdateAdministrativeState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateAdministrativeState`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.UpdateDeviceAdministrativeState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdateAdministrativeState(req.Context(), resourceGroupNameParam, networkDeviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateAdministrativeState = &respr
		n.beginUpdateAdministrativeState.add(req, beginUpdateAdministrativeState)
	}

	resp, err := server.PollerResponderNext(beginUpdateAdministrativeState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdateAdministrativeState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateAdministrativeState) {
		n.beginUpdateAdministrativeState.remove(req)
	}

	return resp, nil
}

func (n *NetworkDevicesServerTransport) dispatchBeginUpgrade(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpgrade == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpgrade not implemented")}
	}
	beginUpgrade := n.beginUpgrade.get(req)
	if beginUpgrade == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkDevices/(?P<networkDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upgrade`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.UpdateVersion](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpgrade(req.Context(), resourceGroupNameParam, networkDeviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpgrade = &respr
		n.beginUpgrade.add(req, beginUpgrade)
	}

	resp, err := server.PollerResponderNext(beginUpgrade, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpgrade.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpgrade) {
		n.beginUpgrade.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to NetworkDevicesServerTransport
var networkDevicesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
