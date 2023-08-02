package spider

import (
	"fmt"
	"lucky/common/logger"
	"os"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

/**
 * config实体
 */

type Config struct {
	DLT struct {
		URL     string `toml:"url"`
		Name    string `toml:"name"`
		Code    string `toml:"code"`
		Period  int    `toml:"period"`
		Start   string `toml:"start"`
		End     string `toml:"end"`
		Page    int    `toml:"page"`
		Sleep   int    `toml:"sleep"`
		Timeout int    `toml:"timeout"`
	} `toml:"dlt"`

	SSQ struct {
		URL     string `toml:"url"`
		Name    string `toml:"name"`
		Code    string `toml:"code"`
		Period  int    `toml:"period"`
		Start   string `toml:"start"`
		End     string `toml:"end"`
		Page    int    `toml:"page"`
		Sleep   int    `toml:"sleep"`
		Timeout int    `toml:"timeout"`
	} `toml:"ssq"`

	KL8 struct {
		URL     string `toml:"url"`
		Name    string `toml:"name"`
		Code    string `toml:"code"`
		Period  int    `toml:"period"`
		Start   string `toml:"start"`
		End     string `toml:"end"`
		Page    int    `toml:"page"`
		Sleep   int    `toml:"sleep"`
		Timeout int    `toml:"timeout"`
	} `toml:"kl8"`

	QXC struct {
		URL     string `toml:"url"`
		Name    string `toml:"name"`
		Code    string `toml:"code"`
		Period  int    `toml:"period"`
		Start   string `toml:"start"`
		End     string `toml:"end"`
		Page    int    `toml:"page"`
		Sleep   int    `toml:"sleep"`
		Timeout int    `toml:"timeout"`
	} `toml:"qxc"`
}

/**
 * 爬虫模块，读取config.toml文件配置信息，初始化爬虫
 */

func init() {
	var config Config
	// 读取配置文件
	data, err := os.ReadFile("spider/config.toml")
	if err != nil {
		logger.Log.Error("无法读取配置文件:", zap.Error(err))
		return
	}

	// 解析配置文件
	if _, err := toml.Decode(string(data), &config); err != nil {
		logger.Log.Error("无法解析配置文件:", zap.Error(err))
		return
	}
	fmt.Println(config)
}

/**
 * 定义爬虫实体
 */

func New() {
	logger.Log.Info("app started")
}
