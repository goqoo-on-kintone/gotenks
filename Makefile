.PHONY: build test run clean lint

# ビルド
build:
	go build -o bin/gotenks ./cmd/gotenks

# テスト実行
test:
	go test -v ./...

# 開発用: ビルドして実行
run: build
	./bin/gotenks

# テストデータで実行
run-test: build
	./bin/gotenks -input ./testdata -output ./examples/output

# クリーンアップ
clean:
	rm -rf bin/
	rm -rf examples/output/*.go

# Lint
lint:
	go vet ./...
	@if command -v staticcheck > /dev/null; then staticcheck ./...; fi

# フォーマット
fmt:
	go fmt ./...

# 依存関係の整理
tidy:
	go mod tidy
