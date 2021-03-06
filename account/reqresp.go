package account

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	CreateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
	GetUserRequest struct {
		ID string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(c context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(c context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)
	req = GetUserRequest{ID: vars["id"]}
	return req, nil
}
