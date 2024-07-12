package main

import "fmt"

// get command args
func readCmd() (args []string, err error) {
	fmt.Print("> ")

	inputs := make([]string, 5)
	potInputs := make([]interface{}, 5)
	for i := range inputs {
		potInputs[i] = &inputs[i]
	}
	_, err = fmt.Scanln(potInputs...)

	if err != nil {
		return nil, err
	}
	// type enter, do nothing
	if inputs[0] == "" {
		return nil, nil
	}
	return inputs, nil
}

// excute command
func excuteCmd(exist *bool) {
	args, err := readCmd()

	if err != nil {
		fmt.Println(err)
		return
	}

	if args != nil && args[0] == ".exit" {
		*exist = true
		return
	}

	fmt.Printf("get args: %v", args)
}

// dir operation:
//
//	ls: -a -d -l
//	cd: relative or regular path
//	pwd: -P show symbol
//	mkdir: -m mode -p
//	rmdir: -p remove multiple empty dir
//	cp:
//	rm: -f ignore nonexists file -r recursion rm file or dri -i ask
//	mv:
//
// file operation:
//
//	cat
//	tac
//	vim
func main() {
	fmt.Println("Welcome")
	exist := false
	for !exist {
		excuteCmd(&exist)
	}
}
