package services

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"
)

type OrderService interface {
	GetAllOrder() []dto.OrderResponse
	GetOrderByID(id int) (dto.OrderResponse, *helpers.AppError)
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return &orderService{orderRepository}
}

func (service *orderService) GetAllOrder() []dto.OrderResponse {
	orders := service.orderRepository.GetAllOrder()

	var orderResponses []dto.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, dto.OrderResponse{
			OrderID:      order.OrderID,
			UserID:       order.UserID,
			CustomerName: order.CustomerName,
			Amount:       order.Amount,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
		})
	}
	return orderResponses
}

func (service *orderService) GetOrderByID(id int) (dto.OrderResponse, *helpers.AppError) {
	order, err := service.orderRepository.GetOrderByID(id)
	if err != nil {
		return dto.OrderResponse{}, helpers.NewUnexpectedError("Internal Server Error")
	}
	return dto.OrderResponse{
		OrderID:      order.OrderID,
		UserID:       order.UserID,
		CustomerName: order.CustomerName,
		Amount:       order.Amount,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}, nil
}
