package adt

type StringArray []String

func NewStringArray(sa []string) StringArray {
	output := []String{}
	for _, str := range sa {
		output = append(output, String(str))
	}
	return output
}

// func (sa StringArray) Map(fn func(string) interface{}) StringArray {
// 	output := []string{}
// 	for _, str := range ([]string)(sa) {
// 		output = append(output, fn(str).(string))
// 	}
// 	return StringArray(output)
// }
