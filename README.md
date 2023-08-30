# Reto 1: Configuración de un servidor Apache con Docker

**Paso 1:** Descargar la imagen de Docker para el servidor Apache:

`docker pull httpd`

 **Paso 2:** Ejecutar un contenedor de Docker con el servidor Apache:
 
`docker run -dit --name apache -p 8080:80 httpd`

**Paso 3:** Copiar archivos al contenedor Apache:

`docker cp index.html apache:/usr/local/apache2/htdocs/ `

`docker cp bloque2.png apache:/usr/local/apache2/htdocs/`

**Paso 4:** Acceder al servidor Apache en tu navegador:

Abre` localhost:8080 `en tu navegador.

# Reto 2-3: Configuración de un proyecto Java con Docker Compose
# Reto 4-5: Cree un programa que dados los datos básicos de un usuario genere un JWT.

**Paso 1:** Crear una imagen Docker para el cliente Java:

`docker build -t client ./client`

**Paso 2:** Iniciar el proyecto con Docker Compose:

`docker-compose up -d`

**Paso 3:** Ver los registros del cliente proyecto:

`docker start my_client`

`docker logs my_client`

**Paso 4:** Ver los validar token generado con las variables de entorno:

[https://jwt.io/](https://jwt.io/ "https://jwt.io/")
