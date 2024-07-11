package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

type Clazz struct {
	Id   int
	Name string
}

func main() {
	movie := Movie{"喜剧之王", 2000, 23, []string{"周星驰", "张柏芝"}}

	//编码的过程 结构体-->json
	marshal, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal is failed.", err)
		return
	}

	fmt.Printf("movie marshal = %s\n", marshal)

	//新版go可以直接识别fieldName-->json
	clazz := Clazz{1, "3班"}
	jsonStr, err := json.Marshal(clazz)
	if err != nil {
		fmt.Println("json marshal is failed.", err)
		return
	}
	fmt.Printf("clazz marshal = %s\n", jsonStr)

	//解码的过程 str -->结构体
	//str =  {"title":"喜剧之王","year":2000,"price":23,"actors":["周星驰","张柏芝"]}
	my_movie := Movie{}
	err = json.Unmarshal(marshal, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal is failed.", err)
		return
	}
	fmt.Printf("movie marshal = %v\n", my_movie)
}
