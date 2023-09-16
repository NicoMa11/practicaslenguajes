//ejercicio 2.1

let rec desplazarIzquierda n lista =
    match n with
    | 0 -> lista
    | _ -> match lista with
           | [] -> []
           | hd::tl -> desplazarIzquierda (n-1) (tl @ [hd])

let rec desplazarDerecha n lista =
    match n with
    | 0 -> lista
    | _ -> match lista with
           | [] -> []
           | _ -> 
               let longitud = List.length lista
               let n' = if n < longitud then n else n % longitud
               let parteDerecha, parteIzquierda = List.splitAt (longitud - n') lista
               parteIzquierda @ parteDerecha

// Ejemplos
let lista1 = [1; 2; 3; 4; 5]
printfn "Lista Orijinal: %A" lista1
let resultado1 = desplazarIzquierda 3 lista1
printfn "Izquierda 3: %A" resultado1  // Salida: [4; 5; 1; 2; 3]

let lista2 = [1; 2; 3; 4; 5]
let resultado2 = desplazarDerecha 4 lista2
printfn "Derecha 4: %A" resultado2  // Salida: [2; 3; 4; 5; 1]

let lista3 = [1; 2; 3; 4; 5]
let resultado3 = desplazarIzquierda 7 lista3
printfn "Izquierda 6: %A" resultado3  // Salida: [3; 4; 5; 6; 7]
