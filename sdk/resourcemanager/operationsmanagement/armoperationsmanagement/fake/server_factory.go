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

// ServerFactory is a fake server for instances of the armoperationsmanagement.ClientFactory type.
type ServerFactory struct {
	// ManagementAssociationsServer contains the fakes for client ManagementAssociationsClient
	ManagementAssociationsServer ManagementAssociationsServer

	// ManagementConfigurationsServer contains the fakes for client ManagementConfigurationsClient
	ManagementConfigurationsServer ManagementConfigurationsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// SolutionsServer contains the fakes for client SolutionsClient
	SolutionsServer SolutionsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armoperationsmanagement.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armoperationsmanagement.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                              *ServerFactory
	trMu                             sync.Mutex
	trManagementAssociationsServer   *ManagementAssociationsServerTransport
	trManagementConfigurationsServer *ManagementConfigurationsServerTransport
	trOperationsServer               *OperationsServerTransport
	trSolutionsServer                *SolutionsServerTransport
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
	case "ManagementAssociationsClient":
		initServer(s, &s.trManagementAssociationsServer, func() *ManagementAssociationsServerTransport {
			return NewManagementAssociationsServerTransport(&s.srv.ManagementAssociationsServer)
		})
		resp, err = s.trManagementAssociationsServer.Do(req)
	case "ManagementConfigurationsClient":
		initServer(s, &s.trManagementConfigurationsServer, func() *ManagementConfigurationsServerTransport {
			return NewManagementConfigurationsServerTransport(&s.srv.ManagementConfigurationsServer)
		})
		resp, err = s.trManagementConfigurationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SolutionsClient":
		initServer(s, &s.trSolutionsServer, func() *SolutionsServerTransport { return NewSolutionsServerTransport(&s.srv.SolutionsServer) })
		resp, err = s.trSolutionsServer.Do(req)
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
