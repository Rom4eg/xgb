// Package xcmisc is the X client API for the XC-MISC extension.
package xcmisc

// This file is automatically generated from xc_misc.xml. Edit at your peril!

import (
	"github.com/Rom4eg/xgb"

	"github.com/Rom4eg/xgb/xproto"
)

// Init must be called before using the XC-MISC extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 7, "XC-MISC").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named XC-MISC could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["XC-MISC"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["XC-MISC"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["XC-MISC"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["XC-MISC"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["XC-MISC"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// GetVersionCookie is a cookie used only for GetVersion requests.
type GetVersionCookie struct {
	*xgb.Cookie
}

// GetVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetVersionCookie.Reply()
func GetVersion(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) GetVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XC-MISC"]; !ok {
		panic("Cannot issue request 'GetVersion' using the uninitialized extension 'XC-MISC'. xcmisc.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return GetVersionCookie{cookie}
}

// GetVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) GetVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XC-MISC"]; !ok {
		panic("Cannot issue request 'GetVersion' using the uninitialized extension 'XC-MISC'. xcmisc.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return GetVersionCookie{cookie}
}

// GetVersionReply represents the data returned from a GetVersion request.
type GetVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

// Reply blocks and returns the reply data for a GetVersion request.
func (cook GetVersionCookie) Reply() (*GetVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getVersionReply(buf), nil
}

// getVersionReply reads a byte slice into a GetVersionReply value.
func getVersionReply(buf []byte) *GetVersionReply {
	v := new(GetVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.ServerMinorVersion = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for GetVersion
// getVersionRequest writes a GetVersion request to a byte slice.
func getVersionRequest(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XC-MISC"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], ClientMajorVersion)
	b += 2

	xgb.Put16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}

// GetXIDListCookie is a cookie used only for GetXIDList requests.
type GetXIDListCookie struct {
	*xgb.Cookie
}

// GetXIDList sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetXIDListCookie.Reply()
func GetXIDList(c *xgb.Conn, Count uint32) GetXIDListCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XC-MISC"]; !ok {
		panic("Cannot issue request 'GetXIDList' using the uninitialized extension 'XC-MISC'. xcmisc.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getXIDListRequest(c, Count), cookie)
	return GetXIDListCookie{cookie}
}

// GetXIDListUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetXIDListUnchecked(c *xgb.Conn, Count uint32) GetXIDListCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XC-MISC"]; !ok {
		panic("Cannot issue request 'GetXIDList' using the uninitialized extension 'XC-MISC'. xcmisc.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getXIDListRequest(c, Count), cookie)
	return GetXIDListCookie{cookie}
}

// GetXIDListReply represents the data returned from a GetXIDList request.
type GetXIDListReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	IdsLen uint32
	// padding: 20 bytes
	Ids []uint32 // size: xgb.Pad((int(IdsLen) * 4))
}

// Reply blocks and returns the reply data for a GetXIDList request.
func (cook GetXIDListCookie) Reply() (*GetXIDListReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getXIDListReply(buf), nil
}

// getXIDListReply reads a byte slice into a GetXIDListReply value.
func getXIDListReply(buf []byte) *GetXIDListReply {
	v := new(GetXIDListReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.IdsLen = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Ids = make([]uint32, v.IdsLen)
	for i := 0; i < int(v.IdsLen); i++ {
		v.Ids[i] = xgb.Get32(buf[b:])
		b += 4
	}

	return v
}

// Write request to wire for GetXIDList
// getXIDListRequest writes a GetXIDList request to a byte slice.
func getXIDListRequest(c *xgb.Conn, Count uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XC-MISC"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Count)
	b += 4

	return buf
}

// GetXIDRangeCookie is a cookie used only for GetXIDRange requests.
type GetXIDRangeCookie struct {
	*xgb.Cookie
}

// GetXIDRange sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetXIDRangeCookie.Reply()
func GetXIDRange(c *xgb.Conn) GetXIDRangeCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XC-MISC"]; !ok {
		panic("Cannot issue request 'GetXIDRange' using the uninitialized extension 'XC-MISC'. xcmisc.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getXIDRangeRequest(c), cookie)
	return GetXIDRangeCookie{cookie}
}

// GetXIDRangeUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetXIDRangeUnchecked(c *xgb.Conn) GetXIDRangeCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XC-MISC"]; !ok {
		panic("Cannot issue request 'GetXIDRange' using the uninitialized extension 'XC-MISC'. xcmisc.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getXIDRangeRequest(c), cookie)
	return GetXIDRangeCookie{cookie}
}

// GetXIDRangeReply represents the data returned from a GetXIDRange request.
type GetXIDRangeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	StartId uint32
	Count   uint32
}

// Reply blocks and returns the reply data for a GetXIDRange request.
func (cook GetXIDRangeCookie) Reply() (*GetXIDRangeReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getXIDRangeReply(buf), nil
}

// getXIDRangeReply reads a byte slice into a GetXIDRangeReply value.
func getXIDRangeReply(buf []byte) *GetXIDRangeReply {
	v := new(GetXIDRangeReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.StartId = xgb.Get32(buf[b:])
	b += 4

	v.Count = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for GetXIDRange
// getXIDRangeRequest writes a GetXIDRange request to a byte slice.
func getXIDRangeRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XC-MISC"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}
