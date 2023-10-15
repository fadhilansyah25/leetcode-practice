func canConstruct(ransomNote string, magazine string) bool {
    letterHashMap := map[rune]int{}

    for _, v := range magazine {
      letterHashMap[v]++
    }

    for _, v := range ransomNote {
      letterHashMap[v]--
      if(letterHashMap[v] < 0) {
        return false
      }
    }

    return true
}