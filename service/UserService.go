package service
import(
	"net/http"
	"api/pojo"
	"github.com/gin-gonic/gin"
)
var  commentList =[]string{"123","456"}
//get all 
func FindAllComment(ctx *gin.Context){
	// ctx.JSON(http.StatusOK,commentList)

}
func GetUserById(ctx *gin.Context){
	 comment  := pojo.FindCommentId(ctx.Param("id")) 
	if  comment. Id == 0 {
		ctx.JSON(http.StatusNotFound,"ERROR")
		return
    }
	 
	ctx.JSON(http.StatusOK,comment)
}
func POSTAllComment(ctx *gin.Context){
	Page := pojo.Page{	}
	err := ctx.BindJSON(&Page)
	if err== nil{
	    ctx.JSON(http.StatusNotAcceptable,err);
			       return
    }
	ctx.JSON(http.StatusOK,"Page post success")
}