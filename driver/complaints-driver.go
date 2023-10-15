package driver

import (
	"context"
	"errors"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type ComplaintDriver interface {
	FindAllComplaints() ([]entity.Complaint, error)
	FindAboutDeliveryComplaints() ([]entity.Complaint, error)
	FindAboutTheAppComplaints() ([]entity.Complaint, error)
	FindImprovementSuggestionComplaints() ([]entity.Complaint, error)
	FindOtherReasonComplaints() ([]entity.Complaint, error)

	AddComplaint(complaint entity.Complaint) error
	AddAboutDelivery(complaint entity.Complaint) error
	AddAboutTheApp(complaint entity.Complaint) error
	AddImprovementSuggestion(complaint entity.Complaint) error
	AddOtherReason(complaint entity.Complaint) error
	DeleteComplaint(wantedId int) error
}

type complaintDriver struct {
}

func NewComplaintDriver() ComplaintDriver {
	return &complaintDriver{}
}

func (driver *complaintDriver) FindAllComplaints() ([]entity.Complaint, error) {
	complaints := make([]entity.Complaint, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, text, date, about_delivery, 
		about_the_app, improvement_suggestion, other_reason 
	from complaints`)
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	defer rows.Close()

	var id, userId int
	var text string
	var date time.Time
	var aboutDelivery, aboutTheApp, improvementSuggestion, otherReason bool

	for rows.Next() {
		err := rows.Scan(&id, &userId, &text, &date, &aboutDelivery, &aboutTheApp, &improvementSuggestion, &otherReason)
		if err != nil {
			return make([]entity.Complaint, 0), err
		}
		complaints = append(complaints, entity.Complaint{
			ID:                    id,
			UserID:                userId,
			Text:                  text,
			Date:                  date,
			AboutDelivery:         aboutDelivery,
			AboutTheApp:           aboutTheApp,
			ImprovementSuggestion: improvementSuggestion,
			OtherReason:           otherReason,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Complaint, 0), err
		}
	}

	// driver.cacheDetails = complaints
	return complaints, nil
}

func (driver *complaintDriver) FindAboutDeliveryComplaints() ([]entity.Complaint, error) {
	complaints := make([]entity.Complaint, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, text, date, about_delivery, 
		about_the_app, improvement_suggestion, other_reason 
	from complaints where about_delivery = true`)
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	defer rows.Close()

	var id, userId int
	var text string
	var date time.Time
	var aboutDelivery, aboutTheApp, improvementSuggestion, otherReason bool

	for rows.Next() {
		err := rows.Scan(&id, &userId, &text, &date, &aboutDelivery, &aboutTheApp, &improvementSuggestion, &otherReason)
		if err != nil {
			return make([]entity.Complaint, 0), err
		}
		complaints = append(complaints, entity.Complaint{
			ID:                    id,
			UserID:                userId,
			Text:                  text,
			Date:                  date,
			AboutDelivery:         aboutDelivery,
			AboutTheApp:           aboutTheApp,
			ImprovementSuggestion: improvementSuggestion,
			OtherReason:           otherReason,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Complaint, 0), err
		}
	}

	// driver.cacheDetails = complaints
	return complaints, nil
}

func (driver *complaintDriver) FindAboutTheAppComplaints() ([]entity.Complaint, error) {
	complaints := make([]entity.Complaint, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, text, date, about_delivery, 
		about_the_app, improvement_suggestion, other_reason 
	from complaints where about_the_app = true`)
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	defer rows.Close()

	var id, userId int
	var text string
	var date time.Time
	var aboutDelivery, aboutTheApp, improvementSuggestion, otherReason bool

	for rows.Next() {
		err := rows.Scan(&id, &userId, &text, &date, &aboutDelivery, &aboutTheApp, &improvementSuggestion, &otherReason)
		if err != nil {
			return make([]entity.Complaint, 0), err
		}
		complaints = append(complaints, entity.Complaint{
			ID:                    id,
			UserID:                userId,
			Text:                  text,
			Date:                  date,
			AboutDelivery:         aboutDelivery,
			AboutTheApp:           aboutTheApp,
			ImprovementSuggestion: improvementSuggestion,
			OtherReason:           otherReason,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Complaint, 0), err
		}
	}

	// driver.cacheDetails = complaints
	return complaints, nil
}

func (driver *complaintDriver) FindImprovementSuggestionComplaints() ([]entity.Complaint, error) {
	complaints := make([]entity.Complaint, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, text, date, about_delivery, 
		about_the_app, improvement_suggestion, other_reason 
	from complaints where improvement_suggestion = true`)
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	defer rows.Close()

	var id, userId int
	var text string
	var date time.Time
	var aboutDelivery, aboutTheApp, improvementSuggestion, otherReason bool

	for rows.Next() {
		err := rows.Scan(&id, &userId, &text, &date, &aboutDelivery, &aboutTheApp, &improvementSuggestion, &otherReason)
		if err != nil {
			return make([]entity.Complaint, 0), err
		}
		complaints = append(complaints, entity.Complaint{
			ID:                    id,
			UserID:                userId,
			Text:                  text,
			Date:                  date,
			AboutDelivery:         aboutDelivery,
			AboutTheApp:           aboutTheApp,
			ImprovementSuggestion: improvementSuggestion,
			OtherReason:           otherReason,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Complaint, 0), err
		}
	}

	// driver.cacheDetails = complaints
	return complaints, nil
}

func (driver *complaintDriver) FindOtherReasonComplaints() ([]entity.Complaint, error) {
	complaints := make([]entity.Complaint, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, text, date, about_delivery, 
		about_the_app, improvement_suggestion, other_reason 
	from complaints where other_reason = true`)
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	defer rows.Close()

	var id, userId int
	var text string
	var date time.Time
	var aboutDelivery, aboutTheApp, improvementSuggestion, otherReason bool

	for rows.Next() {
		err := rows.Scan(&id, &userId, &text, &date, &aboutDelivery, &aboutTheApp, &improvementSuggestion, &otherReason)
		if err != nil {
			return make([]entity.Complaint, 0), err
		}
		complaints = append(complaints, entity.Complaint{
			ID:                    id,
			UserID:                userId,
			Text:                  text,
			Date:                  date,
			AboutDelivery:         aboutDelivery,
			AboutTheApp:           aboutTheApp,
			ImprovementSuggestion: improvementSuggestion,
			OtherReason:           otherReason,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Complaint, 0), err
		}
	}

	// driver.cacheDetails = complaints
	return complaints, nil
}

func (driver *complaintDriver) AddComplaint(complaint entity.Complaint) error {

	if complaint.AboutDelivery {
		err := driver.AddAboutDelivery(complaint)
		if err != nil {
			return err
		}
	} else if complaint.AboutTheApp {
		err := driver.AddAboutTheApp(complaint)
		if err != nil {
			return err
		}
	} else if complaint.ImprovementSuggestion {
		err := driver.AddImprovementSuggestion(complaint)
		if err != nil {
			return err
		}
	} else if complaint.OtherReason {
		err := driver.AddOtherReason(complaint)
		if err != nil {
			return err
		}
	} else {
		return errors.New("complaint type should be specified")
	}

	// defer driver.FindAllDetails()

	return nil
}

func (driver *complaintDriver) AddAboutDelivery(complaint entity.Complaint) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO complaints(user_id, text, date, about_delivery)
	VALUES ($1, $2, $3, $4) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, complaint.UserID, complaint.Text, complaint.Date, true)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("about delivery complaint could not be added")
	}

	return nil
}

func (driver *complaintDriver) AddAboutTheApp(complaint entity.Complaint) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO complaints(user_id, text, date, about_the_app)
	VALUES ($1, $2, $3, $4) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, complaint.UserID, complaint.Text, complaint.Date, true)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("about the app complaint could not be added")
	}

	return nil
}

func (driver *complaintDriver) AddImprovementSuggestion(complaint entity.Complaint) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO complaints(user_id, text, date, improvement_suggestion)
	VALUES ($1, $2, $3, $4) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, complaint.UserID, complaint.Text, complaint.Date, true)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("improvement suggestion complaint could not be added")
	}

	return nil
}

func (driver *complaintDriver) AddOtherReason(complaint entity.Complaint) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO complaints(user_id, text, date, other_reason)
	VALUES ($1, $2, $3, $4) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, complaint.UserID, complaint.Text, complaint.Date, true)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("other reason complaint could not be added")
	}

	return nil
}

func (driver *complaintDriver) DeleteComplaint(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from complaints where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("complaint could not be found")
	}
	return nil
}
