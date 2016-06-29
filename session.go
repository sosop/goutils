package goutils

import (
	"time"
)

const (
	// DefaultAge 默认过期时间
	DefaultAge = time.Minute * 30
)

var (
	store SessionStore
	// ErrSessionNotExist session不存在错误
	ErrSessionNotExist = "Session not exist"
)

// Session session
type Session struct {
	SessionID string
	MaxAge    time.Duration
	Data      interface{}
	newly     int64
}

// NewSession 构造session
func NewSession(data interface{}) Session {
	return Session{SessionID: GenerateUUID(), MaxAge: DefaultAge, Data: data, newly: time.Now().UnixNano()}
}

// NewSessionStore 初始化session存储器
func NewSessionStore(sessionStore SessionStore) {
	store = sessionStore
}

// GetSession 获取session
func GetSession(sessionID string) (Session, error) {
	return store.Get(sessionID)
}

// SessionStore 存储session容器
type SessionStore interface {
	Get(sessionID string) (Session, error)
	Delete(seesionID string) error
	TTL(sessionID string, maxAge time.Duration) error
	SetData(sessionID string, data interface{}) error
}

// GetData 获取数据
func (s Session) GetData() interface{} {
	return s.GetData()
}

// SetData 设置数据
func (s Session) SetData(data interface{}) error {
	return store.SetData(s.SessionID, data)
}

// Destroy 销毁session
func (s Session) Destroy() error {
	return store.Delete(s.SessionID)
}

// SetMaxAge 设置过期时间
func (s Session) SetMaxAge(maxAge time.Duration) error {
	return store.TTL(s.SessionID, maxAge)
}

// Refresh 刷新最近访问
func (s Session) Refresh() {
	store.TTL(s.SessionID, s.MaxAge)
}
