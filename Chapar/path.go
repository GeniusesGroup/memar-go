/* For license and copyright information please see LEGAL file in repository */

package chapar

// Path indicate Chapar switch route plan!
type Path struct {
	path [MaxHopCount]byte
	len  byte
}

func (p *Path) Set(path []byte) {
	copy(p.path[:], path)
	p.len = byte(len(path))
}

func (p *Path) Get() []byte {
	return p.path[:p.len]
}

func (p *Path) GetAsString() string {
	return string(p.path[:p.len])
}

func (p *Path) Len() int {
	return int(p.len)
}

func (p *Path) LenAsByte() byte {
	return p.len
}

// ReadFrom sets path from the given frame
func (p *Path) ReadFrom(frame []byte) {
	copy(p.path[:], frame[FixedHeaderLength:FixedHeaderLength+GetHopCount(frame)])
	p.len = GetHopCount(frame)
}

// WriteTo sets the path in given the frame.
func (p *Path) WriteTo(frame []byte) {
	copy(frame[FixedHeaderLength:], p.path[:])
}

// GetReverse return path in all hops just in reverse.
func (p *Path) GetReverse() (reverse Path) {
	reverse = Path{
		path: ReversePath(p.path[:]),
		len:  p.len,
	}
	return
}
