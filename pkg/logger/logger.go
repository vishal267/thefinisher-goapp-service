package logger

import (
    "log"
    "os"
)

var InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

