package main

import (
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/containers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	containers.Serve()

}
