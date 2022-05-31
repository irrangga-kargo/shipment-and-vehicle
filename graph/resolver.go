package graph

import (
	"fmt"

	"github.com/irrangga-kargo/kargo-trucks/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Trucks    []*model.Truck
	Shipments []*model.Shipment
}

func (r *Resolver) Init() {
	for i := 0; i < 20; i++ {
		truck := &model.Truck{
			ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
			PlateNo: fmt.Sprintf("B %d CD", len(r.Trucks)+1),
		}
		r.Trucks = append(r.Trucks, truck)
	}
}
