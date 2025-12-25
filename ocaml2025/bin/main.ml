let () =
  Files.read_fold Collections.IntMap.empty Advent.Day6.parse_line
  |> Advent.Day6.calc_grid 
  |> string_of_int
  |> print_endline

