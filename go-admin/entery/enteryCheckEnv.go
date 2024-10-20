package entery

import (
	"encoding/json"
	"go-admin/model"
	"go-admin/tools"
	"go-admin/ui"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

var (
	config     = "./config.json"
	JsonEntery model.StructEntery
)

// 检查环境必备条件
func EnteryCheckEnv() {

	// 配置文件如果不存在则需要用户输入
	if !tools.ToolsIsFileExist(config) {
		getUserInput()
	} else {
		// 否则自动从文件读取配置
		readConfig()
	}

	// 给用户看看参数内容有没有问题
	ui.PrintInfo("后端端口: " + JsonEntery.ServicePort)
	ui.PrintInfo("URL指定: " + JsonEntery.Address)
}

// 获取用户输入
func getUserInput() {

	// 询问用户输入第一个端口号
	survey.AskOne(&survey.Input{Message: "请指定服务端口", Default: "9099"}, &JsonEntery.ServicePort)

	// 服务器域名或者ipv4地址，必须提供
	survey.AskOne(&survey.Input{Message: "请指定域名(可提供IPV4)", Default: "localhost"}, &JsonEntery.Address)

	value, err := json.Marshal(JsonEntery)
	if err != nil {
		ui.PrintError("配置序列化失败: " + err.Error())
		os.Exit(1)
	}

	// 将序列化后的数据写出到本地
	err = tools.ToolsWriteFile(config, string(value))
	if err != nil {
		ui.PrintError("配置写出到本地失败: " + err.Error())
		os.Exit(1)
	}
}

// 读取配置文件
func readConfig() {
	value, err := tools.ToolsReadFile(config)
	if err != nil {
		ui.PrintError("配置文件读取失败: " + err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(value), &JsonEntery)
	if err != nil {
		ui.PrintError("配置文件反序列化失败: " + err.Error())
		os.Exit(1)
	}
}
