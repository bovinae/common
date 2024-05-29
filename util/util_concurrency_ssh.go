package util

import (
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type ConcurrencySsh struct {
	config *SshConfig
}

type SshConfig struct {
	Hosts    []string
	User     string
	Password string
	Type     string
	Port     int

	Timeout int // ssh timeout, unit: second
}

func NewConcurrencySsh(config *SshConfig) *ConcurrencySsh {
	return &ConcurrencySsh{config}
}

func (cs *ConcurrencySsh) Call(cmd string) *sync.Map {
	respMap := cs.call(cmd, func(host, cmd string) any {
		resp, err := cs.execCmd(host, cmd)
		if err != nil {
			return err
		}
		return resp
	})
	return respMap
}

func (cs *ConcurrencySsh) call(cmd string, f func(host, cmd string) any) *sync.Map {
	var respMap sync.Map
	var wg sync.WaitGroup
	for _, host := range cs.config.Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			respMap.Store(host, f(host, cmd))
		}(host)
	}
	wg.Wait()
	return &respMap
}

func (cs *ConcurrencySsh) execCmd(host, cmd string) (string, error) {
	cli, err := cs.initSshClient(host)
	if err != nil {
		return "", err
	}
	if cli == nil {
		return "", errors.New("ssh client is nil")
	}
	session, err := cli.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "sshClient.NewSession")
	}
	combo, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", errors.Wrap(err, "session.CombinedOutput")
	}
	return string(combo), nil
}

func (cs *ConcurrencySsh) initSshClient(host string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Duration(cs.config.Timeout) * time.Second,
		User:            cs.config.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // not safe enough
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(cs.config.Password)}

	addr := fmt.Sprintf("%s:%d", host, cs.config.Port)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}

	return sshClient, nil
}
