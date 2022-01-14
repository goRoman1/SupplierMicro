package repository

import (
	"SupplierMicro/proto"
	"context"
	"database/sql"
)

type SupplierMicroRepositoryServer interface {
	CreateStation(ctx context.Context, station *proto.Station) error
	GetLocations(ctx context.Context, req *proto.Request) (*[]proto.Location, error)
	CreateStationInLocation(ctx context.Context, station *proto.Station, location *proto.Location) error
}

type SupplierMicroRepo struct {
	db *sql.DB
}

func NewSupplierMicroRepo(db *sql.DB) *SupplierMicroRepo {
	return &SupplierMicroRepo{db: db}
}

func (repo *SupplierMicroRepo) CreateStation(ctx context.Context, station *proto.Station) (*proto.Station, error) {
	query := `INSERT INTO scooter_stations(id, name, is_active, latitude, longitude)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id`
	row := repo.db.QueryRowContext(ctx, query, station.Id, station.Name, station.IsActive, station.Latitude, station.Longitude)
	err := row.Scan(&station.Id)
	if err != nil {
		return station, err
	}

	return station, nil
}

func (repo *SupplierMicroRepo) GetLocations(ctx context.Context) ([]*proto.Location, error) {
	query := `SELECT id, latitude, longitude, label FROM locations ORDER BY id;`
	row, err := repo.db.QueryContext(ctx, query)

	var result []*proto.Location
	if err != nil {
		return result, err
	}
	defer row.Close()
	for row.Next() {
		var location proto.Location
		err := row.Scan(&location.Id, &location.Latitude, &location.Longitude, &location.Label)
		if err != nil {
			return result, err
		}
		result = append(result, &location)
	}
	return result, nil
}

func (repo *SupplierMicroRepo) CreateStationInLocation(ctx context.Context, station *proto.Station, location *proto.Location) (*proto.Station, error) {
	query := `INSERT INTO scooter_stations (name, is_active, latitude, longitude)
	VALUES($1, $2, $3, $4)
	RETURNING id`
	row := repo.db.QueryRowContext(ctx, query, station.Name, station.IsActive, location.Latitude, location.Longitude)
	err := row.Scan(&station.Id)
	if err != nil {
		return station, err
	}

	return station, nil
}
