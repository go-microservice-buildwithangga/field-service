package controllers

import (
	fieldController "field-service/controllers/field"
	fieldScheduleController "field-service/controllers/fieldschedule"
	timeController "field-service/controllers/time"
	"field-service/services"
)

type Registry struct {
	services services.IServiceRegistry
}

// GetField implements IControllerRegistry.

type IControllerRegistry interface {
	GetField() fieldController.IFieldController
	GetFieldSchedule() fieldScheduleController.IFieldScheduleController

	GetTime() timeController.ITimeController
}

func (r *Registry) NewControllerRegistry(services services.IServiceRegistry) IControllerRegistry {
	return &Registry{
		services: services,
	}
}

// GetField implements IControllerRegistry.
func (r *Registry) GetField() fieldController.IFieldController {
	return fieldController.NewFieldController(r.services)
}

// GetFieldSchedule implements IControllerRegistry.
func (r *Registry) GetFieldSchedule() fieldScheduleController.IFieldScheduleController {
	return fieldScheduleController.NewFieldScheduleController(r.services)
}

// GetTime implements IControllerRegistry.
func (r *Registry) GetTime() timeController.ITimeController {
	return timeController.NewTimeController(r.services)
}
