package logger

import (
	"encoding/json"
	"github.com/gookit/color"
	"log"
	"os"
	"strings"
)

func Log(level string, data interface{}) {
	var logger *log.Logger
	switch strings.ToUpper(level) {
	case "WARN":
		logger = log.New(os.Stderr, color.Yellow.Render("[WARN] "), 0)
	case "FATAL":
		logger = log.New(os.Stderr, color.Red.Render("[FATAL] "), 0)
	case "DEBUG":
		logger = log.New(os.Stderr, color.Blue.Render("[DEBUG] "), 0)
	case "ERROR":
		logger = log.New(os.Stderr, color.Red.Render("[ERR] "), 0)
	default:
		logger = log.New(os.Stderr, color.Cyan.Render("[INFO] "), 0)
	}
	line, _ := json.Marshal(data)
	logger.Println(string(line))
}
