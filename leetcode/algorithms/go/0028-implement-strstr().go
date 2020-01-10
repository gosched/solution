package leetcode

func strStr(haystack string, needle string) int {
    length := len(haystack)
    l := len(needle)
    
    if l == 0 {
        return 0
    }
    
    if length == 0 || length < l {
        return -1
    }
    
    for i := 0; i < length; i++ {
        if (i + l) <= length && haystack[i:i+l] == needle {
            return i
        }
    }
    
    return -1
}

// slice[i:length]
// slice := "12345"
// length == 5
// target == "345"
// fmt.Println(slice[2:2+3]) // 345
// fmt.Println(slice[2:5])   // 345
// fmt.Println(slice[2:6])   // slice bounds out of range