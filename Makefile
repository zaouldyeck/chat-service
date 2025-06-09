# Check to see if ash shell is available, otherwise use zsh
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/zsh)

# =============================================================
# Chat

chat-run:
	go run api/services/cap/main.go | go run api/tooling/logfmt/main.go

# =============================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor
