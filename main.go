package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("请输入: [%s] [需要从文件名中去掉的部分]", os.Args[0])
		return
	} else {
		fmt.Printf("\n当前文件夹中 所有文件名中去掉 %s 字符\n\n将会改变以下文件：\n", os.Args[1])
		process(os.Args[1], false)
		fmt.Printf("\n\n确认应输入Y？")
		cmd := ""
		fmt.Scanf("%s\n", &cmd)
		if cmd == "Y" || cmd == "y" {
			process(os.Args[1], true)
			fmt.Println("操作成功！")
		} else {
			fmt.Println("操作已取消！")
		}
	}
	time.Sleep(time.Second)
}

func process(src string, real_change bool) {
	fis, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("[process] Error=", err)
		return
	}
	for _, v := range fis {
		if v.IsDir() {
			continue
		}
		n := strings.Index(v.Name(), src)
		if n >= 0 {
			name := v.Name()[:n]
			name = fmt.Sprintf("%s%s", name, v.Name()[n+len(src):])

			if len(name) < 1 {
				fmt.Println("无法更改名称：", v.Name(), " 原因：更改后名字为空")
				continue
			} else {
				fmt.Printf("%-20s ---> %s\n", v.Name(), name)
			}

			if real_change {
				err = os.Rename(v.Name(), name)
				if err != nil {
					fmt.Println("[process] 更名失败 Error=", err)
				}
			}
		}
	}

}
