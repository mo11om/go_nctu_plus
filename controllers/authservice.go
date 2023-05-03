package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/joho/godotenv/autoload"
)

var redirect_uri string = os.Getenv("redirect_uri")
var auth_server_url string = os.Getenv("auth_server_url")
var client_id string = os.Getenv("CLIENT_ID")

type tokenResponse struct {
	Access_token string
	Expires_in   int
	Token_type   string

	Scope         string
	Refresh_token string
}

type profileResponse struct {
	Username string
	Email    string
}

func Get_login_uri() string {

	var login_uri string = auth_server_url + "/o/authorize/?" + "client_id=" + client_id + "&response_type=code&scope=profile&redirect_uri=" + redirect_uri + "/api/v1/oauth/code"
	println("auth_server_url", auth_server_url)
	println(login_uri)
	return login_uri
}

func get_token(code string) string {

	//send_redirect_uri := "http://localhost:5173"
	send_redirect_uri := redirect_uri + "/api/v1/oauth/code"
	token_url := auth_server_url + "/o/token/"

	data := url.Values{}

	data.Add("grant_type", "authorization_code")
	data.Add("code", code)
	data.Add("client_id", client_id)
	data.Add("client_secret", os.Getenv("CLIENT_SECRET"))
	data.Add("redirect_uri", send_redirect_uri)

	resp, err := http.PostForm(token_url, data)

	//resp, err := http.Post(token_url, "application/x-www-form-urlencoded", bytes.NewBuffer(send_header))

	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	defer resp.Body.Close()

	//resp_print(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	token := tokenResponse{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("phrase err", err)
		return ""
	}
	fmt.Println("token", token.Access_token)
	return token.Access_token

}

func get_profile(token string) profileResponse {
	profile := profileResponse{}
	client := &http.Client{}

	url := auth_server_url + "/api/profile/"
	fmt.Println("URL:>", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("send err", err)
		return profile
	}

	//resp_print(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(" err", err)
		return profile
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		fmt.Println("phrase err", err)
		return profile
	}
	fmt.Println("profile ", profile)

	defer resp.Body.Close()

	return profile
}

func Check_user_exit_and_create(profile profileResponse) (NCTU_User, error) {
	user := FindUserByStudent_Id(profile.Username)

	if user.UserId == 0 {
		err := CreateUser(profile.Username, profile.Email)
		if err != nil {
			return user, err
		}
		user := FindUserByStudent_Id(profile.Username)

		return user, nil

	}
	return user, nil

}
func get_Jwt(profile profileResponse) (string, error) {
	expiresAt := time.Now().Add(10 * time.Hour).Unix()
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	user, err := Check_user_exit_and_create(profile)
	if err != nil {
		return "", err
	}
	fmt.Println("user_id: ", user.UserId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.UserId,
		"student_id": profile.Username,
		"exp":        expiresAt,
	})

	// Sign and get the complete encoded token as a string using the

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	fmt.Println(tokenString, err)
	return tokenString, nil
}

func Get_jwt_token(code string) (string, error) {

	fmt.Println("get code", code)

	token := get_token(code)
	if token != "" {
		profile := get_profile(token)
		if (profile != profileResponse{}) {
			fmt.Println("ok ")

			return get_Jwt(profile)

		}

	}
	return "", fmt.Errorf("no token found")
}
