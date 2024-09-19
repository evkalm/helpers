package helpers

import (
	"github.com/fatih/color"
)

var colorValue func(a ...interface{}) string
var colorName func(a ...interface{}) string
var colorType func(a ...interface{}) string
var colorFieldName func(a ...interface{}) string
var colorPipe func(a ...interface{}) string
var colorMapKey func(a ...interface{}) string
var colorErr func(a ...interface{}) string
var colorPerf func(a ...interface{}) string

// init
func init() {
	colorValue = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(70)).SprintFunc()
	colorName = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(220)).SprintFunc()
	colorType = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(60)).SprintFunc()
	colorPipe = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(237)).SprintFunc()
	colorFieldName = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(166)).SprintFunc()
	colorMapKey = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(179)).SprintFunc()
	colorErr = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(1)).SprintFunc()
	colorPerf = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(39)).SprintFunc()
}
