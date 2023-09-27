package spider

import (
	"encoding/csv"
	_ "encoding/json"
	"lucky/common/logger"
	"os"
	"strconv"
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
		URL   string `toml:"url"`
		Name  string `toml:"name"`
		Code  string `toml:"code"`
		Start string `toml:"start"`
		End   string `toml:"end"`
	} `toml:"dlt"`

	SSQ struct {
		URL   string `toml:"url"`
		Name  string `toml:"name"`
		Code  string `toml:"code"`
		Start string `toml:"start"`
		End   string `toml:"end"`
	} `toml:"ssq"`

	KL8 struct {
		URL   string `toml:"url"`
		Name  string `toml:"name"`
		Code  string `toml:"code"`
		Start string `toml:"start"`
		End   string `toml:"end"`
	} `toml:"kl8"`

	QXC struct {
		URL   string `toml:"url"`
		Name  string `toml:"name"`
		Code  string `toml:"code"`
		Start string `toml:"start"`
		End   string `toml:"end"`
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
	if LastNumber(config.DLT.Code) == 0 {
		fetchDataAndParse(config.DLT.URL, config.DLT.Start, config.DLT.End, config.DLT.Code)
	} else {
		fetchDataAndParse(config.DLT.URL, strconv.Itoa(LastNumber(config.DLT.Code)+1), config.DLT.End, config.DLT.Code)
	}
	if LastNumber(config.SSQ.Code) == 0 {
		fetchDataAndParse(config.SSQ.URL, config.SSQ.Start, config.SSQ.End, config.SSQ.Code)
	} else {
		fetchDataAndParse(config.SSQ.URL, strconv.Itoa(LastNumber(config.SSQ.Code)+1), config.SSQ.End, config.SSQ.Code)
	}
	//fetchDataAndParse(config.DLT.URL, config.DLT.Start, config.DLT.End, config.DLT.Code)
	//fetchDataAndParse(config.SSQ.URL, config.SSQ.Start, config.SSQ.End, config.SSQ.Code)

}

func New() {}

// 请求URL并解析数据
func fetchDataAndParse(url, start, end, code string) {
	// 构建请求URL
	fullURL := url + "?start=" + start + "&end=" + end

	c := colly.NewCollector()

	// 处理响应
	c.OnResponse(func(r *colly.Response) {
		logger.Log.Info("Visited", zap.String(fullURL, r.Ctx.Get(fullURL)))
	})

	var lotteries []Lottery
	if code == "dlt" {
		c.OnHTML("tbody[id=tdata]", func(e *colly.HTMLElement) {

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				var l Lottery
				l.Issue = el.DOM.Find("td:nth-child(1)").Text()
				l.Num1 = el.DOM.Find("td:nth-child(2)").Text()
				l.Num2 = el.DOM.Find("td:nth-child(3)").Text()
				l.Num3 = el.DOM.Find("td:nth-child(4)").Text()
				l.Num4 = el.DOM.Find("td:nth-child(5)").Text()
				l.Num5 = el.DOM.Find("td:nth-child(6)").Text()
				l.Num6 = el.DOM.Find("td:nth-child(7)").Text()
				l.Num7 = el.DOM.Find("td:nth-child(8)").Text()
				l.Pool = el.DOM.Find("td:nth-child(9)").Text()
				l.ONumberNotes = el.DOM.Find("td:nth-child(10)").Text()
				l.OBonus = el.DOM.Find("td:nth-child(11)").Text()
				l.TNumberNotes = el.DOM.Find("td:nth-child(12)").Text()
				l.TBonus = el.DOM.Find("td:nth-child(13)").Text()
				l.TotalNotes = el.DOM.Find("td:nth-child(14)").Text()
				l.Datatime = el.DOM.Find("td:nth-child(15)").Text()
				lotteries = append(lotteries, l)
				SaveToCSV(lotteries, code)
			})
		})

	}
	if code == "ssq" {
		c.OnHTML("tbody[id=tdata]", func(e *colly.HTMLElement) {
			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				var l Lottery
				l.Issue = el.DOM.Find("td:nth-child(1)").Text()
				l.Num1 = el.DOM.Find("td:nth-child(2)").Text()
				l.Num2 = el.DOM.Find("td:nth-child(3)").Text()
				l.Num3 = el.DOM.Find("td:nth-child(4)").Text()
				l.Num4 = el.DOM.Find("td:nth-child(5)").Text()
				l.Num5 = el.DOM.Find("td:nth-child(6)").Text()
				l.Num6 = el.DOM.Find("td:nth-child(7)").Text()
				l.Num7 = el.DOM.Find("td:nth-child(8)").Text()
				l.Pool = el.DOM.Find("td:nth-child(10)").Text()
				l.ONumberNotes = el.DOM.Find("td:nth-child(11)").Text()
				l.OBonus = el.DOM.Find("td:nth-child(12)").Text()
				l.TNumberNotes = el.DOM.Find("td:nth-child(13)").Text()
				l.TBonus = el.DOM.Find("td:nth-child(14)").Text()
				l.TotalNotes = el.DOM.Find("td:nth-child(15)").Text()
				l.Datatime = el.DOM.Find("td:nth-child(16)").Text()
				lotteries = append(lotteries, l)
				SaveToCSV(lotteries, code)
			})
		})
	}

	// 设置抓取参数
	c.SetRequestTimeout(100 * time.Second)

	// 开始访问
	err := c.Visit(fullURL)
	if err != nil {
		logger.Log.Error("访问失败:", zap.Error(err))
	}

	c.Wait()

}

func SaveToCSV(lotteries []Lottery, code string) {

	var file *os.File
	var err error
	if code == "ssq" {
		file, err = os.Create("data/ssq.csv")
		if err != nil {
			logger.Log.Error("无法创建文件ssq.csv:", zap.Error(err))
		}
	} else if code == "dlt" {
		file, err = os.Create("data/dlt.csv")
		if err != nil {
			logger.Log.Error("无法创建文件dlt.csv:", zap.Error(err))
		}
	}
	if err != nil {
		logger.Log.Error("无法创建文件:", zap.Error(err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Log.Error("无法关闭文件:", zap.Error(err))
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入头部
	header := []string{"Issue", "Num1", "Num2", "Num3", "Num4", "Num5", "Num6", "Num7", "Pool", "ONumberNotes", "OBonus", "TNumberNotes", "TBonus", "TotalNotes", "Datatime"}
	err = writer.Write(header)
	if err != nil {
		logger.Log.Error("Error writing record to csv:", zap.Error(err))
	}

	// 写入数据
	for _, l := range lotteries {
		data := []string{l.Issue, l.Num1, l.Num2, l.Num3, l.Num4, l.Num5, l.Num6, l.Num7, l.Pool, l.ONumberNotes, l.OBonus, l.TNumberNotes, l.TBonus, l.TotalNotes, l.Datatime}
		err := writer.Write(data)
		if err != nil {
			logger.Log.Error("Error writing record to csv:", zap.Error(err))
		}
	}
}

/**
* 读取 csv 文件
* @param code string
* @return int
 */
func LastNumber(code string) int {
	// 读取“data/{code}.csv”文件,获取第二行的第一个单元格的值
	// 拼接csv文件路径
	csvFile := "data/" + code + ".csv"

	// 打开文件
	f, err := os.Open(csvFile)
	if err != nil {
		logger.Log.Error("无法打开文件:", zap.Error(err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logger.Log.Error("无法关闭文件:", zap.Error(err))
		}
	}(f)

	// 读取CSV
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		logger.Log.Error("无法读取CSV:", zap.Error(err))
	}

	// 获取第二行数据
	num, err := strconv.Atoi(records[1][0])
	if err != nil {
		logger.Log.Error("无数据:", zap.Error(err))
	}
	return num
}
