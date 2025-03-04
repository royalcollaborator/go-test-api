package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/royalcollaborator/go-test-api/api/controller"
	"github.com/royalcollaborator/go-test-api/bootstrap"
	mocks "github.com/royalcollaborator/go-test-api/domain/mock"

	"github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

func TestLastTradePrice(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedErr    bool
		setExpec       func() *controller.TradeInfoController
	}{
		{
			name:           "Success Response",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedErr:    false,
			setExpec: func() *controller.TradeInfoController {
				ctrl := gomock.NewController(t)

				res := map[string]map[string]float64{
					"bitcoin": {
						"chf": 75367,
						"eur": 80249,
						"usd": 84080,
					},
				}
				mockTradeUsecase := mocks.NewMockTradeUsecase(ctrl)
				mockTradeUsecase.EXPECT().GetLastTradePrice().Return(http.StatusOK, res, nil)
				tc := &controller.TradeInfoController{
					TradeUsecase: mockTradeUsecase,
					Env:          &bootstrap.Env{},
				}
				return tc
			},
		},
		{
			name:           "Bad Geteway",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedErr:    false,
			setExpec: func() *controller.TradeInfoController {
				ctrl := gomock.NewController(t)

				mockTradeUsecase := mocks.NewMockTradeUsecase(ctrl)
				mockTradeUsecase.EXPECT().GetLastTradePrice().Return(http.StatusBadGateway, nil, errors.New("Request Not Reached"))
				tc := &controller.TradeInfoController{
					TradeUsecase: mockTradeUsecase,
					Env:          &bootstrap.Env{},
				}
				return tc
			},
		},
		{
			name:           "Internal Server Error",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedErr:    false,
			setExpec: func() *controller.TradeInfoController {
				ctrl := gomock.NewController(t)

				mockTradeUsecase := mocks.NewMockTradeUsecase(ctrl)
				mockTradeUsecase.EXPECT().GetLastTradePrice().Return(http.StatusOK, nil, errors.New("Internal Server Error"))
				tc := &controller.TradeInfoController{
					TradeUsecase: mockTradeUsecase,
					Env:          &bootstrap.Env{},
				}
				return tc
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rt := tc.setExpec()
			ginRouter := gin.Default()
			rec := httptest.NewRecorder()
			ginRouter.GET("/api/v1/ltp", rt.LastTradePrice)
			req := httptest.NewRequest(http.MethodGet, "/api/v1/ltp", nil)
			req.Header.Set("Content-Type", "application/json")
			ginRouter.ServeHTTP(rec, req)
		})
	}
}
