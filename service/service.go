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

func (ss *SupplierService) CreateStation(station *proto.ScooterStation) (proto.Response, error) {
	return ss.Repo.CreateStation(station)
}

func (ss *SupplierService) GetLocations(context.Context, *proto.Request) ([]*proto.Location, error) {
	return ss.Repo.GetLocations(context.Background())
}

func (ss *SupplierService) CreateStationInLocation(ctx context.Context, st *proto.StationLocation) (proto.Response, error) {
	return ss.Repo.CreateStationInLocation(ctx, st.ScooterStation, st.Location)
}
