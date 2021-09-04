# Mis notas

Llamemos flow a el conjunto de tasks y pipes.

Los pipes ejecutan plugins. Necesitamos un registro de plugins.

Hay un proceso principal que lo llamamos core, que es el que se encarga de ejecutar los plugins en el orden que estan definidos en el yaml.
Los tasks repiten los pipes, pero con nuevas variables.

El core y tasks es agnostico al loop, el loop repite el manifiesto del core y de los taks.

- core/flow
- core/flow/plugins
- core/loop

La ruta `core/flow/plugins` es diferente a `plugins/` pq la primera es la implementación del plugin dentro del flow, y la segunda es el plugin puro, encapsulado, sin relacion con el resto.

El generate example esta roto.

Dentro de la config de Task, se pone un struct para los loops, y se da el mismo tratamiento que para los plugins.
Es que los de tipo loops puede usarse para un pipe como para todos los pipes, si es para todos los pipes deben estar en tasks. Entonces lo del directorio loop dentro de core no hace falta.

```yaml
---
tasks:
  - name: Today
  	loops:
  		autoincrement:
  			type: date
  		autodecrement:
  			type: number
  			min: 1
  			max: 10
  			offset: 2
```

Si task como core llaman a plugins,
es una fn que se le pase a plugin.?
Si al plugin se le pasa los datos del plugin y las variables?
Cada vez que se hace set de una variable se hace render.
Variables lo meto dentro de core.?
Se mantiene el reset de variables como regla entre task y task, un task es independiente a otro task, no comparten variables entre ellas, solo se comparte variables entre pipes.

/*
 * Loop: Son los ciclos que repiten N veces los pipes
 * pero con diferentes variables y/o sentido del loop.
 * Por ejemplo: Chunk y/o Increment.
 * Estos valores se pasan por argumentos.
 * Y si los pasamos por el JSON?
 */
