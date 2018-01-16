package basic

import (
	"fmt"
	"encoding/json"
)

type MapObject struct {
	Code     string
	IntValue int
}

func RunMapTest() {
	/*	intmap := map[string]int {
			"a" : 1,
			"b" : 2,
		}*/
	objmap := map[string]MapObject{
		"a": MapObject{"a", 1},
		"b": MapObject{"b", 2},
	}
	printmap(objmap)

	for key, _ := range objmap {
		value := objmap[key]
		value.Code = value.Code + "changed"
	}
	printmap(objmap)

}

func printmap(mapdata map[string]MapObject) {
	mapjson, _ := json.Marshal(mapdata)
	fmt.Printf("map data: %s\n", string(mapjson))
}
