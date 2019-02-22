package event

import (
	"github.com/vikashvverma/eventers/pkg/api/event"
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// New creates new event logging service
func New(svc event.Service, logger eventers.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents event logging service
type LogService struct {
	event.Service
	logger eventers.Logger
}
