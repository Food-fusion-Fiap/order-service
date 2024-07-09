package usecases

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/adapters/gateways"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/enums"
	"github.com/Food-fusion-Fiap/order-service/src/utils"
	"time"
)

type ChangeOrderStatusUsecase struct {
	OrderRepository gateways.OrderRepository
}

func (r *ChangeOrderStatusUsecase) Execute(orderId string, changeToStatus string) (*entities.Order, error) {
	var err error
	order, _ := r.OrderRepository.FindById(orderId)

	if order.ID == "" {
		return nil, errors.New("pedido não existe")
	}

	switch changeToStatus {
	case enums.Cancelled:
		order, err = changeToCancelled(order)
	case enums.Received:
		order, err = changeToReceived(order)
	case enums.Preparation:
		order, err = changeToPreparation(order)
	case enums.Ready:
		order, err = changeToReady(order)
	case enums.Finished:
		order, err = changeToFinished(order)
	default:
		return order, errors.New("status desconhecido")
	}

	if err != nil {
		return order, err
	}

	order.UpdatedAt = time.Now().Format(utils.CompleteEnglishDateFormat)
	order, err = r.OrderRepository.Update(order)

	return order, err
}

func changeToCancelled(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Created || order.PaymentStatus != enums.AwaitingPayment {
		return order, errors.New("não é possível cancelar o pedido")
	}

	order.Status = enums.Cancelled
	order.PaymentStatus = enums.Cancelled

	return order, nil
}

func changeToReceived(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Created || order.PaymentStatus != enums.AwaitingPayment {
		return order, errors.New("não é possível mudar o pedido para Recebido")
	}

	order.Status = enums.Received
	order.PaymentStatus = enums.Paid

	return order, nil
}

func changeToPreparation(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Received || order.PaymentStatus != enums.Paid {
		return order, errors.New("não é possível mudar o pedido para Em preparação")
	}

	order.Status = enums.Preparation

	return order, nil
}

func changeToReady(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Preparation || order.PaymentStatus != enums.Paid {
		return order, errors.New("não é possível mudar o pedido para Pronto")
	}

	order.Status = enums.Ready

	return order, nil
}

func changeToFinished(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Ready || order.PaymentStatus != enums.Paid {
		return order, errors.New("não é possível mudar o pedido para Finalizado")
	}

	order.Status = enums.Finished

	return order, nil
}
