package repository

import (
	"SupplierMicro/proto"
	"context"
	"database/sql"
)

type SupplierRepositoryServer interface {
	CreateStation(ctx context.Context, station *proto.ScooterStation) error
	GetLocations(context.Context, *proto.Request) (*proto.Locations, error)
	CreateStationInLocation(ctx context.Context, station *proto.ScooterStation, location *proto.Location) error
}

type SupplierRepo struct {
	db *sql.DB
}

func NewSupplierRepo(db *sql.DB) *SupplierRepo {
	return &SupplierRepo{db: db}
}

func (repo *SupplierRepo) CreateStation(station *proto.ScooterStation) (proto.Response, error) {
	query := `INSERT INTO scooter_stations(id, name, is_active, latitude, longitude)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id`
	row := repo.db.QueryRowContext(context.Background(), query, station.Id, station.Name, station.IsActive, station.Latitude, station.Longitude)
	err := row.Scan(&station.Id)
	if err != nil {
		return proto.Response{}, err
	}

	return proto.Response{}, nil
}

func (repo *SupplierRepo) GetLocations(context.Context) ([]*proto.Location, error) {
	query := `SELECT id, latitude, longitude, label FROM locations ORDER BY id;`
	row, err := repo.db.QueryContext(context.Background(), query)

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

func (repo *SupplierRepo) CreateStationInLocation(station *proto.ScooterStation, location *proto.Location) (proto.Response, error) {
	query := `INSERT INTO scooter_stations (name, is_active, latitude, longitude)
	VALUES($1, $2, $3, $4)
	RETURNING id`
	row := repo.db.QueryRowContext(context.Background(), query, station.Name, station.IsActive, location.Latitude, location.Longitude)
	err := row.Scan(&station.Id)
	if err != nil {
		return proto.Response{}, err
	}

	return proto.Response{}, nil
}
