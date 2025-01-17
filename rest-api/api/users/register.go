package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/lucfek/go-exercises/rest-api/model"
	"github.com/lucfek/go-exercises/rest-api/response"
)

// Register registers a user
func (h Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := userData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}

	if len(strings.TrimSpace(data.Email)) == 0 || len(strings.TrimSpace(data.Password)) == 0 {
		res := response.Resp{
			Status: "error",
			Data:   "Empty values",
		}
		response.Writer(w, res)
		return
	}

	err = h.m.Register(model.User{Email: data.Email, Password: data.Password})

	if err != nil {

		var msg string
		switch err {
		case model.ErrInvalidEmail:
			msg = "Invalid eamil"
		case model.ErrInvalidPass:
			msg = "Invalid password"
		case model.ErrUserAlreadyExist:
			msg = "User already exist"
		default:
			msg = "There was an problem, please try again"
		}
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   msg,
		}
		response.Writer(w, res)
		return
	}
	res := response.Resp{
		Status: "succes",
		Data:   true,
	}
	response.Writer(w, res)
}
