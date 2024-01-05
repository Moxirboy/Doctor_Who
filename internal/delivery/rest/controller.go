package rest

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/pkg/Bot"
	"DoctorWho/internal/pkg/jwt"
	"DoctorWho/internal/pkg/session"
	"DoctorWho/internal/pkg/sms"
	"DoctorWho/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type controller struct {
	usecase usecase.Usecase
	bot     Bot.Bot
}

func NewController(usecase usecase.Usecase) *controller {
	return &controller{usecase: usecase}
}

func (c controller) SignUp(w http.ResponseWriter, r *http.Request) {
	sessions, _ := session.Store.Get(r, "User")
	var NewUser domain.NewUser
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		c.bot.SendErrorNotification(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Invalid json")
	}
	log.Println(NewUser.PhoneNumber)
	id, err := c.usecase.RegisterUser(&NewUser)
	if err != nil {
		c.bot.SendErrorNotification(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Could`nt register")
	}
	sessions.Values["userId"] = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (c controller) Login(w http.ResponseWriter, r *http.Request) {
	sessions, _ := session.Store.Get(r, "User")
	var User dto.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		c.bot.SendErrorNotification(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Invalid json")
	}
	exist, id, err := c.usecase.Login(User.Email)
	if err != nil {
		c.bot.SendErrorNotification(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Could`nt login")
	}
	if exist {
		code := sms.GenerateVerificationCode()
		c.bot.SendNotification(code)
		err = sms.SendEmail(User.Email, code)
		c.bot.SendNotification(User.Email)
		if err != nil {
			c.bot.SendErrorNotification(err)
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
		}
		sessions.Values["userId"] = id

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode("verification code sent")
		if err != nil {
			c.bot.SendErrorNotification(err)
			return
		}
	} else {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}

}
func (c controller) Verification(w http.ResponseWriter, r *http.Request) {

	sessions, _ := session.Store.Get(r, "User")

	var message domain.Sms
	err := json.NewDecoder(r.Body).Decode(&message)
	id := sessions.Values["userId"]
	if id == nil {
		c.bot.SendErrorNotification(err)
		id = message.UserId
		c.bot.SendNotification(id.(string))
	}
	if err != nil {
		c.bot.SendErrorNotification(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Invalid json")
	}
	match, err := c.usecase.Verify(id.(string), message.Code)
	if err != nil {
		c.bot.SendErrorNotification(err)
		http.Error(w, "Not matched", http.StatusUnauthorized)
	}
	if match {
		token, err := jwt.CreateToken(id.(string))
		if err != nil {
			c.bot.SendErrorNotification(err)
			http.Error(w, "error occurred: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"access_token": token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
func (c controller) Logout(w http.ResponseWriter, r *http.Request) {
	sessions, _ := session.Store.Get(r, "User")
	sessions.Values["userId"] = nil
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("success")
}
func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	user := c.usecase.GetAll()
	log.Println(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}
