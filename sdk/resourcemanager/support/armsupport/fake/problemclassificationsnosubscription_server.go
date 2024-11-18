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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"net/http"
	"net/url"
	"regexp"
)

// ProblemClassificationsNoSubscriptionServer is a fake server for instances of the armsupport.ProblemClassificationsNoSubscriptionClient type.
type ProblemClassificationsNoSubscriptionServer struct {
	// ClassifyProblems is the fake for method ProblemClassificationsNoSubscriptionClient.ClassifyProblems
	// HTTP status codes to indicate success: http.StatusOK
	ClassifyProblems func(ctx context.Context, problemServiceName string, problemClassificationsClassificationInput armsupport.ProblemClassificationsClassificationInput, options *armsupport.ProblemClassificationsNoSubscriptionClientClassifyProblemsOptions) (resp azfake.Responder[armsupport.ProblemClassificationsNoSubscriptionClientClassifyProblemsResponse], errResp azfake.ErrorResponder)
}

// NewProblemClassificationsNoSubscriptionServerTransport creates a new instance of ProblemClassificationsNoSubscriptionServerTransport with the provided implementation.
// The returned ProblemClassificationsNoSubscriptionServerTransport instance is connected to an instance of armsupport.ProblemClassificationsNoSubscriptionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProblemClassificationsNoSubscriptionServerTransport(srv *ProblemClassificationsNoSubscriptionServer) *ProblemClassificationsNoSubscriptionServerTransport {
	return &ProblemClassificationsNoSubscriptionServerTransport{srv: srv}
}

// ProblemClassificationsNoSubscriptionServerTransport connects instances of armsupport.ProblemClassificationsNoSubscriptionClient to instances of ProblemClassificationsNoSubscriptionServer.
// Don't use this type directly, use NewProblemClassificationsNoSubscriptionServerTransport instead.
type ProblemClassificationsNoSubscriptionServerTransport struct {
	srv *ProblemClassificationsNoSubscriptionServer
}

// Do implements the policy.Transporter interface for ProblemClassificationsNoSubscriptionServerTransport.
func (p *ProblemClassificationsNoSubscriptionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProblemClassificationsNoSubscriptionClient.ClassifyProblems":
		resp, err = p.dispatchClassifyProblems(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProblemClassificationsNoSubscriptionServerTransport) dispatchClassifyProblems(req *http.Request) (*http.Response, error) {
	if p.srv.ClassifyProblems == nil {
		return nil, &nonRetriableError{errors.New("fake for method ClassifyProblems not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Support/services/(?P<problemServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/classifyProblems`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsupport.ProblemClassificationsClassificationInput](req)
	if err != nil {
		return nil, err
	}
	problemServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("problemServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ClassifyProblems(req.Context(), problemServiceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProblemClassificationsClassificationOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}