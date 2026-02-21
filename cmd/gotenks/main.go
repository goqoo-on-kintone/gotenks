package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goqoo-on-kintone/gotenks/internal/generator"
	"github.com/goqoo-on-kintone/gotenks/internal/parser"
)

func main() {
	// コマンドライン引数の定義
	var input, output, pkg, prefix string

	flag.StringVar(&input, "input", "", "Input .d.ts file or directory")
	flag.StringVar(&input, "i", "", "Input .d.ts file or directory (shorthand)")
	flag.StringVar(&output, "output", "", "Output .go file or directory (default: stdout)")
	flag.StringVar(&output, "o", "", "Output .go file or directory (shorthand)")
	flag.StringVar(&pkg, "package", "kintone", "Package name for generated Go code")
	flag.StringVar(&prefix, "prefix", "K", "Prefix for field names (must start with uppercase letter)")
	flag.Parse()

	if input == "" {
		fmt.Fprintln(os.Stderr, "Error: -input is required")
		flag.Usage()
		os.Exit(1)
	}

	// プレフィックスの検証
	if err := generator.ValidatePrefix(prefix); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// 入力がファイルかディレクトリかを判定
	info, err := os.Stat(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if info.IsDir() {
		// ディレクトリの場合、全 .d.ts ファイルを処理
		err = processDirectory(input, output, pkg, prefix)
	} else {
		// 単一ファイルの場合
		err = processFile(input, output, pkg, prefix)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// processDirectory はディレクトリ内の全 .d.ts ファイルを処理する
func processDirectory(inputDir, outputDir, pkg, prefix string) error {
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".d.ts") {
			continue
		}

		inputPath := filepath.Join(inputDir, entry.Name())

		var outputPath string
		if outputDir != "" {
			// 出力ディレクトリが指定されている場合
			baseName := strings.TrimSuffix(entry.Name(), ".d.ts")
			outputPath = filepath.Join(outputDir, baseName+".go")
		}

		if err := processFile(inputPath, outputPath, pkg, prefix); err != nil {
			return fmt.Errorf("processing %s: %w", entry.Name(), err)
		}
	}

	return nil
}

// processFile は単一の .d.ts ファイルを処理する
func processFile(inputPath, outputPath, pkg, prefix string) error {
	// パース
	result, err := parser.ParseFile(inputPath)
	if err != nil {
		return fmt.Errorf("parsing %s: %w", inputPath, err)
	}

	// コード生成
	config := generator.Config{
		PackageName: pkg,
		Prefix:      prefix,
	}
	code := generator.Generate(result, config)

	// 出力
	if outputPath == "" {
		// 標準出力
		fmt.Print(code)
	} else {
		// ファイル出力
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(outputPath, []byte(code), 0644); err != nil {
			return err
		}
		fmt.Printf("Generated: %s\n", outputPath)
	}

	return nil
}
