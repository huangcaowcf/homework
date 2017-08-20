package main

import(
    "fmt"
    "sort"
)

func RemoveDuplicatesAndEmpty(a []string) (ret []string){
    sort.Strings(a)
    a_len := len(a)
    for i:=0; i < a_len; i++{
        if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
            continue;
        }
        ret = append(ret, a[i])
    }
    return ret
}

func main(){
    a := []string{"hello", "yes", "world", "yes", "hello", "nihao", "shijie", "hello", "yes", "nihao","good"}

    fmt.Println(a)
    fmt.Println(RemoveDuplicatesAndEmpty(a))
}
