//Palindrome test //http://www.golangpro.com/2016/01/check-string-palindrome-go.html
package main

import (   
    "fmt"
    "strings"
)

func main() {
    //variables
    var palindrome string

    fmt.Print("Enter string:")
    fmt.Scanf("%s\n", &palindrome)
    //this converts string to lowercase
    palindrome = strings.ToLower(palindrome)
    //prints out the result
    fmt.Println(ispalindrome(palindrome))
}

func ispalindrome(s string) string {//checks if the string ispalindrome
//the func is type boolean

    mid := len(s) / 2
    last := len(s) - 1
        for i := 0; i < mid; i++ {
             if s[i] != s[last-i] {
                 //returns no if not
                return "NO. It's not a Palimdrome."
             }
        }
    //returns yes if it is
    return "YES! You've entered a Palindrome"
}