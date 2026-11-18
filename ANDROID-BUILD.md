# Cloudflared para Android ARM64

Este proyecto añade una compilación nativa para Android `arm64-v8a`.

## DNS en Android

La compilación se realiza con `GOOS=android`, `GOARCH=arm64` y `CGO_ENABLED=1` usando Android NDK. El archivo `cmd/cloudflared/android_resolver.go` selecciona el resolver cgo de Go, que delega la resolución en Bionic/Android en lugar de depender del archivo `resolv.conf` usado por el resolver puro de Go. Así se evita el fallback incorrecto a ` [::1]:53` y se respeta la configuración DNS activa de Android (incluidos DNS privado y VPN), siempre que el entorno de ejecución permita la resolución del sistema.

## Salida

El workflow genera el ejecutable Android ARM64 con el nombre y la ruta que usa la aplicación:

`app/src/main/jniLibs/arm64-v8a/libcloudflared.so`

**Importante:** aunque el nombre termine en `.so`, el archivo generado es un ejecutable Android (ELF `ET_DYN`/PIE), no una biblioteca JNI para cargar con `System.loadLibrary`. Debe invocarse como ejecutable, tal como se hace con el binario actual. No se debe llamar a `System.loadLibrary("cloudflared")`.

## Compilar

En GitHub: Actions → Build Cloudflared Android ARM64 → Run workflow. Descarga el artifact `libcloudflared-android-arm64` y reemplaza el archivo de la aplicación.
