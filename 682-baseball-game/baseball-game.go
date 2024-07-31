func calPoints(operations []string) int {
    r := []int{}

    for _, ops := range operations {
        switch(ops) {
            case "+":
                if len(r) >= 2 {
                    r = append(r, r[len(r) - 1] + r[len(r) - 2])
                }
            case "D":
                if len(r) > 0 {
                    r = append(r, r[len(r) - 1] * 2)
                }
            case "C":
                if len(r) > 0 {
                    r = r[:len(r)-1]
                }
            default:
                if v, err := strconv.Atoi(ops); err == nil {
                    r = append(r, v)
                }
        }
    }

    sum := 0
    for _, v := range r {
        sum += v
    }

    return sum
}
