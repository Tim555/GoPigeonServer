package server

import (
	"fmt"
	"gopigeon/internal/controllers"
	"gopigeon/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func StartServer() {

	fmt.Println("setting up the server")

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	fmt.Println("Setting up routes.")

	router := mux.NewRouter()

	router.HandleFunc("/", rootPage)
	router.HandleFunc("/chat", chatHandler)

	router.HandleFunc("/users/", usersHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/user/{username}", userHandler)

	http.Handle("/", router)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Go Pidgeon Chat Server.")
}

func chatHandler(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected: ")

	err = ws.WriteMessage(1, []byte("Hi Client!"))

	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	log.Println(r.URL)
	log.Println(username)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	db, _ := controllers.GetDB()
	user := controllers.GetUser(username, db)
	log.Println(user)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "User: %s\n", user.Username)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	db, _ := controllers.GetDB()

	if r.Method == http.MethodGet {
		users := controllers.GetUsers(db)
		fmt.Fprintf(w, "Users: %v\n", users)
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		log.Println(username)
		controllers.CreateUser(&models.User{Username: username, Password: password}, db)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User %s created\n", username)
	}
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
