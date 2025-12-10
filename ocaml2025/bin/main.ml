let () =
  Ocaml2025.Files.read_fold Ocaml2025.Day1.next_rotation2
  |> fun result -> result.score
  |> string_of_int
  |> print_endline
