package leetcode

import "strings"

func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    if len(str) == 0 {
        return 0
    }
    
    var negative = false
    
    switch(str[0]) {
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        case '+':
            str = str[1:]
        case '-':
            str = str[1:]
            negative = true
        default:
            return 0
    }
    
    var max = 1<<31 - 1
    var min = -1<<31
    
    var answer = 0
    
    for _, r := range str {
        if r < '0' || r > '9' {
            break
        }
        
        answer = answer * 10 + int(r-'0')
        
        if !negative && answer > max {
            return max
        }
        
        if negative && -answer < min {
            return min
        }
    }
    
    if negative {
        answer = -answer
    }
    
    return answer
}