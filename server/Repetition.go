package server

import (
	"encoding/csv"
	"fmt"
	"lucky/common/logger"
	"os"
	"strings"

	"go.uber.org/zap"
)

// LotteryRecord /**
type LotteryRecord struct {
	Issue string
	Num1  string
	Num2  string
	Num3  string
	Num4  string
	Num5  string
	Num6  string
	Num7  string
}

func Calculate(code string, number []interface{}) []LotteryRecord {
	filePath := "../data/" + code + ".csv"
	file, err := os.Open(filePath)
	if err != nil {
		logger.Log.Error("无法打开文件:", zap.Error(err))
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Log.Error("无法关闭文件:", zap.Error(err))
		}
	}(file)

	// 创建CSV Reader
	reader := csv.NewReader(file)

	// 读取CSV文件的内容
	records, err := reader.ReadAll()
	if err != nil {
		logger.Log.Error("无法读取CSV:", zap.Error(err))
		return nil
	}

	// 创建用于存储匹配数据的切片
	var matches []LotteryRecord

	// 遍历CSV记录
	for _, record := range records {
		// 创建LotteryRecord结构体并填充数据
		lottery := LotteryRecord{
			Issue: record[0],
			Num1:  record[1],
			Num2:  record[2],
			Num3:  record[3],
			Num4:  record[4],
			Num5:  record[5],
			Num6:  record[6],
			Num7:  record[7],
		}

		// 将读取的号码转换成字符串切片
		var inputNumbers []string
		for _, num := range number {
			inputNumbers = append(inputNumbers, fmt.Sprintf("%v", num))
		}

		// 将CSV记录中的号码拼接成字符串
		recordNumbers := []string{
			lottery.Num1,
			lottery.Num2,
			lottery.Num3,
			lottery.Num4,
			lottery.Num5,
			lottery.Num6,
			lottery.Num7,
		}

		recordNumbersStr := strings.Join(recordNumbers, ",")
		matchingCount := 0
		for _, num := range inputNumbers {
			if strings.Contains(recordNumbersStr, num) {
				matchingCount++
			}
		}

		// 如果有匹配的号码，将记录添加到匹配切片中
		if matchingCount > 4 {
			matches = append(matches, lottery)
		}
	}

	return matches
}
