package main

import (
	"encoding/json"
	"fmt"
	"go_cast/S11P01-game/repository/mysql"
	"go_cast/S11P01-game/service/userservice"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/login", userLoginHandler)
	http.HandleFunc("/health-check", healthCheckHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
	/*
		Second method
		mux := http.NewServeMux()
		mux.HandleFunc("/users/register", userRegisterHandler)

		sever := http.Server{
			Addr:    ":8080",
			Handler: mux,
		}

		sever.ListenAndServe()
	*/
}

func userRegisterHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "Invalid Method: %s", req.Method)
		return
	}

	// Handle POST request
	// res.WriteHeader(http.StatusOK)
	// fmt.Fprint(res, "Registration endpoint reached")
	data, err := io.ReadAll(req.Body)
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	fmt.Printf("Data: %s\n", data)

	var requestBody userservice.RegisterRequest
	err = json.Unmarshal(data, &requestBody)
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, fmt.Errorf("your json structure is not standard"))))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	response, err := userSvc.Register(requestBody)
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(fmt.Sprintf(`{"user": "%v"}`, response.User)))

	return
}

func userLoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "Invalid Method: %s", req.Method)
		return
	}

	// Handle POST request
	// res.WriteHeader(http.StatusOK)
	// fmt.Fprint(res, "Registration endpoint reached")
	data, err := io.ReadAll(req.Body)
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	var requestBody userservice.LoginRequest
	err = json.Unmarshal(data, &requestBody)
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, fmt.Errorf("your json structure is not standard"))))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	response, err := userSvc.Login(requestBody)
	if err != nil {
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write([]byte(fmt.Sprintf(`{"user": "%v"}`, response.User)))

	return
}

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Invalid Method", http.StatusMethodNotAllowed)
		return
	} else {
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, "OK")
		return
	}
}
