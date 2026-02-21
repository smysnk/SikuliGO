local M = {}
local Client = {}
Client.__index = Client

local function file_exists(path)
  local f = io.open(path, "r")
  if f ~= nil then
    f:close()
    return true
  end
  return false
end

local function shell_quote(value)
  return "'" .. tostring(value):gsub("'", "'\"'\"'") .. "'"
end

local function json_escape(value)
  local map = {
    ['"'] = '\\"',
    ["\\"] = "\\\\",
    ["\b"] = "\\b",
    ["\f"] = "\\f",
    ["\n"] = "\\n",
    ["\r"] = "\\r",
    ["\t"] = "\\t"
  }
  return value:gsub('[%z\1-\31\\"]', function(ch)
    return map[ch] or string.format("\\u%04x", string.byte(ch))
  end)
end

local function is_array(t)
  local count = 0
  local max = 0
  for key, _ in pairs(t) do
    if type(key) ~= "number" or key <= 0 or key % 1 ~= 0 then
      return false
    end
    count = count + 1
    if key > max then
      max = key
    end
  end
  return count == max
end

local function encode_json(value)
  local value_type = type(value)
  if value_type == "nil" then
    return "null"
  end
  if value_type == "boolean" then
    return value and "true" or "false"
  end
  if value_type == "number" then
    if value ~= value or value == math.huge or value == -math.huge then
      error("invalid number value for json encoding")
    end
    return tostring(value)
  end
  if value_type == "string" then
    return '"' .. json_escape(value) .. '"'
  end
  if value_type == "table" then
    if is_array(value) then
      local out = {}
      for i = 1, #value do
        out[#out + 1] = encode_json(value[i])
      end
      return "[" .. table.concat(out, ",") .. "]"
    end
    local keys = {}
    for key, _ in pairs(value) do
      keys[#keys + 1] = key
    end
    table.sort(keys, function(a, b)
      return tostring(a) < tostring(b)
    end)
    local out = {}
    for _, key in ipairs(keys) do
      out[#out + 1] = encode_json(tostring(key)) .. ":" .. encode_json(value[key])
    end
    return "{" .. table.concat(out, ",") .. "}"
  end
  error("unsupported value type for json encoding: " .. value_type)
end

function M.new(opts)
  opts = opts or {}
  local self = setmetatable({}, Client)
  self.address = opts.address or os.getenv("SIKULI_GRPC_ADDR") or "127.0.0.1:50051"
  self.auth_token = opts.auth_token or os.getenv("SIKULI_GRPC_AUTH_TOKEN") or ""
  self.trace_id = opts.trace_id or ""
  if opts.plaintext == nil then
    self.plaintext = true
  else
    self.plaintext = opts.plaintext
  end
  self.service = opts.service or "sikuli.v1.SikuliService"
  self.proto_root = opts.proto_root or "../../proto"
  self.proto_file = opts.proto_file or "sikuli/v1/sikuli.proto"
  self.protoset = opts.protoset or "./generated/sikuli.protoset"
  return self
end

function Client:_command(method_name, payload_json, extra_headers)
  local headers = {}
  for key, value in pairs(extra_headers or {}) do
    headers[key] = value
  end
  if self.auth_token ~= "" and headers["x-api-key"] == nil then
    headers["x-api-key"] = self.auth_token
  end
  if self.trace_id ~= "" and headers["x-trace-id"] == nil then
    headers["x-trace-id"] = self.trace_id
  end

  local parts = { "grpcurl" }
  if self.plaintext then
    parts[#parts + 1] = "-plaintext"
  end
  parts[#parts + 1] = "-emit-defaults"

  if self.protoset ~= "" and file_exists(self.protoset) then
    parts[#parts + 1] = "-protoset"
    parts[#parts + 1] = shell_quote(self.protoset)
  else
    parts[#parts + 1] = "-import-path"
    parts[#parts + 1] = shell_quote(self.proto_root)
    parts[#parts + 1] = "-proto"
    parts[#parts + 1] = shell_quote(self.proto_file)
  end

  for key, value in pairs(headers) do
    parts[#parts + 1] = "-H"
    parts[#parts + 1] = shell_quote(key .. ": " .. value)
  end

  parts[#parts + 1] = "-d"
  parts[#parts + 1] = shell_quote(payload_json)
  parts[#parts + 1] = shell_quote(self.address)
  parts[#parts + 1] = shell_quote(self.service .. "/" .. method_name)

  return table.concat(parts, " ") .. " 2>&1"
end

function Client:invoke(method_name, payload, headers)
  local payload_json = payload
  if type(payload) ~= "string" then
    payload_json = encode_json(payload or {})
  end
  local cmd = self:_command(method_name, payload_json, headers)

  local pipe = io.popen(cmd)
  if pipe == nil then
    return nil, "failed to start grpcurl process"
  end
  local output = pipe:read("*a")
  local ok, reason, code = pipe:close()
  if not ok then
    if output ~= nil and output ~= "" then
      return nil, output
    end
    return nil, string.format("grpcurl failed (%s: %s)", tostring(reason), tostring(code))
  end
  return output, nil
end

function Client:find(request, headers)
  return self:invoke("Find", request, headers)
end

function Client:find_all(request, headers)
  return self:invoke("FindAll", request, headers)
end

function Client:read_text(request, headers)
  return self:invoke("ReadText", request, headers)
end

function Client:find_text(request, headers)
  return self:invoke("FindText", request, headers)
end

function Client:move_mouse(request, headers)
  return self:invoke("MoveMouse", request, headers)
end

function Client:click(request, headers)
  return self:invoke("Click", request, headers)
end

function Client:type_text(request, headers)
  return self:invoke("TypeText", request, headers)
end

function Client:hotkey(request, headers)
  return self:invoke("Hotkey", request, headers)
end

function Client:observe_appear(request, headers)
  return self:invoke("ObserveAppear", request, headers)
end

function Client:observe_vanish(request, headers)
  return self:invoke("ObserveVanish", request, headers)
end

function Client:observe_change(request, headers)
  return self:invoke("ObserveChange", request, headers)
end

function Client:open_app(request, headers)
  return self:invoke("OpenApp", request, headers)
end

function Client:focus_app(request, headers)
  return self:invoke("FocusApp", request, headers)
end

function Client:close_app(request, headers)
  return self:invoke("CloseApp", request, headers)
end

function Client:is_app_running(request, headers)
  return self:invoke("IsAppRunning", request, headers)
end

function Client:list_windows(request, headers)
  return self:invoke("ListWindows", request, headers)
end

M.encode_json = encode_json

return M
