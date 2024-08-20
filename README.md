# Generador de Contraseñas

Este es un proyecto de generador de contraseñas seguro desarrollado en Go. Permite a los usuarios generar contraseñas de diferentes longitudes y niveles de dificultad a través de una interfaz web simple y agradable.

## Características

- **Generación de Contraseñas**: Permite generar contraseñas con diferentes niveles de seguridad.
- **Configuración de Longitud**: Los usuarios pueden especificar la longitud de la contraseña, desde 6 caracteres en adelante.
- **Interfaz de Usuario**: Interfaz web visualmente atractiva con un fondo personalizado y un logo.

## Tecnologías Utilizadas

- **Go**: Lenguaje de programación utilizado para el backend.
- **HTML/CSS**: Tecnologías utilizadas para el frontend y el diseño.
- **Math/Rand**: Paquete de Go para la generación de contraseñas aleatorias.

## Estructura del Proyecto

generador_contrasenas/

├── main.go

└── static/

├── index.html

└── styles.css


- `main.go`: Código fuente del servidor Go.
- `static/index.html`: Archivo HTML para la interfaz web.
- `static/styles.css`: Archivo CSS para los estilos de la página.

## Instalación y Ejecución

### Requisitos

- Go 1.18 o superior.

### Pasos para Ejecutar

1. **Clona el Repositorio**

   ```bash
   git clone https://github.com/tu_usuario/generador_contrasenas.git
   cd generador_contrasenas

2. **Ejecuta el Servidor**
  ``` go run main.go```


3. **Accede a la Aplicación

Abre tu navegador web y visita [http://localhost:8080](http://localhost:8080). La interfaz web del generador de contraseñas se mostrará, donde podrás generar contraseñas configurando la longitud y el nivel de dificultad.
