package toolbox

import json "github.com/json-iterator/go"

func ToByte(msg interface{}) []byte {
	data, _ := json.MarshalIndent(msg, "", "  ")
	return data
}

// func (msg *deviceReportData) Bytes() []byte {
// 	// data, _ := json.Marshal(msg)
// 	data, _ := json.MarshalIndent(msg, "", "  ")
// 	return data
// }
