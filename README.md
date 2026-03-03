# go-ecommerce



# Decisiones tomadas
- Se ha elegido PostgreSQL como base de datos por su robustez, escalabilida y para el manejo eficiente de datos relacionales. 

- Se ha optado por utilizar SQLC para generar código Go a partir de consultas SQL. SQLC permite escribir consultas SQL de manera directa y generar código Go eficiente y seguro, lo que facilita la interacción con la base de datos sin sacrificar el rendimiento.

- Se ha decidido implementar el patrón Repository para abstraer la lógica de acceso a datos. Este patrón permite separar la lógica de negocio de la lógica de acceso a datos, lo que mejora la mantenibilidad y escalabilidad del código. Al utilizar un repositorio, podemos cambiar fácilmente la implementación de acceso a datos sin afectar el resto de la aplicación, lo que facilita futuras migraciones o cambios en la base de datos.

- Para las peticiones HTTP, se ha elegido usar net/http, la biblioteca estándar de Go. Esta decisión se basa en la simplicidad y eficiencia que ofrece net/http para manejar solicitudes HTTP. Al utilizar la biblioteca estándar, evitamos dependencias adicionales y mantenemos el proyecto ligero.

- Para el middleware se ha optado por utilizar chi, un router ligero y rápido para Go. Chi ofrece una sintaxis sencilla y un rendimiento excelente, lo que lo hace ideal para manejar rutas y middleware en aplicaciones web.

- Para las migraciones de la base de datos, se ha optado por goose, por su simplicidad y eficacia para gestionar cambios en la estructura de la base de datos. Goose permite escribir migraciones en SQL o Go. Al utilizar goose, podemos mantener un control de versiones claro sobre los cambios en la base de datos y facilitar el proceso de despliegue y mantenimiento.

- Slog para los logs estructurados, paquete estandar de Go para logging estructurado. Slog ofrece una forma eficiente y flexible de registrar eventos y errores en la aplicación, lo que facilita el monitoreo y la depuración.

# Como crear un nuevo endpoint

1. En la carpeta correspondiente al recurso (por ejemplo, `products`) crear un nuevo archivo `handlers.go` y `service.go` si no existen.

2. En el archivo `handlers.go` crear un nuevo handler para el endpoint, por ejemplo, `CreateProductHandler`. Este handler se encargará de recibir la petición HTTP, validar los datos, llamar al servicio correspondiente y enviar la respuesta al cliente.

3. Consulta sql en el archivo `queries.sql` para crear las consultas necesarias para el nuevo endpoint. Por ejemplo, si estamos creando un endpoint para crear un producto, podríamos agregar una consulta SQL para insertar un nuevo producto en la base de datos.

4. Archivo types.go para definir la estructura de los datos que se van a recibir en la petición HTTP. Por ejemplo, podríamos definir una estructura `CreateProductParams` que contenga los campos necesarios para crear un nuevo producto. Y la interfaz del servicio `ProductService` con el método `CreateProduct(ctx context.Context, params CreateProductParams) (repo.Product, error)`.

5. En el archivo `service.go` implementar la lógica de negocio para el nuevo endpoint. Por ejemplo, en el método `CreateProduct`, podríamos validar los datos recibidos como el nombre, precio y cantidad, llamar al repositorio para insertar el nuevo producto en la base de datos y manejar cualquier error que pueda ocurrir durante este proceso.

6. En el archivo `handlers.go` donde recibimos la petición HTTP, validar que el cuerpo de la petición, sea json válido y que tenga la estructura correcta CreateProductParams, luego llamar al servicio `CreateProduct` con los datos validados y manejar la respuesta. Si el producto se crea correctamente, enviar una respuesta con el producto creado. Si ocurre un error, enviar una respuesta con el error correspondiente.

7. Finalmente, agregar la ruta correspondiente en el archivo `api.go` para que el nuevo endpoint esté disponible. Por ejemplo, podríamos agregar `r.Post("/products", productHandler.CreateProduct)` para que el endpoint de creación de productos esté accesible a través de una petición POST a `/products`.


## TODOs
- Manejar correctamente todos los errores htpp en los endpoints, por ejemplo, si el producto no existe, devolver un error 404, si hay un error en la validación de los datos, devolver un error 400, etc.

- Migrar la conexión a la base de datos a un pool de conexiones para mejorar el rendimiento y la escalabilidad de la aplicación. ( *pgx.Conn a pgxpool.Pool* )

- Implementar slog para los logs estructurados en toda la aplicación, para mejorar el monitoreo y la depuración, no usar log.Println, inyectar el logger en la app para que service y handlers puedan usarlo. 
    ```go
    type application struct {
        config config
        db     *pgx.Conn
        logger *slog.Logger //nuevo, inyectar el logger en la aplicación
    }
    ```
- Revisar logs que puedan estar retornando información sensible, como contraseñas, tokens, etc. y asegurarse de que no se estén registrando en los logs. ( cfg.db.dsn )

- Validaciones más robustas tal vez usar validator
    ```go
    import "github.com/go-playground/validator/v10"

    var validate = validator.New()

    type CreateProductParams struct {
        Name         string `json:"name" validate:"required,max=255"`
        PriceInCents int32  `json:"price_in_cents" validate:"gt=0"`
        Quantity     int32  `json:"quantity" validate:"gte=0"`
    }

    // En service
    if err := validate.Struct(tempProduct); err != nil {
        return repo.Product{}, fmt.Errorf("validation error: %w", err)
    }
    ```
- Endpoint traer orden por id.