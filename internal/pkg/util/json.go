package util

import "strings"

/*
		說明: 直接改json值的函式

		用法:
		var inter interface{}
		err = json.Unmarshal(marshal, &inter)
		if err != nil {
			return err
		}
		SetValueInJSON(inter, "欲修改的json key", 想改的值)

		範例:
		s := []byte(`{ "level1a":"aaa", "level1b":{"level2a":"bbb", "level2b":{"level3":"aaa"} } }`)
	    var j interface{}
	    json.Unmarshal(s, &j)
	    SetValueInJSON(j, "level1a", "new value 1a")
	    SetValueInJSON(j, "level1b.level2a", "new value 2a")
	    SetValueInJSON(j, "level1b.level2b.level3", "new value 3")
	    SetValueInJSON(j, "level1b.level2c", "new key")
	    s,_ = json.Marshal(j)
	    fmt.Println(string(s))
	    // result: {"level1a":"new value 1a","level1b":{"level2a":"new value 2a","level2b":{"level3":"new value 3"},"level2c":"new key"}}
*/
func SetValueInJSON(iface interface{}, path string, value interface{}) interface{} {
	m := iface.(map[string]interface{})
	split := strings.Split(path, ".")

	for k, v := range m {
		if strings.EqualFold(k, split[0]) {
			if len(split) == 1 {
				m[k] = value
				return m
			}
			switch v.(type) {
			case map[string]interface{}:
				return SetValueInJSON(v, strings.Join(split[1:], "."), value)
			default:
				return m
			}
		}
	}

	if len(split) == 1 {
		m[split[0]] = value
	} else {
		newMap := make(map[string]interface{})
		newMap[split[len(split)-1]] = value
		for i := len(split) - 2; i > 0; i-- {
			mTmp := make(map[string]interface{})
			mTmp[split[i]] = newMap
			newMap = mTmp
		}
		m[split[0]] = newMap
	}
	return m
}
