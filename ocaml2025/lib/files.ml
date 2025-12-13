let read_fold init fold =
  let ic = open_in "input.txt" in

  let rec read_fold_rec acc =
    try read_fold_rec (fold acc (input_line ic)) with
    | End_of_file -> acc
    | e -> raise e
  in

  let result = read_fold_rec init in
  close_in ic;
  result

let read_single func =
  let ic = open_in "input.txt" in
  let result = input_line ic |> func in
  close_in ic;
  result

let read_map_fold init map_fn fold_fn =
  let ic = open_in "input.txt" in

  let rec read_fold_rec acc =
    try read_fold_rec (input_line ic |> map_fn |> (fun item -> fold_fn item acc)) with
    | End_of_file -> acc
    | e -> raise e
  in

  let result = read_fold_rec init in
  close_in ic;
  result
