package shell

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"restiq/restic"

	escapes "github.com/snugfox/ansi-escapes"
)

type Shell struct {
	Bin  string
	Env  string
	Args []string
	Repo string
	// Dirs []string
}

func New(bin string, env string) *Shell {
	return &Shell{Bin: bin, Env: env}
}

func NewRestic(bin string) *Shell {
	return &Shell{Bin: bin}
}

func (s *Shell) WithArgs(cmdArgs []string) *Shell {
	s.Args = append(s.Args, cmdArgs...)
	fmt.Println("WithArgs", cmdArgs)
	return s
}

func (s *Shell) WithResticSubCmd(cmdArgs string) *Shell {
	s.Args = append(s.Args, "-r", s.Repo, cmdArgs)
	return s
}

func (s *Shell) WithRepo(repo string) *Shell {
	s.Repo = repo
	return s
}

func (s *Shell) WithEnv(kv string) *Shell {
	s.Env = kv
	return s
}

func (s *Shell) WithResticOptionalArg(args []string) *Shell {
	var optArg string
	if len(args) == 2 {
		optArg = args[1]
	} else {
		optArg = "latest"
	}
	s.Args = append(s.Args, optArg)
	return s
}

// func (s *Shell) Notify() *Shell {
// 	app := pushover.New("uQiRzpo4DXghDmr9QzzfQu27cmVRsG")

// 	// Create a new recipient
// 	recipient := pushover.NewRecipient("gznej3rKEVAvPUxu9vvNnqpmZpokzF")

// }

func (s *Shell) Run() *Shell {
	fmt.Println("IN RUN")
	cmd := exec.Command(s.Bin, s.Args...)
	cmd.Env = os.Environ()
	if len(s.Env) > 1 {
		cmd.Env = append(cmd.Env, s.Env)
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	fmt.Printf("Will run: %s %+q\n", s.Bin, s.Args)

	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return s
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return s
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		fmt.Fprintf(os.Stderr, "Exit Code: %d\n", exitError.ExitCode())
		// if exit code, there must not be a repo there
		// assume that for now
	}

	return s
}

func (s *Shell) RunWithJson() *Shell {
	fmt.Print(escapes.EraseScreen)
	fmt.Print(escapes.CursorPos(1, 1))

	fmt.Println("IN RUN JSON")

	cmd := exec.Command(s.Bin, s.Args...)
	cmd.Env = os.Environ()
	if len(s.Env) > 1 {
		cmd.Env = append(cmd.Env, s.Env)
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	var resticJSON restic.ResticBackUp

	go func() {
		for scanner.Scan() {
			err := json.Unmarshal(scanner.Bytes(), &resticJSON)
			if err != nil {
				// fmt.Println(scanner.Text())
				if scanner.Text() == "Is there a repository at the following location?" {
					fmt.Printf("Need to init repo, run:\n%s init -r %s", s.Bin, s.Repo)
				}
			}
			if resticJSON.MessageType == "status" {
				// fmt.Print(`\033[A\033[2K`)
				fmt.Print(escapes.TextDeleteLine)

				fmt.Printf("%-03.02f: Total Bytes: %d / %d - Files: %d\r",
					resticJSON.PercentDone*100, resticJSON.BytesDone, resticJSON.TotalBytes, resticJSON.FilesDone)
				for _, v := range resticJSON.CurrentFiles {
					fmt.Print(escapes.TextDeleteLine)
					fmt.Printf("\tCurrent Files: %s\r",
						v)
				}

			} else if resticJSON.MessageType == "summary" {
				fmt.Print(escapes.TextDeleteLine)
				fmt.Printf("\nSummary: %-03.02f: Total Bytes: %d / %d - Files: %d\n",
					resticJSON.PercentDone*100, resticJSON.BytesDone, resticJSON.TotalBytes, resticJSON.FilesDone)
				return
			}
		}
	}()

	fmt.Printf(">>> Will run: %s %+q\n", s.Bin, s.Args)
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return s
	}

	err = cmd.Wait()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
	// 	return s
	// }

	if exitError, ok := err.(*exec.ExitError); ok {
		fmt.Fprintf(os.Stderr, "%v", scanner.Text())
		fmt.Fprintf(os.Stderr, "\nExit Code: %d\n", exitError.ExitCode())
		//os.Exit(exitError.ExitCode())
		// cmdStrInit := []string{C.ResticBin, "-r", fmt.Sprintf("%s:%s", vv.Type, vv.Location), "stats"}
	}
	return s
}

func (s *Shell) RunRestic() *Shell {
	cmd := exec.Command(s.Bin, s.Args...)
	cmd.Env = os.Environ()
	if len(s.Env) > 1 {
		cmd.Env = append(cmd.Env, s.Env)
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	fmt.Printf("Will run: %s %v\n", s.Bin, s.Args)
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return s
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		// 	return
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		fmt.Fprintf(os.Stderr, "Exit Code: %d\n", exitError.ExitCode())
		//os.Exit(exitError.ExitCode())
		// cmdStrInit := []string{C.ResticBin, "-r", fmt.Sprintf("%s:%s", vv.Type, vv.Location), "stats"}
	}
	return s
}
