package db

import (
	"context"
	"errors"
	"fmt"
	"device-service/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &DB{
		Pool: pool,
	}, nil
}

func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

func (db *DB) CreateDevice(ctx context.Context, device models.Device) (models.Device, error) {
	query := `
		INSERT INTO devices (id, name, type, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, type, status
	`

	var createdDevice models.Device

	err := db.Pool.QueryRow(
		ctx,
		query,
		device.ID,
		device.Name,
		device.Type,
		device.Status,
	).Scan(
		&createdDevice.ID,
		&createdDevice.Name,
		&createdDevice.Type,
		&createdDevice.Status,
	)

	if err != nil {
		return models.Device{}, fmt.Errorf("error creating device: %w", err)
	}

	return createdDevice, nil
}

func (db *DB) GetDevices(ctx context.Context) ([]models.Device, error) {
	query := `
		SELECT id, name, type, status
		FROM devices
		ORDER BY name
	`

	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying devices: %w", err)
	}
	defer rows.Close()

	var devices []models.Device

	for rows.Next() {
		var device models.Device

		if err := rows.Scan(
			&device.ID,
			&device.Name,
			&device.Type,
			&device.Status,
		); err != nil {
			return nil, fmt.Errorf("error scanning device row: %w", err)
		}

		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating device rows: %w", err)
	}

	return devices, nil
}

func (db *DB) GetDeviceByID(ctx context.Context, id string) (models.Device, error) {
	query := `
		SELECT id, name, type, status
		FROM devices
		WHERE id = $1
	`

	var device models.Device

	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&device.ID,
		&device.Name,
		&device.Type,
		&device.Status,
	)

	if err != nil {
		return models.Device{}, fmt.Errorf("error getting device by ID: %w", err)
	}

	return device, nil
}

func (db *DB) UpdateDevice(ctx context.Context, device models.Device) (models.Device, error) {
	query := `
		UPDATE devices
		SET name = $1, type = $2, status = $3
		WHERE id = $4
		RETURNING id, name, type, status
	`

	var updatedDevice models.Device

	err := db.Pool.QueryRow(
		ctx,
		query,
		device.Name,
		device.Type,
		device.Status,
		device.ID,
	).Scan(
		&updatedDevice.ID,
		&updatedDevice.Name,
		&updatedDevice.Type,
		&updatedDevice.Status,
	)

	if err != nil {
		return models.Device{}, fmt.Errorf("error updating device: %w", err)
	}

	return updatedDevice, nil
}

func (db *DB) DeleteDevice(ctx context.Context, id string) error {
	query := `
		DELETE FROM devices
		WHERE id = $1
	`

	result, err := db.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting device: %w", err)
	}

	if result.RowsAffected() == 0 {
		return errors.New("device not found")
	}

	return nil
}