module IntSet = Set.Make(Int)

let rec find_monkeys chars n =
  match chars with
  | [] -> IntSet.empty
  | first :: rest when first = '@' -> IntSet.add n (find_monkeys rest (n + 1))
  | _ :: rest -> find_monkeys rest (n + 1)

let parse_line line =
  String.to_seq line |> List.of_seq |> fun chars -> find_monkeys chars 0

let has_cell grid (cell : Grid.cell) =
  match Files.IntMap.find_opt cell.x grid with
  | Some row -> if IntSet.mem cell.y row then 1 else 0
  | None -> 0
  
let cell_score grid (cell : Grid.cell) =
  List.fold_left (fun acc neighbor -> acc + (has_cell grid neighbor)) 0 (Grid.neighbors cell)
  |> (fun score -> if score > 3 then 0 else 1)

let line_score grid x =
  let rec line_score_rec line =
    match line with
    | [] -> 0
    | first :: rest -> (cell_score grid { x=x; y=first }) + (line_score_rec rest)
  in
  match Files.IntMap.find_opt x grid with
  | Some vals -> IntSet.to_list vals |> line_score_rec
  | None -> 0

let score grid =
  Files.IntMap.bindings grid
  |> List.map (fun (key, _value) -> line_score grid key)
  |> List.fold_left (+) 0
