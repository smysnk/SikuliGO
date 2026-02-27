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

local _, err = client:move_mouse({
  x = 200,
  y = 180,
  opts = { delay_millis = 30 }
})
if err ~= nil then
  io.stderr:write("move_mouse failed:\n" .. err .. "\n")
  os.exit(1)
end

_, err = client:click({
  x = 200,
  y = 180,
  opts = { button = "left", delay_millis = 20 }
})
if err ~= nil then
  io.stderr:write("click failed:\n" .. err .. "\n")
  os.exit(1)
end

_, err = client:type_text({
  text = "hello from lua grpc",
  opts = { delay_millis = 15 }
})
if err ~= nil then
  io.stderr:write("type_text failed:\n" .. err .. "\n")
  os.exit(1)
end

_, err = client:hotkey({
  keys = { "cmd", "a" }
})
if err ~= nil then
  io.stderr:write("hotkey failed:\n" .. err .. "\n")
  os.exit(1)
end

print("input actions sent")
