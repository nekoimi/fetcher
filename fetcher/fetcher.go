package fetcher

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/nekoimi/fetcher/conf"
)

var (
	// 请求远程失败
	ErrStatusCodeNotOk = errors.New("request remote error")
)

func NewHttpClient() *http.Client {
	if conf.EnableProxy {
		url, err := url.Parse(conf.HttpProxyHost)
		if err != nil {
			panic(err)
		}
		return &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(url),
			},
		}
	}
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func Fetch(url string) ([]byte, error) {
	client := NewHttpClient()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Referer", conf.RootUrl)
	request.Header.Set("User-Agent", randUserAgent())

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, ErrStatusCodeNotOk
	}

	return ioutil.ReadAll(response.Body)
}

// 随机User-Agent
func randUserAgent() string {
	UserAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60",
		"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
	}
	rand.Seed(time.Now().UnixNano())
	ri := rand.Intn(len(UserAgents))
	return UserAgents[ri]
}

func RegexpReplaceEmpty(content []byte, regex string) string {
	html := string(content)
	re, _ := regexp.Compile(regex)
	return re.ReplaceAllString(html, "")
}
