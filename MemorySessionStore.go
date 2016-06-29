package goutils

import (
	"errors"
	"time"
)

// MemorySessionStore 内存session
type MemorySessionStore struct {
	SessMap map[string]Session
}

// NewMemorySessionStore 初始化
func NewMemorySessionStore() *MemorySessionStore {
	return &MemorySessionStore{SessMap: make(map[string]Session, 256)}
}

// Get 获取session
func (ms *MemorySessionStore) Get(sessionID string) (Session, error) {
	if s, ok := ms.SessMap[sessionID]; ok {
		return s, nil
	}
	return Session{}, errors.New(ErrSessionNotExist)
}

// Delete 销毁session
func (ms *MemorySessionStore) Delete(sessionID string) error {
	if _, ok := ms.SessMap[sessionID]; ok {
		delete(ms.SessMap, sessionID)
		return nil
	}
	return errors.New(ErrSessionNotExist)
}

// TTL  设置过期
func (ms *MemorySessionStore) TTL(sessionID string, maxAge time.Duration) error {
	if s, ok := ms.SessMap[sessionID]; ok {
		s.MaxAge = maxAge
		s.newly = time.Now().UnixNano()
		ms.SessMap[sessionID] = s
		return nil
	}
	return errors.New(ErrSessionNotExist)
}

// SetData 设置session数据
func (ms *MemorySessionStore) SetData(sessionID string, data interface{}) error {
	if s, ok := ms.SessMap[sessionID]; ok {
		s.Data = data
		ms.SessMap[sessionID] = s
		return nil
	}
	return errors.New(ErrSessionNotExist)
}

func (ms *MemorySessionStore) checkExpire() {
	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 2):
				now := time.Now().UnixNano()
				for id, sess := range ms.SessMap {
					if (sess.newly + int64(sess.MaxAge)) <= now {
						delete(ms.SessMap, id)
					}
				}
			}
			time.Sleep(time.Millisecond)
		}
	}()
}
