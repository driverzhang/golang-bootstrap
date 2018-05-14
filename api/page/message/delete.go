// Delete()
// 删除页面
package message

import (
	"errors"
	"strconv"

	"git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/log"
	"git.zhuzi.me/zzjz/zhuzi-bootstrap/model/page"

	"gopkg.in/gin-gonic/gin.v1"
)

type deleteParms struct {
	SiteId int `json:"siteid"`
	PageId int `json:"pageid"`
}

func (s *deleteParms) Bind(ctx *gin.Context) (err error) {
	if s.SiteId, err = strconv.Atoi(ctx.Param("siteid")); err != nil {
		return errors.New("非法请求 siteid")
	}

	if s.PageId, err = strconv.Atoi(ctx.Param("pageid")); err != nil {
		return errors.New("非法请求 pageid")
	}
	return
}

//删除页面
func Delete(ctx *gin.Context) interface{} {
	var (
		err error

		requestData deleteParms

		pageModel = page.NewPage()
	)
	if err = requestData.Bind(ctx); err != nil {
		return err
	}
	_, err = pageModel.Delete(requestData.SiteId, requestData.PageId)
	if err != nil {
		log.Error("删除页面失败！,err", err.Error(), "pageid", requestData.PageId)
		return err
	}
	return "ok"
}
