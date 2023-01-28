package main


import(
	"fmt"
	"bufio"

	"os"


)
type name struct{

	fname string

	lname string

}


func main(){

	fmt.Print("Enter the file name (full)=")

	var n string

	fmt.Scanln(&n)


	f, _ := os.Open(n)

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanWords)

	var names []name


	for fileScanner.Scan(){

		var nm name

		nm.fname = fileScanner.Text()

		fileScanner.Scan()

		nm.lname = fileScanner.Text()

		names=append(names,nm)

	}


	for _, v:=range names{

		fmt.Printf("First name: %s \tLast name: %s\n",v.fname,v.lname)

	}
}.