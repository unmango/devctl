walk(
  if type == "object" and has("properties") then
    .properties |= with_entries(select(.value["$ref"] != "#"))
  else
    .
  end
)
