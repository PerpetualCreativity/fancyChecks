// Package fancyChecks provides error-checking functions for command-line tools
package fancyChecks

import (
	"os"

	"github.com/fatih/color"
)

type FancyChecker struct {
	successTemplate string
	neutralTemplate string
	errorTemplate   string
}

// Create a checking instance
func New(allPrefix string, successPrefix string, neutralPrefix string, errorPrefix string) FancyChecker {
	if allPrefix != "" {
		allPrefix += ": "
	}
	successTemplate := allPrefix
	neutralTemplate := allPrefix
	errorTemplate := allPrefix
	if successPrefix != "" {
		successTemplate += successPrefix + ": "
	}
	if neutralPrefix != "" {
		neutralTemplate += neutralPrefix + ": "
	}
	if errorPrefix != "" {
		errorTemplate += errorPrefix + ": "
	}
	successTemplate += "%s"
	neutralTemplate += "%s"
	errorTemplate += "%s: %s"
	fc := FancyChecker{
		successTemplate: successTemplate,
		neutralTemplate: neutralTemplate,
		errorTemplate:   errorTemplate,
	}
	return fc
}

// Prints success message, does not exit
func (fc *FancyChecker) Success(message string) {
	color.Green(fc.successTemplate, message)
}

// Prints neutral message, does not exit
func (fc *FancyChecker) Neutral(message string) {
	color.Blue(fc.neutralTemplate, message)
}

// Prints message and err and then exits
func (fc *FancyChecker) fancyExit(message string, err interface{}) {
	color.Red(fc.errorTemplate, message, err)
	os.Exit(1)
}

// fancyExits if err != comp
func (fc *FancyChecker) ErrComp(err interface{}, comp interface{}, message string) {
	if err != comp {
		fc.fancyExit(message, err)
	}
}

// fancyExits if err == comp
func (fc *FancyChecker) ErrNComp(err interface{}, comp interface{}, message string) {
	if err == comp {
		fc.fancyExit(message, err)
	}
}

// fancyExits if err != nil
func (fc *FancyChecker) ErrCheck(err interface{}, message string) {
	fc.ErrComp(err, nil, message)
}

// fancyExits if err != nil and err != exp
func (fc *FancyChecker) ErrExp(err interface{}, exp interface{}, message string) {
	if err != nil && err != exp {
		fc.fancyExit(message, err)
	}
}
