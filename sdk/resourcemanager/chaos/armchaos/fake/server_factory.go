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

// ServerFactory is a fake server for instances of the armchaos.ClientFactory type.
type ServerFactory struct {
	// CapabilitiesServer contains the fakes for client CapabilitiesClient
	CapabilitiesServer CapabilitiesServer

	// CapabilityTypesServer contains the fakes for client CapabilityTypesClient
	CapabilityTypesServer CapabilityTypesServer

	// ExperimentsServer contains the fakes for client ExperimentsClient
	ExperimentsServer ExperimentsServer

	// OperationStatusesServer contains the fakes for client OperationStatusesClient
	OperationStatusesServer OperationStatusesServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// TargetTypesServer contains the fakes for client TargetTypesClient
	TargetTypesServer TargetTypesServer

	// TargetsServer contains the fakes for client TargetsClient
	TargetsServer TargetsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armchaos.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armchaos.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                       *ServerFactory
	trMu                      sync.Mutex
	trCapabilitiesServer      *CapabilitiesServerTransport
	trCapabilityTypesServer   *CapabilityTypesServerTransport
	trExperimentsServer       *ExperimentsServerTransport
	trOperationStatusesServer *OperationStatusesServerTransport
	trOperationsServer        *OperationsServerTransport
	trTargetTypesServer       *TargetTypesServerTransport
	trTargetsServer           *TargetsServerTransport
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
	case "CapabilitiesClient":
		initServer(s, &s.trCapabilitiesServer, func() *CapabilitiesServerTransport { return NewCapabilitiesServerTransport(&s.srv.CapabilitiesServer) })
		resp, err = s.trCapabilitiesServer.Do(req)
	case "CapabilityTypesClient":
		initServer(s, &s.trCapabilityTypesServer, func() *CapabilityTypesServerTransport {
			return NewCapabilityTypesServerTransport(&s.srv.CapabilityTypesServer)
		})
		resp, err = s.trCapabilityTypesServer.Do(req)
	case "ExperimentsClient":
		initServer(s, &s.trExperimentsServer, func() *ExperimentsServerTransport { return NewExperimentsServerTransport(&s.srv.ExperimentsServer) })
		resp, err = s.trExperimentsServer.Do(req)
	case "OperationStatusesClient":
		initServer(s, &s.trOperationStatusesServer, func() *OperationStatusesServerTransport {
			return NewOperationStatusesServerTransport(&s.srv.OperationStatusesServer)
		})
		resp, err = s.trOperationStatusesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "TargetTypesClient":
		initServer(s, &s.trTargetTypesServer, func() *TargetTypesServerTransport { return NewTargetTypesServerTransport(&s.srv.TargetTypesServer) })
		resp, err = s.trTargetTypesServer.Do(req)
	case "TargetsClient":
		initServer(s, &s.trTargetsServer, func() *TargetsServerTransport { return NewTargetsServerTransport(&s.srv.TargetsServer) })
		resp, err = s.trTargetsServer.Do(req)
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
