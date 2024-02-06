Para mi implementación de grafos tome la decisión de usar un grafo ponderado.
Esta decisión es que a base a lo que pude investigar y entender, las aristas contienen un peso, esto considero que es lo que utilizan para saber si una ruta es mejor que otra. Para ello el grafo esta compuesto de nodos, representados por un map[int] que se relacionan con un map[int]int que representa el nodo al cual esta conectado junto con el peso.

El grafo tiene 4 funciones:

- AddNode: Sirve para agregar un nodo al grafo, en las funciones de ejemplo que coloque en el main, no lo utilizo, pero considero que es importante que esta funcion retorne un bool, ya que de esta manera se podría saber si se creo un nuevo nodo o no.
- DeleteNode: Sirve para eliminar un nodo, al igual que la función anterior regresa un bool, pero este es para saber si un nodo fue eliminado o no, algo importante a mencionar es que esta función igual elimina las relaciones (aristas) que comparte con otros nodos, ya que si no existe el nodo, la relación de otro nodo con este nodo recien eliminado no debería existir.
- AddEdge: Crea una arista, recibiendo como parametros el nodo de origen, el nodo de destino y el peso, si el alguno de los nodos (origen o destino) no existe los crea.
- PrintGraph: Imprime el grafo mostrando los nodos y la relación (aristas) que tienen los nodos así como el peso.

Para el desafío adicional del balanceo del grafo, estuve investigando y no encontré información clara de lo que es un balanceo de grafo, algunas páginas se referían a ello como "un algoritmo especifico" sin profundizar en alguno o dar ejemplos de estos, revise papers pero tampoco con mucho éxito y en youtube solo pude encontrar ejercicios de optimización de rutas, pero no se si eso es un sinonimo de balanceo, y no me lo pareció puesto que el resultado final no era un grafo, con más o menos aristas, si no que era un valor fijo.

Para este desafío primero tuve que ir a entender que es un grafo, leí blogs, me vi varios videos y realice apuntes con algunos ejercicios para entender que es un grafo, que no es más que una relación de vertices, conectados con aristas, estos dependiendo del contexto representan algo, a burdo modo es solo información y depende de lo que nosotros deseamos lograr es lo que el grafo representará. Existen varias formas de representarlos como:

- Representación Formal
- Matriz de Adyacencia
- Lista de adyacencia

Creo que lo más difícil de todo fue abstraer lo leído y aprendido a código, ya que en código nosotros no tenemos algo visual para ver el grafo y tenemos que transmitir el conocimiento a una estructura de datos que pueda representar el grafo.
