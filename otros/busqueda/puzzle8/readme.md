# 8 Puzzle deslizante

| | |
|-|-|
| Estado | En progreso |
| Pruebas unitarias | Si |
| Cobertura | 86.2% |

Este problema fue tomado del libro _Inteligencia Artificial. Un enfoque moderno_ por Peter Norvig y Stuart J. Russell. En el capítulo 3, ejercicio 3.2. La consigna es:

> Implemente dos versiones de la función sucesor para el 8-puzle: uno que genere
todos los sucesores a la vez copiando y editando la estructura de datos del 8-puzle, y otro
que genere un nuevo sucesor cada vez, llamando y modificando el estado padre direc-
tamente (haciendo las modificaciones necesarias). Escriba versiones de la búsqueda
primero en profundidad con profundidad iterativa que use estas funciones y compare sus
rendimientos.

El resultado es que la implementacion que muta el estado s mas eficiente en tiempo y memoria. Tomando CopySolver como la implementacion _old_ y MutationSolver como la _new_, vemos que hay una mejora sustancial dado que aplicar una acción en el estado es menos costoso que clonarlo.

```
name                                              old time/op  new time/op  delta
CopySolver/Solving_Puzzle_with_5_movements-8       129µs ± 0%     8µs ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_57_movements-8      396ms ± 0%   178ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_15_movements-8      151ms ± 0%     0ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_39_movements-8      354ms ± 0%    28ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_13_movements-8     4.06ms ± 0%  0.20ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_14_movements-8     8.11ms ± 0%  0.15ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_87_movements-8      277ms ± 0%   179ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_43_movements-8      435ms ± 0%    15ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_77_movements-8      317ms ± 0%   419ms ± 0%     ~     (p=1.000 n=1+1)
CopySolver/Solving_Puzzle_with_39_movements#01-8   551ms ± 0%     1ms ± 0%     ~     (p=1.000 n=1+1)
[Geo mean]                                        67.3ms        3.6ms       -94.71%
```