package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "studiosol.com.br/janjitsu/roman-numbers/app/roman"

type Resolver struct {
	Search roman.Finder
}
