type cell = { x : int; y : int }

let neighbors cell =
  let x = cell.x in
  let y = cell.y in
  [
    { x = x - 1; y = y - 1 };
    { x = x - 1; y };
    { x = x - 1; y = y + 1 };
    { x; y = y - 1 };
    { x; y = y + 1 };
    { x = x + 1; y = y - 1 };
    { x = x + 1; y };
    { x = x + 1; y = y + 1 };
  ]
