package main

import ("flag"
        "fmt"
        "os"
        "strings"
        "unicode/utf8")

func main() {
  // Submit a string with -string=[Your text]
  // os.Args[2:] is how to get all arguments after flag
  str := flag.String("string", "", "String to fix.")
  // help := flag.String("help", "--string <string to fix>", "Help text")
  // char := flag.String("characters", "", "Characters to scan for.")
  flag.Parse()

  if *str == "" {
    flag.PrintDefaults()
    os.Exit(1)
  }

  strSlice := strings.Split(*str, "")
  length := utf8.RuneCountInString(*str)
  spaceBar := makeSpace(length)
  indices := findIndices(strSlice)
  fmt.Println(indices)
  const carrot = "^"
  spaceCarrotBar := replaceAtIndex(spaceBar, carrot, indices)
  // This gets wicked repetitive and seems like it could/should
  // be its own function.
  str1 := strings.Replace(*str, "”", "\"", -1)
  // Would like to better understand why ^ this one requires the value,
  // but the others can just use the direct variable name. Flag thing or
  // parameter thing?
  str2 := strings.Replace(str1, "“", "\"", -1)
  str3 := strings.Replace(str2, "‘", "'", -1)
  str4 := strings.Replace(str3, "’", "'", -1)

  fmt.Println("ftfy:", str4)
  fmt.Println(spaceCarrotBar, "< there it is")
}

func makeSpace(l int) string {
  spaces := make([]string, 0)

  for i := 0; i < l + 5; i++ {
    spaces = append(spaces, " ")
  }

  concatSpaces := strings.Join(spaces, "")

  return concatSpaces
}

func findIndices(slice []string) []int {
  indices := make([]int, 0)

  for i, c := range slice {
    if(c == "”" || c == "“" || c == "‘" || c == "’") {
      indices = append(indices, i)
    }
  }
  return indices
}

func replaceAtIndex(in string, char string, i []int) string {
  strSlice := []string(in)

  for _, c := range i {
    strSlice[c] =     (char)
    // Want to make this reusable and take the character as a parameter,
    // but am stuck on this part where the character also has to
    // become a rune for replacement's sake.
  }

  out := fmt.Sprintf("%c", strSlice)

  return string(out)
}
