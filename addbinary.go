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
    
    result := "" 
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
        
        result = fmt.Sprintf("%d", op % 2 + '0') + result
  
        // Compute carry 
        op = op/ 2
        
        aIdx = aIdx -1
        bIdx = bIdx -1
    } 
    return result
}
