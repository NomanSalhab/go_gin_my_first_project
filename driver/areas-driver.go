package driver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type AreaDriver interface {
	FindAllAreas() ([]entity.Area, error)
	FindActiveAreas() ([]entity.Area, error)
	FindNotActiveAreas() ([]entity.Area, error)
	FindArea(wantedId int) (entity.AreaEditRequest, error)

	AddArea(area entity.Area) error
	DeleteArea(wantedId int) error
	EditArea(areaEditInfo entity.AreaEditRequest) (entity.Area, error)

	ActivateArea(areaEditInfo entity.AreaActivateRequest) error
	DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error
}

type areaDriver struct {
	// cacheAreas []entity.AreaEditRequest
}

func NewAreaDriver() AreaDriver {
	return &areaDriver{}
}

func (driver *areaDriver) FindAllAreas() ([]entity.Area, error) {
	areas := make([]entity.Area, 0)
	rows, err := dbConn.SQL.Query("select id, name, lat, long, active from areas")
	if err != nil {
		return make([]entity.Area, 0), err
	}
	defer rows.Close()

	var id int
	var lat, long float32
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &lat, &long, &active)
		if err != nil {
			return make([]entity.Area, 0), err
		}
		areas = append(areas, entity.Area{
			ID:     id,
			Name:   name,
			Lat:    lat,
			Long:   long,
			Active: active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Area, 0), err
		}
	}

	return areas, nil
}

func (driver *areaDriver) FindActiveAreas() ([]entity.Area, error) {
	areas := make([]entity.Area, 0)
	rows, err := dbConn.SQL.Query("select id, name, lat, long, active from areas where active = true")
	if err != nil {
		return make([]entity.Area, 0), err
	}
	defer rows.Close()

	var id int
	var lat, long float32
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &lat, &long, &active)
		if err != nil {
			return make([]entity.Area, 0), err
		}
		areas = append(areas, entity.Area{
			ID:     id,
			Name:   name,
			Lat:    lat,
			Long:   long,
			Active: active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Area, 0), err
		}
	}

	return areas, nil
}

func (driver *areaDriver) FindNotActiveAreas() ([]entity.Area, error) {
	areas := make([]entity.Area, 0)
	rows, err := dbConn.SQL.Query("select id, name, lat, long, active from areas where active = false")
	if err != nil {
		return make([]entity.Area, 0), err
	}
	defer rows.Close()

	var id int
	var lat, long float32
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &lat, &long, &active)
		if err != nil {
			return make([]entity.Area, 0), err
		}
		areas = append(areas, entity.Area{
			ID:     id,
			Name:   name,
			Lat:    lat,
			Long:   long,
			Active: active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Area, 0), err
		}
	}

	return areas, nil
}

func (driver *areaDriver) AddArea(area entity.Area) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO areas(name, lat, long, active)
	VALUES ($1, $2, $3, $4) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, area.Name, area.Lat, area.Long, area.Active)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("area could not be added")
	}

	return nil
}

func (driver *areaDriver) DeleteArea(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from areas where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("area could not be found")
	}

	return nil
}

func (driver *areaDriver) FindArea(wantedId int) (entity.AreaEditRequest, error) {

	query := `select id, name, lat, long, active from areas where id=$1`
	var id int
	var name string
	var active bool
	var lat, long float32
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &lat, &long, &active)
	if err != nil {
		return entity.AreaEditRequest{
			ID:     0,
			Name:   "",
			Lat:    0.0,
			Long:   0.0,
			Active: false,
		}, err
	}
	area := entity.AreaEditRequest{
		ID:     id,
		Name:   name,
		Lat:    lat,
		Long:   long,
		Active: active,
	}

	return area, nil
}

func (driver *areaDriver) EditArea(areaEditInfo entity.AreaEditRequest) (entity.Area, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := GetEditAreaStatementString(areaEditInfo)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return entity.Area{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.Area{}, errors.New("area could not be found")
	}
	return entity.Area{}, nil
}

func GetEditAreaStatementString(areaEditInfo entity.AreaEditRequest) string {
	stmt := `UPDATE areas SET `
	if areaEditInfo.Name != "" {
		stmt = stmt + `name = '` + areaEditInfo.Name + `', `
	}
	if areaEditInfo.Lat != 0 {
		stmt = stmt + `lat = ` + fmt.Sprint(areaEditInfo.Lat) + `, `
	}
	if areaEditInfo.Long != 0 {
		stmt = stmt + `long = ` + fmt.Sprint(areaEditInfo.Long) + `, `
	}
	if areaEditInfo.Active {
		stmt = stmt + `active = true `
	} else {
		stmt = stmt + `active = false `
	}
	stmt = stmt + `where id = ` + fmt.Sprint(areaEditInfo.ID) + ` RETURNING *`
	return stmt
}

func (driver *areaDriver) ActivateArea(areaEditInfo entity.AreaActivateRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE areas SET active = true WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, areaEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("area could not be found")
	}
	return nil
}

func (driver *areaDriver) DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE areas SET active = false WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, areaEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("area could not be found")
	}
	return nil
}
