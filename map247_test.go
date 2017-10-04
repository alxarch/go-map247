package map247_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alxarch/go-map247"
	"github.com/stretchr/testify/assert"
)

func Test_TimeIndex(t *testing.T) {
	i := map247.Index(3, 15)
	assert.Equal(t, 3*24+15, i)
	tm := time.Date(2017, time.March, 14, 15, 0, 0, 0, time.UTC)
	assert.Equal(t, 2*24+15, map247.TimeIndex(tm))

}

func Test_Mask(t *testing.T) {
	m := map247.Maskb{}
	m = m.Set(time.Wednesday, 5, true)
	if !m.Get(time.Wednesday, 5) {
		t.Errorf("Invalid value %d", m[int(time.Wednesday)])
	}
	m = m.Set(time.Wednesday, 5, false)
	if m.Get(time.Wednesday, 5) {
		t.Errorf("Invalid inset %d", m[int(time.Wednesday)])
	}
	m = map247.MaskFromIntSliceDayHour([]int{0, 2, 0, 4})
	assert.Equal(t, uint8(0), m[0])
	assert.Equal(t, uint8(1), m[1])
	assert.Equal(t, uint8(0), m[2])
	assert.Equal(t, uint8(1), m[3])
	m.String()

}
func Test_Float64JSON(t *testing.T) {
	data := []byte("[1,2,3]")
	f := &map247.Float64{}
	err := json.Unmarshal(data, f)
	assert.NoError(t, err)
	assert.Equal(t, []float64{1, 2, 3}, f.Index[:3])
}
func Test_Uint64JSON(t *testing.T) {
	data := []byte("[1,2,3]")
	f := &map247.Uint64{}
	err := json.Unmarshal(data, f)
	assert.NoError(t, err)
	assert.Equal(t, []uint64{1, 2, 3}, f.Index[:3])
}
