package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func main()  {
	sshHost := "10.252.39.6"
	sshUser := "root"
	sshPasswrod := "root"
	sshType := "password"  // password或者key
	//sshKeyPath := "" // ssh id_rsa.id路径
	sshPort := 22

	// 创建ssh登录配置
	config := &ssh.ClientConfig{

		Timeout: time.Second, // ssh连接time out时间一秒钟,如果ssh验证错误会在一秒钟返回
		User: sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),  // 这个可以,但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
		//Auth
	}

	// sshType ssh类型
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPasswrod)}
	} else {
		//config.Auth = []ssh.AuthMethod(publicKeyAuthFunc(sshKeyPath))
		return
	}

	// dial 获取ssh client
	addr := fmt.Sprintf("%s:%d",sshHost,sshPort)
	fmt.Println("================= addr = ", addr)

	sshClient,err := ssh.Dial("tcp",addr,config)
	if err != nil {
		log.Fatal("创建ssh client 失败",err)
	}
	defer sshClient.Close()


	// 创建ssh-session
	session,err := sshClient.NewSession()
	if err != nil {
		//log.Fatal("创建ssh session失败",err)
		fmt.Println("创建ssh session失败", err)
	}

	defer session.Close()



	// 执行远程命令
	combo,err := session.CombinedOutput("whoami; cd /root/abc; ls -al;")
	if err != nil {
		log.Fatal("远程执行cmd失败",err)
	}
	log.Println("命令输出:",string(combo))


}