let () =
  Ocaml2025.Files.read_single Ocaml2025.Day2.sum_invalids
  |> string_of_int |> print_endline
