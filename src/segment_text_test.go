package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextSegmentTemplate(t *testing.T) {
	// set date for unit test
	cases := []struct {
		Case            string
		ExpectedString  string
		Text            string
	}{
		{
			Case:            "no template",
			Text:            "",
			ExpectedString:  "",
		},
		{
			Case:            "static value",
			Text:            "some text",
			ExpectedString:  "some text",
		},
        {
			Case:            "template value",
			Text:            "{{.Env.VALUE}}",
			ExpectedString:  "some value",
		},
        {
            Case:            "  value",
            Text:            "{{if .Root}}root{{else}}no root{{end}}",
            ExpectedString:  "root",
        },
	}

	for _, tc := range cases {
		env := new(MockedEnvironment)
        env.On("getenv", "VALUE").Return("some value")
        env.On("isRunningAsRoot", nil).Return(true)
        env.On("getcwd", nil).Return("/some/path")
        env.On("homeDir", nil).Return("/home/user")
        env.On("getPathSeperator", nil).Return("/")
        env.On("getShellName", nil).Return("testsh")
        env.On("getCurrentUser", nil).Return("testuser")
        env.On("getHostName", nil).Return("test.com", nil)
		props := &properties{
			values: map[Property]interface{}{
				TextProperty: tc.Text,
			},
		}
		text := &text{
			env:         env,
			props:       props,
		}

        assert.Equal(t, tc.ExpectedString, text.string(), tc.Case)
	}
}
