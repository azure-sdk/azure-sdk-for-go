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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ProtectedItemServer is a fake server for instances of the armrecoveryservicesdatareplication.ProtectedItemClient type.
type ProtectedItemServer struct {
	// BeginCreate is the fake for method ProtectedItemClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, resource armrecoveryservicesdatareplication.ProtectedItemModel, options *armrecoveryservicesdatareplication.ProtectedItemClientBeginCreateOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ProtectedItemClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, options *armrecoveryservicesdatareplication.ProtectedItemClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProtectedItemClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, options *armrecoveryservicesdatareplication.ProtectedItemClientGetOptions) (resp azfake.Responder[armrecoveryservicesdatareplication.ProtectedItemClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ProtectedItemClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vaultName string, options *armrecoveryservicesdatareplication.ProtectedItemClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesdatareplication.ProtectedItemClientListResponse])

	// BeginPlannedFailover is the fake for method ProtectedItemClient.BeginPlannedFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPlannedFailover func(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, body armrecoveryservicesdatareplication.PlannedFailoverModel, options *armrecoveryservicesdatareplication.ProtectedItemClientBeginPlannedFailoverOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientPlannedFailoverResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ProtectedItemClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, properties armrecoveryservicesdatareplication.ProtectedItemModelUpdate, options *armrecoveryservicesdatareplication.ProtectedItemClientBeginUpdateOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewProtectedItemServerTransport creates a new instance of ProtectedItemServerTransport with the provided implementation.
// The returned ProtectedItemServerTransport instance is connected to an instance of armrecoveryservicesdatareplication.ProtectedItemClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProtectedItemServerTransport(srv *ProtectedItemServer) *ProtectedItemServerTransport {
	return &ProtectedItemServerTransport{
		srv:                  srv,
		beginCreate:          newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientCreateResponse]](),
		beginDelete:          newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientDeleteResponse]](),
		newListPager:         newTracker[azfake.PagerResponder[armrecoveryservicesdatareplication.ProtectedItemClientListResponse]](),
		beginPlannedFailover: newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientPlannedFailoverResponse]](),
		beginUpdate:          newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientUpdateResponse]](),
	}
}

// ProtectedItemServerTransport connects instances of armrecoveryservicesdatareplication.ProtectedItemClient to instances of ProtectedItemServer.
// Don't use this type directly, use NewProtectedItemServerTransport instead.
type ProtectedItemServerTransport struct {
	srv                  *ProtectedItemServer
	beginCreate          *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientCreateResponse]]
	beginDelete          *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientDeleteResponse]]
	newListPager         *tracker[azfake.PagerResponder[armrecoveryservicesdatareplication.ProtectedItemClientListResponse]]
	beginPlannedFailover *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientPlannedFailoverResponse]]
	beginUpdate          *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.ProtectedItemClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ProtectedItemServerTransport.
func (p *ProtectedItemServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProtectedItemServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if protectedItemServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = protectedItemServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProtectedItemClient.BeginCreate":
				res.resp, res.err = p.dispatchBeginCreate(req)
			case "ProtectedItemClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "ProtectedItemClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "ProtectedItemClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
			case "ProtectedItemClient.BeginPlannedFailover":
				res.resp, res.err = p.dispatchBeginPlannedFailover(req)
			case "ProtectedItemClient.BeginUpdate":
				res.resp, res.err = p.dispatchBeginUpdate(req)
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

func (p *ProtectedItemServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := p.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesdatareplication.ProtectedItemModel](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreate(req.Context(), resourceGroupNameParam, vaultNameParam, protectedItemNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		p.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		p.beginCreate.remove(req)
	}

	return resp, nil
}

func (p *ProtectedItemServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		forceDeleteUnescaped, err := url.QueryUnescape(qp.Get("forceDelete"))
		if err != nil {
			return nil, err
		}
		forceDeleteParam, err := parseOptional(forceDeleteUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
		if err != nil {
			return nil, err
		}
		var options *armrecoveryservicesdatareplication.ProtectedItemClientBeginDeleteOptions
		if forceDeleteParam != nil {
			options = &armrecoveryservicesdatareplication.ProtectedItemClientBeginDeleteOptions{
				ForceDelete: forceDeleteParam,
			}
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, vaultNameParam, protectedItemNameParam, options)
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

func (p *ProtectedItemServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, protectedItemNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectedItemModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectedItemServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems`
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
		odataOptionsUnescaped, err := url.QueryUnescape(qp.Get("odataOptions"))
		if err != nil {
			return nil, err
		}
		odataOptionsParam := getOptional(odataOptionsUnescaped)
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		pageSizeUnescaped, err := url.QueryUnescape(qp.Get("pageSize"))
		if err != nil {
			return nil, err
		}
		pageSizeParam, err := parseOptional(pageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		var options *armrecoveryservicesdatareplication.ProtectedItemClientListOptions
		if odataOptionsParam != nil || continuationTokenParam != nil || pageSizeParam != nil {
			options = &armrecoveryservicesdatareplication.ProtectedItemClientListOptions{
				OdataOptions:      odataOptionsParam,
				ContinuationToken: continuationTokenParam,
				PageSize:          pageSizeParam,
			}
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, vaultNameParam, options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicesdatareplication.ProtectedItemClientListResponse, createLink func() string) {
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

func (p *ProtectedItemServerTransport) dispatchBeginPlannedFailover(req *http.Request) (*http.Response, error) {
	if p.srv.BeginPlannedFailover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPlannedFailover not implemented")}
	}
	beginPlannedFailover := p.beginPlannedFailover.get(req)
	if beginPlannedFailover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/plannedFailover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesdatareplication.PlannedFailoverModel](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginPlannedFailover(req.Context(), resourceGroupNameParam, vaultNameParam, protectedItemNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPlannedFailover = &respr
		p.beginPlannedFailover.add(req, beginPlannedFailover)
	}

	resp, err := server.PollerResponderNext(beginPlannedFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginPlannedFailover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPlannedFailover) {
		p.beginPlannedFailover.remove(req)
	}

	return resp, nil
}

func (p *ProtectedItemServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesdatareplication.ProtectedItemModelUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, vaultNameParam, protectedItemNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ProtectedItemServerTransport
var protectedItemServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
