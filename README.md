# lightflow

Creo que debería renombrar el proyecto a Helium, vamos a ver que otro nombre hay mejor.

## Todo:

- convertirlo en un servicio con cron (esto con el k8 se resuelve y no hace falta desarrollar)
- Permitir varios pipe; a,b,c
- Hacer un warn si una variable esta declarada y se sobre escribe. valor antes y despues
- agregar --list [all|token|loop|pipes]
- validar el JSON
	- formato valido, por el lio del *
	- In the loop requiere name field, each name in all structure require value single characteres, allo separation with - or _ not space.
- validate all args on cli
	- Validate name of variable use lowercase
	- fechas validas.
	- validar los args de la cli, si el json esta mal detener. si una pipe, task o looping no existe salir.
- validar las variables por loop, y no me acuerdo en que otro nivel
- forzar que no se puede usar el template en variables.
- se podria especificar en cada loop que pipes usar?
- si el command recibe un exit en el pipe salir del pipe y seguir el loop, solo haceer retry si tiene configurado eso del retry.
