package slices

import "testing"

func TestC128Slice(t *testing.T) {
	sxp := C128List(1)
	switch {
	case sxp.Len() != 1:				t.Fatalf("C128List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.C128At(0) != 1:			t.Fatalf("C128List(1) element 0 should be 1 and not %v", sxp.C128At(0))
	}

	sxp = C128List(1, 2)
	switch {
	case sxp.Len() != 2:				t.Fatalf("C128List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.C128At(0) != 1:			t.Fatalf("C128List(1 2) element 0 should be 1 and not %v", sxp.C128At(0))
	case sxp.C128At(1) != 2:			t.Fatalf("C128List(1 2) element 1 should be 2 and not %v", sxp.C128At(1))
	}

	sxp = C128List(1, 2, 3)
	switch {
	case sxp.Len() != 3:				t.Fatalf("C128List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.C128At(0) != 1:			t.Fatalf("C128List(1 2 3) element 0 should be 1 and not %v", sxp.C128At(0))
	case sxp.C128At(1) != 2:			t.Fatalf("C128List(1 2 3) element 1 should be 2 and not %v", sxp.C128At(1))
	case sxp.C128At(2) != 3:			t.Fatalf("C128List(1 2 3) element 2 should be 3 and not %v", sxp.C128At(2))
	}
}

func TestC128SliceString(t *testing.T) {
	ConfirmString := func(s *C128Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(C128List(), "()")
	ConfirmString(C128List(0), "((0+0i))")
	ConfirmString(C128List(0, 1), "((0+0i) (1+0i))")
	ConfirmString(C128List(0, 1i), "((0+0i) (0+1i))")
}

func TestC128SliceLen(t *testing.T) {
	ConfirmLength := func(s *C128Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(C128List(0), 1)
	ConfirmLength(C128List(0, 1), 2)
}

func TestC128SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *C128Slice, i, j int, r *C128Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(C128List(0, 1, 2), 0, 1, C128List(1, 0, 2))
	ConfirmSwap(C128List(0, 1, 2), 0, 2, C128List(2, 1, 0))
}

func TestC128SliceCut(t *testing.T) {
	ConfirmCut := func(s *C128Slice, start, end int, r *C128Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 0, 1, C128List(1, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 1, 2, C128List(0, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 2, 3, C128List(0, 1, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 3, 4, C128List(0, 1, 2, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 4, 5, C128List(0, 1, 2, 3, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 5, 6, C128List(0, 1, 2, 3, 4))

	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), -1, 1, C128List(1, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 0, 2, C128List(2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 1, 3, C128List(0, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 2, 4, C128List(0, 1, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 3, 5, C128List(0, 1, 2, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 4, 6, C128List(0, 1, 2, 3))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 5, 7, C128List(0, 1, 2, 3, 4))
}

func TestC128SliceDelete(t *testing.T) {
	ConfirmCut := func(s *C128Slice, index int, r *C128Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), -1, C128List(0, 1, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 0, C128List(1, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 1, C128List(0, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 2, C128List(0, 1, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 3, C128List(0, 1, 2, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 4, C128List(0, 1, 2, 3, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 5, C128List(0, 1, 2, 3, 4))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 6, C128List(0, 1, 2, 3, 4, 5))
}

func TestC128SliceEach(t *testing.T) {
	var count	complex128
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestC128SliceEachWithIndex(t *testing.T) {
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithIndex(func(index int, i interface{}) {
		if index != int(real(i.(complex128))) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestC128SliceEachWithKey(t *testing.T) {
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithKey(func(key, i interface{}) {
		if complex(float64(key.(int)), 0) != i {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestC128SliceC128Each(t *testing.T) {
	var count	complex128
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).C128Each(func(i complex128) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestC128SliceC128EachWithIndex(t *testing.T) {
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).C128EachWithIndex(func(index int, i complex128) {
		if int(real(i)) != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestC128SliceC128EachWithKey(t *testing.T) {
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).C128EachWithKey(func(key interface{}, i complex128) {
		if key.(int) != int(real(i)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestC128SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *C128Slice, destination, source, count int, r *C128Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, C128List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, C128List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestC128SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *C128Slice, start, count int, r *C128Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, C128List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, C128List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestC128SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *C128Slice, offset int, v, r *C128Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, C128List(10, 9, 8, 7), C128List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, C128List(10, 9, 8, 7), C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, C128List(11, 12, 13, 14), C128List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestC128SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *C128Slice, l, c int, r *C128Slice) {
		o := s.String()
		el := l
		if el > c {
			el = c
		}
		switch s.Reallocate(l, c); {
		case s == nil:				t.Fatalf("%v.Reallocate(%v, %v) created a nil value for Slice", o, l, c)
		case s.Cap() != c:			t.Fatalf("%v.Reallocate(%v, %v) capacity should be %v but is %v", o, l, c, c, s.Cap())
		case s.Len() != el:			t.Fatalf("%v.Reallocate(%v, %v) length should be %v but is %v", o, l, c, el, s.Len())
		case !r.Equal(s):			t.Fatalf("%v.Reallocate(%v, %v) should be %v but is %v", o, l, c, r, s)
		}
	}

	ConfirmReallocate(C128List(), 0, 10, C128List())
	ConfirmReallocate(C128List(0, 1, 2, 3, 4), 3, 10, C128List(0, 1, 2))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4), 5, 10, C128List(0, 1, 2, 3, 4))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4), 10, 10, C128List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, C128List(0))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, C128List(0, 1, 2, 3, 4))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, C128List(0, 1, 2, 3, 4))
}

func TestC128SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *C128Slice, n int, r *C128Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(C128List(), 1, C128List(0))
	ConfirmExtend(C128List(), 2, C128List(0, 0))
}

func TestC128SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *C128Slice, i, n int, r *C128Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(C128List(), -1, 1, C128List(0))
	ConfirmExpand(C128List(), 0, 1, C128List(0))
	ConfirmExpand(C128List(), 1, 1, C128List(0))
	ConfirmExpand(C128List(), 0, 2, C128List(0, 0))

	ConfirmExpand(C128List(0, 1, 2), -1, 2, C128List(0, 0, 0, 1, 2))
	ConfirmExpand(C128List(0, 1, 2), 0, 2, C128List(0, 0, 0, 1, 2))
	ConfirmExpand(C128List(0, 1, 2), 1, 2, C128List(0, 0, 0, 1, 2))
	ConfirmExpand(C128List(0, 1, 2), 2, 2, C128List(0, 1, 0, 0, 2))
	ConfirmExpand(C128List(0, 1, 2), 3, 2, C128List(0, 1, 2, 0, 0))
	ConfirmExpand(C128List(0, 1, 2), 4, 2, C128List(0, 1, 2, 0, 0))
}

func TestC128SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *C128Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(C128List(0, 1), 0)
}

func TestC128SliceReverse(t *testing.T) {
	sxp := C128List(1, 2, 3, 4, 5)
	rxp := C128List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestC128SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *C128Slice, v interface{}, r *C128Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(C128List(), complex128(0), C128List(0))
}

func TestC128SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *C128Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(C128List(), C128List(0), C128List(0))
	ConfirmAppendSlice(C128List(), C128List(0, 1), C128List(0, 1))
	ConfirmAppendSlice(C128List(0, 1, 2), C128List(3, 4), C128List(0, 1, 2, 3, 4))
}

func TestC128SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *C128Slice, v interface{}, r *C128Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(C128List(), complex128(0), C128List(0))
	ConfirmPrepend(C128List(0), complex128(1), C128List(1, 0))
}

func TestC128SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *C128Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(C128List(), C128List(0), C128List(0))
	ConfirmPrependSlice(C128List(), C128List(0, 1), C128List(0, 1))
	ConfirmPrependSlice(C128List(0, 1, 2), C128List(3, 4), C128List(3, 4, 0, 1, 2))
}

func TestC128SliceSubslice(t *testing.T) {
	ConfirmSubslice := func(s *C128Slice, start, end int, r *C128Slice) {
		if x := s.Subslice(start, end); !r.Equal(x) {
			t.Fatalf("Subslice(%v, %v) should be %v but is %v", start, end, r, x)
		}
	}
	t.Fatal()
	ConfirmSubslice(C128List(), 0, 1, nil)
}

func TestC128SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *C128Slice, count int, r *C128Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(C128List(), 5, C128List())
	ConfirmRepeat(C128List(0), 1, C128List(0))
	ConfirmRepeat(C128List(0), 2, C128List(0, 0))
	ConfirmRepeat(C128List(0), 3, C128List(0, 0, 0))
	ConfirmRepeat(C128List(0), 4, C128List(0, 0, 0, 0))
	ConfirmRepeat(C128List(0), 5, C128List(0, 0, 0, 0, 0))
}

func TestC128SliceCar(t *testing.T) {
	ConfirmCar := func(s *C128Slice, r complex128) {
		n := s.Car().(complex128)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(C128List(1, 2, 3), 1)
}

func TestC128SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *C128Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(C128List(1, 2, 3), C128List(2, 3))
}

func TestC128SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *C128Slice, v interface{}, r *C128Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(C128List(1, 2, 3, 4, 5), complex128(0), C128List(0, 2, 3, 4, 5))
}

func TestC128SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *C128Slice, v interface{}, r *C128Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(C128List(1, 2, 3, 4, 5), nil, C128List(1))
	ConfirmRplacd(C128List(1, 2, 3, 4, 5), complex128(10), C128List(1, 10))
	ConfirmRplacd(C128List(1, 2, 3, 4, 5), C128List(5, 4, 3, 2), C128List(1, 5, 4, 3, 2))
	ConfirmRplacd(C128List(1, 2, 3, 4, 5), C128List(2, 4, 8, 16, 32), C128List(1, 2, 4, 8, 16, 32))
}