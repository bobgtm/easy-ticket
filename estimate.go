package main

import (
	"time"
)

var eob = "3:30"
var layout = "3:04"
var t, _ = time.Parse(layout, eob)

// Trying to set the time of the end of my day
// Once we accomplish this, we will find the time remaining by taking the end of day and subtracting time.Now
// This will let us know whether we will have to get to someone's request in a given day.
// If we cannot, then it will require they wait longer for their request to be serviced.
// fmt.Println(t.)

// var timeString string
// wait := 20 * countEm(paylaod) / 60
// if wait >= 60 {
// 	timeString = "minutes"
// }

// var x = fmt.Println("Your esitmated wait is:", 20*countEm(payload)/60, "hours")
