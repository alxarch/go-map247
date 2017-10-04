package map247

import (
	"encoding/json"
	"time"
)

const ScheduleSize = 24 * 7

func Index(d time.Weekday, h int) int {
	return int(d)*24 + h
}
func TimeIndex(t time.Time) int {
	return int(t.Weekday())*24 + t.Hour()
}

type Maskb [24]uint8

func (m Maskb) Set(d time.Weekday, h int, b bool) Maskb {
	if 0 <= h && h <= 24 {
		bit := uint8(1) << uint8(d)
		if b {
			m[h] |= bit
		} else {
			m[h] &^= bit
		}
	}
	return m
}

func (m Maskb) Match(t time.Time) bool {
	return m.get(t.Weekday(), t.Hour())
}

func (m Maskb) get(d time.Weekday, h int) bool {
	return m[h]&(1<<uint8(d)) != 0
}

func (m Maskb) Get(d time.Weekday, h int) bool {
	if 0 <= h && h < 24 {
		return m.get(d, h)
	}
	return false
}

func (m Maskb) Day(d time.Weekday) []bool {
	day := [24]bool{}
	h := uint8(d)
	for i := 0; i < 24; i++ {
		day[i] = (m[i] & 1 << h) == m[i]
	}
	return day[:]
}

func MaskFromIntSliceDayHour(s []int) (m Maskb) {
	if s == nil {
		return
	}
	for i := 0; i < 24; i++ {
		for j := 0; j < 7; j++ {
			n := j*24 + i
			if n < len(s) && s[n] != 0 {
				m[i] |= 1 << uint8(j)
			}
		}
	}
	return
}

var days = []byte("SMTWTFS")
var header = []byte("H 00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23\n")

func (m Maskb) String() string {
	w := []byte{}
	w = append(w, header...)
	for i := 0; i < 7; i++ {
		w = append(w, days[i], ' ')
		for j := 0; j < 24; j++ {
			v := m[j] & (1 << uint8(i))
			if v == 1 {
				v = ' '
			} else {
				v = days[i]
			}
			w = append(w, v, v, ' ')
		}
		w = append(w, '\n')
	}
	return string(w)
}

// func MaskFromBoolSlice(s []bool) (m Mask) {
// 	if s == nil {
// 		return
// 	}
// 	n := len(s)
// 	if n > ScheduleSize {
// 		n = ScheduleSize
// 	}
// 	for i := 0; i < n; i++ {
// 		if s[i] {
// 			m |= 1 << Mask(i)
// 		}
// 	}
// 	return
// }

type Float64 struct {
	Index [ScheduleSize]float64
}

func (m *Float64) Get(t time.Time) float64 {
	if m != nil {
		return m.Index[TimeIndex(t)]
	}
	return 0
}

func (m *Float64) SetAll(v float64) {
	for i := 0; i < ScheduleSize; i++ {
		m.Index[i] = v
	}
}
func (m *Float64) Set(d time.Weekday, h int, v float64) {
	if m != nil {
		m.Index[Index(d, h)] = v
	}
}

func (m *Float64) UnmarshalJSON(data []byte) (err error) {
	values := make([]float64, 0, ScheduleSize)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	copy(m.Index[:], values)
	return nil
}
func (m Maskb) Empty() bool {
	for i := 0; i < 24; i++ {
		if m[i] != 0 {
			return false
		}
	}
	return true
}
func NewMask(b []byte) (m Maskb) {
	return m.SetBytes(b)
}
func (m Maskb) SetBytes(b []byte) Maskb {
	if b == nil {
		return m
	}
	for i, h := range b {
		if i < 24 {
			m[i] = h
		} else {
			break
		}
	}
	return m
}

type Uint64 struct {
	Index [ScheduleSize]uint64
}

func (m *Uint64) Get(t time.Time) uint64 {
	if m != nil {
		return m.Index[TimeIndex(t)]
	}
	return 0
}

func (m *Uint64) SetAll(v uint64) {
	for i := 0; i < ScheduleSize; i++ {
		m.Index[i] = v
	}
}
func (m *Uint64) Set(d time.Weekday, h int, v uint64) {
	if m != nil {
		m.Index[Index(d, h)] = v
	}
}

func (m *Uint64) UnmarshalJSON(data []byte) (err error) {
	values := make([]uint64, 0, ScheduleSize)
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	copy(m.Index[:], values)
	return nil
}
