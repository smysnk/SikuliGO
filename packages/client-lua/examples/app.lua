local script_dir = arg[0]:match("(.*/)")
if script_dir == nil then
  script_dir = "./"
end

package.path = package.path .. ";" .. script_dir .. "../?.lua"

local sikuligo = require("sikuligo_client")
local client = sikuligo.new({
  proto_root = script_dir .. "../../proto",
  proto_file = "sikuli/v1/sikuli.proto",
  protoset = script_dir .. "../generated/sikuli.protoset"
})

local app_name = os.getenv("SIKULI_APP_NAME") or "Calculator"

local _, err = client:open_app({
  name = app_name,
  args = {}
})
if err ~= nil then
  io.stderr:write("open_app failed:\n" .. err .. "\n")
  os.exit(1)
end

local running, running_err = client:is_app_running({ name = app_name })
if running_err ~= nil then
  io.stderr:write("is_app_running failed:\n" .. running_err .. "\n")
  os.exit(1)
end
print("is_app_running response:")
print(running)

local windows, windows_err = client:list_windows({ name = app_name })
if windows_err ~= nil then
  io.stderr:write("list_windows failed:\n" .. windows_err .. "\n")
  os.exit(1)
end
print("list_windows response:")
print(windows)

_, err = client:focus_app({ name = app_name })
if err ~= nil then
  io.stderr:write("focus_app failed:\n" .. err .. "\n")
  os.exit(1)
end

_, err = client:close_app({ name = app_name })
if err ~= nil then
  io.stderr:write("close_app failed:\n" .. err .. "\n")
  os.exit(1)
end

print("app control actions sent")
