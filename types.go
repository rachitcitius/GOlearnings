// types
package main

import (
	"fmt"
)

type Vehicle struct {
	engine string
	wheels int
	chasis string
	cylinders int
}

type Motorcycle struct {
	Vehicle
	Handlebar string
}

type Car struct {
	Vehicle
	steering string
	window string
}

func main() {
	// sameLenArr()
	distribute50bitCoins()
}

func sameLenArr() {
	//Group all the names with same length
	var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
						"Emma", "Isabella", "Emily", "Madison",
						"Ava", "Olivia", "Sophia", "Abigail",
						"Elizabeth", "Chloe", "Samantha",
						"Addison", "Natalie", "Mia", "Alexis", "Foo", "R"}
						
	//Getting a string with max length
	var maxLen int
	for _, val:=range names {
		//fmt.Println(x,val)
		if l:=len(val); l > maxLen {
			maxLen=l
		}		
	}
	
	output := make([][]string, maxLen)
	for _, val:=range names {
		output[len(val)-1] = append(output[len(val)-1], val)
	}
	
	fmt.Printf("%v\n", output)
}

func distribute50bitCoins() {
/*50 bitcoins to distribute to 10 users
The coins will be distributed based on the vowels contained in each name where:
a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins
*/

var (
	//coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
	distribution = make(map[string]int, len(users))
	)
	
	NoOfBitCoins := func(name string) int {
		var total int
		for y:=0; y < len(name); y++ {
			switch string(name[y]) {
				case "a","A","e","E":
					total++
				case "i","I":
					total=total+2
				case "o","O":
					total=total+3
				case "u","U":
					total=total+4
			}
		}
		return total
	}
	
	for _, user:=range users {
		distribution[user]=NoOfBitCoins(user)
	}
	fmt.Println(distribution)
}

func commaOk () {
	var m = map[string]int {"A":20, "B":50, "C":60}
	fmt.Println(m)
	
	//comma ok usage
	elm, ok:=m["B"]
	if ok { fmt.Println("Present")   
	} else { 
	fmt.Println(elm)  
	}
}

func type1 () {
	//switch case
	var value interface{}
	switch str:= value.(type) {
	case string: {
		fmt.Println(str) }
	default: { fmt.Println(str) }
	}
	
	//comma ok idiom
	str1, ok:= value.(string)
	if ok {
		fmt.Println(str1)
	}
	
	//initializing with new keyword
	x:=new(int)
	fmt.Println(*x)
	
	a := [...]string{"hello", "world","nice","cool"}
	fmt.Printf("%s\n",a)
}
