paas: [ "$E2E_WITH_COVERAGE" = 'true' ] && COMMAND='GOCOVERDIR=${PAAS_COVERAGE_DIR:-/tmp/coverage/paas} go run -cover cmd/manager/main.go' || COMMAND='go run cmd/manager/main.go' && sh -c "$COMMAND"
