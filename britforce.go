package main
import (
"fmt"
)
func BruteForceSearch(text string, pattern string) int {
j := 0
n := len(text)
m := len(pattern)
for i := 0; i <= n-m; i++ {
j = 0
for j < m && pattern[j] == text[i+j] {
j++
}
if j == m {
return i
}
}
return -1
}
