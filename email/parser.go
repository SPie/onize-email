package email

import (
    "bytes"
    "fmt"
    "text/template"
)

type ParserContract interface {
    Parse(templateName string, data map[string]interface{}) ([]byte, error)
}

type Parser struct {
    templatesDir string
}

func NewParser(templatesDir string) ParserContract {
    return Parser{templatesDir}
}

func (parser Parser) Parse(templateName string, data map[string]interface{}) ([]byte, error) {
    parseTemplate, err := template.ParseFiles(fmt.Sprintf("%s/%s.txt", parser.templatesDir, templateName))
    if err != nil {
	return nil, err
    }

    buffer := new(bytes.Buffer)
    err = parseTemplate.Execute(buffer, data)
    if err != nil {
	return nil, err
    }

    return buffer.Bytes(), nil
}
