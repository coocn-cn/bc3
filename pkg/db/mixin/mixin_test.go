// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package mixin_test

import (
	"testing"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/mixin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type annotation string

func (annotation) Name() string { return "" }

func TestAnnotateFields(t *testing.T) {
	annotations := []schema.Annotation{
		annotation("foo"),
		annotation("bar"),
		annotation("baz"),
	}
	fields := mixin.AnnotateFields(
		mixin.Time{}, annotations...,
	).Fields()
	require.Len(t, fields, 2)
	for _, f := range fields {
		desc := f.Descriptor()
		require.Len(t, desc.Annotations, len(annotations))
		for i := range desc.Annotations {
			assert.Equal(t, annotations[i], desc.Annotations[i])
		}
	}
}

type TestSchema struct {
	ent.Schema
}

func (TestSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("one", TestSchema.Type),
		edge.From("two", TestSchema.Type).
			Ref("one"),
	}
}

func TestAnnotateEdges(t *testing.T) {
	annotations := []schema.Annotation{
		annotation("foo"),
		annotation("bar"),
		annotation("baz"),
	}
	edges := mixin.AnnotateEdges(
		TestSchema{}, annotations...,
	).Edges()
	require.Len(t, edges, 2)
	for _, e := range edges {
		desc := e.Descriptor()
		require.Len(t, desc.Annotations, len(annotations))
		for i := range desc.Annotations {
			assert.Equal(t, annotations[i], desc.Annotations[i])
		}
	}
}
