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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
	"net/http"
	"net/url"
	"regexp"
)

// PredictionsServer is a fake server for instances of the armcustomerinsights.PredictionsClient type.
type PredictionsServer struct {
	// BeginCreateOrUpdate is the fake for method PredictionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters armcustomerinsights.PredictionResourceFormat, options *armcustomerinsights.PredictionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcustomerinsights.PredictionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PredictionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *armcustomerinsights.PredictionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcustomerinsights.PredictionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PredictionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *armcustomerinsights.PredictionsClientGetOptions) (resp azfake.Responder[armcustomerinsights.PredictionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetModelStatus is the fake for method PredictionsClient.GetModelStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetModelStatus func(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *armcustomerinsights.PredictionsClientGetModelStatusOptions) (resp azfake.Responder[armcustomerinsights.PredictionsClientGetModelStatusResponse], errResp azfake.ErrorResponder)

	// GetTrainingResults is the fake for method PredictionsClient.GetTrainingResults
	// HTTP status codes to indicate success: http.StatusOK
	GetTrainingResults func(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *armcustomerinsights.PredictionsClientGetTrainingResultsOptions) (resp azfake.Responder[armcustomerinsights.PredictionsClientGetTrainingResultsResponse], errResp azfake.ErrorResponder)

	// NewListByHubPager is the fake for method PredictionsClient.NewListByHubPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHubPager func(resourceGroupName string, hubName string, options *armcustomerinsights.PredictionsClientListByHubOptions) (resp azfake.PagerResponder[armcustomerinsights.PredictionsClientListByHubResponse])

	// ModelStatus is the fake for method PredictionsClient.ModelStatus
	// HTTP status codes to indicate success: http.StatusOK
	ModelStatus func(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters armcustomerinsights.PredictionModelStatus, options *armcustomerinsights.PredictionsClientModelStatusOptions) (resp azfake.Responder[armcustomerinsights.PredictionsClientModelStatusResponse], errResp azfake.ErrorResponder)
}

// NewPredictionsServerTransport creates a new instance of PredictionsServerTransport with the provided implementation.
// The returned PredictionsServerTransport instance is connected to an instance of armcustomerinsights.PredictionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPredictionsServerTransport(srv *PredictionsServer) *PredictionsServerTransport {
	return &PredictionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcustomerinsights.PredictionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcustomerinsights.PredictionsClientDeleteResponse]](),
		newListByHubPager:   newTracker[azfake.PagerResponder[armcustomerinsights.PredictionsClientListByHubResponse]](),
	}
}

// PredictionsServerTransport connects instances of armcustomerinsights.PredictionsClient to instances of PredictionsServer.
// Don't use this type directly, use NewPredictionsServerTransport instead.
type PredictionsServerTransport struct {
	srv                 *PredictionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcustomerinsights.PredictionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcustomerinsights.PredictionsClientDeleteResponse]]
	newListByHubPager   *tracker[azfake.PagerResponder[armcustomerinsights.PredictionsClientListByHubResponse]]
}

// Do implements the policy.Transporter interface for PredictionsServerTransport.
func (p *PredictionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PredictionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if predictionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = predictionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PredictionsClient.BeginCreateOrUpdate":
				res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
			case "PredictionsClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PredictionsClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PredictionsClient.GetModelStatus":
				res.resp, res.err = p.dispatchGetModelStatus(req)
			case "PredictionsClient.GetTrainingResults":
				res.resp, res.err = p.dispatchGetTrainingResults(req)
			case "PredictionsClient.NewListByHubPager":
				res.resp, res.err = p.dispatchNewListByHubPager(req)
			case "PredictionsClient.ModelStatus":
				res.resp, res.err = p.dispatchModelStatus(req)
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

func (p *PredictionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predictions/(?P<predictionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcustomerinsights.PredictionResourceFormat](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		predictionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predictionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, hubNameParam, predictionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PredictionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predictions/(?P<predictionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		predictionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predictionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, hubNameParam, predictionNameParam, nil)
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

func (p *PredictionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predictions/(?P<predictionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	predictionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predictionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, hubNameParam, predictionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PredictionResourceFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PredictionsServerTransport) dispatchGetModelStatus(req *http.Request) (*http.Response, error) {
	if p.srv.GetModelStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetModelStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predictions/(?P<predictionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getModelStatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	predictionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predictionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetModelStatus(req.Context(), resourceGroupNameParam, hubNameParam, predictionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PredictionModelStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PredictionsServerTransport) dispatchGetTrainingResults(req *http.Request) (*http.Response, error) {
	if p.srv.GetTrainingResults == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetTrainingResults not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predictions/(?P<predictionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getTrainingResults`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	predictionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predictionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetTrainingResults(req.Context(), resourceGroupNameParam, hubNameParam, predictionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PredictionTrainingResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PredictionsServerTransport) dispatchNewListByHubPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByHubPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHubPager not implemented")}
	}
	newListByHubPager := p.newListByHubPager.get(req)
	if newListByHubPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predictions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByHubPager(resourceGroupNameParam, hubNameParam, nil)
		newListByHubPager = &resp
		p.newListByHubPager.add(req, newListByHubPager)
		server.PagerResponderInjectNextLinks(newListByHubPager, req, func(page *armcustomerinsights.PredictionsClientListByHubResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHubPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByHubPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHubPager) {
		p.newListByHubPager.remove(req)
	}
	return resp, nil
}

func (p *PredictionsServerTransport) dispatchModelStatus(req *http.Request) (*http.Response, error) {
	if p.srv.ModelStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method ModelStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predictions/(?P<predictionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/modelStatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcustomerinsights.PredictionModelStatus](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	predictionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predictionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ModelStatus(req.Context(), resourceGroupNameParam, hubNameParam, predictionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PredictionsServerTransport
var predictionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
