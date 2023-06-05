package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	userinterface "shortTermStrategy/userinterface/facade"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

func main() {
	r := gin.Default()
	userinterface.ServiceGroupApp(r)
	var base Config
	data, _ := ioutil.ReadFile("config.yaml")
	_ = yaml.Unmarshal(data, &base)
	p := fmt.Sprintf(":%d", base.Server.Port)
	err := r.Run(p)
	if err != nil {
		log.Fatalln(err)
	}

}
