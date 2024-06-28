package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	SetLogger("log", logrus.TraceLevel)
	InitializeDatabase()
	StartBot()
	StartServer()
}
