package classify

import (
	"fmt"
	"slices"
	"strings"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

var prohibitedFactors = []string{
	"social_scoring",
	"subliminal_manipulation",
	"emotion_inference_workplace",
	"biometric_categorisation_sensitive",
}

var highRiskDomains = []string{
	"employment",
	"creditworthiness",
	"education",
	"essential_services",
	"law_enforcement",
	"migration",
	"justice",
	"critical_infrastructure",
}

var transparencyCapabilities = []string{
	"conversational_agent",
	"content_generation",
	"synthetic_media",
	"emotion_recognition",
}

// Evaluate provides deterministic risk routing for governance workflow triage.
func Evaluate(system domain.System) domain.Classification {
	tier := "minimal"
	reasons := []string{"No configured high-risk or transparency trigger matched."}

	if factor, ok := firstMatch(system.RiskFactors, prohibitedFactors); ok {
		tier = "prohibited-review"
		reasons = []string{fmt.Sprintf("Risk factor %q requires prohibited-practice review.", factor)}
	} else if slices.Contains(highRiskDomains, strings.ToLower(system.Domain)) {
		tier = "high"
		reasons = []string{fmt.Sprintf("Domain %q is routed to the high-impact control set.", system.Domain)}
	} else if capability, ok := firstMatch(system.Capabilities, transparencyCapabilities); ok {
		tier = "transparency"
		reasons = []string{fmt.Sprintf("Capability %q triggers transparency controls.", capability)}
	}

	declared := strings.ToLower(system.DeclaredRiskTier)
	return domain.Classification{
		Tier:                tier,
		Reasons:             reasons,
		DeclaredTier:        declared,
		DeclarationMismatch: declared != "" && declared != tier,
	}
}

func firstMatch(values, candidates []string) (string, bool) {
	for _, value := range values {
		if slices.Contains(candidates, strings.ToLower(value)) {
			return value, true
		}
	}
	return "", false
}
