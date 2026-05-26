package readerresult

// Bracket makes sure that a resource is cleaned up in the event of an error. The release action is called regardless of
// whether the body action returns and error or not.
func Bracket[
	R, A, B, ANY any](

	acquire Lazy[ReaderResult[R, A]],
	use Kleisli[R, A, B],
	release func(A, B, error) ReaderResult[R, ANY],
) ReaderResult[R, B] {
	_ = "STUB: not implemented"
	return nil
}

// acquire the resource

func WithResource[B, R, A, ANY any](
	onCreate Lazy[ReaderResult[R, A]],
	onRelease Kleisli[R, A, ANY],
) Kleisli[R, Kleisli[R, A, B], B] {
	_ = "STUB: not implemented"
	return nil
}
