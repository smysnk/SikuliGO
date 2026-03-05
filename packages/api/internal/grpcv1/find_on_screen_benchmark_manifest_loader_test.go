package grpcv1

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFindBenchScenariosFromManifestExample(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		// Fall back to package-relative when repo root auto-detection fails.
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", manifestPath)

	scenarios, source, err := loadFindBenchScenariosFromManifest(true, false)
	if err != nil {
		t.Fatalf("load manifest scenarios: %v", err)
	}
	if source == "" {
		t.Fatalf("expected resolved manifest path")
	}
	if len(scenarios) == 0 {
		t.Fatalf("expected materialized scenarios")
	}

	// Example manifest runs 10 scenarios across 4 resolutions.
	if got, want := len(scenarios), 40; got != want {
		t.Fatalf("scenario count mismatch: got=%d want=%d", got, want)
	}

	for _, sc := range scenarios {
		if strings.TrimSpace(sc.scenarioTypeID) == "" {
			t.Fatalf("expected scenarioTypeID to be populated scenario=%s", sc.name)
		}
		if sc.screenW <= 0 || sc.screenH <= 0 {
			t.Fatalf("invalid resolution in scenario %+v", sc)
		}
		if sc.size < 16 {
			t.Fatalf("invalid size in scenario %+v", sc)
		}
		if sc.tolerance < 0 || sc.tolerance > 1 {
			t.Fatalf("invalid tolerance in scenario %+v", sc)
		}
		if sc.maxAreaRatio < 1 {
			t.Fatalf("invalid area ratio in scenario %+v", sc)
		}
		if strings.TrimSpace(sc.sourceImagePath) == "" || len(sc.sourceTargets) == 0 {
			t.Fatalf("expected region-spec source data scenario=%s type=%s image=%q targets=%d", sc.name, sc.scenarioTypeID, sc.sourceImagePath, len(sc.sourceTargets))
		}
		if _, err := os.Stat(sc.sourceImagePath); err != nil {
			t.Fatalf("expected source image to exist scenario=%s path=%s err=%v", sc.name, sc.sourceImagePath, err)
		}
		if strings.TrimSpace(sc.benchmarkImagePath) == "" {
			t.Fatalf("expected benchmark image path to be populated scenario=%s", sc.name)
		}
		if _, err := os.Stat(sc.benchmarkImagePath); err != nil {
			t.Fatalf("expected benchmark image to exist scenario=%s path=%s err=%v", sc.name, sc.benchmarkImagePath, err)
		}
		if !strings.EqualFold(sc.kind, "photographic") {
			if !strings.Contains(sc.sourceImagePath, filepath.Join("packages", "api", "internal", "grpcv1", "testdata", "find-bench-assets", "scenario")) {
				t.Fatalf("expected non-photographic scenario to use testdata scenario image path=%s", sc.sourceImagePath)
			}
		}
		if strings.EqualFold(sc.kind, "photographic") {
			if strings.TrimSpace(sc.targetAsset) == "" {
				t.Fatalf("expected photographic scenario to resolve target asset path, scenario=%s", sc.name)
			}
			if _, err := os.Stat(sc.targetAsset); err != nil {
				t.Fatalf("expected photographic target asset to exist path=%s err=%v", sc.targetAsset, err)
			}
		}
	}
}

func TestLoadFindBenchScenariosFromManifestStrictValidationFailure(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	manifestPath := filepath.Join(root, "docs", "bench", "find-on-screen-scenarios.example.json")
	raw, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("read manifest fixture: %v", err)
	}

	var doc map[string]any
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("unmarshal manifest fixture: %v", err)
	}
	doc["unexpected_field"] = true
	mutated, err := json.Marshal(doc)
	if err != nil {
		t.Fatalf("marshal mutated manifest: %v", err)
	}

	tmpDir := t.TempDir()
	badManifest := filepath.Join(tmpDir, "bad-manifest.json")
	if err := os.WriteFile(badManifest, mutated, 0o644); err != nil {
		t.Fatalf("write mutated manifest: %v", err)
	}

	t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", badManifest)
	_, _, err = loadFindBenchScenariosFromManifest(true, false)
	if err == nil {
		t.Fatalf("expected strict schema validation error")
	}
	if !strings.Contains(err.Error(), "strict schema validation") {
		t.Fatalf("expected strict validation error, got: %v", err)
	}
}

func TestFindBenchManifestPreflightFromEnv(t *testing.T) {
	path := strings.TrimSpace(os.Getenv("FIND_BENCH_SCENARIO_MANIFEST"))
	if path == "" {
		t.Skip("FIND_BENCH_SCENARIO_MANIFEST not set")
	}
	scenarios, source, err := loadFindBenchScenariosFromManifest(true, true)
	if err != nil {
		t.Fatalf("manifest preflight failed: %v", err)
	}
	if source == "" {
		t.Fatalf("manifest preflight resolved empty source")
	}
	if len(scenarios) == 0 {
		t.Fatalf("manifest preflight produced zero scenarios")
	}
	schemaPath, err := resolveFindBenchSchemaPath()
	if err != nil {
		t.Fatalf("manifest preflight schema resolve failed: %v", err)
	}
	t.Logf("manifest preflight ok resolved_manifest=%s resolved_schema=%s scenario_count=%d", source, schemaPath, len(scenarios))
}

func TestFindBenchManifestFixtureRegression(t *testing.T) {
	root := findBenchRepoRoot(".")
	if root == "" {
		root = filepath.Clean(filepath.Join("..", "..", ".."))
	}
	fixtureDir := filepath.Join(root, "packages", "api", "internal", "grpcv1", "testdata", "find-bench-manifest")
	cases := []struct {
		file       string
		shouldPass bool
	}{
		{file: "valid-full.json", shouldPass: true},
		{file: "valid-minimal.json", shouldPass: true},
		{file: "invalid-unknown-field.json", shouldPass: false},
		{file: "invalid-schema-major.json", shouldPass: false},
		{file: "invalid-missing-required.json", shouldPass: false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.file, func(t *testing.T) {
			t.Setenv("FIND_BENCH_SCENARIO_MANIFEST", filepath.Join(fixtureDir, tc.file))
			scenarios, source, err := loadFindBenchScenariosFromManifest(true, true)
			if tc.shouldPass {
				if err != nil {
					t.Fatalf("expected success, got error: %v", err)
				}
				if source == "" || len(scenarios) == 0 {
					t.Fatalf("expected non-empty source/scenarios source=%q count=%d", source, len(scenarios))
				}
				return
			}
			if err == nil {
				t.Fatalf("expected failure, got success count=%d source=%s", len(scenarios), source)
			}
			t.Logf("expected failure: %v", err)
		})
	}
}

func TestResolveFindBenchManifestPathExampleAlias(t *testing.T) {
	path, err := resolveFindBenchManifestPath("example")
	if err != nil {
		t.Fatalf("resolve example alias: %v", err)
	}
	if !strings.HasSuffix(path, filepath.Join("docs", "bench", "find-on-screen-scenarios.example.json")) {
		t.Fatalf("unexpected resolved path: %s", path)
	}
}

func TestValidateFindBenchSchemaVersion(t *testing.T) {
	for _, tc := range []struct {
		version   string
		shouldErr bool
	}{
		{version: "1.0.0", shouldErr: false},
		{version: "1.8.2", shouldErr: false},
		{version: "2.0.0", shouldErr: true},
		{version: "0.9.0", shouldErr: true},
		{version: "broken", shouldErr: true},
	} {
		tc := tc
		t.Run(fmt.Sprintf("v_%s", strings.ReplaceAll(tc.version, ".", "_")), func(t *testing.T) {
			err := validateFindBenchSchemaVersion(tc.version)
			if tc.shouldErr && err == nil {
				t.Fatalf("expected error for version=%s", tc.version)
			}
			if !tc.shouldErr && err != nil {
				t.Fatalf("unexpected error for version=%s: %v", tc.version, err)
			}
		})
	}
}
