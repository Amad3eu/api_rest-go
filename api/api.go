package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkbhowmick/go-rest-api/auth"
	"github.com/pkbhowmick/go-rest-api/model"
)

var users map[string]model.User

var router = mux.NewRouter()

func InitializeDB() {
	users = make(map[string]model.User)
	users["1"] = model.User{
		"1",
		"Pulak",
		"Kanti",
		[]model.Repository{
			{
				"1001",
				"go-rest-api",
				"public",
				1,
			},
		},
		time.Now(),
	}
	users["2"] = model.User{
		"2",
		"Mehedi",
		"Hasan",
		[]model.Repository{
			{
				"1002",
				"go-api-server",
				"public",
				2,
			},
		},
		time.Now(),
	}
	users["3"] = model.User{
		"3",
		"Prangon",
		"Majumdar",
		[]model.Repository{
			{
				"1003",
				"go-http-api-server",
				"private",
				3,
			},
		},
		time.Now(),
	}
	users["4"] = model.User{
		"4",
		"Sakib",
		"Alamin",
		[]model.Repository{
			{
				"1004",
				"go-httpapi-server",
				"private",
				5,
			},
		},
		time.Now(),
	}
	users["5"] = model.User{
		"5",
		"Sahadat",
		"Sahin",
		[]model.Repository{
			{
				"1005",
				"go-http-server",
				"public",
				5,
			},
		},
		time.Now(),
	}
}

func userToArray() []model.User {
	items := make([]model.User, 0)
	for _, item := range users {
		items = append(items, item)
	}
	return items
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	allUsers := userToArray()
	err := json.NewEncoder(res).Encode(allUsers)
	res.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	if user, ok := users[id]; ok {
		res.WriteHeader(http.StatusOK)
		err := json.NewEncoder(res).Encode(user)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	errMsg := "User with id " + id + " doesn't exist"
	http.Error(res, errMsg, http.StatusNotFound)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(res, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if user.ID == "" || user.FirstName == "" || user.LastName == "" {
		http.Error(res, "Missing required fields", http.StatusBadRequest)
		return
	}
	if _, ok := users[user.ID]; ok {
		http.Error(res, "User with given ID already exist", http.StatusConflict)
		return
	}
	user.CreatedAt = time.Now()
	users[user.ID] = user
	res.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(res).Encode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(res, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	var newUser, oldUser model.User
	oldUser, ok := users[id]
	if !ok {
		http.Error(res, "User doesn't exist", http.StatusNotFound)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.ID = oldUser.ID
	newUser.CreatedAt = oldUser.CreatedAt
	users[oldUser.ID] = newUser
	err = json.NewEncoder(res).Encode(&newUser)
	res.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	if user, ok := users[id]; ok {
		delete(users, id)
		err := json.NewEncoder(res).Encode(&user)
		res.WriteHeader(http.StatusOK)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(res, "User doesn't exist", http.StatusNotFound)
}

func Homepage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"status" : "OK"}`))
}

func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	token, err := auth.GenerateToken("testuser")
	if err != nil {
		http.Error(res, "Wrong username or password!", http.StatusUnauthorized)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"token" : "` + token + `"}`))
}

var port string = "8080"
var authStatus bool = true

func SetFlags(serverPort string, auth bool) {
	port = serverPort
	authStatus = auth
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s%s %s", req.Method, req.Host, req.URL, req.Proto)
		next.ServeHTTP(res, req)
	})
}

func Init() {
	InitializeDB()
	router.Use(Logger)
	if authStatus {
		router.Use(auth.Authentication)
	}
	router.HandleFunc("/", Homepage).Methods("GET")
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/login", Login).Methods("POST")
}

func StartServer() {
	Init()
	log.Println("Server is listening on port " + port)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
