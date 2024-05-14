package rocket

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get rocket by id", func(t *testing.T) {
		mockStore := NewMockStore(mockCtrl)
		id := "UUID-1"

		mockStore.EXPECT().
			GetRocketByID(id).
			Return(Rocket{
				ID: id,
			}, nil)

		service := New(mockStore)
		rocket, err := service.GetRocketByID(
			context.Background(),
			id,
		)
		assert.NoError(t, err)
		assert.Equal(t, id, rocket.ID)
	})

	t.Run("tests insert rocket", func(t *testing.T) {
		mockStore := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocket := Rocket{
			ID: id,
		}

		mockStore.EXPECT().
			InsertRocket(rocket).
			Return(rocket, nil)

		service := New(mockStore)
		rocket, err := service.InsertRocket(
			context.Background(),
			rocket,
		)
		assert.NoError(t, err)
		assert.Equal(t, id, rocket.ID)
	})

	t.Run("tests delete rocket", func(t *testing.T) {
		mockStore := NewMockStore(mockCtrl)
		id := "UUID-1"

		mockStore.EXPECT().
			DeleteRocket(id).
			Return(nil)

		service := New(mockStore)
		err := service.DeleteRocket(
			context.Background(),
			id,
		)
		assert.NoError(t, err)
	})
}
