.Phony:build
build:
	cd cmd/ddd_on_golang && \
	GOOS=linux go build -o ../../bin/main

.Phony:format
format:
	go fmt ./...

.Phony:test
# ANSI color
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m
COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''
test:
	go test -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

.Phony:attack_hunter
attack_hunter:
	curl -X PUT \
		-H 'Content-Type: application/json' \
		-d '{"hunterId":1}' \
		"http://localhost:8080/monsters/1/attack" | jq

.Phony:attack_monster
attack_monster:
	curl -X PUT \
		-H 'Content-Type: application/json' \
		-d '{"monsterId":1}' \
		"http://localhost:8080/hunters/1/attack" | jq

.Phony:find_all_hunters
find_all_hunters:
	curl -X GET \
	-H 'Content-Type: application/json' \
	"http://localhost:8080/hunters/" | jq

.Phony:find_all_monsters
find_all_monsters:
	curl -X GET \
	-H 'Content-Type: application/json' \
	"http://localhost:8080/monsters/" | jq

.Phony:find_monster
find_monster:
	curl -X GET \
		-H 'Content-Type: application/json' \
		"http://localhost:8080/monsters/10" | jq

.Phony:get_material_from_monster
get_material_from_monster:
	curl -X POST \
	-H 'Content-Type: application/json' \
	-d '{"monsterId": "1"}' \
	"http://localhost:8080/hunters/1/get_material_from_monster" | jq
