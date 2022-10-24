package main

import (
	"encoding/json"
	"fmt"
	"github.com/drone/envsubst/v2"
	"io/ioutil"
	"k8s-study/core"
	"log"
	"os"
)

func setupConfig() {
	if err := PathExists(config); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	configJson, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	newJson, err := envsubst.EvalEnv(string(configJson))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(newJson), &core.AppConf)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// PathExists 检查文件夹是否存在
func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("file not found")
	}
	return err
}
