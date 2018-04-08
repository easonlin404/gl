package gl

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {

	assert.Equal(t,Encode(0),"a")
	assert.Equal(t,Encode(125),"cb")
	assert.Equal(t,Encode(273984),"bjrg")
	assert.Equal(t,Encode(3792586),"p4MU")
}

func TestBase62(t *testing.T) {
	assert.Equal(t,Base62(138), []int{14, 2})
	assert.Equal(t,Base62(125), []int{1,2})
	assert.Equal(t,Base62(0), []int{0})
}
func TestDecode(t *testing.T) {

	assert.Equal(t,Decode("cb"),125)
	assert.Equal(t,Decode("e9a"),19158)
}