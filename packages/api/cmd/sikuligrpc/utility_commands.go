package main

import (
	"bufio"
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	iofs "io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

//go:embed examples/node examples/python
var embeddedExamples embed.FS

func maybeRunUtilityCommands(args []string) (bool, error) {
	if len(args) == 0 {
		return false, nil
	}
	switch args[0] {
	case "help", "-h", "--help":
		printCommandHelp(os.Stdout)
		return true, nil
	case "doctor":
		return true, runDoctor(args[1:])
	case "install-binary":
		return true, runInstallBinary(args[1:])
	case "init:js-examples":
		return true, runInitJSExamples(args[1:])
	case "init:py-examples":
		return true, runInitPyExamples(args[1:])
	default:
		return false, nil
	}
}

func printCommandHelp(w io.Writer) {
	_, _ = fmt.Fprintln(w, "Usage:")
	_, _ = fmt.Fprintln(w, "  sikuligo [flags]")
	_, _ = fmt.Fprintln(w, "  sikuligo init:js-examples [--dir <targetDir>] [--skip-install]")
	_, _ = fmt.Fprintln(w, "  sikuligo init:py-examples [--dir <targetDir>] [--skip-install]")
	_, _ = fmt.Fprintln(w, "  sikuligo install-binary [--dir <binDir>] [--yes] [--no-shell-update]")
	_, _ = fmt.Fprintln(w, "  sikuligo doctor")
	_, _ = fmt.Fprintln(w, "")
	_, _ = fmt.Fprintln(w, "Server Flags:")
	_, _ = fmt.Fprintln(w, "  -listen string")
	_, _ = fmt.Fprintln(w, "        gRPC listen address (default \":50051\")")
	_, _ = fmt.Fprintln(w, "  -admin-listen string")
	_, _ = fmt.Fprintln(w, "        admin HTTP listen address for health/metrics/dashboard; empty disables admin server (default \":8080\")")
	_, _ = fmt.Fprintln(w, "  -sqlite-path string")
	_, _ = fmt.Fprintln(w, "        sqlite datastore path for API sessions, client sessions, and interactions (default \"sikuligo.db\")")
	_, _ = fmt.Fprintln(w, "  -auth-token string")
	_, _ = fmt.Fprintln(w, "        shared API token; accepted via metadata x-api-key or Authorization: Bearer <token>")
	_, _ = fmt.Fprintln(w, "  -enable-reflection")
	_, _ = fmt.Fprintln(w, "        enable gRPC reflection (default true)")
}

func runDoctor(args []string) error {
	fs := flag.NewFlagSet("doctor", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(fs.Args(), " "))
	}
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("resolve executable: %w", err)
	}
	exeReal, err := filepath.EvalSymlinks(exe)
	if err != nil {
		exeReal = exe
	}
	fmt.Println("sikuligo doctor: ok")
	fmt.Printf("binary: %s\n", exeReal)
	fmt.Printf("platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	depsMissing := doctorMissingRuntimeDeps()
	if len(depsMissing) == 0 {
		fmt.Println("runtime deps: ok")
		return nil
	}
	for _, dep := range depsMissing {
		fmt.Printf("runtime deps: missing %s\n", dep)
	}
	return fmt.Errorf("missing runtime dependencies: %s", strings.Join(depsMissing, ", "))
}

func doctorMissingRuntimeDeps() []string {
	missing := make([]string, 0, 2)
	switch runtime.GOOS {
	case "darwin":
		if _, err := exec.LookPath("cliclick"); err != nil {
			missing = append(missing, `cliclick (install with "brew install cliclick")`)
		}
	case "linux":
		if _, err := exec.LookPath("xdotool"); err != nil {
			missing = append(missing, "xdotool")
		}
	case "windows":
		if _, err := exec.LookPath("powershell"); err != nil {
			missing = append(missing, "powershell")
		}
	}
	return missing
}

func runInstallBinary(args []string) error {
	fs := flag.NewFlagSet("install-binary", flag.ContinueOnError)
	targetDir := fs.String("dir", defaultInstallDir(), "target directory")
	yes := fs.Bool("yes", false, "auto-approve shell profile PATH update prompt")
	noShellUpdate := fs.Bool("no-shell-update", false, "skip adding target dir to shell profile")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: sikuligo install-binary [--dir <binDir>] [--yes] [--no-shell-update]\n")
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(fs.Args(), " "))
	}

	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("resolve executable: %w", err)
	}
	exeReal, err := filepath.EvalSymlinks(exe)
	if err != nil {
		exeReal = exe
	}
	runtimes := discoverRuntimeSources(exeReal)
	if len(runtimes) == 0 {
		return fmt.Errorf("no runtimes found near executable: %s", exeReal)
	}

	if err := os.MkdirAll(*targetDir, 0o755); err != nil {
		return fmt.Errorf("create target dir: %w", err)
	}
	var copied []string
	for _, src := range runtimes {
		base := filepath.Base(src)
		targets := map[string]struct{}{base: {}}
		if strings.HasPrefix(strings.ToLower(base), "sikuligrpc") {
			targets[strings.Replace(base, "sikuligrpc", "sikuligo", 1)] = struct{}{}
		}
		for targetBase := range targets {
			dst := filepath.Join(*targetDir, targetBase)
			if err := copyFile(src, dst); err != nil {
				return fmt.Errorf("copy %s -> %s: %w", src, dst, err)
			}
			if runtime.GOOS != "windows" {
				if err := os.Chmod(dst, 0o755); err != nil {
					return fmt.Errorf("chmod %s: %w", dst, err)
				}
			}
			copied = append(copied, dst)
		}
	}
	for _, dst := range copied {
		fmt.Println(dst)
	}

	if !*noShellUpdate {
		if profile, sourceCmd, ok := detectShellProfile(); ok {
			shouldUpdate := *yes
			if !shouldUpdate {
				answer, err := promptYesNo(fmt.Sprintf("Add %s to PATH in %s?", *targetDir, profile))
				if err != nil {
					return err
				}
				shouldUpdate = answer
			}
			if shouldUpdate {
				updated, err := ensurePathExport(profile, *targetDir)
				if err != nil {
					return err
				}
				if updated {
					fmt.Printf("Run %s to reload PATH in this shell.\n", sourceCmd)
				} else {
					fmt.Printf("PATH already configured in %s.\n", profile)
				}
			} else {
				fmt.Printf("Ensure %s is on PATH for new shells.\n", *targetDir)
			}
		}
	}
	return nil
}

func runInitJSExamples(args []string) error {
	fs := flag.NewFlagSet("init:js-examples", flag.ContinueOnError)
	targetDir := fs.String("dir", "", "project directory")
	skipInstall := fs.Bool("skip-install", false, "skip running yarn install")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: sikuligo init:js-examples [--dir <targetDir>] [--skip-install]\n")
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(fs.Args(), " "))
	}

	projectDir := strings.TrimSpace(*targetDir)
	if projectDir == "" {
		value, err := promptWithDefault("Project directory name", "sikuligo-demo")
		if err != nil {
			return err
		}
		projectDir = value
	}
	if !filepath.IsAbs(projectDir) {
		projectDir = mustAbs(projectDir)
	}
	if err := os.MkdirAll(projectDir, 0o755); err != nil {
		return fmt.Errorf("create project directory: %w", err)
	}
	if err := ensureNodePackageJSON(projectDir); err != nil {
		return err
	}
	if !*skipInstall {
		cmd := exec.Command("yarn", "install")
		cmd.Dir = projectDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("yarn install failed: %w", err)
		}
	}
	if err := runInitExamples([]string{"--lang", "node", "--dir", projectDir, "--force"}); err != nil {
		return err
	}
	fmt.Printf("Initialized SikuliGO project in: %s\n", projectDir)
	fmt.Printf("Examples copied to: %s\n", filepath.Join(projectDir, "examples"))
	return nil
}

func runInitPyExamples(args []string) error {
	fs := flag.NewFlagSet("init:py-examples", flag.ContinueOnError)
	targetDir := fs.String("dir", "", "project directory")
	skipInstall := fs.Bool("skip-install", false, "skip setting up .venv and pip install")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: sikuligo init:py-examples [--dir <targetDir>] [--skip-install]\n")
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(fs.Args(), " "))
	}

	projectDir := strings.TrimSpace(*targetDir)
	if projectDir == "" {
		value, err := promptWithDefault("Project directory name", "sikuligo-demo")
		if err != nil {
			return err
		}
		projectDir = value
	}
	if !filepath.IsAbs(projectDir) {
		projectDir = mustAbs(projectDir)
	}
	if err := os.MkdirAll(projectDir, 0o755); err != nil {
		return fmt.Errorf("create project directory: %w", err)
	}
	if err := ensurePythonRequirements(projectDir); err != nil {
		return err
	}
	if !*skipInstall {
		if err := setupPythonEnvironment(projectDir); err != nil {
			return err
		}
	}
	if err := runInitExamples([]string{"--lang", "python", "--dir", projectDir, "--force"}); err != nil {
		return err
	}
	fmt.Printf("Initialized SikuliGO project in: %s\n", projectDir)
	fmt.Printf("Examples copied to: %s\n", filepath.Join(projectDir, "examples"))
	return nil
}

func runInitExamples(args []string) error {
	fs := flag.NewFlagSet("init-examples", flag.ContinueOnError)
	lang := fs.String("lang", "node", "example language: node|javascript|python")
	targetDir := fs.String("dir", ".", "target directory")
	force := fs.Bool("force", false, "overwrite existing examples directory")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(fs.Args(), " "))
	}

	normalizedLang, sourceRoot, err := resolveExampleLanguage(*lang)
	if err != nil {
		return err
	}
	outputExamplesDir := filepath.Join(*targetDir, "examples")
	if stat, err := os.Stat(outputExamplesDir); err == nil && stat.IsDir() {
		if !*force {
			return fmt.Errorf("target already exists: %s (use --force to overwrite)", outputExamplesDir)
		}
		if err := os.RemoveAll(outputExamplesDir); err != nil {
			return fmt.Errorf("remove existing target: %w", err)
		}
	}
	if err := copyEmbeddedExamples(sourceRoot, outputExamplesDir); err != nil {
		return err
	}

	fmt.Printf("Initialized %s examples in: %s\n", normalizedLang, outputExamplesDir)
	return nil
}

func resolveExampleLanguage(lang string) (string, string, error) {
	switch strings.ToLower(strings.TrimSpace(lang)) {
	case "node", "javascript", "js":
		return "node", "examples/node", nil
	case "python", "py":
		return "python", "examples/python", nil
	default:
		return "", "", fmt.Errorf("unsupported language %q (expected node|python)", lang)
	}
}

func copyEmbeddedExamples(sourceRoot, outputDir string) error {
	return iofs.WalkDir(embeddedExamples, sourceRoot, func(entryPath string, d iofs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel := strings.TrimPrefix(entryPath, sourceRoot)
		rel = strings.TrimPrefix(rel, "/")
		if rel == "" {
			return nil
		}
		target := filepath.Join(outputDir, filepath.FromSlash(rel))
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		content, err := embeddedExamples.ReadFile(entryPath)
		if err != nil {
			return fmt.Errorf("read embedded example %s: %w", entryPath, err)
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return fmt.Errorf("create output dir for %s: %w", target, err)
		}
		if err := os.WriteFile(target, content, 0o644); err != nil {
			return fmt.Errorf("write output file %s: %w", target, err)
		}
		return nil
	})
}

func ensureNodePackageJSON(projectDir string) error {
	pkgPath := filepath.Join(projectDir, "package.json")
	pkg := map[string]any{
		"name":    filepath.Base(projectDir),
		"private": true,
		"type":    "module",
	}
	if strings.TrimSpace(pkg["name"].(string)) == "" {
		pkg["name"] = "sikuligo-project"
	}
	if b, err := os.ReadFile(pkgPath); err == nil {
		var existing map[string]any
		if err := json.Unmarshal(b, &existing); err == nil {
			pkg = existing
		}
	}
	if _, ok := pkg["name"]; !ok {
		pkg["name"] = filepath.Base(projectDir)
	}
	if _, ok := pkg["private"]; !ok {
		pkg["private"] = true
	}
	pkg["type"] = "module"
	deps, _ := pkg["dependencies"].(map[string]any)
	if deps == nil {
		deps = map[string]any{}
	}
	version := strings.TrimSpace(os.Getenv("SIKULIGO_NODE_PACKAGE_VERSION"))
	if version == "" {
		version = "latest"
	}
	deps["@sikuligo/sikuligo"] = version
	pkg["dependencies"] = deps
	out, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return fmt.Errorf("encode package.json: %w", err)
	}
	out = append(out, '\n')
	if err := os.WriteFile(pkgPath, out, 0o644); err != nil {
		return fmt.Errorf("write package.json: %w", err)
	}
	return nil
}

func ensurePythonRequirements(projectDir string) error {
	requirementsPath := filepath.Join(projectDir, "requirements.txt")
	packageVersion := strings.TrimSpace(os.Getenv("SIKULIGO_PY_PACKAGE_VERSION"))
	requirement := "sikuligo"
	if packageVersion != "" && packageVersion != "latest" {
		requirement = fmt.Sprintf("sikuligo==%s", packageVersion)
	}
	content := requirement + "\n"
	if err := os.WriteFile(requirementsPath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write requirements.txt: %w", err)
	}
	return nil
}

func setupPythonEnvironment(projectDir string) error {
	pythonBin, err := resolvePythonBinary()
	if err != nil {
		return err
	}

	venvCmd := exec.Command(pythonBin, "-m", "venv", ".venv")
	venvCmd.Dir = projectDir
	venvCmd.Stdout = os.Stdout
	venvCmd.Stderr = os.Stderr
	venvCmd.Stdin = os.Stdin
	if err := venvCmd.Run(); err != nil {
		return fmt.Errorf("python venv setup failed: %w", err)
	}

	venvPython := filepath.Join(projectDir, ".venv", "bin", "python")
	if runtime.GOOS == "windows" {
		venvPython = filepath.Join(projectDir, ".venv", "Scripts", "python.exe")
	}
	pipCmd := exec.Command(venvPython, "-m", "pip", "install", "-r", "requirements.txt")
	pipCmd.Dir = projectDir
	pipCmd.Stdout = os.Stdout
	pipCmd.Stderr = os.Stderr
	pipCmd.Stdin = os.Stdin
	if err := pipCmd.Run(); err != nil {
		return fmt.Errorf("pip install failed: %w", err)
	}
	return nil
}

func resolvePythonBinary() (string, error) {
	candidates := []string{"python3", "python"}
	for _, candidate := range candidates {
		path, err := exec.LookPath(candidate)
		if err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("python not found in PATH (tried: %s)", strings.Join(candidates, ", "))
}

func promptWithDefault(label, fallback string) (string, error) {
	fmt.Printf("%s [%s]: ", label, fallback)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			if strings.TrimSpace(line) == "" {
				return fallback, nil
			}
			return strings.TrimSpace(line), nil
		}
		return "", err
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return fallback, nil
	}
	return line, nil
}

func promptYesNo(question string) (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return false, nil
	}
	fmt.Printf("%s [y/N]: ", question)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}
	switch strings.ToLower(strings.TrimSpace(line)) {
	case "y", "yes":
		return true, nil
	default:
		return false, nil
	}
}

func detectShellProfile() (profile string, sourceCmd string, ok bool) {
	shell := os.Getenv("SHELL")
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", false
	}
	switch {
	case strings.Contains(shell, "zsh"):
		return filepath.Join(home, ".zshrc"), "source ~/.zshrc", true
	case strings.Contains(shell, "bash"):
		return filepath.Join(home, ".bash_profile"), "source ~/.bash_profile", true
	default:
		return "", "", false
	}
}

func ensurePathExport(profilePath, binDir string) (bool, error) {
	exportLine := fmt.Sprintf("export PATH=\"%s:$PATH\"", binDir)
	marker := "# Added by sikuligo install-binary"
	snippet := marker + "\n" + exportLine + "\n"
	existing, _ := os.ReadFile(profilePath)
	content := string(existing)
	if strings.Contains(content, exportLine) {
		return false, nil
	}
	if content != "" && !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	content += snippet
	if err := os.WriteFile(profilePath, []byte(content), 0o644); err != nil {
		return false, fmt.Errorf("update shell profile %s: %w", profilePath, err)
	}
	return true, nil
}

func defaultInstallDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return filepath.Join(home, ".local", "bin")
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = out.Close() }()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}

var runtimeNamePattern = regexp.MustCompile(`(?i)^sikuli.*(\.exe)?$`)

func discoverRuntimeSources(primary string) []string {
	candidates := map[string]struct{}{primary: {}}
	dir := filepath.Dir(primary)
	entries, err := os.ReadDir(dir)
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if !runtimeNamePattern.MatchString(name) {
				continue
			}
			candidates[filepath.Join(dir, name)] = struct{}{}
		}
	}
	out := make([]string, 0, len(candidates))
	for candidate := range candidates {
		if stat, err := os.Stat(candidate); err == nil && !stat.IsDir() {
			out = append(out, candidate)
		}
	}
	return out
}

func mustAbs(p string) string {
	abs, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	return abs
}
