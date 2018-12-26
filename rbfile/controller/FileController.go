package controller

import (
	R "../../rbstruct/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../rbutil"
	"../../rbwork/base"
	"net/http"
	"os"
	"io"
)



func Index(w http.ResponseWriter, r *http.Request)  {
	url :=  constant.DownloadHost+ r.RequestURI
	http.Redirect(w,r,url,http.StatusMovedPermanently)
}

func Upload(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	if r.Method == "POST" {
		err := r.ParseMultipartForm(102400)
		base.CheckErr(err)
		m := r.MultipartForm
		files := m.File["files"]
		for i, _ := range files {
			file, err := files[i].Open()
			if err != nil {
				hc.ReturnMsg(R.ErrorMsg("参数错误"))
				return
			}
			f, err := os.OpenFile(rbutil.FilePath+files[i].Filename, os.O_WRONLY|os.O_CREATE, 0666)
			base.CheckErr(err)
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