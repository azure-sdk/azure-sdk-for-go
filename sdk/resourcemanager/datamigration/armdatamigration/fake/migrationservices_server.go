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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
	"net/http"
	"net/url"
	"regexp"
)

// MigrationServicesServer is a fake server for instances of the armdatamigration.MigrationServicesClient type.
type MigrationServicesServer struct {
	// BeginCreateOrUpdate is the fake for method MigrationServicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters armdatamigration.MigrationService, options *armdatamigration.MigrationServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdatamigration.MigrationServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method MigrationServicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, migrationServiceName string, options *armdatamigration.MigrationServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdatamigration.MigrationServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method MigrationServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, migrationServiceName string, options *armdatamigration.MigrationServicesClientGetOptions) (resp azfake.Responder[armdatamigration.MigrationServicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method MigrationServicesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armdatamigration.MigrationServicesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armdatamigration.MigrationServicesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method MigrationServicesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armdatamigration.MigrationServicesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armdatamigration.MigrationServicesClientListBySubscriptionResponse])

	// NewListMigrationsPager is the fake for method MigrationServicesClient.NewListMigrationsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMigrationsPager func(resourceGroupName string, migrationServiceName string, options *armdatamigration.MigrationServicesClientListMigrationsOptions) (resp azfake.PagerResponder[armdatamigration.MigrationServicesClientListMigrationsResponse])

	// BeginUpdate is the fake for method MigrationServicesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, migrationServiceName string, parameters armdatamigration.MigrationServiceUpdate, options *armdatamigration.MigrationServicesClientBeginUpdateOptions) (resp azfake.PollerResponder[armdatamigration.MigrationServicesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewMigrationServicesServerTransport creates a new instance of MigrationServicesServerTransport with the provided implementation.
// The returned MigrationServicesServerTransport instance is connected to an instance of armdatamigration.MigrationServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMigrationServicesServerTransport(srv *MigrationServicesServer) *MigrationServicesServerTransport {
	return &MigrationServicesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armdatamigration.MigrationServicesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armdatamigration.MigrationServicesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armdatamigration.MigrationServicesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armdatamigration.MigrationServicesClientListBySubscriptionResponse]](),
		newListMigrationsPager:      newTracker[azfake.PagerResponder[armdatamigration.MigrationServicesClientListMigrationsResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armdatamigration.MigrationServicesClientUpdateResponse]](),
	}
}

// MigrationServicesServerTransport connects instances of armdatamigration.MigrationServicesClient to instances of MigrationServicesServer.
// Don't use this type directly, use NewMigrationServicesServerTransport instead.
type MigrationServicesServerTransport struct {
	srv                         *MigrationServicesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armdatamigration.MigrationServicesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armdatamigration.MigrationServicesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armdatamigration.MigrationServicesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armdatamigration.MigrationServicesClientListBySubscriptionResponse]]
	newListMigrationsPager      *tracker[azfake.PagerResponder[armdatamigration.MigrationServicesClientListMigrationsResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armdatamigration.MigrationServicesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for MigrationServicesServerTransport.
func (m *MigrationServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MigrationServicesClient.BeginCreateOrUpdate":
		resp, err = m.dispatchBeginCreateOrUpdate(req)
	case "MigrationServicesClient.BeginDelete":
		resp, err = m.dispatchBeginDelete(req)
	case "MigrationServicesClient.Get":
		resp, err = m.dispatchGet(req)
	case "MigrationServicesClient.NewListByResourceGroupPager":
		resp, err = m.dispatchNewListByResourceGroupPager(req)
	case "MigrationServicesClient.NewListBySubscriptionPager":
		resp, err = m.dispatchNewListBySubscriptionPager(req)
	case "MigrationServicesClient.NewListMigrationsPager":
		resp, err = m.dispatchNewListMigrationsPager(req)
	case "MigrationServicesClient.BeginUpdate":
		resp, err = m.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MigrationServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/migrationServices/(?P<migrationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatamigration.MigrationService](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		migrationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, migrationServiceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *MigrationServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/migrationServices/(?P<migrationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		migrationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, migrationServiceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *MigrationServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/migrationServices/(?P<migrationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	migrationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, migrationServiceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MigrationService, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MigrationServicesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := m.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/migrationServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		m.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armdatamigration.MigrationServicesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		m.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (m *MigrationServicesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := m.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/migrationServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		m.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armdatamigration.MigrationServicesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		m.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (m *MigrationServicesServerTransport) dispatchNewListMigrationsPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListMigrationsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMigrationsPager not implemented")}
	}
	newListMigrationsPager := m.newListMigrationsPager.get(req)
	if newListMigrationsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/migrationServices/(?P<migrationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listMigrations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		migrationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationServiceName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListMigrationsPager(resourceGroupNameParam, migrationServiceNameParam, nil)
		newListMigrationsPager = &resp
		m.newListMigrationsPager.add(req, newListMigrationsPager)
		server.PagerResponderInjectNextLinks(newListMigrationsPager, req, func(page *armdatamigration.MigrationServicesClientListMigrationsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListMigrationsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListMigrationsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMigrationsPager) {
		m.newListMigrationsPager.remove(req)
	}
	return resp, nil
}

func (m *MigrationServicesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := m.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/migrationServices/(?P<migrationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatamigration.MigrationServiceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		migrationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrationServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdate(req.Context(), resourceGroupNameParam, migrationServiceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		m.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		m.beginUpdate.remove(req)
	}

	return resp, nil
}
