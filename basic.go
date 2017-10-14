package main

import (
  "fmt"
)

/*
  loops,functions,arrya,slice,maps,pointers
  structs,,interfaces,erros,concurrency,
*/

type Node struct {
  data int
  name string
}

// method declaration on struct's
func (t Node) combine() string {
    return "test"
}

func main() {

  // // decleration and loops
  // count := 10
  //  for i := 0;i < count;i++ {
  //    //fmt.Println(i);
  //  }
  //
  //  a,b := 1,2 // shorthand
  //  var t11,t22 int = 0,1
  //  var t int // decleration
  //  t = 0 // initialization
  //  var t1 int = 0 // decleration,initialization
  //  fmt.Printf("t:%d,t1:%d,t11:%d,t22:%d",)
  //
  // // arrays[length],slices[no-lenght],ranges
  // var a1213 [10]int
  // var a12 []int
  //
  // var array1 = [2]int{1,2} // size,type,values, var,type,value
  // array2 := [3]int{1,2,3};

  // array3 := [...]int{1,2,3}
  // fmt.Println(array3)
  // count := 0

  // for range time.Tick(time.Second) {
    // count++
    // fmt.Println(count)
  // }

  // for _,e := range array3 {
  //   for range time.Tick(time.Millisecond) {
  //   }
  // }

  var p *int
  r := 10
  p = &r
  fmt.Println(*p)

    // creating a struct
    var store = Node{10,"hi"}
    var store2 = Node{data: 11,name: "chan"}
    fmt.Println(store,store2)

    fmt.Println(store2.data)

    var store3 = []Node{{1,"new"},{2,"another"}}
    fmt.Println(store3[1])

    fmt.Println(store.combine)
}
