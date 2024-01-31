package db

import (
	"database/sql"
	"fmt"
	"model"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var (
	connStr = "postgresql://max:1234@localhost:5432/WB?sslmode=disable"
)

func openConn() *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func DbAction(action string, id int, date string) error {
	// open a database connection
	db := openConn()
	defer db.Close()

	// execute the appropriate SQL statement based on the specified action
	switch action {
	case "create":
		// insert a new event into the Events table
		_, err := db.Exec("INSERT INTO Events (id, event_date) VALUES ($1, $2)", id, date)
		if err != nil {
			return err
		}
	case "update":
		// update the event date in the Events table based on the user ID
		_, err := db.Exec("UPDATE Events SET event_date = $1 WHERE id = $2", date, id)
		if err != nil {
			return err
		}
	case "delete":
		// delete an event from the Events table based on the user ID
		_, err := db.Exec("DELETE FROM Events WHERE id = $1", id)
		if err != nil {
			return err
		}
	}

	// return nil if the database action is successful
	return nil
}

func GetEvents(period string, id int, date string) ([]model.Event, error) {
	// Initialize an empty slice to store events
	var eventStore []model.Event

	// open a database connection
	db := openConn()
	defer db.Close()

	var rows *sql.Rows
	var err error
	params := strings.Split(date, "-")
	year, month, day := params[0], params[1], params[2]

	// construct the date string in the "2006-01-02" format
	formattedDate := fmt.Sprintf("%s-%s-%s", year, month, day)

	// parse the formatted date string
	cpecDate, _ := time.Parse("2006-01-02", formattedDate)

	// execute SQL query based on the specified period
	switch period {
	case "day":
		rows, err = db.Query("SELECT * FROM Events WHERE event_date::date = $1::date", date)
	case "week":
		// use temporary values for the start and end of the week
		startOfWeek := cpecDate.AddDate(0, 0, -int(cpecDate.Weekday()))
		endOfWeek := startOfWeek.AddDate(0, 0, 6)

		rows, err = db.Query("SELECT * FROM Events WHERE event_date::date BETWEEN $1::date AND $2::date", startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02"))
	case "month":
		// use temporary values for the start and end of the month
		startOfMonth := time.Date(cpecDate.Year(), cpecDate.Month(), 1, 0, 0, 0, 0, time.UTC)
		endOfMonth := startOfMonth.AddDate(0, 1, -1)

		rows, err = db.Query("SELECT * FROM Events WHERE event_date::date BETWEEN $1::date AND $2::date", startOfMonth.Format("2006-01-02"), endOfMonth.Format("2006-01-02"))
	}

	// check for errors during the SQL query execution
	if err != nil {
		return nil, err
	}

	// iterate over the query result and populate the eventStore slice
	for rows.Next() {
		var event model.Event
		err := rows.Scan(&event.ID, &event.Date)
		if err != nil {
			return nil, err
		}
		eventStore = append(eventStore, event)
	}

	// return the populated eventStore slice and nil error if successful
	return eventStore, nil
}
