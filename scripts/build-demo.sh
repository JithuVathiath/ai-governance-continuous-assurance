#!/bin/sh
set -eu

GO_BIN="${GO_BIN:-go}"
ASSESSMENT_TIME="2026-09-13T12:00:00Z"

mkdir -p bin artifacts/public artifacts/run reports
"$GO_BIN" build -trimpath -o bin/aegis ./cmd/aegis
rm -f artifacts/run/aegis-ledger.db

run_case() {
  case_name="$1"
  output_name="$2"
  bin/aegis assess \
    --system "cases/$case_name/system.json" \
    --evidence "cases/$case_name/evidence.json" \
    --policies policies \
    --out "artifacts/public/$output_name.assessment.json" \
    --report "reports/$output_name.md" \
    --sarif "artifacts/public/$output_name.sarif" \
    --ledger artifacts/run/aegis-ledger.db \
    --at "$ASSESSMENT_TIME"
}

run_case recruitment-screening-v1 recruitment-screening-v1
run_case recruitment-screening-v2 recruitment-screening-v2
run_case consumer-credit-scoring consumer-credit-scoring
run_case customer-support-assistant customer-support-assistant
run_case foundation-model-procurement foundation-model-procurement

bin/aegis diff \
  --before artifacts/public/recruitment-screening-v1.assessment.json \
  --after artifacts/public/recruitment-screening-v2.assessment.json \
  --out artifacts/public/recruitment-screening-drift.json \
  --report reports/recruitment-screening-drift.md

bin/aegis verify-ledger \
  --ledger artifacts/run/aegis-ledger.db \
  --out artifacts/public/ledger-verification.json
