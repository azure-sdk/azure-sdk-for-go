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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// BlobContainersServer is a fake server for instances of the armstorage.BlobContainersClient type.
type BlobContainersServer struct {
	// ClearLegalHold is the fake for method BlobContainersClient.ClearLegalHold
	// HTTP status codes to indicate success: http.StatusOK
	ClearLegalHold func(ctx context.Context, resourceGroupName string, accountName string, containerName string, legalHold armstorage.LegalHold, options *armstorage.BlobContainersClientClearLegalHoldOptions) (resp azfake.Responder[armstorage.BlobContainersClientClearLegalHoldResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method BlobContainersClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, accountName string, containerName string, blobContainer armstorage.BlobContainer, options *armstorage.BlobContainersClientCreateOptions) (resp azfake.Responder[armstorage.BlobContainersClientCreateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateImmutabilityPolicy is the fake for method BlobContainersClient.CreateOrUpdateImmutabilityPolicy
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdateImmutabilityPolicy func(ctx context.Context, resourceGroupName string, accountName string, containerName string, options *armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyOptions) (resp azfake.Responder[armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method BlobContainersClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, containerName string, options *armstorage.BlobContainersClientDeleteOptions) (resp azfake.Responder[armstorage.BlobContainersClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteImmutabilityPolicy is the fake for method BlobContainersClient.DeleteImmutabilityPolicy
	// HTTP status codes to indicate success: http.StatusOK
	DeleteImmutabilityPolicy func(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string, options *armstorage.BlobContainersClientDeleteImmutabilityPolicyOptions) (resp azfake.Responder[armstorage.BlobContainersClientDeleteImmutabilityPolicyResponse], errResp azfake.ErrorResponder)

	// ExtendImmutabilityPolicy is the fake for method BlobContainersClient.ExtendImmutabilityPolicy
	// HTTP status codes to indicate success: http.StatusOK
	ExtendImmutabilityPolicy func(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string, options *armstorage.BlobContainersClientExtendImmutabilityPolicyOptions) (resp azfake.Responder[armstorage.BlobContainersClientExtendImmutabilityPolicyResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BlobContainersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, containerName string, options *armstorage.BlobContainersClientGetOptions) (resp azfake.Responder[armstorage.BlobContainersClientGetResponse], errResp azfake.ErrorResponder)

	// GetImmutabilityPolicy is the fake for method BlobContainersClient.GetImmutabilityPolicy
	// HTTP status codes to indicate success: http.StatusOK
	GetImmutabilityPolicy func(ctx context.Context, resourceGroupName string, accountName string, containerName string, options *armstorage.BlobContainersClientGetImmutabilityPolicyOptions) (resp azfake.Responder[armstorage.BlobContainersClientGetImmutabilityPolicyResponse], errResp azfake.ErrorResponder)

	// Lease is the fake for method BlobContainersClient.Lease
	// HTTP status codes to indicate success: http.StatusOK
	Lease func(ctx context.Context, resourceGroupName string, accountName string, containerName string, options *armstorage.BlobContainersClientLeaseOptions) (resp azfake.Responder[armstorage.BlobContainersClientLeaseResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BlobContainersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armstorage.BlobContainersClientListOptions) (resp azfake.PagerResponder[armstorage.BlobContainersClientListResponse])

	// LockImmutabilityPolicy is the fake for method BlobContainersClient.LockImmutabilityPolicy
	// HTTP status codes to indicate success: http.StatusOK
	LockImmutabilityPolicy func(ctx context.Context, resourceGroupName string, accountName string, containerName string, ifMatch string, options *armstorage.BlobContainersClientLockImmutabilityPolicyOptions) (resp azfake.Responder[armstorage.BlobContainersClientLockImmutabilityPolicyResponse], errResp azfake.ErrorResponder)

	// BeginObjectLevelWorm is the fake for method BlobContainersClient.BeginObjectLevelWorm
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginObjectLevelWorm func(ctx context.Context, resourceGroupName string, accountName string, containerName string, options *armstorage.BlobContainersClientBeginObjectLevelWormOptions) (resp azfake.PollerResponder[armstorage.BlobContainersClientObjectLevelWormResponse], errResp azfake.ErrorResponder)

	// SetLegalHold is the fake for method BlobContainersClient.SetLegalHold
	// HTTP status codes to indicate success: http.StatusOK
	SetLegalHold func(ctx context.Context, resourceGroupName string, accountName string, containerName string, legalHold armstorage.LegalHold, options *armstorage.BlobContainersClientSetLegalHoldOptions) (resp azfake.Responder[armstorage.BlobContainersClientSetLegalHoldResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method BlobContainersClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, containerName string, blobContainer armstorage.BlobContainer, options *armstorage.BlobContainersClientUpdateOptions) (resp azfake.Responder[armstorage.BlobContainersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewBlobContainersServerTransport creates a new instance of BlobContainersServerTransport with the provided implementation.
// The returned BlobContainersServerTransport instance is connected to an instance of armstorage.BlobContainersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBlobContainersServerTransport(srv *BlobContainersServer) *BlobContainersServerTransport {
	return &BlobContainersServerTransport{
		srv:                  srv,
		newListPager:         newTracker[azfake.PagerResponder[armstorage.BlobContainersClientListResponse]](),
		beginObjectLevelWorm: newTracker[azfake.PollerResponder[armstorage.BlobContainersClientObjectLevelWormResponse]](),
	}
}

// BlobContainersServerTransport connects instances of armstorage.BlobContainersClient to instances of BlobContainersServer.
// Don't use this type directly, use NewBlobContainersServerTransport instead.
type BlobContainersServerTransport struct {
	srv                  *BlobContainersServer
	newListPager         *tracker[azfake.PagerResponder[armstorage.BlobContainersClientListResponse]]
	beginObjectLevelWorm *tracker[azfake.PollerResponder[armstorage.BlobContainersClientObjectLevelWormResponse]]
}

// Do implements the policy.Transporter interface for BlobContainersServerTransport.
func (b *BlobContainersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BlobContainersClient.ClearLegalHold":
		resp, err = b.dispatchClearLegalHold(req)
	case "BlobContainersClient.Create":
		resp, err = b.dispatchCreate(req)
	case "BlobContainersClient.CreateOrUpdateImmutabilityPolicy":
		resp, err = b.dispatchCreateOrUpdateImmutabilityPolicy(req)
	case "BlobContainersClient.Delete":
		resp, err = b.dispatchDelete(req)
	case "BlobContainersClient.DeleteImmutabilityPolicy":
		resp, err = b.dispatchDeleteImmutabilityPolicy(req)
	case "BlobContainersClient.ExtendImmutabilityPolicy":
		resp, err = b.dispatchExtendImmutabilityPolicy(req)
	case "BlobContainersClient.Get":
		resp, err = b.dispatchGet(req)
	case "BlobContainersClient.GetImmutabilityPolicy":
		resp, err = b.dispatchGetImmutabilityPolicy(req)
	case "BlobContainersClient.Lease":
		resp, err = b.dispatchLease(req)
	case "BlobContainersClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	case "BlobContainersClient.LockImmutabilityPolicy":
		resp, err = b.dispatchLockImmutabilityPolicy(req)
	case "BlobContainersClient.BeginObjectLevelWorm":
		resp, err = b.dispatchBeginObjectLevelWorm(req)
	case "BlobContainersClient.SetLegalHold":
		resp, err = b.dispatchSetLegalHold(req)
	case "BlobContainersClient.Update":
		resp, err = b.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchClearLegalHold(req *http.Request) (*http.Response, error) {
	if b.srv.ClearLegalHold == nil {
		return nil, &nonRetriableError{errors.New("fake for method ClearLegalHold not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clearLegalHold`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.LegalHold](req)
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.ClearLegalHold(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LegalHold, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if b.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.BlobContainer](req)
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BlobContainer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchCreateOrUpdateImmutabilityPolicy(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrUpdateImmutabilityPolicy == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateImmutabilityPolicy not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/immutabilityPolicies/(?P<immutabilityPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.ImmutabilityPolicy](req)
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyOptions
	if ifMatchParam != nil || !reflect.ValueOf(body).IsZero() {
		options = &armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyOptions{
			IfMatch:    ifMatchParam,
			Parameters: &body,
		}
	}
	respr, errRespr := b.srv.CreateOrUpdateImmutabilityPolicy(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImmutabilityPolicy, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if b.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Delete(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchDeleteImmutabilityPolicy(req *http.Request) (*http.Response, error) {
	if b.srv.DeleteImmutabilityPolicy == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteImmutabilityPolicy not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/immutabilityPolicies/(?P<immutabilityPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.DeleteImmutabilityPolicy(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, getHeaderValue(req.Header, "If-Match"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImmutabilityPolicy, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchExtendImmutabilityPolicy(req *http.Request) (*http.Response, error) {
	if b.srv.ExtendImmutabilityPolicy == nil {
		return nil, &nonRetriableError{errors.New("fake for method ExtendImmutabilityPolicy not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/immutabilityPolicies/default/extend`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.ImmutabilityPolicy](req)
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	var options *armstorage.BlobContainersClientExtendImmutabilityPolicyOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armstorage.BlobContainersClientExtendImmutabilityPolicyOptions{
			Parameters: &body,
		}
	}
	respr, errRespr := b.srv.ExtendImmutabilityPolicy(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, getHeaderValue(req.Header, "If-Match"), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImmutabilityPolicy, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BlobContainer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchGetImmutabilityPolicy(req *http.Request) (*http.Response, error) {
	if b.srv.GetImmutabilityPolicy == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetImmutabilityPolicy not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/immutabilityPolicies/(?P<immutabilityPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armstorage.BlobContainersClientGetImmutabilityPolicyOptions
	if ifMatchParam != nil {
		options = &armstorage.BlobContainersClientGetImmutabilityPolicyOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := b.srv.GetImmutabilityPolicy(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImmutabilityPolicy, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchLease(req *http.Request) (*http.Response, error) {
	if b.srv.Lease == nil {
		return nil, &nonRetriableError{errors.New("fake for method Lease not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/lease`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.LeaseContainerRequest](req)
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	var options *armstorage.BlobContainersClientLeaseOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armstorage.BlobContainersClientLeaseOptions{
			Parameters: &body,
		}
	}
	respr, errRespr := b.srv.Lease(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LeaseContainerResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers`
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
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam := getOptional(maxpagesizeUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		includeUnescaped, err := url.QueryUnescape(qp.Get("$include"))
		if err != nil {
			return nil, err
		}
		includeParam := getOptional(armstorage.ListContainersInclude(includeUnescaped))
		var options *armstorage.BlobContainersClientListOptions
		if maxpagesizeParam != nil || filterParam != nil || includeParam != nil {
			options = &armstorage.BlobContainersClientListOptions{
				Maxpagesize: maxpagesizeParam,
				Filter:      filterParam,
				Include:     includeParam,
			}
		}
		resp := b.srv.NewListPager(resourceGroupNameParam, accountNameParam, options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armstorage.BlobContainersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchLockImmutabilityPolicy(req *http.Request) (*http.Response, error) {
	if b.srv.LockImmutabilityPolicy == nil {
		return nil, &nonRetriableError{errors.New("fake for method LockImmutabilityPolicy not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/immutabilityPolicies/default/lock`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.LockImmutabilityPolicy(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, getHeaderValue(req.Header, "If-Match"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImmutabilityPolicy, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchBeginObjectLevelWorm(req *http.Request) (*http.Response, error) {
	if b.srv.BeginObjectLevelWorm == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginObjectLevelWorm not implemented")}
	}
	beginObjectLevelWorm := b.beginObjectLevelWorm.get(req)
	if beginObjectLevelWorm == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/migrate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginObjectLevelWorm(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginObjectLevelWorm = &respr
		b.beginObjectLevelWorm.add(req, beginObjectLevelWorm)
	}

	resp, err := server.PollerResponderNext(beginObjectLevelWorm, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		b.beginObjectLevelWorm.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginObjectLevelWorm) {
		b.beginObjectLevelWorm.remove(req)
	}

	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchSetLegalHold(req *http.Request) (*http.Response, error) {
	if b.srv.SetLegalHold == nil {
		return nil, &nonRetriableError{errors.New("fake for method SetLegalHold not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/setLegalHold`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.LegalHold](req)
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.SetLegalHold(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LegalHold, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BlobContainersServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobServices/default/containers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.BlobContainer](req)
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Update(req.Context(), resourceGroupNameParam, accountNameParam, containerNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BlobContainer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
