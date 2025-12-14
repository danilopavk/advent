let () =
  Ocaml2025.Files.read_map_fold 0
    (fun batteries -> Ocaml2025.Day3.score batteries 2)
    ( + )
  |> string_of_int |> print_endline

let () =
  Ocaml2025.Files.read_map_fold 0
    (fun batteries -> Ocaml2025.Day3.score batteries 12)
    ( + )
  |> string_of_int |> print_endline
