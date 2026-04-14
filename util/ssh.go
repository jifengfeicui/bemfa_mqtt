package util

import (
	"bafa/global"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func SSHCommand(host, port, user, password, command string) error {
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return fmt.Errorf("ssh dial failed: %w", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("ssh session failed: %w", err)
	}
	defer session.Close()

	global.SugarLogger.Info("SSH执行: " + command)
	err = session.Run(command)
	if err != nil {
		return fmt.Errorf("ssh run failed: %w", err)
	}
	return nil
}
