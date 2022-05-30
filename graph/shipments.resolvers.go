package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/irrangga-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, page int, first int, truckID string) (*model.Shipment, error) {
	truck := &model.Truck{}
	for _, v := range r.Trucks {
		if v.ID == truckID {
			truck = &model.Truck{
				ID:      truckID,
				PlateNo: v.PlateNo,
			}
		}
	}

	shipment := &model.Shipment{
		ID:           fmt.Sprintf("SHIPMENT-%d", len(r.Trucks)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: deliveryDate,
		Page:         page,
		First:        first,
		Truck:        truck,
	}
	r.Shipments = append(r.Shipments, shipment)
	return shipment, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context) ([]*model.Shipment, error) {
	return r.Shipments, nil
}
