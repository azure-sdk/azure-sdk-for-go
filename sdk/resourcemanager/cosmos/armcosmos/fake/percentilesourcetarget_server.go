// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
	"net/http"
	"net/url"
	"regexp"
)

// PercentileSourceTargetServer is a fake server for instances of the armcosmos.PercentileSourceTargetClient type.
type PercentileSourceTargetServer struct {
	// NewListMetricsPager is the fake for method PercentileSourceTargetClient.NewListMetricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricsPager func(resourceGroupName string, accountName string, sourceRegion string, targetRegion string, filter string, options *armcosmos.PercentileSourceTargetClientListMetricsOptions) (resp azfake.PagerResponder[armcosmos.PercentileSourceTargetClientListMetricsResponse])
}

// NewPercentileSourceTargetServerTransport creates a new instance of PercentileSourceTargetServerTransport with the provided implementation.
// The returned PercentileSourceTargetServerTransport instance is connected to an instance of armcosmos.PercentileSourceTargetClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPercentileSourceTargetServerTransport(srv *PercentileSourceTargetServer) *PercentileSourceTargetServerTransport {
	return &PercentileSourceTargetServerTransport{
		srv:                 srv,
		newListMetricsPager: newTracker[azfake.PagerResponder[armcosmos.PercentileSourceTargetClientListMetricsResponse]](),
	}
}

// PercentileSourceTargetServerTransport connects instances of armcosmos.PercentileSourceTargetClient to instances of PercentileSourceTargetServer.
// Don't use this type directly, use NewPercentileSourceTargetServerTransport instead.
type PercentileSourceTargetServerTransport struct {
	srv                 *PercentileSourceTargetServer
	newListMetricsPager *tracker[azfake.PagerResponder[armcosmos.PercentileSourceTargetClientListMetricsResponse]]
}

// Do implements the policy.Transporter interface for PercentileSourceTargetServerTransport.
func (p *PercentileSourceTargetServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PercentileSourceTargetServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if percentileSourceTargetServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = percentileSourceTargetServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PercentileSourceTargetClient.NewListMetricsPager":
				res.resp, res.err = p.dispatchNewListMetricsPager(req)
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

func (p *PercentileSourceTargetServerTransport) dispatchNewListMetricsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListMetricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricsPager not implemented")}
	}
	newListMetricsPager := p.newListMetricsPager.get(req)
	if newListMetricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sourceRegion/(?P<sourceRegion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetRegion/(?P<targetRegion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/percentile/metrics`
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
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		sourceRegionParam, err := url.PathUnescape(matches[regex.SubexpIndex("sourceRegion")])
		if err != nil {
			return nil, err
		}
		targetRegionParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetRegion")])
		if err != nil {
			return nil, err
		}
		filterParam, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListMetricsPager(resourceGroupNameParam, accountNameParam, sourceRegionParam, targetRegionParam, filterParam, nil)
		newListMetricsPager = &resp
		p.newListMetricsPager.add(req, newListMetricsPager)
	}
	resp, err := server.PagerResponderNext(newListMetricsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListMetricsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricsPager) {
		p.newListMetricsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PercentileSourceTargetServerTransport
var percentileSourceTargetServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
