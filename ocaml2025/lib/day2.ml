let rec parts l size =
  match l with
  | [] -> []
  | _ -> List.take size l :: parts (List.drop size l) size

let range until =
  let rec aux i acc =
    if i > until then acc
    else aux (i + 1) (i :: acc)
  in
  aux 1 []

let invalid n partition_size =
  let digits = Number_util.digits_of_int n in
  if List.length digits mod partition_size > 0 then false
  else
    match parts digits partition_size with
    | first :: second :: rest ->
        first = second && List.for_all (fun partition -> partition = first) rest
    | _ -> false

let invalid_all n =
  let partition_sizes = range (string_of_int n |> String.length) in
  List.exists (fun partition_size -> invalid n partition_size) partition_sizes

let rec range_score from until invalid_func =
  if from > until then 0
  else (if invalid_func from then from else 0) + range_score (from + 1) until invalid_func

let sum_invalids input part =
  let invalid_func = if part = 1 then
    (fun n ->
      let len = string_of_int n |> String.length in
      if len mod 2 <> 0 then false  (* skip odd-length numbers *)
      else invalid n (len / 2))
  else invalid_all in
  String.split_on_char ',' input |> fun ranges ->
  List.map (fun range -> String.split_on_char '-' range) ranges |> fun ranges ->
  List.map (fun range -> List.map int_of_string range) ranges |> fun ranges ->
  List.map
    (fun range ->
      match range with [ first; second ] -> range_score first second invalid_func | _ -> 0)
    ranges
  |> fun scores -> List.fold_left ( + ) 0 scores
