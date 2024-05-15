package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"

	"bookstore/cmd/dao"
	"bookstore/cmd/model"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	loggedIn, _ := dao.IsLoggedIn(r)
	if loggedIn {
		FetchBooksByPriceRangeHandler(w, r)
	} else {

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		user, _ := dao.GetUserByUsernameAndPassword(username, password)
		switch {
		case user != nil:
			sessionID := uuid.NewString()
			session := &model.Session{
				SessionID: sessionID,
				Username:  user.Username,
				UserID:    user.ID,
			}
			if err := dao.CreateSession(session); err != nil {
				log.Printf("CreateSession failed: %v", err)
				http.Error(w, "Failed to create session", http.StatusInternalServerError)
				return
			}

			cookie := http.Cookie{
				Name:     "user",
				Value:    sessionID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			tmpl := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			if err := tmpl.Execute(w, user); err != nil {
				log.Printf("Failed to execute LoginHandler template: %s", err)
				http.Error(w, "Failed to render template", http.StatusInternalServerError)
				return
			}
		default:
			tmpl := template.Must(template.ParseFiles("views/pages/user/login.html"))
			if err := tmpl.Execute(w, "帳號或密碼不正確"); err != nil {
				log.Printf("Failed to execute LoginHandler template: %s", err)
				http.Error(w, "Failed to render template", http.StatusInternalServerError)
				return
			}
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		if err := dao.DeleteSession(cookieValue); err != nil {
			log.Printf("DeleteSession failed: %v", err)
			http.Error(w, "Failed to delete session", http.StatusInternalServerError)
			return
		}
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}

	FetchBooksByPriceRangeHandler(w, r)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := dao.GetUserByUsername(username)
	switch {
	case user != nil:
		tmpl := template.Must(template.ParseFiles("views/pages/user/register.html"))
		if err := tmpl.Execute(w, "帳號名已存在"); err != nil {
			log.Printf("Failed to execute RegisterHandler template: %s", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	default:
		if err := dao.SaveUser(username, password, email); err != nil {
			log.Printf("SaveUser failed: %v", err)
			http.Error(w, "Failed to save user", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("views/pages/user/register_success.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("Failed to execute RegisterHandler template: %s", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	}
}

func CheckUserNameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")

	user, _ := dao.GetUserByUsername(username)
	switch {
	case user != nil:
		if _, err := w.Write([]byte("<font style='color:red'>帳號名已存在</font>")); err != nil {
			log.Printf("GetUserByUsername failed: %v", err)
			http.Error(w, "Failed to get user", http.StatusInternalServerError)
			return
		}
	default:
		if _, err := w.Write([]byte("<font style='color:green'>帳號名可使用</font>")); err != nil {
			log.Printf("Failed to write response: %v", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}
