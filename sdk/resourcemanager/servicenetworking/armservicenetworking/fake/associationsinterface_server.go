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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
	"net/http"
	"net/url"
	"regexp"
)

// AssociationsInterfaceServer is a fake server for instances of the armservicenetworking.AssociationsInterfaceClient type.
type AssociationsInterfaceServer struct {
	// BeginCreateOrUpdate is the fake for method AssociationsInterfaceClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, resource armservicenetworking.Association, options *armservicenetworking.AssociationsInterfaceClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armservicenetworking.AssociationsInterfaceClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AssociationsInterfaceClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *armservicenetworking.AssociationsInterfaceClientBeginDeleteOptions) (resp azfake.PollerResponder[armservicenetworking.AssociationsInterfaceClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AssociationsInterfaceClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *armservicenetworking.AssociationsInterfaceClientGetOptions) (resp azfake.Responder[armservicenetworking.AssociationsInterfaceClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByTrafficControllerPager is the fake for method AssociationsInterfaceClient.NewListByTrafficControllerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByTrafficControllerPager func(resourceGroupName string, trafficControllerName string, options *armservicenetworking.AssociationsInterfaceClientListByTrafficControllerOptions) (resp azfake.PagerResponder[armservicenetworking.AssociationsInterfaceClientListByTrafficControllerResponse])

	// Update is the fake for method AssociationsInterfaceClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, properties armservicenetworking.AssociationUpdate, options *armservicenetworking.AssociationsInterfaceClientUpdateOptions) (resp azfake.Responder[armservicenetworking.AssociationsInterfaceClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAssociationsInterfaceServerTransport creates a new instance of AssociationsInterfaceServerTransport with the provided implementation.
// The returned AssociationsInterfaceServerTransport instance is connected to an instance of armservicenetworking.AssociationsInterfaceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssociationsInterfaceServerTransport(srv *AssociationsInterfaceServer) *AssociationsInterfaceServerTransport {
	return &AssociationsInterfaceServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armservicenetworking.AssociationsInterfaceClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armservicenetworking.AssociationsInterfaceClientDeleteResponse]](),
		newListByTrafficControllerPager: newTracker[azfake.PagerResponder[armservicenetworking.AssociationsInterfaceClientListByTrafficControllerResponse]](),
	}
}

// AssociationsInterfaceServerTransport connects instances of armservicenetworking.AssociationsInterfaceClient to instances of AssociationsInterfaceServer.
// Don't use this type directly, use NewAssociationsInterfaceServerTransport instead.
type AssociationsInterfaceServerTransport struct {
	srv                             *AssociationsInterfaceServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armservicenetworking.AssociationsInterfaceClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armservicenetworking.AssociationsInterfaceClientDeleteResponse]]
	newListByTrafficControllerPager *tracker[azfake.PagerResponder[armservicenetworking.AssociationsInterfaceClientListByTrafficControllerResponse]]
}

// Do implements the policy.Transporter interface for AssociationsInterfaceServerTransport.
func (a *AssociationsInterfaceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AssociationsInterfaceClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AssociationsInterfaceClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AssociationsInterfaceClient.Get":
		resp, err = a.dispatchGet(req)
	case "AssociationsInterfaceClient.NewListByTrafficControllerPager":
		resp, err = a.dispatchNewListByTrafficControllerPager(req)
	case "AssociationsInterfaceClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AssociationsInterfaceServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/associations/(?P<associationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armservicenetworking.Association](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
		if err != nil {
			return nil, err
		}
		associationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("associationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, trafficControllerNameParam, associationNameParam, body, nil)
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

func (a *AssociationsInterfaceServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/associations/(?P<associationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
		if err != nil {
			return nil, err
		}
		associationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("associationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, trafficControllerNameParam, associationNameParam, nil)
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

func (a *AssociationsInterfaceServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/associations/(?P<associationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
	if err != nil {
		return nil, err
	}
	associationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("associationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, trafficControllerNameParam, associationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Association, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssociationsInterfaceServerTransport) dispatchNewListByTrafficControllerPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByTrafficControllerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByTrafficControllerPager not implemented")}
	}
	newListByTrafficControllerPager := a.newListByTrafficControllerPager.get(req)
	if newListByTrafficControllerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/associations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByTrafficControllerPager(resourceGroupNameParam, trafficControllerNameParam, nil)
		newListByTrafficControllerPager = &resp
		a.newListByTrafficControllerPager.add(req, newListByTrafficControllerPager)
		server.PagerResponderInjectNextLinks(newListByTrafficControllerPager, req, func(page *armservicenetworking.AssociationsInterfaceClientListByTrafficControllerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByTrafficControllerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByTrafficControllerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByTrafficControllerPager) {
		a.newListByTrafficControllerPager.remove(req)
	}
	return resp, nil
}

func (a *AssociationsInterfaceServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceNetworking/trafficControllers/(?P<trafficControllerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/associations/(?P<associationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicenetworking.AssociationUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	trafficControllerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("trafficControllerName")])
	if err != nil {
		return nil, err
	}
	associationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("associationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameParam, trafficControllerNameParam, associationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Association, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
