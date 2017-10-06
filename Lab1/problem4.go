//Nuzhafiq - G00324934 - Program to find Factorial of number
//http://www.golangprograms.com/go-program-to-find-factorial-of-a-number.html
package main
import "fmt"
 
/* Variable Declaration */
var factVal float64 = 1 // uint64 is the set of all unsigned 64-bit integers.
                       // Range: 0 through 18446744073709551615. 
var i int = 1
var n int
 
/*     function declaration     */
func factorial(n int) float64 {   
    if(n < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{        
        for i:=1; i<=n; i++ {
            factVal *= float64(i)  // mismatched types int64 and int
        }
         
    }    
    return factVal  /* return from function*/
}
 
func main(){  

    fmt.Print("Enter a positive integer between 0 - 100 : ")
    fmt.Scan(&n)   
    fmt.Println("Factorial is: ",factorial(n))
	fmt.Print("The sum is: ",sum(n))
}

func sum(n int )int{
	add := 0
	storeInt := [250] int{};
	storeInt[0] = 1;

	for i := 2; i <= n; i++{
		for j := 0; j < len(storeInt); j++{
			storeInt[j] *= i;
			if j > 0 && storeInt[j-1] > 9{
				storeInt[j] += int(storeInt[j-1]/10)
				storeInt[j-1] %= 10;
			}
		} 
	}
	for i := 0; i < len(storeInt); i++{
		add += storeInt[i];
	}
	return add;
}