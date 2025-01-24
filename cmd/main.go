package main

import (
	"database/sql"
	"database_practice/models"
	"database_practice/postgres"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var guests []models.Guest

var db *sql.DB

func createGuestHandler(w http.ResponseWriter, r *http.Request) {
	var guest models.Guest
	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	guests = append(guests, guest)
	fmt.Fprintf(w, "post new guest: '%v'\n", guest)

	id, err := postgres.CreateGuest(db, guest.First_name, guest.Second_name, guest.Current_room_num)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(id)
}

func checkInHandler(w http.ResponseWriter, r *http.Request) {
	var guest models.Guest

	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := postgres.CheckInGuest(db, guest.Current_room_num, guest.First_name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to check in guest: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking rows affected: %v", err), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No guest found to check in", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Guest '%s' checked in to room %d\n", guest.First_name, guest.Current_room_num)
}

func checkOutHandler(w http.ResponseWriter, r *http.Request) {
	var guest models.Guest

	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := postgres.CheckOutGuest(db, guest.Current_room_num)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to check out guest: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking rows affected: %v", err), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No guest found to check out", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Guest '%s' checked out from room %d\n", guest.First_name, guest.Current_room_num)
}

func deleteGuestHandler(w http.ResponseWriter, r *http.Request) {
	var guest models.Guest

	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := postgres.DeleteGuest(db, guest.First_name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete guest: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking rows affected: %v", err), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No guest found to delte", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Guest '%s' deleted %d\n", guest.First_name, guest.Current_room_num)
}

func changeRoomHandler(w http.ResponseWriter, r *http.Request) {
	var guest models.Guest

	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := postgres.ChangeGuestRoom(db, guest.Current_room_num, guest.First_name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to change guest's room: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking rows affected: %v", err), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No guest's room found to be changed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Room '%d' was assigned to guest %s\n", guest.Current_room_num, guest.First_name)
}

func main() {

	// ПОДКЛЮЧЕНИЕ К БД
	connStr := "user=postgres password=postgres dbname=hotel_manager sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// АДРЕСА ЗАПРОСОВ НА СЕРВЕР
	http.HandleFunc("/guest/create-guest", createGuestHandler)
	http.HandleFunc("/guest/check-in", checkInHandler)
	http.HandleFunc("/guest/check-out", checkOutHandler)
	http.HandleFunc("/guest/delete-guest", deleteGuestHandler)
	http.HandleFunc("/guest/change-room", changeRoomHandler)

	// Начало обработки запросов
	log.Println("server start listening on port 8080")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("s02432432342")

}

// как развернуть веб-сервер
// написать методы и протестить в main()
// поднять http сервис, на каждый метод сделать обработчик который будет создавать записи в базе через постман
