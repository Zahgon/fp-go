package readerreaderioeither

import (
	"github.com/IBM/fp-go/v2/either"
	"github.com/IBM/fp-go/v2/io"
	"github.com/IBM/fp-go/v2/ioeither"
	"github.com/IBM/fp-go/v2/reader"
	"github.com/IBM/fp-go/v2/readerioeither"
)

//go:inline
func Local[C, E, A, R1, R2 any](f func(R2) R1) func(ReaderReaderIOEither[R1, C, E, A]) ReaderReaderIOEither[R2, C, E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func LocalIOK[C, E, A, R1, R2 any](f io.Kleisli[R2, R1]) func(ReaderReaderIOEither[R1, C, E, A]) ReaderReaderIOEither[R2, C, E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func LocalIOEitherK[C, A, R1, R2, E any](f ioeither.Kleisli[E, R2, R1]) func(ReaderReaderIOEither[R1, C, E, A]) ReaderReaderIOEither[R2, C, E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func LocalEitherK[C, A, R1, R2, E any](f either.Kleisli[E, R2, R1]) func(ReaderReaderIOEither[R1, C, E, A]) ReaderReaderIOEither[R2, C, E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func LocalReaderIOEitherK[A, C, E, R1, R2 any](f readerioeither.Kleisli[C, E, R2, R1]) func(ReaderReaderIOEither[R1, C, E, A]) ReaderReaderIOEither[R2, C, E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func LocalReaderK[E, A, C, R1, R2 any](f reader.Kleisli[C, R2, R1]) func(ReaderReaderIOEither[R1, C, E, A]) ReaderReaderIOEither[R2, C, E, A] {
	_ = "STUB: not implemented"
	return nil
}

//go:inline
func LocalReaderReaderIOEitherK[A, C, E, R1, R2 any](f Kleisli[R2, C, E, R2, R1]) func(ReaderReaderIOEither[R1, C, E, A]) ReaderReaderIOEither[R2, C, E, A] {
	_ = "STUB: not implemented"
	return nil
}
