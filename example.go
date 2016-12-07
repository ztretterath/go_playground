package main

import "fmt"

func main(){

  // fmt.Println("hello world") //like consoloe.log or puts

  // var age int = 40
  //
  // var num float64 = 1.6184
  //
  // // randNum := 1
  //
  // fmt.Println(age, num)
  //
  // var numOne = 1.000
  // var num99 = .999
  //
  // fmt.Println(numOne - num99) //floats not always accurate

  // const pi float64 = 3.14159265
  // //
  // // var myName string = "Zach R"
  // //
  // // fmt.Println(len(myName))
  //
  // // var isOver20 bool = true
  //
  // fmt.Printf("%.3f \n", pi) //specify float length
  //
  // fmt.Printf("%T \n", pi) //type

  // fmt.Println("true && false = ", true && false)
  // fmt.Println("true || false = ", true || false)

  /////FOR LOOPS/////
  // i := 1
  // for i <= 10 {
  //   fmt.Println(i)
  //   i++
  // }
  //
  // for j := 0; j < 5; j++ {
  //   fmt.Println(j)
  // }

  /////IF ELSE/////
  // yourAge := 18
  // if yourAge >=16 {
  //   fmt.Println("you can drive")
  // } else if yourAge >= 18{
  //   fmt.Println("you can vote")
  // }


  /////SWITCH/////
  // yourAge := 5
  // switch yourAge{
  // case 16 : fmt.Println("Go drive!")
  // case 18 : fmt.Println("Go vote!")
  // default : fmt.Println("Go have fun!")
  // }

  /////ARRAYS/////
  // favNums2 := [5]float64 {1, 2, 3, 4, 5}
  //
  // for i, value := range favNums2 {
  //   fmt.Println(value, i)
  // }

  /////SLICE/////
  // numSlice := []int {5, 4, 3, 2, 1}

  // numSlice2 := numSlice[3:5]
  //
  // fmt.Println("numSlice2[0] = ", numSlice2[0])

  // fmt.Println("numSlice[:2] = ", numSlice[:2])

  // numSlice3 := make([]int, 5, 10)
  //
  // copy(numSlice3, numSlice)
  //
  // fmt.Println(numSlice3[0])
  //
  // numSlice3 = append(numSlice3, 0, -1)
  //
  // fmt.Println(numSlice3)

  /////MAPS/////
  // presAge := make(map[string] int)
  //
  // presAge["TheodoreRoosevelt"] = 42
  //
  // fmt.Println(len(presAge))

  // num1, num2 := next2values(5)
  // fmt.Println(num1, num2)
  //
  // listNums := []float64{1, 2, 3, 4, 5}
  //
  // fmt.Println(addThemUp(listNums))

  // fmt.Println(subtractThem(1,2,3,4,5))

  /////CLOSURE/////
  // num3 := 3
  // doubleNum := func() int{
  //   num3 *=2
  //   return num3
  // }
  // fmt.Println(doubleNum())
  // fmt.Println(doubleNum())

  // fmt.Println(factorial(3))

  /////DEFER/////
  defer printTwo()
  printOne() //defer executes AFTER main finishes
  //use last func as cleanup func

}

func printOne(){ fmt.Println(1)}
func printTwo(){ fmt.Println(2)}


// factorial(3)
// 3 * factorial(2) == 3 * 2
// 2 * factorial(1) == 2 * 1
// factorial(0) == 1
// func factorial(num int) int {
//   if num == 0 {
//     return 1
//   }
//   return num * factorial(num - 1)
// }

// func subtractThem(args ...int) int {
//   finalValue := 0
//   for _, value := range args{
//     finalValue -= value
//   }
//   return finalValue
// } //(args ...int) b/c unsure num of inputs


// func addThemUp(numbers []float64) float64 {
//   sum := 0.0
//
//   for _, val := range numbers {
//     sum += val
//   }
//   return sum
// }
//
// func next2values(number int) (int, int){
//   return number+1, number+2
// } //(int, int) b/c returning 2 vals
