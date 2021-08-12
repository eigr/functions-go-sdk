package main

import (
	"log"

	"github.com/eigr/functions-go-sdk/cloudstate"
	"github.com/eigr/functions-go-sdk/cloudstate/crdt"
	"github.com/eigr/functions-go-sdk/cloudstate/protocol"
	"github.com/eigr/functions-go-sdk/example/crdt_shoppingcart/shoppingcart"
)

// main creates a CloudState instance and registers the ShoppingCart
// as a event sourced entity.
func main() {
	server, err := cloudstate.New(protocol.Config{
		ServiceName:    "shopping-cart",
		ServiceVersion: "0.1.0",
	})
	if err != nil {
		log.Fatalf("cloudstate.New failed: %v", err)
	}

	// tag::register-crdt[]
	err = server.RegisterCRDT(&crdt.Entity{
		ServiceName: "example.shoppingcart.ShoppingCartService",
		EntityFunc:  shoppingcart.NewShoppingCart,
	}, protocol.DescriptorConfig{
		Service: "shoppingcart.proto",
	}.AddDomainDescriptor("domain.proto", "hotitems.proto"))
	// end::register-crdt[]
	if err != nil {
		log.Fatalf("CloudState failed to register entity: %v", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("CloudState failed to run: %v", err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}
