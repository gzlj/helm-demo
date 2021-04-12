package main

import (

	"fmt"
	"github.com/fatih/structs"
	"github.com/ghodss/yaml"
)

func main() {

	u1 := UserInfo{Name: "奇奇", Age: 18,
		Profile: Profile{"双色球"},
		Svcs: map[string]interface{}{
			"cas": SvcOverrideConfig{
				Image: ImageConfig{
					Repository: "demo.io/cas",
					PullPolicy: "Always",
				},
			},
		},

	}
	m3 := structs.Map(&u1)
	for k, v := range m3 {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
	//b, _ := json.Marshal(&u1)
	b, _ := yaml.Marshal(&u1)
	fmt.Println()
	var m map[string]interface{}
	_ = yaml.Unmarshal(b, &m)
	for k, v := range m {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}

}

// UserInfo 用户信息
type UserInfo struct {
	Name string `json:"name" structs:"name"`
	Age  int    `json:"age" structs:"age"`
	Profile `json:"profile" structs:"profile"`
	Svcs map[string]interface{}
}

type SvcOverrideConfig struct {
	Image ImageConfig `json:"image,omitempty structs:"image,omitempty"`
}

type ImageConfig struct {
	Repository string `json:"repository,omitempty structs:"repository,omitempty"`
	PullPolicy string `json:"pullPolicy,omitempty" structs:"pullPolicy,omitempty"`
}

// Profile 配置信息
type Profile struct {
	Hobby string `json:"hobby" structs:"hobby"`
}

