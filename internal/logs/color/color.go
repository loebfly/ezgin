package color

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"os"
	"strconv"
	"strings"
)

var (
	NoColor = noColorExists() || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

	// Output defines the standard output of the print functions. By default,
	// os.Stdout is used.
	Output = colorable.NewColorableStdout()
)

// noColorExists returns true if the environment variable NO_COLOR exists.
func noColorExists() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params  []Attribute
	noColor *bool
}

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Reset Base attributes
const (
	Reset Attribute = iota
)

// Foreground text colors
const (
	FgRed     Attribute = 31
	FgGreen   Attribute = 32
	FgYellow  Attribute = 33
	FgMagenta Attribute = 35
)

// New returns a newly created color object.
func New(value ...Attribute) *Color {
	c := &Color{
		params: make([]Attribute, 0),
	}

	if noColorExists() {
		c.noColor = boolPtr(true)
	}

	c.Add(value...)
	return c
}

// Unset resets all escape attributes and clears the output. Usually should
// be called after Set().
func Unset() {
	if NoColor {
		return
	}

	_, _ = fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	if c.isNoColorSet() {
		return c
	}

	_, _ = fmt.Fprint(Output, c.format())
	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// wrap wraps the s string with the colors attributes. The string is ready to
// be printed.
func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unFormat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unFormat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

func (c *Color) isNoColorSet() bool {
	// check first if we have user set action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

func boolPtr(v bool) *bool {
	return &v
}
