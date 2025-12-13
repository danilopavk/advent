let () =
  Ocaml2025.Files.read_map_fold 0 Ocaml2025.Day3.score (+)
  |> string_of_int |> print_endline
