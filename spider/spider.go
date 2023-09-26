package spider

import (
	"encoding/csv"
	_ "encoding/json"
	"fmt"
	"log"
	"lucky/common/logger"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gocolly/colly"
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

type Lottery struct {
	Issue        string // 开奖期数
	Num1         string // 开奖号码1
	Num2         string // 开奖号码2
	Num3         string // 开奖号码3
	Num4         string // 开奖号码4
	Num5         string // 开奖号码5
	Num6         string // 开奖号码6
	Num7         string // 开奖号码7
	Pool         string //
	ONumberNotes string //注数
	OBonus       string //奖金
	TNumberNotes string //注数
	TBonus       string //奖金
	TotalNotes   string //总注数
	Datatime     string //开奖时间
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
	fetchDataAndParse(config.DLT.URL, config.DLT.Start, config.DLT.End)
	//fetchDataAndParse(config.SSQ.URL, config.SSQ.Start, config.SSQ.End)
	//SSQ := fetchDataAndParse(config.SSQ.URL, config.SSQ.Start, config.SSQ.End)
	//fmt.Println(SSQ)

}

func New() {
	logger.Log.Info("app started")
}

// 请求URL并解析数据
func fetchDataAndParse(url, start, end string) {
	// 构建请求URL
	fullURL := url + "?start=" + start + "&end=" + end

	c := colly.NewCollector()

	// 处理响应
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Ctx.Get("url"))
	})

	var lotteries []Lottery
	c.OnHTML("tr[class=t_tr1]", func(e *colly.HTMLElement) {
		var l Lottery
		l.Issue = e.ChildText("td:nth-child(1)")
		l.Num1 = e.ChildText("td:nth-child(2)")
		l.Num2 = e.ChildText("td:nth-child(3)")
		l.Num3 = e.ChildText("td:nth-child(4)")
		l.Num4 = e.ChildText("td:nth-child(5)")
		l.Num5 = e.ChildText("td:nth-child(6)")
		l.Num6 = e.ChildText("td:nth-child(7)")
		l.Num7 = e.ChildText("td:nth-child(8)")
		l.Pool = e.ChildText("td:nth-child(9)")
		l.ONumberNotes = e.ChildText("td:nth-child(10)")
		l.OBonus = e.ChildText("td:nth-child(11)")
		l.TNumberNotes = e.ChildText("td:nth-child(12)")
		l.TBonus = e.ChildText("td:nth-child(13)")
		l.TotalNotes = e.ChildText("td:nth-child(14)")
		l.Datatime = e.ChildText("td:nth-child(15)")
		lotteries = append(lotteries, l)
		SaveToCSV(lotteries)
		// lotteries保存为 data.csv

	})

	// 设置抓取参数
	c.SetRequestTimeout(100 * time.Second)

	// 开始访问
	err := c.Visit(fullURL)
	if err != nil {
		logger.Log.Error("访问失败:", zap.Error(err))
	}

	c.Wait()

}

func SaveToCSV(lotteries []Lottery) {

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入头部
	header := []string{"Issue", "Num1", "Num2", "Num3", "Num4", "Num5", "Num6", "Num7", "Pool", "ONumberNotes", "OBonus", "TNumberNotes", "TBonus", "TotalNotes", "Datatime"}
	err = writer.Write(header)
	if err != nil {
		log.Fatalln("Error writing record to csv:", err)

	}

	// 写入数据
	for _, l := range lotteries {
		data := []string{l.Issue, l.Num1, l.Num2, l.Num3, l.Num4, l.Num5, l.Num6, l.Num7, l.Pool, l.ONumberNotes, l.OBonus, l.TNumberNotes, l.TBonus, l.TotalNotes, l.Datatime}
		err := writer.Write(data)
		if err != nil {
			log.Fatalln("Error writing record to csv:", err)

		}
	}
}
