func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    
    for _, cust := range accounts {
        currentWealth := 0
        for _, bank := range cust {
            currentWealth += bank
        }
        
        maxWealth = int(math.Max(float64(maxWealth), float64(currentWealth)))
    }
    
    return maxWealth
}