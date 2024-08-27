// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armcontoso.ClientFactory type.
type ServerFactory struct {
	// EmployeesServer contains the fakes for client EmployeesClient
	EmployeesServer EmployeesServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armcontoso.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armcontoso.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                *ServerFactory
	trMu               sync.Mutex
	trEmployeesServer  *EmployeesServerTransport
	trOperationsServer *OperationsServerTransport
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
	case "EmployeesClient":
		initServer(s, &s.trEmployeesServer, func() *EmployeesServerTransport { return NewEmployeesServerTransport(&s.srv.EmployeesServer) })
		resp, err = s.trEmployeesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
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
