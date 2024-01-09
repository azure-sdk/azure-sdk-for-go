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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
	"net/http"
	"net/url"
	"regexp"
)

// EntityRelationsServer is a fake server for instances of the armsecurityinsights.EntityRelationsClient type.
type EntityRelationsServer struct {
	// GetRelation is the fake for method EntityRelationsClient.GetRelation
	// HTTP status codes to indicate success: http.StatusOK
	GetRelation func(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, relationName string, options *armsecurityinsights.EntityRelationsClientGetRelationOptions) (resp azfake.Responder[armsecurityinsights.EntityRelationsClientGetRelationResponse], errResp azfake.ErrorResponder)
}

// NewEntityRelationsServerTransport creates a new instance of EntityRelationsServerTransport with the provided implementation.
// The returned EntityRelationsServerTransport instance is connected to an instance of armsecurityinsights.EntityRelationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEntityRelationsServerTransport(srv *EntityRelationsServer) *EntityRelationsServerTransport {
	return &EntityRelationsServerTransport{srv: srv}
}

// EntityRelationsServerTransport connects instances of armsecurityinsights.EntityRelationsClient to instances of EntityRelationsServer.
// Don't use this type directly, use NewEntityRelationsServerTransport instead.
type EntityRelationsServerTransport struct {
	srv *EntityRelationsServer
}

// Do implements the policy.Transporter interface for EntityRelationsServerTransport.
func (e *EntityRelationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EntityRelationsClient.GetRelation":
		resp, err = e.dispatchGetRelation(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EntityRelationsServerTransport) dispatchGetRelation(req *http.Request) (*http.Response, error) {
	if e.srv.GetRelation == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRelation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/entities/(?P<entityId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/relations/(?P<relationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	entityIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("entityId")])
	if err != nil {
		return nil, err
	}
	relationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("relationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetRelation(req.Context(), resourceGroupNameParam, workspaceNameParam, entityIDParam, relationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Relation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
