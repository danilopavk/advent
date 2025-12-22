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
    try
      read_fold_rec (input_line ic |> map_fn |> fun item -> fold_fn item acc)
    with
    | End_of_file -> acc
    | e -> raise e
  in

  let result = read_fold_rec init in
  close_in ic;
  result

let read_and_index map_fn =
  let ic = open_in "input.txt" in

  let rec parse_to_grid grid index =
    try
      let grid_value = input_line ic |> map_fn in
      parse_to_grid (Collections.IntMap.add index grid_value grid) (index + 1)
    with
    | End_of_file -> grid
    | e -> raise e
  in
  let result = parse_to_grid Collections.IntMap.empty 0 in
  close_in ic;
  result
