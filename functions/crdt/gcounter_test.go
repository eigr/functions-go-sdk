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

package crdt

import (
	"testing"

	"github.com/eigr/functions-go-sdk/functions/entity"
)

func TestGCounter(t *testing.T) {
	delta := func(incr uint64) *entity.CrdtDelta {
		return &entity.CrdtDelta{
			Delta: &entity.CrdtDelta_Gcounter{
				Gcounter: &entity.GCounterDelta{
					Increment: incr,
				},
			},
		}
	}

	t.Run("should reflect a delta update", func(t *testing.T) {
		c := NewGCounter()
		if err := c.applyDelta(encDecDelta(delta(10))); err != nil {
			t.Fatal(err)
		}
		if v := c.Value(); v != 10 {
			t.Fatalf("c.Value: %v; want: %d", v, 10)
		}
		if err := c.applyDelta(encDecDelta(delta(5))); err != nil {
			t.Fatal(err)
		}
		if v := c.Value(); v != 15 {
			t.Fatalf("c.Value: %v; want: %d", v, 15)
		}
	})

	t.Run("should generate deltas", func(t *testing.T) {
		c := GCounter{}
		c.Increment(10)
		if c.Delta().GetGcounter().GetIncrement() != 10 {
			t.Fatalf("counter increment: %v; want: %v", c.delta, 10)
		}
		c.resetDelta()
		if c.Delta().GetGcounter().GetIncrement() != 0 {
			t.Fatalf("counter increment: %v; want: %v", c.delta, 0)
		}
		c.Increment(3)
		if c.Value() != 13 {
			t.Fatalf("counter increment: %v; want: %v", c.delta, 13)
		}
		c.Increment(4)
		if c.Value() != 17 {
			t.Fatalf("counter increment: %v; want: %v", c.delta, 17)
		}
		// TODO: c.Delta().GetGcounter().GetIncrement() is 0 even if Delta() returns nil
		if c.Delta().GetGcounter().GetIncrement() != 7 {
			t.Fatalf("counter increment: %v; want: %v", c.delta, 7)
		}
		c.resetDelta()
		if d := c.Delta(); d == nil {
			t.Fatalf("c.Delta() should be not nil, but was")
		}
	})

	t.Run("should return its state", func(t *testing.T) {
		c := GCounter{}
		c.Increment(10)
		if v := c.Value(); v != 10 {
			t.Fatalf("c.Value: %v; want: %d", v, 10)
		}
		c.resetDelta()
		if d := c.Delta(); d == nil {
			t.Fatalf("c.Delta() should be not nil, but was")
		}
	})
}

func TestGCounterAdditional(t *testing.T) {
	t.Run("should report hasDelta", func(t *testing.T) {
		c := NewGCounter()
		if c.HasDelta() {
			t.Fatalf("c.HasDelta() but should not")
		}
		c.Increment(29)
		if !c.HasDelta() {
			t.Fatalf("c.HasDelta() is false, but should not")
		}
	})

	t.Run("should catch illegal delta applied", func(t *testing.T) {
		c := NewGCounter()
		err := c.applyDelta(&entity.CrdtDelta{
			Delta: &entity.CrdtDelta_Pncounter{
				Pncounter: &entity.PNCounterDelta{
					Change: 11,
				},
			},
		})
		if err == nil {
			t.Fatalf("c.applyDelta() has to err, but did not")
		}
	})
}
