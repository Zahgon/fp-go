package readerio

//go:inline
func Bracket[
	R, A, B, ANY any](

	acquire ReaderIO[R, A],
	use Kleisli[R, A, B],
	release func(A, B) ReaderIO[R, ANY],
) ReaderIO[R, B] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func WithResource[R, A, B, ANY any](
	onCreate ReaderIO[R, A], onRelease Kleisli[R, A, ANY]) Kleisli[R, Kleisli[R, A, B], B] {
	_ = "STUB: not implemented"
	return nil
}
