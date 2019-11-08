package main

import "github.com/grsmv/goweek"

import "fmt"

func days() {
	week, _ := goweek.NewWeek(2015, 46)
	fmt.Println(week.Days)

}
