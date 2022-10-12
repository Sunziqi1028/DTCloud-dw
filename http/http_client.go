/**
 * @Author: Sun
 * @Description:
 * @File:  http_client
 * @Version: 1.0.0
 * @Date: 2022/6/15 22:37
 */

package http

import (
	"bytes"
	"fmt"
	"net/http"
	"wg315/global"
)

func PostData(data []byte) int {
	code, err := httpPost(data)
	if err != nil {
		fmt.Println(err)
	}
	return code
}

func httpPost(data []byte) (code int, err error) {
	url := fmt.Sprintf("http://%s:%s/%s", global.Config.RemoteAddr.Addr, global.Config.RemoteAddr.Port, global.Config.RemoteAddr.URl)
	//host := strings.TrimSpace(url)
	fmt.Println(url, "line:28")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("post data err:", err)
		return 500, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response err:", err)
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
