package main

import (
	"errors"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli"
)

var (
	withBM      bool
	withGRPC    bool
	withSwagger bool
)

func protocAction(ctx *cli.Context) (err error) {
	if err = checkProtoc(); err != nil {
		return err
	}
	if !withGRPC && !withBM && !withSwagger {
		withBM = true
		withGRPC = true
		withSwagger = true
	}
	if withBM {
		if err = installBMGen(); err != nil {
			return
		}
		if err = genBM(ctx); err != nil {
			return
		}
	}
	if withGRPC {
		if err = installGRPCGen(); err != nil {
			return err
		}
		if err = genGRPC(ctx); err != nil {
			return
		}
	}
	if withSwagger {
		if err = installSwaggerGen(); err != nil {
			return
		}
		if err = genSwagger(ctx); err != nil {
			return
		}
	}
	log.Printf("generate %v success.\n", ctx.Args())
	return nil
}

func checkProtoc() error {
	if _, err := exec.LookPath("protoc"); err != nil {
		switch runtime.GOOS {
		case "darwin":
			fmt.Println("brew install protobuf")
			cmd := exec.Command("brew", "install", "protobuf")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				return err
			}
		case "linux":
			fmt.Println("snap install --classic protobuf")
			cmd := exec.Command("snap", "install", "--classic", "protobuf")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				return err
			}
		default:
			return errors.New("您还没安装protobuf，请进行手动安装：https://github.com/protocolbuffers/protobuf/releases")
		}
	}
	return nil
}

func generate(ctx *cli.Context, protoc string) error {
	pwd, _ := os.Getwd()
	gosrc := path.Join(gopath(), "src")
	ext, err := latestKratos()
	if err != nil {
		return err
	}
	line := fmt.Sprintf(protoc, gosrc, ext, pwd)
	//log.Println(line, strings.Join(ctx.Args(), " "))
	args := strings.Split(line, " ")
	args = append(args, ctx.Args()...)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = pwd
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func goget(url string) error {
	args := strings.Split(url, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(url)
	return cmd.Run()
}

func latestKratos() (string, error) {
	gopath := gopath()
	ext := path.Join(gopath, "src/github.com/inside-the-mirror/kratos/third_party")
	if _, err := os.Stat(ext); !os.IsNotExist(err) {
		return ext, nil
	}
	baseMod := path.Join(gopath, "pkg/mod/github.com/bilibili")
	files, err := ioutil.ReadDir(baseMod)
	if err != nil {
		return "", err
	}
	for i := len(files) - 1; i >= 0; i-- {
		if strings.HasPrefix(files[i].Name(), "kratos@") {
			return path.Join(baseMod, files[i].Name(), "third_party"), nil
		}
	}
	return "", errors.New("not found kratos package")
}

func gopath() (gp string) {
	gopaths := strings.Split(os.Getenv("GOPATH"), ":")
	if len(gopaths) == 1 {
		return gopaths[0]
	}
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	abspwd, err := filepath.Abs(pwd)
	if err != nil {
		return
	}
	for _, gopath := range gopaths {
		absgp, err := filepath.Abs(gopath)
		if err != nil {
			return
		}
		if strings.HasPrefix(abspwd, absgp) {
			return absgp
		}
	}
	return build.Default.GOPATH
}
