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

// ServerFactory is a fake server for instances of the armdevopsinfrastructure.ClientFactory type.
type ServerFactory struct {
	ImageVersionsServer      ImageVersionsServer
	OperationsServer         OperationsServer
	PoolsServer              PoolsServer
	ResourceDetailsServer    ResourceDetailsServer
	SKUServer                SKUServer
	SubscriptionUsagesServer SubscriptionUsagesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdevopsinfrastructure.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdevopsinfrastructure.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                        *ServerFactory
	trMu                       sync.Mutex
	trImageVersionsServer      *ImageVersionsServerTransport
	trOperationsServer         *OperationsServerTransport
	trPoolsServer              *PoolsServerTransport
	trResourceDetailsServer    *ResourceDetailsServerTransport
	trSKUServer                *SKUServerTransport
	trSubscriptionUsagesServer *SubscriptionUsagesServerTransport
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
	case "ImageVersionsClient":
		initServer(s, &s.trImageVersionsServer, func() *ImageVersionsServerTransport {
			return NewImageVersionsServerTransport(&s.srv.ImageVersionsServer)
		})
		resp, err = s.trImageVersionsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PoolsClient":
		initServer(s, &s.trPoolsServer, func() *PoolsServerTransport { return NewPoolsServerTransport(&s.srv.PoolsServer) })
		resp, err = s.trPoolsServer.Do(req)
	case "ResourceDetailsClient":
		initServer(s, &s.trResourceDetailsServer, func() *ResourceDetailsServerTransport {
			return NewResourceDetailsServerTransport(&s.srv.ResourceDetailsServer)
		})
		resp, err = s.trResourceDetailsServer.Do(req)
	case "SKUClient":
		initServer(s, &s.trSKUServer, func() *SKUServerTransport { return NewSKUServerTransport(&s.srv.SKUServer) })
		resp, err = s.trSKUServer.Do(req)
	case "SubscriptionUsagesClient":
		initServer(s, &s.trSubscriptionUsagesServer, func() *SubscriptionUsagesServerTransport {
			return NewSubscriptionUsagesServerTransport(&s.srv.SubscriptionUsagesServer)
		})
		resp, err = s.trSubscriptionUsagesServer.Do(req)
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
