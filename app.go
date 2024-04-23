package main

import (
	"fmt"
	"strconv"//imported this for conversion of string to integer
)
//FeatureFlags defines a map from feature names corresponding to boolean values
type FeatureFlags map[string]bool

var featureFlags = FeatureFlags{
	"same day delivery": true, //default value i have set it to enabled/true
}

func main() {

	//getting input from the user(age and order amount)
	var ageStr, orderAmountStr string
	fmt.Print("Enter your age: ")
	fmt.Scanln(&ageStr)
	fmt.Print("Enter your order amount: ")
	fmt.Scanln(&orderAmountStr)

	//converting the age and order input strings to integers
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Invalid age input")
		return
	}
	orderAmount, err := strconv.Atoi(orderAmountStr)
	if err != nil {
		fmt.Println("Invalid order amount input")
		return
	}

	//the conditional expression required inorder to use the feature
	conditionalExpression := "(age > 25) OR (orderAmount > 10000)"

	featureName := "same day delivery"

	//check if the user is allowed to use the same day delivery feature
	allowed := isAllowed(conditionalExpression, featureName, age, orderAmount, featureFlags)

	//printing the result if the user is allowed to use the feature or not
	if allowed {
		fmt.Println("User is allowed to use the feature:", featureName)
	} else {
		fmt.Println("User is not allowed to use the feature:", featureName)
	}
}


func isAllowed(conditionalExpression string, featureName string, age int, orderAmount int, flags FeatureFlags) bool {
	// check if the feature is present in the map and if the boolean value is true.(&& checks both conditions inorder to return true)
	if flagValue, ok := flags[featureName]; ok && flagValue {
		switch conditionalExpression {
		//checking if the conditions are met
		case "(age > 25) OR (orderAmount > 10000)":
			return age > 25 && orderAmount > 10000
		default:
			return false
		}
	}
	//returns false if the feature is disabled or if the conditions are not met
	return false
}
