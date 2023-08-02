package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/google/shlex"
)

func getIndexSeparator(str string) [][]int {
	re := regexp.MustCompile(`[^\\](?:\\\\)*(\|)`)
	return re.FindAllStringIndex(str, -1)
}

func splitBySeparator(str string, matched [][]int) (arr []string) {
	left := 0
	for i := 0; i < len(matched); i++ {
		arr = append(arr, str[left:matched[i][1]-1])
		left = matched[i][1]
	}
	arr = append(arr, str[left:])
	return arr
}

func splitString(str string) []string {
	matched := getIndexSeparator(str)
	return splitBySeparator(str, matched)
}

func getCurrPath(stdout io.Writer) {
	dir, ok := os.Getwd()
	if ok != nil {
		fmt.Fprintln(stdout, ok)
	} else {
		fmt.Fprintln(stdout, dir)
	}
}

func changeDir(stdout io.Writer, path string) {
	fmt.Println(path)
	err := os.Chdir(path)
	if err != nil {
		fmt.Println(err)
	}
	getCurrPath(stdout)
}

func outputToConsole(stdout io.Writer, row []string) {
	str := strings.Join(row, " ")
	if len(str) != 0 {
		io.WriteString(stdout, str)
		io.WriteString(stdout, "\n")
	}
}

type ICommands interface {
	Start() error
	Wait() error
}

type EmbeddedCommandCb func(io.Reader, io.Writer) error

type embeddedCommand struct {
	stdin   io.Reader
	stdout  io.Writer
	command EmbeddedCommandCb
	wg      sync.WaitGroup
}

func newEmbeddedCommand(stdin io.Reader, stdout io.Writer, command EmbeddedCommandCb) *embeddedCommand {
	return &embeddedCommand{
		stdin:   stdin,
		stdout:  stdout,
		command: command,
	}
}

func (ecomm *embeddedCommand) Start() error {
	// check if already started
	ecomm.wg.Add(1)
	go func() {
		if err := ecomm.command(ecomm.stdin, ecomm.stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		ecomm.wg.Done()
	}()
	return nil
}

func (ecomm *embeddedCommand) Wait() error {
	// check if not started
	ecomm.wg.Wait()
	return nil
}

func getСommandFromLine(text string, stdout io.WriteCloser, stdin io.ReadCloser) (error, ICommands) {
	command, err := shlex.Split(text)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	switch command[0] {
	case "cd":
		// Обработать ошибку panic: runtime error:
		// index out of range [1] with length 1

		return nil, newEmbeddedCommand(stdin, stdout, func(_ io.Reader, stdout io.Writer) error {
			changeDir(stdout, command[1])
			return nil
		})
	case "pwd":
		return nil, newEmbeddedCommand(stdin, stdout, func(_ io.Reader, stdout io.Writer) error {
			getCurrPath(stdout)
			return nil
		})
	case "echo":
		return nil, newEmbeddedCommand(stdin, stdout, func(_ io.Reader, stdout io.Writer) error {
			outputToConsole(stdout, command[1:])
			return nil
		})
	case "kill":
		return nil, newEmbeddedCommand(stdin, stdout, func(_ io.Reader, _ io.Writer) error {
			pid, err := strconv.ParseInt(command[1], 10, 32)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}
			syscall.Kill(int(pid), syscall.SIGTERM)
			fmt.Fprintln(os.Stderr, "убить процесс ", pid)
			return nil
		})
	case "ps":
		return nil, newEmbeddedCommand(stdin, stdout, func(_ io.Reader, stdout io.Writer) error {
			files, err := ioutil.ReadDir("/proc/")
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				if !file.IsDir() {
					continue
				}
				pid, err := strconv.ParseInt(file.Name(), 10, 32)
				if err != nil {
					continue
				}
				b, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/comm", pid))
				if err != nil {
					continue
				}
				fmt.Fprintln(stdout, pid, string(b))
			}

			return nil
		})
	default:
		cmd := exec.Command(command[0], command[1:]...)
		// stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Stdin = stdin
		cmd.Stdout = stdout
		cmd.Stderr = os.Stderr
		// err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		return nil, cmd
	}
	return nil, nil
}
func main() {

	fmt.Println("Ghbdtn")
	// arr := make([][]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		// Сканирует строку из Stdin (Консоль)
		ok := scanner.Scan()
		if !ok {
			return
		}
		// Содержит отсканированный текст строки

		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		commands := splitString(text)
		var nextStdin io.ReadCloser = os.Stdin
		cmds := []ICommands{}
		// get rid if the pipes array
		pipes := []io.Closer{}
		for idx, v := range commands {
			var stdin io.ReadCloser
			var stdout io.WriteCloser

			stdin = nextStdin
			if idx+1 == len(commands) {
				stdout = os.Stdout
				nextStdin = nil
			} else {
				nextStdin, stdout = io.Pipe()
				pipes = append(pipes, stdin)
				pipes = append(pipes, stdout)
			}

			_, cmd := getСommandFromLine(v, stdout, stdin)
			cmds = append(cmds, cmd)
		}
		for idx, cmd := range cmds {
			if err := cmd.Start(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				if idx != 0 && idx+1 != len(cmds) {
					pipes[idx*2].Close()
				}
				if idx+1 != len(cmds) {
					pipes[idx*2+1].Close()
				}
			}
		}
		for idx, cmd := range cmds {
			if err := cmd.Wait(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			if idx != 0 && idx+1 != len(cmds) {
				pipes[idx*2].Close()
			}
			if idx+1 != len(cmds) {
				pipes[idx*2+1].Close()
			}
		}
	}
}
