package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

func spider(username string, password string, imagecode string, c *http.Client) {
	url1 := "http://xk1.ahu.cn/default2.aspx"
	//url2 := "http://xk1.ahu.cn/CheckCode.aspx?"
	//url3 := "http://xk1.ahu.cn/xs_main.aspx?xh=P71514011"
	v := url.Values{}
	enc := mahonia.NewEncoder("gbk")
	but := enc.ConvertString("学生")
	v.Add("__VIEWSTATE", "/wEPDwUJODk4OTczODQxZGQhFC7x2TzAGZQfpidAZYYjo/LeoQ==")
	v.Add("txtUserName", username)
	v.Add("TextBox2", password)
	v.Add("txtSecretCode", imagecode)
	v.Add("RadioButtonList1", but)
	v.Add("Button1", "")
	v.Add("lbLanguage", "")
	v.Add("hidPdrs", "")
	v.Add("hidsc", "")
	v.Add("__EVENTVALIDATION", "/wEWDgKX/4yyDQKl1bKzCQLs0fbZDAKEs66uBwK/wuqQDgKAqenNDQLN7c0VAuaMg+INAveMotMNAoznisYGArursYYIAt+RzN8IApObsvIHArWNqOoPqeRyuQR+OEZezxvi70FKdYMjxzk=")
	//建立client发送POST请求
	body := strings.NewReader(v.Encode())
	r, _ := http.NewRequest("POST", url1, body)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Referer", "http://xk1.ahu.cn/default2.aspx")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.104 Safari/537.36")
	response, err := c.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	file3, err := os.OpenFile("spider.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()
	body2, _ := ioutil.ReadAll(response.Body)
	num, err := file3.Write(body2)
	if num != len(body2) {
		log.Fatal(err)
	}
	fmt.Println(response.Status)
}
