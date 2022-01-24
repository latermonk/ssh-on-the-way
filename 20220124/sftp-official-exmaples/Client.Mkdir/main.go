package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Example of mimicing 'mkdir --parents'; I.E. recursively create
	// directoryies and don't error if any directories already exists.
	var conn *ssh.Client

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	sshFxFailure := uint32(4)
	mkdirParents := func(client *sftp.Client, dir string) (err error) {
		var parents string

		if path.IsAbs(dir) {
			// Otherwise, an absolute path given below would be turned in to a relative one
			// by splitting on "/"
			parents = "/"
		}

		for _, name := range strings.Split(dir, "/") {
			if name == "" {
				// Paths with double-/ in them should just move along
				// this will also catch the case of the first character being a "/", i.e. an absolute path
				continue
			}
			parents = path.Join(parents, name)
			err = client.Mkdir(parents)
			if status, ok := err.(*sftp.StatusError); ok {
				if status.Code == sshFxFailure {
					var fi os.FileInfo
					fi, err = client.Stat(parents)
					if err == nil {
						if !fi.IsDir() {
							return fmt.Errorf("file exists: %s", parents)
						}
					}
				}
			}
			if err != nil {
				break
			}
		}
		return err
	}

	err = mkdirParents(client, "/tmp/foo/bar")
	if err != nil {
		log.Fatal(err)
	}
}
