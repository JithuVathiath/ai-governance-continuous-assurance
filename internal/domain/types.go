package domain

// System describes an AI use case at a specific point in its lifecycle.
type System struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Version          string   `json:"version"`
	Owner            string   `json:"owner"`
	Role             string   `json:"role"`
	Domain           string   `json:"domain"`
	LifecycleStage   string   `json:"lifecycle_stage"`
	IntendedPurpose  string   `json:"intended_purpose"`
	DecisionImpact   string   `json:"decision_impact"`
	Geographies      []string `json:"geographies"`
	Capabilities     []string `json:"capabilities"`
	RiskFactors      []string `json:"risk_factors"`
	DataSensitivity  []string `json:"data_sensitivity"`
	AffectedGroups   []string `json:"affected_groups"`
	DeclaredRiskTier string   `json:"declared_risk_tier,omitempty"`
	Vendor           string   `json:"vendor,omitempty"`
	DeploymentDate   string   `json:"deployment_date,omitempty"`
}

// Evidence is one artefact offered in support of one or more controls.
type Evidence struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Owner       string `json:"owner"`
	Artifact    string `json:"artifact"`
	SHA256      string `json:"sha256"`
	CollectedAt string `json:"collected_at"`
	ValidUntil  string `json:"valid_until,omitempty"`
	Integrity   string `json:"integrity,omitempty"`
}

// EvidenceBundle groups the evidence for one system version.
type EvidenceBundle struct {
	SystemID string     `json:"system_id"`
	Version  string     `json:"version"`
	Items    []Evidence `json:"items"`
}

// PolicyPack is a versioned collection of independently testable controls.
type PolicyPack struct {
	ID         string `json:"id"`
	Framework  string `json:"framework"`
	Version    string `json:"version"`
	AsOf       string `json:"as_of"`
	SourceURL  string `json:"source_url"`
	Disclaimer string `json:"disclaimer"`
	Rules      []Rule `json:"rules"`
	Digest     string `json:"digest,omitempty"`
}

// Rule defines applicability and the evidence requirements for one control.
type Rule struct {
	ID             string        `json:"id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	SourceRef      string        `json:"source_ref"`
	Severity       string        `json:"severity"`
	ApplicableWhen Conditions    `json:"applicable_when"`
	Requirements   []Requirement `json:"requirements"`
}

// Conditions uses OR within a field and AND between populated fields.
type Conditions struct {
	RiskTiers      []string `json:"risk_tiers,omitempty"`
	Domains        []string `json:"domains,omitempty"`
	Roles          []string `json:"roles,omitempty"`
	Geographies    []string `json:"geographies,omitempty"`
	Capabilities   []string `json:"capabilities,omitempty"`
	RiskFactors    []string `json:"risk_factors,omitempty"`
	LifecycleStage []string `json:"lifecycle_stages,omitempty"`
}

// Requirement describes the accepted evidence for a control outcome.
type Requirement struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	EvidenceTypes []string `json:"evidence_types"`
	Minimum       int      `json:"minimum"`
}

// Classification is a traceable risk-routing result, not a legal opinion.
type Classification struct {
	Tier                string   `json:"tier"`
	Reasons             []string `json:"reasons"`
	DeclaredTier        string   `json:"declared_tier,omitempty"`
	DeclarationMismatch bool     `json:"declaration_mismatch"`
}

// PolicyVersion records the exact policy material used by an assessment.
type PolicyVersion struct {
	PackID     string `json:"pack_id"`
	Framework  string `json:"framework"`
	Version    string `json:"version"`
	AsOf       string `json:"as_of"`
	SourceURL  string `json:"source_url"`
	PackSHA256 string `json:"pack_sha256"`
}

// RequirementResult explains which evidence supported one requirement.
type RequirementResult struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Status      string   `json:"status"`
	EvidenceIDs []string `json:"evidence_ids"`
	Reason      string   `json:"reason"`
}

// ControlResult is the evidence-backed output for one applicable policy rule.
type ControlResult struct {
	ID           string              `json:"id"`
	Framework    string              `json:"framework"`
	PackID       string              `json:"pack_id"`
	Title        string              `json:"title"`
	SourceRef    string              `json:"source_ref"`
	Severity     string              `json:"severity"`
	Applicable   bool                `json:"applicable"`
	Status       string              `json:"status"`
	Reason       string              `json:"reason"`
	Requirements []RequirementResult `json:"requirements"`
}

// Finding is a machine-readable gap suitable for reports and SARIF.
type Finding struct {
	ID        string `json:"id"`
	ControlID string `json:"control_id"`
	Framework string `json:"framework"`
	Severity  string `json:"severity"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Message   string `json:"message"`
}

// Summary deliberately reports evidence coverage rather than legal compliance.
type Summary struct {
	Applicable       int     `json:"applicable"`
	Supported        int     `json:"supported"`
	Partial          int     `json:"partial"`
	Missing          int     `json:"missing"`
	Expired          int     `json:"expired"`
	EvidenceCoverage float64 `json:"evidence_coverage_percent"`
	Decision         string  `json:"decision"`
}

// GraphNode is one entity in the traceable governance evidence graph.
type GraphNode struct {
	ID    string            `json:"id"`
	Type  string            `json:"type"`
	Label string            `json:"label"`
	Meta  map[string]string `json:"meta,omitempty"`
}

// GraphEdge links evidence and governance entities.
type GraphEdge struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Relation string `json:"relation"`
}

// EvidenceGraph makes every conclusion traceable to a source and artefact.
type EvidenceGraph struct {
	Nodes []GraphNode `json:"nodes"`
	Edges []GraphEdge `json:"edges"`
}

// LedgerReceipt locates an assessment in the tamper-evident event chain.
type LedgerReceipt struct {
	Sequence     int64  `json:"sequence"`
	EventHash    string `json:"event_hash"`
	PreviousHash string `json:"previous_hash"`
	PayloadHash  string `json:"payload_hash"`
}

// Assessment is the complete continuous-assurance output.
type Assessment struct {
	SchemaVersion  string          `json:"schema_version"`
	AssessmentID   string          `json:"assessment_id"`
	GeneratedAt    string          `json:"generated_at"`
	System         System          `json:"system"`
	Classification Classification  `json:"classification"`
	PolicyVersions []PolicyVersion `json:"policy_versions"`
	Evidence       []Evidence      `json:"evidence"`
	Controls       []ControlResult `json:"controls"`
	Findings       []Finding       `json:"findings"`
	Summary        Summary         `json:"summary"`
	Graph          EvidenceGraph   `json:"evidence_graph"`
	Ledger         *LedgerReceipt  `json:"ledger,omitempty"`
	Disclaimer     string          `json:"disclaimer"`
}

// DriftItem describes one control change between two assessments.
type DriftItem struct {
	Key        string `json:"key"`
	Framework  string `json:"framework"`
	ControlID  string `json:"control_id"`
	Title      string `json:"title"`
	Before     string `json:"before"`
	After      string `json:"after"`
	ChangeType string `json:"change_type"`
	Severity   string `json:"severity"`
}

// DriftReport captures changes in scope and evidence readiness.
type DriftReport struct {
	SchemaVersion     string      `json:"schema_version"`
	SystemID          string      `json:"system_id"`
	BeforeAssessment  string      `json:"before_assessment"`
	AfterAssessment   string      `json:"after_assessment"`
	BeforeTier        string      `json:"before_tier"`
	AfterTier         string      `json:"after_tier"`
	NewControls       int         `json:"new_controls"`
	ResolvedControls  int         `json:"resolved_controls"`
	RegressedControls int         `json:"regressed_controls"`
	ChangedControls   int         `json:"changed_controls"`
	Items             []DriftItem `json:"items"`
}
