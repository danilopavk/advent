let invalid n =
  let digits = string_of_int n |> String.to_seq |> List.of_seq in
  if List.length digits mod 2 = 1 then false
  else
    let half_len = List.length digits / 2 in
    let first_half = List.take half_len digits in
    let second_half = List.drop half_len digits in
    List.for_all2 ( == ) first_half second_half

let rec range_score from until =
  if from > until then 0
  else (if invalid from then from else 0) + range_score (from + 1) until

let sum_invalids input =
  String.split_on_char ',' input |> fun ranges ->
  List.map (fun range -> String.split_on_char '-' range) ranges |> fun ranges ->
  List.map (fun range -> List.map int_of_string range) ranges |> fun ranges ->
  List.map
    (fun range ->
      match range with [ first; second ] -> range_score first second | _ -> 0)
    ranges
  |> fun scores -> List.fold_left ( + ) 0 scores
