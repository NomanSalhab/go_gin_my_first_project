package driver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

func FindAllAreas() ([]entity.Area, error) {
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

func FindActiveAreas() ([]entity.Area, error) {
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

func FindNotActiveAreas() ([]entity.Area, error) {
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

func AddArea(area entity.Area) error {

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

func DeleteArea(wantedId int) error {

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

func EditArea(areaEditInfo entity.AreaEditRequest) (entity.Area, error) {

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

func ActivateArea(areaEditInfo entity.AreaActivateRequest) error {

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

func DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error {

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
