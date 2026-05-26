package readert

func Sequence[
	HKTR2HKTR1A ~func(R2) HKTR1HKTA,
	R1, R2, HKTR1HKTA, HKTA any](
	mchain func(func(func(R1) HKTA) HKTA) func(HKTR1HKTA) HKTA,
	ma HKTR2HKTR1A,
) func(R1) func(R2) HKTA {
	_ = "STUB: not implemented"
	return nil
}

func SequenceReader[
	HKTR2HKTR1A ~func(R2) HKTR1HKTA,
	R1, R2, A, HKTR1HKTA, HKTA any](
	mmap func(func(func(R1) A) A) func(HKTR1HKTA) HKTA,
	ma HKTR2HKTR1A,
) func(R1) func(R2) HKTA {
	_ = "STUB: not implemented"
	return nil
}

func Traverse[
	HKTR2A ~func(R2) HKTA,
	HKTR1B ~func(R1) HKTB,
	R1, R2, A, HKTR1HKTB, HKTA, HKTB any](
	mmap func(func(A) HKTR1B) func(HKTA) HKTR1HKTB,
	mchain func(func(func(R1) HKTB) HKTB) func(HKTR1HKTB) HKTB,
	f func(A) HKTR1B,
) func(HKTR2A) func(R1) func(R2) HKTB {
	_ = "STUB: not implemented"
	return nil
}

func TraverseReader[
	HKTR2A ~func(R2) HKTA,
	HKTR1B ~func(R1) B,
	R1, R2, A, B, HKTR1HKTB, HKTA, HKTB any](
	mmap1 func(func(A) HKTR1B) func(HKTA) HKTR1HKTB,
	mmap2 func(func(func(R1) B) B) func(HKTR1HKTB) HKTB,
	f func(A) HKTR1B,
) func(HKTR2A) func(R1) func(R2) HKTB {
	_ = "STUB: not implemented"
	return nil
}
