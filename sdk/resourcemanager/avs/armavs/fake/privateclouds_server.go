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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PrivateCloudsServer is a fake server for instances of the armavs.PrivateCloudsClient type.
type PrivateCloudsServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateCloudsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud armavs.PrivateCloud, options *armavs.PrivateCloudsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armavs.PrivateCloudsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateCloudsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateCloudName string, options *armavs.PrivateCloudsClientBeginDeleteOptions) (resp azfake.PollerResponder[armavs.PrivateCloudsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateCloudsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateCloudName string, options *armavs.PrivateCloudsClientGetOptions) (resp azfake.Responder[armavs.PrivateCloudsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PrivateCloudsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armavs.PrivateCloudsClientListOptions) (resp azfake.PagerResponder[armavs.PrivateCloudsClientListResponse])

	// ListAdminCredentials is the fake for method PrivateCloudsClient.ListAdminCredentials
	// HTTP status codes to indicate success: http.StatusOK
	ListAdminCredentials func(ctx context.Context, resourceGroupName string, privateCloudName string, options *armavs.PrivateCloudsClientListAdminCredentialsOptions) (resp azfake.Responder[armavs.PrivateCloudsClientListAdminCredentialsResponse], errResp azfake.ErrorResponder)

	// NewListInSubscriptionPager is the fake for method PrivateCloudsClient.NewListInSubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListInSubscriptionPager func(options *armavs.PrivateCloudsClientListInSubscriptionOptions) (resp azfake.PagerResponder[armavs.PrivateCloudsClientListInSubscriptionResponse])

	// BeginRotateNsxtPassword is the fake for method PrivateCloudsClient.BeginRotateNsxtPassword
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginRotateNsxtPassword func(ctx context.Context, resourceGroupName string, privateCloudName string, options *armavs.PrivateCloudsClientBeginRotateNsxtPasswordOptions) (resp azfake.PollerResponder[armavs.PrivateCloudsClientRotateNsxtPasswordResponse], errResp azfake.ErrorResponder)

	// BeginRotateVcenterPassword is the fake for method PrivateCloudsClient.BeginRotateVcenterPassword
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginRotateVcenterPassword func(ctx context.Context, resourceGroupName string, privateCloudName string, options *armavs.PrivateCloudsClientBeginRotateVcenterPasswordOptions) (resp azfake.PollerResponder[armavs.PrivateCloudsClientRotateVcenterPasswordResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method PrivateCloudsClient.Update
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Update func(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate armavs.PrivateCloudUpdate, options *armavs.PrivateCloudsClientUpdateOptions) (resp azfake.Responder[armavs.PrivateCloudsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPrivateCloudsServerTransport creates a new instance of PrivateCloudsServerTransport with the provided implementation.
// The returned PrivateCloudsServerTransport instance is connected to an instance of armavs.PrivateCloudsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateCloudsServerTransport(srv *PrivateCloudsServer) *PrivateCloudsServerTransport {
	return &PrivateCloudsServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armavs.PrivateCloudsClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armavs.PrivateCloudsClientDeleteResponse]](),
		newListPager:               newTracker[azfake.PagerResponder[armavs.PrivateCloudsClientListResponse]](),
		newListInSubscriptionPager: newTracker[azfake.PagerResponder[armavs.PrivateCloudsClientListInSubscriptionResponse]](),
		beginRotateNsxtPassword:    newTracker[azfake.PollerResponder[armavs.PrivateCloudsClientRotateNsxtPasswordResponse]](),
		beginRotateVcenterPassword: newTracker[azfake.PollerResponder[armavs.PrivateCloudsClientRotateVcenterPasswordResponse]](),
	}
}

// PrivateCloudsServerTransport connects instances of armavs.PrivateCloudsClient to instances of PrivateCloudsServer.
// Don't use this type directly, use NewPrivateCloudsServerTransport instead.
type PrivateCloudsServerTransport struct {
	srv                        *PrivateCloudsServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armavs.PrivateCloudsClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armavs.PrivateCloudsClientDeleteResponse]]
	newListPager               *tracker[azfake.PagerResponder[armavs.PrivateCloudsClientListResponse]]
	newListInSubscriptionPager *tracker[azfake.PagerResponder[armavs.PrivateCloudsClientListInSubscriptionResponse]]
	beginRotateNsxtPassword    *tracker[azfake.PollerResponder[armavs.PrivateCloudsClientRotateNsxtPasswordResponse]]
	beginRotateVcenterPassword *tracker[azfake.PollerResponder[armavs.PrivateCloudsClientRotateVcenterPasswordResponse]]
}

// Do implements the policy.Transporter interface for PrivateCloudsServerTransport.
func (p *PrivateCloudsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateCloudsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "PrivateCloudsClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PrivateCloudsClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateCloudsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateCloudsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PrivateCloudsClient.ListAdminCredentials":
		resp, err = p.dispatchListAdminCredentials(req)
	case "PrivateCloudsClient.NewListInSubscriptionPager":
		resp, err = p.dispatchNewListInSubscriptionPager(req)
	case "PrivateCloudsClient.BeginRotateNsxtPassword":
		resp, err = p.dispatchBeginRotateNsxtPassword(req)
	case "PrivateCloudsClient.BeginRotateVcenterPassword":
		resp, err = p.dispatchBeginRotateVcenterPassword(req)
	case "PrivateCloudsClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (p *PrivateCloudsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.PrivateCloud](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PrivateCloudsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateCloudNameParam, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateCloudsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, privateCloudNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateCloud, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateCloudsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armavs.PrivateCloudsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
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

func (p *PrivateCloudsServerTransport) dispatchListAdminCredentials(req *http.Request) (*http.Response, error) {
	if p.srv.ListAdminCredentials == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAdminCredentials not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listAdminCredentials`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ListAdminCredentials(req.Context(), resourceGroupNameParam, privateCloudNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AdminCredentials, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateCloudsServerTransport) dispatchNewListInSubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListInSubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListInSubscriptionPager not implemented")}
	}
	newListInSubscriptionPager := p.newListInSubscriptionPager.get(req)
	if newListInSubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListInSubscriptionPager(nil)
		newListInSubscriptionPager = &resp
		p.newListInSubscriptionPager.add(req, newListInSubscriptionPager)
		server.PagerResponderInjectNextLinks(newListInSubscriptionPager, req, func(page *armavs.PrivateCloudsClientListInSubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListInSubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListInSubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListInSubscriptionPager) {
		p.newListInSubscriptionPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateCloudsServerTransport) dispatchBeginRotateNsxtPassword(req *http.Request) (*http.Response, error) {
	if p.srv.BeginRotateNsxtPassword == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRotateNsxtPassword not implemented")}
	}
	beginRotateNsxtPassword := p.beginRotateNsxtPassword.get(req)
	if beginRotateNsxtPassword == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rotateNsxtPassword`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginRotateNsxtPassword(req.Context(), resourceGroupNameParam, privateCloudNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRotateNsxtPassword = &respr
		p.beginRotateNsxtPassword.add(req, beginRotateNsxtPassword)
	}

	resp, err := server.PollerResponderNext(beginRotateNsxtPassword, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginRotateNsxtPassword.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRotateNsxtPassword) {
		p.beginRotateNsxtPassword.remove(req)
	}

	return resp, nil
}

func (p *PrivateCloudsServerTransport) dispatchBeginRotateVcenterPassword(req *http.Request) (*http.Response, error) {
	if p.srv.BeginRotateVcenterPassword == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRotateVcenterPassword not implemented")}
	}
	beginRotateVcenterPassword := p.beginRotateVcenterPassword.get(req)
	if beginRotateVcenterPassword == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rotateVcenterPassword`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginRotateVcenterPassword(req.Context(), resourceGroupNameParam, privateCloudNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRotateVcenterPassword = &respr
		p.beginRotateVcenterPassword.add(req, beginRotateVcenterPassword)
	}

	resp, err := server.PollerResponderNext(beginRotateVcenterPassword, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginRotateVcenterPassword.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRotateVcenterPassword) {
		p.beginRotateVcenterPassword.remove(req)
	}

	return resp, nil
}

func (p *PrivateCloudsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armavs.PrivateCloudUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Update(req.Context(), resourceGroupNameParam, privateCloudNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateCloud, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}
