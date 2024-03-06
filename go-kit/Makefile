.SHELL := /bin/bash
.DEFAULT_GOAL := help

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: init
init: copy-env ## 開発環境用のDBを作成
	bash hack/docker-run.sh

.PHONY: init-test
init-test: copy-env setup-cognito-test ## テスト環境用のDB作成とCognitoセットアップ
	bash hack/docker-run-test.sh

.PHONY: copy-env
copy-env: ## 各ディレクトリの.envをコピー
	bash ./hack/copy-env.sh

.PHONY: run
run: init migrate seed server ## migrationとseedを実行してサーバーを起動

.PHONY: server
server: init migrate setup-cognito ## サーバーを起動
	docker-compose up api

.PHONY: migrate
migrate: init ## DBのmigrationを実行
	docker-compose run --rm migration

.PHONY: down
down: ## コンテナを削除
	docker-compose down --remove-orphans

.PHONY: gen-entity
gen-entity: init migrate ## DBスキーマからentityを生成
	docker-compose run --rm go-cmd ./cmd/gorm-gen

.PHONY: seed
seed: init migrate ## DBにseedデータを投入
	docker-compose run --rm go-cmd ./cmd/db/seed/dev

.PHONY: seed-stg
seed-stg: init migrate ## DBにseedデータを投入
	docker-compose run --rm go-cmd ./cmd/db/seed/stg

.PHONY: setup-cognito
setup-cognito: init ## Cognitoのセットアップ
	docker-compose run --rm go-cmd ./cmd/setup-cognito .env

.PHONY: send-push-notification
send-push-notification: ## プッシュ通知を送信
	docker-compose run --rm go-cmd ./service/exercise/internal/cmd/send-push-notification/dev

.PHONY: test
test: init-test ## テストを実行
	docker-compose run --rm go-test go test -v ./...

.PHONY: test-func
test-func: func = ## 指定した関数のみテストを実行
test-func: init-test
	docker-compose run --rm go-test go test -v -run $(func) ./...

.PHONY: setup-cognito-test
setup-cognito-test: init ## テスト用のCognitoのセットアップ
	docker-compose run --rm go-cmd ./cmd/setup-cognito .env.test
