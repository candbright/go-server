package core

import (
	"fmt"
	"github.com/candbright/go-log/log"
	"github.com/candbright/server-mc/internal/mc-server/config"
	"github.com/pkg/errors"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/candbright/go-ssh/ssh"
)

type Process struct {
	Version string
	RootDir string
	Screen  Screen
	backup  bool
}

type ProcessConfig struct {
	Version string
	RootDir string
	Session ssh.Session
}

func NewProcess(cfg ProcessConfig) (*Process, error) {
	p := &Process{
		Version: cfg.Version,
		RootDir: cfg.RootDir,
		Screen: Screen{
			Session: cfg.Session,
			Name:    fmt.Sprintf("mc-%s", cfg.Version),
		},
	}
	return p, nil
}

func (p *Process) Active() bool {
	return p.Screen.Exists()
}

func (p *Process) ExecFile() string {
	return path.Join(p.RootDir, "bedrock_server")
}

func (p *Process) ScreenName() string {
	return fmt.Sprintf("mc-%s", p.Version)
}

func (p *Process) Start() error {
	if !p.Active() {
		err := p.Screen.Create()
		if err != nil {
			return err
		}
		err = p.Screen.ExecCmd("LD_LIBRARY_PATH=.", p.ExecFile())
		if err != nil {
			return err
		}
	}
	return errors.Errorf("screen %s already started!", p.ScreenName())
}

func (p *Process) Stop() error {
	if p.Active() {
		return p.Screen.Exit()
	}
	return errors.Errorf("screen %s not found!", p.ScreenName())
}

func (p *Process) Restart() error {
	err := p.Stop()
	if err != nil {
		return err
	}
	err = p.Start()
	if err != nil {
		return err
	}
	return nil
}

func (p *Process) ExecCmd(arg ...string) error {
	return p.Screen.ExecCmd(arg...)
}

func (p *Process) startBackup() {
	if !p.backup {
		p.backup = true
		go func() {
			for p.backup {
				p.backupTick()
				time.Sleep(time.Hour * 1)
			}
		}()
		go func() {
			time.Sleep(time.Hour * 1)
			for p.backup {
				p.backupClearTick()
				time.Sleep(time.Hour * 24)
			}
		}()
	}
}

func (p *Process) WorldsDir() string {
	return path.Join(p.RootDir, "worlds")
}

func (p *Process) BackupDir() string {
	return path.Join(config.ServerConfig.Get("mc.path"), fmt.Sprintf("backup-"+p.RootDir))
}

func (p *Process) stopBackup() {
	p.backup = false
}

func (p *Process) backupTick() {
	if !p.Active() {
		return
	}
	//zip data
	sourceDir := p.WorldsDir()
	backupDir := p.BackupDir()
	backupFile := fmt.Sprintf("%s/backup-%s.zip", backupDir, time.Now().Format("20060102-150405"))
	err := p.Screen.Session.MakeDirAll(backupDir, 0777)
	if err != nil {
		log.WithError(err).Error("make backup dir failed")
		return
	}
	err = p.Screen.Session.Run("zip", "-r", backupFile, sourceDir)
	if err != nil {
		log.WithError(err).Error("zip failed")
		return
	}
	log.Infof("backup has been saved to: %s\n", backupFile)
}

func (p *Process) backupClearTick() {
	if !p.Active() {
		return
	}
	// read dir
	files, err := p.Screen.Session.ReadDir(p.BackupDir())
	if err != nil {
		log.WithError(err).Error("read backup dir failed")
		return
	}
	getFileTime := func(name string) string {
		return strings.Split(strings.Split(name, "-")[1], ".")[0]
	}
	// sort files by time
	sort.Slice(files, func(i, j int) bool {
		fi1, _ := os.Stat(getFileTime(files[i].Name))
		fi2, _ := os.Stat(getFileTime(files[j].Name))
		return fi1.ModTime().Before(fi2.ModTime())
	})

	// remove previous backup file
	for i := 0; i < len(files)-24; i++ {
		e := p.Screen.Session.Remove(files[i].Name)
		if e != nil {
			log.WithError(e).Error("remove previous backup file failed")
		}
	}
}
