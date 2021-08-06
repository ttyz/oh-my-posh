package main

type text struct {
	props *properties
	env   environmentInfo
}

const (
	// TextProperty represents text to write
	TextProperty Property = "text"
)

func (t *text) enabled() bool {
	return true
}

func (t *text) string() string {
    template := &textTemplate{
        Template: t.props.getString(TextProperty, "!!text property not defined!!"),
        Context: t,
        Env: t.env,
    }
	text := template.renderPlainContextTemplate(nil)
	return text;
}

func (t *text) init(props *properties, env environmentInfo) {
	t.props = props
	t.env = env
}
