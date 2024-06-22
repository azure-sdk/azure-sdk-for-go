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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics/v2"
	"net/http"
	"net/url"
	"regexp"
)

// DataProductsServer is a fake server for instances of the armnetworkanalytics.DataProductsClient type.
type DataProductsServer struct {
	// AddUserRole is the fake for method DataProductsClient.AddUserRole
	// HTTP status codes to indicate success: http.StatusOK
	AddUserRole func(ctx context.Context, resourceGroupName string, dataProductName string, body armnetworkanalytics.RoleAssignmentCommonProperties, options *armnetworkanalytics.DataProductsClientAddUserRoleOptions) (resp azfake.Responder[armnetworkanalytics.DataProductsClientAddUserRoleResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method DataProductsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, dataProductName string, resource armnetworkanalytics.DataProduct, options *armnetworkanalytics.DataProductsClientBeginCreateOptions) (resp azfake.PollerResponder[armnetworkanalytics.DataProductsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DataProductsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, dataProductName string, options *armnetworkanalytics.DataProductsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkanalytics.DataProductsClientDeleteResponse], errResp azfake.ErrorResponder)

	// GenerateStorageAccountSasToken is the fake for method DataProductsClient.GenerateStorageAccountSasToken
	// HTTP status codes to indicate success: http.StatusOK
	GenerateStorageAccountSasToken func(ctx context.Context, resourceGroupName string, dataProductName string, body armnetworkanalytics.AccountSas, options *armnetworkanalytics.DataProductsClientGenerateStorageAccountSasTokenOptions) (resp azfake.Responder[armnetworkanalytics.DataProductsClientGenerateStorageAccountSasTokenResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DataProductsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, dataProductName string, options *armnetworkanalytics.DataProductsClientGetOptions) (resp azfake.Responder[armnetworkanalytics.DataProductsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method DataProductsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetworkanalytics.DataProductsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetworkanalytics.DataProductsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method DataProductsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetworkanalytics.DataProductsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetworkanalytics.DataProductsClientListBySubscriptionResponse])

	// ListRolesAssignments is the fake for method DataProductsClient.ListRolesAssignments
	// HTTP status codes to indicate success: http.StatusOK
	ListRolesAssignments func(ctx context.Context, resourceGroupName string, dataProductName string, body any, options *armnetworkanalytics.DataProductsClientListRolesAssignmentsOptions) (resp azfake.Responder[armnetworkanalytics.DataProductsClientListRolesAssignmentsResponse], errResp azfake.ErrorResponder)

	// RemoveUserRole is the fake for method DataProductsClient.RemoveUserRole
	// HTTP status codes to indicate success: http.StatusNoContent
	RemoveUserRole func(ctx context.Context, resourceGroupName string, dataProductName string, body armnetworkanalytics.RoleAssignmentDetail, options *armnetworkanalytics.DataProductsClientRemoveUserRoleOptions) (resp azfake.Responder[armnetworkanalytics.DataProductsClientRemoveUserRoleResponse], errResp azfake.ErrorResponder)

	// RotateKey is the fake for method DataProductsClient.RotateKey
	// HTTP status codes to indicate success: http.StatusNoContent
	RotateKey func(ctx context.Context, resourceGroupName string, dataProductName string, body armnetworkanalytics.KeyVaultInfo, options *armnetworkanalytics.DataProductsClientRotateKeyOptions) (resp azfake.Responder[armnetworkanalytics.DataProductsClientRotateKeyResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method DataProductsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, dataProductName string, properties armnetworkanalytics.DataProductUpdate, options *armnetworkanalytics.DataProductsClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetworkanalytics.DataProductsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDataProductsServerTransport creates a new instance of DataProductsServerTransport with the provided implementation.
// The returned DataProductsServerTransport instance is connected to an instance of armnetworkanalytics.DataProductsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDataProductsServerTransport(srv *DataProductsServer) *DataProductsServerTransport {
	return &DataProductsServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armnetworkanalytics.DataProductsClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetworkanalytics.DataProductsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetworkanalytics.DataProductsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armnetworkanalytics.DataProductsClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armnetworkanalytics.DataProductsClientUpdateResponse]](),
	}
}

// DataProductsServerTransport connects instances of armnetworkanalytics.DataProductsClient to instances of DataProductsServer.
// Don't use this type directly, use NewDataProductsServerTransport instead.
type DataProductsServerTransport struct {
	srv                         *DataProductsServer
	beginCreate                 *tracker[azfake.PollerResponder[armnetworkanalytics.DataProductsClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetworkanalytics.DataProductsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetworkanalytics.DataProductsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armnetworkanalytics.DataProductsClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armnetworkanalytics.DataProductsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for DataProductsServerTransport.
func (d *DataProductsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DataProductsClient.AddUserRole":
		resp, err = d.dispatchAddUserRole(req)
	case "DataProductsClient.BeginCreate":
		resp, err = d.dispatchBeginCreate(req)
	case "DataProductsClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DataProductsClient.GenerateStorageAccountSasToken":
		resp, err = d.dispatchGenerateStorageAccountSasToken(req)
	case "DataProductsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DataProductsClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DataProductsClient.NewListBySubscriptionPager":
		resp, err = d.dispatchNewListBySubscriptionPager(req)
	case "DataProductsClient.ListRolesAssignments":
		resp, err = d.dispatchListRolesAssignments(req)
	case "DataProductsClient.RemoveUserRole":
		resp, err = d.dispatchRemoveUserRole(req)
	case "DataProductsClient.RotateKey":
		resp, err = d.dispatchRotateKey(req)
	case "DataProductsClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DataProductsServerTransport) dispatchAddUserRole(req *http.Request) (*http.Response, error) {
	if d.srv.AddUserRole == nil {
		return nil, &nonRetriableError{errors.New("fake for method AddUserRole not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/addUserRole`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetworkanalytics.RoleAssignmentCommonProperties](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.AddUserRole(req.Context(), resourceGroupNameParam, dataProductNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignmentDetail, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := d.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkanalytics.DataProduct](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreate(req.Context(), resourceGroupNameParam, dataProductNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		d.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		d.beginCreate.remove(req)
	}

	return resp, nil
}

func (d *DataProductsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, dataProductNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DataProductsServerTransport) dispatchGenerateStorageAccountSasToken(req *http.Request) (*http.Response, error) {
	if d.srv.GenerateStorageAccountSasToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method GenerateStorageAccountSasToken not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateStorageAccountSasToken`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetworkanalytics.AccountSas](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GenerateStorageAccountSasToken(req.Context(), resourceGroupNameParam, dataProductNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccountSasToken, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, dataProductNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataProduct, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := d.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		d.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetworkanalytics.DataProductsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		d.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := d.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		d.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armnetworkanalytics.DataProductsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		d.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchListRolesAssignments(req *http.Request) (*http.Response, error) {
	if d.srv.ListRolesAssignments == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListRolesAssignments not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listRolesAssignments`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[any](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.ListRolesAssignments(req.Context(), resourceGroupNameParam, dataProductNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListRoleAssignments, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchRemoveUserRole(req *http.Request) (*http.Response, error) {
	if d.srv.RemoveUserRole == nil {
		return nil, &nonRetriableError{errors.New("fake for method RemoveUserRole not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/removeUserRole`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetworkanalytics.RoleAssignmentDetail](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RemoveUserRole(req.Context(), resourceGroupNameParam, dataProductNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchRotateKey(req *http.Request) (*http.Response, error) {
	if d.srv.RotateKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method RotateKey not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rotateKey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetworkanalytics.KeyVaultInfo](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RotateKey(req.Context(), resourceGroupNameParam, dataProductNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataProductsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := d.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkAnalytics/dataProducts/(?P<dataProductName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkanalytics.DataProductUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		dataProductNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataProductName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameParam, dataProductNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		d.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		d.beginUpdate.remove(req)
	}

	return resp, nil
}
