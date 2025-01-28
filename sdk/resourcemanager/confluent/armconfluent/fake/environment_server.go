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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// EnvironmentServer is a fake server for instances of the armconfluent.EnvironmentClient type.
type EnvironmentServer struct {
	// CreateOrUpdate is the fake for method EnvironmentClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, organizationName string, environmentID string, options *armconfluent.EnvironmentClientCreateOrUpdateOptions) (resp azfake.Responder[armconfluent.EnvironmentClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method EnvironmentClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, organizationName string, environmentID string, options *armconfluent.EnvironmentClientBeginDeleteOptions) (resp azfake.PollerResponder[armconfluent.EnvironmentClientDeleteResponse], errResp azfake.ErrorResponder)
}

// NewEnvironmentServerTransport creates a new instance of EnvironmentServerTransport with the provided implementation.
// The returned EnvironmentServerTransport instance is connected to an instance of armconfluent.EnvironmentClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEnvironmentServerTransport(srv *EnvironmentServer) *EnvironmentServerTransport {
	return &EnvironmentServerTransport{
		srv:         srv,
		beginDelete: newTracker[azfake.PollerResponder[armconfluent.EnvironmentClientDeleteResponse]](),
	}
}

// EnvironmentServerTransport connects instances of armconfluent.EnvironmentClient to instances of EnvironmentServer.
// Don't use this type directly, use NewEnvironmentServerTransport instead.
type EnvironmentServerTransport struct {
	srv         *EnvironmentServer
	beginDelete *tracker[azfake.PollerResponder[armconfluent.EnvironmentClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for EnvironmentServerTransport.
func (e *EnvironmentServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EnvironmentClient.CreateOrUpdate":
		resp, err = e.dispatchCreateOrUpdate(req)
	case "EnvironmentClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EnvironmentServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Confluent/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armconfluent.SCEnvironmentRecord](req)
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
	environmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentId")])
	if err != nil {
		return nil, err
	}
	var options *armconfluent.EnvironmentClientCreateOrUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armconfluent.EnvironmentClientCreateOrUpdateOptions{
			Body: &body,
		}
	}
	respr, errRespr := e.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, organizationNameParam, environmentIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SCEnvironmentRecord, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnvironmentServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Confluent/organizations/(?P<organizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		environmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, organizationNameParam, environmentIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}
