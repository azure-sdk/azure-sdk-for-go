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

// ServerFactory is a fake server for instances of the armstreamanalytics.ClientFactory type.
type ServerFactory struct {
	// ClustersServer contains the fakes for client ClustersClient
	ClustersServer ClustersServer

	// FunctionsServer contains the fakes for client FunctionsClient
	FunctionsServer FunctionsServer

	// InputsServer contains the fakes for client InputsClient
	InputsServer InputsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// OutputsServer contains the fakes for client OutputsClient
	OutputsServer OutputsServer

	// PrivateEndpointsServer contains the fakes for client PrivateEndpointsClient
	PrivateEndpointsServer PrivateEndpointsServer

	// StreamingJobsServer contains the fakes for client StreamingJobsClient
	StreamingJobsServer StreamingJobsServer

	// SubscriptionsServer contains the fakes for client SubscriptionsClient
	SubscriptionsServer SubscriptionsServer

	// TransformationsServer contains the fakes for client TransformationsClient
	TransformationsServer TransformationsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armstreamanalytics.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armstreamanalytics.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                      *ServerFactory
	trMu                     sync.Mutex
	trClustersServer         *ClustersServerTransport
	trFunctionsServer        *FunctionsServerTransport
	trInputsServer           *InputsServerTransport
	trOperationsServer       *OperationsServerTransport
	trOutputsServer          *OutputsServerTransport
	trPrivateEndpointsServer *PrivateEndpointsServerTransport
	trStreamingJobsServer    *StreamingJobsServerTransport
	trSubscriptionsServer    *SubscriptionsServerTransport
	trTransformationsServer  *TransformationsServerTransport
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
	case "ClustersClient":
		initServer(s, &s.trClustersServer, func() *ClustersServerTransport { return NewClustersServerTransport(&s.srv.ClustersServer) })
		resp, err = s.trClustersServer.Do(req)
	case "FunctionsClient":
		initServer(s, &s.trFunctionsServer, func() *FunctionsServerTransport { return NewFunctionsServerTransport(&s.srv.FunctionsServer) })
		resp, err = s.trFunctionsServer.Do(req)
	case "InputsClient":
		initServer(s, &s.trInputsServer, func() *InputsServerTransport { return NewInputsServerTransport(&s.srv.InputsServer) })
		resp, err = s.trInputsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "OutputsClient":
		initServer(s, &s.trOutputsServer, func() *OutputsServerTransport { return NewOutputsServerTransport(&s.srv.OutputsServer) })
		resp, err = s.trOutputsServer.Do(req)
	case "PrivateEndpointsClient":
		initServer(s, &s.trPrivateEndpointsServer, func() *PrivateEndpointsServerTransport {
			return NewPrivateEndpointsServerTransport(&s.srv.PrivateEndpointsServer)
		})
		resp, err = s.trPrivateEndpointsServer.Do(req)
	case "StreamingJobsClient":
		initServer(s, &s.trStreamingJobsServer, func() *StreamingJobsServerTransport {
			return NewStreamingJobsServerTransport(&s.srv.StreamingJobsServer)
		})
		resp, err = s.trStreamingJobsServer.Do(req)
	case "SubscriptionsClient":
		initServer(s, &s.trSubscriptionsServer, func() *SubscriptionsServerTransport {
			return NewSubscriptionsServerTransport(&s.srv.SubscriptionsServer)
		})
		resp, err = s.trSubscriptionsServer.Do(req)
	case "TransformationsClient":
		initServer(s, &s.trTransformationsServer, func() *TransformationsServerTransport {
			return NewTransformationsServerTransport(&s.srv.TransformationsServer)
		})
		resp, err = s.trTransformationsServer.Do(req)
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
