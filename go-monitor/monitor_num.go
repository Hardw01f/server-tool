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

// init var
var Timestamp int = 0
var counter int = 0


func main(){
		//get commandline args
		flag.Parse()
		args := flag.Args()
		FILENAME := args[0]

		// for debug
		//fmt.Println(FILENAME)
		//fmt.Println("get() : ",get(FILENAME))

		//get Targetfile timestamp
		Timestamp = get(FILENAME)

		//fmt.Println(reflect.TypeOf(Timestamp))

		//Create result file to ./log/
		file, _ := os.Create("./log/res_"+strconv.Itoa(Timestamp)+".txt")
		defer file.Close()

		FirstTime := fmt.Sprint("Starting-Timestamp : ",Timestamp)
		FirstTime = FirstTime + "\n"
		//fmt.Println(FirstTime)

		//Write to resultfile
		fmt.Fprint(file,FirstTime)

		time.Sleep(1 * time.Second)

		for{
				//Jadge Targetfile timestamp is change? 
				if Timestamp != get(FILENAME){
						//done Targetfile timestamp was changed

						//fmt.Println("OverWriting file now !! : ",get(FILENAME))
						WrittenTime := fmt.Sprint("OverWriting file now !! : ",get(FILENAME))
						WrittenTime = WrittenTime + "\n"
						//fmt.Println(WrittenTime)

						if counter == 1{
								//write Timestamp chenging time
								fmt.Fprint(file,WrittenTime)

						}else{
								if counter == 5{
								// wait 5 second, Exit Program
								os.Exit(0)
								}
						}
						counter += 1
				}

		time.Sleep(1 * time.Second)
		}
}


// Function of for getting Targetfile Timestamp(int)
func get(FILENAME string)int{
		res , err := exec.Command("ls","-la",FILENAME).Output()
		if err != nil{
				fmt.Println("error : ",err)
		}
		convert := string(res)
		//fmt.Println(convert)

		array := strings.Split(convert," ")
		//fmt.Println("array : ",array)
		splited := strings.Split(array[11],":")
		//fmt.Println(array[9]+array[10]+array[6]+splited[0]+splited[1])
		resint, _  := strconv.Atoi(array[9]+array[10]+array[6]+splited[0]+splited[1])
		return resint
}


