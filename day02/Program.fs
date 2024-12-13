open System
open System.IO


let lineToIntList (line: string) =
    line.Split [|' '|]
    |> Array.toList
    |> List.map (fun x ->
        match Int32.TryParse x with
        | (true, n) -> n
        | _ -> -1
    )


let rec checkUp entry =
    match entry with
    | prev :: x :: xs -> x > prev && x - prev <= 3 && checkUp (x :: xs)
    | _ :: xs -> checkUp xs
    | [] -> true


let rec checkDown entry = 
    match entry with
    | prev :: x :: xs -> x < prev && prev - x <= 3 && checkDown (x :: xs)
    | _ :: xs -> checkDown xs
    | [] -> true


let readLines path =
    File.ReadLines(path)
    |> Seq.map lineToIntList


let task01 entries = 
    entries
    |> Seq.filter (fun x -> checkUp x || checkDown x)
    |> Seq.length


let main =
    let entries = readLines "input"

    // Task01
    task01 entries
    |> printfn "%d"
