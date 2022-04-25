# 8 Puzzle deslizante

| | |
|-|-|
| Estado | En progreso |
| Pruebas unitarias | Si |
| Cobertura | 86.2% |

## Consigna

Este problema fue tomado del libro _Inteligencia Artificial. Un enfoque moderno_ por Peter Norvig y Stuart J. Russell. En el capítulo 3, ejercicio 3.2.

> Implemente dos versiones de la función sucesor para el 8-puzle: uno que genere
todos los sucesores a la vez copiando y editando la estructura de datos del 8-puzle, y otro
que genere un nuevo sucesor cada vez, llamando y modificando el estado padre direc-
tamente (haciendo las modificaciones necesarias). Escriba versiones de la búsqueda
primero en profundidad con profundidad iterativa que use estas funciones y compare sus
rendimientos.

## Solución

Se implementaron ambos algoritmos mediante una misma interfaz:

* [CopySolver](./copy_solver.go)
* [MutationSolver](./mutation_solver.go)

Los casos de prueba son exactamente iguales para ambas implementaciones, estos son generados en el momento de compilación del programa. Se encontró que la implementacion que muta el estado es mas eficiente en tiempo y memoria. Tomando CopySolver como la implementacion _old_ y MutationSolver como la _new_, vemos que hay una mejora sustancial, dado que aplicar una acción en el estado es menos costoso que clonarlo.

```
name                                             old time/op  new time/op   delta
CopySolver/Solving_Puzzle_with_0_movements-8     8.74µs ± 0%   8.03µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#01-8  41.6ms ± 0%  176.7ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#02-8   113µs ± 0%    267µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#03-8  54.8ms ± 0%   23.9ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#04-8   231µs ± 0%    197µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#05-8   396µs ± 0%    130µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#06-8   241ms ± 0%    162ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#07-8  1.66ms ± 0%  12.93ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#08-8   524ms ± 0%    323ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_0_movements#09-8  64.4ms ± 0%    0.7ms ± 0%     ~     (p=1.000 n=1+1)
[Geo mean]                                       4.41ms        3.20ms       -27.40%
```