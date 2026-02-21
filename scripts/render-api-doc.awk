function trim(s) {
  sub(/^[[:space:]]+/, "", s)
  sub(/[[:space:]]+$/, "", s)
  return s
}

function lcase(s,    out, i, c, pos) {
  out = s
  for (i = 1; i <= length(out); i++) {
    c = substr(out, i, 1)
    pos = index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", c)
    if (pos > 0) {
      out = substr(out, 1, i - 1) substr("abcdefghijklmnopqrstuvwxyz", pos, 1) substr(out, i + 1)
    }
  }
  return out
}

function anchor(symbol_kind, symbol_name, recv_name) {
  if (symbol_kind == "type") {
    return "type-" lcase(symbol_name)
  }
  if (symbol_kind == "func") {
    return "func-" lcase(symbol_name)
  }
  return "method-" lcase(recv_name) "-" lcase(symbol_name)
}

function has_word(sig, word,    pat) {
  pat = "(^|[^A-Za-z0-9_])" word "([^A-Za-z0-9_]|$)"
  return sig ~ pat
}

function add_use(sig, self_type,    t, out, sep) {
  out = ""
  sep = ""
  for (t in type_map) {
    if (t == self_type) {
      continue
    }
    if (has_word(sig, t)) {
      out = out sep "[`" t "`](#type-" lcase(t) ")"
      sep = ", "
    }
  }
  return out
}

{
  n++
  line[n] = $0
  if ($0 ~ /^type[[:space:]]+[A-Za-z_][A-Za-z0-9_]*/) {
    split($0, parts, /[[:space:]]+/)
    t = parts[2]
    type_map[t] = 1
  }
}

END {
  print "# API: `" rel "`"
  print ""
  print "[Back to API Index](./)"
  print ""
  print "<style>"
  print "  .api-type { color: #0f766e; font-weight: 700; }"
  print "  .api-func { color: #1d4ed8; font-weight: 700; }"
  print "  .api-method { color: #7c3aed; font-weight: 700; }"
  print "  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, \"Liberation Mono\", monospace; }"
  print "</style>"
  print ""
  print "Legend: <span class=\"api-type\">Type</span>, <span class=\"api-func\">Function</span>, <span class=\"api-method\">Method</span>"
  print ""

  if (n > 0) {
    print "Package: `" trim(line[1]) "`"
    print ""
  }

  print "## Symbol Index"
  print ""

  type_count = 0
  func_count = 0
  method_count = 0
  for (i = 1; i <= n; i++) {
    s = line[i]
    if (s ~ /^type[[:space:]]+[A-Za-z_][A-Za-z0-9_]*/) {
      split(s, parts, /[[:space:]]+/)
      name = parts[2]
      type_count++
      type_name[type_count] = name
      type_sig[name] = s
      desc = ""
      j = i + 1
      while (j <= n && line[j] ~ /^[[:space:]]{4}[^[:space:]].*$/) {
        piece = trim(line[j])
        if (desc == "") {
          desc = piece
        } else {
          desc = desc " " piece
        }
        j++
      }
      type_desc[name] = desc
      continue
    }

    if (s ~ /^func[[:space:]]+\([^)]*\)[[:space:]]+[A-Za-z_][A-Za-z0-9_]*\(/) {
      method_count++
      symbol = s
      sub(/^func[[:space:]]+\(/, "", symbol)
      split(symbol, recv_parts, /\)[[:space:]]+/)
      receiver = recv_parts[1]
      gsub(/^\*/, "", receiver)
      split(receiver, recv_tokens, /[[:space:]]+/)
      receiver = recv_tokens[length(recv_tokens)]
      gsub(/^\*/, "", receiver)
      split(recv_parts[2], meth_parts, /\(/)
      meth = meth_parts[1]
      method_recv[method_count] = receiver
      method_name[method_count] = meth
      method_sig[method_count] = s
      desc = ""
      j = i + 1
      while (j <= n && line[j] ~ /^[[:space:]]{4}[^[:space:]].*$/) {
        piece = trim(line[j])
        if (desc == "") {
          desc = piece
        } else {
          desc = desc " " piece
        }
        j++
      }
      method_desc[method_count] = desc
      continue
    }

    if (s ~ /^func[[:space:]]+[A-Za-z_][A-Za-z0-9_]*\(/) {
      func_count++
      split(s, func_parts, /\(/)
      split(func_parts[1], name_parts, /[[:space:]]+/)
      fname = name_parts[2]
      func_name[func_count] = fname
      func_sig[func_count] = s
      desc = ""
      j = i + 1
      while (j <= n && line[j] ~ /^[[:space:]]{4}[^[:space:]].*$/) {
        piece = trim(line[j])
        if (desc == "") {
          desc = piece
        } else {
          desc = desc " " piece
        }
        j++
      }
      func_desc[func_count] = desc
    }
  }

  print "### Types"
  print ""
  if (type_count == 0) {
    print "- none"
  } else {
    for (i = 1; i <= type_count; i++) {
      t = type_name[i]
      a = anchor("type", t, "")
      printf("- <span class=\"api-type\">[`%s`](#%s)</span>\n", t, a)
    }
  }
  print ""

  print "### Functions"
  print ""
  if (func_count == 0) {
    print "- none"
  } else {
    for (i = 1; i <= func_count; i++) {
      f = func_name[i]
      a = anchor("func", f, "")
      printf("- <span class=\"api-func\">[`%s`](#%s)</span>\n", f, a)
    }
  }
  print ""

  print "### Methods"
  print ""
  if (method_count == 0) {
    print "- none"
  } else {
    for (i = 1; i <= method_count; i++) {
      m = method_name[i]
      r = method_recv[i]
      a = anchor("method", m, r)
      printf("- <span class=\"api-method\">[`%s.%s`](#%s)</span>\n", r, m, a)
    }
  }
  print ""

  print "## Declarations"
  print ""

  if (type_count > 0) {
    print "### Types"
    print ""
    for (i = 1; i <= type_count; i++) {
      t = type_name[i]
      a = anchor("type", t, "")
      printf("#### <a id=\"%s\"></a><span class=\"api-type\">Type</span> `%s`\n\n", a, t)
      printf("- Signature: <span class=\"api-signature\">`%s`</span>\n", type_sig[t])
      uses = add_use(type_sig[t], t)
      if (uses != "") {
        printf("- Uses: %s\n", uses)
      }
      if (type_desc[t] != "") {
        printf("- Notes: %s\n", type_desc[t])
      }
      print ""
    }
  }

  if (func_count > 0) {
    print "### Functions"
    print ""
    for (i = 1; i <= func_count; i++) {
      f = func_name[i]
      a = anchor("func", f, "")
      printf("#### <a id=\"%s\"></a><span class=\"api-func\">Function</span> `%s`\n\n", a, f)
      printf("- Signature: <span class=\"api-signature\">`%s`</span>\n", func_sig[i])
      uses = add_use(func_sig[i], "")
      if (uses != "") {
        printf("- Uses: %s\n", uses)
      }
      if (func_desc[i] != "") {
        printf("- Notes: %s\n", func_desc[i])
      }
      print ""
    }
  }

  if (method_count > 0) {
    print "### Methods"
    print ""
    for (i = 1; i <= method_count; i++) {
      m = method_name[i]
      r = method_recv[i]
      a = anchor("method", m, r)
      printf("#### <a id=\"%s\"></a><span class=\"api-method\">Method</span> `%s.%s`\n\n", a, r, m)
      printf("- Signature: <span class=\"api-signature\">`%s`</span>\n", method_sig[i])
      uses = add_use(method_sig[i], r)
      if (uses != "") {
        printf("- Uses: %s\n", uses)
      }
      if (method_desc[i] != "") {
        printf("- Notes: %s\n", method_desc[i])
      }
      print ""
    }
  }

  print "## Raw Package Doc"
  print ""
  print "```text"
  for (i = 1; i <= n; i++) {
    print line[i]
  }
  print "```"
}
