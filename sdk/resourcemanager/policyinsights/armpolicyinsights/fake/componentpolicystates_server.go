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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// ComponentPolicyStatesServer is a fake server for instances of the armpolicyinsights.ComponentPolicyStatesClient type.
type ComponentPolicyStatesServer struct {
	// ListQueryResultsForPolicyDefinition is the fake for method ComponentPolicyStatesClient.ListQueryResultsForPolicyDefinition
	// HTTP status codes to indicate success: http.StatusOK
	ListQueryResultsForPolicyDefinition func(ctx context.Context, subscriptionID string, policyDefinitionName string, componentPolicyStatesResource armpolicyinsights.ComponentPolicyStatesResource, options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionOptions) (resp azfake.Responder[armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse], errResp azfake.ErrorResponder)

	// ListQueryResultsForResource is the fake for method ComponentPolicyStatesClient.ListQueryResultsForResource
	// HTTP status codes to indicate success: http.StatusOK
	ListQueryResultsForResource func(ctx context.Context, resourceID string, componentPolicyStatesResource armpolicyinsights.ComponentPolicyStatesResource, options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceOptions) (resp azfake.Responder[armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceResponse], errResp azfake.ErrorResponder)

	// ListQueryResultsForResourceGroup is the fake for method ComponentPolicyStatesClient.ListQueryResultsForResourceGroup
	// HTTP status codes to indicate success: http.StatusOK
	ListQueryResultsForResourceGroup func(ctx context.Context, subscriptionID string, resourceGroupName string, componentPolicyStatesResource armpolicyinsights.ComponentPolicyStatesResource, options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions) (resp azfake.Responder[armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse], errResp azfake.ErrorResponder)

	// ListQueryResultsForResourceGroupLevelPolicyAssignment is the fake for method ComponentPolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment
	// HTTP status codes to indicate success: http.StatusOK
	ListQueryResultsForResourceGroupLevelPolicyAssignment func(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, componentPolicyStatesResource armpolicyinsights.ComponentPolicyStatesResource, options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions) (resp azfake.Responder[armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse], errResp azfake.ErrorResponder)

	// ListQueryResultsForSubscription is the fake for method ComponentPolicyStatesClient.ListQueryResultsForSubscription
	// HTTP status codes to indicate success: http.StatusOK
	ListQueryResultsForSubscription func(ctx context.Context, subscriptionID string, componentPolicyStatesResource armpolicyinsights.ComponentPolicyStatesResource, options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionOptions) (resp azfake.Responder[armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse], errResp azfake.ErrorResponder)

	// ListQueryResultsForSubscriptionLevelPolicyAssignment is the fake for method ComponentPolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment
	// HTTP status codes to indicate success: http.StatusOK
	ListQueryResultsForSubscriptionLevelPolicyAssignment func(ctx context.Context, subscriptionID string, policyAssignmentName string, componentPolicyStatesResource armpolicyinsights.ComponentPolicyStatesResource, options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions) (resp azfake.Responder[armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse], errResp azfake.ErrorResponder)
}

// NewComponentPolicyStatesServerTransport creates a new instance of ComponentPolicyStatesServerTransport with the provided implementation.
// The returned ComponentPolicyStatesServerTransport instance is connected to an instance of armpolicyinsights.ComponentPolicyStatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewComponentPolicyStatesServerTransport(srv *ComponentPolicyStatesServer) *ComponentPolicyStatesServerTransport {
	return &ComponentPolicyStatesServerTransport{srv: srv}
}

// ComponentPolicyStatesServerTransport connects instances of armpolicyinsights.ComponentPolicyStatesClient to instances of ComponentPolicyStatesServer.
// Don't use this type directly, use NewComponentPolicyStatesServerTransport instead.
type ComponentPolicyStatesServerTransport struct {
	srv *ComponentPolicyStatesServer
}

// Do implements the policy.Transporter interface for ComponentPolicyStatesServerTransport.
func (c *ComponentPolicyStatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ComponentPolicyStatesClient.ListQueryResultsForPolicyDefinition":
		resp, err = c.dispatchListQueryResultsForPolicyDefinition(req)
	case "ComponentPolicyStatesClient.ListQueryResultsForResource":
		resp, err = c.dispatchListQueryResultsForResource(req)
	case "ComponentPolicyStatesClient.ListQueryResultsForResourceGroup":
		resp, err = c.dispatchListQueryResultsForResourceGroup(req)
	case "ComponentPolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment":
		resp, err = c.dispatchListQueryResultsForResourceGroupLevelPolicyAssignment(req)
	case "ComponentPolicyStatesClient.ListQueryResultsForSubscription":
		resp, err = c.dispatchListQueryResultsForSubscription(req)
	case "ComponentPolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment":
		resp, err = c.dispatchListQueryResultsForSubscriptionLevelPolicyAssignment(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ComponentPolicyStatesServerTransport) dispatchListQueryResultsForPolicyDefinition(req *http.Request) (*http.Response, error) {
	if c.srv.ListQueryResultsForPolicyDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQueryResultsForPolicyDefinition not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<authorizationNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PolicyInsights/componentPolicyStates/(?P<componentPolicyStatesResource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryResults`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	componentPolicyStatesResourceParam, err := parseWithCast(matches[regex.SubexpIndex("componentPolicyStatesResource")], func(v string) (armpolicyinsights.ComponentPolicyStatesResource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpolicyinsights.ComponentPolicyStatesResource(p), nil
	})
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
	orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(selectUnescaped)
	fromUnescaped, err := url.QueryUnescape(qp.Get("$from"))
	if err != nil {
		return nil, err
	}
	fromParam, err := parseOptional(fromUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	toUnescaped, err := url.QueryUnescape(qp.Get("$to"))
	if err != nil {
		return nil, err
	}
	toParam, err := parseOptional(toUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	applyUnescaped, err := url.QueryUnescape(qp.Get("$apply"))
	if err != nil {
		return nil, err
	}
	applyParam := getOptional(applyUnescaped)
	var options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionOptions
	if topParam != nil || orderByParam != nil || selectParam != nil || fromParam != nil || toParam != nil || filterParam != nil || applyParam != nil {
		options = &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionOptions{
			Top:     topParam,
			OrderBy: orderByParam,
			Select:  selectParam,
			From:    fromParam,
			To:      toParam,
			Filter:  filterParam,
			Apply:   applyParam,
		}
	}
	respr, errRespr := c.srv.ListQueryResultsForPolicyDefinition(req.Context(), subscriptionIDParam, policyDefinitionNameParam, componentPolicyStatesResourceParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentPolicyStatesQueryResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ComponentPolicyStatesServerTransport) dispatchListQueryResultsForResource(req *http.Request) (*http.Response, error) {
	if c.srv.ListQueryResultsForResource == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQueryResultsForResource not implemented")}
	}
	const regexStr = `/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PolicyInsights/componentPolicyStates/(?P<componentPolicyStatesResource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryResults`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceId")])
	if err != nil {
		return nil, err
	}
	componentPolicyStatesResourceParam, err := parseWithCast(matches[regex.SubexpIndex("componentPolicyStatesResource")], func(v string) (armpolicyinsights.ComponentPolicyStatesResource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpolicyinsights.ComponentPolicyStatesResource(p), nil
	})
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
	orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(selectUnescaped)
	fromUnescaped, err := url.QueryUnescape(qp.Get("$from"))
	if err != nil {
		return nil, err
	}
	fromParam, err := parseOptional(fromUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	toUnescaped, err := url.QueryUnescape(qp.Get("$to"))
	if err != nil {
		return nil, err
	}
	toParam, err := parseOptional(toUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	applyUnescaped, err := url.QueryUnescape(qp.Get("$apply"))
	if err != nil {
		return nil, err
	}
	applyParam := getOptional(applyUnescaped)
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceOptions
	if topParam != nil || orderByParam != nil || selectParam != nil || fromParam != nil || toParam != nil || filterParam != nil || applyParam != nil || expandParam != nil {
		options = &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceOptions{
			Top:     topParam,
			OrderBy: orderByParam,
			Select:  selectParam,
			From:    fromParam,
			To:      toParam,
			Filter:  filterParam,
			Apply:   applyParam,
			Expand:  expandParam,
		}
	}
	respr, errRespr := c.srv.ListQueryResultsForResource(req.Context(), resourceIDParam, componentPolicyStatesResourceParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentPolicyStatesQueryResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ComponentPolicyStatesServerTransport) dispatchListQueryResultsForResourceGroup(req *http.Request) (*http.Response, error) {
	if c.srv.ListQueryResultsForResourceGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQueryResultsForResourceGroup not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PolicyInsights/componentPolicyStates/(?P<componentPolicyStatesResource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryResults`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	componentPolicyStatesResourceParam, err := parseWithCast(matches[regex.SubexpIndex("componentPolicyStatesResource")], func(v string) (armpolicyinsights.ComponentPolicyStatesResource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpolicyinsights.ComponentPolicyStatesResource(p), nil
	})
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
	orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(selectUnescaped)
	fromUnescaped, err := url.QueryUnescape(qp.Get("$from"))
	if err != nil {
		return nil, err
	}
	fromParam, err := parseOptional(fromUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	toUnescaped, err := url.QueryUnescape(qp.Get("$to"))
	if err != nil {
		return nil, err
	}
	toParam, err := parseOptional(toUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	applyUnescaped, err := url.QueryUnescape(qp.Get("$apply"))
	if err != nil {
		return nil, err
	}
	applyParam := getOptional(applyUnescaped)
	var options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions
	if topParam != nil || orderByParam != nil || selectParam != nil || fromParam != nil || toParam != nil || filterParam != nil || applyParam != nil {
		options = &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions{
			Top:     topParam,
			OrderBy: orderByParam,
			Select:  selectParam,
			From:    fromParam,
			To:      toParam,
			Filter:  filterParam,
			Apply:   applyParam,
		}
	}
	respr, errRespr := c.srv.ListQueryResultsForResourceGroup(req.Context(), subscriptionIDParam, resourceGroupNameParam, componentPolicyStatesResourceParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentPolicyStatesQueryResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ComponentPolicyStatesServerTransport) dispatchListQueryResultsForResourceGroupLevelPolicyAssignment(req *http.Request) (*http.Response, error) {
	if c.srv.ListQueryResultsForResourceGroupLevelPolicyAssignment == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQueryResultsForResourceGroupLevelPolicyAssignment not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<authorizationNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policyAssignments/(?P<policyAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PolicyInsights/componentPolicyStates/(?P<componentPolicyStatesResource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryResults`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyAssignmentName")])
	if err != nil {
		return nil, err
	}
	componentPolicyStatesResourceParam, err := parseWithCast(matches[regex.SubexpIndex("componentPolicyStatesResource")], func(v string) (armpolicyinsights.ComponentPolicyStatesResource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpolicyinsights.ComponentPolicyStatesResource(p), nil
	})
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
	orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(selectUnescaped)
	fromUnescaped, err := url.QueryUnescape(qp.Get("$from"))
	if err != nil {
		return nil, err
	}
	fromParam, err := parseOptional(fromUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	toUnescaped, err := url.QueryUnescape(qp.Get("$to"))
	if err != nil {
		return nil, err
	}
	toParam, err := parseOptional(toUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	applyUnescaped, err := url.QueryUnescape(qp.Get("$apply"))
	if err != nil {
		return nil, err
	}
	applyParam := getOptional(applyUnescaped)
	var options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions
	if topParam != nil || orderByParam != nil || selectParam != nil || fromParam != nil || toParam != nil || filterParam != nil || applyParam != nil {
		options = &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions{
			Top:     topParam,
			OrderBy: orderByParam,
			Select:  selectParam,
			From:    fromParam,
			To:      toParam,
			Filter:  filterParam,
			Apply:   applyParam,
		}
	}
	respr, errRespr := c.srv.ListQueryResultsForResourceGroupLevelPolicyAssignment(req.Context(), subscriptionIDParam, resourceGroupNameParam, policyAssignmentNameParam, componentPolicyStatesResourceParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentPolicyStatesQueryResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ComponentPolicyStatesServerTransport) dispatchListQueryResultsForSubscription(req *http.Request) (*http.Response, error) {
	if c.srv.ListQueryResultsForSubscription == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQueryResultsForSubscription not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PolicyInsights/componentPolicyStates/(?P<componentPolicyStatesResource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryResults`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	componentPolicyStatesResourceParam, err := parseWithCast(matches[regex.SubexpIndex("componentPolicyStatesResource")], func(v string) (armpolicyinsights.ComponentPolicyStatesResource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpolicyinsights.ComponentPolicyStatesResource(p), nil
	})
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
	orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(selectUnescaped)
	fromUnescaped, err := url.QueryUnescape(qp.Get("$from"))
	if err != nil {
		return nil, err
	}
	fromParam, err := parseOptional(fromUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	toUnescaped, err := url.QueryUnescape(qp.Get("$to"))
	if err != nil {
		return nil, err
	}
	toParam, err := parseOptional(toUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	applyUnescaped, err := url.QueryUnescape(qp.Get("$apply"))
	if err != nil {
		return nil, err
	}
	applyParam := getOptional(applyUnescaped)
	var options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionOptions
	if topParam != nil || orderByParam != nil || selectParam != nil || fromParam != nil || toParam != nil || filterParam != nil || applyParam != nil {
		options = &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionOptions{
			Top:     topParam,
			OrderBy: orderByParam,
			Select:  selectParam,
			From:    fromParam,
			To:      toParam,
			Filter:  filterParam,
			Apply:   applyParam,
		}
	}
	respr, errRespr := c.srv.ListQueryResultsForSubscription(req.Context(), subscriptionIDParam, componentPolicyStatesResourceParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentPolicyStatesQueryResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ComponentPolicyStatesServerTransport) dispatchListQueryResultsForSubscriptionLevelPolicyAssignment(req *http.Request) (*http.Response, error) {
	if c.srv.ListQueryResultsForSubscriptionLevelPolicyAssignment == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQueryResultsForSubscriptionLevelPolicyAssignment not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<authorizationNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policyAssignments/(?P<policyAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PolicyInsights/componentPolicyStates/(?P<componentPolicyStatesResource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryResults`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	policyAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyAssignmentName")])
	if err != nil {
		return nil, err
	}
	componentPolicyStatesResourceParam, err := parseWithCast(matches[regex.SubexpIndex("componentPolicyStatesResource")], func(v string) (armpolicyinsights.ComponentPolicyStatesResource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpolicyinsights.ComponentPolicyStatesResource(p), nil
	})
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
	orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(selectUnescaped)
	fromUnescaped, err := url.QueryUnescape(qp.Get("$from"))
	if err != nil {
		return nil, err
	}
	fromParam, err := parseOptional(fromUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	toUnescaped, err := url.QueryUnescape(qp.Get("$to"))
	if err != nil {
		return nil, err
	}
	toParam, err := parseOptional(toUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	applyUnescaped, err := url.QueryUnescape(qp.Get("$apply"))
	if err != nil {
		return nil, err
	}
	applyParam := getOptional(applyUnescaped)
	var options *armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions
	if topParam != nil || orderByParam != nil || selectParam != nil || fromParam != nil || toParam != nil || filterParam != nil || applyParam != nil {
		options = &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions{
			Top:     topParam,
			OrderBy: orderByParam,
			Select:  selectParam,
			From:    fromParam,
			To:      toParam,
			Filter:  filterParam,
			Apply:   applyParam,
		}
	}
	respr, errRespr := c.srv.ListQueryResultsForSubscriptionLevelPolicyAssignment(req.Context(), subscriptionIDParam, policyAssignmentNameParam, componentPolicyStatesResourceParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentPolicyStatesQueryResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}