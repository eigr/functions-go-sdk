package valueentity

import (
	"errors"

	"github.com/eigr/functions-go-sdk/functions/encoding"
	"github.com/eigr/functions-go-sdk/functions/protocol"
	"github.com/eigr/functions-go-sdk/functions/value"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

type ValueEntityTckModelEntity struct {
	state string
}

func NewValueEntityTckModelEntity(id value.EntityID) value.EntityHandler {
	return &ValueEntityTckModelEntity{}
}

func (e *ValueEntityTckModelEntity) HandleCommand(ctx *value.Context, name string, msg proto.Message) (*any.Any, error) {
	var forward = false
	var failure error = nil
	switch m := msg.(type) {
	case *Request:
		for _, action := range m.GetActions() {
			switch a := action.Action.(type) {
			case *RequestAction_Update:
				e.state = a.Update.Value
				failure = ctx.Update(encoding.MarshalAny(&Persisted{
					Value: e.state,
				}))
			case *RequestAction_Delete:
				ctx.Delete()
				e.state = ""
			case *RequestAction_Forward:
				forward = true
				req, err := encoding.MarshalAny(&Request{
					Id: a.Forward.Id,
				})
				if err != nil {
					return nil, err
				}
				ctx.Forward(&protocol.Forward{
					ServiceName: "functions.tck.model.valueentity.ValueEntityTwo",
					CommandName: "Call",
					Payload:     req,
				})
			case *RequestAction_Effect:
				req, err := encoding.MarshalAny(&Request{
					Id: a.Effect.Id,
				})
				if err != nil {
					return nil, err
				}
				ctx.SideEffect(&protocol.SideEffect{
					ServiceName: "functions.tck.model.valueentity.ValueEntityTwo",
					CommandName: "Call",
					Payload:     req,
					Synchronous: a.Effect.GetSynchronous(),
				})
			case *RequestAction_Fail:
				failure = protocol.ClientError{Err: errors.New(a.Fail.GetMessage())}
			}
		}
	}
	if forward {
		return nil, nil
	}
	if failure != nil {
		return nil, failure
	}
	return encoding.MarshalAny(&Response{
		Message: e.state,
	})
}

func (e *ValueEntityTckModelEntity) HandleState(ctx *value.Context, state *any.Any) error {
	var p Persisted
	if err := encoding.UnmarshalAny(state, &p); err != nil {
		return err
	}
	e.state = p.GetValue()
	return nil
}

type ValueEntityTckModelEntityTwo struct {
}

func NewValueEntityTckModelEntityTwo(id value.EntityID) value.EntityHandler {
	return &ValueEntityTckModelEntityTwo{}
}

func (e *ValueEntityTckModelEntityTwo) HandleCommand(ctx *value.Context, name string, msg proto.Message) (*any.Any, error) {
	return encoding.MarshalAny(&Response{})
}

func (e *ValueEntityTckModelEntityTwo) HandleState(ctx *value.Context, state *any.Any) error {
	return nil
}

type ValueEntityConfiguredEntity string

func NewValueEntityConfiguredEntity(id value.EntityID) value.EntityHandler {
	r := ValueEntityConfiguredEntity("state")
	return &r
}

func (v *ValueEntityConfiguredEntity) HandleCommand(ctx *value.Context, name string, msg proto.Message) (*any.Any, error) {
	return encoding.MarshalAny(&Response{})
}

func (v *ValueEntityConfiguredEntity) HandleState(ctx *value.Context, state *any.Any) error {
	return nil
}
