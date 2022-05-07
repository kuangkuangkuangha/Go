package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

// 处理 /upload  逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()                                    // 返回一个新的使用MD5校验的hash.Hash接口。
		io.WriteString(h, strconv.FormatInt(crutime, 10)) // formatint返回crutime的10进制字符串表示, 然后将字符串写入h中
		token := fmt.Sprintf("%x", h.Sum(nil))            // h.Sum返回nil的MD5校验和，然后以十六进制输出

		t, _ := template.ParseFiles("upload.gtpl")
		//ParseFiles方法解析匹配pattern的文件里的模板定义并将解析结果与t关联。
		//如果发生错误，会停止解析并返回nil，否则返回(t, nil)。至少要存在一个匹配的文件。

		t.Execute(w, token)
		// Execute方法将解析好的模板应用到data上，并将输出写入wr。
		// 如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行。

	} else {
		r.ParseMultipartForm(32 << 20)
		// ParseMultipartForm将请求的主体作为multipart/form-data解析。
		// 请求的整个主体都会被解析，得到的文件记录最多maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里。
		// 如果必要，ParseMultipartForm会自行调用ParseForm。重复调用本方法是无意义的。

		file, handler, err := r.FormFile("uploadfile")
		// FormFile返回以key为键查询r.MultipartForm字段得到结果中的第一个文件和它的信息。
		// 如果必要，本函数会隐式调用ParseMultipartForm和ParseForm。查询失败会返回ErrMissingFile错误。

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe("9090", nil)
}
