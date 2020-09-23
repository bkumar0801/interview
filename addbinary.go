/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

func addBinary(a string, b string) string {
    sz := 0
    if len(a) < len(b) {
        sz = len(b)
    } else {
        sz = len(a)
    }
    result := make ([]rune, sz + 1)
    idx := sz

    op := 0 
  
    aIdx := len(a) - 1
    bIdx := len(b) - 1
    
    for (aIdx >= 0 || bIdx >= 0 || op == 1) { 
        if aIdx >= 0 {
            op = op + int(a[aIdx])- '0'
        } 
        
        if bIdx >= 0 {
            op = op + int(b[bIdx]) - '0'
        } 
            
        op = op + 0
        
        result[idx] = rune(op % 2 + '0')
        
        idx = idx - 1  
        // Compute carry 
        op = op/ 2
        
        aIdx = aIdx -1
        bIdx = bIdx -1
    } 
    return string(result)
    
}
