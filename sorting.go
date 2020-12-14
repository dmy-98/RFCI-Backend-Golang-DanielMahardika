package main

import "fmt"

func Sort(array []int) {
    total:=0
    for i := 0;i < len(array)-1;i++{
        if array[i]>array[i+1]{
            total++
            array[i],array[i+1] = array[i+1],array[i]
            fmt.Printf("%d. [%d,%d] -> %d\n",total,array[i],array[i+1],array);
            i=-1
        }
        fmt.Printf("i = %d\n",i)
    }
    fmt.Printf("Jumlah swap: %d\n",total)
}

func main() {
    input := []int{4, 9, 7, 5, 8, 9, 3}
    fmt.Printf("soal : %d\n",input)
    Sort(input)
}