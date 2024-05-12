package main

import "github.com/juanmiguelar/gin_study/entrypoints"

func main() {
	entrypoints.CreateRoutes().Run()
}
