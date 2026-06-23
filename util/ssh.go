package util

import (
	"bafa/global"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func SSHCommand(host, port, user, password, command string) error {
	return sshCommand(host, port, user, password, command, "")
}

func SSHCommandWithStdin(host, port, user, password, command, stdin string) error {
	return sshCommand(host, port, user, password, command, stdin)
}

func sshCommand(host, port, user, password, command, stdin string) error {
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
	if stdin != "" {
		session.Stdin = strings.NewReader(stdin)
	}
	err = session.Run(command)
	if err != nil {
		return fmt.Errorf("ssh run failed: %w", err)
	}
	return nil
}
