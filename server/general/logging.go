package general

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"path"
	"runtime"
	"strings"
)



const (
	LogRequestID = "go-microservice-request"
)


type BaseFormatter struct {
	Formatter        easy.Formatter
	CallerPrettyfier func(*runtime.Frame) (function string, file string)
}

func (f *BaseFormatter) Format(entry *log.Entry) ([]byte, error) {
	var funcVal, fileVal string

	if entry != nil {
		if f.CallerPrettyfier != nil {
			funcVal, fileVal = f.CallerPrettyfier(entry.Caller)
		} else {
			funcVal = entry.Caller.Function
			fileVal = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		}
		entry.Data["func"] = funcVal
		entry.Data["file"] = fileVal

		return f.Formatter.Format(entry)
	}

	return []byte{}, errors.New("not found configuration")
}

func SetupLogging() {
	log.SetReportCaller(true)
	log.SetFormatter(&BaseFormatter{
		Formatter: easy.Formatter{
			TimestampFormat: "2006-01-02T15:04:05.1483386-00:00",
			LogFormat:       "[%time%][%lvl%] [%file%][%func%] | [%requestID%] | %msg%\n",
		},
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			var funcVal, fileVal = "[not_found_function_name]", "[not_found_file_name]"
			if frame != nil {
				filename := path.Base(frame.File)
				function := strings.Split(fmt.Sprintf("%s()", frame.Function), "/")
				funcVal, fileVal = function[len(function)-1], fmt.Sprintf("%s:%d", filename, frame.Line)
			}

			return funcVal, fileVal
		},
	})
}


func GetContextFromEcho(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), LogRequestID, c.Response().Header().Get(echo.HeaderXRequestID))
}

func BuildContext() context.Context {
	return context.WithValue(context.Background(), LogRequestID, generateRequestId())
}

func ServerLog(c context.Context) *log.Entry {
	rawRequestId := c.Value(LogRequestID)
	if rawRequestId == nil {
		rawRequestId = generateRequestId()
	}
	return log.WithFields(log.Fields{
		"requestID": rawRequestId.(string),
	})
}

func generateRequestId() string {
	return fmt.Sprintf("al-%s", random.String(32))
}

func Log() *log.Entry {
	return log.WithFields(log.Fields{
		"requestID": "-",
	})
}


