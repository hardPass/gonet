package types

import (
	"encoding/json"
)

import (
	"types/estates"
	"types/grid"
)

type IPCObject struct {
	Sender    int32 // sender id
	Multicast bool  // indicate wheather this message is a multicast message(group)
	Service   int16
	Object    []byte // json formatted object
	Time      int64  // send time
}

func (obj *IPCObject) Json() []byte {
	val, _ := json.Marshal(obj)
	return val
}

type Session struct {
	MQ chan IPCObject // Player's Internal Message Queue
	// user data
	User    User
	Estates estates.Manager
	Grid    *grid.Grid // Building's bitmap, online constructing...

	// session related
	LoggedIn bool // flag for weather the user is logged in
	KickOut  bool // flag for player is kicked out

	// time related
	ConnectTime    int64 // tcp connection establish time
	LastPacketTime int64 // last packet arrive time
	LastFlushTime  int64 // last flush to db time
	Dirty          bool  // mark the data as dirty
	OpCount        int   // num of operations since last sync
}
