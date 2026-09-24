/* For license and copyright information please see the LEGAL file in the code repository */

package rule_p

import (
	"encoding/json"
)

// https://github.com/gorules/zen-go/blob/master/zen.go

type EvaluationOptions struct {
	Trace    bool
	MaxDepth uint8
}

type EvaluationResponse struct {
	Performance string
	Result      json.RawMessage
	Trace       *json.RawMessage
}

type Engine interface {
	Evaluate(key string, context any) (*EvaluationResponse, error)
	EvaluateWithOpts(key string, context any, options EvaluationOptions) (*EvaluationResponse, error)
	GetDecision(key string) (Decision, error)
	CreateDecision(data []byte) (Decision, error)
	Dispose()

	// https://www.ibm.com/docs/en/imdmcdhsa/10.1.0?topic=rules-understanding-rule-engine-methods
	// https://www.ibm.com/docs/en/imdmcdhsa/11.6.0?topic=features-external-business-rules
	loadRules()
	assertFact()
	fireRules()
	clearFacts()
}

type Decision interface {
	Evaluate(context any) (*EvaluationResponse, error)
	EvaluateWithOpts(context any, options EvaluationOptions) (*EvaluationResponse, error)
	Dispose()
}

// type Condition
// type Action
