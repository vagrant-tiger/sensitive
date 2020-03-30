package main

import (
	"fmt"
	"github.com/importcjj/sensitive"
	"github.com/labstack/echo"
	"net/http"
)

var (
	filter *sensitive.Filter
	e *echo.Echo
)

/**
 * 定义返回值结构
 */
type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type (
	// text
	Text struct {
		Str string `json:"str" form:"str" query:"str"`
	}

	// word
	Word struct {
		Word string `json:"word" form:"word" query:"word"`
	}

	// result
	Result struct {
		Result int `json:"result"`
		Words []string `json:"words"`
	}
)

/**
 * 返回json
 */
func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

func main() {
	fmt.Println("start")

	// 从文件载入到内存
	initStore()

	e = echo.New()

	e.POST("/replace", replace)
	e.POST("/filter", filterWord)
	e.POST("/validate", validateWord)
	e.POST("/findAll", findAll)
	e.POST("/addWord", addWord)
	e.POST("/delWord", delWord)

	// 启动
	e.Logger.Fatal(e.Start(":80"))
}

/**
 * 敏感词载入
 */
func initStore() {
	fmt.Println("data start load...")

	filter = sensitive.New()
	err := filter.LoadWordDict("dic/mgc.txt")
	if err != nil {
		e.Logger.Error(err.Error())
		fmt.Println(err.Error())
	}

	fmt.Println("data loaded success")
}

/**
 * 替换敏感词
 */
func replace(c echo.Context) error {
	result := NewBaseJsonBean()
	str := new(Text)

	if err := c.Bind(str); err != nil {
		result.Code = 1
		result.Message = err.Error()
		return c.JSON(http.StatusOK, result)
	}

	if str.Str == "" {
		result.Code = 1
		result.Message = "参数错误"
		return c.JSON(http.StatusOK, result)
	}

	res := filter.Replace(str.Str, '*')

	result.Data = res

	//向客户端返回JSON数据
	return c.JSON(http.StatusOK, result)
}

/**
 * 过滤敏感词
 */
func filterWord(c echo.Context) error {
	result := NewBaseJsonBean()
	str := new(Text)

	if err := c.Bind(str); err != nil {
		result.Code = 1
		result.Message = err.Error()
		return c.JSON(http.StatusOK, result)
	}

	if str.Str == "" {
		result.Code = 1
		result.Message = "参数错误"
		return c.JSON(http.StatusOK, result)
	}

	res := filter.Filter(str.Str)

	result.Data = res

	//向客户端返回JSON数据
	return c.JSON(http.StatusOK, result)
}

/**
 * 验证内容是否包含敏感词，有的话返回0和第一个敏感词，没有的话返回1
 */
func validateWord(c echo.Context) error {
	result := NewBaseJsonBean()
	str := new(Text)

	if err := c.Bind(str); err != nil {
		result.Code = 1
		result.Message = err.Error()
		return c.JSON(http.StatusOK, result)
	}

	if str.Str == "" {
		result.Code = 1
		result.Message = "参数错误"
		return c.JSON(http.StatusOK, result)
	}

	res, sWord := filter.Validate(str.Str)

	resData := new(Result)
	if res {
		resData.Result = 1
	} else {
		resData.Result = 0
		resData.Words = append(resData.Words, sWord)
	}

	result.Data = resData

	//向客户端返回JSON数据
	return c.JSON(http.StatusOK, result)
}

/**
 * 返回全部敏感词
 */
func findAll(c echo.Context) error {
	result := NewBaseJsonBean()
	str := new(Text)

	if err := c.Bind(str); err != nil {
		result.Code = 1
		result.Message = err.Error()
		return c.JSON(http.StatusOK, result)
	}

	if str.Str == "" {
		result.Code = 1
		result.Message = "参数错误"
		return c.JSON(http.StatusOK, result)
	}

	res := filter.FindAll(str.Str)

	result.Data = res

	//向客户端返回JSON数据
	return c.JSON(http.StatusOK, result)
}

/**
 * 添加敏感词
 */
func addWord(c echo.Context) error {
	result := NewBaseJsonBean()
	word := new(Word)

	if err := c.Bind(word); err != nil {
		result.Code = 1
		result.Message = err.Error()
		return c.JSON(http.StatusOK, result)
	}

	if word.Word == "" {
		result.Code = 1
		result.Message = "参数错误"
		return c.JSON(http.StatusOK, result)
	}

	filter.AddWord(word.Word)

	// 只添加到了内存中，没有保存到文件，需要处理（文件 or 数据库）todo


	return c.JSON(http.StatusOK, result)
}

/**
 * 删除敏感词
 */
func delWord(c echo.Context) error {
	result := NewBaseJsonBean()
	word := new(Word)

	if err := c.Bind(word); err != nil {
		result.Code = 1
		result.Message = err.Error()
		return c.JSON(http.StatusOK, result)
	}

	if word.Word == "" {
		result.Code = 1
		result.Message = "参数错误"
		return c.JSON(http.StatusOK, result)
	}

	filter.DelWord(word.Word)

	// 只在内存中删除，没有同步到文件，需要处理（文件 or 数据库）todo


	return c.JSON(http.StatusOK, result)
}