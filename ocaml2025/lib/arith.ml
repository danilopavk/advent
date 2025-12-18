let digits_of_string n =
  String.to_seq n |> List.of_seq
  |> List.map (fun c -> int_of_char c - int_of_char '0')

let digits_of_int n = string_of_int n |> digits_of_string

(*
Source - https://stackoverflow.com/a
Posted by gasche, modified by community. See post 'Timeline' for change history
Retrieved 2025-12-14, License - CC BY-SA 3.0
*)
let rec pow a = function
  | 0 -> 1
  | 1 -> a
  | n ->
      let b = pow a (n / 2) in
      b * b * if n mod 2 = 0 then 1 else a
