# go-ecommerce



# Decisiones tomadas
- Se ha elegido PostgreSQL como base de datos por su robustez, escalabilida y para el manejo eficiente de datos relacionales. 

- Se ha optado por utilizar SQLC para generar código Go a partir de consultas SQL. SQLC permite escribir consultas SQL de manera directa y generar código Go eficiente y seguro, lo que facilita la interacción con la base de datos sin sacrificar el rendimiento.

- Se ha decidido implementar el patrón Repository para abstraer la lógica de acceso a datos. Este patrón permite separar la lógica de negocio de la lógica de acceso a datos, lo que mejora la mantenibilidad y escalabilidad del código. Al utilizar un repositorio, podemos cambiar fácilmente la implementación de acceso a datos sin afectar el resto de la aplicación, lo que facilita futuras migraciones o cambios en la base de datos.

- Para las peticiones HTTP, se ha elegido usar net/http, la biblioteca estándar de Go. Esta decisión se basa en la simplicidad y eficiencia que ofrece net/http para manejar solicitudes HTTP. Al utilizar la biblioteca estándar, evitamos dependencias adicionales y mantenemos el proyecto ligero.

- Para el middleware se ha optado por utilizar chi, un router ligero y rápido para Go. Chi ofrece una sintaxis sencilla y un rendimiento excelente, lo que lo hace ideal para manejar rutas y middleware en aplicaciones web.

- Para las migraciones de la base de datos, se ha optado por goose, por su simplicidad y eficacia para gestionar cambios en la estructura de la base de datos. Goose permite escribir migraciones en SQL o Go. Al utilizar goose, podemos mantener un control de versiones claro sobre los cambios en la base de datos y facilitar el proceso de despliegue y mantenimiento.