// Copyright 2019 Lightbend Inc.
// Copyright 2022 The eigr.io Authors.
// Use of this source code is governed by the Apache License
// that can be found in the LICENSE file.

package functions

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/eigr/functions-go-sdk/functions/action"
	"github.com/eigr/functions-go-sdk/functions/crdt"
	"github.com/eigr/functions-go-sdk/functions/discovery"
	"github.com/eigr/functions-go-sdk/functions/entity"
	"github.com/eigr/functions-go-sdk/functions/eventsourced"
	"github.com/eigr/functions-go-sdk/functions/protocol"
	"github.com/eigr/functions-go-sdk/functions/value"
	"google.golang.org/grpc"
)

// Functions is an instance of a Eigr Functions User Function.
type Functions struct {
	grpcServer            *grpc.Server
	entityDiscoveryServer *discovery.EntityDiscoveryServer
	eventSourcedServer    *eventsourced.Server
	crdtServer            *crdt.Server
	actionServer          *action.Server
	valueServer           *value.Server
}

// New returns a new Functions instance.
func New(c protocol.Config) (*Functions, error) {
	cs := &Functions{
		grpcServer:            grpc.NewServer(),
		entityDiscoveryServer: discovery.NewServer(c),
		eventSourcedServer:    eventsourced.NewServer(),
		crdtServer:            crdt.NewServer(),
		actionServer:          action.NewServer(),
		valueServer:           value.NewServer(),
	}
	protocol.RegisterEntityDiscoveryServer(cs.grpcServer, cs.entityDiscoveryServer)
	entity.RegisterEventSourcedServer(cs.grpcServer, cs.eventSourcedServer)
	entity.RegisterCrdtServer(cs.grpcServer, cs.crdtServer)
	entity.RegisterValueEntityServer(cs.grpcServer, cs.valueServer)
	entity.RegisterActionProtocolServer(cs.grpcServer, cs.actionServer)
	return cs, nil
}

// RegisterEventSourced registers an event sourced entity.
func (cs *Functions) RegisterEventSourced(entity *eventsourced.Entity, config protocol.DescriptorConfig, options ...eventsourced.Option) error {
	entity.Options(options...)
	if err := cs.eventSourcedServer.Register(entity); err != nil {
		return err
	}
	if err := cs.entityDiscoveryServer.RegisterEventSourcedEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// RegisterCRDT registers a CRDT entity.
func (cs *Functions) RegisterCRDT(entity *crdt.Entity, config protocol.DescriptorConfig, options ...crdt.Option) error {
	entity.Options(options...)
	if err := cs.crdtServer.Register(entity); err != nil {
		return err
	}
	if err := cs.entityDiscoveryServer.RegisterCRDTEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// RegisterAction registers an action entity.
func (cs *Functions) RegisterAction(entity *action.Entity, config protocol.DescriptorConfig) error {
	if err := cs.actionServer.Register(entity); err != nil {
		return err
	}
	if err := cs.entityDiscoveryServer.RegisterActionEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// RegisterValueEntity registers a Value entity.
func (cs *Functions) RegisterValueEntity(entity *value.Entity, config protocol.DescriptorConfig, options ...value.Option) error {
	entity.Options(options...)
	if err := cs.valueServer.Register(entity); err != nil {
		return err
	}
	if err := cs.entityDiscoveryServer.RegisterValueEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// Run runs the Functions instance on the interface and port defined by
// the HOST and PORT environment variable.
func (cs *Functions) Run() error {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		return errors.New("unable to get environment variable \"HOST\"")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("unable to get environment variable \"PORT\"")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	if err := cs.RunWithListener(lis); err != nil {
		return fmt.Errorf("failed to RunWithListener for: %v with: %w", lis, err)
	}
	return nil
}

// Run runs the Functions instance with a listener provided.
func (cs *Functions) RunWithListener(lis net.Listener) error {
	return cs.grpcServer.Serve(lis)
}

// Stop gracefully stops the Eigr Functions instance.
func (cs *Functions) Stop() {
	cs.grpcServer.GracefulStop()
	log.Println("Functions stopped")
}
