//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armdevhub.ClientFactory type.
type ServerFactory struct {
	DeveloperHubServiceServer DeveloperHubServiceServer
	IacProfilesServer         IacProfilesServer
	OperationsServer          OperationsServer
	WorkflowServer            WorkflowServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdevhub.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdevhub.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                         *ServerFactory
	trMu                        sync.Mutex
	trDeveloperHubServiceServer *DeveloperHubServiceServerTransport
	trIacProfilesServer         *IacProfilesServerTransport
	trOperationsServer          *OperationsServerTransport
	trWorkflowServer            *WorkflowServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "DeveloperHubServiceClient":
		initServer(s, &s.trDeveloperHubServiceServer, func() *DeveloperHubServiceServerTransport {
			return NewDeveloperHubServiceServerTransport(&s.srv.DeveloperHubServiceServer)
		})
		resp, err = s.trDeveloperHubServiceServer.Do(req)
	case "IacProfilesClient":
		initServer(s, &s.trIacProfilesServer, func() *IacProfilesServerTransport { return NewIacProfilesServerTransport(&s.srv.IacProfilesServer) })
		resp, err = s.trIacProfilesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "WorkflowClient":
		initServer(s, &s.trWorkflowServer, func() *WorkflowServerTransport { return NewWorkflowServerTransport(&s.srv.WorkflowServer) })
		resp, err = s.trWorkflowServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
