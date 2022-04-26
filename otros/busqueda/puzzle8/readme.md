# 8 Puzzle deslizante

| | |
|-|-|
| Estado | En progreso |
| Pruebas unitarias | Si |
| Cobertura | 86.2% |

## Consigna

Este problema fue tomado del libro _Inteligencia Artificial. Un enfoque moderno_ por Peter Norvig y Stuart J. Russell. En el capítulo 3, ejercicio 3.10.

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
name                                              old time/op  new time/op   delta
CopySolver/Solving_Puzzle_with_5_movements-8      8.83µs ± 0%   7.58µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_57_movements-8     40.6ms ± 0%  181.3ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_15_movements-8      116µs ± 0%    338µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_39_movements-8     57.9ms ± 0%   30.7ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_13_movements-8      229µs ± 0%    190µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_14_movements-8      374µs ± 0%    129µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_87_movements-8      246ms ± 0%    177ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_43_movements-8     1.49ms ± 0%  13.21ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_77_movements-8      532ms ± 0%    326ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_39_movements#01-8  67.2ms ± 0%    0.7ms ± 0%     ~     (p=1.000 n=1+1)
[Geo mean]                                        4.40ms        3.36ms       -23.50%
```