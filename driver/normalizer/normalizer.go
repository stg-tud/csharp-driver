package normalizer

import (
	"gopkg.in/bblfsh/sdk.v2/uast"
	. "gopkg.in/bblfsh/sdk.v2/uast/transformer"
)

var Preprocess = Transformers([][]Transformer{
	{Mappings(Preprocessors...)},
}...)

var Normalize = Transformers([][]Transformer{

	{Mappings(Normalizers...)},
}...)

// Preprocessors is a block of AST preprocessing rules rules.
var Preprocessors = []Mapping{
	Map(
		Part("_", Obj{
			"FullSpan": Obj{
				uast.KeyType: String("TextSpan"),
				"Length":     AnyNode(nil),
				"Start":      Var("start"),
				"End":        Var("end"),
			},
			// TODO(dennwc): add those as a custom position fields?
			"Span":      AnyNode(nil),
			"SpanStart": AnyNode(nil),
		}),
		Part("_", Obj{
			// remap to temporary keys and let ObjectToNode to pick them up
			"spanStart": Var("start"),
			"spanEnd":   Var("end"),
		}),
	),

	ObjectToNode{
		OffsetKey:    "spanStart",
		EndOffsetKey: "spanEnd",
	}.Mapping(),
}

// Normalizers is the main block of normalization rules to convert native AST to semantic UAST.
var Normalizers = []Mapping{
	MapSemantic("IdentifierNameSyntax", uast.Identifier{}, MapObj(
		Obj{
			"Identifier": Obj{
				uast.KeyType: String("SyntaxToken"),
				// TODO(dennwc): assert that it's the same as in parent
				uast.KeyPos: AnyNode(nil),

				"LeadingTrivia":  Arr(),
				"TrailingTrivia": Arr(),
				"RawKind":        Int(8508),
				"Text":           Var("name"),
				"Value":          Var("name"),
				"ValueText":      Var("name"),
			},
			"RawKind": Int(8616),
		},
		Obj{
			"Name": Var("name"),
		},
	)),
}
