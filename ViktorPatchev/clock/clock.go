package clock

import "fmt"

func New(hour, minute int) Clock {
	var c Clock
	c.hour = hour
	c.minute = minute
	return c.FixTime()
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.FixTime()

}

func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes
	return c.FixTime()
}

type Clock struct {
	hour   int
	minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) FixTime() Clock {
	for c.minute < 0 {
		c.minute += 60
		c.hour--
	}
	for c.hour < 0 {
		c.hour += 24
	}

	c.hour += int(c.minute / 60)
	c.minute = c.minute % 60
	c.hour = c.hour % 24
	return c
}
