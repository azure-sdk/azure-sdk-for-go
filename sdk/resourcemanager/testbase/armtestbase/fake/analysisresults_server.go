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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
	"net/http"
	"net/url"
	"regexp"
)

// AnalysisResultsServer is a fake server for instances of the armtestbase.AnalysisResultsClient type.
type AnalysisResultsServer struct {
	// Get is the fake for method AnalysisResultsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string, analysisResultName armtestbase.AnalysisResultName, options *armtestbase.AnalysisResultsClientGetOptions) (resp azfake.Responder[armtestbase.AnalysisResultsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AnalysisResultsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, testBaseAccountName string, packageName string, testResultName string, analysisResultType armtestbase.AnalysisResultType, options *armtestbase.AnalysisResultsClientListOptions) (resp azfake.PagerResponder[armtestbase.AnalysisResultsClientListResponse])
}

// NewAnalysisResultsServerTransport creates a new instance of AnalysisResultsServerTransport with the provided implementation.
// The returned AnalysisResultsServerTransport instance is connected to an instance of armtestbase.AnalysisResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAnalysisResultsServerTransport(srv *AnalysisResultsServer) *AnalysisResultsServerTransport {
	return &AnalysisResultsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armtestbase.AnalysisResultsClientListResponse]](),
	}
}

// AnalysisResultsServerTransport connects instances of armtestbase.AnalysisResultsClient to instances of AnalysisResultsServer.
// Don't use this type directly, use NewAnalysisResultsServerTransport instead.
type AnalysisResultsServerTransport struct {
	srv          *AnalysisResultsServer
	newListPager *tracker[azfake.PagerResponder[armtestbase.AnalysisResultsClientListResponse]]
}

// Do implements the policy.Transporter interface for AnalysisResultsServerTransport.
func (a *AnalysisResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AnalysisResultsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if analysisResultsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = analysisResultsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AnalysisResultsClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "AnalysisResultsClient.NewListPager":
				res.resp, res.err = a.dispatchNewListPager(req)
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

func (a *AnalysisResultsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packages/(?P<packageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/testResults/(?P<testResultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/analysisResults/(?P<analysisResultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
	if err != nil {
		return nil, err
	}
	packageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packageName")])
	if err != nil {
		return nil, err
	}
	testResultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testResultName")])
	if err != nil {
		return nil, err
	}
	analysisResultNameParam, err := parseWithCast(matches[regex.SubexpIndex("analysisResultName")], func(v string) (armtestbase.AnalysisResultName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armtestbase.AnalysisResultName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, testBaseAccountNameParam, packageNameParam, testResultNameParam, analysisResultNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AnalysisResultSingletonResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AnalysisResultsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packages/(?P<packageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/testResults/(?P<testResultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/analysisResults`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
		if err != nil {
			return nil, err
		}
		packageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("packageName")])
		if err != nil {
			return nil, err
		}
		testResultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testResultName")])
		if err != nil {
			return nil, err
		}
		analysisResultTypeParam, err := parseWithCast(qp.Get("analysisResultType"), func(v string) (armtestbase.AnalysisResultType, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armtestbase.AnalysisResultType(p), nil
		})
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, testBaseAccountNameParam, packageNameParam, testResultNameParam, analysisResultTypeParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AnalysisResultsServerTransport
var analysisResultsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
