#!/usr/bin/env bash

# Shared macOS OCR/OpenCV CGO environment setup for local client/build scripts.
# Source this file, then call `configure_macos_ocr_env`.

append_flag_if_missing() {
  local current="$1"
  local flag="$2"
  if [[ -z "${current}" ]]; then
    printf '%s' "$flag"
    return
  fi
  if [[ " ${current} " == *" ${flag} "* ]]; then
    printf '%s' "$current"
    return
  fi
  printf '%s %s' "$current" "$flag"
}

normalize_flag_var() {
  local name="$1"
  local value="${!name:-}"
  if [[ -z "${value//[[:space:]]/}" ]]; then
    export "${name}="
    return
  fi
  export "${name}=$(echo "${value}" | xargs 2>/dev/null || true)"
}

configure_macos_ocr_env() {
  if [[ "$(uname -s)" != "Darwin" ]]; then
    return 0
  fi
  if ! command -v brew >/dev/null 2>&1; then
    return 0
  fi

  local homebrew_prefix lept_prefix tess_prefix
  homebrew_prefix="$(brew --prefix 2>/dev/null || true)"
  lept_prefix="$(brew --prefix leptonica 2>/dev/null || true)"
  tess_prefix="$(brew --prefix tesseract 2>/dev/null || true)"

  if [[ -n "${homebrew_prefix}" ]]; then
    export PKG_CONFIG_PATH="${homebrew_prefix}/lib/pkgconfig${PKG_CONFIG_PATH:+:${PKG_CONFIG_PATH}}"
    export LIBRARY_PATH="${homebrew_prefix}/lib${LIBRARY_PATH:+:${LIBRARY_PATH}}"
    export CGO_CFLAGS="$(append_flag_if_missing "${CGO_CFLAGS:-}" "-I${homebrew_prefix}/include")"
    export CGO_CPPFLAGS="$(append_flag_if_missing "${CGO_CPPFLAGS:-}" "-I${homebrew_prefix}/include")"
    export CGO_CXXFLAGS="$(append_flag_if_missing "${CGO_CXXFLAGS:-}" "-I${homebrew_prefix}/include")"
    export CGO_LDFLAGS="$(append_flag_if_missing "${CGO_LDFLAGS:-}" "-L${homebrew_prefix}/lib")"
  fi

  if [[ -n "${lept_prefix}" ]]; then
    export CGO_CFLAGS="$(append_flag_if_missing "${CGO_CFLAGS:-}" "-I${lept_prefix}/include")"
    export CGO_CPPFLAGS="$(append_flag_if_missing "${CGO_CPPFLAGS:-}" "-I${lept_prefix}/include")"
    export CGO_CXXFLAGS="$(append_flag_if_missing "${CGO_CXXFLAGS:-}" "-I${lept_prefix}/include")"
    export CGO_LDFLAGS="$(append_flag_if_missing "${CGO_LDFLAGS:-}" "-L${lept_prefix}/lib")"
  fi

  if [[ -n "${tess_prefix}" ]]; then
    export CGO_CFLAGS="$(append_flag_if_missing "${CGO_CFLAGS:-}" "-I${tess_prefix}/include")"
    export CGO_CPPFLAGS="$(append_flag_if_missing "${CGO_CPPFLAGS:-}" "-I${tess_prefix}/include")"
    export CGO_CXXFLAGS="$(append_flag_if_missing "${CGO_CXXFLAGS:-}" "-I${tess_prefix}/include")"
    export CGO_LDFLAGS="$(append_flag_if_missing "${CGO_LDFLAGS:-}" "-L${tess_prefix}/lib")"
  fi

  export CGO_LDFLAGS="$(append_flag_if_missing "${CGO_LDFLAGS:-}" "-llept")"
  export CGO_LDFLAGS="$(append_flag_if_missing "${CGO_LDFLAGS:-}" "-ltesseract")"

  normalize_flag_var PKG_CONFIG_PATH
  normalize_flag_var LIBRARY_PATH
  normalize_flag_var CGO_CFLAGS
  normalize_flag_var CGO_CPPFLAGS
  normalize_flag_var CGO_CXXFLAGS
  normalize_flag_var CGO_LDFLAGS
}
