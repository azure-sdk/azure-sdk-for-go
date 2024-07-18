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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilepacketcore/armmobilepacketcore"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// UpfDeploymentsServer is a fake server for instances of the armmobilepacketcore.UpfDeploymentsClient type.
type UpfDeploymentsServer struct {
	// CreateOrUpdate is the fake for method UpfDeploymentsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, upfDeploymentName string, resource armmobilepacketcore.UpfDeploymentResource, options *armmobilepacketcore.UpfDeploymentsClientCreateOrUpdateOptions) (resp azfake.Responder[armmobilepacketcore.UpfDeploymentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method UpfDeploymentsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, upfDeploymentName string, options *armmobilepacketcore.UpfDeploymentsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmobilepacketcore.UpfDeploymentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method UpfDeploymentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, upfDeploymentName string, options *armmobilepacketcore.UpfDeploymentsClientGetOptions) (resp azfake.Responder[armmobilepacketcore.UpfDeploymentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method UpfDeploymentsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmobilepacketcore.UpfDeploymentsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmobilepacketcore.UpfDeploymentsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method UpfDeploymentsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmobilepacketcore.UpfDeploymentsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmobilepacketcore.UpfDeploymentsClientListBySubscriptionResponse])

	// UpdateTags is the fake for method UpfDeploymentsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, upfDeploymentName string, properties armmobilepacketcore.UpfDeploymentResourceTagsUpdate, options *armmobilepacketcore.UpfDeploymentsClientUpdateTagsOptions) (resp azfake.Responder[armmobilepacketcore.UpfDeploymentsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewUpfDeploymentsServerTransport creates a new instance of UpfDeploymentsServerTransport with the provided implementation.
// The returned UpfDeploymentsServerTransport instance is connected to an instance of armmobilepacketcore.UpfDeploymentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUpfDeploymentsServerTransport(srv *UpfDeploymentsServer) *UpfDeploymentsServerTransport {
	return &UpfDeploymentsServerTransport{
		srv:                         srv,
		beginDelete:                 newTracker[azfake.PollerResponder[armmobilepacketcore.UpfDeploymentsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmobilepacketcore.UpfDeploymentsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmobilepacketcore.UpfDeploymentsClientListBySubscriptionResponse]](),
	}
}

// UpfDeploymentsServerTransport connects instances of armmobilepacketcore.UpfDeploymentsClient to instances of UpfDeploymentsServer.
// Don't use this type directly, use NewUpfDeploymentsServerTransport instead.
type UpfDeploymentsServerTransport struct {
	srv                         *UpfDeploymentsServer
	beginDelete                 *tracker[azfake.PollerResponder[armmobilepacketcore.UpfDeploymentsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmobilepacketcore.UpfDeploymentsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmobilepacketcore.UpfDeploymentsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for UpfDeploymentsServerTransport.
func (u *UpfDeploymentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UpfDeploymentsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "UpfDeploymentsClient.CreateOrUpdate":
		resp, err = u.dispatchCreateOrUpdate(req)
	case "UpfDeploymentsClient.BeginDelete":
		resp, err = u.dispatchBeginDelete(req)
	case "UpfDeploymentsClient.Get":
		resp, err = u.dispatchGet(req)
	case "UpfDeploymentsClient.NewListByResourceGroupPager":
		resp, err = u.dispatchNewListByResourceGroupPager(req)
	case "UpfDeploymentsClient.NewListBySubscriptionPager":
		resp, err = u.dispatchNewListBySubscriptionPager(req)
	case "UpfDeploymentsClient.UpdateTags":
		resp, err = u.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (u *UpfDeploymentsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if u.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobilePacketCore/upfDeployments/(?P<upfDeploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmobilepacketcore.UpfDeploymentResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	upfDeploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("upfDeploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, upfDeploymentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UpfDeploymentResource, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

func (u *UpfDeploymentsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if u.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := u.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobilePacketCore/upfDeployments/(?P<upfDeploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		upfDeploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("upfDeploymentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := u.srv.BeginDelete(req.Context(), resourceGroupNameParam, upfDeploymentNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		u.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		u.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		u.beginDelete.remove(req)
	}

	return resp, nil
}

func (u *UpfDeploymentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobilePacketCore/upfDeployments/(?P<upfDeploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	upfDeploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("upfDeploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Get(req.Context(), resourceGroupNameParam, upfDeploymentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UpfDeploymentResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UpfDeploymentsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := u.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobilePacketCore/upfDeployments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := u.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		u.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmobilepacketcore.UpfDeploymentsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		u.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (u *UpfDeploymentsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := u.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobilePacketCore/upfDeployments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := u.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		u.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmobilepacketcore.UpfDeploymentsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		u.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (u *UpfDeploymentsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if u.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MobilePacketCore/upfDeployments/(?P<upfDeploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmobilepacketcore.UpfDeploymentResourceTagsUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	upfDeploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("upfDeploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.UpdateTags(req.Context(), resourceGroupNameParam, upfDeploymentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UpfDeploymentResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
