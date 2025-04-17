package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitea.qwertysystem.net/BETS/ts-utils/entity"
)

func TestMakeMap(t *testing.T) {
	t.Run("map is not nil", func(t *testing.T) {
		m := entity.MakeMap[[]int, []int]()

		assert.NotNil(t, m)
	})
}

func TestMap(t *testing.T) {
	t.Run("successfully test each function", func(t *testing.T) {
		m := entity.MakeMap[[]int, []int]()

		// Set
		m.Set([]int{1, 2, 3}, []int{3, 2, 1})
		m.Set([]int{3, 2, 1}, []int{1, 2, 3})

		// Exist
		assert.True(t, m.Exist([]int{1, 2, 3}))
		assert.True(t, m.Exist([]int{3, 2, 1}))

		// Get
		assert.Equal(t, []int{3, 2, 1}, m.Get([]int{1, 2, 3}, []int{}))
		assert.Equal(t, []int{1, 2, 3}, m.Get([]int{3, 2, 1}, []int{}))

		// Delete
		m.Delete([]int{1, 2, 3})
		m.Delete([]int{3, 2, 1})

		assert.False(t, m.Exist([]int{1, 2, 3}))
		assert.False(t, m.Exist([]int{3, 2, 1}))
	})

	t.Run("set: if key exists previously, overwrite the value", func(t *testing.T) {
		m := entity.MakeMap[[]int, []int]()

		m.Set([]int{1, 2, 3}, []int{3, 2, 1})

		assert.Equal(t, []int{3, 2, 1}, m.Get([]int{1, 2, 3}, []int{}))

		m.Set([]int{1, 2, 3}, []int{1, 2, 3})

		assert.Equal(t, []int{1, 2, 3}, m.Get([]int{1, 2, 3}, []int{}))
	})

	t.Run("get: use fallback value if key doesn't exists", func(t *testing.T) {
		m := entity.MakeMap[[]int, []int]()

		assert.Equal(t, []int{}, m.Get([]int{1, 2, 3}, []int{}))
	})
}
