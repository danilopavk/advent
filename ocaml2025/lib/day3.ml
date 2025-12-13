let rec max_second batteries =
match batteries with
| 9 :: _ -> 9
| first :: rest -> max first (max_second rest)
| [] -> 0


let rec max_first batteries =
match batteries with
| _ :: [] -> 0
| 9 :: _ -> 9
| [] -> 0
| first :: rest -> max first (max_first rest)

let rec find_score batteries target =
match batteries with
| first :: [] -> first
| first :: second :: [] -> 10 * first + second
| first :: rest when first = target -> 10 * target + (max_second rest)
| _ :: rest -> find_score rest target
| [] -> 0

let line_score batteries = find_score batteries (max_first batteries)

let score line = Number_util.digits_of_string line |> line_score
