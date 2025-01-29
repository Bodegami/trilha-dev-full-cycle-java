package cli_test

import (
	"fmt"
	"github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/adapter/cli"
	mock_application "github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var productName = "Product Test"
var productPrice = 25.99
var productId = "abc"
var productStatus = "enabled"
var productMock *mock_application.MockProductInterface
var service *mock_application.MockProductServiceInterface

func setUp(ctrl *gomock.Controller) {
	productMock = mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service = mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
}

func TestRun_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	setUp(ctrl)

	resultExpected := fmt.Sprintf("Product ID %s with the name %s and the price %f  and status %s was created",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	setUp(ctrl)

	resultExpected := fmt.Sprintf("Product %s has been enabled", productName)

	result, err := cli.Run(service, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	setUp(ctrl)

	resultExpected := fmt.Sprintf("Product %s has been disabled", productName)

	result, err := cli.Run(service, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	setUp(ctrl)

	resultExpected := fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "get", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
