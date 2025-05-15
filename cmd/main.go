package main

import (
	"encoding/json"
	"fmt"
	"go_cast/S11P01-game/repository/mysql"
	"go_cast/S11P01-game/service/userservice"
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/login", userLoginHandler)
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/users/profile", getProfileHandler)

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

func getProfileHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Invalid Method", http.StatusMethodNotAllowed)
		return
	} else {
		// I have to get UserID from URL query
		userID := req.URL.Query().Get("user_id")
		if userID == "" {
			res.Write([]byte(`{"error": "user_id is needed"}`))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		mysqlRepo := mysql.New()
		userSvc := userservice.New(mysqlRepo)
		id, err := strconv.Atoi(userID)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "Invalid user_id parameter: %s", userID)
			return
		}
		reqUser := userservice.ProfileRequest{
			UserID: uint(id),
		}
		response, err := userSvc.GetProfile(reqUser)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
		data, err := json.Marshal(response)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Write(data)
		return
	}
}
