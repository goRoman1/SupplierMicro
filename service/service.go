package service

import (
	"SupplierMicro/proto"
	"SupplierMicro/repository"
	"context"
	"database/sql"
)

type SupplierMicroService struct {
	Repo *repository.SupplierMicroRepo
	*proto.UnimplementedSupplierMicroServiceServer
}

func NewSupplierMicroService(repo *sql.DB) *SupplierMicroService {
	return &SupplierMicroService{
		Repo: repository.NewSupplierMicroRepo(repo),
	}
}

func (ss *SupplierMicroService) CreateStation(ctx context.Context, station *proto.Station) (*proto.Response, error) {
	stationCreated, err := ss.Repo.CreateStation(ctx, station)
	return &proto.Response{
		Success:        err == nil,
		ScooterStation: stationCreated,
	}, err
}

func (ss *SupplierMicroService) GetLocations(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	_ = req
	locations, err := ss.Repo.GetLocations(ctx)
	return &proto.Response{
		Success:   err == nil,
		Locations: locations,
	}, err
}

func (ss *SupplierMicroService) GetStations(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	_ = req
	stations, err := ss.Repo.GetStations(ctx)
	return &proto.Response{
		Success:         err == nil,
		ScooterStations: stations,
	}, err
}

func (ss *SupplierMicroService) CreateStationInLocation(ctx context.Context, st *proto.StationLocation) (*proto.Response, error) {
	stationCreated, err := ss.Repo.CreateStationInLocation(ctx, st.ScooterStation, st.Location)
	return &proto.Response{
		Success:        err == nil,
		ScooterStation: stationCreated,
	}, err
}
