//
// Copyright 2019 Lightbend Inc.
// Copyright 2021 The eigr.io Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package action

import (
	"errors"

	"github.com/eigr/functions-go-sdk/functions/action"
	"github.com/eigr/functions-go-sdk/functions/encoding"
	"github.com/eigr/functions-go-sdk/functions/protocol"
	"github.com/golang/protobuf/proto"
)

type TestModel struct {
}

func NewTestModel() action.EntityHandler {
	return &TestModel{}
}

func (m *TestModel) HandleCommand(ctx *action.Context, name string, msg proto.Message) error {
	ctx.CloseFunc(func(c *action.Context) error {
		return nil
	})

	var failure error
	switch r := msg.(type) {
	case *Request:
		for _, g := range r.GetGroups() {
			// The single request may contain multiple grouped steps, each group corresponding to an expected response.
			for _, step := range g.GetSteps() {
				switch s := step.Step.(type) {
				case *ProcessStep_Reply:
					resp, err := encoding.MarshalAny(&Response{
						Message: s.Reply.GetMessage(),
					})
					if err != nil {
						return err
					}
					ctx.RespondWith(resp)
				case *ProcessStep_Forward:
					payload, err := encoding.MarshalAny(&OtherRequest{
						Id: s.Forward.GetId(),
					})
					if err != nil {
						return err
					}
					ctx.Forward(&protocol.Forward{
						ServiceName: "functions.tck.model.action.ActionTwo",
						CommandName: "Call",
						Payload:     payload,
						Metadata:    ctx.Metadata(),
					})
				case *ProcessStep_Effect:
					payload, err := encoding.MarshalAny(&OtherRequest{
						Id: s.Effect.GetId(),
					})
					if err != nil {
						return err
					}
					ctx.SideEffect(&protocol.SideEffect{
						ServiceName: "functions.tck.model.action.ActionTwo",
						CommandName: "Call",
						Payload:     payload,
						Synchronous: s.Effect.GetSynchronous(),
						Metadata:    ctx.Metadata(),
					})
				case *ProcessStep_Fail:
					failure = protocol.ClientError{Err: errors.New(s.Fail.Message)}
				}
			}
			if name == "ProcessStreamedOut" || name == "ProcessStreamed" {
				err := ctx.Respond(failure)
				if err != nil {
					return err
				}
			}
		}
	}
	if name == "ProcessStreamedOut" || name == "ProcessStreamed" {
		ctx.Cancel()
		return nil
	}
	return failure
}

type TestModelTwo struct {
}

func NewTestModelTwo() action.EntityHandler {
	return &TestModelTwo{}
}

func (m *TestModelTwo) HandleCommand(ctx *action.Context, name string, msg proto.Message) error {
	resp, err := encoding.MarshalAny(&Response{})
	if err != nil {
		return err
	}
	ctx.RespondWith(resp)
	return nil
}
