//training sri

package main
import ("fmt"
"os"
)

func finished() {
fmt.Println ("Finished work")
}

func somefunction() int {
defer finished()
return 101

}

func main() {

f, err := os.Open("abc.txt")
defer f.Close()

if err != nil {
fmt.Println(err)
return
}

// call the function which calls a defer
somefunction()

}
