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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// IotSecuritySolutionsAnalyticsRecommendationServer is a fake server for instances of the armsecurity.IotSecuritySolutionsAnalyticsRecommendationClient type.
type IotSecuritySolutionsAnalyticsRecommendationServer struct {
	// Get is the fake for method IotSecuritySolutionsAnalyticsRecommendationClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, solutionName string, aggregatedRecommendationName string, options *armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientGetOptions) (resp azfake.Responder[armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IotSecuritySolutionsAnalyticsRecommendationClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, solutionName string, options *armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientListOptions) (resp azfake.PagerResponder[armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientListResponse])
}

// NewIotSecuritySolutionsAnalyticsRecommendationServerTransport creates a new instance of IotSecuritySolutionsAnalyticsRecommendationServerTransport with the provided implementation.
// The returned IotSecuritySolutionsAnalyticsRecommendationServerTransport instance is connected to an instance of armsecurity.IotSecuritySolutionsAnalyticsRecommendationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIotSecuritySolutionsAnalyticsRecommendationServerTransport(srv *IotSecuritySolutionsAnalyticsRecommendationServer) *IotSecuritySolutionsAnalyticsRecommendationServerTransport {
	return &IotSecuritySolutionsAnalyticsRecommendationServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientListResponse]](),
	}
}

// IotSecuritySolutionsAnalyticsRecommendationServerTransport connects instances of armsecurity.IotSecuritySolutionsAnalyticsRecommendationClient to instances of IotSecuritySolutionsAnalyticsRecommendationServer.
// Don't use this type directly, use NewIotSecuritySolutionsAnalyticsRecommendationServerTransport instead.
type IotSecuritySolutionsAnalyticsRecommendationServerTransport struct {
	srv          *IotSecuritySolutionsAnalyticsRecommendationServer
	newListPager *tracker[azfake.PagerResponder[armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientListResponse]]
}

// Do implements the policy.Transporter interface for IotSecuritySolutionsAnalyticsRecommendationServerTransport.
func (i *IotSecuritySolutionsAnalyticsRecommendationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IotSecuritySolutionsAnalyticsRecommendationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if iotSecuritySolutionsAnalyticsRecommendationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = iotSecuritySolutionsAnalyticsRecommendationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "IotSecuritySolutionsAnalyticsRecommendationClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "IotSecuritySolutionsAnalyticsRecommendationClient.NewListPager":
				res.resp, res.err = i.dispatchNewListPager(req)
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

func (i *IotSecuritySolutionsAnalyticsRecommendationServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/iotSecuritySolutions/(?P<solutionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/analyticsModels/default/aggregatedRecommendations/(?P<aggregatedRecommendationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	solutionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("solutionName")])
	if err != nil {
		return nil, err
	}
	aggregatedRecommendationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("aggregatedRecommendationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, solutionNameParam, aggregatedRecommendationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IoTSecurityAggregatedRecommendation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IotSecuritySolutionsAnalyticsRecommendationServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/iotSecuritySolutions/(?P<solutionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/analyticsModels/default/aggregatedRecommendations`
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
		solutionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("solutionName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientListOptions
		if topParam != nil {
			options = &armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientListOptions{
				Top: topParam,
			}
		}
		resp := i.srv.NewListPager(resourceGroupNameParam, solutionNameParam, options)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to IotSecuritySolutionsAnalyticsRecommendationServerTransport
var iotSecuritySolutionsAnalyticsRecommendationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
