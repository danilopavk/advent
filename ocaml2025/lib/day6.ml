type column = { operation: string; elements: int list}

let update_column elem column =
  match elem with
  | "*" -> { operation = "*" ; elements = column.elements }
  | "+" -> { operation = "+" ; elements = column.elements }
  | i -> { operation = column.operation ; elements = (int_of_string i) :: column.elements }

let update_in_grid elem index grid =
  Collections.IntMap.update index (fun current_val -> match current_val with
  | Some column -> Some (update_column elem column)
  | None -> Some (update_column elem {elements = []; operation = "+"})) grid

let rec update_grid elements index grid =
  match elements with
  | [] -> grid
  | "" :: rest -> update_grid rest index grid
  | first :: rest -> update_grid rest (index + 1) (update_in_grid first index grid)

let parse_line grid line = String.split_on_char ' ' line |> fun elements -> update_grid elements 0 grid

let calc_column column = List.fold_left (if column.operation = "*" then ( * ) else (+)) (if column.operation = "*" then 1 else 0) column.elements

let calc_grid grid = 
  let rec calc_grid_rec grid index =
    (match Collections.IntMap.find_opt index grid with
    | Some column -> (calc_column column) + (calc_grid_rec grid (index + 1))
    | None -> 0) in
  calc_grid_rec grid 0
