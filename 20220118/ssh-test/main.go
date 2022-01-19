package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"time"
)

func publicKeyAuthFunc(kPath string)ssh.AuthMethod{
	key,err := ioutil.ReadFile(kPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func main()  {
	//可以使用 password 或者 sshkey 2种方式来认证。
	sshHost := "10.252.39.6" // 主机名
	sshUser := "root"     //用户名
	sshPassword := "root" //密码
	sshType := "password"   //ssh认证类型
	sshKeyPath := ""        //ssh id_rsa.id路径
	sshPort := 22

	//创建ssh登陆配置
	config := &ssh.ClientConfig{
		Timeout: time.Second, //ssh 连接timeout时间一秒钟，如果ssh验证错误 会在1秒内返回
		User: sshUser, //指定ssh连接用户
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以，但是不够安全
	}

	if sshType == "password"{
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	}else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	//dial获取ssh Client
	addr := fmt.Sprintf("%s:%d",sshHost,sshPort)
	sshClient,err := ssh.Dial("tcp",addr,config)
	if err != nil {
		log.Fatal("创建ssh client 失败",err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session,err := sshClient.NewSession()
	if err != nil{
		log.Fatal("创建ssh session 失败",err)
	}
	defer  session.Close()

	//执行远程命令
	combo,err := session.CombinedOutput("whoami; cd /root/abc; ls -al")
	if err != nil{
		log.Fatal("远程执行cmd 失败",err)
	}
	log.Println("命令输出:",string(combo))
}
