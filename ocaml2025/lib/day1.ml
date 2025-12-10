type state_score = {
    state: int;
    score: int;
}

let rotate state amount direction = (if direction then state + amount else state - amount)
    |> (fun x -> x mod 100) 
    |> (fun y -> if y >= 0 then y else 100 + y)

let amount input = String.sub input 1 (String.length input - 1) |> int_of_string
let direction input = input.[0] == 'R'

let next_rotation state_score input = 
    let new_state = rotate state_score.state (amount input) (direction input) in
    let new_score = state_score.score + if new_state = 0 then 1 else 0 in
    { state = new_state; score = new_score }
    
