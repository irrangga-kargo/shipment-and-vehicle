package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/irrangga-kargo/kargo-trucks/graph/generated"
	"github.com/irrangga-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	truck := &model.Truck{
		ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
		PlateNo: plateNo,
	}

	listPlateNo := strings.Split(plateNo, " ")
	if len(listPlateNo) == 3 {
		plateStart := listPlateNo[0]
		if len(plateStart) < 1 || len(plateStart) > 2 {
			return nil, errors.New("PLATE_NO_INVALID")
		}
		for _, v := range plateStart {
			if v < 'A' || v > 'Z' {
				return nil, errors.New("PLATE_NO_INVALID")
			}
		}

		nomor, errNomor := strconv.Atoi(listPlateNo[1])
		if errNomor != nil {
			return nil, errors.New("PLATE_NO_INVALID")
		}
		if nomor < 1 && nomor > 9999 {
			return nil, errors.New("PLATE_NO_INVALID")
		}

		plateEnd := listPlateNo[2]
		if len(plateEnd) < 1 || len(plateEnd) > 3 {
			return nil, errors.New("PLATE_NO_INVALID")
		}
		for _, v := range plateEnd {
			if v < 'A' || v > 'Z' {
				return nil, errors.New("PLATE_NO_INVALID")
			}
		}
	} else {
		return nil, errors.New("PLATE_NO_INVALID")
	}

	r.Trucks = append(r.Trucks, truck)
	return truck, nil
}

func (r *queryResolver) PaginatedTrucks(ctx context.Context) ([]*model.Truck, error) {
	return r.Trucks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
