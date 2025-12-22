type range = { from : int; until : int }
type ingredients = { fresh : int; ranges : range list }
type ingredients_sorted = { ranges : range list }

let init_ingredients = { fresh = 0; ranges = [] }
let init_ingredients_sorted = {ranges = []}

let add_range ingredients range =
  { fresh = ingredients.fresh; ranges = range :: ingredients.ranges }

let parse_range line =
  String.split_on_char '-' line |> fun splitted ->
  match splitted with
  | [ first; second ] ->
      { from = int_of_string first; until = int_of_string second }
  | _ -> raise (Invalid_argument line)

let rec is_in_range ranges ingredient =
  match ranges with
  | [] -> false
  | { from; until } :: _ when ingredient >= from && ingredient <= until -> true
  | _ :: rest -> is_in_range rest ingredient

let process_ingredient (ingredients : ingredients) ingredient =
  if is_in_range ingredients.ranges ingredient then
    { fresh = ingredients.fresh + 1; ranges = ingredients.ranges }
  else ingredients

let parse_line ingredients line =
  match line with
  | "" -> ingredients
  | line when String.contains line '-' ->
      parse_range line |> add_range ingredients
  | line -> int_of_string line |> process_ingredient ingredients

let add_range_sorted (ingredients : ingredients_sorted) range =
  let rec add_range_to_ranges ranges range =
    match ranges with
    | [] -> range :: []
    | first :: rest when first.from < range.from ->
        first :: add_range_to_ranges rest range
    | first :: rest -> range :: add_range_to_ranges rest first
  in
  { ranges = add_range_to_ranges ingredients.ranges range }

let parse_line_sorted ingredients line =
  match line with
  | line when String.contains line '-' -> parse_range line |> add_range_sorted ingredients
  | _ -> ingredients

let calculate (ingredients : ingredients_sorted) =
  let rec calculate ranges least most =
    match ranges with
    | [] -> most - least + 1
    | first :: rest when first.from > most -> most - least + 1 + (calculate rest first.from first.until)
    | first :: rest -> calculate rest (min first.from least) (max first.until most)
  in
  match ingredients.ranges with
  | [] -> 0
  | first :: rest -> calculate rest first.from first.until
