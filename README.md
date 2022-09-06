Soy un aficionado a los juegos de mesa, y a los TCG. Desde hace unos años que vengo soñando con la idea de poder crear un juego
y finalmente me decidí a empezar. Así nace este proyecto, el cual busca facilitar el complicado proceso de creación de cartas y su 
almacenamiento. 

El objetivo de este CLI es el siguiente:

- Crear las cartas en un formato JSON predefinido por mi.
- Almacenar dichos JSON en un archivo .txt llamado por definicion JSON_NombreDeLaCarta.txt, el cual irá dentro de una carpeta llamada 
NombreDeLaCarta, las cuales irán dentro de la carpeta /Cards (es posible redefinir la carpeta a utilizar en el .env)
- El programa recorrerá todas las carpetas dentro de /Cards y, si encuentra algún .txt que contenga JSON en su nombre, leerá el archivo y
almacenará su contenido en la struct definida anteriormente.
- Tras este proceso, se guardará el JSON en una coleccion de Firestore.

De esta forma se podrá hace un "copy - paste" del formato JSON que se quiera utilizar, cambiando así solo los atributos propios de cada carta,
y el programa se encargará de tomar todas las cartas y guardarlas en una base de datos.

Este es el primer objetivo. El segundo es poder ampliar las capacidades de este programa e implementar una nueva función que permita leer
todos los archivos .txt que contengan JSON en su nombre, almacenar su contenido en la struct predefinida y luego, con esa información, reescribir el JSON en una nueva struct. Esto para flexibilizar el cambio de versiones de cartas y ahorrar el trabajo manual.

-------------------

.env

#La carpeta debe estar en la misma carpeta raiz del proyecto, osea, donde esté el archivo main.go
Cards_Folder_Path=NombreDeLaCarpeta
