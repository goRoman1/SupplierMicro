package service

import (
	"SupplierMicro/proto"
	"SupplierMicro/repository"
	"context"
	"database/sql"
)

type SupplierService struct {
	Repo *repository.SupplierRepo
	*proto.UnimplementedSupplierServiceServer
}

func NewSupplierService(repo *sql.DB) *SupplierService {
	return &SupplierService{
		Repo: repository.NewSupplierRepo(repo),
	}
}

func (ss *SupplierService) CreateStation(ctx context.Context, station *proto.Station) (*proto.Response, error) {
	stationCreated, err := ss.Repo.CreateStation(ctx, station)
	//goland:noinspection ALL
	return &proto.Response{
		Success:        err == nil,
		ScooterStation: stationCreated,
	}, err
}

func (ss *SupplierService) GetLocations(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	_ = req
	locations, err := ss.Repo.GetLocations(ctx)
	return &proto.Response{
		Success:   err == nil,
		Locations: locations,
	}, err
}

func (ss *SupplierService) CreateStationInLocation(ctx context.Context, st *proto.StationLocation) (*proto.Response, error) {
	stationCreated, err := ss.Repo.CreateStationInLocation(ctx, st.ScooterStation, st.Location)
	return &proto.Response{
		Success:        err == nil,
		ScooterStation: stationCreated,
	}, err
}
