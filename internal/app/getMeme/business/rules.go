package business

import "fmt"

var urls []string = []string{"a", "b", "c", "d"}

func Execute(input Input) (Output, error) {
  output := output{}
  for _, v := range urls {
    item := item{
      url: v,
    }
    output.items = append(output.items, item)
  }

  return output, nil 
}
