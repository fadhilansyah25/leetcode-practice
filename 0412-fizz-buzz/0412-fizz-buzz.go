func fizzBuzz(n int) []string {
    res := make([]string, n)

    for i := 0; i < n; i++ {
        if (i + 1) % 3 == 0 {
            res[i] += "Fizz"
        }

        if (i + 1) % 5 == 0 {
            res[i] += "Buzz"
        }

        if res[i] == "" {
            res[i] = strconv.Itoa(i+1)
        }   
        
    }

    return res
}