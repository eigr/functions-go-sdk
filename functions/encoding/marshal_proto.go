//
// Copyright 2019 Lightbend Inc.
// Copyright 2022 The eigr.io Authors.
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

package encoding

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

var ErrMarshal = errors.New("unable to marshal a message")

// MarshalAny marshals a proto.Message to an any.Any value.
func MarshalAny(pb interface{}) (*any.Any, error) {
	// TODO: protobufs are expected here, but Cloudstate supports other formats
	message, ok := pb.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("got a non-proto message as protobuf: %v", pb)
	}
	bytes, err := proto.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, ErrMarshal)
	}
	return &any.Any{
		TypeUrl: fmt.Sprintf("%s/%s", ProtoAnyBase, proto.MessageName(message)),
		Value:   bytes,
	}, nil
}

func UnmarshalAny(x *any.Any, p proto.Message) error {
	return proto.Unmarshal(x.Value, p)
}
