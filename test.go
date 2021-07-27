package main

import (
    "crypto/rand"
    "fmt"
    "math/big"
)

type FeedItem struct {
     Author int// 作者
     Category int // 类别
}
func createRand() int {
    result, _ := rand.Int(rand.Reader, big.NewInt(5))
    return int(result.Int64())
}

func reorderFeedItems(inputItems []FeedItem)([]FeedItem, error){
    inputItems1 := inputItems[:10]
    // m := len(c)-1
    fmt.Println(inputItems1) 
    for i,v := range inputItems1{
       //判断是否作者重复
       if v.Author==inputItems[i+1].Author{
            getdiffAuthor(v.Author,i+1,inputItems)
            fmt.Println(inputItems1) 
       }
       //判断是否类别重复

       if i>0 && i<9  { 
            // fmt.Println(i) 
            if v.Category==inputItems[i+1].Category && v.Category==inputItems[i-1].Category{
                getdiffCategory(v.Category,i+1,inputItems)
            }
       }
       
    }
    return inputItems1,nil
}
func getdiffAuthor(Author int ,index int ,inputItems []FeedItem){
    inputItems1 := inputItems[index:]
    // fmt.Println(inputItems)
    for i,v := range inputItems1{
        if v.Author!=Author {
            changeData(index,i+index,inputItems)
            // return
            break 
        }
    }
}
func getdiffCategory(Category int ,index int ,inputItems []FeedItem){
    inputItems1 := inputItems[index:]
    // fmt.Println(inputItems)
    for i,v := range inputItems1{
        if v.Category!=Category {
            changeData(index,i+index,inputItems)
            break
        }
    }
}
func changeData(index1 int ,index2 int ,inputItems []FeedItem){
    // fmt.Println(index1)
    // fmt.Println(index2)
    inputItems[index2],inputItems[index1] = inputItems[index1],inputItems[index2]
}


func main() {
    // var inputItems []FeedItem
    inputItems := make([]FeedItem, 50)
    
    for i := 0; i < 50; i++ {
        
        inputItems[i].Author = createRand()
        inputItems[i].Category = createRand()
    }
    
    FeedItem,err := reorderFeedItems(inputItems)
    // fmt.Println(inputItems)
    fmt.Println(FeedItem)
    fmt.Println(err)
}