package readerioresult

import (
	"github.com/IBM/fp-go/v2/reader"
)

//go:inline
func Sequence[R1, R2, A any](ma ReaderIOResult[R2, ReaderIOResult[R1, A]]) Kleisli[R2, R1, A] {
	_ = "STUB: not implemented"
	return nil

	//go:inline
}

func SequenceReader[R1, R2, A any](ma ReaderIOResult[R2, Reader[R1, A]]) Kleisli[R2, R1, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func SequenceReaderIO[R1, R2, A any](ma ReaderIOResult[R2, ReaderIO[R1, A]]) Kleisli[R2, R1, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func SequenceReaderEither[R1, R2, A any](ma ReaderIOResult[R2, ReaderResult[R1, A]]) Kleisli[R2, R1, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func SequenceReaderResult[R1, R2, A any](ma ReaderIOResult[R2, ReaderResult[R1, A]]) Kleisli[R2, R1, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func Traverse[R2, R1, A, B any](
	f Kleisli[R1, A, B],
) func(ReaderIOResult[R2, A]) Kleisli[R2, R1, B] {
	_ = "STUB: not implemented"
	return nil
}

func TraverseReader[R2, R1, A, B any](
	f reader.Kleisli[R1, A, B],
) func(ReaderIOResult[R2, A]) Kleisli[R2, R1, B] {
	_ = "STUB: not implemented"
	return nil
}
