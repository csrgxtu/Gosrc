package main

import (
    "github.com/ddliu/go-httpclient"
)

func main() {
    httpclient.Defaults(httpclient.Map {
        httpclient.OPT_USERAGENT: "Googlebot",
        "Accept-Language": "en-us",
    })

    res, err := httpclient.Get("http://www.baidu.com/s", map[string]string{
        "wd": "csrgxtu",
    })

    println(res.StatusCode, err)
}
