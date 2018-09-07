package merak

import (
	"io"
	"net/http"
	neturl "net/url"
	"os"
	"strings"

	"github.com/luoyancn/dubhe/logging"
)

func Download(url string, proxy string) error {
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		logging.LOG.Errorf("Cannot download file:%v\n", err)
		return err
	}
	req.Header.Set("Connection", "close")
	var resp *http.Response
	if "" != proxy {
		_proxy, err := neturl.Parse(proxy)
		if err != nil {
			return err
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(_proxy),
			},
		}
		resp, err = client.Do(req)
	} else {
		resp, err = http.DefaultClient.Do(req)
	}
	if nil != err || resp.StatusCode != http.StatusOK {
		logging.LOG.Errorf("Cannot download file, HTTP Error:%v\n",
			resp.StatusCode)
		return err
	}

	_str := strings.SplitAfter(url, "/")
	_download, err := os.Create(_str[len(_str)-1])
	if nil != err {
		logging.LOG.Errorf("Cannot create local file:%v\n", err)
		return err
	}
	defer resp.Body.Close()
	defer _download.Close()
	_, err = io.Copy(_download, resp.Body)
	if nil != err {
		logging.LOG.Errorf("Fail to save file:%v\n", err)
		return err
	}
	return nil
}
