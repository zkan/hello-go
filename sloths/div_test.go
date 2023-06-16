package sloths

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRand struct{ mock.Mock }

func (m *mockRand) randomInt(max int) int {
	args := m.Called(max)
	return args.Int(0)
}

func newMockRand() *mockRand { return &mockRand{} }

func TestDivByRand(t *testing.T) {
	// get our mockRand
	m := newMockRand()
	// specify our return value. Since the code in divByRand
	// passes 10 into randomInt, we pass 10 in as the argument
	// to go with randomInt, and specify that we want the
	// method to return 6.
	m.On("randomInt", 10).Return(6)

	// now run divByRand and assert that we got back the
	// return value we expected, just like in a Go test that
	// doesn't use Testify Mock.
	quotient := divByRand(30, m)
	expected := 5
	if quotient != expected {
		t.Errorf("expected quotient to be %d, got %d", expected, quotient)
	}

	// check that randomInt was called with the number 10;
	// if not then the test fails
	m.AssertCalled(t, "randomInt", 10)
}
