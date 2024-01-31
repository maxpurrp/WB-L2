package main

import (
	"db"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"model"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо
	работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event
			POST /update_event
			POST /delete_event
			GET /events_for_day
			GET /events_for_week
			GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."}
	в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки
		входных данных (невалидный int например) сервер должен возвращать HTTP 400.
		 В случае остальных ошибок сервер должен возвращать HTTP 500.
		Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func dateValidate(date string) error {
	// parse the date string using the "2006-01-02" format
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		// return an error if parsing fails, indicating an invalid date format
		return fmt.Errorf("invalid date format")
	}

	// get the current date and time
	currentDate := time.Now()

	// check if the parsed date is before the current date
	if parsedDate.Before(currentDate) {
		// return an error if the parsed date is in the past
		return fmt.Errorf("it is not possible to create an event in the past")
	}

	// return nil if the date validation is successful
	return nil
}

func decodeEvent(data *http.Request) (model.Event, error) {
	// initialize an empty model.Event
	event := model.Event{}

	// create a JSON decoder for the request body
	decoder := json.NewDecoder(data.Body)

	// decode the JSON data into the model.Event struct
	err := decoder.Decode(&event)
	if err != nil {
		// return an error if decoding fails, indicating an invalid request body
		return model.Event{}, fmt.Errorf("invalid request body")
	}

	// validate that 'id' and 'date' are present and non-empty in the decoded event
	if event.ID == 0 || event.Date == "" {
		// return an error if validation fails, indicating missing or empty required fields
		return model.Event{}, fmt.Errorf("'id' and 'date' are expected in the required fields")
	}

	// return the decoded event and nil error if successful
	return event, nil
}

func getQueryParams(r *http.Request) (model.Event, error) {
	// retrieve the "id" and "date" parameters from the URL query
	id := r.URL.Query().Get("id")
	date := r.URL.Query().Get("date")

	// convert the "id" parameter to an integer
	eventID, err := strconv.Atoi(id)
	if err != nil {
		// return an error if conversion fails, indicating an invalid ID format
		return model.Event{}, fmt.Errorf("invalid ID format: %v", err)
	}

	// create a model.Event instance with the extracted values
	event := model.Event{
		ID:   eventID,
		Date: date,
	}

	// return the created model.Event and nil error
	return event, nil
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	// set the Content-Type header to indicate a JSON response
	w.Header().Set("Content-Type", "application/json")
	// set the HTTP status code for the response
	w.WriteHeader(statusCode)

	// encode the provided data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		// log an error if encoding fails
		log.Printf("Error encoding to JSON: %v", err)
		return
	}

	// write the JSON data to the HTTP response
	_, err = w.Write(jsonData)
	if err != nil {
		// log an error if writing to the response fails
		log.Printf("Error writing JSON to response: %v", err)
	}
}

func handleEvents(w http.ResponseWriter, r *http.Request, eventType string) {
	log.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)
	// validate that the HTTP method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// extract query parameters from the request
	event, err := getQueryParams(r)
	if err != nil {
		// respond with a JSON-encoded error message if parameter extraction fails
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// retrieve events from the database based on the specified period
	events, err := db.GetEvents(eventType, event.ID, event.Date)
	if err != nil {
		// respond with a JSON-encoded error message if database retrieval fails
		sendJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// create a response object containing the retrieved events
	resp := model.Response{Events: events}
	// send a JSON response with the retrieved events and a status code of 200 (OK)
	sendJSONResponse(w, http.StatusOK, resp)
}

func handleEventAction(w http.ResponseWriter, r *http.Request, action string, successMessage string) {
	log.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)
	// validate that the HTTP method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// decode event data from the request body
	event, err := decodeEvent(r)
	if err != nil {
		// respond with a JSON-encoded error message if decoding fails
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// validate the date of the event
	err = dateValidate(event.Date)
	if err != nil {
		// Respond with a JSON-encoded error message if date validation fails
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// execute the specified database action (create, update, or delete)
	db.DbAction(action, event.ID, event.Date)

	// send a JSON response with a success message and a status code of 200 (OK)
	sendJSONResponse(w, http.StatusOK, map[string]string{"result": successMessage})
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	// delegate the handling of the create action to a common function
	handleEventAction(w, r, "create", "Event successfully created")
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	// delegate the handling of the update action to a common function
	handleEventAction(w, r, "update", "Event successfully updated")
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	// delegate the handling of the delete action to a common function
	handleEventAction(w, r, "delete", "Event successfully deleted")
}

func dayEvents(w http.ResponseWriter, r *http.Request) {
	// delegate the handling of the day period to a common function
	handleEvents(w, r, "day")
}

func weekEvents(w http.ResponseWriter, r *http.Request) {
	// delegate the handling of the week period to a common function
	handleEvents(w, r, "week")
}

func monthEvents(w http.ResponseWriter, r *http.Request) {
	// delegate the handling of the month period to a common function
	handleEvents(w, r, "month")
}

func readConfig() (string, error) {
	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	var config model.Config
	if err := viper.ReadInConfig(); err != nil {
		return config.Port, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config.Port, err
	}

	return config.Port, nil
}

func main() {
	// create a new ServeMux to handle HTTP requests
	mux := http.NewServeMux()
	readConfig()

	// define routes and their corresponding handlers
	mux.HandleFunc("/create_event", createEvent)
	mux.HandleFunc("/update_event", updateEvent)
	mux.HandleFunc("/delete_event", deleteEvent)
	mux.HandleFunc("/events_for_day", dayEvents)
	mux.HandleFunc("/events_for_week", weekEvents)
	mux.HandleFunc("/events_for_month", monthEvents)

	// start the HTTP server using the defined ServeMux

	port, err := readConfig()
	if err != nil {
		fmt.Println("cant read config file, use default (8000)")
		port = "8000"
	}
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
