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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AlertRuleServer is a fake server for instances of the armsecurityinsights.AlertRuleClient type.
type AlertRuleServer struct {
	// BeginTriggerRuleRun is the fake for method AlertRuleClient.BeginTriggerRuleRun
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginTriggerRuleRun func(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, analyticsRuleRunTriggerParameter armsecurityinsights.AnalyticsRuleRunTrigger, options *armsecurityinsights.AlertRuleClientBeginTriggerRuleRunOptions) (resp azfake.PollerResponder[armsecurityinsights.AlertRuleClientTriggerRuleRunResponse], errResp azfake.ErrorResponder)
}

// NewAlertRuleServerTransport creates a new instance of AlertRuleServerTransport with the provided implementation.
// The returned AlertRuleServerTransport instance is connected to an instance of armsecurityinsights.AlertRuleClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAlertRuleServerTransport(srv *AlertRuleServer) *AlertRuleServerTransport {
	return &AlertRuleServerTransport{
		srv:                 srv,
		beginTriggerRuleRun: newTracker[azfake.PollerResponder[armsecurityinsights.AlertRuleClientTriggerRuleRunResponse]](),
	}
}

// AlertRuleServerTransport connects instances of armsecurityinsights.AlertRuleClient to instances of AlertRuleServer.
// Don't use this type directly, use NewAlertRuleServerTransport instead.
type AlertRuleServerTransport struct {
	srv                 *AlertRuleServer
	beginTriggerRuleRun *tracker[azfake.PollerResponder[armsecurityinsights.AlertRuleClientTriggerRuleRunResponse]]
}

// Do implements the policy.Transporter interface for AlertRuleServerTransport.
func (a *AlertRuleServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AlertRuleClient.BeginTriggerRuleRun":
		resp, err = a.dispatchBeginTriggerRuleRun(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AlertRuleServerTransport) dispatchBeginTriggerRuleRun(req *http.Request) (*http.Response, error) {
	if a.srv.BeginTriggerRuleRun == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTriggerRuleRun not implemented")}
	}
	beginTriggerRuleRun := a.beginTriggerRuleRun.get(req)
	if beginTriggerRuleRun == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/alertRules/(?P<ruleId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggerRuleRun`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.AnalyticsRuleRunTrigger](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		ruleIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginTriggerRuleRun(req.Context(), resourceGroupNameParam, workspaceNameParam, ruleIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTriggerRuleRun = &respr
		a.beginTriggerRuleRun.add(req, beginTriggerRuleRun)
	}

	resp, err := server.PollerResponderNext(beginTriggerRuleRun, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		a.beginTriggerRuleRun.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTriggerRuleRun) {
		a.beginTriggerRuleRun.remove(req)
	}

	return resp, nil
}
