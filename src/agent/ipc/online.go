package ipc

import "sync"
import . "types"

var _active map[int32]*Session
var _lock sync.RWMutex

//----------------------------------------------- register a user as online user
func RegisterOnline(sess *Session, id int32) {
	defer _lock.Unlock()
	_lock.Lock()

	_active[id] = sess
}

//----------------------------------------------- unregister a user from online users
func UnregisterOnline(id int32) {
	defer _lock.Unlock()
	_lock.Lock()

	delete(_active, id)
}

//----------------------------------------------- query a online user
func QueryOnline(id int32) *Session {
	defer _lock.RUnlock()
	_lock.RLock()

	return _active[id]
}

func init() {
	_active = make(map[int32]*Session)
}
