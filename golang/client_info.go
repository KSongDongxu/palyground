// 获取Http客户端常用信息
package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type ClientInfo struct {
	UserAgent      string
	AcceptLanguage string
	ClientIp       string
}

func NewClientInfo(userAgent, acceptLanguage, clientIp string) (clientInfo *ClientInfo) {
	return &ClientInfo{
		UserAgent:      userAgent, // User-Agent
		AcceptLanguage: acceptLanguage, // Accept-Language
		ClientIp:       clientIp,
	}
}

func DefaultClientInfo(userAgent string) (clientInfo *ClientInfo) {
	return &ClientInfo{
		UserAgent:      userAgent,
		AcceptLanguage: "zh-CN",
		ClientIp:       "127.0.0.1",
	}
}

// 是否包含
func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

// 是否匹配 不区分大小写
func isMatch(s, sub string) bool {
	pattern := "(?i:" + sub + ")"
	match, _ := regexp.MatchString(pattern, s)
	return match
}

// 匹配
func match(s, pattern string) string {
	r, _ := regexp.Compile(pattern)
	sub := r.FindString(s)
	return sub
}

func (clientInfo *ClientInfo) GetBrowser() string {
	userAgent := clientInfo.UserAgent
	switch {
	case userAgent == "":
		return ""
	case contains(userAgent, "Chrome"):
		return "Google Chrome"
	case contains(userAgent, "Firefox"):
		return "Mozilla Firefox"
	case contains(userAgent, "iphone") || contains(userAgent, "ipod"):
		return "iphone"
	case contains(userAgent, "ipad"):
		return "ipad"
	case contains(userAgent, "android"):
		return "android"
	case contains(userAgent, "MSIE 12.0") ||
		contains(userAgent, "MSIE 11.0") ||
		contains(userAgent, "MSIE 10.0") ||
		contains(userAgent, "MSIE 9.0") ||
		contains(userAgent, "MSIE 8.0") ||
		contains(userAgent, "MSIE 7.0") ||
		contains(userAgent, "MSIE 6.0"):
		return "Internet Explorer"
	case contains(userAgent, "Maxthon"):
		return "Maxthon"
	case contains(userAgent, "NetCaptor"):
		return "NetCaptor"
	case contains(userAgent, "Netscape"):
		return "Netscape"
	case contains(userAgent, "Lynx"):
		return "Lynx"
	case contains(userAgent, "OPR"):
		return "Opera"
	default:
		return "other"
	}
}

func (clientInfo *ClientInfo) GetBrowserVer() string {
	var (
		userAgent string
		sub       string
	)
	userAgent = clientInfo.UserAgent
	if userAgent == "" {
		return ""
	}
	// 谷歌
	sub = match(userAgent, "(?i:Chrome/(\\d+)\\..*)")
	if sub != "" {
		sArr := strings.Split(sub, " ")
		version := strings.Split(sArr[0], "/")[1]
		return version
	}
	// 火狐
	sub = match(userAgent, "(?i:FireFox/(\\d+)\\..*)")
	if sub != "" {
		sArr := strings.Split(sub, "/")
		version := sArr[1]
		return version
	}
	// IE
	sub = match(userAgent, "(?i:MSIE\\s(\\d+)\\..*)")
	if sub != "" {
		sArr := strings.Split(sub, ";")
		version := strings.Split(sArr[0], " ")[1]
		return version
	}
	// Safari
	sub = match(userAgent, "(?i:AppleWebKit/(\\d+)\\..*)")
	if sub != "" {
		sArr := strings.Split(sub, "/")
		version := strings.Split(sArr[0], " ")[1]
		return version
	}

	// Opera
	sub = match(userAgent, "(?i:OPR[\\s|/](\\d+)\\..*)")
	if sub != "" {
		sArr := strings.Split(sub, "/")
		version := sArr[1]
		return version
	}
	return ""
}

func (clientInfo *ClientInfo) GetOs() string {
	userAgent := clientInfo.UserAgent
	switch {
	case userAgent == "":
		return ""
	case isMatch(userAgent, "win") && contains(userAgent, "95"):
		return "Windows 95"
	case isMatch(userAgent, "win 9x") && contains(userAgent, "4.9.0"):
		return "Windows ME"
	case isMatch(userAgent, "win") && contains(userAgent, "98"):
		return "Windows 98"
	case isMatch(userAgent, "win") && contains(userAgent, "nt 5.0"):
		return "Windows 2000"
	case isMatch(userAgent, "win") && contains(userAgent, "nt 6.0"):
		return "Windows Vista"
	case isMatch(userAgent, "win") && contains(userAgent, "nt 10"):
		return "Windows 10"
	case isMatch(userAgent, "win") && contains(userAgent, "nt 8"):
		return "Windows 8"
	case isMatch(userAgent, "win") && contains(userAgent, "nt 6"):
		return "Windows 7"
	case isMatch(userAgent, "win") && contains(userAgent, "nt 5"):
		return "Windows XP"
	case isMatch(userAgent, "win") && contains(userAgent, "nt"):
		return "Windows NT"
	case isMatch(userAgent, "win") && contains(userAgent, "32"):
		return "Windows 32"
	case isMatch(userAgent, "Linux; Android"):
		return "Android"
	case isMatch(userAgent, "Macintosh;"):
		return "Mac OS"
	case isMatch(userAgent, "iPhone;") && isMatch(userAgent, "Mac OS"):
		return "iPhone OS"
	case isMatch(userAgent, "iPad;") && isMatch(userAgent, "Mac OS"):
		return "iPad OS"
	case isMatch(userAgent, "linux"):
		return "Linux"
	case isMatch(userAgent, "unix"):
		return "Unix"
	case isMatch(userAgent, "sun") && isMatch(userAgent, "os"):
		return "SunOS"
	case isMatch(userAgent, "ibm") && isMatch(userAgent, "os"):
		return "IBM OS/2"
	case isMatch(userAgent, "Mac") && isMatch(userAgent, "PC"):
		return "Macintosh"
	case isMatch(userAgent, "PowerPC"):
		return "PowerPC"
	case isMatch(userAgent, "AIX"):
		return "AIX"
	case isMatch(userAgent, "HPUX"):
		return "HPUX"
	case isMatch(userAgent, "NetBSD"):
		return "NetBSD"
	case isMatch(userAgent, "BSD"):
		return "BSD"
	case isMatch(userAgent, "OSF1"):
		return "OSF1"
	case isMatch(userAgent, "IRIX"):
		return "IRIX"
	case isMatch(userAgent, "FreeBSD"):
		return "FreeBSD"
	default:
		return "未知"
	}
}

func (clientInfo *ClientInfo) GetBrowserLang() string {
	acceptLang := clientInfo.AcceptLanguage
	if acceptLang == "" {
		return ""
	}
	lang := acceptLang[:5]
	if isMatch(lang, "zh-cn") {
		return "简体中文"
	} else if isMatch(lang, "zh") {
		return "繁体中文"
	} else {
		return "English"
	}
}

func (clientInfo *ClientInfo) GetRegion() string {
	clientIp := clientInfo.ClientIp
	if clientIp == "" || clientIp == "127.0.0.1" || clientIp == "::1" {
		return ""
	}
	url := fmt.Sprintf("%s%s", "http://ip.taobao.com/service/getIpInfo.php?ip=", clientIp)
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return ""
	}

	var data struct{
		Code int
		Data map[string]interface{}
	}

	var result map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return ""
	}
	if data.Code != 0 {
		return ""
	}
	// 返回数据
	result = data.Data
	country := result["country"].(string)
	city := result["city"].(string)
	region := result["region"].(string)

	// city 存在
	var s1 = ""
	var s2 = ""
	if region != "" && region != "XX" {
		s1 = region
	}
	if city != "" && city != "XX" && region != s1 {
		s2 = city
	}
	return strings.TrimSpace(country + " " + s1 + " " + s2)
}

func (clientInfo *ClientInfo) IsiPhone() bool {
	userAgent := clientInfo.UserAgent
	return isMatch(userAgent, "iPhone")
}

func (clientInfo *ClientInfo) IsiPad() bool {
	userAgent := clientInfo.UserAgent
	return isMatch(userAgent, "iPad")
}

func (clientInfo *ClientInfo) IsiOS() bool {
	return clientInfo.IsiPad() || clientInfo.IsiPhone()
}

func (clientInfo *ClientInfo) IsAndroid() bool {
	userAgent := clientInfo.UserAgent
	return isMatch(userAgent, "Android")
}

func (clientInfo *ClientInfo) IsMobile() bool {
	return clientInfo.IsiOS() || clientInfo.IsAndroid()
}

