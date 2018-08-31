package main

import(
		"fmt"
		"os/exec"
		"strings"
		"flag"
		"strconv"
		"time"
		"os"
)

var Timestamp int = 0
var counter int = 0


func main(){
		flag.Parse()
		args := flag.Args()
		FILENAME := args[0]

		//fmt.Println(FILENAME)
		//fmt.Println("get() : ",get(FILENAME))
		Timestamp = get(FILENAME)

		//fmt.Println(reflect.TypeOf(Timestamp))

		file, _ := os.Create("./log/res_"+strconv.Itoa(Timestamp)+".txt")
		defer file.Close()

		FirstTime := fmt.Sprint("Starting-Timestamp : ",Timestamp)
		FirstTime = FirstTime + "\n"
		//fmt.Println(FirstTime)

		fmt.Fprint(file,FirstTime)

		time.Sleep(1 * time.Second)

		for{
				if Timestamp != get(FILENAME){
						//fmt.Println("OverWriting file now !! : ",get(FILENAME))
						WrittenTime := fmt.Sprint("OverWriting file now !! : ",get(FILENAME))
						WrittenTime = WrittenTime + "\n"
						//fmt.Println(WrittenTime)
						if counter == 1{
								fmt.Fprint(file,WrittenTime)
						}else{
								if counter == 5{
								os.Exit(0)
								}
						}
						counter += 1
				}

		time.Sleep(1 * time.Second)
		}
}


func get(FILENAME string)int{
		res , err := exec.Command("ls","-la",FILENAME).Output()
		if err != nil{
				fmt.Println("error : ",err)
		}
		convert := string(res)
		//fmt.Println(convert)

		array := strings.Split(convert," ")
		//fmt.Println(array)
		splited := strings.Split(array[7],":")
		//fmt.Println(array[6]+splited[0]+splited[1])
		resint, _  := strconv.Atoi(array[6]+splited[0]+splited[1])
		return resint
}



