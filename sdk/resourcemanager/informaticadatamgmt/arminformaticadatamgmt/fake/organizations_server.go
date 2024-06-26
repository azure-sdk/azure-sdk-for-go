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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
	"net/http"
	"net/url"
	"regexp"
)

// OrganizationsServer is a fake server for instances of the arminformaticadatamgmt.OrganizationsClient type.
type OrganizationsServer struct {
	// BeginCreateOrUpdate is the fake for method OrganizationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, organizationName string, resource arminformaticadatamgmt.InformaticaOrganizationResource, options *arminformaticadatamgmt.OrganizationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[arminformaticadatamgmt.OrganizationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method OrganizationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, organizationName string, options *arminformaticadatamgmt.OrganizationsClientBeginDeleteOptions) (resp azfake.PollerResponder[arminformaticadatamgmt.OrganizationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OrganizationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, organizationName string, options *arminformaticadatamgmt.OrganizationsClientGetOptions) (resp azfake.Responder[arminformaticadatamgmt.OrganizationsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAllServerlessRuntimes is the fake for method OrganizationsClient.GetAllServerlessRuntimes
	// HTTP status codes to indicate success: http.StatusOK
	GetAllServerlessRuntimes func(ctx context.Context, resourceGroupName string, organizationName string, options *arminformaticadatamgmt.OrganizationsClientGetAllServerlessRuntimesOptions) (resp azfake.Responder[arminformaticadatamgmt.OrganizationsClientGetAllServerlessRuntimesResponse], errResp azfake.ErrorResponder)

	// GetServerlessMetadata is the fake for method OrganizationsClient.GetServerlessMetadata
	// HTTP status codes to indicate success: http.StatusOK
	GetServerlessMetadata func(ctx context.Context, resourceGroupName string, organizationName string, options *arminformaticadatamgmt.OrganizationsClientGetServerlessMetadataOptions) (resp azfake.Responder[arminformaticadatamgmt.OrganizationsClientGetServerlessMetadataResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method OrganizationsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *arminformaticadatamgmt.OrganizationsClientListByResourceGroupOptions) (resp azfake.PagerResponder[arminformaticadatamgmt.OrganizationsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method OrganizationsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *arminformaticadatamgmt.OrganizationsClientListBySubscriptionOptions) (resp azfake.PagerResponder[arminformaticadatamgmt.OrganizationsClientListBySubscriptionResponse])

	// Update is the fake for method OrganizationsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, organizationName string, properties arminformaticadatamgmt.InformaticaOrganizationResourceUpdate, options *arminformaticadatamgmt.OrganizationsClientUpdateOptions) (resp azfake.Responder[arminformaticadatamgmt.OrganizationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewOrganizationsServerTransport creates a new instance of OrganizationsServerTransport with the provided implementation.
// The returned OrganizationsServerTransport instance is connected to an instance of arminformaticadatamgmt.OrganizationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOrganizationsServerTransport(srv *OrganizationsServer) *OrganizationsServerTransport {
	return &OrganizationsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[arminformaticadatamgmt.OrganizationsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[arminformaticadatamgmt.OrganizationsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[arminformaticadatamgmt.OrganizationsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[arminformaticadatamgmt.OrganizationsClientListBySubscriptionResponse]](),
	}
}

// OrganizationsServerTransport connects instances of arminformaticadatamgmt.OrganizationsClient to instances of OrganizationsServer.
// Don't use this type directly, use NewOrganizationsServerTransport instead.
type OrganizationsServerTransport struct {
	srv                         *OrganizationsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[arminformaticadatamgmt.OrganizationsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[arminformaticadatamgmt.OrganizationsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[arminformaticadatamgmt.OrganizationsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[arminformaticadatamgmt.OrganizationsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for OrganizationsServerTransport.
func (o *OrganizationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OrganizationsClient.BeginCreateOrUpdate":
		resp, err = o.dispatchBeginCreateOrUpdate(req)
	case "OrganizationsClient.BeginDelete":
		resp, err = o.dispatchBeginDelete(req)
	case "OrganizationsClient.Get":
		resp, err = o.dispatchGet(req)
	case "OrganizationsClient.GetAllServerlessRuntimes":
		resp, err = o.dispatchGetAllServerlessRuntimes(req)
	case "OrganizationsClient.GetServerlessMetadata":
		resp, err = o.dispatchGetServerlessMetadata(req)
	case "OrganizationsClient.NewListByResourceGroupPager":
		resp, err = o.dispatchNewListByResourceGroupPager(req)
	case "OrganizationsClient.NewListBySubscriptionPager":
		resp, err = o.dispatchNewListBySubscriptionPager(req)
	case "OrganizationsClient.Update":
		resp, err = o.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := o.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[arminformaticadatamgmt.InformaticaOrganizationResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, organizationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		o.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		o.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		o.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if o.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := o.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginDelete(req.Context(), resourceGroupNameParam, organizationNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		o.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		o.beginDelete.remove(req)
	}

	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, organizationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InformaticaOrganizationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchGetAllServerlessRuntimes(req *http.Request) (*http.Response, error) {
	if o.srv.GetAllServerlessRuntimes == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAllServerlessRuntimes not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getAllServerlessRuntimes`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.GetAllServerlessRuntimes(req.Context(), resourceGroupNameParam, organizationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InformaticaServerlessRuntimeResourceList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchGetServerlessMetadata(req *http.Request) (*http.Response, error) {
	if o.srv.GetServerlessMetadata == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetServerlessMetadata not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getServerlessMetadata`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.GetServerlessMetadata(req.Context(), resourceGroupNameParam, organizationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerlessMetadataResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := o.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := o.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		o.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *arminformaticadatamgmt.OrganizationsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		o.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := o.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := o.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		o.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *arminformaticadatamgmt.OrganizationsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		o.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (o *OrganizationsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Informatica\.DataManagement/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[arminformaticadatamgmt.InformaticaOrganizationResourceUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	organizationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("organizationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Update(req.Context(), resourceGroupNameParam, organizationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InformaticaOrganizationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
