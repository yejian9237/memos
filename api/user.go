package api

import (
	"encoding/json"
	"memos/api/e"
	"memos/store"
	"net/http"

	"github.com/gorilla/mux"
)

func handleGetMyUserInfo(w http.ResponseWriter, r *http.Request) {
	userId, _ := GetUserIdInSession(r)

	user, err := store.GetUserById(userId)

	if err != nil {
		e.ErrorHandler(w, "USER_NOT_FOUND", err.Error())
		return
	}

	json.NewEncoder(w).Encode(Response{
		Succeed: true,
		Message: "",
		Data:    user,
	})
}

func handleUpdateMyUserInfo(w http.ResponseWriter, r *http.Request) {
	userId, _ := GetUserIdInSession(r)

	updateUserPatch := store.UpdateUserPatch{}
	err := json.NewDecoder(r.Body).Decode(&updateUserPatch)

	if err != nil {
		e.ErrorHandler(w, "REQUEST_BODY_ERROR", "Bad request")
		return
	}

	if updateUserPatch.Username != nil {
		usernameUsable, _ := store.CheckUsernameUsable(*updateUserPatch.Username)
		if !usernameUsable {
			json.NewEncoder(w).Encode(Response{
				Succeed: false,
				Message: "Username is existed",
				Data:    nil,
			})
			return
		}
	}

	user, err := store.UpdateUser(userId, &updateUserPatch)

	if err != nil {
		e.ErrorHandler(w, "DATABASE_ERROR", err.Error())
		return
	}

	json.NewEncoder(w).Encode(Response{
		Succeed: true,
		Message: "",
		Data:    user,
	})
}

func handleResetUserOpenId(w http.ResponseWriter, r *http.Request) {
	userId, _ := GetUserIdInSession(r)

	openId, err := store.ResetUserOpenId(userId)

	if err != nil {
		e.ErrorHandler(w, "DATABASE_ERROR", err.Error())
		return
	}

	json.NewEncoder(w).Encode(Response{
		Succeed: true,
		Message: "",
		Data:    openId,
	})
}

func handleCheckUsername(w http.ResponseWriter, r *http.Request) {
	type CheckUsernameDataBody struct {
		Username string
	}

	checkUsername := CheckUsernameDataBody{}
	err := json.NewDecoder(r.Body).Decode(&checkUsername)

	if err != nil {
		e.ErrorHandler(w, "REQUEST_BODY_ERROR", "Bad request")
		return
	}

	usable, err := store.CheckUsernameUsable(checkUsername.Username)

	if err != nil {
		e.ErrorHandler(w, "DATABASE_ERROR", err.Error())
		return
	}

	json.NewEncoder(w).Encode(Response{
		Succeed: true,
		Message: "",
		Data:    usable,
	})
}

func handleValidPassword(w http.ResponseWriter, r *http.Request) {
	type ValidPasswordDataBody struct {
		Password string
	}

	userId, _ := GetUserIdInSession(r)
	validPassword := ValidPasswordDataBody{}
	err := json.NewDecoder(r.Body).Decode(&validPassword)

	if err != nil {
		e.ErrorHandler(w, "REQUEST_BODY_ERROR", "Bad request")
		return
	}

	valid, err := store.CheckPasswordValid(userId, validPassword.Password)

	if err != nil {
		e.ErrorHandler(w, "DATABASE_ERROR", err.Error())
		return
	}

	json.NewEncoder(w).Encode(Response{
		Succeed: true,
		Message: "",
		Data:    valid,
	})
}

func RegisterUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/api/user").Subrouter()

	userRouter.Use(JSONResponseMiddleWare)
	userRouter.Use(AuthCheckerMiddleWare)

	userRouter.HandleFunc("/me", handleGetMyUserInfo).Methods("GET")
	userRouter.HandleFunc("/me", handleUpdateMyUserInfo).Methods("PATCH")
	userRouter.HandleFunc("/open_id/new", handleResetUserOpenId).Methods("POST")
	userRouter.HandleFunc("/checkusername", handleCheckUsername).Methods("POST")
	userRouter.HandleFunc("/validpassword", handleValidPassword).Methods("POST")
}
