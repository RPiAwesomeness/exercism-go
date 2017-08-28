package clock

import (
	"fmt"
	"strconv"
)

const testVersion = 4

// Clock is an object that keeps track of hours and minutes in 24 hour form
type Clock struct {
	Hour   int
	Minute int
}

// New creates a new Clock object and returns it from the given time in hours & minutes
func New(hour, minute int) Clock {
	var createHour, createMin int
	// TODO: Add logic for handling roll-over, continuous roll-over, and negatives
	// Midnight is {0, 0}
	// Negative hours are from {0, 0}
	// 60 minutes = 1 hour (minutes roll over into hours)

	// Parse the minute value
	if minute >= 60 {
		// It's greater than 60, roll over into X hours
		createHour += minute / 60
		// Subtract that many hours worth of minutes
		createMin = minute - (createHour * 60)
	} else if minute <= 0 {
		// It's negative щ(゜ロ゜щ)
		// += because it'll be a negative number from the division
		createHour += minute / 60
		// Subtract the number of hours in minutes from the minutes
		createMin = minute - (createHour * 60)
	} else {
		// It's a positive minute less than an hour
		createMin = minute
	}

	// Parse the hour values
	if hour >= 24 {
		// It's more than 1 day, roll over
		// How many days?
		days := hour / 24
		// Remove that many days worth of hours
		createHour -= hour - (days * 24)
	} else if hour <= 0 {
		// It's a negative hour! щ(゜ロ゜щ) Roll over backwards
		days := hour / 24
		// Remove that many days worth of hours
		createHour -= hour + (days * 24)
	} else {
		// It's a positive hour less than a day
		createHour += hour
	}

	// Set up the Clock object and return it
	return Clock{
		Hour:   createHour,
		Minute: createMin,
	}
}

// String is a method for objects of type Clock that returns the current
// time in string form, formatted in HH:MM format
func (c Clock) String() string {
	var hour string
	var minute string

	// These if statements convert the integer values into HH and MM format
	if c.Hour < 10 {
		hour = fmt.Sprintf("0%d", c.Hour)
	} else {
		hour = strconv.Itoa(c.Hour)
	}
	if c.Minute < 10 {
		minute = fmt.Sprintf("0%d", c.Minute)
	} else {
		minute = strconv.Itoa(c.Minute)
	}
	// Return the HH:MM formatted string
	return fmt.Sprintf("%s:%s", hour, minute)
}

// Add is a method for objects of type Clock that allows you to add minutes
// to the clock's time in int form
func (c Clock) Add(minutes int) Clock {
	// TODO: Must increment hours
	// Subtract minutes and then remove hours (across days)
	return c
}
