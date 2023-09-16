// ejercicio 2.2
open System

let filtrarPorSubcadena (pal:string) (lista:string list) =
        lista
        |> List.filter (fun str -> System.Text.RegularExpressions.Regex(sprintf "\\b%s\\b"  pal).IsMatch(str))
let lista = ["perro bravo"; "la casa"; "el perro"; "pintando la cerca"; "mi mascota es un perro"]
let subcadena = "perro"
let resultado = filtrarPorSubcadena subcadena lista
printfn "%A" resultado
