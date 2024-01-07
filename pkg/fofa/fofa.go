package fofa

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FateBug403/FoFa/pkg/config"
	"github.com/FateBug403/FoFa/pkg/model"
	"github.com/FateBug403/FoFa/pkg/result"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/projectdiscovery/retryablehttp-go"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type FoFa struct {
	Config *config.FoFa
	HttpClient *retryablehttp.Client

}

func NewFoFa(option *config.FoFa) (*FoFa,error) {
	var err error
	// 创建一个 RetryableClient
	client := retryablehttp.NewClient(retryablehttp.Options{
		RetryWaitMin:    1 * time.Second,
		RetryWaitMax:    3 * time.Second,
		Timeout:         15 * time.Second,
		RetryMax:        5,
		RespReadLimit:   0,
		Verbose:         false,
		KillIdleConn:    false,
		CheckRetry:      nil,
		Backoff:         nil,
		NoAdjustTimeout: false,
	})
	// 身份验证
	req, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/info/my?email=%s&key=%s", option.Baseurl, option.Email, option.Key), nil)
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/112.0")
	resp,err:=client.Do(req)
	if err != nil {
		return nil,err
	}
	// 确保在函数结束时关闭响应的主体
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	sj, err := simplejson.NewJson(body)
	if err != nil {
		return nil,err
	}
	var publicKey bool = sj.Get("error").MustBool(true)
	if publicKey{
		return nil,errors.New("FoFa身份信息验证失败")
	}

	return &FoFa{HttpClient: client,Config: option},nil
}

// SearchAll 查询指定的内容返回数据
func (receiver *FoFa) SearchAll(search string) (*result.Result,error){
	var err error
	searchbase64 := base64.StdEncoding.EncodeToString([]byte(search))

	//通过反射获取自定义的结构体的每个字段
	var fields []string
	t := reflect.TypeOf(&result.InFo{}).Elem()
	for i := 1; i < t.NumField(); i++ {
		field:=strings.ToLower(t.Field(i).Name)
		fields=append(fields, field)
	}
	field :=strings.Join(fields,",")

	//获取查询返回数据
	req, err := retryablehttp.NewRequest(http.MethodGet,fmt.Sprintf("%s/api/v1/search/all?email=%s&key=%s&qbase64=%s&fields=%s&size=%d", receiver.Config.Baseurl, receiver.Config.Email, receiver.Config.Key, searchbase64, field,receiver.Config.Size), nil)
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/112.0")
	resp,err:=receiver.HttpClient.Do(req)
	if err != nil {
		return nil,err
	}
	// 确保在函数结束时关闭响应的主体
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var searchResult model.FoFa
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		return nil,err
	}

	//<-------------------------返回------------------------->
	//封装返回的数据
	var Result result.Result
	for num, values := range searchResult.Results {
		var target result.InFo
		// 使用反射包中的 ValueOf 方法获取结构体变量的值
		v := reflect.ValueOf(&target).Elem()
		// 使用循环设置结构体变量的每个字段的值
		for i := 1; i < v.NumField(); i++ {
			//将值转化为Value类型
			fieldValue := reflect.ValueOf(values[i-1])
			//设置某个字段的值
			v.Field(i).Set(fieldValue)
		}
		target.Id= int64(num)
		// 如果host中包含有http,则进行解析，只要域名加端口
		if strings.Contains(target.Host,"://"){
			parsedURL, err := url.Parse(target.Host)
			if err == nil {
				target.Host=parsedURL.Host
			}
		}
		Result.InFos=append(Result.InFos,target)
	}
	log.Println("从fofa搜索到"+fmt.Sprint(len(Result.InFos)))

	return &Result,nil
}