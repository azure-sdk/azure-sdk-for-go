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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
	"net/http"
	"net/url"
	"regexp"
)

// AutoUpgradeProfilesServer is a fake server for instances of the armcontainerservicefleet.AutoUpgradeProfilesClient type.
type AutoUpgradeProfilesServer struct {
	// BeginCreateOrUpdate is the fake for method AutoUpgradeProfilesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, resource armcontainerservicefleet.AutoUpgradeProfile, options *armcontainerservicefleet.AutoUpgradeProfilesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AutoUpgradeProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *armcontainerservicefleet.AutoUpgradeProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AutoUpgradeProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *armcontainerservicefleet.AutoUpgradeProfilesClientGetOptions) (resp azfake.Responder[armcontainerservicefleet.AutoUpgradeProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFleetPager is the fake for method AutoUpgradeProfilesClient.NewListByFleetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFleetPager func(resourceGroupName string, fleetName string, options *armcontainerservicefleet.AutoUpgradeProfilesClientListByFleetOptions) (resp azfake.PagerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientListByFleetResponse])
}

// NewAutoUpgradeProfilesServerTransport creates a new instance of AutoUpgradeProfilesServerTransport with the provided implementation.
// The returned AutoUpgradeProfilesServerTransport instance is connected to an instance of armcontainerservicefleet.AutoUpgradeProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAutoUpgradeProfilesServerTransport(srv *AutoUpgradeProfilesServer) *AutoUpgradeProfilesServerTransport {
	return &AutoUpgradeProfilesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientDeleteResponse]](),
		newListByFleetPager: newTracker[azfake.PagerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientListByFleetResponse]](),
	}
}

// AutoUpgradeProfilesServerTransport connects instances of armcontainerservicefleet.AutoUpgradeProfilesClient to instances of AutoUpgradeProfilesServer.
// Don't use this type directly, use NewAutoUpgradeProfilesServerTransport instead.
type AutoUpgradeProfilesServerTransport struct {
	srv                 *AutoUpgradeProfilesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientDeleteResponse]]
	newListByFleetPager *tracker[azfake.PagerResponder[armcontainerservicefleet.AutoUpgradeProfilesClientListByFleetResponse]]
}

// Do implements the policy.Transporter interface for AutoUpgradeProfilesServerTransport.
func (a *AutoUpgradeProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AutoUpgradeProfilesClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AutoUpgradeProfilesClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AutoUpgradeProfilesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AutoUpgradeProfilesClient.NewListByFleetPager":
		resp, err = a.dispatchNewListByFleetPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AutoUpgradeProfilesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/autoUpgradeProfiles/(?P<autoUpgradeProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerservicefleet.AutoUpgradeProfile](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		autoUpgradeProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autoUpgradeProfileName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.AutoUpgradeProfilesClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armcontainerservicefleet.AutoUpgradeProfilesClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, fleetNameParam, autoUpgradeProfileNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AutoUpgradeProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/autoUpgradeProfiles/(?P<autoUpgradeProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		autoUpgradeProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autoUpgradeProfileName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.AutoUpgradeProfilesClientBeginDeleteOptions
		if ifMatchParam != nil {
			options = &armcontainerservicefleet.AutoUpgradeProfilesClientBeginDeleteOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, fleetNameParam, autoUpgradeProfileNameParam, options)
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

func (a *AutoUpgradeProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/autoUpgradeProfiles/(?P<autoUpgradeProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
	if err != nil {
		return nil, err
	}
	autoUpgradeProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autoUpgradeProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, fleetNameParam, autoUpgradeProfileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AutoUpgradeProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutoUpgradeProfilesServerTransport) dispatchNewListByFleetPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByFleetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFleetPager not implemented")}
	}
	newListByFleetPager := a.newListByFleetPager.get(req)
	if newListByFleetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/autoUpgradeProfiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByFleetPager(resourceGroupNameParam, fleetNameParam, nil)
		newListByFleetPager = &resp
		a.newListByFleetPager.add(req, newListByFleetPager)
		server.PagerResponderInjectNextLinks(newListByFleetPager, req, func(page *armcontainerservicefleet.AutoUpgradeProfilesClientListByFleetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFleetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByFleetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFleetPager) {
		a.newListByFleetPager.remove(req)
	}
	return resp, nil
}