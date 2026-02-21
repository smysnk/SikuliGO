local script_dir = arg[0]:match("(.*/)")
if script_dir == nil then
  script_dir = "./"
end

package.path = package.path .. ";" .. script_dir .. "../?.lua"

local sikuligo = require("sikuligo_client")

local function gray_image(name, rows)
  local pix = {}
  for y = 1, #rows do
    for x = 1, #rows[y] do
      pix[#pix + 1] = rows[y][x]
    end
  end
  return {
    name = name,
    width = #rows[1],
    height = #rows,
    pix = pix
  }
end

local client = sikuligo.new({
  proto_root = script_dir .. "../../proto",
  proto_file = "sikuli/v1/sikuli.proto",
  protoset = script_dir .. "../generated/sikuli.protoset"
})

local source = gray_image("source", {
  { 10, 10, 10, 10, 10, 10, 10, 10 },
  { 10, 0, 255, 10, 10, 10, 10, 10 },
  { 10, 255, 0, 10, 0, 255, 10, 10 },
  { 10, 10, 10, 10, 255, 0, 10, 10 },
  { 10, 10, 10, 10, 10, 10, 10, 10 }
})

local needle = gray_image("needle", {
  { 0, 255 },
  { 255, 0 }
})

local response, err = client:find({
  source = source,
  pattern = {
    image = needle,
    exact = true
  }
})

if err ~= nil then
  io.stderr:write("find failed:\n" .. err .. "\n")
  os.exit(1)
end

print(response)
