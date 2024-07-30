package gtl

import "testing"

func TestAtomicPointer(t *testing.T) {
	var p AtomicPointer[int]
	if v := p.Load(); v != nil {
		t.Fatalf("initial value = %v, want nil", v)
	}
	vv1 := 42
	p.Store(&vv1)
	if v := p.Load(); *v != 42 {
		t.Fatalf("value = %v, want 42", *v)
	}
	vv2 := 84
	p.Store(&vv2)
	if v := p.Load(); *v != 84 {
		t.Fatalf("value = %v, want 84", *v)
	}
}
