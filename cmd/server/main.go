package main

import "github.com/web_zktnotify/internal/app/initialize"

func main(){
	initialize.SetupConfig()
	initialize.SetupServer()
}
