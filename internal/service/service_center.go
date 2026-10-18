package service

import "github.com/jojomoeys-source/VyshKat/internal/domain"

type ServiceCenter interface {
	Inspect(transport domain.TransportItem) bool
}

type OperationalServiceCenter struct{}

func NewOperationalServiceCenter() *OperationalServiceCenter {
	return &OperationalServiceCenter{}
}

func (s *OperationalServiceCenter) Inspect(transport domain.TransportItem) bool {
	return transport.IsOperational()
}
