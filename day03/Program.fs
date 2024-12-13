open System
open System.IO
open System.Text.RegularExpressions


let performMult (mult: string) = 
    mult.Split([|','; '('; ')'|], StringSplitOptions.RemoveEmptyEntries)
    |> fun split -> 
        match split with
        | [|"mul"; v1; v2|] -> 
            match (Int32.TryParse v1, Int32.TryParse v2) with
            | ((true, n1), (true, n2)) -> n1 * n2
            | _ ->  0
        | _ -> 0


let task01 input =
    Regex.Matches (input, "mul\(\d{1,3},\d{1,3}\)")
    |>  Seq.map (fun m ->
        performMult m.Value
    )
    |> Seq.sum


let rec processOperations willDo operations =
    match operations with
    | "do" :: ops -> processOperations true ops
    | "don't" :: ops -> processOperations false ops
    | mult :: ops -> if willDo then performMult mult + processOperations willDo ops else processOperations willDo ops
    | [] -> 0


let task02 input = 
    Regex.Matches (input, "mul\(\d{1,3},\d{1,3}\)|do(n't)?")
    |> Seq.toList
    |> List.map (fun m -> m.Value)
    |> processOperations true


let main =
    let lines = File.ReadAllText("input")

    // Task01
    task01 lines
    |> printfn "Task01: %d"

    task02 lines
    |> printfn "Task02: %d"
