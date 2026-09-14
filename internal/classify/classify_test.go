package classify

import (
	"testing"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

func TestEvaluateRoutesRiskTiers(t *testing.T) {
	tests := []struct {
		name   string
		system domain.System
		want   string
	}{
		{"prohibited review takes priority", domain.System{RiskFactors: []string{"social_scoring"}}, "prohibited-review"},
		{"high impact domain", domain.System{Domain: "employment"}, "high"},
		{"transparency capability", domain.System{Capabilities: []string{"conversational_agent"}}, "transparency"},
		{"minimal fallback", domain.System{Domain: "inventory"}, "minimal"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Evaluate(test.system).Tier; got != test.want {
				t.Fatalf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestEvaluateFlagsDeclarationMismatch(t *testing.T) {
	result := Evaluate(domain.System{Domain: "creditworthiness", DeclaredRiskTier: "minimal"})
	if !result.DeclarationMismatch || result.Tier != "high" {
		t.Fatalf("unexpected classification: %#v", result)
	}
}
