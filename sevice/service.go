package sevice

import (
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"os"
	"path/filepath"
	"time"
)

type program struct {
}

func (p *program) Start(s service.Service) error {
	logger.Info("Starting")
	//go  p.run()
	go p.run()
	return nil
}
func (p *program) run() {
	path := GetCurrentDirectory()
	f, err := os.OpenFile("F:\\1.TXT", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logger.Error(err)
	}
	defer f.Close()
	fmt.Println(123)

	for {
		fmt.Println(time.Now().Format("2006-01-02 15:03:04"))
		time.Sleep(3 * time.Second)
		f.WriteString(time.Now().Format("2006-01-02 15:03:04"))
		f.WriteString(path)
		f.WriteString(`
`)
	}
	return
}
func (p *program) Stop(s service.Service) error {
	logger.Info("Stopping")
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//return strings.Replace(dir, "\\", "/", -1) //将\替换成/
	return dir
}

var logger service.Logger

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()
	fmt.Println(1)
	flag.Parse()
	path := GetCurrentDirectory()

	svcConfig := &service.Config{
		Name:             "111111",
		DisplayName:      "111111",
		Description:      "地磅系统",
		WorkingDirectory: path,
	}
	prg := &program{}
	s, err := service.ChosenSystem().New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger, err = s.Logger(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			fmt.Printf("Valid actions: %q\n", service.ControlAction)
			fmt.Println(err)
			return
		}
		return
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = s.Install()
			if err != nil {
				fmt.Println("服务安装失败")
				fmt.Println(err)
				return
			}
			fmt.Println("服务安装成功")
			return
		}

		if os.Args[1] == "remove" {
			err = s.Uninstall()
			if err != nil {
				fmt.Println("服务卸载失败")
				fmt.Println(err)
				return
			}
			fmt.Println("服务卸载成功")
			return
		}
		if os.Args[1] == "restart" {
			err = s.Restart()
			if err != nil {
				fmt.Println("服务重启失败")
				fmt.Println(err)
				return
			}
			fmt.Println("服务重启成功")
			return
		}
		if os.Args[1] == "stop" {
			err = s.Stop()
			if err != nil {
				fmt.Println("服务停止失败")
				fmt.Println(err)
				return
			}
			fmt.Println("服务停止成功")
			return
		}
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
