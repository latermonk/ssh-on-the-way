package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	//"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//可以使用 password 或者 sshkey 2种方式来认证。
	sshHost := "10.252.39.6" // 主机名
	sshUser := "root"        //用户名
	sshPassword := "root"    //密码
	sshType := "password"    //ssh认证类型
	//sshKeyPath := ""          //ssh id_rsa.id路径
	sshPort := 22

	//创建ssh登陆配置
	config := &ssh.ClientConfig{
		Timeout: time.Second, //ssh 连接timeout时间一秒钟，如果ssh验证错误 会在1秒内返回
		User:    sshUser,     //指定ssh连接用户
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以，但是不够安全
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		//config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	//dial获取ssh Client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)


	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()

	//将当前终端的stdin文件句柄设置给远程给远程终端，这样就可以使用tab键

	fd := int(os.Stdin.Fd())
	fmt.Println("=========MakeRaw=====111======== fd", fd)

	state, err := terminal.MakeRaw(fd)

	fmt.Println("==============haha=======111111111111111111111111=========== state =", state)
	if err != nil {
		fmt.Println("=========MakeRaw=====222========")
		panic(err)
	}

	fmt.Println("==============haha======22222222222222222222222222============")

	defer terminal.Restore(fd, state)

	fmt.Println("==============haha=======3333333333333333333333333333===========")

	session.Stdout = os.Stdout // 会话输出关联到系统标准输出设备
	session.Stderr = os.Stderr // 会话错误输出关联到系统标准错误输出设备
	session.Stdin = os.Stdin   // 会话输入关联到系统标准输入设备

	//设置终端模式
	fmt.Println("set  terminal mode ")
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     //禁止回显 （0 禁止,1 启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}

	// 请求伪终端
	fmt.Println("===========request fake terminal======================")
	if err = session.RequestPty("linux", 32, 160, modes); err != nil {
		log.Fatalf("request pty error: %s", err.Error())
	}

	//启动远程shell
	fmt.Println("=========launch remote shell==================")
	if err = session.Shell(); err != nil {
		log.Fatalf("start shell error: %s", err.Error())
	}

	//等待远程命令（终端）退出
	fmt.Println("===================wait for remote disconection===========")
	if err = session.Wait(); err != nil {
		log.Fatalf("return error: %s", err.Error())
	}
}
