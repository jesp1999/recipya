package server_test

import (
	"encoding/json"
	"github.com/reaper47/recipya/internal/auth"
	"github.com/reaper47/recipya/internal/models"
	"github.com/reaper47/recipya/internal/server"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlers_Auth_Confirm(t *testing.T) {
	srv := newServerTest()
	srv.Repository = &mockRepository{
		UsersRegistered: []models.User{
			{ID: 1, Email: "test@example.com"},
		},
	}

	const uri = "/auth/confirm"

	t.Run("missing token", func(t *testing.T) {
		rr := sendRequest(srv, http.MethodGet, uri, noHeader, nil)

		assertStatus(t, rr.Code, http.StatusSeeOther)
		assertHeader(t, rr, "Location", "/")
	})

	t.Run("invalid confirmation token for existing user", func(t *testing.T) {
		token, _ := auth.CreateToken(map[string]any{"userID": 2}, 1*time.Nanosecond)

		rr := sendRequest(srv, http.MethodGet, uri+"?token="+token, noHeader, nil)

		assertStatus(t, rr.Code, http.StatusBadRequest)
		want := []string{
			`<title hx-swap-oob="true">Token Expired | Recipya</title>`,
			"The token associated with the URL expired.",
		}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})

	t.Run("user does not exist", func(t *testing.T) {
		token, _ := auth.CreateToken(map[string]any{"userID": 2}, 1*time.Hour)

		rr := sendRequest(srv, http.MethodGet, uri+"?token="+token, noHeader, nil)

		assertStatus(t, rr.Code, http.StatusNotFound)
		want := []string{
			`<title hx-swap-oob="true">Confirm | Recipya</title>`,
			"An error occurred when you requested to confirm your account.",
		}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})

	t.Run("valid confirmation token for existing user", func(t *testing.T) {
		token, _ := auth.CreateToken(map[string]any{"userID": 1}, 1*time.Hour)

		rr := sendRequest(srv, http.MethodGet, uri+"?token="+token, noHeader, nil)

		assertStatus(t, rr.Code, http.StatusOK)
		want := []string{
			`<title hx-swap-oob="true">Confirm | Recipya</title>`,
			"Your account has been confirmed.",
		}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})
}

func TestHandlers_Auth_Login(t *testing.T) {
	srv := newServerTest()
	repo := &mockRepository{
		UsersRegistered: []models.User{
			{ID: 1, Email: "test@example.com"},
		},
	}
	srv.Repository = repo

	const uri = "/auth/login"

	testcases := []struct {
		name string
		form string
	}{
		{name: "email invalid", form: "email=hello@test&password=123"},
		{name: "email invalid", form: "email=hello@test.com"},
		{name: "user not found", form: "email=hello@test.com&password=123"},
	}
	for _, tc := range testcases {
		t.Run("invalid credentials when "+tc.name, func(t *testing.T) {
			rr := sendRequest(srv, http.MethodPost, uri, formHeader, strings.NewReader(tc.form))

			assertStatus(t, rr.Code, http.StatusNoContent)
			var got map[string]string
			_ = json.Unmarshal([]byte(rr.Header().Get("HX-Trigger")), &got)
			wantHeader := `{"message":"Credentials are invalid.","backgroundColor":"bg-red-500"}`
			if got["showToast"] != wantHeader {
				t.Fatalf("got\n%q\nbut want\n%q", got, wantHeader)
			}
		})
	}

	t.Run("redirect to home when logged in", func(t *testing.T) {
		rr := sendRequestAsLoggedIn(srv, http.MethodPost, uri, formHeader, nil)

		assertStatus(t, rr.Code, http.StatusSeeOther)
		assertHeader(t, rr, "Location", "/")
	})

	t.Run("redirect to accessed uri after logged in", func(t *testing.T) {
		otherUri := "/recipes/add"
		r := httptest.NewRequest(http.MethodPost, uri, strings.NewReader("email=test@example.com&password=123&remember-me=false"))
		r.Header.Set("Content-Type", string(formHeader))
		r.AddCookie(server.NewRedirectCookie(otherUri))

		rr := httptest.NewRecorder()
		srv.Router.ServeHTTP(rr, r)

		assertStatus(t, rr.Code, http.StatusSeeOther)
		assertHeader(t, rr, "HX-Redirect", otherUri)
	})

	t.Run("login  successful", func(t *testing.T) {
		maps.Clear(server.SessionData)

		rr := sendRequest(srv, http.MethodPost, uri, formHeader, strings.NewReader("email=test@example.com&password=123&remember-me=false"))

		assertStatus(t, rr.Code, http.StatusSeeOther)
		assertHeader(t, rr, "HX-Redirect", "/")
		var (
			isUserInSession   bool
			isCookieStoresSID bool
		)
		for sid, userID := range server.SessionData {
			if userID == 1 {
				isUserInSession = true
				isCookieStoresSID = slices.ContainsFunc(rr.Result().Cookies(), func(cookie *http.Cookie) bool {
					return cookie.Name == "session" && cookie.Value == sid.String()
				})
				break
			}
		}
		if !isUserInSession {
			t.Fatal("expected user to be in the server's session data")
		}
		if !isCookieStoresSID {
			t.Fatal("expected session ID to be stored in a cookie named 'session'")
		}
	})

	t.Run("user checked remember me", func(t *testing.T) {
		numAuthTokensBefore := len(repo.AuthTokens)

		rr := sendRequest(srv, http.MethodPost, uri, formHeader, strings.NewReader("email=test@example.com&password=123&remember-me=true"))

		assertStatus(t, rr.Code, http.StatusSeeOther)
		assertHeader(t, rr, "HX-Redirect", "/")

		cookies := rr.Result().Cookies()
		index := slices.IndexFunc(cookies, func(cookie *http.Cookie) bool { return cookie.Name == "remember_me" })
		if index == -1 {
			t.Fatal("there must be a session cookie")
		}
		if cookies[index].Expires.Before(time.Now().Add(30 * 24 * time.Hour).Add(-1 * time.Minute)) {
			t.Fatalf("got expiration %v but want an expiration of 1 month", cookies[index].Expires)
		}

		if len(repo.AuthTokens) != numAuthTokensBefore+1 {
			t.Fatal("expected an authentication token to be added to the database")
		}
	})

	t.Run("user checked remember me and accesses login page", func(t *testing.T) {
		rr := sendRequest(srv, http.MethodPost, uri, formHeader, strings.NewReader("email=test@example.com&password=123&remember-me=true"))
		r := httptest.NewRequest(http.MethodGet, uri, nil)
		for _, c := range rr.Result().Cookies() {
			if c.Name != "session" {
				r.AddCookie(c)
			}
		}

		rr = httptest.NewRecorder()
		srv.Router.ServeHTTP(rr, r)

		assertStatus(t, rr.Code, http.StatusSeeOther)
		assertHeader(t, rr, "HX-Redirect", "/")
	})
}

func TestHandlers_Auth_Logout(t *testing.T) {
	repo := &mockRepository{}
	srv := server.NewServer(repo, &mockEmail{}, &mockFiles{})

	const uri = "/auth/logout"

	t.Run("cannot log out a user who is already logged out", func(t *testing.T) {
		originalNumSessions := len(server.SessionData)

		rr := sendRequest(srv, http.MethodPost, uri, noHeader, nil)

		assertStatus(t, rr.Code, http.StatusOK)
		if originalNumSessions != len(server.SessionData) {
			t.Fatalf("expected same number of sessions")
		}
	})

	t.Run("valid logout for a logged-in user", func(t *testing.T) {
		maps.Clear(server.SessionData)
		originalNumSessions := len(server.SessionData) + 1

		rr := sendRequestAsLoggedIn(srv, http.MethodPost, uri, noHeader, nil)

		assertStatus(t, rr.Code, http.StatusSeeOther)
		if len(server.SessionData) != originalNumSessions-1 {
			t.Fatalf("expected one less number of sessions")
		}
		var isCookieInvalid bool
		for _, c := range rr.Result().Cookies() {
			if c.Name == "session" {
				isCookieInvalid = c.MaxAge == -1
			}
		}
		if !isCookieInvalid {
			t.Fatal("expected the session cookie to be invalidated")
		}
	})

	t.Run("remember me user has its token deleted on logout", func(t *testing.T) {
		originalNumAuthTokens := len(repo.AuthTokens) + 1

		rr := repo.sendRequestAsLoggedInRememberMe(srv, http.MethodPost, uri, noHeader, nil)

		assertStatus(t, rr.Code, http.StatusSeeOther)
		var isCookieInvalid bool
		for _, c := range rr.Result().Cookies() {
			if c.Name == "remember_me" {
				isCookieInvalid = c.MaxAge == -1
			}
		}
		if !isCookieInvalid {
			t.Fatal("expected the remember me cookie to be invalidated")
		}
		if len(repo.AuthTokens) != originalNumAuthTokens-1 {
			t.Fatal("expected one less auth token in the database")
		}
	})
}

func TestHandlers_Auth_Register(t *testing.T) {
	srv := newServerTest()

	const uri = "/auth/register"

	t.Run("valid email", func(t *testing.T) {
		rr := sendRequest(srv, http.MethodPost, uri+"/validate-email", formHeader, strings.NewReader("email=test@example.com&password=test123&password-confirm=test123"))

		want := []string{

			`<input type="email" class="w-full rounded-lg bg-gray-100 px-4 py-2 border border-green-500" id="email" name="email" placeholder="Enter your email address..." hx-post="/auth/register/validate-email" hx-indicator="#ind" required value="test@example.com"/>`,
		}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})

	t.Run("invalid email", func(t *testing.T) {
		rr := sendRequest(srv, http.MethodPost, uri+"/validate-email", formHeader, strings.NewReader("email=invalid-email&password=test123&password-confirm=test123"))

		want := []string{`<div class="text-red-500 text-xs">Please enter a valid email.</div>`}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})

	t.Run("email already taken", func(t *testing.T) {
		srv.Repository = &mockRepository{
			UsersRegistered: []models.User{
				{ID: 1, Email: "test@example.com"},
			},
		}

		rr := sendRequest(srv, http.MethodPost, uri+"/validate-email", formHeader, strings.NewReader("email=test@example.com&password=test123&password-confirm=test123"))

		want := []string{`<div class="text-red-500 text-xs">This email is already taken. Please enter another one.</div>`}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})

	t.Run("valid passwords", func(t *testing.T) {
		rr := sendRequest(srv, http.MethodPost, uri+"/validate-password", formHeader, strings.NewReader("email=invalid-email&password=test123&password-confirm=test123"))

		want := []string{`<input class="w-full rounded-lg bg-gray-100 px-4 py-2 border border-green-500" id="password-confirm" name="password-confirm"`}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})

	t.Run("invalid password", func(t *testing.T) {
		rr := sendRequest(srv, http.MethodPost, uri+"/validate-password", formHeader, strings.NewReader("email=test@example.com&password=test123&password-confirm=test"))

		want := []string{`<div class="text-red-500 text-xs">Passwords do not match.</div>`}
		assertStringsInHTML(t, getBodyHTML(rr), want)
	})

	t.Run("redirect to home when logged in", func(t *testing.T) {
		rr := sendRequestAsLoggedIn(srv, http.MethodPost, uri, formHeader, nil)

		assertStatus(t, rr.Code, http.StatusSeeOther)
	})

	t.Run("invalid registration user already exists", func(t *testing.T) {
		email := "invalid@register.com"
		_ = sendRequest(srv, http.MethodPost, uri, formHeader, strings.NewReader("email="+email+"&password=test123&password-confirm=test123"))
		originalNumUsers := len(srv.Repository.Users())

		rr := sendRequest(srv, http.MethodPost, uri, formHeader, strings.NewReader("email="+email+"&password=test123&password-confirm=test123"))

		assertStatus(t, rr.Code, http.StatusUnprocessableEntity)
		assertHeader(t, rr, "HX-Trigger", `{"showToast":"{\"message\":\"Registration failed.\",\"backgroundColor\":\"bg-red-500\"}"}`)
		if len(srv.Repository.Users()) != originalNumUsers {
			t.Fatalf("expected no user to be added to the db of %d users", originalNumUsers)
		}
	})

	t.Run("valid registration for new user", func(t *testing.T) {
		email := "regsiter@valid.com"
		originalNumUsers := len(srv.Repository.Users())

		rr := sendRequest(srv, http.MethodPost, uri, formHeader, strings.NewReader("email="+email+"&password=test123&password-confirm=test123"))

		assertStatus(t, rr.Code, http.StatusSeeOther)
		assertHeader(t, rr, "HX-Redirect", "/auth/login")
		users := srv.Repository.Users()
		if len(users) != originalNumUsers+1 {
			t.Fatalf("expected %d users but got %d", originalNumUsers+1, len(users))
		}
		if users[len(users)-1].Email != email {
			t.Fatalf("got user %+v but want %s", users[len(users)-1], email)
		}
	})
}
