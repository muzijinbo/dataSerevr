package demo

import (
	//"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	//"log"
	/*	"github.com/xuri/excelize"*/
	"logging"
	"net/http"
	"net/url"
	/*	"os"*/
	"strconv"
	"strings"
	"time"
	"webcrawler/analyzer"
	base "webcrawler/base"
	pipeline "webcrawler/itempipeline"
	sched "webcrawler/scheduler"
	"webcrawler/tool"
	toolcr "webcrawler/tool/createResult"
)

// 日志记录器。
var logger logging.Logger = logging.NewSimpleLogger()

// 条目处理器。
func processItem(item base.Item) (result base.Item, err error) {
	if item == nil {
		return nil, errors.New("Invalid item!")
	}
	// 生成结果
	result = make(map[string]interface{})
	for k, v := range item {
		result[k] = v
	}
	if _, ok := result["number"]; !ok {
		result["number"] = len(result)
	}
	time.Sleep(10 * time.Millisecond)
	return result, nil
}

/*
// 条目处理器。
func processItem(item base.Item) (result base.Item, err error) {
	if item == nil {
		return nil, errors.New("Invalid item!///无效的条目")
	}

	//fmt.Println(item)
	index, ok1 := item["index"].(int)
	sText, ok2 := item["text"].(string)
	sIndex := strconv.Itoa(index + 1)
	if ok1 && ok2 {
		mystring := sIndex + " " + sText + "\n"
		toolcr.WriteStringInFile("mydata1.txt", mystring)
	}
	//time.Sleep(10 * time.Millisecond)
	return result, nil
}*/
func processItem2(item base.Item) (result base.Item, err error) {
	if item == nil {
		return nil, errors.New("Invalid item!///无效的条目")
	}

	//fmt.Println(item)
	index, ok1 := item["index"].(int)
	sText, ok2 := item["text"].(string)
	if ok1 && ok2 {
		sIndex := strconv.Itoa(index + 1)
		mystringA := "A" + sIndex
		mystringB := "B" + sIndex
		toolcr.WriteIntInXlsxFile("mydata.xlsx", mystringA, index+1)
		toolcr.WriteStringInXlsxFile("mydata.xlsx", mystringB, sText)
	}
	//time.Sleep(10 * time.Millisecond)
	return result, nil
}

// 响应解析函数。只解析“A”标签。
func parseForATag(httpResp *http.Response, respDepth uint32) ([]base.Data, []error) {
	// TODO 支持更多的HTTP响应状态
	if httpResp.StatusCode != 200 {
		err := errors.New(
			fmt.Sprintf("Unsupported status code %d. (httpResponse=%v)", httpResp))
		return nil, []error{err}
	}
	var reqUrl *url.URL = httpResp.Request.URL
	var httpRespBody io.ReadCloser = httpResp.Body
	defer func() {
		if httpRespBody != nil {
			httpRespBody.Close()
		}
	}()
	dataList := make([]base.Data, 0)
	errs := make([]error, 0)
	// 开始解析
	doc, err := goquery.NewDocumentFromReader(httpRespBody)
	if err != nil {
		errs = append(errs, err)
		return dataList, errs
	}
	// 查找“A”标签并提取链接地址
	doc.Find("a").Each(func(index int, sel *goquery.Selection) {
		href, exists := sel.Attr("href")
		// 前期过滤
		if !exists || href == "" || href == "#" || href == "/" {
			return
		}
		href = strings.TrimSpace(href)
		lowerHref := strings.ToLower(href)
		// 暂不支持对Javascript代码的解析。
		if href != "" && !strings.HasPrefix(lowerHref, "javascript") {
			aUrl, err := url.Parse(href)
			if err != nil {
				errs = append(errs, err)
				return
			}
			if !aUrl.IsAbs() {
				//本方法根据一个绝对URI将一个URI补全为一个绝对URI，参见RFC 3986 节 5.2。参数ref可以是绝对URI或者相对URI。ResolveReference总是返回一个新的URL实例，即使该实例和u或者ref完全一样。如果ref是绝对URI，本方法会忽略参照URI并返回ref的一个拷贝。
				aUrl = reqUrl.ResolveReference(aUrl)
			}
			httpReq, err := http.NewRequest("GET", aUrl.String(), nil)
			if err != nil {
				errs = append(errs, err)
			} else {
				req := base.NewRequest(httpReq, respDepth)
				dataList = append(dataList, req)
			}
		}
		text := strings.TrimSpace(sel.Text())
		if text != "" {
			imap := make(map[string]interface{})
			//srequrl := reqUrl.String()
			imap["parent_url"] = reqUrl.String()
			imap["a.text"] = text
			imap["a.index"] = index
			item := base.Item(imap)
			dataList = append(dataList, &item)
		}
	})
	return dataList, errs
}

// 响应解析函数。只解析“A”标签。
func parseForATag2(httpResp *http.Response, respDepth uint32) ([]base.Data, []error) {
	// TODO 支持更多的HTTP响应状态
	if httpResp.StatusCode != 200 {
		err := errors.New(
			fmt.Sprintf("Unsupported status code %d. (httpResponse=%v)", httpResp))
		return nil, []error{err}
	}
	//var reqUrl *url.URL = httpResp.Request.URL
	var httpRespBody io.ReadCloser = httpResp.Body
	defer func() {
		if httpRespBody != nil {
			httpRespBody.Close()
		}
	}()
	dataList := make([]base.Data, 0)
	errs := make([]error, 0)
	// 开始解析
	doc, err := goquery.NewDocumentFromReader(httpRespBody)
	if err != nil {
		errs = append(errs, err)
		return dataList, errs
	}
	// 查找“A”标签并提取链接地址
	doc.Find(".topics .topic").Each(func(index int, sel *goquery.Selection) {
		text1 := sel.Find(".title a").Text()
		//log.Println("第", index+1, "个帖子的标题：", text1)
		// 前期过滤
		text := strings.TrimSpace(text1)
		if text != "" {
			imap := make(map[string]interface{})
			imap["index"] = index
			imap["text"] = text
			item := base.Item(imap)
			dataList = append(dataList, &item)
		}
	})
	return dataList, errs
}

// 获得响应解析函数的序列。
func getResponseParsers() []analyzer.ParseResponse {
	parsers := []analyzer.ParseResponse{
		//parseForATag,
		parseForATag2,
	}
	return parsers
}

// 获得条目处理器的序列。
func getItemProcessors() []pipeline.ProcessItem {
	itemProcessors := []pipeline.ProcessItem{
		processItem,
		//processItem2,
	}
	return itemProcessors
}

// 生成HTTP客户端。
func genHttpClient() *http.Client {
	return &http.Client{}
}

func record(level byte, content string) {
	if content == "" {
		return
	}
	switch level {
	case 0:
		logger.Infoln(content)
	case 1:
		logger.Warnln(content)
	case 2:
		logger.Infoln(content)
	}
}

func Crewl(address string) {
	// 创建调度器
	scheduler := sched.NewScheduler()

	// 准备监控参数
	intervalNs := 10 * time.Millisecond
	maxIdleCount := uint(1000)
	// 开始监控
	checkCountChan := tool.Monitoring(
		scheduler,
		intervalNs,
		maxIdleCount,
		true,
		false,
		record)

	// 准备启动参数
	channelArgs := base.NewChannelArgs(10, 10, 10, 10)
	poolBaseArgs := base.NewPoolBaseArgs(3, 3)
	crawlDepth := uint32(1)
	httpClientGenerator := genHttpClient
	respParsers := getResponseParsers()
	itemProcessors := getItemProcessors()
	//startUrl := "http://studygolang.com/topics"
	startUrl := address
	firstHttpReq, err := http.NewRequest("GET", startUrl, nil)
	if err != nil {
		logger.Errorln(err)
		return
	}
	// 开启调度器
	scheduler.Start(
		channelArgs,
		poolBaseArgs,
		crawlDepth,
		httpClientGenerator,
		respParsers,
		itemProcessors,
		firstHttpReq)

	// 等待监控结束
	<-checkCountChan
}
