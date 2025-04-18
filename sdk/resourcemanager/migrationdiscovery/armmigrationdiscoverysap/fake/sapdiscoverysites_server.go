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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscoverysap"
	"net/http"
	"net/url"
	"regexp"
)

// SapDiscoverySitesServer is a fake server for instances of the armmigrationdiscoverysap.SapDiscoverySitesClient type.
type SapDiscoverySitesServer struct {
	// BeginCreate is the fake for method SapDiscoverySitesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, resource armmigrationdiscoverysap.SAPDiscoverySite, options *armmigrationdiscoverysap.SapDiscoverySitesClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SapDiscoverySitesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *armmigrationdiscoverysap.SapDiscoverySitesClientBeginDeleteOptions) (resp azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SapDiscoverySitesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *armmigrationdiscoverysap.SapDiscoverySitesClientGetOptions) (resp azfake.Responder[armmigrationdiscoverysap.SapDiscoverySitesClientGetResponse], errResp azfake.ErrorResponder)

	// BeginImportEntities is the fake for method SapDiscoverySitesClient.BeginImportEntities
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginImportEntities func(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, options *armmigrationdiscoverysap.SapDiscoverySitesClientBeginImportEntitiesOptions) (resp azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientImportEntitiesResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method SapDiscoverySitesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmigrationdiscoverysap.SapDiscoverySitesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method SapDiscoverySitesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmigrationdiscoverysap.SapDiscoverySitesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientListBySubscriptionResponse])

	// Update is the fake for method SapDiscoverySitesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, sapDiscoverySiteName string, properties armmigrationdiscoverysap.SAPDiscoverySiteTagsUpdate, options *armmigrationdiscoverysap.SapDiscoverySitesClientUpdateOptions) (resp azfake.Responder[armmigrationdiscoverysap.SapDiscoverySitesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSapDiscoverySitesServerTransport creates a new instance of SapDiscoverySitesServerTransport with the provided implementation.
// The returned SapDiscoverySitesServerTransport instance is connected to an instance of armmigrationdiscoverysap.SapDiscoverySitesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSapDiscoverySitesServerTransport(srv *SapDiscoverySitesServer) *SapDiscoverySitesServerTransport {
	return &SapDiscoverySitesServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientDeleteResponse]](),
		beginImportEntities:         newTracker[azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientImportEntitiesResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientListBySubscriptionResponse]](),
	}
}

// SapDiscoverySitesServerTransport connects instances of armmigrationdiscoverysap.SapDiscoverySitesClient to instances of SapDiscoverySitesServer.
// Don't use this type directly, use NewSapDiscoverySitesServerTransport instead.
type SapDiscoverySitesServerTransport struct {
	srv                         *SapDiscoverySitesServer
	beginCreate                 *tracker[azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientDeleteResponse]]
	beginImportEntities         *tracker[azfake.PollerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientImportEntitiesResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmigrationdiscoverysap.SapDiscoverySitesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for SapDiscoverySitesServerTransport.
func (s *SapDiscoverySitesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SapDiscoverySitesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sapDiscoverySitesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sapDiscoverySitesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SapDiscoverySitesClient.BeginCreate":
				res.resp, res.err = s.dispatchBeginCreate(req)
			case "SapDiscoverySitesClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "SapDiscoverySitesClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SapDiscoverySitesClient.BeginImportEntities":
				res.resp, res.err = s.dispatchBeginImportEntities(req)
			case "SapDiscoverySitesClient.NewListByResourceGroupPager":
				res.resp, res.err = s.dispatchNewListByResourceGroupPager(req)
			case "SapDiscoverySitesClient.NewListBySubscriptionPager":
				res.resp, res.err = s.dispatchNewListBySubscriptionPager(req)
			case "SapDiscoverySitesClient.Update":
				res.resp, res.err = s.dispatchUpdate(req)
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

func (s *SapDiscoverySitesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapDiscoverySites/(?P<sapDiscoverySiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscoverysap.SAPDiscoverySite](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapDiscoverySiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapDiscoverySiteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreate(req.Context(), resourceGroupNameParam, sapDiscoverySiteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		s.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		s.beginCreate.remove(req)
	}

	return resp, nil
}

func (s *SapDiscoverySitesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapDiscoverySites/(?P<sapDiscoverySiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapDiscoverySiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapDiscoverySiteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, sapDiscoverySiteNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SapDiscoverySitesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapDiscoverySites/(?P<sapDiscoverySiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	sapDiscoverySiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapDiscoverySiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, sapDiscoverySiteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPDiscoverySite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SapDiscoverySitesServerTransport) dispatchBeginImportEntities(req *http.Request) (*http.Response, error) {
	if s.srv.BeginImportEntities == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginImportEntities not implemented")}
	}
	beginImportEntities := s.beginImportEntities.get(req)
	if beginImportEntities == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapDiscoverySites/(?P<sapDiscoverySiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importEntities`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapDiscoverySiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapDiscoverySiteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginImportEntities(req.Context(), resourceGroupNameParam, sapDiscoverySiteNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginImportEntities = &respr
		s.beginImportEntities.add(req, beginImportEntities)
	}

	resp, err := server.PollerResponderNext(beginImportEntities, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginImportEntities.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginImportEntities) {
		s.beginImportEntities.remove(req)
	}

	return resp, nil
}

func (s *SapDiscoverySitesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapDiscoverySites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		s.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmigrationdiscoverysap.SapDiscoverySitesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		s.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (s *SapDiscoverySitesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := s.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapDiscoverySites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		s.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmigrationdiscoverysap.SapDiscoverySitesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		s.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (s *SapDiscoverySitesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapDiscoverySites/(?P<sapDiscoverySiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrationdiscoverysap.SAPDiscoverySiteTagsUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	sapDiscoverySiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapDiscoverySiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Update(req.Context(), resourceGroupNameParam, sapDiscoverySiteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPDiscoverySite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SapDiscoverySitesServerTransport
var sapDiscoverySitesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
