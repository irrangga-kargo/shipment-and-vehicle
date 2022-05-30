package graph

import "github.com/irrangga-kargo/kargo-trucks/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Trucks    []*model.Truck
	Shipments []*model.Shipment
}
