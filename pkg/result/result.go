package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回的结果的内容：
type ResultCont struct {
	Code	int  `json:"code"`     //提示代码
	Msg string  `json:"msg"`       //提示信息
	Data interface{} `json:"data"` //出错
}

//放回结果
type Result struct {
	Ctx *gin.Context
}

//生成result
func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

//生成一个error
func NewError(code int, msg string) ResultCont {
	return ResultCont{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	}
}

//成功
func (r *Result) Success(data interface{}) {
	if (data == nil) {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK,res)
	r.Ctx.Abort()
}

//出错,接受code和msg
func (r *Result)ErrorCode(code int,msg string) {
	res := ResultCont{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	//if (code == http.StatusBadRequest)
	r.Ctx.JSON(http.StatusOK,res)
	r.Ctx.Abort()
}


//出错,接受resultcont
func (r *Result)Error(res ResultCont) {
	r.Ctx.JSON(http.StatusOK,res)
	r.Ctx.Abort()
}
