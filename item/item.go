package item

// PUBLIC INTERFACE
//

// STARTIFACE OMIT
type Item interface {
	Type() string
	Time() int
	unique()
}
// STOPIFACE OMIT

// HIDDEN STRUCT IMPLEMENTING INTERFACE
//

// STARTSTRUCT OMIT
type item struct {
	taype string
	time  int
}
// STOPSTRUCT OMIT

// METHODS
//

func (it *item) unique() {}

func (it *item) Type() string {
	return it.taype
}

func (it *item) Time() int {
	return it.time
}

// PACKAGE FUNCTIONS
//

// STARTNEW OMIT
func New(taype string, time int) Item {
	return &item{taype: taype, time: time}
}
// STOPNEW OMIT
