package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v72"
)

func main() {
	params := &stripe.CustomerParams{
		Description: stripe.String("Stripe Developer"),
		Email:       stripe.String("gostripe@stripe.com"),
	}
	fmt.Println(&params)
	fmt.Println("Hello World!")
}
