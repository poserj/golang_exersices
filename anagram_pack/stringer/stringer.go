package stringer
import ("sort"
       "strings"
       )
func AreAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
    str1 = strings.ToLower(str1)
    str2 = strings.ToLower(str2)
	str1Array := []rune(str1)
	sort.Slice(str1Array, func(i, j int) bool {
		return str1Array[i] < str1Array[j]
	})
	str2Array := []rune(str2)
	sort.Slice(str2Array, func(i, j int) bool {
		return str2Array[i] < str2Array[j]
	})
	    for i := 0; i < len(str1Array); i++ {
		if str1Array[i] != str2Array[i] {
			return false
		}
	}
	
	
	return true

}
