/*
 Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:

Input: "aba"
Output: True

Example 2:

Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
*/

func isPalindrome(s string) bool {
    if len(s) == 0 {
        return false
    }
    
    middleIndex := len(s) / 2
    rev := len(s[:middleIndex]) -1
    for _, v := range s[middleIndex+1:] {
        if rune(s[rev]) != v {
            return false
        }
        rev = rev - 1
    }
    return true
}

func validPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    
    if len(s) % 2 == 0 {
        for i, _ := range s {
            if i == 0 && isPalindrome(s[i+1:]) {
                    return true
            } else if i == len(s) - 1 && isPalindrome(s[0:i]){
                 return true
            } else if isPalindrome(s[0:i] + s[i+1:]) {
                return true
            }
        }
        return false
    }
    return isPalindrome(s)
}
