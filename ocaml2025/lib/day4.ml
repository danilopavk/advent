let rec find_monkeys chars n =
  match chars with
  | [] -> Collections.IntSet.empty
  | first :: rest when first = '@' ->
      Collections.IntSet.add n (find_monkeys rest (n + 1))
  | _ :: rest -> find_monkeys rest (n + 1)

let parse_line line =
  String.to_seq line |> List.of_seq |> fun chars -> find_monkeys chars 0

let has_cell grid (cell : Grid.cell) =
  match Collections.IntMap.find_opt cell.x grid with
  | Some row -> if Collections.IntSet.mem cell.y row then 1 else 0
  | None -> 0

let cell_approachable grid (cell : Grid.cell) =
  List.fold_left
    (fun acc neighbor -> acc + has_cell grid neighbor)
    0 (Grid.neighbors cell)
  |> fun score -> score < 4

let line_score grid x =
  let rec line_score_rec line =
    match line with
    | [] -> []
    | first :: rest when cell_approachable grid { x; y = first } ->
        ({ x; y = first } : Grid.cell) :: line_score_rec rest
    | _ :: rest -> line_score_rec rest
  in
  match Collections.IntMap.find_opt x grid with
  | Some vals -> Collections.IntSet.to_list vals |> line_score_rec
  | None -> []

let find_approachable grid =
  Collections.IntMap.bindings grid
  |> List.map (fun (key, _value) -> line_score grid key)
  |> List.concat

let remove_cell grid (cell : Grid.cell) =
  match Collections.IntMap.find_opt cell.x grid with
  | Some vals -> (
      match Collections.IntSet.remove cell.y vals with
      | new_set when Collections.IntSet.is_empty new_set ->
          Collections.IntMap.remove cell.x grid
      | new_set -> Collections.IntMap.add cell.x new_set grid)
  | None -> grid

let remove_cells grid cells = List.fold_left remove_cell grid cells
let score grid = find_approachable grid |> List.length

let rec score_rec grid =
  let cells_to_remove = find_approachable grid in
  match List.length cells_to_remove with
  | 0 -> 0
  | n -> n + (remove_cells grid cells_to_remove |> score_rec)
