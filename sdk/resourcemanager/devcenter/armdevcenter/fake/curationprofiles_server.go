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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CurationProfilesServer is a fake server for instances of the armdevcenter.CurationProfilesClient type.
type CurationProfilesServer struct {
	// BeginCreateOrUpdate is the fake for method CurationProfilesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, devCenterName string, curationProfileName string, body armdevcenter.CurationProfile, options *armdevcenter.CurationProfilesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdevcenter.CurationProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CurationProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, devCenterName string, curationProfileName string, options *armdevcenter.CurationProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdevcenter.CurationProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CurationProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, devCenterName string, curationProfileName string, options *armdevcenter.CurationProfilesClientGetOptions) (resp azfake.Responder[armdevcenter.CurationProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDevCenterPager is the fake for method CurationProfilesClient.NewListByDevCenterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDevCenterPager func(resourceGroupName string, devCenterName string, options *armdevcenter.CurationProfilesClientListByDevCenterOptions) (resp azfake.PagerResponder[armdevcenter.CurationProfilesClientListByDevCenterResponse])

	// BeginUpdate is the fake for method CurationProfilesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, devCenterName string, curationProfileName string, body armdevcenter.CurationProfileUpdate, options *armdevcenter.CurationProfilesClientBeginUpdateOptions) (resp azfake.PollerResponder[armdevcenter.CurationProfilesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCurationProfilesServerTransport creates a new instance of CurationProfilesServerTransport with the provided implementation.
// The returned CurationProfilesServerTransport instance is connected to an instance of armdevcenter.CurationProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCurationProfilesServerTransport(srv *CurationProfilesServer) *CurationProfilesServerTransport {
	return &CurationProfilesServerTransport{
		srv:                     srv,
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armdevcenter.CurationProfilesClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armdevcenter.CurationProfilesClientDeleteResponse]](),
		newListByDevCenterPager: newTracker[azfake.PagerResponder[armdevcenter.CurationProfilesClientListByDevCenterResponse]](),
		beginUpdate:             newTracker[azfake.PollerResponder[armdevcenter.CurationProfilesClientUpdateResponse]](),
	}
}

// CurationProfilesServerTransport connects instances of armdevcenter.CurationProfilesClient to instances of CurationProfilesServer.
// Don't use this type directly, use NewCurationProfilesServerTransport instead.
type CurationProfilesServerTransport struct {
	srv                     *CurationProfilesServer
	beginCreateOrUpdate     *tracker[azfake.PollerResponder[armdevcenter.CurationProfilesClientCreateOrUpdateResponse]]
	beginDelete             *tracker[azfake.PollerResponder[armdevcenter.CurationProfilesClientDeleteResponse]]
	newListByDevCenterPager *tracker[azfake.PagerResponder[armdevcenter.CurationProfilesClientListByDevCenterResponse]]
	beginUpdate             *tracker[azfake.PollerResponder[armdevcenter.CurationProfilesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for CurationProfilesServerTransport.
func (c *CurationProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CurationProfilesClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CurationProfilesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CurationProfilesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CurationProfilesClient.NewListByDevCenterPager":
		resp, err = c.dispatchNewListByDevCenterPager(req)
	case "CurationProfilesClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CurationProfilesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/curationProfiles/(?P<curationProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevcenter.CurationProfile](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		curationProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("curationProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, devCenterNameParam, curationProfileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *CurationProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/curationProfiles/(?P<curationProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		curationProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("curationProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, devCenterNameParam, curationProfileNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *CurationProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/curationProfiles/(?P<curationProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	curationProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("curationProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, devCenterNameParam, curationProfileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CurationProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CurationProfilesServerTransport) dispatchNewListByDevCenterPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByDevCenterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDevCenterPager not implemented")}
	}
	newListByDevCenterPager := c.newListByDevCenterPager.get(req)
	if newListByDevCenterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/curationProfiles`
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
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
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
		var options *armdevcenter.CurationProfilesClientListByDevCenterOptions
		if topParam != nil {
			options = &armdevcenter.CurationProfilesClientListByDevCenterOptions{
				Top: topParam,
			}
		}
		resp := c.srv.NewListByDevCenterPager(resourceGroupNameParam, devCenterNameParam, options)
		newListByDevCenterPager = &resp
		c.newListByDevCenterPager.add(req, newListByDevCenterPager)
		server.PagerResponderInjectNextLinks(newListByDevCenterPager, req, func(page *armdevcenter.CurationProfilesClientListByDevCenterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDevCenterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByDevCenterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDevCenterPager) {
		c.newListByDevCenterPager.remove(req)
	}
	return resp, nil
}

func (c *CurationProfilesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/curationProfiles/(?P<curationProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevcenter.CurationProfileUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		curationProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("curationProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, devCenterNameParam, curationProfileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}
