func canConstruct(ransomNote string, magazine string) bool {
    letterHashMap := make([]int, 26)

    for _, v := range magazine {
      letterHashMap[v - 'a']++
    }

    for _, v := range ransomNote {
      letterHashMap[v - 'a']--
      if letterHashMap[v - 'a'] < 0 {
        return false
      }
    }

    return true
}