import "strconv"

func fizzBuzz(n int) []string {
    var result []string
    if n<= 0 {
        return result
    }
    
    result = make([]string, n)

    for i:= 1; i<=n; i++ {
        result[i-1] = strconv.Itoa(i)  
    } 
    
    for i:= 3; i<=n;{
        result[i-1] = "Fizz"
        i += 3
    }
    
    for i:= 5; i<=n;{
        result[i-1] = "Buzz"
        i += 5
    }
    
    for i:= 15; i<=n;{
        result[i-1] = "FizzBuzz"
        i += 15
    }
    return result
}
