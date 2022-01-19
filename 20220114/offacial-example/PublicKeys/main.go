package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	var hostKey ssh.PublicKey
	hostKey , _ = ioutil.ReadFile("./key_10.253.21.252_22.pub")
	//if err != nil {
	//	log.Fatalf("unable to read pub key: %v", err)
	//}

	//hostKey, err = ssh.ParsePublicKey(pubkey)
	//if err != nil {
	//	log.Fatalf("unable to parse public key: %v", err)
	//}


	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	key, err := ioutil.ReadFile("./projectmanager.txt")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}



	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", "10.253.21.252:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()
}
