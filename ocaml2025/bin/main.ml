let () =
  Files.read_fold Advent.Day5.init_ingredients Advent.Day5.parse_line
  |> fun ingredients -> ingredients.fresh |> string_of_int |> print_endline

let () =
  Files.read_fold Advent.Day5.init_ingredients_sorted Advent.Day5.parse_line_sorted
  |> Advent.Day5.calculate |> string_of_int |> print_endline
