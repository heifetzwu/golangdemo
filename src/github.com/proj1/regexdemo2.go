package main

import "fmt"

import "bytes"
import "regexp"

func regexdemo2() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	fmt.Println(match)
	
	r, _ := regexp.Compile("p([a-z]+)ch")

	
	fmt.Println(1, r.MatchString("peach"))

	fmt.Println(2, r.FindString("peach punch"))

	fmt.Println(3, r.FindStringIndex("peach punch"))

	fmt.Println(4, r.FindStringSubmatch("peach punch"))
	
	fmt.Println(5, r.FindStringSubmatchIndex("peach punch"))
	
	fmt.Println(6, r.FindAllString("peach punch pinch", -1))  

    fmt.Println(7.1, r.FindAllStringSubmatch("peach punch pinch", -1))
    fmt.Println(7.2, r.FindStringSubmatch("peach punch pinch"))
    
    fmt.Println(8, r.FindAllStringSubmatchIndex("peach punch pinch", -1))  
 
    fmt.Println(9, r.FindAllString("peach punch pinch", 2)) 
    fmt.Println(10, r.FindAllString("peach punch pinch", 3)) 
    
    fmt.Println(11, r.Match([]byte("peach")))  
    
    r = regexp.MustCompile("p([a-z]+)ch")  
    fmt.Println(12, r)  
    
    fmt.Println(13, r.ReplaceAllString("a peach", "<fruit>")) 
    
    in := []byte("a peach")  
    out := r.ReplaceAllFunc(in, bytes.ToUpper)  
    fmt.Println(14, string(out))  
// 	fmt.Println("match")
}
