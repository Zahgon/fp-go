package reader

//go:inline
func Bracket[
	R, A, B, ANY any](

	acquire Reader[R, A],
	use Kleisli[R, A, B],
	release func(A, B) Reader[R, ANY],
) Reader[R, B] {
	_ = "STUB: not implemented"
	return nil
}
