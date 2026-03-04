// Package sikuli provides the compatibility-facing automation API used by SikuliGO.
//
// The surface is intentionally aligned with common SikuliX concepts so existing
// script flows can migrate with minimal rewriting:
//   - Pattern and similarity tuning
//   - Region scoped search and wait semantics
//   - Screen level orchestration
//   - Input control (click, type, hotkey)
//   - OCR and observe events
//
// Java SikuliX and SikuliGO are not byte-for-byte identical, but the exported
// contracts in this package are designed to preserve the same mental model.
package sikuli
