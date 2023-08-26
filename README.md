# desarrollo-aplicaciones-empresariales

## Backend
#### Prerequisitos
- Tener instalado go. Se puede validar con el comando 
 ```bash
go version
```
### Ejecutar proyecto localmente
1. Ingresar en el directorio del aplicativo. Si nos encontramos en el directorio raiz del repositorio podemos ejecutar el comando para ingresar al aplicativo 
```bash
cd back
```
2. Instalar la base de datos mediante docker, es importante verificar que nos encontremos en el directorio raiz.
```bash
docker compose up
```
3. Levantar el aplicativo, es importante verificar que nos encontremos en el directorio back.
```bash
go run main.go
```

## Frontend

#### Prerequisitos
- Tener instalado node js `version 18.17.1`. Se puede validar con el comando 
 ```bash
node --version
```
**Importante** Se recomienda tener instalado nvm para poder cambiar entre versiones
mas info [NVM documentacion](https://github.com/nvm-sh/nvm). Para cualquier sistema operativo

### Ejecutar proyecto localmente
1. Ingresar en el directorio del aplicativo. Si nos encontramos en el directorio raiz del repositorio podemos ejecutar el comando para ingresar al aplicativo 
```bash
cd front/to-do-list
```
2. Instalar las dependencias, es importante verificar que nos encontremos en el directorio raiz; para ello validar que nos encontremos en el directorio del archivo `package.json`
```bash
yarn install
```
3. Levantar el aplicativo
```bash
yarn dev
```

