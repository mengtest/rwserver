package controller

import (
	R "../../rbstruct/base"
	TQC "../../rbwork/constant"
	"net/http"
	"../../rbwork/network"
	"../rbutil"
	"os"
	"io"
)

func checkErr(err error) {
	if err != nil {
		err.Error()
	}
}

func Index(w http.ResponseWriter, r *http.Request)  {
	url :=  TQC.DownloadHost+ r.RequestURI
	http.Redirect(w,r,url,http.StatusMovedPermanently)
}

func Upload(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	if r.Method == "POST" {
		err := r.ParseMultipartForm(102400)
		checkErr(err)
		m := r.MultipartForm
		files := m.File["files"]
		for i, _ := range files {
			file, err := files[i].Open()
			if err != nil {
				hc.ReturnMsg(R.ErrorMsg("参数错误"))
				return
			}
			f, err := os.OpenFile(rbutil.FilePath+files[i].Filename, os.O_WRONLY|os.O_CREATE, 0666)
			checkErr(err)
			io.Copy(f, file)
			defer f.Close()
			defer file.Close()
		}
	}else{
		hc.ReturnMsg(R.ErrorMsg("不支持该请求方式"))
		return
	}

	hc.ReturnMsg(R.OK())
}


//下载指定文件
func Download(w http.ResponseWriter, r *http.Request)  {
	rbutil.StaticHandler.ServeHTTP(w,r)
	return
}