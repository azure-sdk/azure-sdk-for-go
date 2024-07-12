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

// ServerFactory is a fake server for instances of the armguestconfiguration.ClientFactory type.
type ServerFactory struct {
	AssignmentReportsServer                        AssignmentReportsServer
	AssignmentReportsVMSSServer                    AssignmentReportsVMSSServer
	AssignmentsServer                              AssignmentsServer
	AssignmentsVMSSServer                          AssignmentsVMSSServer
	ConnectedVMwarevSphereAssignmentsServer        ConnectedVMwarevSphereAssignmentsServer
	ConnectedVMwarevSphereAssignmentsReportsServer ConnectedVMwarevSphereAssignmentsReportsServer
	HCRPAssignmentReportsServer                    HCRPAssignmentReportsServer
	HCRPAssignmentsServer                          HCRPAssignmentsServer
	OperationsServer                               OperationsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armguestconfiguration.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armguestconfiguration.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                              *ServerFactory
	trMu                                             sync.Mutex
	trAssignmentReportsServer                        *AssignmentReportsServerTransport
	trAssignmentReportsVMSSServer                    *AssignmentReportsVMSSServerTransport
	trAssignmentsServer                              *AssignmentsServerTransport
	trAssignmentsVMSSServer                          *AssignmentsVMSSServerTransport
	trConnectedVMwarevSphereAssignmentsServer        *ConnectedVMwarevSphereAssignmentsServerTransport
	trConnectedVMwarevSphereAssignmentsReportsServer *ConnectedVMwarevSphereAssignmentsReportsServerTransport
	trHCRPAssignmentReportsServer                    *HCRPAssignmentReportsServerTransport
	trHCRPAssignmentsServer                          *HCRPAssignmentsServerTransport
	trOperationsServer                               *OperationsServerTransport
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
	case "AssignmentReportsClient":
		initServer(s, &s.trAssignmentReportsServer, func() *AssignmentReportsServerTransport {
			return NewAssignmentReportsServerTransport(&s.srv.AssignmentReportsServer)
		})
		resp, err = s.trAssignmentReportsServer.Do(req)
	case "AssignmentReportsVMSSClient":
		initServer(s, &s.trAssignmentReportsVMSSServer, func() *AssignmentReportsVMSSServerTransport {
			return NewAssignmentReportsVMSSServerTransport(&s.srv.AssignmentReportsVMSSServer)
		})
		resp, err = s.trAssignmentReportsVMSSServer.Do(req)
	case "AssignmentsClient":
		initServer(s, &s.trAssignmentsServer, func() *AssignmentsServerTransport { return NewAssignmentsServerTransport(&s.srv.AssignmentsServer) })
		resp, err = s.trAssignmentsServer.Do(req)
	case "AssignmentsVMSSClient":
		initServer(s, &s.trAssignmentsVMSSServer, func() *AssignmentsVMSSServerTransport {
			return NewAssignmentsVMSSServerTransport(&s.srv.AssignmentsVMSSServer)
		})
		resp, err = s.trAssignmentsVMSSServer.Do(req)
	case "ConnectedVMwarevSphereAssignmentsClient":
		initServer(s, &s.trConnectedVMwarevSphereAssignmentsServer, func() *ConnectedVMwarevSphereAssignmentsServerTransport {
			return NewConnectedVMwarevSphereAssignmentsServerTransport(&s.srv.ConnectedVMwarevSphereAssignmentsServer)
		})
		resp, err = s.trConnectedVMwarevSphereAssignmentsServer.Do(req)
	case "ConnectedVMwarevSphereAssignmentsReportsClient":
		initServer(s, &s.trConnectedVMwarevSphereAssignmentsReportsServer, func() *ConnectedVMwarevSphereAssignmentsReportsServerTransport {
			return NewConnectedVMwarevSphereAssignmentsReportsServerTransport(&s.srv.ConnectedVMwarevSphereAssignmentsReportsServer)
		})
		resp, err = s.trConnectedVMwarevSphereAssignmentsReportsServer.Do(req)
	case "HCRPAssignmentReportsClient":
		initServer(s, &s.trHCRPAssignmentReportsServer, func() *HCRPAssignmentReportsServerTransport {
			return NewHCRPAssignmentReportsServerTransport(&s.srv.HCRPAssignmentReportsServer)
		})
		resp, err = s.trHCRPAssignmentReportsServer.Do(req)
	case "HCRPAssignmentsClient":
		initServer(s, &s.trHCRPAssignmentsServer, func() *HCRPAssignmentsServerTransport {
			return NewHCRPAssignmentsServerTransport(&s.srv.HCRPAssignmentsServer)
		})
		resp, err = s.trHCRPAssignmentsServer.Do(req)
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
