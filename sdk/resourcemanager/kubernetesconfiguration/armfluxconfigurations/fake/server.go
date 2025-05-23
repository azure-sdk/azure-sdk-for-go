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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armfluxconfigurations"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// Server is a fake server for instances of the armfluxconfigurations.Client type.
type Server struct {
	// BeginCreateOrUpdate is the fake for method Client.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfiguration armfluxconfigurations.FluxConfiguration, options *armfluxconfigurations.ClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armfluxconfigurations.ClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method Client.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, options *armfluxconfigurations.ClientBeginDeleteOptions) (resp azfake.PollerResponder[armfluxconfigurations.ClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method Client.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, options *armfluxconfigurations.ClientGetOptions) (resp azfake.Responder[armfluxconfigurations.ClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method Client.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *armfluxconfigurations.ClientListOptions) (resp azfake.PagerResponder[armfluxconfigurations.ClientListResponse])

	// BeginUpdate is the fake for method Client.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfigurationPatch armfluxconfigurations.FluxConfigurationPatch, options *armfluxconfigurations.ClientBeginUpdateOptions) (resp azfake.PollerResponder[armfluxconfigurations.ClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armfluxconfigurations.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armfluxconfigurations.ClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armfluxconfigurations.ClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armfluxconfigurations.ClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armfluxconfigurations.ClientUpdateResponse]](),
	}
}

// ServerTransport connects instances of armfluxconfigurations.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv                 *Server
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armfluxconfigurations.ClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armfluxconfigurations.ClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armfluxconfigurations.ClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armfluxconfigurations.ClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "Client.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "Client.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "Client.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "Client.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
			case "Client.BeginUpdate":
				res.resp, res.err = s.dispatchBeginUpdate(req)
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

func (s *ServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<clusterRp>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesConfiguration/fluxConfigurations/(?P<fluxConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfluxconfigurations.FluxConfiguration](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterRpParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterRp")])
		if err != nil {
			return nil, err
		}
		clusterResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterResourceName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		fluxConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fluxConfigurationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, clusterRpParam, clusterResourceNameParam, clusterNameParam, fluxConfigurationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *ServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<clusterRp>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesConfiguration/fluxConfigurations/(?P<fluxConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterRpParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterRp")])
		if err != nil {
			return nil, err
		}
		clusterResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterResourceName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		fluxConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fluxConfigurationName")])
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
		var options *armfluxconfigurations.ClientBeginDeleteOptions
		if forceDeleteParam != nil {
			options = &armfluxconfigurations.ClientBeginDeleteOptions{
				ForceDelete: forceDeleteParam,
			}
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterRpParam, clusterResourceNameParam, clusterNameParam, fluxConfigurationNameParam, options)
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

func (s *ServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<clusterRp>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesConfiguration/fluxConfigurations/(?P<fluxConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterRpParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterRp")])
	if err != nil {
		return nil, err
	}
	clusterResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterResourceName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	fluxConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fluxConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, clusterRpParam, clusterResourceNameParam, clusterNameParam, fluxConfigurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FluxConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<clusterRp>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesConfiguration/fluxConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterRpParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterRp")])
		if err != nil {
			return nil, err
		}
		clusterResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterResourceName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, clusterRpParam, clusterResourceNameParam, clusterNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armfluxconfigurations.ClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<clusterRp>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesConfiguration/fluxConfigurations/(?P<fluxConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfluxconfigurations.FluxConfigurationPatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterRpParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterRp")])
		if err != nil {
			return nil, err
		}
		clusterResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterResourceName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		fluxConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fluxConfigurationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, clusterRpParam, clusterResourceNameParam, clusterNameParam, fluxConfigurationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerTransport
var serverTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
