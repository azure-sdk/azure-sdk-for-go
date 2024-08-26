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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
)

// RegistryDataReferencesServer is a fake server for instances of the armmachinelearning.RegistryDataReferencesClient type.
type RegistryDataReferencesServer struct {
	// GetBlobReferenceSAS is the fake for method RegistryDataReferencesClient.GetBlobReferenceSAS
	// HTTP status codes to indicate success: http.StatusOK
	GetBlobReferenceSAS func(ctx context.Context, resourceGroupName string, registryName string, name string, version string, body armmachinelearning.GetBlobReferenceSASRequestDto, options *armmachinelearning.RegistryDataReferencesClientGetBlobReferenceSASOptions) (resp azfake.Responder[armmachinelearning.RegistryDataReferencesClientGetBlobReferenceSASResponse], errResp azfake.ErrorResponder)
}

// NewRegistryDataReferencesServerTransport creates a new instance of RegistryDataReferencesServerTransport with the provided implementation.
// The returned RegistryDataReferencesServerTransport instance is connected to an instance of armmachinelearning.RegistryDataReferencesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegistryDataReferencesServerTransport(srv *RegistryDataReferencesServer) *RegistryDataReferencesServerTransport {
	return &RegistryDataReferencesServerTransport{srv: srv}
}

// RegistryDataReferencesServerTransport connects instances of armmachinelearning.RegistryDataReferencesClient to instances of RegistryDataReferencesServer.
// Don't use this type directly, use NewRegistryDataReferencesServerTransport instead.
type RegistryDataReferencesServerTransport struct {
	srv *RegistryDataReferencesServer
}

// Do implements the policy.Transporter interface for RegistryDataReferencesServerTransport.
func (r *RegistryDataReferencesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RegistryDataReferencesClient.GetBlobReferenceSAS":
		resp, err = r.dispatchGetBlobReferenceSAS(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegistryDataReferencesServerTransport) dispatchGetBlobReferenceSAS(req *http.Request) (*http.Response, error) {
	if r.srv.GetBlobReferenceSAS == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBlobReferenceSAS not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datareferences/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.GetBlobReferenceSASRequestDto](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetBlobReferenceSAS(req.Context(), resourceGroupNameParam, registryNameParam, nameParam, versionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GetBlobReferenceSASResponseDto, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}