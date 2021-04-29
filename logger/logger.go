package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ckhero/go-common/config"
	"github.com/ckhero/go-common/util/trace"
	"github.com/go-kit/kit/sd/lb"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"regexp"
	"runtime"
	"sync"
	"time"
)

var (
	logger      *logrus.Logger
	project     string
	application string
	loggerOnce  sync.Once
)

/**
 * 初始化日志组件
 */
func InitLogger(projectName string, applicationName string, config config.Logger) {
	loggerOnce.Do(func() {
		project, application = projectName, applicationName
		logger = logrus.New()
		level, _ := logrus.ParseLevel(config.Level)
		logger.Level = level
		logger.SetReportCaller(true)
		//控制台打印即可
		logger.SetOutput(os.Stdout)
		//json格式输出，PrettyPrint不能使用，否则k8s采集的有问题
		var re = regexp.MustCompile(`^dev-gitlab.wanxingrowth.com/`)
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   time.RFC3339Nano,
			//DisableTimestamp:  false,
			//DisableHTMLEscape: false,
			//DataKey:           "",
			//FieldMap:          nil,
			CallerPrettyfier:  func(f *runtime.Frame) (string, string) {
				fileName := path.Base(f.File)
				return fmt.Sprintf("%s()", re.ReplaceAllString(f.Function, "")), fmt.Sprintf("%s:%d", fileName, f.Line)
			},
			//PrettyPrint:       false,
		})
	})
}

func GetLogger(ctx context.Context) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"project":     project,
		"application": application,
		"traceId":     trace.GetTraceId(ctx),
	})
}

func GetLoggerWithBody(ctx context.Context, body interface{}) *logrus.Entry {
	logger := GetLogger(ctx)
	return logger.WithField("body", FormatMessage(body))
}

func GetLoggerWithErr(ctx context.Context, err error) *logrus.Entry {
	logger := GetLogger(ctx)
	return logger.WithError(err)
}

func GetLoggerWithRedisErr(ctx context.Context, err error) *logrus.Entry {
	logger := GetLogger(ctx)
	return logger.WithField("redis err", err.Error())
}

func LoggerRedisErr(ctx context.Context, err error) {
	logger := GetLogger(ctx)
	logger.WithField("err", err.Error()).Error("get redis conn fail !")
}

/**
 * 格式化消息
 */
func FormatMessage(body interface{}) string {
	switch body.(type) {
	case string:
		return body.(string)
	case lb.RetryError:
		return body.(lb.RetryError).Final.Error()
	default:
		logData, err := json.Marshal(body)
		if err != nil {
			logrus.Error("format log message error: ", err.Error())
			return ""
		}
		//转换一次格式
		var str bytes.Buffer
		_ = json.Indent(&str, logData, "", "    ")

		return str.String()
	}
}
