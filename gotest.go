package rules

import (
	"go/ast"

	"github.com/securego/gosec"
)

type weakRand struct {
	gosec.MetaData
	funcNames   []string
	packagePath string
}

func (w *weakRand) ID() string {
	return w.MetaData.ID
}

func (w *weakRand) Match(n ast.Node, c *gosec.Context) (*gosec.Issue, error) {
	for _, funcName := range w.funcNames {
		if _, matched := gosec.MatchCallByPackage(n, c, w.packagePath, funcName); matched {
			return gosec.NewIssue(c, n, w.ID(), w.What, w.Severity, w.Confidence), nil
		}
	}

	return nil, nil
}
