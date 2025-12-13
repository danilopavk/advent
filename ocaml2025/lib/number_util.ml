let digits_of_string n = String.to_seq n |> List.of_seq |> List.map (fun c -> int_of_char c - int_of_char '0')

let digits_of_int n = string_of_int n |> digits_of_string
