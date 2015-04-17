package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

type ConfigStruct struct {
	Host   string   `json:"host"`
	Fruits []string `json:"fruits"`
	Othres []Other  `json:"others"`
}

type Other struct {
	SerTcpHost string `json:"serTcpHost"`
	SerTcpPort int    `json:"serTcpPort"`
}

var jsonStr = `{"host": "http://localhost:9090","fruits": ["apple", "peach"],"others":[{"serTcpHost":"serHost1","SerTcpPort":123},{"serTcpHost":"serHost2","SerTcpPort":456}]}`

func main() {
	//JsonStrToMap()
	JsonStrToStructAll()
}

func JsonStrToMap() {
	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
		fmt.Println(dat["host"])
		fmt.Println(dat["fruits"])

	}
}

//json str 转struct
func JsonStrToStructAll() {
	var config ConfigStruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(config)
		fmt.Println(config.Host)
		fmt.Println(config.Fruits)
		fmt.Println(config.Othres)
		for k, v := range config.Othres {
			fmt.Println(k)
			fmt.Println(v.SerTcpHost)
		}
	}
}

/*
//json str 转struct(部份字段)
func JsonStrToStruct() {
	var part Other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(part)
		fmt.Println(part.SerTcpSocketPort)
	}
}

func StructToJsonStr() {

	//struct 到json str
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))
	}
}

//map 到json str
func MapToJsonStr() {
	fmt.Println("================map 到json str=====================")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)
}

//array 到 json str
func ArrayToJsonStr() {
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}
}

//json 到 []string
func JsonToStrings() {
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(wo)
	}
}
*/
