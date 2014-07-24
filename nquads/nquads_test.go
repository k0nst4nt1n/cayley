// Copyright 2014 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nquads

import (
	"reflect"
	"testing"

	"github.com/google/cayley/graph"
)

var testNTriples = []struct {
	message string
	input   string
	expect  *graph.Triple
	err     error
}{
	// Tests taken from http://www.w3.org/TR/n-quads/ and http://www.w3.org/TR/n-triples/.

	// N-Triples example 1.
	{
		message: "parse triple with commment",
		input:   `<http://one.example/subject1> <http://one.example/predicate1> <http://one.example/object1> . # comments here`,
		expect: &graph.Triple{
			Subject:    "<http://one.example/subject1>",
			Predicate:  "<http://one.example/predicate1>",
			Object:     "<http://one.example/object1>",
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with blank subject node, literal object and no comment (1)",
		input:   `_:subject1 <http://an.example/predicate1> "object1" .`,
		expect: &graph.Triple{
			Subject:    "_:subject1",
			Predicate:  "<http://an.example/predicate1>",
			Object:     `"object1"`,
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with blank subject node, literal object and no comment (2)",
		input:   `_:subject2 <http://an.example/predicate2> "object2" .`,
		expect: &graph.Triple{
			Subject:    "_:subject2",
			Predicate:  "<http://an.example/predicate2>",
			Object:     `"object2"`,
			Provenance: "",
		},
		err: nil,
	},

	// N-Triples example 2.
	{
		message: "parse triple with three IRIREFs",
		input:   `<http://example.org/#spiderman> <http://www.perceive.net/schemas/relationship/enemyOf> <http://example.org/#green-goblin> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/#spiderman>",
			Predicate:  "<http://www.perceive.net/schemas/relationship/enemyOf>",
			Object:     "<http://example.org/#green-goblin>",
			Provenance: "",
		},
		err: nil,
	},

	// N-Triples example 3.
	{
		message: "parse triple with blank node labelled subject and object and IRIREF predicate (1)",
		input:   `_:alice <http://xmlns.com/foaf/0.1/knows> _:bob .`,
		expect: &graph.Triple{
			Subject:    "_:alice",
			Predicate:  "<http://xmlns.com/foaf/0.1/knows>",
			Object:     "_:bob",
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with blank node labelled subject and object and IRIREF predicate (2)",
		input:   `_:bob <http://xmlns.com/foaf/0.1/knows> _:alice .`,
		expect: &graph.Triple{
			Subject:    "_:bob",
			Predicate:  "<http://xmlns.com/foaf/0.1/knows>",
			Object:     "_:alice",
			Provenance: "",
		},
		err: nil,
	},

	// N-Quads example 1.
	{
		message: "parse quad with commment",
		input:   `<http://one.example/subject1> <http://one.example/predicate1> <http://one.example/object1> <http://example.org/graph3> . # comments here`,
		expect: &graph.Triple{
			Subject:    "<http://one.example/subject1>",
			Predicate:  "<http://one.example/predicate1>",
			Object:     "<http://one.example/object1>",
			Provenance: "<http://example.org/graph3>",
		},
		err: nil,
	},
	{
		message: "parse quad with blank subject node, literal object, IRIREF predicate and label, and no comment (1)",
		input:   `_:subject1 <http://an.example/predicate1> "object1" <http://example.org/graph1> .`,
		expect: &graph.Triple{
			Subject:    "_:subject1",
			Predicate:  "<http://an.example/predicate1>",
			Object:     `"object1"`,
			Provenance: "<http://example.org/graph1>",
		},
		err: nil,
	},
	{
		message: "parse quad with blank subject node, literal object, IRIREF predicate and label, and no comment (2)",
		input:   `_:subject2 <http://an.example/predicate2> "object2" <http://example.org/graph5> .`,
		expect: &graph.Triple{
			Subject:    "_:subject2",
			Predicate:  "<http://an.example/predicate2>",
			Object:     `"object2"`,
			Provenance: "<http://example.org/graph5>",
		},
		err: nil,
	},

	// N-Quads example 2.
	{
		message: "parse quad with all IRIREF parts",
		input:   `<http://example.org/#spiderman> <http://www.perceive.net/schemas/relationship/enemyOf> <http://example.org/#green-goblin> <http://example.org/graphs/spiderman> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/#spiderman>",
			Predicate:  "<http://www.perceive.net/schemas/relationship/enemyOf>",
			Object:     "<http://example.org/#green-goblin>",
			Provenance: "<http://example.org/graphs/spiderman>",
		},
		err: nil,
	},

	// N-Quads example 3.
	{
		message: "parse quad with blank node labelled subject and object and IRIREF predicate and label (1)",
		input:   `_:alice <http://xmlns.com/foaf/0.1/knows> _:bob <http://example.org/graphs/john> .`,
		expect: &graph.Triple{
			Subject:    "_:alice",
			Predicate:  "<http://xmlns.com/foaf/0.1/knows>",
			Object:     "_:bob",
			Provenance: "<http://example.org/graphs/john>",
		},
		err: nil,
	},
	{
		message: "parse quad with blank node labelled subject and object and IRIREF predicate and label (2)",
		input:   `_:bob <http://xmlns.com/foaf/0.1/knows> _:alice <http://example.org/graphs/james> .`,
		expect: &graph.Triple{
			Subject:    "_:bob",
			Predicate:  "<http://xmlns.com/foaf/0.1/knows>",
			Object:     "_:alice",
			Provenance: "<http://example.org/graphs/james>",
		},
		err: nil,
	},

	// N-Triples tests.
	{
		message: "parse triple with all IRIREF parts",
		input:   `<http://example.org/bob#me> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://xmlns.com/foaf/0.1/Person> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>",
			Object:     "<http://xmlns.com/foaf/0.1/Person>",
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with all IRIREF parts",
		input:   `<http://example.org/bob#me> <http://xmlns.com/foaf/0.1/knows> <http://example.org/alice#me> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://xmlns.com/foaf/0.1/knows>",
			Object:     "<http://example.org/alice#me>",
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with IRIREF schema on literal object",
		input:   `<http://example.org/bob#me> <http://schema.org/birthDate> "1990-07-04"^^<http://www.w3.org/2001/XMLSchema#date> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://schema.org/birthDate>",
			Object:     `"1990-07-04"^^<http://www.w3.org/2001/XMLSchema#date>`,
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse commented IRIREF in triple",
		input:   `<http://example.org/bob#me> <http://xmlns.com/foaf/0.1/topic_interest> <http://www.wikidata.org/entity/Q12418> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://xmlns.com/foaf/0.1/topic_interest>",
			Object:     "<http://www.wikidata.org/entity/Q12418>",
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with literal subject",
		input:   `<http://www.wikidata.org/entity/Q12418> <http://purl.org/dc/terms/title> "Mona Lisa" .`,
		expect: &graph.Triple{
			Subject:    "<http://www.wikidata.org/entity/Q12418>",
			Predicate:  "<http://purl.org/dc/terms/title>",
			Object:     `"Mona Lisa"`,
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with all IRIREF parts (1)",
		input:   `<http://www.wikidata.org/entity/Q12418> <http://purl.org/dc/terms/creator> <http://dbpedia.org/resource/Leonardo_da_Vinci> .`,
		expect: &graph.Triple{
			Subject:    "<http://www.wikidata.org/entity/Q12418>",
			Predicate:  "<http://purl.org/dc/terms/creator>",
			Object:     "<http://dbpedia.org/resource/Leonardo_da_Vinci>",
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse triple with all IRIREF parts (2)",
		input:   `<http://data.europeana.eu/item/04802/243FA8618938F4117025F17A8B813C5F9AA4D619> <http://purl.org/dc/terms/subject> <http://www.wikidata.org/entity/Q12418> .`,
		expect: &graph.Triple{
			Subject:    "<http://data.europeana.eu/item/04802/243FA8618938F4117025F17A8B813C5F9AA4D619>",
			Predicate:  "<http://purl.org/dc/terms/subject>",
			Object:     "<http://www.wikidata.org/entity/Q12418>",
			Provenance: "",
		},
		err: nil,
	},

	// N-Quads tests.
	{
		message: "parse commented IRIREF in quad (1)",
		input:   `<http://example.org/bob#me> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://xmlns.com/foaf/0.1/Person> <http://example.org/bob> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>",
			Object:     "<http://xmlns.com/foaf/0.1/Person>",
			Provenance: "<http://example.org/bob>",
		},
		err: nil,
	},
	{
		message: "parse quad with all IRIREF parts",
		input:   `<http://example.org/bob#me> <http://xmlns.com/foaf/0.1/knows> <http://example.org/alice#me> <http://example.org/bob> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://xmlns.com/foaf/0.1/knows>",
			Object:     "<http://example.org/alice#me>",
			Provenance: "<http://example.org/bob>",
		},
		err: nil,
	},
	{
		message: "parse quad with IRIREF schema on literal object",
		input:   `<http://example.org/bob#me> <http://schema.org/birthDate> "1990-07-04"^^<http://www.w3.org/2001/XMLSchema#date> <http://example.org/bob> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://schema.org/birthDate>",
			Object:     `"1990-07-04"^^<http://www.w3.org/2001/XMLSchema#date>`,
			Provenance: "<http://example.org/bob>",
		},
		err: nil,
	},
	{
		message: "parse commented IRIREF in quad (2)",
		input:   `<http://example.org/bob#me> <http://xmlns.com/foaf/0.1/topic_interest> <http://www.wikidata.org/entity/Q12418> <http://example.org/bob> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob#me>",
			Predicate:  "<http://xmlns.com/foaf/0.1/topic_interest>",
			Object:     "<http://www.wikidata.org/entity/Q12418>",
			Provenance: "<http://example.org/bob>",
		},
		err: nil,
	},
	{
		message: "parse literal object and colon qualified label in quad",
		input:   `<http://www.wikidata.org/entity/Q12418> <http://purl.org/dc/terms/title> "Mona Lisa" <https://www.wikidata.org/wiki/Special:EntityData/Q12418> .`,
		expect: &graph.Triple{
			Subject:    "<http://www.wikidata.org/entity/Q12418>",
			Predicate:  "<http://purl.org/dc/terms/title>",
			Object:     `"Mona Lisa"`,
			Provenance: "<https://www.wikidata.org/wiki/Special:EntityData/Q12418>",
		},
		err: nil,
	},
	{
		message: "parse all IRIREF parts with colon qualified label in quad (1)",
		input:   `<http://www.wikidata.org/entity/Q12418> <http://purl.org/dc/terms/creator> <http://dbpedia.org/resource/Leonardo_da_Vinci> <https://www.wikidata.org/wiki/Special:EntityData/Q12418> .`,
		expect: &graph.Triple{
			Subject:    "<http://www.wikidata.org/entity/Q12418>",
			Predicate:  "<http://purl.org/dc/terms/creator>",
			Object:     "<http://dbpedia.org/resource/Leonardo_da_Vinci>",
			Provenance: "<https://www.wikidata.org/wiki/Special:EntityData/Q12418>",
		},
		err: nil,
	},
	{
		message: "parse all IRIREF parts with colon qualified label in quad (2)",
		input:   `<http://data.europeana.eu/item/04802/243FA8618938F4117025F17A8B813C5F9AA4D619> <http://purl.org/dc/terms/subject> <http://www.wikidata.org/entity/Q12418> <https://www.wikidata.org/wiki/Special:EntityData/Q12418> .`,
		expect: &graph.Triple{
			Subject:    "<http://data.europeana.eu/item/04802/243FA8618938F4117025F17A8B813C5F9AA4D619>",
			Predicate:  "<http://purl.org/dc/terms/subject>",
			Object:     "<http://www.wikidata.org/entity/Q12418>",
			Provenance: "<https://www.wikidata.org/wiki/Special:EntityData/Q12418>",
		},
		err: nil,
	},
	{
		message: "parse all IRIREF parts (quad section - 1)",
		input:   `<http://example.org/bob> <http://purl.org/dc/terms/publisher> <http://example.org> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob>",
			Predicate:  "<http://purl.org/dc/terms/publisher>",
			Object:     "<http://example.org>",
			Provenance: "",
		},
		err: nil,
	},
	{
		message: "parse all IRIREF parts (quad section - 2)",
		input:   `<http://example.org/bob> <http://purl.org/dc/terms/rights> <http://creativecommons.org/licenses/by/3.0/> .`,
		expect: &graph.Triple{
			Subject:    "<http://example.org/bob>",
			Predicate:  "<http://purl.org/dc/terms/rights>",
			Object:     "<http://creativecommons.org/licenses/by/3.0/>",
			Provenance: "",
		},
		err: nil,
	},
}

func TestParse(t *testing.T) {
	for _, test := range testNTriples {
		got, err := Parse(test.input)
		if err != test.err {
			t.Errorf("Unexpected error when %s: got:%v expect:%v", test.message, err, test.err)
		}
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Failed to %s, %q, got:%q expect:%q", test.message, test.input, got, test.expect)
		}
	}
}

var result *graph.Triple

func BenchmarkParser(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result, _ = Parse("<http://example/s> <http://example/p> \"object of some real\\tlength\"@en . # comment")
	}
}