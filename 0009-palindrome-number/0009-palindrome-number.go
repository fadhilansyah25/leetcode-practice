func isPalindrome(x int) bool {
    rev := 0
    temp := x

    for temp > 0 {
        rev = rev * 10 + temp % 10
        temp /= 10
    }

    return x == rev
}