.SHELL := /bin/bash
.DEFAULT_GOAL := help

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: init
init: copy-env ## 開発環境用のDBを作成
	bash hack/docker-run.sh

.PHONY: copy-env
copy-env: ## 各ディレクトリの.envをコピー
	bash ./hack/copy-env.sh

.PHONY: run
run: init server ## migrationとseedを実行してサーバーを起動

.PHONY: server
server: init ## サーバーを起動
	docker-compose up api-math

.PHONY: down
down: ## コンテナを削除
	docker-compose down --remove-orphans
