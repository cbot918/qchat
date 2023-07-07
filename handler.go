package main

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/cbot918/liby/jwty"
)

type Handler struct {
	Svc Service
}

func NewHandler(s Storage) *Handler {
	svc := NewService(s)
	return &Handler{
		Svc: svc,
	}
}

type LoginParam struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Token   string `json:"token"`
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// set allow origin

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	user := &LoginParam{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseError(w, http.StatusBadRequest, "request body wrong")
		return
	}

	result, err := h.Svc.LoginService(user)
	if err != nil {
		responseError(w, http.StatusBadRequest, result)
		return
	}

	// get user name from email
	pattern := `^([A-Za-z0-9._%+-]+)@`
	re := regexp.MustCompile(pattern)
	name := re.FindStringSubmatch(user.Email)

	res := &LoginResponse{
		Message: "auth successful",
		Name:    name[1],
		Token:   result,
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
}

type ListFriendParam struct {
}

type ListFriendResponse struct {
	Message string `json:"message"`
}

func (h *Handler) ListFriend(w http.ResponseWriter, r *http.Request) {

	j := jwty.New()
	token, err := j.FastJwt(1, "yale918@gmail.com")
	if err != nil {
		log("create jwt token failed")
		return
	}
	log("token: " + token)

	clams := j.DecodeJwt(token)
	log("clams: ", clams)
	log(clams.Email)
	log(clams.Id)

	resFriends, err := h.Svc.ListFriendService()
	if err != nil {
		log("listfriendservice failed")
		responseError(w, 400, err.Error())
		return
	}

	log(resFriends)

	responseOk(w, ListFriendResponse{
		Message: "test",
	})

}

func responseOk(w http.ResponseWriter, a any) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(a)
	w.Write(jsonData)
}

func responseError(w http.ResponseWriter, code int, message string) {
	// w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	w.Write(jsonData)
}
