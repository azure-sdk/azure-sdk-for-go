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

// ServerFactory is a fake server for instances of the armdurabletask.ClientFactory type.
type ServerFactory struct {
	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// SchedulersServer contains the fakes for client SchedulersClient
	SchedulersServer SchedulersServer

	// TaskHubsServer contains the fakes for client TaskHubsClient
	TaskHubsServer TaskHubsServer

}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdurabletask.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdurabletask.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv *ServerFactory
	trMu sync.Mutex
	trOperationsServer *OperationsServerTransport
	trSchedulersServer *SchedulersServerTransport
	trTaskHubsServer *TaskHubsServerTransport
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
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SchedulersClient":
		initServer(s, &s.trSchedulersServer, func() *SchedulersServerTransport { return NewSchedulersServerTransport(&s.srv.SchedulersServer) })
		resp, err = s.trSchedulersServer.Do(req)
	case "TaskHubsClient":
		initServer(s, &s.trTaskHubsServer, func() *TaskHubsServerTransport { return NewTaskHubsServerTransport(&s.srv.TaskHubsServer) })
		resp, err = s.trTaskHubsServer.Do(req)
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
