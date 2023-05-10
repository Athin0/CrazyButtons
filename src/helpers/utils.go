package utils

import "sync"

type LastCommand struct {
	UserCommand map[int64]string
	mx          *sync.RWMutex
}

func NewLastCommand() *LastCommand {
	return &LastCommand{UserCommand: make(map[int64]string), mx: &sync.RWMutex{}}
}

func (c LastCommand) Get(userId int64) string {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.UserCommand[userId]
}

func (c LastCommand) Set(userId int64, command string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.UserCommand[userId] = command
}
