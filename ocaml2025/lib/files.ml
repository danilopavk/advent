let read_fold fold =
  let ic = open_in "input.txt" in

  let rec read_fold_rec acc =
    try read_fold_rec (fold acc (input_line ic)) with
    | End_of_file -> acc
    | e -> raise e
  in

  let result = read_fold_rec Day1.{ state = 50; score = 0 } in
  close_in ic;
  result

let read_single func =
  let ic = open_in "input.txt" in
  let result = input_line ic |> func in
  close_in ic;
  result
