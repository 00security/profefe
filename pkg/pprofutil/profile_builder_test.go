package pprofutil

import (
	"testing"

	pprofProfile "github.com/00security/profefe/internal/pprof/profile"
	"github.com/00security/profefe/pkg/profile"
	"github.com/stretchr/testify/assert"
)

func TestProfileBuilder_IsEmpty(t *testing.T) {
	b := NewProfileBuilder(profile.TypeCPU)
	assert.True(t, b.IsEmpty())

	b.AddSample(&pprofProfile.Sample{})
	assert.False(t, b.IsEmpty())
}
