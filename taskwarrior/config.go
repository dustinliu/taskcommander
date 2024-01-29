package warrior

//import (
//"strings"
//)

// var config *TaskWarriorConfig

//const (
//dataLocationKey = "data.location"
//)

//func init() {
//config = &TaskWarriorConfig{}
//data, err := Taskwarrior("show")
//if err != nil {
//panic(err)
//}

//lines := strings.Split(string(data), "\n")
//for _, line := range lines {
//if strings.Contains(line, dataLocationKey) {
//str, _ := strings.CutPrefix(line, dataLocationKey)
//config.Data_location = strings.TrimSpace(str)
//}
//}
//}

//type TaskWarriorConfig struct {
//Data_location string
//}

//func GetTaskWarriorConfig() *TaskWarriorConfig {
//return config
//}
