package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

var redirect_uri string = os.Getenv("redirect_uri")
var auth_server_url string = os.Getenv("auth_server_url")
var client_id string = os.Getenv("CLIENT_ID")
var client_secret string = os.Getenv("CLIENT_SECRET")

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

func Nycu_Oauth_get_login_uri() string {

	var login_uri string = auth_server_url + "/o/authorize/?" + "client_id=" + client_id + "&response_type=code&scope=profile&redirect_uri=" + redirect_uri + "/api/v1/oauth/code"
	return login_uri
}
func Nycu_Oauth_redirect(ctx *gin.Context) {

	ctx.Redirect(303, Nycu_Oauth_get_login_uri())
}

func Nycu_Oauth_Get_code(ctx *gin.Context) {
	code := ctx.Query("code")

	fmt.Println("get code", code)

	get_token_profile(code)

	// send_header, err := json.MarshalIndent(header,"","\t" )
	// if err != nil {
	//     fmt.Println("err",err)
	//     return
	// }

	// fmt. Println(string (send_header))

}

func get_token_profile(code string) {
	send_redirect_uri := redirect_uri + "/api/v1/oauth/code"
	token_url := auth_server_url + "/o/token/"

	data := url.Values{}

	data.Add("grant_type", "authorization_code")
	data.Add("code", code)
	data.Add("client_id", client_id)
	data.Add("client_secret", client_secret)
	data.Add("redirect_uri", send_redirect_uri)

	resp, err := http.PostForm(token_url, data)

	//resp, err := http.Post(token_url, "application/x-www-form-urlencoded", bytes.NewBuffer(send_header))

	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer resp.Body.Close()

	//resp_print(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	token := tokenResponse{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("phrase err", err)
		return
	}
	fmt.Println("token", token.Access_token)

	get_profile(token.Access_token)
}

func resp_print(resp *http.Response) {
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("response Body:", string(body))
}
func get_profile(token string) {
	url := auth_server_url + "/api/profile/"
	fmt.Println("URL:>", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("send err", err)
		return
	}

	//resp_print(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(" err", err)
		return
	}

	profile := profileResponse{}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		fmt.Println("phrase err", err)
		return
	}
	fmt.Println("profile ", profile)

	defer resp.Body.Close()
}
