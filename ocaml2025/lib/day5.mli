type range = { from : int; until : int }
type ingredients = { fresh : int; ranges : range list }
type ingredients_sorted = { ranges : range list }

val init_ingredients : ingredients
val init_ingredients_sorted : ingredients_sorted
val parse_line : ingredients -> string -> ingredients
val parse_line_sorted : ingredients_sorted -> string -> ingredients_sorted
val calculate : ingredients_sorted -> int
