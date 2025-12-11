type state_score = { state : int; score : int }

let rotate state amount direction =
  ((if direction then state + amount else state - amount) |> fun x -> x mod 100)
  |> fun y -> if y >= 0 then y else 100 + y

let amount input = String.sub input 1 (String.length input - 1) |> int_of_string
let direction input = input.[0] == 'R'

let score state_score =
  state_score.score + if state_score.state = 0 then 1 else 0

let final_rotation amount direction state =
  ( amount mod 100 |> fun amount_mod_100 ->
    state + if direction then amount_mod_100 else -1 * amount_mod_100 )
  |> fun score_normalized ->
  state != 0 && (score_normalized <= 0 || score_normalized > 99)

let score2 state_score amount direction =
  state_score.score + (amount / 100)
  + if final_rotation amount direction state_score.state then 1 else 0

let next_rotation state_score input =
  {
    state = rotate state_score.state (amount input) (direction input);
    score = score state_score;
  }

let next_rotation2 state_score input =
  let amount = amount input in
  let direction = direction input in
  {
    state = rotate state_score.state amount direction;
    score = score2 state_score amount direction;
  }
