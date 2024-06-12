package shorten_path

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortenPath(t *testing.T) {
	t.Run("", func(t *testing.T) {
		expected := "/foo/bar/baz"
		output := ShortenPath("/foo/../test/../test/../foo//bar/./baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/foo/bar/baz"
		output := ShortenPath("/../../foo/bar/baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/foo/bar/baz"
		output := ShortenPath("/foo/bar/baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "foo/bar/baz"
		output := ShortenPath("foo/bar/baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/foo/bar/baz"
		output := ShortenPath("/../../foo/bar/baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "../../foo/bar/baz"
		output := ShortenPath("../../foo/bar/baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/bar/baz"
		output := ShortenPath("/../../foo/../../bar/baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "../../../bar/baz"
		output := ShortenPath("../../foo/../../bar/baz")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/foo/kappa"
		output := ShortenPath("/foo/./././bar/./baz///////////test/../../../kappa")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "../../../../../../../.."
		output := ShortenPath("../../../this////one/./../../is/../../going/../../to/be/./././../../../just/eight/double/dots/../../../../../..")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/"
		output := ShortenPath("/../../../this////one/./../../is/../../going/../../to/be/./././../../../just/a/forward/slash/../../../../../..")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "../../../../../../../../foo"
		output := ShortenPath("../../../this////one/./../../is/../../going/../../to/be/./././../../../just/eight/double/dots/../../../../../../foo")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/foo"
		output := ShortenPath("/../../../this////one/./../../is/../../going/../../to/be/./././../../../just/a/forward/slash/../../../../../../foo")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "foo"
		output := ShortenPath("foo/bar/..")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "foo/bar"
		output := ShortenPath("./foo/bar")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := ".."
		output := ShortenPath("foo/../..")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "/"
		output := ShortenPath("/")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := ".."
		output := ShortenPath("./..")
		require.Equal(t, expected, output)
	})
}
