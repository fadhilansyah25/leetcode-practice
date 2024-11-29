func lengthOfLastWord(s string) int {
    result := strings.Fields(s);

    return len(result[len(result) -1 ]);
}