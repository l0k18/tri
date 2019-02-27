package tri

import (
	"testing"
)

func TestBrief(t *testing.T) {

	// one item only
	tb1 := Brief{
		"item1", "item2",
	}
	e := tb1.Validate()
	if e == nil {
		t.Error("validator allowed more than one")
	}

	// string typed item
	tb2 := Brief{
		1,
	}
	e = tb2.Validate()
	if e == nil {
		t.Error("validator permitted other than a string")
	}

	// string length < 80
	tb3 := Brief{
		"123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
	}
	e = tb3.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	// no control characters
	tb4 := Brief{
		"this should not have a cr at the end\n",
	}
	e = tb4.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	tb5 := Brief{
		"this should not have a cr at the end\n",
	}
	e = tb5.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	tb6 := Brief{
		"this is ok",
	}
	e = tb6.Validate()
	if e != nil {
		t.Error("validator rejected correct input")
	}

}

func TestCommand(t *testing.T) {

	//string in index 0

	//string is valid name (letters only)

	//more than one brief not allowed

	//more than one handler not allowed

	//handler not nil

	//Brief field present

	//Handler present

}

func TestCommands(t *testing.T) {

}

func TestDefault(t *testing.T) {

	// only one item
	td1 := Default{
		"item1", "item2",
	}
	e := td1.Validate()
	if e == nil {
		t.Error("validator allowed more than one")
	}

	// item is string
	td2 := Default{
		1,
	}
	e = td2.Validate()
	if e == nil {
		t.Error("validator permitted other than a string")
	}

	// item is a ValidName
	td3 := Default{
		"abc123",
	}
	e = td3.Validate()
	if e == nil {
		t.Error("validator permitted an invalid name")
	}

	// item is a valid
	td4 := Default{
		"abc",
	}
	e = td4.Validate()
	if e != nil {
		t.Error("validator rejected a valid name")
	}

}

func TestDefaultCommand(t *testing.T) {

	// only one item
	tdc1 := DefaultCommand{
		"item1", "item2",
	}
	e := tdc1.Validate()
	if e == nil {
		t.Error("validator allowed more than one")
	}

	// item is string
	tdc2 := DefaultCommand{
		1,
	}
	e = tdc2.Validate()
	if e == nil {
		t.Error("validator permitted other than a string")
	}

	// item is a ValidName
	tdc3 := DefaultCommand{
		"abc123",
	}
	e = tdc3.Validate()
	if e == nil {
		t.Error("validator permitted an invalid name")
	}

	// item is a valid
	tdc4 := DefaultCommand{
		"abc",
	}
	e = tdc4.Validate()
	if e != nil {
		t.Error("validator rejected a valid name")
	}
}

func TestDefaultOn(t *testing.T) {

	// must be empty

	// must not be empty

	// must have pairs of elements

	// elements must be strings

	// even numbered (first in pair) elements have no control characters

}

func TestExamples(t *testing.T) {

	// must not be empty

	// must have pairs of elements

	// elements must be strings

	// even numbered (first in pair) elements have no control characters

	// contains only one element

	// element is a string

	// string is a ValidName

}

func TestGroup(t *testing.T) {

	// contains only one element

	// element is a string

	// string is a ValidName
}

func TestHelp(t *testing.T) {

	// contains only one element

	// element is a string

}

func TestRunAfter(t *testing.T) {

	// may not contain anything

}

func TestShort(t *testing.T) {

	// contains only one elemennt

	// element is a rune (single character/unicode point)

}

func TestSlot(t *testing.T) {

	// slots are all the same type (pointer to said type)

}

func TestTerminates(t *testing.T) {

	// contains no elements
}

func TestTri(t *testing.T) {

	// contains at least 4 elements

	// first element is a string

	// string is a ValidName

	// contains (only) one Brief

	// contains (only) one Version

	// contains (only) one Commands

	// contains no more than one DefaultCommand

	// only contains element from set of possible elements

	// Brief is missing

	// Version is missing

	// Commands is missing

}

func TestTrigger(t *testing.T) {

	// must contain name, Brief and Handler

	// first element is name

	// name is a ValidName

	// only one Brief

	// only one Handler

	// rest of items only (one of) Short, Usage, Help, DefaultOn, Terminates and RunAfter

}

func TestUsage(t *testing.T) {

	// only one element

	// element is string

	// string is no more than 80 chars long

	// string contains no control characters

}

func TestValidName(t *testing.T) {

	// contains at least 3 elements

	// first is string

	// name is ValidName

	// only one Brief

	// only one handler

	// has one each of Brief and Handler

}

func TestVar(t *testing.T) {

	// has no more than 4 fields

	// has at least 3 fields

	// first three are integers

	// integers are under 100

	// field 4 is a string

	// string contains only letters and numbers

}

func TestVersion(t *testing.T) {

	//contains only letters

}
