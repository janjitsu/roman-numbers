package graph

import "studiosol.com.br/janjitsu/roman-numbers/app/roman"

// here we pass Finder interface to Resolver
type Resolver struct {
	Search roman.Finder
}
