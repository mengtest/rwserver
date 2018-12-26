package base

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path"
	"time"
	"github.com/shiena/ansicolor"
	"os"
)



func Init(logDir string,logFileName string) {
	baseLogPaht := path.Join(logDir, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d",
		//rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24), // 设置文件清理前的最长保存时间
		rotatelogs.WithRotationTime(time.Hour*24), //日志切割时间间隔
		rotatelogs.WithRotationCount(1), //设置文件清理前最多保存的个数
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	},&log.TextFormatter{ForceColors: true,TimestampFormat:"yyyy-MM-dd HH:mm:ss.SSS"}) //开启颜色

	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.AddHook(lfHook)
}

func LogDebug(args ...interface{})  {
	log.Debug(args)
}

func LogInfo(args ...interface{})  {
	log.Info(args)
}

func LogError(args ...interface{})  {
	log.Error(args)
}

func LogFatal(args ...interface{})  {
	log.Fatal(args)
}

func LogErr(err error){
	if err != nil {
		log.Error(err)
	}
}
