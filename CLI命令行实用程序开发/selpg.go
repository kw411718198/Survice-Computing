package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag"github.com/spf13/pflag"
)

type sp_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int
	page_type   bool
	print_dest  string
}

func main() {
	var sa sp_args
	init_args(&sa)
	process_args(&sa)
	process_input(&sa)
}

func init_args(sa *sp_args) {
	flag.IntVarP(&(sa.start_page), "start_page", "s", -1, "Define startpage")
	flag.IntVarP(&(sa.end_page), "end_page", "e", -1, "Define endpage")
	flag.IntVarP(&(sa.page_len), "page_len", "l", 72, "Define page length")
	flag.BoolVarP(&(sa.page_type), "page_type", "f", false, "Define page type")
	flag.StringVarP(&(sa.print_dest), "print_dest", "d", "", "Define print_dest")

	flag.Usage = func() {
		fmt.Printf("Usage:\n\n")
		fmt.Printf("\tselpg -s start_page -e end_page [-f (speciy how the page is sperated)| -l lines_per_page_default_72] [-d dest] [filename]\n\n")
	}
	flag.Parse()
}

//参数检查
func process_args(sa *sp_args) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "\nnot enough arguments\n")
		flag.Usage()
		os.Exit(1)
	}

	if (sa.start_page == -1) || (sa.end_page == -1) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be empty! Please check your command!\n")
		flag.Usage()
		os.Exit(2)
	}

	if (sa.start_page <= 0) || (sa.end_page <= 0) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be negative! Please check your command!\n")
		flag.Usage()
		os.Exit(3)
	}

	if sa.start_page > sa.end_page {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage can't be bigger than the endPage! Please check your command!\n")
		flag.Usage()
		os.Exit(4)
	}
	if len(flag.Args()) == 1 {
		_, err := os.Stat(flag.Args()[0])
		if err != nil && os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "\ninput file \"%s\" does not exist\n", flag.Args()[0])
			os.Exit(5)
		}
		sa.in_filename = flag.Args()[0]
	}

	if (sa.page_type == true) && (sa.page_len != 72) {
		fmt.Fprintf(os.Stderr, "\n[Error]The command -l and -f are exclusive, you can't use them together!\n")
		flag.Usage()
		os.Exit(6)
	}

	if sa.page_len <= 0 {
		fmt.Fprintf(os.Stderr, "\n[Error]The pageLen can't be less than 1 ! Please check your command!\n")
		flag.Usage()
		os.Exit(7)
	} else {
		fmt.Printf("\n[ArgsStart]\n")
		fmt.Printf("startPage: %d\nendPage: %d\ninputFile: %s\npageLength: %d\npageType: %s\nprintDestation: %s\n[ArgsEnd]", sa.start_page, sa.end_page, sa.in_filename, sa.page_len, sa.page_type, sa.print_dest)
	}

}

//数据输出
func process_input(sa *sp_args) {
	var fin *os.File
	if sa.in_filename == "" {
		fin = os.Stdin
	} else {
		var err error
		fin, err = os.Open(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n[Error]%s:", sa.in_filename)
			os.Exit(8)
		}
	}
	line_count := 0
	page_count := 1
	buf := bufio.NewReader(fin)

	cmd := &exec.Cmd{}
	var fout io.WriteCloser
	if sa.print_dest == "" {
		fout = os.Stdout
	} else {
		cmd = exec.Command("cat")
		var err error
		cmd.Stdout, err = os.OpenFile(sa.print_dest, os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n[Error]%s:", "Input pipe open\n")
			os.Exit(9)
		}
		fout, _ = cmd.StdinPipe()
		cmd.Start()
	}

	for true {
		var line string
		var err error
		if sa.page_type {
			line, err = buf.ReadString('\f')
			page_count++
		} else {
			line, err = buf.ReadString('\n')
			line_count++
			if line_count > sa.page_len {
				page_count++
				line_count = 1
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n[Error]%s:", "Input pipe open\n", "file read in\n")
			os.Exit(10)
		}
		if (page_count >= sa.start_page) && (page_count <= sa.end_page) {
			var outputErr error
			_, outputErr = fout.Write([]byte(line))
			if outputErr != nil {
				fmt.Fprintf(os.Stderr, "\n[Error]%s:", "pipe input")
				os.Exit(11)
			}

			if outputErr != nil {
				fmt.Fprintf(os.Stderr, "\n[Error]%s:", "Error happend when output the pages.")
				os.Exit(12)
			}
		}
	}
}
