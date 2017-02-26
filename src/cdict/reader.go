package main

import (
	"debug/dwarf"
	"debug/elf"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	//"strings"
)

func main() {

	path := getCurrentPath()
	fmt.Println("path =", path)

	filename := "..\\testdata\\test1_c.o"

	file, err := elf.Open(filename)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s, err = %s", filename, err.Error())
		return
	}
	defer file.Close()

	data, err := file.DWARF()
	if err != nil {
		fmt.Println("ERROR: cannot get dwarf")
		return
	}

	reader := data.Reader()
	for {
		e, err := reader.Next()
		if err != nil {
			fmt.Println("ERROR: failed to read dwarf")
			return
		}
		fmt.Println("e =", e)
		//fmt.Println("e.Offset =", e.Offset)
		//fmt.Println("e.Tag =", e.Tag)
		if e == nil {
			break
		}
		if e.Tag == dwarf.TagTypedef {
			typ, err := data.Type(e.Offset)
			if err != nil {
				fmt.Println("ERROR: failed to read data.Type()")
				return
			}

			t1 := typ.(*dwarf.TypedefType)
			var typstr string
			if ts, ok := t1.Type.(*dwarf.StructType); ok {
				typstr = ts.Defn()
			} else {
				typstr = t1.Type.String()
			}
			fmt.Println("type =", typstr)
		}
	}

}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)

	path, _ := filepath.Abs(s)
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
