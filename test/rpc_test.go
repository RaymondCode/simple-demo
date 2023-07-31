package test

import (
	"Momotok-Server/rpc"
	"fmt"
	"io"
	"testing"
)

// avatar, _ := http.Get("https://acg.suyanw.cn/sjtx/random.php")                  //prepare to change to rpc, but actually they don't need them
// background_image, _ := http.Get("https://acg.suyanw.cn/api.php")
func TestRpc(t *testing.T) {
	resp, _ := rpc.HttpRequest("GET", "https://acg.suyanw.cn/sjtx/random.php", nil)
	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(resp.Body)
	}
	fmt.Println()

}
