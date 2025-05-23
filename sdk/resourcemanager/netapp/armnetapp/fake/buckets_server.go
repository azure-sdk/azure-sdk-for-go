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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
	"net/http"
	"net/url"
	"regexp"
)

// BucketsServer is a fake server for instances of the armnetapp.BucketsClient type.
type BucketsServer struct {
	// BeginCreateOrUpdate is the fake for method BucketsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, bucketName string, body armnetapp.Bucket, options *armnetapp.BucketsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetapp.BucketsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BucketsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, bucketName string, options *armnetapp.BucketsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetapp.BucketsClientDeleteResponse], errResp azfake.ErrorResponder)

	// GenerateCredentials is the fake for method BucketsClient.GenerateCredentials
	// HTTP status codes to indicate success: http.StatusOK
	GenerateCredentials func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, bucketName string, body armnetapp.BucketCredentialsExpiry, options *armnetapp.BucketsClientGenerateCredentialsOptions) (resp azfake.Responder[armnetapp.BucketsClientGenerateCredentialsResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BucketsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, bucketName string, options *armnetapp.BucketsClientGetOptions) (resp azfake.Responder[armnetapp.BucketsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BucketsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, poolName string, volumeName string, options *armnetapp.BucketsClientListOptions) (resp azfake.PagerResponder[armnetapp.BucketsClientListResponse])

	// BeginUpdate is the fake for method BucketsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, bucketName string, body armnetapp.BucketPatch, options *armnetapp.BucketsClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetapp.BucketsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewBucketsServerTransport creates a new instance of BucketsServerTransport with the provided implementation.
// The returned BucketsServerTransport instance is connected to an instance of armnetapp.BucketsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBucketsServerTransport(srv *BucketsServer) *BucketsServerTransport {
	return &BucketsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetapp.BucketsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetapp.BucketsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetapp.BucketsClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armnetapp.BucketsClientUpdateResponse]](),
	}
}

// BucketsServerTransport connects instances of armnetapp.BucketsClient to instances of BucketsServer.
// Don't use this type directly, use NewBucketsServerTransport instead.
type BucketsServerTransport struct {
	srv                 *BucketsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetapp.BucketsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetapp.BucketsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetapp.BucketsClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armnetapp.BucketsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for BucketsServerTransport.
func (b *BucketsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BucketsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if bucketsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = bucketsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BucketsClient.BeginCreateOrUpdate":
				res.resp, res.err = b.dispatchBeginCreateOrUpdate(req)
			case "BucketsClient.BeginDelete":
				res.resp, res.err = b.dispatchBeginDelete(req)
			case "BucketsClient.GenerateCredentials":
				res.resp, res.err = b.dispatchGenerateCredentials(req)
			case "BucketsClient.Get":
				res.resp, res.err = b.dispatchGet(req)
			case "BucketsClient.NewListPager":
				res.resp, res.err = b.dispatchNewListPager(req)
			case "BucketsClient.BeginUpdate":
				res.resp, res.err = b.dispatchBeginUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (b *BucketsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buckets/(?P<bucketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.Bucket](req)
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		bucketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bucketName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, bucketNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *BucketsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buckets/(?P<bucketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		bucketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bucketName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, bucketNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BucketsServerTransport) dispatchGenerateCredentials(req *http.Request) (*http.Response, error) {
	if b.srv.GenerateCredentials == nil {
		return nil, &nonRetriableError{errors.New("fake for method GenerateCredentials not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buckets/(?P<bucketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateCredentials`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetapp.BucketCredentialsExpiry](req)
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
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
	if err != nil {
		return nil, err
	}
	bucketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bucketName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.GenerateCredentials(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, bucketNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BucketGenerateCredentials, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BucketsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buckets/(?P<bucketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
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
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
	if err != nil {
		return nil, err
	}
	bucketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bucketName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, bucketNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Bucket, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BucketsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buckets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListPager(resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetapp.BucketsClientListResponse, createLink func() string) {
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

func (b *BucketsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := b.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buckets/(?P<bucketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.BucketPatch](req)
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		bucketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bucketName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginUpdate(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, bucketNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		b.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		b.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		b.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to BucketsServerTransport
var bucketsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
