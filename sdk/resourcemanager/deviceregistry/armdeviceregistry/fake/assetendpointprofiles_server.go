// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
	"net/http"
	"net/url"
	"regexp"
)

// AssetEndpointProfilesServer is a fake server for instances of the armdeviceregistry.AssetEndpointProfilesClient type.
type AssetEndpointProfilesServer struct {
	// BeginCreateOrReplace is the fake for method AssetEndpointProfilesClient.BeginCreateOrReplace
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrReplace func(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, resource armdeviceregistry.AssetEndpointProfile, options *armdeviceregistry.AssetEndpointProfilesClientBeginCreateOrReplaceOptions) (resp azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientCreateOrReplaceResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AssetEndpointProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, options *armdeviceregistry.AssetEndpointProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AssetEndpointProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, options *armdeviceregistry.AssetEndpointProfilesClientGetOptions) (resp azfake.Responder[armdeviceregistry.AssetEndpointProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method AssetEndpointProfilesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armdeviceregistry.AssetEndpointProfilesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armdeviceregistry.AssetEndpointProfilesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method AssetEndpointProfilesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armdeviceregistry.AssetEndpointProfilesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armdeviceregistry.AssetEndpointProfilesClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method AssetEndpointProfilesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, properties armdeviceregistry.AssetEndpointProfile, options *armdeviceregistry.AssetEndpointProfilesClientBeginUpdateOptions) (resp azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAssetEndpointProfilesServerTransport creates a new instance of AssetEndpointProfilesServerTransport with the provided implementation.
// The returned AssetEndpointProfilesServerTransport instance is connected to an instance of armdeviceregistry.AssetEndpointProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssetEndpointProfilesServerTransport(srv *AssetEndpointProfilesServer) *AssetEndpointProfilesServerTransport {
	return &AssetEndpointProfilesServerTransport{
		srv:                         srv,
		beginCreateOrReplace:        newTracker[azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientCreateOrReplaceResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armdeviceregistry.AssetEndpointProfilesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armdeviceregistry.AssetEndpointProfilesClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientUpdateResponse]](),
	}
}

// AssetEndpointProfilesServerTransport connects instances of armdeviceregistry.AssetEndpointProfilesClient to instances of AssetEndpointProfilesServer.
// Don't use this type directly, use NewAssetEndpointProfilesServerTransport instead.
type AssetEndpointProfilesServerTransport struct {
	srv                         *AssetEndpointProfilesServer
	beginCreateOrReplace        *tracker[azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientCreateOrReplaceResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armdeviceregistry.AssetEndpointProfilesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armdeviceregistry.AssetEndpointProfilesClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armdeviceregistry.AssetEndpointProfilesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for AssetEndpointProfilesServerTransport.
func (a *AssetEndpointProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AssetEndpointProfilesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "AssetEndpointProfilesClient.BeginCreateOrReplace":
		resp, err = a.dispatchBeginCreateOrReplace(req)
	case "AssetEndpointProfilesClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AssetEndpointProfilesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AssetEndpointProfilesClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "AssetEndpointProfilesClient.NewListBySubscriptionPager":
		resp, err = a.dispatchNewListBySubscriptionPager(req)
	case "AssetEndpointProfilesClient.BeginUpdate":
		resp, err = a.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (a *AssetEndpointProfilesServerTransport) dispatchBeginCreateOrReplace(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrReplace == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrReplace not implemented")}
	}
	beginCreateOrReplace := a.beginCreateOrReplace.get(req)
	if beginCreateOrReplace == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceRegistry/assetEndpointProfiles/(?P<assetEndpointProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdeviceregistry.AssetEndpointProfile](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		assetEndpointProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetEndpointProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrReplace(req.Context(), resourceGroupNameParam, assetEndpointProfileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrReplace = &respr
		a.beginCreateOrReplace.add(req, beginCreateOrReplace)
	}

	resp, err := server.PollerResponderNext(beginCreateOrReplace, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrReplace.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrReplace) {
		a.beginCreateOrReplace.remove(req)
	}

	return resp, nil
}

func (a *AssetEndpointProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceRegistry/assetEndpointProfiles/(?P<assetEndpointProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		assetEndpointProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetEndpointProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, assetEndpointProfileNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AssetEndpointProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceRegistry/assetEndpointProfiles/(?P<assetEndpointProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	assetEndpointProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetEndpointProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, assetEndpointProfileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssetEndpointProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssetEndpointProfilesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceRegistry/assetEndpointProfiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armdeviceregistry.AssetEndpointProfilesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (a *AssetEndpointProfilesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := a.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceRegistry/assetEndpointProfiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		a.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armdeviceregistry.AssetEndpointProfilesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		a.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (a *AssetEndpointProfilesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceRegistry/assetEndpointProfiles/(?P<assetEndpointProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdeviceregistry.AssetEndpointProfile](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		assetEndpointProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assetEndpointProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), resourceGroupNameParam, assetEndpointProfileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}
