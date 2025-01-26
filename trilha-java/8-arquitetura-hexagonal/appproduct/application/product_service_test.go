package application_test

import (
	"github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/application"
	mock_application "github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
