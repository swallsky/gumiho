package app

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
)

// 执行脚本
func (h *Http) scriptBin(sh string) string {
	shfile := path.Join(h.RootDir, "/script/", sh) //执行的文件
	// //判断文件是否存在
	if _, err := os.Stat(shfile); err == nil {
		cmd := exec.Command("/bin/bash", "-c", shfile)
		output, err := cmd.Output()
		if err == nil {
			return fmt.Sprintf("Execute shell:%s finished with output:\n%s", shfile, string(output))
		} else {
			return fmt.Sprintf("Execute shell:%s failed with error:%s", shfile, err.Error())
		}
	} else {
		return fmt.Sprintf("Shell file:%s is not found.\n", shfile)
	}
}

// 读取相应的脚本并执行
func (h *Http) script(c *gin.Context) {
	route := c.Param("R")              //读取shell路由配置
	if _, ok := h.Router[route]; !ok { //判断路由是否在配置文件中
		c.JSON(404, gin.H{
			"code":  404,
			"error": "no route",
		})
	} else { // 正常执行相应的shell
		msg := h.scriptBin(h.Router[route])
		c.JSON(200, gin.H{
			"code":  200,
			"route": route,
			"msg":   msg,
		})
	}
}
