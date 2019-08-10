package email

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateNewParser(t *testing.T) {
    templatesDir := "/templatesDir"

    parser := NewParser(templatesDir)

    assert.Implements(t, new(ParserContract), parser)
}

func TestParse(t *testing.T) {
    text, err := NewParser("templates").Parse("test", map[string]interface{}{"testVariable":"A Test Variable"})

    assert.Empty(t, err)
    assert.Equal(t, "Test variable: A Test Variable\n", string(text))
}

func TestParseWithInvalidTemplateName(t *testing.T) {
    _, err := NewParser("templates").Parse("InvalidTemplateName", map[string]interface{}{"testVariable":"A Test Variable"})

    assert.NotEmpty(t, err)
}

func TestParseWithError(t *testing.T) {
    _, err := NewParser("templates").Parse("invalidTest", map[string]interface{}{})

    assert.NotEmpty(t, err)
}
