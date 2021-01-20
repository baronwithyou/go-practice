package main

import (
	"fmt"

	"github.com/baronwithyou/go-practice/20ex/08poster/poster"
)

/**
 * 在iterm上监听用户的输入
 * 用户每输入一个string就查询一次，打印相关信息并且将图片下载下来
 */

const DIR = "/tmp"

func main() {

	res, err := poster.SearchMoive("iron man")

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = poster.DownloadImage(res.Poster, DIR, res.Title); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
