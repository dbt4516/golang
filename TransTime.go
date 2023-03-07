package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {

	args := os.Args

	// 打印程序名称和所有参数
	fmt.Println("Program name:", args[0])
	fmt.Println("Arguments:", args[1:])


	// 打开文件
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建一个Scanner对象来逐行读取文件内容
	scanner := bufio.NewScanner(file)

	// 循环遍历文件中的每一行
	for scanner.Scan() {
		// 打印当前行的内容

		// 输入的字符串
		input := scanner.Text()

		// 编译正则表达式，匹配13位数字
		re := regexp.MustCompile(`\d{13}`)

		// 匹配所有时间戳
		matches := re.FindAllString(input, -1)

		// 循环遍历所有匹配到的时间戳
		for _, match := range matches {
			// 将时间戳转换为int64类型
			timestamp, err := strconv.ParseInt(match, 10, 64)
			if err != nil {
				fmt.Println("Error parsing timestamp:", err)
				continue
			}

			// 使用time包将时间戳转换为时间类型
			t := time.Unix(0, timestamp*int64(time.Millisecond))

			// 将时间转换为字符串
			dateString := t.Format("2006-01-02 15:04:05")

			// 替换原字符串中的时间戳
			input = re.ReplaceAllString(input, dateString)
		}

		// 输出替换后的字符串
		fmt.Println(input)
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

}
