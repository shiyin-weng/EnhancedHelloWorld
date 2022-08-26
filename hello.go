package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	argsWithoutProg := os.Args[1:]
	envParamMap := make(map[string]string)
	name := "World"
	greetingMsg := "Hello"

	if os.Getenv("HELLO_NAME") != "" {
		envParamMap["HELLO_NAME"] = os.Getenv("HELLO_NAME")
		name = os.Getenv("HELLO_NAME")
	}

	if os.Getenv("HELLO_MSG") != "" {
		envParamMap["HELLO_MSG"] = os.Getenv("HELLO_MSG")
		greetingMsg = os.Getenv("HELLO_MSG")
	}

	// check the env params and line args existance
	// 1. no env vars or command line args
	if len(argsWithoutProg) == 0 && len(envParamMap) == 0 {
		sugar.Infof("No input environment variables or command line args")
		fmt.Printf("%s, %s!\n", greetingMsg, name)

		return
	}

	// 2. only command line args
	if len(argsWithoutProg) > 0 && len(envParamMap) == 0 {
		sugar.Infof("Only command line args")
		fmt.Printf("%s, %s!\n", greetingMsg, argsWithoutProg[0])

		return
	}

	// 3. only environment variables
	if len(argsWithoutProg) == 0 && len(envParamMap) > 0 {
		sugar.Infof("Only environment variables")
		fmt.Printf("%s, %s!\n", greetingMsg, name)

		return
	}
}
