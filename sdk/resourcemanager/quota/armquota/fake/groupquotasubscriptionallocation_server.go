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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota/v2"
	"net/http"
	"net/url"
	"regexp"
)

// GroupQuotaSubscriptionAllocationServer is a fake server for instances of the armquota.GroupQuotaSubscriptionAllocationClient type.
type GroupQuotaSubscriptionAllocationServer struct {
	// List is the fake for method GroupQuotaSubscriptionAllocationClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, managementGroupID string, groupQuotaName string, resourceProviderName string, location string, options *armquota.GroupQuotaSubscriptionAllocationClientListOptions) (resp azfake.Responder[armquota.GroupQuotaSubscriptionAllocationClientListResponse], errResp azfake.ErrorResponder)
}

// NewGroupQuotaSubscriptionAllocationServerTransport creates a new instance of GroupQuotaSubscriptionAllocationServerTransport with the provided implementation.
// The returned GroupQuotaSubscriptionAllocationServerTransport instance is connected to an instance of armquota.GroupQuotaSubscriptionAllocationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGroupQuotaSubscriptionAllocationServerTransport(srv *GroupQuotaSubscriptionAllocationServer) *GroupQuotaSubscriptionAllocationServerTransport {
	return &GroupQuotaSubscriptionAllocationServerTransport{srv: srv}
}

// GroupQuotaSubscriptionAllocationServerTransport connects instances of armquota.GroupQuotaSubscriptionAllocationClient to instances of GroupQuotaSubscriptionAllocationServer.
// Don't use this type directly, use NewGroupQuotaSubscriptionAllocationServerTransport instead.
type GroupQuotaSubscriptionAllocationServerTransport struct {
	srv *GroupQuotaSubscriptionAllocationServer
}

// Do implements the policy.Transporter interface for GroupQuotaSubscriptionAllocationServerTransport.
func (g *GroupQuotaSubscriptionAllocationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GroupQuotaSubscriptionAllocationClient.List":
		resp, err = g.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GroupQuotaSubscriptionAllocationServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if g.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Quota/groupQuotas/(?P<groupQuotaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceProviders/(?P<resourceProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotaAllocations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	groupQuotaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupQuotaName")])
	if err != nil {
		return nil, err
	}
	resourceProviderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderName")])
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.List(req.Context(), managementGroupIDParam, groupQuotaNameParam, resourceProviderNameParam, locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionQuotaAllocationsList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
