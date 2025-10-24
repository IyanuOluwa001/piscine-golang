func LoafOfBread(str string) string {
    // Check if string is less than 5 characters
    if len(str) < 5 {
        return "Invalid Output\n"
    }

    // Remove spaces from string for processing
    cleaned := ""
    for _, char := range str {
        if char != ' ' {
            cleaned += string(char)
        }
    }

    // If cleaned string is less than 5 characters
    if len(cleaned) < 5 {
        return "Invalid Output\n"
    }

    result := ""
    for i := 0; i < len(cleaned); i += 6 {
        // If remaining characters are less than 5, break
        if i+5 > len(cleaned) {
            break
        }
        // Add 5 characters to result
        result += cleaned[i:i+5]
        // Add space after 5 characters
        if i+5 < len(cleaned) {
            result += " "
        }
    }
    // Add newline at the end
    result += "\n"
    return result
}