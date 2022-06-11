package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(12) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "It's Sunday!"
		//Don't need to use break keyword once condition is true (at this point its optional to post it for flow style)
		break
		//fallthrough
	case 2:
		result = "Oh great it's Monday.."
		break
		//fallthrough
	case 3:
		result = "Tuesday.... second day of the week"
		break
	case 4:
		result = "Wed... hump day"
		break
	case 5:
		result = "Thursday, could be Bon Jovi Friday"
		break
	case 6:
		result = "Friday, some say it's Saturday 0.25"
		break
	case 7:
		result = "Finally it's Saturday. Don't let the day end."
		break
	default:
		result = "I don't know what day it is. Are we on Earth?"
	}

	fmt.Println(result)
}
