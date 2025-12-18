let () =
  Files.read_and_index Advent.Day4.parse_line
  |> Advent.Day4.score
  |> string_of_int
  |> print_endline
