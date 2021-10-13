package adt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnescapePunctuationCharacters(t *testing.T) {
	escapedString := `\!\"\#\$\%\&\'\(\)\*\+\,\-\.\/\:\;\<\=\>\?\@\[\\\]\^\_\` + "`" + `\{\|\}\~`
	unescapedString := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	result := String(escapedString).Unescape()

	assert.Equal(t, string(result), unescapedString)
}

func TestUnescapeOtherCharacters(t *testing.T) {
	escapedString := `\→\A\a\ \3\φ\«`
	unescapedString := escapedString

	result := String(escapedString).Unescape()

	assert.Equal(t, string(result), unescapedString)
}
