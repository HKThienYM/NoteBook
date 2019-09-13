package services

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	brandManagerment "thien.com/gRPC_Gateway/pb"
	repo "thien.com/server/Repository"
)

type BrandService struct{}

func (b *BrandService) GetBrand(context.Context, *brandManagerment.Foo) (*brandManagerment.ListBrands, error) {
	lb := repo.GetBrand()
	if lb == nil {
		return nil, errors.New("Get Brand failed!")
	}
	var response brandManagerment.ListBrands
	for _, v := range lb {
		response.ListBrands = append(response.ListBrands, &brandManagerment.Brand{Name: v.Name})
	}
	return &response, nil
}

func (b *BrandService) GetBrandVehicle(ctx context.Context, brand *brandManagerment.Brand) (*brandManagerment.ListVehicle, error) {
	lv := repo.GetBrandVehicle(brand.GetName())
	if lv != nil {
		return nil, errors.New("Can't get Vehicle from that brand!")
	}
	var response brandManagerment.ListVehicle
	for _, v := range lv {
		response.ListVehicles = append(response.ListVehicles, &brandManagerment.Vehicle{Name: v.Name})
	}
	return &response, nil
}

func RegisterService(ser *grpc.Server) {
	brandManagerment.RegisterBrandManagerServer(ser, &BrandService{})
}

func GetAllVehicle() []repo.Vehicle {
	return repo.GetAllVehicle()
}
