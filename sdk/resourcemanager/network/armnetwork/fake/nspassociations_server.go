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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// NspAssociationsServer is a fake server for instances of the armnetwork.NspAssociationsClient type.
type NspAssociationsServer struct {
	// BeginCreateOrUpdate is the fake for method NspAssociationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, parameters armnetwork.NspAssociation, options *armnetwork.NspAssociationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.NspAssociationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NspAssociationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, options *armnetwork.NspAssociationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.NspAssociationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NspAssociationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, associationName string, options *armnetwork.NspAssociationsClientGetOptions) (resp azfake.Responder[armnetwork.NspAssociationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method NspAssociationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkSecurityPerimeterName string, options *armnetwork.NspAssociationsClientListOptions) (resp azfake.PagerResponder[armnetwork.NspAssociationsClientListResponse])
}

// NewNspAssociationsServerTransport creates a new instance of NspAssociationsServerTransport with the provided implementation.
// The returned NspAssociationsServerTransport instance is connected to an instance of armnetwork.NspAssociationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNspAssociationsServerTransport(srv *NspAssociationsServer) *NspAssociationsServerTransport {
	return &NspAssociationsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.NspAssociationsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.NspAssociationsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.NspAssociationsClientListResponse]](),
	}
}

// NspAssociationsServerTransport connects instances of armnetwork.NspAssociationsClient to instances of NspAssociationsServer.
// Don't use this type directly, use NewNspAssociationsServerTransport instead.
type NspAssociationsServerTransport struct {
	srv                 *NspAssociationsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.NspAssociationsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.NspAssociationsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.NspAssociationsClientListResponse]]
}

// Do implements the policy.Transporter interface for NspAssociationsServerTransport.
func (n *NspAssociationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NspAssociationsClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NspAssociationsClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NspAssociationsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NspAssociationsClient.NewListPager":
		resp, err = n.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NspAssociationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceAssociations/(?P<associationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.NspAssociation](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
		if err != nil {
			return nil, err
		}
		associationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("associationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, associationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NspAssociationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceAssociations/(?P<associationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
		if err != nil {
			return nil, err
		}
		associationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("associationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, associationNameParam, nil)
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

func (n *NspAssociationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceAssociations/(?P<associationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
	if err != nil {
		return nil, err
	}
	associationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("associationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, associationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NspAssociation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NspAssociationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := n.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceAssociations`
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
		networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
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
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.NspAssociationsClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.NspAssociationsClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := n.srv.NewListPager(resourceGroupNameParam, networkSecurityPerimeterNameParam, options)
		newListPager = &resp
		n.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.NspAssociationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		n.newListPager.remove(req)
	}
	return resp, nil
}
