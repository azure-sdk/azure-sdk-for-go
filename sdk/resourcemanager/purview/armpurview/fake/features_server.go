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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
	"net/http"
	"net/url"
	"regexp"
)

// FeaturesServer is a fake server for instances of the armpurview.FeaturesClient type.
type FeaturesServer struct {
	// AccountGet is the fake for method FeaturesClient.AccountGet
	// HTTP status codes to indicate success: http.StatusOK
	AccountGet func(ctx context.Context, resourceGroupName string, accountName string, featureRequest armpurview.BatchFeatureRequest, options *armpurview.FeaturesClientAccountGetOptions) (resp azfake.Responder[armpurview.FeaturesClientAccountGetResponse], errResp azfake.ErrorResponder)

	// SubscriptionGet is the fake for method FeaturesClient.SubscriptionGet
	// HTTP status codes to indicate success: http.StatusOK
	SubscriptionGet func(ctx context.Context, locations string, featureRequest armpurview.BatchFeatureRequest, options *armpurview.FeaturesClientSubscriptionGetOptions) (resp azfake.Responder[armpurview.FeaturesClientSubscriptionGetResponse], errResp azfake.ErrorResponder)
}

// NewFeaturesServerTransport creates a new instance of FeaturesServerTransport with the provided implementation.
// The returned FeaturesServerTransport instance is connected to an instance of armpurview.FeaturesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFeaturesServerTransport(srv *FeaturesServer) *FeaturesServerTransport {
	return &FeaturesServerTransport{srv: srv}
}

// FeaturesServerTransport connects instances of armpurview.FeaturesClient to instances of FeaturesServer.
// Don't use this type directly, use NewFeaturesServerTransport instead.
type FeaturesServerTransport struct {
	srv *FeaturesServer
}

// Do implements the policy.Transporter interface for FeaturesServerTransport.
func (f *FeaturesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FeaturesClient.AccountGet":
		resp, err = f.dispatchAccountGet(req)
	case "FeaturesClient.SubscriptionGet":
		resp, err = f.dispatchSubscriptionGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FeaturesServerTransport) dispatchAccountGet(req *http.Request) (*http.Response, error) {
	if f.srv.AccountGet == nil {
		return nil, &nonRetriableError{errors.New("fake for method AccountGet not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Purview/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listFeatures`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpurview.BatchFeatureRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.AccountGet(req.Context(), resourceGroupNameParam, accountNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BatchFeatureStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FeaturesServerTransport) dispatchSubscriptionGet(req *http.Request) (*http.Response, error) {
	if f.srv.SubscriptionGet == nil {
		return nil, &nonRetriableError{errors.New("fake for method SubscriptionGet not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Purview/locations/(?P<locations>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listFeatures`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpurview.BatchFeatureRequest](req)
	if err != nil {
		return nil, err
	}
	locationsParam, err := url.PathUnescape(matches[regex.SubexpIndex("locations")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.SubscriptionGet(req.Context(), locationsParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BatchFeatureStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
