package sarif

import (
	"fmt"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/iojson"
)

// Write emits governance gaps in SARIF 2.1.0 for GitHub code-scanning ingestion.
func Write(path, systemPath string, assessment domain.Assessment) error {
	rules := make([]map[string]any, 0, len(assessment.Findings))
	results := make([]map[string]any, 0, len(assessment.Findings))
	for _, finding := range assessment.Findings {
		rules = append(rules, map[string]any{
			"id":               finding.ID,
			"shortDescription": map[string]string{"text": finding.Title},
			"help":             map[string]string{"text": "Supply reviewed evidence and obtain independent governance approval."},
			"properties":       map[string]any{"framework": finding.Framework, "severity": finding.Severity},
		})
		results = append(results, map[string]any{
			"ruleId":  finding.ID,
			"level":   sarifLevel(finding.Severity),
			"message": map[string]string{"text": finding.Message},
			"locations": []map[string]any{{
				"physicalLocation": map[string]any{
					"artifactLocation": map[string]string{"uri": systemPath},
					"region":           map[string]int{"startLine": 1},
				},
			}},
		})
	}
	document := map[string]any{
		"version": "2.1.0",
		"$schema": "https://json.schemastore.org/sarif-2.1.0.json",
		"runs": []map[string]any{{
			"tool": map[string]any{"driver": map[string]any{
				"name": "Aegis Continuous Assurance", "version": "0.1.0",
				"informationUri": "https://github.com/JithuVathiath/ai-governance-continuous-assurance",
				"rules":          rules,
			}},
			"results": results,
		}},
	}
	if err := iojson.Write(path, document); err != nil {
		return fmt.Errorf("write SARIF: %w", err)
	}
	return nil
}

func sarifLevel(severity string) string {
	switch severity {
	case "critical", "high":
		return "error"
	case "medium":
		return "warning"
	default:
		return "note"
	}
}
