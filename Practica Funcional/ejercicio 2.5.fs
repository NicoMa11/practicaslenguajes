// ejercicio 2.5

let grafo = [
    [('i', 0); ('a', 5); ('b', 2)];
    [('a', 5); ('i', 0); ('c', 3); ('d', 7)];
    [('b', 2); ('i', 0); ('c', 1); ('d', 4)];
    [('c', 3); ('a', 1); ('b', 1); ('x', 6)];
    [('d', 7); ('a', 4); ('b', 8); ('f', 3)];
    [('f', 3); ('d', 2)];
    [('x', 6); ('c', 3)]
]

//verifica si un elemento e está en una lista de
//tuplas y se utiliza para comprobar si un nodo es miembro de una ruta.
let miembro e (lista: ('a * 'b) list) =
    lista
    |> List.map (fun (x,_) -> x = e)
    |> List.reduce (fun x y -> x || y)

//toma un nodo y un grafo y devuelve los vecinos del nodo en el grafo.
let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | ((head, peso)::adyacencias)::tail when nodo = head -> adyacencias
    | _::tail -> vecinos nodo tail 

let extender (ruta: _ list) grafo =
    (vecinos (fst ruta.Head) grafo)
    |> List.map (fun (x,peso) -> if (miembro x ruta) then [] else (x,peso)::ruta )
    |> List.filter (fun x -> x <> [])

//realiza una búsqueda en profundidad
let rec prof_aux fin (rutas: _ list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        List.append
            ([List.rev rutas.Head])
            (prof_aux fin rutas.Tail grafo)
    else
        prof_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo
        
let prof ini fin grafo =
    prof_aux fin [[(ini,0)]] grafo
    

//encuentra la ruta más corta entre un nodo inicial y un nodo final 
let ruta_mas_corta ini fin grafo =
    let rec prof_aux fin (rutas: _ list list) grafo =
        if rutas = [] then
            None
        elif fin = (fst rutas.Head.Head) then
            Some(List.rev rutas.Head)
        else
            prof_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo
    match prof_aux fin [[(ini,0)]] grafo with
    | None -> None
    | Some(ruta) ->
        let longitud = ruta |> List.map snd |> List.sum
        Some(ruta, longitud)

let resultado = ruta_mas_corta 'i' 'f' grafo
match resultado with
| None -> printf "No se encontró una ruta entre los nodos especificados"
| Some(ruta, longitud) -> printf "Ruta más corta: %A longitud: %d" ruta longitud

