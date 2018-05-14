package page

import (
	"errors"
	"fmt"
)

type Page struct {
	PageId     int    `json:"pageid" xorm:"'pageid' not null int(11) pk autoincr"`
	SiteId     int    `json:"siteid" xorm:"'siteid' not null int(11) index"`
	Type       int    `json:"type" 		  xorm:"'type' not null tinyint(1)"`
	PageName   string `json:"page_name"   xorm:"'page_name' varchar(255)"`
	CustomUrl  string `json:"custom_url"  xorm:"'custom_url' varchar(255) index"`
	Sort       int    `json:"sort" 		  xorm:"'sort' default 0 int(11)"`
	SeoTitle   string `json:"seo_title"   xorm:"'seo_title' not null json text"`
	SeoKeyword string `json:"seo_keyword" xorm:"'seo_keyword' not null json text"`
	SeoDesc    string `json:"seo_desc" 	  xorm:"'seo_desc' not null json text"`
	Ico        string `json:"ico" 		  xorm:"'ico' not null text"`
	IsHome     bool   `json:"is_home" 	  xorm:"'is_home' not null default 0 tinyint(1)"`
	IsPublish  bool   `json:"is_publish"  xorm:"'is_publish' not null default 0 tinyint(1)"`
}

type PageExtend struct {
	PageId       int    `json:"pageid" xorm:"'pageid' not null int(11) pk"`
	SiteId       int    `json:"siteid" xorm:"'siteid' not null int(11) index"`
	OnlineLayout string `json:"online_layout" xorm:"'online_layout' null varchar(20)"`
	DraftLayout  string `json:"draft_layout" xorm:"'draft_layout' null varchar(20)"`
}

func NewPageExtend() *PageExtend {
	return &PageExtend{}
}
func (p PageExtend) TableName() string {
	return "page_extend"
}

func NewPage() *Page {
	return &Page{}
}

func (p Page) TableName() string {
	return "page"
}

// 获取扩展表 page_extend 的信息
func (p *PageExtend) Get() (exist bool, err error) {

	if p.PageId < 1 {
		fmt.Println("enter pageid", p.PageId)
		return
	}
	if p.SiteId < 1 {
		err = errors.New("invalid siteid")
		return
	}
	exist, err = engine.Read.Get(p)
	return
}

// 获取一条页面信息
func (p *Page) Get() (exist bool, err error) {

	if p.PageId < 1 {
		fmt.Println("enter pageid", p.PageId)
		return
	}
	if p.SiteId < 1 {
		err = errors.New("invalid siteid")
		return
	}
	exist, err = engine.Read.Get(p)
	return
}

// 获取COPY
func (p *Page) GetOne() (exist bool, err error) {

	if p.PageId < 1 {
		fmt.Println("enter pageid", p.PageId)
		return
	}

	exist, err = engine.Read.Get(p)
	return
}

//获取列表
func (p *Page) List(where string, whereArgs ...interface{}) (data []Page, err error) {
	session := engine.Read.NewSession()
	if where != "" {
		session.Where(where, whereArgs...)
	}
	err = session.Find(&data)
	if len(data) == 0 {
		data = make([]Page, 0)
	}
	return
}

// 删除页面
func (p *Page) Delete(siteid, pageid int) (num int64, err error) {
	session := engine.Write.NewSession()
	session.Where("pageid=?", pageid)

	return session.Delete(&Page{})
}

// 将站点下所有页面 is_home 设为 0 false
func (p *Page) UpdateSiteHome(siteid int, cols ...string) (num int64, err error) {
	session := engine.Write.NewSession()
	session.Where("siteid=?", siteid)
	if len(cols) > 0 {
		session.Cols(cols...)
	} else {
		session.AllCols()
	}
	return session.Update(p)
}

// 设置主页
func (p *Page) SendPageMsgHome(pageid int, cols ...string) (num int64, err error) {
	session := engine.Write.NewSession()
	session.Id(pageid)
	if len(cols) > 0 {
		session.Cols(cols...)
	} else {
		session.AllCols()
	}
	return session.Update(p)
}

// 设置seo-新增 SEO信息
func (p *Page) SendPageSeo() (num int64, err error) {
	if p.PageId < 1 {
		err = errors.New("invalid pageid")
		return
	}
	if p.SiteId < 1 {
		err = errors.New("invalid siteid")
		return
	}
	num, err = engine.Write.Insert(p)
	return
}

//设置seo-更新 SEO信息
func (p *Page) UpdatePageSeo(clos ...string) (num int64, err error) {
	if p.PageId < 1 {
		err = errors.New("invalid pageid")
		return
	}
	if p.SiteId < 1 {
		err = errors.New("invalid siteid")
		return
	}
	session := engine.Write.NewSession()
	if len(clos) > 0 {
		session.Cols(clos...)
	} else {
		session.AllCols()
	}
	num, err = session.Id(p.PageId).Update(p)
	return
}

// 更新页面信息
func (p *Page) UpdatePage(clos ...string) (num int64, err error) {
	if p.PageId < 1 {
		err = errors.New("invalid pageid")
		return
	}
	if p.SiteId < 1 {
		err = errors.New("invalid siteid")
		return
	}
	session := engine.Write.NewSession()
	if len(clos) > 0 {
		session.Cols(clos...)
	} else {
		session.AllCols()
	}
	num, err = session.Id(p.PageId).Update(p)
	return
}

// 根据 传入siteid 插入数据到该siteid 下
func InsertPage(copyPage interface{}) (num int64, err error) {
	session := engine.Write.NewSession()
	num, err = session.Insert(copyPage)
	return
}
