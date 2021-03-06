package main
import (
	"os"
	. "flag"
	"fmt"
	"io/ioutil"
	"strconv"
	. "strings"
)

const (
	OK = 0
	NO_SUCH_FILE
)

var name, spacer *string
var message, text_block	string
var repeats int

func init() {
	var def_c int
	if def_r := os.Getenv("DEF_REPS"); len(def_r) != 0 {
		var err error
		if def_c, err = strconv.Atoi(def_r); err != nil {
			fmt.Println("DEF_REPS must be an integer:", def_r)
		}
	}

	name = String("n", defaultName(), "n: name of person to greet")
	spacer = String("s", ",", "s: separator between name and message")
	file := String("f", "", "f: name of a file containing a block of text to display")
	IntVar(&repeats, "c", def_c, "c: number of times to display the message")
	Parse()
	if len(*file) > 0 {
		if text, err := ioutil.ReadFile(*file); err == nil {
			text_block = string(text)
		} else {
			fmt.Println("no such file:", *file)
		}
	}
	message = Join(Args(), " ")
}

func main() {
	switch {
	case len(message) > 0:
		message = fmt.Sprintf("hello %v%v %v", *name, *spacer, message)
		if len(text_block) > 0 {
			message += "\n" + text_block
		}
	case len(text_block) > 0:
		message = fmt.Sprintf("hello %v\n%v", *name, text_block)
	default:
		message = fmt.Sprintf("hello %v", *name)
	}

	for ; repeats > 0; repeats-- {
		fmt.Println(message)
	}
}

func defaultName() (r string) {
	if r = os.Getenv("DEF_NAME"); len(r) == 0 {
		r = "world"
	}
	return
}