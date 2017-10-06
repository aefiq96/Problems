package main

import "fmt"

func main() {
    //variable
    var input string
    //ask user 
    fmt.Println("Enter the string for reverse")
    fmt.Scanf("%s",&input)
    //print the result
    reversed1 := reverse(input)
    fmt.Println(reversed1)
}
//function to reverse
func reverse(value string) string {
   
    data := []rune(value)
    result := []rune{}
    
    for i := len(data) - 1; i >= 0; i-- {
        result = append(result, data[i])
    }

    return string(result)
}

