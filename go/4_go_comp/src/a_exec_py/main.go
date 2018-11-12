package main

import "fmt"
import "os/exec"

func main() {
	cmd := exec.Command("python",  "-c", "import pythonfile; print pythonfile.cat_strings('foo', 'bar')")
	fmt.Println("-----11111")
    fmt.Println(cmd.Args)
    out, err := cmd.CombinedOutput()
	if err != nil { fmt.Println(err); }
	fmt.Println("222beg -----------")
	fmt.Println(string(out))
	fmt.Println("222end--------------")
	
	fmt.Println("-----------------------")
	cmd = exec.Command("python",  "a.py")
	out1, _ := cmd.CombinedOutput()
	fmt.Println(out1)
	fmt.Println(string(out1))

}