// interface2

/* Created an interface of Mobile. Created two structures for different mobile mfg. with different signatures

*/
package main

import (
	"fmt"
)

type IMobile interface{
	WhatPhone() string
}

type Samsung struct{
	OStype string
	version string
}

type Nokia struct{
	modelName string
	price int
	OStype string
}

func (s Samsung) WhatPhone() string {
	return fmt.Sprintf("This is a Samsung phone running on %s & version %s", s.OStype, s.version)
}

func (n Nokia) WhatPhone() string {
	return fmt.Sprintf("This is a Nokia phone. Model is %s, running on %s & is available for %d",n.modelName,n.OStype,n.price)
}

func getPhoneDetails(iMob IMobile) () {
	fmt.Println(iMob.WhatPhone())
}

func main() {
	sams := Samsung{"Android", "6.2.0"}
	nok := Nokia{"6500",25000,"Symbian"}
	
	getPhoneDetails(sams)
	getPhoneDetails(nok)
}
