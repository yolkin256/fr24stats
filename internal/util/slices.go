package util

func SliceChunk[T any](slice []T, size int) [][]T {
	chunks := make([][]T, 0, (len(slice)+size-1)/size)
	for size < len(slice) {
		slice, chunks = slice[size:], append(chunks, slice[0:size:size])
	}

	chunks = append(chunks, slice)
	return chunks
}
