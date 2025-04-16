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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedInstanceEncryptionProtectorsServer is a fake server for instances of the armsql.ManagedInstanceEncryptionProtectorsClient type.
type ManagedInstanceEncryptionProtectorsServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedInstanceEncryptionProtectorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName armsql.EncryptionProtectorName, parameters armsql.ManagedInstanceEncryptionProtector, options *armsql.ManagedInstanceEncryptionProtectorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedInstanceEncryptionProtectorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedInstanceEncryptionProtectorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName armsql.EncryptionProtectorName, options *armsql.ManagedInstanceEncryptionProtectorsClientGetOptions) (resp azfake.Responder[armsql.ManagedInstanceEncryptionProtectorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByInstancePager is the fake for method ManagedInstanceEncryptionProtectorsClient.NewListByInstancePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInstancePager func(resourceGroupName string, managedInstanceName string, options *armsql.ManagedInstanceEncryptionProtectorsClientListByInstanceOptions) (resp azfake.PagerResponder[armsql.ManagedInstanceEncryptionProtectorsClientListByInstanceResponse])

	// BeginRevalidate is the fake for method ManagedInstanceEncryptionProtectorsClient.BeginRevalidate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginRevalidate func(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName armsql.EncryptionProtectorName, options *armsql.ManagedInstanceEncryptionProtectorsClientBeginRevalidateOptions) (resp azfake.PollerResponder[armsql.ManagedInstanceEncryptionProtectorsClientRevalidateResponse], errResp azfake.ErrorResponder)
}

// NewManagedInstanceEncryptionProtectorsServerTransport creates a new instance of ManagedInstanceEncryptionProtectorsServerTransport with the provided implementation.
// The returned ManagedInstanceEncryptionProtectorsServerTransport instance is connected to an instance of armsql.ManagedInstanceEncryptionProtectorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedInstanceEncryptionProtectorsServerTransport(srv *ManagedInstanceEncryptionProtectorsServer) *ManagedInstanceEncryptionProtectorsServerTransport {
	return &ManagedInstanceEncryptionProtectorsServerTransport{
		srv:                    srv,
		beginCreateOrUpdate:    newTracker[azfake.PollerResponder[armsql.ManagedInstanceEncryptionProtectorsClientCreateOrUpdateResponse]](),
		newListByInstancePager: newTracker[azfake.PagerResponder[armsql.ManagedInstanceEncryptionProtectorsClientListByInstanceResponse]](),
		beginRevalidate:        newTracker[azfake.PollerResponder[armsql.ManagedInstanceEncryptionProtectorsClientRevalidateResponse]](),
	}
}

// ManagedInstanceEncryptionProtectorsServerTransport connects instances of armsql.ManagedInstanceEncryptionProtectorsClient to instances of ManagedInstanceEncryptionProtectorsServer.
// Don't use this type directly, use NewManagedInstanceEncryptionProtectorsServerTransport instead.
type ManagedInstanceEncryptionProtectorsServerTransport struct {
	srv                    *ManagedInstanceEncryptionProtectorsServer
	beginCreateOrUpdate    *tracker[azfake.PollerResponder[armsql.ManagedInstanceEncryptionProtectorsClientCreateOrUpdateResponse]]
	newListByInstancePager *tracker[azfake.PagerResponder[armsql.ManagedInstanceEncryptionProtectorsClientListByInstanceResponse]]
	beginRevalidate        *tracker[azfake.PollerResponder[armsql.ManagedInstanceEncryptionProtectorsClientRevalidateResponse]]
}

// Do implements the policy.Transporter interface for ManagedInstanceEncryptionProtectorsServerTransport.
func (m *ManagedInstanceEncryptionProtectorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *ManagedInstanceEncryptionProtectorsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if managedInstanceEncryptionProtectorsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = managedInstanceEncryptionProtectorsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ManagedInstanceEncryptionProtectorsClient.BeginCreateOrUpdate":
				res.resp, res.err = m.dispatchBeginCreateOrUpdate(req)
			case "ManagedInstanceEncryptionProtectorsClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "ManagedInstanceEncryptionProtectorsClient.NewListByInstancePager":
				res.resp, res.err = m.dispatchNewListByInstancePager(req)
			case "ManagedInstanceEncryptionProtectorsClient.BeginRevalidate":
				res.resp, res.err = m.dispatchBeginRevalidate(req)
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

func (m *ManagedInstanceEncryptionProtectorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/encryptionProtector/(?P<encryptionProtectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedInstanceEncryptionProtector](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		encryptionProtectorNameParam, err := parseWithCast(matches[regex.SubexpIndex("encryptionProtectorName")], func(v string) (armsql.EncryptionProtectorName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.EncryptionProtectorName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, encryptionProtectorNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedInstanceEncryptionProtectorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/encryptionProtector/(?P<encryptionProtectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	encryptionProtectorNameParam, err := parseWithCast(matches[regex.SubexpIndex("encryptionProtectorName")], func(v string) (armsql.EncryptionProtectorName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.EncryptionProtectorName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, encryptionProtectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedInstanceEncryptionProtector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedInstanceEncryptionProtectorsServerTransport) dispatchNewListByInstancePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByInstancePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInstancePager not implemented")}
	}
	newListByInstancePager := m.newListByInstancePager.get(req)
	if newListByInstancePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/encryptionProtector`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByInstancePager(resourceGroupNameParam, managedInstanceNameParam, nil)
		newListByInstancePager = &resp
		m.newListByInstancePager.add(req, newListByInstancePager)
		server.PagerResponderInjectNextLinks(newListByInstancePager, req, func(page *armsql.ManagedInstanceEncryptionProtectorsClientListByInstanceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInstancePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByInstancePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInstancePager) {
		m.newListByInstancePager.remove(req)
	}
	return resp, nil
}

func (m *ManagedInstanceEncryptionProtectorsServerTransport) dispatchBeginRevalidate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginRevalidate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRevalidate not implemented")}
	}
	beginRevalidate := m.beginRevalidate.get(req)
	if beginRevalidate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/encryptionProtector/(?P<encryptionProtectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revalidate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		encryptionProtectorNameParam, err := parseWithCast(matches[regex.SubexpIndex("encryptionProtectorName")], func(v string) (armsql.EncryptionProtectorName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.EncryptionProtectorName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginRevalidate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, encryptionProtectorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRevalidate = &respr
		m.beginRevalidate.add(req, beginRevalidate)
	}

	resp, err := server.PollerResponderNext(beginRevalidate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginRevalidate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRevalidate) {
		m.beginRevalidate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ManagedInstanceEncryptionProtectorsServerTransport
var managedInstanceEncryptionProtectorsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
