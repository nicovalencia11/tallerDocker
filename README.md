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

# Reto 2: Configuración de un proyecto Java con Docker Compose

**Paso 1:** Crear una imagen Docker para el cliente Java:

`docker build -t myclient ./client`

**Paso 2:** Iniciar el proyecto con Docker Compose:

`docker-compose up -d`

**Paso 3:** Ver los registros del proyecto:

`docker-compose logs`

# Reto 3:
