package service
import (
	"os"
	"net/http"
	"encoding/json"
	"fmt"
	 "bytes"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)
	
 




 
var redirect_uri string= os.Getenv("redirect_uri");
var auth_server_url string=os.Getenv("auth_server_url")
var client_id string=os.Getenv("CLIENT_ID")
var client_secret string=os.Getenv("CLIENT_SECRET")

type tokenResponse struct {
	Grant_type string `json:"grant_type"`
	Code string `json:"code"`
	Client_id string `json:"client_id"` 
	Client_secret string `json:"client_secret"`
	Redirect_uri string `json:"redirect_uri"`
}

 
func Nycu_Oauth_get_login_uri  ()string{

	var login_uri string=auth_server_url+"/o/authorize/?"+"client_id="+client_id+"&response_type=code&scope=profile&redirect_uri="+redirect_uri+"/api/v1/oauth/code";
	return login_uri;
}
func Nycu_Oauth_redirect(ctx *gin.Context)  {
	  
	 ctx.Redirect(303,Nycu_Oauth_get_login_uri())
}
func Nycu_Oauth_Get_code(ctx *gin.Context)  {
	code:= ctx.Query("code")
	header := tokenResponse{
	 "authorization_code",
	 code,
	 client_id,
	 client_secret,
	 redirect_uri+"/api/v1/oauth/token",}
	
    fmt.Println("get code",code )
	 
	send_header, err := json.MarshalIndent(header,"","\t" )
    if err != nil {
        fmt.Println("err",err)
        return
    }
	 
	fmt. Println(string (send_header))
	
	
	
	token_url := auth_server_url+"/o/token/"
	fmt.Println("URL:>", token_url)
	resp, err := http.Post(token_url, "application/x-www-form-urlencoded",bytes.NewBuffer(send_header))
	//resp, err := http.Post(token_url, "application/x-www-form-urlencoded", bytes.NewBuffer(send_header))
	
	if err!= nil {
        fmt.Println("err",err)
        return
    }
    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, err := ioutil.ReadAll(resp.Body)
	if
	    err!= nil {
        fmt.Println("err",err)
        return
    }
    fmt.Println("response Body:", string(body))
    defer resp.Body.Close()
    
}	
		