// Code generated from parser/Language.g4 by ANTLR 4.13.2. DO NOT EDIT.

package interpreter // Language
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LanguageParser struct {
	*antlr.BaseParser
}

var LanguageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func languageParserInit() {
	staticData := &LanguageParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'-'", "'!'", "'*'", "'/'", "'%'", "'+'", "'>'", "'<'",
		"'>='", "'<='", "'=='", "'!='", "'&&'", "'||'", "'.'", "'('", "')'",
		"'['", "','", "']'", "'='", "'+='", "'-='", "'++'", "'--'", "'nil'",
		"'if'", "'else'", "'switch'", "'case'", "':'", "'default'", "'for'",
		"';'", "'range'", "'break'", "'continue'", "'return'", "'atoi'", "'parseFloat'",
		"'typeOf'", "'print'", "'mut'", "'[]'", "'[][]'", "':='", "'struct'",
		"'fn'", "'void'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "TYPE", "INTEGER", "BOOLEAN", "FLOAT", "STRING", "ID", "ONELINECOMMENT",
		"MULTILINECOMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "nonDeclaration", "blockStatement", "expressionStatement",
		"ifStatement", "switchStatement", "switchCase", "defaultCase", "forStatement",
		"transferenceStatement", "atoiStatement", "parseFloatStatement", "typeOfStatement",
		"printStatement", "variableDeclaration", "sliceElements", "matrixElements",
		"structInitialization", "structFieldInit", "structDeclaration", "structAtribute",
		"functionDeclaration", "parameterList", "parameter", "argumentList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 60, 469, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 63,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		75, 8, 2, 1, 3, 1, 3, 5, 3, 79, 8, 3, 10, 3, 12, 3, 82, 9, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 92, 8, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 5, 4, 99, 8, 4, 10, 4, 12, 4, 102, 9, 4, 3, 4, 104, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 148, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 165, 8, 4, 3, 4, 167, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 198,
		8, 4, 10, 4, 12, 4, 201, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		209, 8, 5, 3, 5, 211, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 217, 8, 6, 10,
		6, 12, 6, 220, 9, 6, 1, 6, 3, 6, 223, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 231, 8, 7, 10, 7, 12, 7, 234, 9, 7, 1, 8, 1, 8, 1, 8, 5,
		8, 239, 8, 8, 10, 8, 12, 8, 242, 9, 8, 1, 9, 1, 9, 1, 9, 3, 9, 247, 8,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 267, 8, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 273, 8, 10, 3, 10, 275, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 297, 8, 14, 10, 14, 12,
		14, 300, 9, 14, 3, 14, 302, 8, 14, 1, 14, 1, 14, 1, 15, 3, 15, 307, 8,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 313, 8, 15, 1, 15, 3, 15, 316, 8,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 324, 8, 15, 3, 15,
		326, 8, 15, 1, 15, 3, 15, 329, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		3, 15, 336, 8, 15, 1, 15, 3, 15, 339, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 3, 15, 346, 8, 15, 1, 15, 3, 15, 349, 8, 15, 1, 15, 3, 15, 352,
		8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 359, 8, 15, 1, 15, 1,
		15, 3, 15, 363, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		371, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 379, 8, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 384, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5,
		16, 390, 8, 16, 10, 16, 12, 16, 393, 9, 16, 3, 16, 395, 8, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 403, 8, 17, 10, 17, 12, 17, 406,
		9, 17, 3, 17, 408, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 5, 18, 415,
		8, 18, 10, 18, 12, 18, 418, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 5, 20, 428, 8, 20, 10, 20, 12, 20, 431, 9, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 442, 8,
		22, 1, 22, 1, 22, 3, 22, 446, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		5, 23, 453, 8, 23, 10, 23, 12, 23, 456, 9, 23, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 5, 25, 464, 8, 25, 10, 25, 12, 25, 467, 9, 25, 1, 25,
		0, 1, 8, 26, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 0, 9, 1, 0, 3, 4, 1, 0, 24, 25,
		1, 0, 26, 27, 1, 0, 5, 7, 2, 0, 3, 3, 8, 8, 1, 0, 9, 12, 1, 0, 13, 14,
		2, 0, 52, 52, 57, 57, 2, 0, 51, 52, 57, 57, 535, 0, 55, 1, 0, 0, 0, 2,
		62, 1, 0, 0, 0, 4, 74, 1, 0, 0, 0, 6, 76, 1, 0, 0, 0, 8, 166, 1, 0, 0,
		0, 10, 202, 1, 0, 0, 0, 12, 212, 1, 0, 0, 0, 14, 226, 1, 0, 0, 0, 16, 235,
		1, 0, 0, 0, 18, 266, 1, 0, 0, 0, 20, 274, 1, 0, 0, 0, 22, 276, 1, 0, 0,
		0, 24, 281, 1, 0, 0, 0, 26, 286, 1, 0, 0, 0, 28, 291, 1, 0, 0, 0, 30, 383,
		1, 0, 0, 0, 32, 385, 1, 0, 0, 0, 34, 398, 1, 0, 0, 0, 36, 411, 1, 0, 0,
		0, 38, 419, 1, 0, 0, 0, 40, 423, 1, 0, 0, 0, 42, 434, 1, 0, 0, 0, 44, 437,
		1, 0, 0, 0, 46, 449, 1, 0, 0, 0, 48, 457, 1, 0, 0, 0, 50, 460, 1, 0, 0,
		0, 52, 54, 3, 2, 1, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53,
		1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 1, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0,
		58, 63, 3, 4, 2, 0, 59, 63, 3, 30, 15, 0, 60, 63, 3, 40, 20, 0, 61, 63,
		3, 44, 22, 0, 62, 58, 1, 0, 0, 0, 62, 59, 1, 0, 0, 0, 62, 60, 1, 0, 0,
		0, 62, 61, 1, 0, 0, 0, 63, 3, 1, 0, 0, 0, 64, 75, 3, 6, 3, 0, 65, 75, 3,
		28, 14, 0, 66, 75, 3, 22, 11, 0, 67, 75, 3, 24, 12, 0, 68, 75, 3, 26, 13,
		0, 69, 75, 3, 10, 5, 0, 70, 75, 3, 12, 6, 0, 71, 75, 3, 18, 9, 0, 72, 75,
		3, 20, 10, 0, 73, 75, 3, 8, 4, 0, 74, 64, 1, 0, 0, 0, 74, 65, 1, 0, 0,
		0, 74, 66, 1, 0, 0, 0, 74, 67, 1, 0, 0, 0, 74, 68, 1, 0, 0, 0, 74, 69,
		1, 0, 0, 0, 74, 70, 1, 0, 0, 0, 74, 71, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0,
		74, 73, 1, 0, 0, 0, 75, 5, 1, 0, 0, 0, 76, 80, 5, 1, 0, 0, 77, 79, 3, 2,
		1, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81,
		1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 84, 5, 2, 0, 0,
		84, 7, 1, 0, 0, 0, 85, 86, 6, 4, -1, 0, 86, 87, 7, 0, 0, 0, 87, 167, 3,
		8, 4, 29, 88, 89, 5, 57, 0, 0, 89, 91, 5, 18, 0, 0, 90, 92, 3, 50, 25,
		0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 167,
		5, 19, 0, 0, 94, 103, 5, 20, 0, 0, 95, 100, 3, 8, 4, 0, 96, 97, 5, 21,
		0, 0, 97, 99, 3, 8, 4, 0, 98, 96, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100,
		98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1,
		0, 0, 0, 103, 95, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0,
		0, 105, 167, 5, 22, 0, 0, 106, 107, 5, 57, 0, 0, 107, 108, 5, 20, 0, 0,
		108, 109, 3, 8, 4, 0, 109, 110, 5, 22, 0, 0, 110, 111, 5, 20, 0, 0, 111,
		112, 3, 8, 4, 0, 112, 113, 5, 22, 0, 0, 113, 114, 5, 23, 0, 0, 114, 115,
		3, 8, 4, 19, 115, 167, 1, 0, 0, 0, 116, 117, 5, 57, 0, 0, 117, 118, 5,
		20, 0, 0, 118, 119, 3, 8, 4, 0, 119, 120, 5, 22, 0, 0, 120, 121, 5, 23,
		0, 0, 121, 122, 3, 8, 4, 18, 122, 167, 1, 0, 0, 0, 123, 124, 5, 57, 0,
		0, 124, 125, 7, 1, 0, 0, 125, 167, 3, 8, 4, 15, 126, 127, 5, 57, 0, 0,
		127, 167, 7, 2, 0, 0, 128, 167, 3, 22, 11, 0, 129, 167, 3, 24, 12, 0, 130,
		167, 3, 26, 13, 0, 131, 132, 5, 57, 0, 0, 132, 133, 5, 20, 0, 0, 133, 134,
		3, 8, 4, 0, 134, 135, 5, 22, 0, 0, 135, 136, 5, 20, 0, 0, 136, 137, 3,
		8, 4, 0, 137, 138, 5, 22, 0, 0, 138, 167, 1, 0, 0, 0, 139, 140, 5, 57,
		0, 0, 140, 141, 5, 20, 0, 0, 141, 142, 3, 8, 4, 0, 142, 143, 5, 22, 0,
		0, 143, 167, 1, 0, 0, 0, 144, 145, 5, 57, 0, 0, 145, 147, 5, 1, 0, 0, 146,
		148, 3, 36, 18, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 167, 5, 2, 0, 0, 150, 167, 5, 53, 0, 0, 151, 167, 5, 54,
		0, 0, 152, 167, 5, 55, 0, 0, 153, 167, 5, 56, 0, 0, 154, 167, 5, 57, 0,
		0, 155, 167, 5, 28, 0, 0, 156, 157, 5, 18, 0, 0, 157, 158, 3, 8, 4, 0,
		158, 159, 5, 19, 0, 0, 159, 165, 1, 0, 0, 0, 160, 161, 5, 20, 0, 0, 161,
		162, 3, 8, 4, 0, 162, 163, 5, 22, 0, 0, 163, 165, 1, 0, 0, 0, 164, 156,
		1, 0, 0, 0, 164, 160, 1, 0, 0, 0, 165, 167, 1, 0, 0, 0, 166, 85, 1, 0,
		0, 0, 166, 88, 1, 0, 0, 0, 166, 94, 1, 0, 0, 0, 166, 106, 1, 0, 0, 0, 166,
		116, 1, 0, 0, 0, 166, 123, 1, 0, 0, 0, 166, 126, 1, 0, 0, 0, 166, 128,
		1, 0, 0, 0, 166, 129, 1, 0, 0, 0, 166, 130, 1, 0, 0, 0, 166, 131, 1, 0,
		0, 0, 166, 139, 1, 0, 0, 0, 166, 144, 1, 0, 0, 0, 166, 150, 1, 0, 0, 0,
		166, 151, 1, 0, 0, 0, 166, 152, 1, 0, 0, 0, 166, 153, 1, 0, 0, 0, 166,
		154, 1, 0, 0, 0, 166, 155, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 199,
		1, 0, 0, 0, 168, 169, 10, 28, 0, 0, 169, 170, 7, 3, 0, 0, 170, 198, 3,
		8, 4, 29, 171, 172, 10, 27, 0, 0, 172, 173, 7, 4, 0, 0, 173, 198, 3, 8,
		4, 28, 174, 175, 10, 26, 0, 0, 175, 176, 7, 5, 0, 0, 176, 198, 3, 8, 4,
		27, 177, 178, 10, 25, 0, 0, 178, 179, 7, 6, 0, 0, 179, 198, 3, 8, 4, 26,
		180, 181, 10, 24, 0, 0, 181, 182, 5, 15, 0, 0, 182, 198, 3, 8, 4, 25, 183,
		184, 10, 23, 0, 0, 184, 185, 5, 16, 0, 0, 185, 198, 3, 8, 4, 24, 186, 187,
		10, 17, 0, 0, 187, 188, 5, 17, 0, 0, 188, 189, 5, 57, 0, 0, 189, 190, 5,
		23, 0, 0, 190, 198, 3, 8, 4, 18, 191, 192, 10, 16, 0, 0, 192, 193, 5, 23,
		0, 0, 193, 198, 3, 8, 4, 17, 194, 195, 10, 22, 0, 0, 195, 196, 5, 17, 0,
		0, 196, 198, 5, 57, 0, 0, 197, 168, 1, 0, 0, 0, 197, 171, 1, 0, 0, 0, 197,
		174, 1, 0, 0, 0, 197, 177, 1, 0, 0, 0, 197, 180, 1, 0, 0, 0, 197, 183,
		1, 0, 0, 0, 197, 186, 1, 0, 0, 0, 197, 191, 1, 0, 0, 0, 197, 194, 1, 0,
		0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0,
		200, 9, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203, 5, 29, 0, 0, 203, 204,
		3, 8, 4, 0, 204, 210, 3, 6, 3, 0, 205, 208, 5, 30, 0, 0, 206, 209, 3, 10,
		5, 0, 207, 209, 3, 6, 3, 0, 208, 206, 1, 0, 0, 0, 208, 207, 1, 0, 0, 0,
		209, 211, 1, 0, 0, 0, 210, 205, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211,
		11, 1, 0, 0, 0, 212, 213, 5, 31, 0, 0, 213, 214, 5, 57, 0, 0, 214, 218,
		5, 1, 0, 0, 215, 217, 3, 14, 7, 0, 216, 215, 1, 0, 0, 0, 217, 220, 1, 0,
		0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0,
		220, 218, 1, 0, 0, 0, 221, 223, 3, 16, 8, 0, 222, 221, 1, 0, 0, 0, 222,
		223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 5, 2, 0, 0, 225, 13, 1,
		0, 0, 0, 226, 227, 5, 32, 0, 0, 227, 228, 3, 8, 4, 0, 228, 232, 5, 33,
		0, 0, 229, 231, 3, 2, 1, 0, 230, 229, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0,
		232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 15, 1, 0, 0, 0, 234, 232,
		1, 0, 0, 0, 235, 236, 5, 34, 0, 0, 236, 240, 5, 33, 0, 0, 237, 239, 3,
		2, 1, 0, 238, 237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0,
		0, 240, 241, 1, 0, 0, 0, 241, 17, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243,
		246, 5, 35, 0, 0, 244, 247, 3, 30, 15, 0, 245, 247, 3, 8, 4, 0, 246, 244,
		1, 0, 0, 0, 246, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 5, 36,
		0, 0, 249, 250, 3, 8, 4, 0, 250, 251, 5, 36, 0, 0, 251, 252, 3, 8, 4, 0,
		252, 253, 3, 6, 3, 0, 253, 267, 1, 0, 0, 0, 254, 255, 5, 35, 0, 0, 255,
		256, 3, 8, 4, 0, 256, 257, 3, 6, 3, 0, 257, 267, 1, 0, 0, 0, 258, 259,
		5, 35, 0, 0, 259, 260, 5, 57, 0, 0, 260, 261, 5, 21, 0, 0, 261, 262, 5,
		57, 0, 0, 262, 263, 5, 23, 0, 0, 263, 264, 5, 37, 0, 0, 264, 265, 5, 57,
		0, 0, 265, 267, 3, 6, 3, 0, 266, 243, 1, 0, 0, 0, 266, 254, 1, 0, 0, 0,
		266, 258, 1, 0, 0, 0, 267, 19, 1, 0, 0, 0, 268, 275, 5, 38, 0, 0, 269,
		275, 5, 39, 0, 0, 270, 272, 5, 40, 0, 0, 271, 273, 3, 8, 4, 0, 272, 271,
		1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 268, 1, 0,
		0, 0, 274, 269, 1, 0, 0, 0, 274, 270, 1, 0, 0, 0, 275, 21, 1, 0, 0, 0,
		276, 277, 5, 41, 0, 0, 277, 278, 5, 18, 0, 0, 278, 279, 3, 8, 4, 0, 279,
		280, 5, 19, 0, 0, 280, 23, 1, 0, 0, 0, 281, 282, 5, 42, 0, 0, 282, 283,
		5, 18, 0, 0, 283, 284, 3, 8, 4, 0, 284, 285, 5, 19, 0, 0, 285, 25, 1, 0,
		0, 0, 286, 287, 5, 43, 0, 0, 287, 288, 5, 18, 0, 0, 288, 289, 3, 8, 4,
		0, 289, 290, 5, 19, 0, 0, 290, 27, 1, 0, 0, 0, 291, 292, 5, 44, 0, 0, 292,
		301, 5, 18, 0, 0, 293, 298, 3, 8, 4, 0, 294, 295, 5, 21, 0, 0, 295, 297,
		3, 8, 4, 0, 296, 294, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0,
		0, 0, 298, 299, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0,
		301, 293, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303,
		304, 5, 19, 0, 0, 304, 29, 1, 0, 0, 0, 305, 307, 5, 45, 0, 0, 306, 305,
		1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 5, 57,
		0, 0, 309, 312, 5, 52, 0, 0, 310, 311, 5, 23, 0, 0, 311, 313, 3, 8, 4,
		0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 384, 1, 0, 0, 0, 314,
		316, 5, 45, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 317,
		1, 0, 0, 0, 317, 318, 5, 57, 0, 0, 318, 319, 5, 46, 0, 0, 319, 325, 5,
		52, 0, 0, 320, 323, 5, 23, 0, 0, 321, 324, 3, 32, 16, 0, 322, 324, 3, 8,
		4, 0, 323, 321, 1, 0, 0, 0, 323, 322, 1, 0, 0, 0, 324, 326, 1, 0, 0, 0,
		325, 320, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 384, 1, 0, 0, 0, 327,
		329, 5, 45, 0, 0, 328, 327, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 330,
		1, 0, 0, 0, 330, 331, 5, 57, 0, 0, 331, 332, 5, 47, 0, 0, 332, 335, 5,
		52, 0, 0, 333, 334, 5, 23, 0, 0, 334, 336, 3, 34, 17, 0, 335, 333, 1, 0,
		0, 0, 335, 336, 1, 0, 0, 0, 336, 384, 1, 0, 0, 0, 337, 339, 5, 45, 0, 0,
		338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340,
		341, 5, 57, 0, 0, 341, 348, 5, 57, 0, 0, 342, 343, 5, 23, 0, 0, 343, 345,
		5, 1, 0, 0, 344, 346, 3, 36, 18, 0, 345, 344, 1, 0, 0, 0, 345, 346, 1,
		0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 349, 5, 2, 0, 0, 348, 342, 1, 0, 0,
		0, 348, 349, 1, 0, 0, 0, 349, 384, 1, 0, 0, 0, 350, 352, 5, 45, 0, 0, 351,
		350, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 354,
		5, 57, 0, 0, 354, 355, 5, 48, 0, 0, 355, 356, 5, 57, 0, 0, 356, 358, 5,
		1, 0, 0, 357, 359, 3, 36, 18, 0, 358, 357, 1, 0, 0, 0, 358, 359, 1, 0,
		0, 0, 359, 360, 1, 0, 0, 0, 360, 384, 5, 2, 0, 0, 361, 363, 5, 45, 0, 0,
		362, 361, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364,
		365, 5, 57, 0, 0, 365, 366, 5, 48, 0, 0, 366, 367, 5, 46, 0, 0, 367, 368,
		5, 52, 0, 0, 368, 384, 3, 32, 16, 0, 369, 371, 5, 45, 0, 0, 370, 369, 1,
		0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 373, 5, 57, 0,
		0, 373, 374, 5, 48, 0, 0, 374, 375, 5, 47, 0, 0, 375, 376, 5, 52, 0, 0,
		376, 384, 3, 34, 17, 0, 377, 379, 5, 45, 0, 0, 378, 377, 1, 0, 0, 0, 378,
		379, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 5, 57, 0, 0, 381, 382,
		5, 48, 0, 0, 382, 384, 3, 8, 4, 0, 383, 306, 1, 0, 0, 0, 383, 315, 1, 0,
		0, 0, 383, 328, 1, 0, 0, 0, 383, 338, 1, 0, 0, 0, 383, 351, 1, 0, 0, 0,
		383, 362, 1, 0, 0, 0, 383, 370, 1, 0, 0, 0, 383, 378, 1, 0, 0, 0, 384,
		31, 1, 0, 0, 0, 385, 394, 5, 1, 0, 0, 386, 391, 3, 8, 4, 0, 387, 388, 5,
		21, 0, 0, 388, 390, 3, 8, 4, 0, 389, 387, 1, 0, 0, 0, 390, 393, 1, 0, 0,
		0, 391, 389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393,
		391, 1, 0, 0, 0, 394, 386, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 396,
		1, 0, 0, 0, 396, 397, 5, 2, 0, 0, 397, 33, 1, 0, 0, 0, 398, 407, 5, 1,
		0, 0, 399, 404, 3, 32, 16, 0, 400, 401, 5, 21, 0, 0, 401, 403, 3, 32, 16,
		0, 402, 400, 1, 0, 0, 0, 403, 406, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 404,
		405, 1, 0, 0, 0, 405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 407, 399,
		1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 5, 2,
		0, 0, 410, 35, 1, 0, 0, 0, 411, 416, 3, 38, 19, 0, 412, 413, 5, 21, 0,
		0, 413, 415, 3, 38, 19, 0, 414, 412, 1, 0, 0, 0, 415, 418, 1, 0, 0, 0,
		416, 414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 37, 1, 0, 0, 0, 418, 416,
		1, 0, 0, 0, 419, 420, 5, 57, 0, 0, 420, 421, 5, 33, 0, 0, 421, 422, 3,
		8, 4, 0, 422, 39, 1, 0, 0, 0, 423, 424, 5, 49, 0, 0, 424, 425, 5, 57, 0,
		0, 425, 429, 5, 1, 0, 0, 426, 428, 3, 42, 21, 0, 427, 426, 1, 0, 0, 0,
		428, 431, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430,
		432, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 432, 433, 5, 2, 0, 0, 433, 41, 1,
		0, 0, 0, 434, 435, 7, 7, 0, 0, 435, 436, 5, 57, 0, 0, 436, 43, 1, 0, 0,
		0, 437, 438, 5, 50, 0, 0, 438, 439, 5, 57, 0, 0, 439, 441, 5, 18, 0, 0,
		440, 442, 3, 46, 23, 0, 441, 440, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442,
		443, 1, 0, 0, 0, 443, 445, 5, 19, 0, 0, 444, 446, 7, 8, 0, 0, 445, 444,
		1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 448, 3, 6,
		3, 0, 448, 45, 1, 0, 0, 0, 449, 454, 3, 48, 24, 0, 450, 451, 5, 21, 0,
		0, 451, 453, 3, 48, 24, 0, 452, 450, 1, 0, 0, 0, 453, 456, 1, 0, 0, 0,
		454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 47, 1, 0, 0, 0, 456, 454,
		1, 0, 0, 0, 457, 458, 5, 57, 0, 0, 458, 459, 7, 7, 0, 0, 459, 49, 1, 0,
		0, 0, 460, 465, 3, 8, 4, 0, 461, 462, 5, 21, 0, 0, 462, 464, 3, 8, 4, 0,
		463, 461, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 465,
		466, 1, 0, 0, 0, 466, 51, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 50, 55, 62,
		74, 80, 91, 100, 103, 147, 164, 166, 197, 199, 208, 210, 218, 222, 232,
		240, 246, 266, 272, 274, 298, 301, 306, 312, 315, 323, 325, 328, 335, 338,
		345, 348, 351, 358, 362, 370, 378, 383, 391, 394, 404, 407, 416, 429, 441,
		445, 454, 465,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LanguageParserInit initializes any static state used to implement LanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LanguageParserInit() {
	staticData := &LanguageParserStaticData
	staticData.once.Do(languageParserInit)
}

// NewLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLanguageParser(input antlr.TokenStream) *LanguageParser {
	LanguageParserInit()
	this := new(LanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LanguageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Language.g4"

	return this
}

// LanguageParser tokens.
const (
	LanguageParserEOF              = antlr.TokenEOF
	LanguageParserT__0             = 1
	LanguageParserT__1             = 2
	LanguageParserT__2             = 3
	LanguageParserT__3             = 4
	LanguageParserT__4             = 5
	LanguageParserT__5             = 6
	LanguageParserT__6             = 7
	LanguageParserT__7             = 8
	LanguageParserT__8             = 9
	LanguageParserT__9             = 10
	LanguageParserT__10            = 11
	LanguageParserT__11            = 12
	LanguageParserT__12            = 13
	LanguageParserT__13            = 14
	LanguageParserT__14            = 15
	LanguageParserT__15            = 16
	LanguageParserT__16            = 17
	LanguageParserT__17            = 18
	LanguageParserT__18            = 19
	LanguageParserT__19            = 20
	LanguageParserT__20            = 21
	LanguageParserT__21            = 22
	LanguageParserT__22            = 23
	LanguageParserT__23            = 24
	LanguageParserT__24            = 25
	LanguageParserT__25            = 26
	LanguageParserT__26            = 27
	LanguageParserT__27            = 28
	LanguageParserT__28            = 29
	LanguageParserT__29            = 30
	LanguageParserT__30            = 31
	LanguageParserT__31            = 32
	LanguageParserT__32            = 33
	LanguageParserT__33            = 34
	LanguageParserT__34            = 35
	LanguageParserT__35            = 36
	LanguageParserT__36            = 37
	LanguageParserT__37            = 38
	LanguageParserT__38            = 39
	LanguageParserT__39            = 40
	LanguageParserT__40            = 41
	LanguageParserT__41            = 42
	LanguageParserT__42            = 43
	LanguageParserT__43            = 44
	LanguageParserT__44            = 45
	LanguageParserT__45            = 46
	LanguageParserT__46            = 47
	LanguageParserT__47            = 48
	LanguageParserT__48            = 49
	LanguageParserT__49            = 50
	LanguageParserT__50            = 51
	LanguageParserTYPE             = 52
	LanguageParserINTEGER          = 53
	LanguageParserBOOLEAN          = 54
	LanguageParserFLOAT            = 55
	LanguageParserSTRING           = 56
	LanguageParserID               = 57
	LanguageParserONELINECOMMENT   = 58
	LanguageParserMULTILINECOMMENT = 59
	LanguageParserWS               = 60
)

// LanguageParser rules.
const (
	LanguageParserRULE_program               = 0
	LanguageParserRULE_statement             = 1
	LanguageParserRULE_nonDeclaration        = 2
	LanguageParserRULE_blockStatement        = 3
	LanguageParserRULE_expressionStatement   = 4
	LanguageParserRULE_ifStatement           = 5
	LanguageParserRULE_switchStatement       = 6
	LanguageParserRULE_switchCase            = 7
	LanguageParserRULE_defaultCase           = 8
	LanguageParserRULE_forStatement          = 9
	LanguageParserRULE_transferenceStatement = 10
	LanguageParserRULE_atoiStatement         = 11
	LanguageParserRULE_parseFloatStatement   = 12
	LanguageParserRULE_typeOfStatement       = 13
	LanguageParserRULE_printStatement        = 14
	LanguageParserRULE_variableDeclaration   = 15
	LanguageParserRULE_sliceElements         = 16
	LanguageParserRULE_matrixElements        = 17
	LanguageParserRULE_structInitialization  = 18
	LanguageParserRULE_structFieldInit       = 19
	LanguageParserRULE_structDeclaration     = 20
	LanguageParserRULE_structAtribute        = 21
	LanguageParserRULE_functionDeclaration   = 22
	LanguageParserRULE_parameterList         = 23
	LanguageParserRULE_parameter             = 24
	LanguageParserRULE_argumentList          = 25
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LanguageParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280982157937344538) != 0 {
		{
			p.SetState(52)
			p.Statement()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NonDeclaration() INonDeclarationContext
	VariableDeclaration() IVariableDeclarationContext
	StructDeclaration() IStructDeclarationContext
	FunctionDeclaration() IFunctionDeclarationContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) NonDeclaration() INonDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonDeclarationContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) StructDeclaration() IStructDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LanguageParserRULE_statement)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.NonDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.VariableDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.StructDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.FunctionDeclaration()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonDeclarationContext is an interface to support dynamic dispatch.
type INonDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BlockStatement() IBlockStatementContext
	PrintStatement() IPrintStatementContext
	AtoiStatement() IAtoiStatementContext
	ParseFloatStatement() IParseFloatStatementContext
	TypeOfStatement() ITypeOfStatementContext
	IfStatement() IIfStatementContext
	SwitchStatement() ISwitchStatementContext
	ForStatement() IForStatementContext
	TransferenceStatement() ITransferenceStatementContext
	ExpressionStatement() IExpressionStatementContext

	// IsNonDeclarationContext differentiates from other interfaces.
	IsNonDeclarationContext()
}

type NonDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonDeclarationContext() *NonDeclarationContext {
	var p = new(NonDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_nonDeclaration
	return p
}

func InitEmptyNonDeclarationContext(p *NonDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_nonDeclaration
}

func (*NonDeclarationContext) IsNonDeclarationContext() {}

func NewNonDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonDeclarationContext {
	var p = new(NonDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_nonDeclaration

	return p
}

func (s *NonDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NonDeclarationContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *NonDeclarationContext) PrintStatement() IPrintStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *NonDeclarationContext) AtoiStatement() IAtoiStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtoiStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtoiStatementContext)
}

func (s *NonDeclarationContext) ParseFloatStatement() IParseFloatStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParseFloatStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParseFloatStatementContext)
}

func (s *NonDeclarationContext) TypeOfStatement() ITypeOfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeOfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeOfStatementContext)
}

func (s *NonDeclarationContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *NonDeclarationContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *NonDeclarationContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *NonDeclarationContext) TransferenceStatement() ITransferenceStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransferenceStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransferenceStatementContext)
}

func (s *NonDeclarationContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *NonDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterNonDeclaration(s)
	}
}

func (s *NonDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitNonDeclaration(s)
	}
}

func (s *NonDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitNonDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) NonDeclaration() (localctx INonDeclarationContext) {
	localctx = NewNonDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LanguageParserRULE_nonDeclaration)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.BlockStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.PrintStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.AtoiStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.ParseFloatStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.TypeOfStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.IfStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.SwitchStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(71)
			p.ForStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(72)
			p.TransferenceStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(73)
			p.expressionStatement(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_blockStatement
	return p
}

func InitEmptyBlockStatementContext(p *BlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_blockStatement
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LanguageParserRULE_blockStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(LanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280982157937344538) != 0 {
		{
			p.SetState(77)
			p.Statement()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(83)
		p.Match(LanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_expressionStatement
	return p
}

func InitEmptyExpressionStatementContext(p *ExpressionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_expressionStatement
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) CopyAll(ctx *ExpressionStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceAssignmentContext struct {
	ExpressionStatementContext
}

func NewSliceAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAssignmentContext {
	var p = new(SliceAssignmentContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *SliceAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *SliceAssignmentContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *SliceAssignmentContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *SliceAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterSliceAssignment(s)
	}
}

func (s *SliceAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitSliceAssignment(s)
	}
}

func (s *SliceAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitSliceAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	ExpressionStatementContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceLiteralContext struct {
	ExpressionStatementContext
}

func NewSliceLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceLiteralContext {
	var p = new(SliceLiteralContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *SliceLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceLiteralContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *SliceLiteralContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *SliceLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterSliceLiteral(s)
	}
}

func (s *SliceLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitSliceLiteral(s)
	}
}

func (s *SliceLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitSliceLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParseFloatExprContext struct {
	ExpressionStatementContext
}

func NewParseFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParseFloatExprContext {
	var p = new(ParseFloatExprContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *ParseFloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseFloatExprContext) ParseFloatStatement() IParseFloatStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParseFloatStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParseFloatStatementContext)
}

func (s *ParseFloatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterParseFloatExpr(s)
	}
}

func (s *ParseFloatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitParseFloatExpr(s)
	}
}

func (s *ParseFloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitParseFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExpressionStatementContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ExpressionStatementContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(LanguageParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructAssignmentContext struct {
	ExpressionStatementContext
}

func NewStructAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAssignmentContext {
	var p = new(StructAssignmentContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *StructAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAssignmentContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *StructAssignmentContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StructAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *StructAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStructAssignment(s)
	}
}

func (s *StructAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStructAssignment(s)
	}
}

func (s *StructAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStructAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrixAccessContext struct {
	ExpressionStatementContext
}

func NewMatrixAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixAccessContext {
	var p = new(MatrixAccessContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *MatrixAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *MatrixAccessContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *MatrixAccessContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *MatrixAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterMatrixAccess(s)
	}
}

func (s *MatrixAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitMatrixAccess(s)
	}
}

func (s *MatrixAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitMatrixAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentContext struct {
	ExpressionStatementContext
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructAccessContext struct {
	ExpressionStatementContext
}

func NewStructAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAccessContext {
	var p = new(StructAccessContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *StructAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAccessContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StructAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *StructAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStructAccess(s)
	}
}

func (s *StructAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStructAccess(s)
	}
}

func (s *StructAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStructAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubOperatorContext struct {
	ExpressionStatementContext
}

func NewAddSubOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubOperatorContext {
	var p = new(AddSubOperatorContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *AddSubOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubOperatorContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *AddSubOperatorContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *AddSubOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAddSubOperator(s)
	}
}

func (s *AddSubOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAddSubOperator(s)
	}
}

func (s *AddSubOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAddSubOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModContext struct {
	ExpressionStatementContext
}

func NewMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModContext {
	var p = new(MulDivModContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *MulDivModContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

func (s *MulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtoiExprContext struct {
	ExpressionStatementContext
}

func NewAtoiExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtoiExprContext {
	var p = new(AtoiExprContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *AtoiExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtoiExprContext) AtoiStatement() IAtoiStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtoiStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtoiStatementContext)
}

func (s *AtoiExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAtoiExpr(s)
	}
}

func (s *AtoiExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAtoiExpr(s)
	}
}

func (s *AtoiExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAtoiExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierContext struct {
	ExpressionStatementContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualContext struct {
	ExpressionStatementContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *EqualContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitEqual(s)
	}
}

func (s *EqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterLessContext struct {
	ExpressionStatementContext
}

func NewGreaterLessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterLessContext {
	var p = new(GreaterLessContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *GreaterLessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterLessContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *GreaterLessContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *GreaterLessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterGreaterLess(s)
	}
}

func (s *GreaterLessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitGreaterLess(s)
	}
}

func (s *GreaterLessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitGreaterLess(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExpressionStatementContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ExpressionStatementContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(LanguageParserBOOLEAN, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceAccessContext struct {
	ExpressionStatementContext
}

func NewSliceAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAccessContext {
	var p = new(SliceAccessContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *SliceAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *SliceAccessContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *SliceAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterSliceAccess(s)
	}
}

func (s *SliceAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitSliceAccess(s)
	}
}

func (s *SliceAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitSliceAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExpressionStatementContext
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrixAssignmentContext struct {
	ExpressionStatementContext
}

func NewMatrixAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *MatrixAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *MatrixAssignmentContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *MatrixAssignmentContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *MatrixAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterMatrixAssignment(s)
	}
}

func (s *MatrixAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitMatrixAssignment(s)
	}
}

func (s *MatrixAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitMatrixAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructInstantiationContext struct {
	ExpressionStatementContext
}

func NewStructInstantiationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInstantiationContext {
	var p = new(StructInstantiationContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *StructInstantiationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstantiationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *StructInstantiationContext) StructInitialization() IStructInitializationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInitializationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInitializationContext)
}

func (s *StructInstantiationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStructInstantiation(s)
	}
}

func (s *StructInstantiationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStructInstantiation(s)
	}
}

func (s *StructInstantiationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStructInstantiation(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerContext struct {
	ExpressionStatementContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LanguageParserINTEGER, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilContext struct {
	ExpressionStatementContext
}

func NewNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilContext {
	var p = new(NilContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterNil(s)
	}
}

func (s *NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitNil(s)
	}
}

func (s *NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeOfExprContext struct {
	ExpressionStatementContext
}

func NewTypeOfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOfExprContext {
	var p = new(TypeOfExprContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *TypeOfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOfExprContext) TypeOfStatement() ITypeOfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeOfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeOfStatementContext)
}

func (s *TypeOfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterTypeOfExpr(s)
	}
}

func (s *TypeOfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitTypeOfExpr(s)
	}
}

func (s *TypeOfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitTypeOfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	ExpressionStatementContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(LanguageParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	ExpressionStatementContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegateContext struct {
	ExpressionStatementContext
}

func NewNegateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateContext {
	var p = new(NegateContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *NegateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *NegateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterNegate(s)
	}
}

func (s *NegateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitNegate(s)
	}
}

func (s *NegateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitNegate(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrementDecrementContext struct {
	ExpressionStatementContext
}

func NewIncrementDecrementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementDecrementContext {
	var p = new(IncrementDecrementContext)

	InitEmptyExpressionStatementContext(&p.ExpressionStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionStatementContext))

	return p
}

func (s *IncrementDecrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementDecrementContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *IncrementDecrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterIncrementDecrement(s)
	}
}

func (s *IncrementDecrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitIncrementDecrement(s)
	}
}

func (s *IncrementDecrementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitIncrementDecrement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	return p.expressionStatement(0)
}

func (p *LanguageParser) expressionStatement(_p int) (localctx IExpressionStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, LanguageParserRULE_expressionStatement, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(86)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LanguageParserT__2 || _la == LanguageParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(87)
			p.expressionStatement(29)
		}

	case 2:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Match(LanguageParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279238570329505816) != 0 {
			{
				p.SetState(90)
				p.ArgumentList()
			}

		}
		{
			p.SetState(93)
			p.Match(LanguageParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSliceLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.Match(LanguageParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279238570329505816) != 0 {
			{
				p.SetState(95)
				p.expressionStatement(0)
			}
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == LanguageParserT__20 {
				{
					p.SetState(96)
					p.Match(LanguageParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)
					p.expressionStatement(0)
				}

				p.SetState(102)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(105)
			p.Match(LanguageParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewMatrixAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(106)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(LanguageParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.expressionStatement(0)
		}
		{
			p.SetState(109)
			p.Match(LanguageParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(LanguageParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.expressionStatement(0)
		}
		{
			p.SetState(112)
			p.Match(LanguageParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(LanguageParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.expressionStatement(19)
		}

	case 5:
		localctx = NewSliceAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(116)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(LanguageParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.expressionStatement(0)
		}
		{
			p.SetState(119)
			p.Match(LanguageParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(LanguageParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.expressionStatement(18)
		}

	case 6:
		localctx = NewAddSubOperatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LanguageParserT__23 || _la == LanguageParserT__24) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(125)
			p.expressionStatement(15)
		}

	case 7:
		localctx = NewIncrementDecrementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(126)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LanguageParserT__25 || _la == LanguageParserT__26) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 8:
		localctx = NewAtoiExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(128)
			p.AtoiStatement()
		}

	case 9:
		localctx = NewParseFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(129)
			p.ParseFloatStatement()
		}

	case 10:
		localctx = NewTypeOfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(130)
			p.TypeOfStatement()
		}

	case 11:
		localctx = NewMatrixAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(LanguageParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expressionStatement(0)
		}
		{
			p.SetState(134)
			p.Match(LanguageParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(LanguageParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.expressionStatement(0)
		}
		{
			p.SetState(137)
			p.Match(LanguageParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewSliceAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(139)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(LanguageParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.expressionStatement(0)
		}
		{
			p.SetState(142)
			p.Match(LanguageParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewStructInstantiationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(LanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserID {
			{
				p.SetState(146)
				p.StructInitialization()
			}

		}
		{
			p.SetState(149)
			p.Match(LanguageParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(150)
			p.Match(LanguageParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(151)
			p.Match(LanguageParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(152)
			p.Match(LanguageParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			p.Match(LanguageParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(154)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 19:
		localctx = NewNilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(155)
			p.Match(LanguageParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 20:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LanguageParserT__17:
			{
				p.SetState(156)
				p.Match(LanguageParserT__17)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(157)
				p.expressionStatement(0)
			}
			{
				p.SetState(158)
				p.Match(LanguageParserT__18)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case LanguageParserT__19:
			{
				p.SetState(160)
				p.Match(LanguageParserT__19)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(161)
				p.expressionStatement(0)
			}
			{
				p.SetState(162)
				p.Match(LanguageParserT__21)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(197)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(169)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&224) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(170)
					p.expressionStatement(29)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				{
					p.SetState(172)
					_la = p.GetTokenStream().LA(1)

					if !(_la == LanguageParserT__2 || _la == LanguageParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(173)
					p.expressionStatement(28)
				}

			case 3:
				localctx = NewGreaterLessContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				{
					p.SetState(175)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7680) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(176)
					p.expressionStatement(27)
				}

			case 4:
				localctx = NewEqualContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				{
					p.SetState(178)
					_la = p.GetTokenStream().LA(1)

					if !(_la == LanguageParserT__12 || _la == LanguageParserT__13) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(179)
					p.expressionStatement(26)
				}

			case 5:
				localctx = NewAndContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(181)
					p.Match(LanguageParserT__14)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(182)
					p.expressionStatement(25)
				}

			case 6:
				localctx = NewOrContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(184)
					p.Match(LanguageParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(185)
					p.expressionStatement(24)
				}

			case 7:
				localctx = NewStructAssignmentContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(187)
					p.Match(LanguageParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(188)
					p.Match(LanguageParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(189)
					p.Match(LanguageParserT__22)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(190)
					p.expressionStatement(18)
				}

			case 8:
				localctx = NewAssignmentContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(192)
					p.Match(LanguageParserT__22)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(193)
					p.expressionStatement(17)
				}

			case 9:
				localctx = NewStructAccessContext(p, NewExpressionStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_expressionStatement)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(195)
					p.Match(LanguageParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(196)
					p.Match(LanguageParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionStatement() IExpressionStatementContext
	AllBlockStatement() []IBlockStatementContext
	BlockStatement(i int) IBlockStatementContext
	IfStatement() IIfStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *IfStatementContext) AllBlockStatement() []IBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStatementContext); ok {
			tst[i] = t.(IBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) BlockStatement(i int) IBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LanguageParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(LanguageParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.expressionStatement(0)
	}
	{
		p.SetState(204)
		p.BlockStatement()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LanguageParserT__29 {
		{
			p.SetState(205)
			p.Match(LanguageParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LanguageParserT__28:
			{
				p.SetState(206)
				p.IfStatement()
			}

		case LanguageParserT__0:
			{
				p.SetState(207)
				p.BlockStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	DefaultCase() IDefaultCaseContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *SwitchStatementContext) AllSwitchCase() []ISwitchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseContext); ok {
			tst[i] = t.(ISwitchCaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) SwitchCase(i int) ISwitchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCaseContext)
}

func (s *SwitchStatementContext) DefaultCase() IDefaultCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultCaseContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LanguageParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(LanguageParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(LanguageParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(LanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserT__31 {
		{
			p.SetState(215)
			p.SwitchCase()
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LanguageParserT__33 {
		{
			p.SetState(221)
			p.DefaultCase()
		}

	}
	{
		p.SetState(224)
		p.Match(LanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionStatement() IExpressionStatementContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *SwitchCaseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LanguageParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(LanguageParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.expressionStatement(0)
	}
	{
		p.SetState(228)
		p.Match(LanguageParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280982157937344538) != 0 {
		{
			p.SetState(229)
			p.Statement()
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultCaseContext is an interface to support dynamic dispatch.
type IDefaultCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsDefaultCaseContext differentiates from other interfaces.
	IsDefaultCaseContext()
}

type DefaultCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultCaseContext() *DefaultCaseContext {
	var p = new(DefaultCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_defaultCase
	return p
}

func InitEmptyDefaultCaseContext(p *DefaultCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_defaultCase
}

func (*DefaultCaseContext) IsDefaultCaseContext() {}

func NewDefaultCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_defaultCase

	return p
}

func (s *DefaultCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultCaseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *DefaultCaseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) DefaultCase() (localctx IDefaultCaseContext) {
	localctx = NewDefaultCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LanguageParserRULE_defaultCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(LanguageParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Match(LanguageParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280982157937344538) != 0 {
		{
			p.SetState(237)
			p.Statement()
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) CopyAll(ctx *ForStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForConditionalContext struct {
	ForStatementContext
}

func NewForConditionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForConditionalContext {
	var p = new(ForConditionalContext)

	InitEmptyForStatementContext(&p.ForStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStatementContext))

	return p
}

func (s *ForConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForConditionalContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ForConditionalContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ForConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterForConditional(s)
	}
}

func (s *ForConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitForConditional(s)
	}
}

func (s *ForConditionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitForConditional(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForSimpleContext struct {
	ForStatementContext
}

func NewForSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForSimpleContext {
	var p = new(ForSimpleContext)

	InitEmptyForStatementContext(&p.ForStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStatementContext))

	return p
}

func (s *ForSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForSimpleContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForSimpleContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ForSimpleContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ForSimpleContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterForSimple(s)
	}
}

func (s *ForSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitForSimple(s)
	}
}

func (s *ForSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitForSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForSliceContext struct {
	ForStatementContext
}

func NewForSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForSliceContext {
	var p = new(ForSliceContext)

	InitEmptyForStatementContext(&p.ForStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStatementContext))

	return p
}

func (s *ForSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForSliceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserID)
}

func (s *ForSliceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserID, i)
}

func (s *ForSliceContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ForSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterForSlice(s)
	}
}

func (s *ForSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitForSlice(s)
	}
}

func (s *ForSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitForSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LanguageParserRULE_forStatement)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.Match(LanguageParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(244)
				p.VariableDeclaration()
			}

		case 2:
			{
				p.SetState(245)
				p.expressionStatement(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(248)
			p.Match(LanguageParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.expressionStatement(0)
		}
		{
			p.SetState(250)
			p.Match(LanguageParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.expressionStatement(0)
		}
		{
			p.SetState(252)
			p.BlockStatement()
		}

	case 2:
		localctx = NewForConditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(LanguageParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.expressionStatement(0)
		}
		{
			p.SetState(256)
			p.BlockStatement()
		}

	case 3:
		localctx = NewForSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(258)
			p.Match(LanguageParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(LanguageParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Match(LanguageParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Match(LanguageParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.BlockStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransferenceStatementContext is an interface to support dynamic dispatch.
type ITransferenceStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransferenceStatementContext differentiates from other interfaces.
	IsTransferenceStatementContext()
}

type TransferenceStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransferenceStatementContext() *TransferenceStatementContext {
	var p = new(TransferenceStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_transferenceStatement
	return p
}

func InitEmptyTransferenceStatementContext(p *TransferenceStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_transferenceStatement
}

func (*TransferenceStatementContext) IsTransferenceStatementContext() {}

func NewTransferenceStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferenceStatementContext {
	var p = new(TransferenceStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_transferenceStatement

	return p
}

func (s *TransferenceStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferenceStatementContext) CopyAll(ctx *TransferenceStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TransferenceStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferenceStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakStatementContext struct {
	TransferenceStatementContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	InitEmptyTransferenceStatementContext(&p.TransferenceStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferenceStatementContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	TransferenceStatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyTransferenceStatementContext(&p.TransferenceStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferenceStatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	TransferenceStatementContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	InitEmptyTransferenceStatementContext(&p.TransferenceStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferenceStatementContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) TransferenceStatement() (localctx ITransferenceStatementContext) {
	localctx = NewTransferenceStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LanguageParserRULE_transferenceStatement)
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LanguageParserT__37:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(LanguageParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LanguageParserT__38:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
			p.Match(LanguageParserT__38)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LanguageParserT__39:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.Match(LanguageParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(271)
				p.expressionStatement(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtoiStatementContext is an interface to support dynamic dispatch.
type IAtoiStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionStatement() IExpressionStatementContext

	// IsAtoiStatementContext differentiates from other interfaces.
	IsAtoiStatementContext()
}

type AtoiStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtoiStatementContext() *AtoiStatementContext {
	var p = new(AtoiStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_atoiStatement
	return p
}

func InitEmptyAtoiStatementContext(p *AtoiStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_atoiStatement
}

func (*AtoiStatementContext) IsAtoiStatementContext() {}

func NewAtoiStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtoiStatementContext {
	var p = new(AtoiStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_atoiStatement

	return p
}

func (s *AtoiStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AtoiStatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *AtoiStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtoiStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtoiStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAtoiStatement(s)
	}
}

func (s *AtoiStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAtoiStatement(s)
	}
}

func (s *AtoiStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAtoiStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) AtoiStatement() (localctx IAtoiStatementContext) {
	localctx = NewAtoiStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LanguageParserRULE_atoiStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(LanguageParserT__40)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(LanguageParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.expressionStatement(0)
	}
	{
		p.SetState(279)
		p.Match(LanguageParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParseFloatStatementContext is an interface to support dynamic dispatch.
type IParseFloatStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionStatement() IExpressionStatementContext

	// IsParseFloatStatementContext differentiates from other interfaces.
	IsParseFloatStatementContext()
}

type ParseFloatStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseFloatStatementContext() *ParseFloatStatementContext {
	var p = new(ParseFloatStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_parseFloatStatement
	return p
}

func InitEmptyParseFloatStatementContext(p *ParseFloatStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_parseFloatStatement
}

func (*ParseFloatStatementContext) IsParseFloatStatementContext() {}

func NewParseFloatStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseFloatStatementContext {
	var p = new(ParseFloatStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_parseFloatStatement

	return p
}

func (s *ParseFloatStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseFloatStatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ParseFloatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseFloatStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseFloatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterParseFloatStatement(s)
	}
}

func (s *ParseFloatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitParseFloatStatement(s)
	}
}

func (s *ParseFloatStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitParseFloatStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) ParseFloatStatement() (localctx IParseFloatStatementContext) {
	localctx = NewParseFloatStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LanguageParserRULE_parseFloatStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(LanguageParserT__41)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(LanguageParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.expressionStatement(0)
	}
	{
		p.SetState(284)
		p.Match(LanguageParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeOfStatementContext is an interface to support dynamic dispatch.
type ITypeOfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionStatement() IExpressionStatementContext

	// IsTypeOfStatementContext differentiates from other interfaces.
	IsTypeOfStatementContext()
}

type TypeOfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeOfStatementContext() *TypeOfStatementContext {
	var p = new(TypeOfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_typeOfStatement
	return p
}

func InitEmptyTypeOfStatementContext(p *TypeOfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_typeOfStatement
}

func (*TypeOfStatementContext) IsTypeOfStatementContext() {}

func NewTypeOfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeOfStatementContext {
	var p = new(TypeOfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_typeOfStatement

	return p
}

func (s *TypeOfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeOfStatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *TypeOfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeOfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterTypeOfStatement(s)
	}
}

func (s *TypeOfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitTypeOfStatement(s)
	}
}

func (s *TypeOfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitTypeOfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) TypeOfStatement() (localctx ITypeOfStatementContext) {
	localctx = NewTypeOfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LanguageParserRULE_typeOfStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(LanguageParserT__42)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(LanguageParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.expressionStatement(0)
	}
	{
		p.SetState(289)
		p.Match(LanguageParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionStatement() []IExpressionStatementContext
	ExpressionStatement(i int) IExpressionStatementContext

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_printStatement
	return p
}

func InitEmptyPrintStatementContext(p *PrintStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_printStatement
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *PrintStatementContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LanguageParserRULE_printStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(LanguageParserT__43)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(LanguageParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279238570329505816) != 0 {
		{
			p.SetState(293)
			p.expressionStatement(0)
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LanguageParserT__20 {
			{
				p.SetState(294)
				p.Match(LanguageParserT__20)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(295)
				p.expressionStatement(0)
			}

			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(303)
		p.Match(LanguageParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) CopyAll(ctx *VariableDeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExplicitSliceDeclarationContext struct {
	VariableDeclarationContext
}

func NewExplicitSliceDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExplicitSliceDeclarationContext {
	var p = new(ExplicitSliceDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ExplicitSliceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplicitSliceDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *ExplicitSliceDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *ExplicitSliceDeclarationContext) SliceElements() ISliceElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceElementsContext)
}

func (s *ExplicitSliceDeclarationContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ExplicitSliceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterExplicitSliceDeclaration(s)
	}
}

func (s *ExplicitSliceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitExplicitSliceDeclaration(s)
	}
}

func (s *ExplicitSliceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitExplicitSliceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImplicitStructDeclarationContext struct {
	VariableDeclarationContext
}

func NewImplicitStructDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImplicitStructDeclarationContext {
	var p = new(ImplicitStructDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ImplicitStructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplicitStructDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserID)
}

func (s *ImplicitStructDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserID, i)
}

func (s *ImplicitStructDeclarationContext) StructInitialization() IStructInitializationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInitializationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInitializationContext)
}

func (s *ImplicitStructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterImplicitStructDeclaration(s)
	}
}

func (s *ImplicitStructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitImplicitStructDeclaration(s)
	}
}

func (s *ImplicitStructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitImplicitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExplicitStructDeclarationContext struct {
	VariableDeclarationContext
}

func NewExplicitStructDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExplicitStructDeclarationContext {
	var p = new(ExplicitStructDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ExplicitStructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplicitStructDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserID)
}

func (s *ExplicitStructDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserID, i)
}

func (s *ExplicitStructDeclarationContext) StructInitialization() IStructInitializationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInitializationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInitializationContext)
}

func (s *ExplicitStructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterExplicitStructDeclaration(s)
	}
}

func (s *ExplicitStructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitExplicitStructDeclaration(s)
	}
}

func (s *ExplicitStructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitExplicitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExplicitMatrixDeclarationContext struct {
	VariableDeclarationContext
}

func NewExplicitMatrixDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExplicitMatrixDeclarationContext {
	var p = new(ExplicitMatrixDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ExplicitMatrixDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplicitMatrixDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *ExplicitMatrixDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *ExplicitMatrixDeclarationContext) MatrixElements() IMatrixElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixElementsContext)
}

func (s *ExplicitMatrixDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterExplicitMatrixDeclaration(s)
	}
}

func (s *ExplicitMatrixDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitExplicitMatrixDeclaration(s)
	}
}

func (s *ExplicitMatrixDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitExplicitMatrixDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImplicitMatrixDeclarationContext struct {
	VariableDeclarationContext
}

func NewImplicitMatrixDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImplicitMatrixDeclarationContext {
	var p = new(ImplicitMatrixDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ImplicitMatrixDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplicitMatrixDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *ImplicitMatrixDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *ImplicitMatrixDeclarationContext) MatrixElements() IMatrixElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixElementsContext)
}

func (s *ImplicitMatrixDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterImplicitMatrixDeclaration(s)
	}
}

func (s *ImplicitMatrixDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitImplicitMatrixDeclaration(s)
	}
}

func (s *ImplicitMatrixDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitImplicitMatrixDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImplicitDeclarationContext struct {
	VariableDeclarationContext
}

func NewImplicitDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImplicitDeclarationContext {
	var p = new(ImplicitDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ImplicitDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplicitDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *ImplicitDeclarationContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ImplicitDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterImplicitDeclaration(s)
	}
}

func (s *ImplicitDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitImplicitDeclaration(s)
	}
}

func (s *ImplicitDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitImplicitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExplicitDeclarationContext struct {
	VariableDeclarationContext
}

func NewExplicitDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExplicitDeclarationContext {
	var p = new(ExplicitDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ExplicitDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplicitDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *ExplicitDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *ExplicitDeclarationContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ExplicitDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterExplicitDeclaration(s)
	}
}

func (s *ExplicitDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitExplicitDeclaration(s)
	}
}

func (s *ExplicitDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitExplicitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImplicitSliceDeclarationContext struct {
	VariableDeclarationContext
}

func NewImplicitSliceDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImplicitSliceDeclarationContext {
	var p = new(ImplicitSliceDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ImplicitSliceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplicitSliceDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *ImplicitSliceDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *ImplicitSliceDeclarationContext) SliceElements() ISliceElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceElementsContext)
}

func (s *ImplicitSliceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterImplicitSliceDeclaration(s)
	}
}

func (s *ImplicitSliceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitImplicitSliceDeclaration(s)
	}
}

func (s *ImplicitSliceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitImplicitSliceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LanguageParserRULE_variableDeclaration)
	var _la int

	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExplicitDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(305)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(308)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Match(LanguageParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__22 {
			{
				p.SetState(310)
				p.Match(LanguageParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(311)
				p.expressionStatement(0)
			}

		}

	case 2:
		localctx = NewExplicitSliceDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(314)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(317)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Match(LanguageParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Match(LanguageParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__22 {
			{
				p.SetState(320)
				p.Match(LanguageParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case LanguageParserT__0:
				{
					p.SetState(321)
					p.SliceElements()
				}

			case LanguageParserT__2, LanguageParserT__3, LanguageParserT__17, LanguageParserT__19, LanguageParserT__27, LanguageParserT__40, LanguageParserT__41, LanguageParserT__42, LanguageParserINTEGER, LanguageParserBOOLEAN, LanguageParserFLOAT, LanguageParserSTRING, LanguageParserID:
				{
					p.SetState(322)
					p.expressionStatement(0)
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}

	case 3:
		localctx = NewExplicitMatrixDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(327)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(330)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(LanguageParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(LanguageParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__22 {
			{
				p.SetState(333)
				p.Match(LanguageParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(334)
				p.MatrixElements()
			}

		}

	case 4:
		localctx = NewExplicitStructDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(337)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(340)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__22 {
			{
				p.SetState(342)
				p.Match(LanguageParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(343)
				p.Match(LanguageParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == LanguageParserID {
				{
					p.SetState(344)
					p.StructInitialization()
				}

			}
			{
				p.SetState(347)
				p.Match(LanguageParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		localctx = NewImplicitStructDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(350)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(353)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Match(LanguageParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(LanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserID {
			{
				p.SetState(357)
				p.StructInitialization()
			}

		}
		{
			p.SetState(360)
			p.Match(LanguageParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewImplicitSliceDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(361)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(364)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(LanguageParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(LanguageParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(LanguageParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.SliceElements()
		}

	case 7:
		localctx = NewImplicitMatrixDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(369)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(372)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Match(LanguageParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(LanguageParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(LanguageParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.MatrixElements()
		}

	case 8:
		localctx = NewImplicitDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LanguageParserT__44 {
			{
				p.SetState(377)
				p.Match(LanguageParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(380)
			p.Match(LanguageParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(LanguageParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.expressionStatement(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceElementsContext is an interface to support dynamic dispatch.
type ISliceElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionStatement() []IExpressionStatementContext
	ExpressionStatement(i int) IExpressionStatementContext

	// IsSliceElementsContext differentiates from other interfaces.
	IsSliceElementsContext()
}

type SliceElementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceElementsContext() *SliceElementsContext {
	var p = new(SliceElementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_sliceElements
	return p
}

func InitEmptySliceElementsContext(p *SliceElementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_sliceElements
}

func (*SliceElementsContext) IsSliceElementsContext() {}

func NewSliceElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceElementsContext {
	var p = new(SliceElementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_sliceElements

	return p
}

func (s *SliceElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceElementsContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *SliceElementsContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *SliceElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterSliceElements(s)
	}
}

func (s *SliceElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitSliceElements(s)
	}
}

func (s *SliceElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitSliceElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) SliceElements() (localctx ISliceElementsContext) {
	localctx = NewSliceElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LanguageParserRULE_sliceElements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(LanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279238570329505816) != 0 {
		{
			p.SetState(386)
			p.expressionStatement(0)
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LanguageParserT__20 {
			{
				p.SetState(387)
				p.Match(LanguageParserT__20)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(388)
				p.expressionStatement(0)
			}

			p.SetState(393)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(396)
		p.Match(LanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatrixElementsContext is an interface to support dynamic dispatch.
type IMatrixElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSliceElements() []ISliceElementsContext
	SliceElements(i int) ISliceElementsContext

	// IsMatrixElementsContext differentiates from other interfaces.
	IsMatrixElementsContext()
}

type MatrixElementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixElementsContext() *MatrixElementsContext {
	var p = new(MatrixElementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_matrixElements
	return p
}

func InitEmptyMatrixElementsContext(p *MatrixElementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_matrixElements
}

func (*MatrixElementsContext) IsMatrixElementsContext() {}

func NewMatrixElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixElementsContext {
	var p = new(MatrixElementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_matrixElements

	return p
}

func (s *MatrixElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixElementsContext) AllSliceElements() []ISliceElementsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISliceElementsContext); ok {
			len++
		}
	}

	tst := make([]ISliceElementsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISliceElementsContext); ok {
			tst[i] = t.(ISliceElementsContext)
			i++
		}
	}

	return tst
}

func (s *MatrixElementsContext) SliceElements(i int) ISliceElementsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceElementsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceElementsContext)
}

func (s *MatrixElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterMatrixElements(s)
	}
}

func (s *MatrixElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitMatrixElements(s)
	}
}

func (s *MatrixElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitMatrixElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) MatrixElements() (localctx IMatrixElementsContext) {
	localctx = NewMatrixElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LanguageParserRULE_matrixElements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(LanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LanguageParserT__0 {
		{
			p.SetState(399)
			p.SliceElements()
		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LanguageParserT__20 {
			{
				p.SetState(400)
				p.Match(LanguageParserT__20)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(401)
				p.SliceElements()
			}

			p.SetState(406)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(409)
		p.Match(LanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructInitializationContext is an interface to support dynamic dispatch.
type IStructInitializationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructFieldInit() []IStructFieldInitContext
	StructFieldInit(i int) IStructFieldInitContext

	// IsStructInitializationContext differentiates from other interfaces.
	IsStructInitializationContext()
}

type StructInitializationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructInitializationContext() *StructInitializationContext {
	var p = new(StructInitializationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structInitialization
	return p
}

func InitEmptyStructInitializationContext(p *StructInitializationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structInitialization
}

func (*StructInitializationContext) IsStructInitializationContext() {}

func NewStructInitializationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructInitializationContext {
	var p = new(StructInitializationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_structInitialization

	return p
}

func (s *StructInitializationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructInitializationContext) AllStructFieldInit() []IStructFieldInitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldInitContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldInitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldInitContext); ok {
			tst[i] = t.(IStructFieldInitContext)
			i++
		}
	}

	return tst
}

func (s *StructInitializationContext) StructFieldInit(i int) IStructFieldInitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldInitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldInitContext)
}

func (s *StructInitializationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitializationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructInitializationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStructInitialization(s)
	}
}

func (s *StructInitializationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStructInitialization(s)
	}
}

func (s *StructInitializationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStructInitialization(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) StructInitialization() (localctx IStructInitializationContext) {
	localctx = NewStructInitializationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LanguageParserRULE_structInitialization)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.StructFieldInit()
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserT__20 {
		{
			p.SetState(412)
			p.Match(LanguageParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.StructFieldInit()
		}

		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFieldInitContext is an interface to support dynamic dispatch.
type IStructFieldInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ExpressionStatement() IExpressionStatementContext

	// IsStructFieldInitContext differentiates from other interfaces.
	IsStructFieldInitContext()
}

type StructFieldInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldInitContext() *StructFieldInitContext {
	var p = new(StructFieldInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structFieldInit
	return p
}

func InitEmptyStructFieldInitContext(p *StructFieldInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structFieldInit
}

func (*StructFieldInitContext) IsStructFieldInitContext() {}

func NewStructFieldInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldInitContext {
	var p = new(StructFieldInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_structFieldInit

	return p
}

func (s *StructFieldInitContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldInitContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *StructFieldInitContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StructFieldInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStructFieldInit(s)
	}
}

func (s *StructFieldInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStructFieldInit(s)
	}
}

func (s *StructFieldInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStructFieldInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) StructFieldInit() (localctx IStructFieldInitContext) {
	localctx = NewStructFieldInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LanguageParserRULE_structFieldInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(LanguageParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.Match(LanguageParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.expressionStatement(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclarationContext is an interface to support dynamic dispatch.
type IStructDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllStructAtribute() []IStructAtributeContext
	StructAtribute(i int) IStructAtributeContext

	// IsStructDeclarationContext differentiates from other interfaces.
	IsStructDeclarationContext()
}

type StructDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclarationContext() *StructDeclarationContext {
	var p = new(StructDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structDeclaration
	return p
}

func InitEmptyStructDeclarationContext(p *StructDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structDeclaration
}

func (*StructDeclarationContext) IsStructDeclarationContext() {}

func NewStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_structDeclaration

	return p
}

func (s *StructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(LanguageParserID, 0)
}

func (s *StructDeclarationContext) AllStructAtribute() []IStructAtributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructAtributeContext); ok {
			len++
		}
	}

	tst := make([]IStructAtributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructAtributeContext); ok {
			tst[i] = t.(IStructAtributeContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclarationContext) StructAtribute(i int) IStructAtributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructAtributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructAtributeContext)
}

func (s *StructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) StructDeclaration() (localctx IStructDeclarationContext) {
	localctx = NewStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LanguageParserRULE_structDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Match(LanguageParserT__48)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(424)
		p.Match(LanguageParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(425)
		p.Match(LanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserTYPE || _la == LanguageParserID {
		{
			p.SetState(426)
			p.StructAtribute()
		}

		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(432)
		p.Match(LanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructAtributeContext is an interface to support dynamic dispatch.
type IStructAtributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	TYPE() antlr.TerminalNode

	// IsStructAtributeContext differentiates from other interfaces.
	IsStructAtributeContext()
}

type StructAtributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructAtributeContext() *StructAtributeContext {
	var p = new(StructAtributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structAtribute
	return p
}

func InitEmptyStructAtributeContext(p *StructAtributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_structAtribute
}

func (*StructAtributeContext) IsStructAtributeContext() {}

func NewStructAtributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructAtributeContext {
	var p = new(StructAtributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_structAtribute

	return p
}

func (s *StructAtributeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructAtributeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserID)
}

func (s *StructAtributeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserID, i)
}

func (s *StructAtributeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *StructAtributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAtributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructAtributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStructAtribute(s)
	}
}

func (s *StructAtributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStructAtribute(s)
	}
}

func (s *StructAtributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStructAtribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) StructAtribute() (localctx IStructAtributeContext) {
	localctx = NewStructAtributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LanguageParserRULE_structAtribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LanguageParserTYPE || _la == LanguageParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(435)
		p.Match(LanguageParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	BlockStatement() IBlockStatementContext
	ParameterList() IParameterListContext
	TYPE() antlr.TerminalNode

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserID)
}

func (s *FunctionDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserID, i)
}

func (s *FunctionDeclarationContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *FunctionDeclarationContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LanguageParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.Match(LanguageParserT__49)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.Match(LanguageParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		p.Match(LanguageParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LanguageParserID {
		{
			p.SetState(440)
			p.ParameterList()
		}

	}
	{
		p.SetState(443)
		p.Match(LanguageParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&150870587516911616) != 0 {
		{
			p.SetState(444)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&150870587516911616) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(447)
		p.BlockStatement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LanguageParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Parameter()
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserT__20 {
		{
			p.SetState(450)
			p.Match(LanguageParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.Parameter()
		}

		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	TYPE() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserID)
}

func (s *ParameterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserID, i)
}

func (s *ParameterContext) TYPE() antlr.TerminalNode {
	return s.GetToken(LanguageParserTYPE, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LanguageParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		p.Match(LanguageParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LanguageParserTYPE || _la == LanguageParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionStatement() []IExpressionStatementContext
	ExpressionStatement(i int) IExpressionStatementContext

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LanguageParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpressionStatement() []IExpressionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			len++
		}
	}

	tst := make([]IExpressionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionStatementContext); ok {
			tst[i] = t.(IExpressionStatementContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) ExpressionStatement(i int) IExpressionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LanguageParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.expressionStatement(0)
	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserT__20 {
		{
			p.SetState(461)
			p.Match(LanguageParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.expressionStatement(0)
		}

		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *LanguageParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionStatementContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionStatementContext)
		}
		return p.ExpressionStatement_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LanguageParser) ExpressionStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 22)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
