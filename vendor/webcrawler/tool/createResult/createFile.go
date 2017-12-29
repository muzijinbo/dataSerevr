package createResult

import (
	"fmt"
	"github.com/xuri/excelize"
	"io"
	"os"
	"strconv"
	"sync"
)

func WriteStringInFile(fileName string, mystring string) error {
	_, err := os.Stat(fileName)
	var myFile *os.File
	if err != nil && os.IsNotExist(err) {
		myFile, err = os.Create(fileName)
		fmt.Println(fileName + "文件不存在，已经重新创建！")
	}
	myFile, err = os.OpenFile(fileName, os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("打开" + fileName + "文件失败")
	}
	_, err = io.WriteString(myFile, mystring)
	if err != nil {
		fmt.Println("写入文件失败！")
	}
	defer myFile.Close()
	return err
}
func WriteIntInFile(fileName string, myInt int) error {
	s := strconv.Itoa(myInt)
	err := WriteStringInFile(fileName, s)
	return err
}
func WriteBytesInFile(fileName string, mybytes []byte) error {
	_, err := os.Stat(fileName)
	var myFile *os.File
	if err != nil && os.IsNotExist(err) {
		myFile, err = os.Create(fileName)
		fmt.Println(fileName + "文件不存在，已经重新创建！")
	}
	myFile, err = os.OpenFile(fileName, os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("打开" + fileName + "文件失败")
	}
	_, err = myFile.Write(mybytes)

	if err != nil {
		fmt.Println("写入文件失败！")
	}
	defer myFile.Close()
	return err
}

func WriteStringInXlsxFile(fileName string, location string, mydata string) error {
	_, err := os.Stat(fileName)
	var xlsx *excelize.File

	if err != nil && os.IsNotExist(err) {
		xlsx = excelize.NewFile()
	} else {
		xlsx, err = excelize.OpenFile(fileName)
	}
	MyXlsFile := newMyXlsFile(xlsx)
	MyXlsFile.XlsFile.SetCellValue("Sheet1", location, mydata)
	fmt.Println(location + "_" + mydata)
	err = MyXlsFile.XlsFile.SaveAs(fileName)
	defer MyXlsFile.m.Unlock()
	return err
}

type MyXlsFile struct {
	XlsFile *excelize.File
	m       sync.Mutex
}

func newMyXlsFile(XlsFile *excelize.File) MyXlsFile {
	//XlsFile := excelize.NewFile()
	return MyXlsFile{XlsFile: XlsFile}
}
func WriteIntInXlsxFile(fileName string, location string, mydata int) error {
	_, err := os.Stat(fileName)
	var xlsx *excelize.File

	if err != nil && os.IsNotExist(err) {
		xlsx = excelize.NewFile()
	} else {
		xlsx, err = excelize.OpenFile(fileName)
	}
	MyXlsFile := newMyXlsFile(xlsx)
	MyXlsFile.m.Lock()
	MyXlsFile.XlsFile.SetCellValue("Sheet1", location, mydata)
	fmt.Println(location)
	err = MyXlsFile.XlsFile.SaveAs(fileName)
	defer MyXlsFile.m.Unlock()
	return err
}
