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
    | prev :: x :: xs -> if x > prev && x - prev <= 3 then checkUp (x :: xs) else List.length entry
    | _ :: xs -> checkUp xs
    | [] -> -1


let rec checkDown entry = 
    match entry with
    | prev :: x :: xs -> if x < prev && prev - x <= 3 then checkDown (x :: xs) else List.length entry
    | _ :: xs -> checkDown xs
    | [] -> -1


let readLines path =
    File.ReadLines(path)
    |> Seq.map lineToIntList


let task01 entries = 
    entries
    |> Seq.filter (fun x -> checkUp x = -1 || checkDown x = -1)
    |> Seq.length


let task02 entries =
    entries
    |> Seq.filter (fun entry ->
        match checkUp entry with
        | -1 -> true
        | up -> 
            match checkDown entry with
            | -1 -> true
            | down ->
                let faultUpI = List.length entry - up
                let faultDownI = List.length entry - down

                checkUp (List.removeAt faultUpI entry) = -1 ||
                checkUp (List.removeAt (faultUpI + 1) entry) = -1 ||
                checkDown (List.removeAt faultDownI entry) = -1 ||
                checkDown (List.removeAt (faultDownI + 1) entry) = -1
    )
    |> Seq.length


let main =
    let entries = readLines "input"

    // Task01
    task01 entries
    |> printfn "Task01: %d"

    // Task02
    task02 entries
    |> printfn "Task02: %d"
