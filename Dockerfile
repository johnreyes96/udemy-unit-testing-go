FROM golang:1.22

# Define la ruta de trabajo dentro del contenedor
WORKDIR /go/src/app
COPY . .

# No necesitas copiar ninguna aplicación, ya que no tienes ninguna
RUN go build -o mi_aplicacion