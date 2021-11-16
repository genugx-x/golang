package testlib

import "fmt"

var kpop map[string]string

func init() {
	kpop = make(map[string]string)
	kpop["CRUSH"] = "어떻게 지내"
	kpop["JANNABI"] = "뜨거운 여름은 가고 남은 건 볼품없지만"
	kpop["IU"] = "밤편지 (Through the Night)"
}

func GetMusic(singer string) string {
	return kpop[singer]
}

func getKeys() {
	for _, kv := range kpop {
		fmt.Println(kv)
	}
}
