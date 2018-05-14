package id_generator

/*
id生成器客户端

ex:

libs.GetIdGenerator().send("test")

*/
import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type idGenerator struct {
	Server string // 服务器地址
}

var IdGenerator = &idGenerator{}

// 设置服务器地址
func (this *idGenerator) SetServer(url string) {
	this.Server = url
}

// 向服务器发送数据
// key为要请求的id项目的名字，num是要请求的个数
// 当请求成功时将会返因一个id的slice
func (this *idGenerator) Send(key string, num ...int64) ([]int64, error) {
	if len(num) == 0 {
		num = []int64{1}
	}
	urlParames := &url.Values{}
	urlParames.Add("name", key)
	urlParames.Add("num", strconv.FormatInt(num[0], 10))
	r, err := http.Get(this.Server + "?" + urlParames.Encode())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// 获取返回的数据内容
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	bodyStr := string(body)

	// 判断返回的状态码
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(bodyStr)
	}

	rs := []int64{}
	// 当请求的ID的数量比较大时，说明请求的内容将会按一个数组的方式返回
	if num[0] > 1 {
		var arr []string
		arr = strings.Split(bodyStr, ",")
		for _, v := range arr {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			rs = append(rs, id)
		}
		return rs, nil
	} else {
		id, err := strconv.ParseInt(bodyStr, 10, 64)
		if err != nil {
			return nil, err
		}
		return []int64{id}, nil
	}
	return nil, errors.New("can`t get data")
}

// 生成Ids, 36进制的字符串, 使用前请先Init
func GenIds(num int64) (ids []string, err error) {
	idList, err := IdGenerator.Send("PagesItems", num)
	if err != nil {
		return
	}

	if len(idList) != int(num) {
		err = fmt.Errorf("IdGenerator return len of ids want %d,but %d", num, len(idList))
		return
	}

	ids = make([]string, len(idList))
	for k, v := range idList {
		ids[k] = strconv.FormatInt(v, 36)
	}

	return
}

// 生成ids, 10进制的字符串
func GenIdsInt(num int64) (ids []int, err error) {
	idList, err := IdGenerator.Send("PagesItems", num)
	if err != nil {
		return
	}

	if len(idList) != int(num) {
		err = fmt.Errorf("IdGenerator return len of ids want %d,but %d", num, len(idList))
		return
	}

	ids = make([]int, len(idList))
	for k, v := range idList {
		ids[k] = int(v)
	}

	return
}

func Init(server string) {
	if server != "" {
		IdGenerator.SetServer(server)
	} else {
		log.Panic("id generator server not exists or empty")
	}
}
