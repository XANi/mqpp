package util

import (
	"github.com/XANi/mqpp/common"
	"github.com/fatih/color"
	"fmt"
	"strings"
)

type FormatConfig struct {
	MetaFormat string
	HeaderFormat string
	BodyFormat string

}

var Formatting FormatConfig

func init() {
	Formatting = FormatConfig{
		MetaFormat: "%-60s",
		HeaderFormat: "\n   ^--%-54s",
		BodyFormat: "%s",

	}
}




func Format(m common.Message) string {
	 var formattedSource, outHeaders, out string
		if len(m.Source) > 0 {
		formattedSource = color.HiGreenString(strings.Join(m.Source,`/`))
	}
	var formattedHeaders string
	if len(m.Headers) > 0 {
		var h []string
		for k, v := range m.Headers {
			h = append(h, fmt.Sprintf("%s→%v",k,v))
		}
		formattedHeaders = color.HiBlueString(strings.Join(h," "))
	}
	if len(formattedSource) > 0 {
		if len(formattedHeaders) > 0 {
			outHeaders = fmt.Sprintf("%s" + Formatting.HeaderFormat, formattedSource, formattedHeaders)
		} else {
			outHeaders = fmt.Sprintf("%s", formattedSource)
		}
	} else {
		outHeaders = color.HiRedString("^---")
	}
	out = fmt.Sprintf(Formatting.MetaFormat + ": " + Formatting.BodyFormat, outHeaders, string(m.Body))
	return out
}