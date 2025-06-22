package compiler

import (
	parser "backend/parser"

)

// Visitor para recorrer el árbol de sintaxis
type Visitor struct {
	*parser.BaseLanguageVisitor
};

// NewVisitor crea una nueva instancia del Visitor
func NewVisitor() *Visitor {
	return &Visitor{
		BaseLanguageVisitor: &parser.BaseLanguageVisitor{},
	}
}

