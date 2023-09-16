// ejercicio 2.3

let n_esimo n lista =
    let indices = [0..(List.length lista - 1)]
    let elementos = List.map (fun i -> if i = n then Some (List.nth lista i) else None) indices
    match List.choose id elementos with
    | [resultado] -> Some resultado
    | _ -> None
    
//Salidas
let lista1 = [1; 2; 3; 4; 5]
printfn "Lista original: %A" lista1
let resultado1 = n_esimo 2 lista1
printfn "n_esimo 2: %A" resultado1 

let resultado2 = n_esimo 3 lista1
printfn "n_esimo 3: %A" resultado2  

let resultado3 = n_esimo 6 lista1
printfn "n_esimo 6: %A" resultado3  
