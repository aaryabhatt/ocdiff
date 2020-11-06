package main

import (
	"fmt"
    "os"
    "os/exec"
)

func main() {
	fmt.Println("This program ocdiff will help to get difference between current and last replicas of a DC/STS")
     
	//Find oc executable path
	ocExecPath, err := exec.LookPath("oc")
		if err != nil {
			fmt.Println("Error is:", err )
		}else {
	        fmt.Println("OC excecutable path",ocExecPath)
		      }

# it will add all strint to profName , including script name and argument
#	progName := os.Args

# To capture all argument passed
#	allArguments := os.Args[1:]

## Get first argument
	dcName := os.Args[1]

#	fmt.Println("Name of program", progName)
#	fmt.Println("All Arguments", allArguments)
	fmt.Println("Name of DC", dcName)

}
