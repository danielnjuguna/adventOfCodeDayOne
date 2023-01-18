package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// Declaring the Calories type

type Calories []int
var elf = map[int]int{}
func main() {
  //Loading Input text using the os bultin package
	f, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }

  //Declaring and initializing an empty map and creating an instance of newReader
    var reader = bufio.NewReader(f)   
    max, err := partOne(reader)
    if err != nil{
      log.Fatal(err)
    }
    fmt.Println(max)// My answer for Part 1 = 75501

    sum, err := partTwo()
    if err != nil{
      log.Fatal(err)
    }
    fmt.Println(sum) // My Answer for Part 2 = 215594

  

}

func partOne(b *bufio.Reader) (int, error){
  i := 0
  for {

    //Reading each line and returning the slice bytes read.
    line, _, err := b.ReadLine()
    if err != nil {
      //EOF for End of File
      if err == io.EOF {
        break
      }
      log.Fatal(err)
    }
    if len(line) == 0 {
      i++
      continue
    }

    //Converting the int type
    calorie, _ := strconv.Atoi(string(line))

    _, ok := elf[i]
    if !ok {
      elf[i] = calorie
      continue
    }
    elf[i] += calorie    
  }
  
  max := 0
  for _, cals := range elf {
    if cals > max {
      max = cals
    }
  }
  return max, nil
}

func partTwo() (int, error){
  calList := make(Calories, len(elf))
  for i, c := range elf {
    //Populating the slice int of type Calories with the value of elf map
    calList[i] = c
  }

  //Sorting the populated slice of type Calories in descending order
  sort.Slice(calList, func(i, j int) bool {
		return calList[i] > calList[j]
	})

  sum := 0
  top3 := 3
  for i := 0; i < top3; i++ {
    sum += calList[i]
  }
  return sum, nil
}
