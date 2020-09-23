# lightflow

Creo que debería renombrar el proyecto a Helium, vamos a ver que otro nombre hay mejor.

## Todo:

- convertirlo en un servicio con cron
- Permitir varios pipe; a,b,c
- Hacer un warn si una variable esta declarada y se sobre escribe. valor antes y despues
- Tener el execution time por task y por looping, y mostrar un resumen al final.
- agregar --list [all|token|loop|pipes]
- validar el JSON
	- formato valido, por el lio del *
	- In the loop requiere name field, each name in all structure require value single characteres, allo separation with - or _ not space.
- validate all args on cli
	- Validate name of variable use lowercase
	- fechas validas.
	- validar los args de la cli, si el json esta mal detener. si una pipe, task o looping no existe salir.
- validar las variables por loop, y no me acuerdo en que otro nivel
