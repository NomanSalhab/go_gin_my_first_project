package driver

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type UserDriver interface {
	FindAllUsers() ([]entity.User, error)
	FindActiveUsers() ([]entity.User, error)
	FindNotActiveUsers() ([]entity.User, error)

	FindUser(wantedId int) (entity.User, error)
	FindUserAddresses(wantedId int) ([]entity.Address, error)
	FindUserCircles(wantedId int) (entity.UserCirclesResponse, error)
	FindUserRateAndCircles(wantedId int) (entity.UserCirclesResponse, error)

	DeleteUser(wantedId int) error
	AddUser(user entity.User) error
	LoginUser(userLoginInfo entity.UserLoginRequest) (entity.User, error)
	EditUser(userEditInfo entity.UserEditRequest) (entity.User, error)
	ActivateUser(userInfo entity.UserInfoRequest) error
	DeactivateUser(userInfo entity.UserInfoRequest) error

	UserAddAddress(address entity.AddAddressRequest) error
	UserDeleteAddress(wantedId int) error

	GetEditUserStatementString(userEditInfo entity.UserEditRequest) string

	EditUserBalanceAndCircles(userId int, deliveryCost int, productsCost int) error
	EditDeliveryWorkerBalanceAndCircles(userId int, deliveryCost int) error

	SpecializeUser(userInfo entity.UserInfoRequest) error
	NormalizeUser(userInfo entity.UserInfoRequest) error
	ChangeUserRole(userInfo entity.UserChangeRoleRequest) error
}

type userDriver struct {
	// cachedUsers []entity.User
}

func NewUserDriver() UserDriver {
	return &userDriver{}
}

func (driver *userDriver) FindAllUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := dbConn.SQL.Query("select id, name, phone, balance, active, circles, role, user_discount, area_id, special_user from users")
	if err != nil {
		return make([]entity.User, 0), err
	}
	defer rows.Close()

	var id, balance, circles, role, areaId int
	// var deliveryTime time.Time
	var name, phone string
	var active, specialUser bool
	var userDiscount float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &phone, &balance, &active, &circles, &role, &userDiscount, &areaId, &specialUser)
		if err != nil {
			// log.Println(err)
			return make([]entity.User, 0), err
		}
		area, err := ad.FindArea(areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.User, 0), err
		}
		users = append(users, entity.User{
			ID:           id,
			Name:         name,
			Phone:        phone,
			Password:     "",
			Balance:      balance,
			Active:       active,
			Circles:      circles,
			Role:         role,
			UserDiscount: userDiscount,
			Area:         area,
			SpecialUser:  specialUser,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.User, 0), err
		}
		// fmt.Println("------------------------")
	}

	return users, nil
}

func (driver *userDriver) FindActiveUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := dbConn.SQL.Query("select id, name, phone, balance, active, circles, role, user_discount, area_id, special_user from users where active = true")
	if err != nil {
		return make([]entity.User, 0), err
	}
	defer rows.Close()

	var id, balance, circles, role, areaId int
	// var deliveryTime time.Time
	var name, phone string
	var active, specialUser bool
	var userDiscount float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &phone, &balance, &active, &circles, &role, &userDiscount, &areaId, &specialUser)
		if err != nil {
			// log.Println(err)
			return make([]entity.User, 0), err
		}
		area, err := ad.FindArea(areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.User, 0), err
		}
		users = append(users, entity.User{
			ID:           id,
			Name:         name,
			Phone:        phone,
			Password:     "",
			Balance:      balance,
			Active:       active,
			Circles:      circles,
			Role:         role,
			UserDiscount: userDiscount,
			Area:         area,
			SpecialUser:  specialUser,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.User, 0), err
		}
		// fmt.Println("------------------------")
	}

	return users, nil
}

func (driver *userDriver) FindNotActiveUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := dbConn.SQL.Query("select id, name, phone, balance, active, circles, role, user_discount, area_id, special_user from users where active = false")
	if err != nil {
		return make([]entity.User, 0), err
	}
	defer rows.Close()

	var id, balance, circles, role, areaId int
	// var deliveryTime time.Time
	var name, phone string
	var active, specialUser bool
	var userDiscount float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &phone, &balance, &active, &circles, &role, &userDiscount, &areaId, &specialUser)
		if err != nil {
			// log.Println(err)
			return make([]entity.User, 0), err
		}
		area, err := ad.FindArea(areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.User, 0), err
		}
		users = append(users, entity.User{
			ID:           id,
			Name:         name,
			Phone:        phone,
			Password:     "",
			Balance:      balance,
			Active:       active,
			Circles:      circles,
			Role:         role,
			UserDiscount: userDiscount,
			Area:         area,
			SpecialUser:  specialUser,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.User, 0), err
		}
		// fmt.Println("------------------------")
	}

	return users, nil
}

func (driver *userDriver) FindUser(wantedId int) (entity.User, error) {

	query := `
	select 
		id, name, phone, balance, active, circles, role, user_discount, area_id, special_user 
	from users where id = $1`
	var id, balance, circles, role, areaId int
	var name, phone string
	var active, specialUser bool
	var userDiscount float32
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &phone, &balance, &active, &circles, &role, &userDiscount, &areaId, &specialUser)
	// fmt.Println("User Data:", id, name, phone)
	if err != nil {
		return entity.User{
			ID:       0,
			Name:     "",
			Phone:    "",
			Password: "",
			Balance:  0,
			Active:   false,
			Circles:  0,
			Role:     0,
		}, err
	}
	area, err := ad.FindArea(areaId)
	if err != nil {
		// log.Println(err)
		return entity.User{}, err
	}
	user := entity.User{
		ID:           id,
		Name:         name,
		Phone:        phone,
		Password:     "",
		Balance:      balance,
		Active:       active,
		Circles:      circles,
		Role:         role,
		UserDiscount: userDiscount,
		Area:         area,
		SpecialUser:  specialUser,
	}
	return user, nil
}

func (driver *userDriver) FindUserAddresses(wantedId int) ([]entity.Address, error) {
	var addresses []entity.Address
	rows, err := dbConn.SQL.Query("select id, name, description, latitude, longitude, user_id from addresses where user_id = $1", wantedId)
	if err != nil {
		return make([]entity.Address, 0), err
	}
	defer rows.Close()

	var id, userId int
	var name, description string
	var latitude, longitude float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &latitude, &longitude, &userId)
		if err != nil {
			// log.Println(err)
			return make([]entity.Address, 0), err
		}
		addresses = append(addresses, entity.Address{
			ID:           id,
			Name:         name,
			Descripition: description,
			Latitude:     latitude,
			Longitude:    longitude,
			UserId:       userId,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Address, 0), err
		}
		// fmt.Println("------------------------")
	}

	return addresses, nil
}

func (driver *userDriver) FindUserCircles(wantedId int) (entity.UserCirclesResponse, error) {
	// var userCircles entity.UserCirclesResponse
	userCirclesAndRate, err1 := driver.FindUserRateAndCircles(wantedId)
	if err1 != nil {
		return entity.UserCirclesResponse{
			Circles: userCirclesAndRate.Circles,
			Rate:    userCirclesAndRate.Rate,
		}, err1
	}
	// var circles int

	// rows, err := dbConn.SQL.Query("select circles from users where user_id = $1", wantedId)
	// if err != nil {
	// 	return entity.UserCirclesResponse{
	// 		Circles: 0,
	// 		Rate:    10000,
	// 	}, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	err := rows.Scan(&circles)
	// 	if err != nil {
	// 		// log.Println(err)
	// 		return entity.UserCirclesResponse{Circles: 0, Rate: rate}, err
	// 	}
	// 	userCircles = entity.UserCirclesResponse{
	// 		Circles: circles,
	// 		Rate:    rate,
	// 	}
	// 	// fmt.Println("Record is:", userId, deliveryTime, products)
	// 	if err = rows.Err(); err != nil {
	// 		// log.Fatal("error Scanning Rows!")
	// 		return entity.UserCirclesResponse{
	// 			Circles: 0,
	// 			Rate:    rate,
	// 		}, err
	// 	}
	// 	// fmt.Println("------------------------")
	// }

	return userCirclesAndRate, nil
}

func (driver *userDriver) FindUserRateAndCircles(wantedId int) (entity.UserCirclesResponse, error) {
	rows, err := dbConn.SQL.Query("select row_number() over(order by circles desc), id, circles from users where role = 0")
	if err != nil {
		return entity.UserCirclesResponse{
			Circles: 0,
			Rate:    10000,
		}, err
	}
	defer rows.Close()

	var id, circles, rate int

	for rows.Next() {
		err := rows.Scan(&rate, &id, &circles)
		if err != nil {
			// log.Println(err)
			return entity.UserCirclesResponse{
				Circles: 0,
				Rate:    10000,
			}, err
		}
		if id == wantedId {
			return entity.UserCirclesResponse{
				Circles: circles,
				Rate:    rate,
			}, nil
		}
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return entity.UserCirclesResponse{
				Circles: 0,
				Rate:    10000,
			}, err
		}
		// fmt.Println("------------------------")
	}

	return entity.UserCirclesResponse{
		Circles: 0,
		Rate:    10000,
	}, errors.New("user is not a customer or can not be found")
}

func (driver *userDriver) DeleteUser(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from users where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be found")
	}

	return nil
}

func (driver *userDriver) AddUser(user entity.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO users(name, phone, password)
	VALUES ($1, $2, $3) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, user.Name, user.Phone, user.Password)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be added")
	}

	return nil
}

func (driver *userDriver) UserAddAddress(address entity.AddAddressRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO addresses(user_id, name, description, latitude, longitude)
	VALUES ($1, $2, $3, $4) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, address.UserId, address.Name, address.Descripition, address.Latitude, address.Longitude)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("address could not be added")
	}

	return nil
}

func (driver *userDriver) UserDeleteAddress(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from addresses where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("address could not be found")
	}

	return nil
}

func (driver *userDriver) LoginUser(userLoginInfo entity.UserLoginRequest) (entity.User, error) {
	// user, err := FindUser(userLoginInfo)
	query := `select id, name, phone, password, active, role, user_discount, area_id from users where phone = $1`
	var id, role, areaId int
	var name, phone, password string
	var active bool
	var userDiscount float32
	row := dbConn.SQL.QueryRow(query, userLoginInfo.Phone)
	err := row.Scan(&id, &name, &phone, &password, &active, &role, &userDiscount, &areaId)
	if id == 0 {
		return entity.User{
			ID:       0,
			Name:     "",
			Phone:    "",
			Password: "",
			Balance:  0,
			Active:   false,
			Circles:  0,
			Role:     0,
		}, errors.New("user phone does not exist")
	}
	if userLoginInfo.Password != password {
		return entity.User{
			ID:       0,
			Name:     "",
			Phone:    "",
			Password: "",
			Balance:  0,
			Active:   false,
			Circles:  0,
			Role:     0,
		}, errors.New("user password does not match")
	}
	if err != nil {
		return entity.User{
			ID:       0,
			Name:     "",
			Phone:    "",
			Password: "",
			Balance:  0,
			Active:   false,
			Circles:  0,
			Role:     0,
		}, err
	}
	area, err := ad.FindArea(areaId)
	if err != nil {
		// log.Println(err)
		return entity.User{}, err
	}
	user := entity.User{
		ID:           id,
		Name:         name,
		Phone:        phone,
		Password:     "",
		Active:       active,
		Role:         role,
		UserDiscount: userDiscount,
		Area:         area,
	}
	return user, nil
}

func (driver *userDriver) EditUser(userEditInfo entity.UserEditRequest) (entity.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := driver.GetEditUserStatementString(userEditInfo)
	// fmt.Println("Statement Is:", stmt)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return entity.User{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.User{}, errors.New("user could not be found")
	}
	user, err := driver.FindUser(userEditInfo.ID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (driver *userDriver) GetEditUserStatementString(userEditInfo entity.UserEditRequest) string {
	stmt := `UPDATE users SET `
	if userEditInfo.Name != "" {
		stmt = stmt + `name = '` + userEditInfo.Name + `', `
	}
	// if userEditInfo.Role != 0 {
	// 	stmt = stmt + `role = ` + fmt.Sprint(userEditInfo.Role) + `, `
	// }
	if userEditInfo.Password != "" {
		stmt = stmt + `password = '` + userEditInfo.Password + `', `
	}
	if userEditInfo.Balance != 0 {
		stmt = stmt + `balance = ` + fmt.Sprint(userEditInfo.Balance) + `, `
	}
	if userEditInfo.Area.ID != 0 {
		stmt = stmt + `area_id = ` + fmt.Sprint(userEditInfo.Area.ID) + `, `
	}
	if userEditInfo.Active {
		stmt = stmt + `active = true `
	} else {
		stmt = stmt + `active = false `
	}
	stmt = stmt + `where id = ` + fmt.Sprint(userEditInfo.ID) + ` RETURNING *`
	return stmt
}

func (driver *userDriver) EditUserDiscount(userInfo entity.User) (entity.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE users SET user_discount = $1 where id = $2 RETURNING *`
	// fmt.Println("Statement Is:", stmt)

	result, err := dbConn.SQL.ExecContext(ctx, stmt, userInfo.UserDiscount, userInfo.ID)
	if err != nil {
		return entity.User{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.User{}, errors.New("user could not be found")
	}
	user, err := driver.FindUser(userInfo.ID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (driver *userDriver) ActivateUser(userInfo entity.UserInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE users SET active = true WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, userInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be found")
	}
	return nil
}

func (driver *userDriver) DeactivateUser(userInfo entity.UserInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE users SET active = false WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, userInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be found")
	}
	return nil
}

func (driver *userDriver) SpecializeUser(userInfo entity.UserInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE users SET special_user = true WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, userInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be found")
	}
	return nil
}

func (driver *userDriver) NormalizeUser(userInfo entity.UserInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE users SET special_user = false WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, userInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be found")
	}
	return nil
}

func (driver *userDriver) ChangeUserRole(userInfo entity.UserChangeRoleRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE users SET role = $1 WHERE id = $2 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, userInfo.Role, userInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be found")
	}
	return nil
}

func (driver *userDriver) EditUserBalanceAndCircles(userId int, deliveryCost int, productsCost int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE users SET balance = balance + $1, circles = circles + $2 where id = $3 RETURNING id`
	balanceIncrease := int(math.Floor(float64(productsCost) / float64(1000)))
	circlesIncrease := int(math.Floor(float64(deliveryCost) / float64(1000)))

	result, err := dbConn.SQL.ExecContext(ctx, stmt, balanceIncrease, circlesIncrease, userId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user could not be found")
	}
	user, err := driver.FindUser(userId)
	if err != nil {
		return err
	}
	if !user.SpecialUser {
		rateAndCircles, _ := driver.FindUserRateAndCircles(userId)
		// fmt.Println("User ID:", userId)
		// fmt.Println("Rate and Circles:", rateAndCircles)
		if rateAndCircles.Rate == 1 {
			driver.EditUserDiscount(entity.User{ID: userId, UserDiscount: 0.75})
		} else if rateAndCircles.Rate == 2 {
			driver.EditUserDiscount(entity.User{ID: userId, UserDiscount: 0.5})
		} else if rateAndCircles.Rate == 3 {
			driver.EditUserDiscount(entity.User{ID: userId, UserDiscount: 0.25})
		} else {
			driver.EditUserDiscount(entity.User{ID: userId, UserDiscount: 0.0})
		}
	}
	return nil
}

func (driver *userDriver) EditDeliveryWorkerBalanceAndCircles(userId int, deliveryCost int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	deliveryWorker, err := driver.FindUser(userId)
	if err != nil {
		return err
	}

	if deliveryWorker.Role == 1 {
		// time.Now().Day() == 1 && && deliveryWorker.Circles == 0 && deliveryWorker.Balance == 0
		// balanceIncrease := int(math.Floor(float64(deliveryCost) / float64(1000)))
		stmt := `UPDATE users SET balance = balance + $1, circles = circles + 1 where id = $2 RETURNING id`
		result, err := dbConn.SQL.ExecContext(ctx, stmt, deliveryCost, userId)
		if err != nil {
			return err
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return errors.New("user could not be found")
		}
		_, err = driver.FindUser(userId)
		if err != nil {
			return err
		}
	}
	/* else {
		stmt := `UPDATE users SET balance = $1, circles = 1 where id = $3 RETURNING id`
		result, err := dbConn.SQL.ExecContext(ctx, stmt, deliveryCost, userId)
		if err != nil {
			return err
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return errors.New("user could not be found")
		}
		_, err = driver.FindUser(userId)
		if err != nil {
			return err
		}
	} */

	return nil
}
