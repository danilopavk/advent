let rec max_nth batteries len batteries_needed highest =
  match batteries with
  | first :: _ when batteries_needed == len -> first
  | first :: _ when first == highest -> first
  | first :: rest -> max first (max_nth rest (len - 1) batteries_needed highest)
  | [] -> 0

let calc_for_n n digit_no = Number_util.pow 10 (digit_no - 1) * n

let rec line_score_nth batteries len batteries_needed =
  if batteries_needed = 0 then 0
  else
    let target = max_nth batteries len batteries_needed 9 in
    match batteries with
    | first :: rest when target = first ->
        calc_for_n first batteries_needed
        + line_score_nth rest (len - 1) (batteries_needed - 1)
    | _ :: rest -> line_score_nth rest (len - 1) batteries_needed
    | [] -> 0

let score line batteries_needed =
  let digits = Number_util.digits_of_string line in
  let len = List.length digits in
  line_score_nth digits len batteries_needed
