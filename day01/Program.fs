open System
open System.IO

let getColumns filePath = 
    File.ReadLines(filePath)
    |> Seq.fold (fun (col1, col2) line ->
        let split = line.Split([|' '|], StringSplitOptions.RemoveEmptyEntries)
        match split with
        | [|first; second|] -> 
            match Int32.TryParse first, Int32.TryParse second with
            | ((true, n1), (true, n2)) -> (n1 :: col1, n2 :: col2)
            | _ -> (col1, col2)
        | _ -> (col1, col2)
    ) ([], [])
    |> fun (col1, col2) -> (List.rev col1, List.rev col2)


let task01 (col1, col2) =
    (List.sort(col1), List.sort(col2))
    |> fun(col1Sort, col2Sort) -> List.map2 (fun n1 n2 -> abs(n1 - n2)) col1Sort col2Sort
    |> List.sum


let task02(col1, col2) =
    col1
    |> List.map (fun x -> 
        col2 
        |> List.filter (fun y -> y = x)
        |> List.length
        |> (*) x
    )
    |> List.sum


let main = 
    let (col1, col2) = getColumns "input"

    // Task01
    task01 (col1, col2)
    |> printfn "%d"

    // Task02
    task02 (col1, col2)
    |> printfn "%d"
