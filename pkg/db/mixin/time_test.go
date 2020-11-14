package mixin

import (
	"testing"

	"github.com/facebook/ent/schema/mixin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeMixin(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		fields := mixin.CreateTime{}.Fields()
		require.Len(t, fields, 1)
		desc := fields[0].Descriptor()
		assert.Equal(t, "create_time", desc.Name)
		assert.True(t, desc.Immutable)
		assert.NotNil(t, desc.Default)
		assert.Nil(t, desc.UpdateDefault)
	})
	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		fields := mixin.UpdateTime{}.Fields()
		require.Len(t, fields, 1)
		desc := fields[0].Descriptor()
		assert.Equal(t, "update_time", desc.Name)
		assert.True(t, desc.Immutable)
		assert.NotNil(t, desc.Default)
		assert.NotNil(t, desc.UpdateDefault)
	})
	t.Run("Compose", func(t *testing.T) {
		t.Parallel()
		fields := mixin.Time{}.Fields()
		require.Len(t, fields, 2)
		assert.Equal(t, "create_time", fields[0].Descriptor().Name)
		assert.Equal(t, "update_time", fields[1].Descriptor().Name)
	})
}
